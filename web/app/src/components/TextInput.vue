<template>
  <div
    :id="'text-input'+uid"
    :class="{ 'error': container.error, 'disabled': container.disabled }"
    :style="{ width: containerWidth }"
    class="text-input"
  >
    <input
      :id="'text-input-field'+uid"
      v-model="field.model"
      :style="{ width: fieldWidth }"
      class="text-input-field"
      type="text"
      @focus="$emit('focus')"
      @blur="$emit('blur')"
    >
    <label
      class="text-input-label"
      :class="{ 'outside': label.outside }"
      @click="focus()"
    >
      <slot />
    </label>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "TextInput",
  props: {
    width: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    error: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      label: {
        outside: false,
      },
      container: {
        error: false,
        disabled: false,
      },
      field: {
        model: String(),
      },
    }
  },
  computed: {
    containerWidth(): string {
      return String(this.width) + "px"
    },
    fieldWidth(): string {
      return String(this.width - 20) + "px"
    },
  },
  watch: {
    disabled(current: boolean): void {
      this.setDisabled(current)
    },
    error(current: string): void {
      this.setError(current)
    },
  },
  mounted() {
    this.$nextTick(() => {
      this.setDisabled(this.disabled)
      this.setError(this.error)
    })
    this.$watch(
      () => {
        return this.field.model
      },
      (current, previous) => {
        this.label.outside = Boolean(current)
        if (current !== previous) {
          this.$emit("input", current)
        }
      }
    )
  },
  methods: {
    focus(): void {
      let field = document.getElementById("text-input-field" + this.uid)
      if (field == null) {
        return
      }
      field.focus()
    },
    setDisabled(bool: boolean): void {
      let field = document.getElementById("text-input-field" + this.uid)
      if (field == null) {
        return
      }
      this.container.disabled = bool
      bool
        ? field.setAttribute("disabled", "true")
        : field.removeAttribute("disabled")
    },
    setError(str: string): void {
      let container = document.getElementById("text-input" + this.uid)
      if (container == null) {
        return
      }
      container.setAttribute("data-before", str)
      if (str === "") {
        this.container.error = false
        return
      }
      this.container.error = true
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .text-input {
    position: relative;

    margin: 7px 0 13px 0;

    overflow: visible;

    &::before {
      position: absolute;
      bottom: -13px;
      left: 1px;

      height: 11px;
      width: 100%;

      content: attr(data-before);

      font-size: 11px;
      color: transparent;

      transition: $transition-default;
      transition-property: color;
    }
    &.error::before {
      color: $font-color-error-default;
    }

    &.disabled {
      .text-input-field:disabled {
        border: 1px dashed rgba(204, 209, 224, 0.3) !important;
      }
      .text-input-label {
        color: rgba(204, 209, 224, 0.3) !important;
      }
    }

    &.error {
      .text-input-field {
        border: 1px solid $color-error-default !important;
      }
      .text-input-label {
        color: $font-color-error-default !important;
      }
    }

    &.disabled.error {
      .text-input-field:disabled {
        border: 1px dashed rgba(204, 209, 224, 0.3) !important;
      }
      .text-input-label {
        color: rgba(204, 209, 224, 0.3) !important;
      }
      &::before {
        color: transparent;
      }
    }

    .text-input-field {
      padding: 9px;

      border: 1px solid $color-tertiary-default;
      border-radius: $border-radius-default;

      background-color: transparent;

      color: $font-color-primary-default;
      font-size: 14px;

      transition: $transition-default;
      transition-property: border-color;

      &:focus {
        outline: none;
        border-color: $color-accent-default;
      }

      &:disabled {
        border: 1px dashed rgba(204, 209, 224, 0.3);

        cursor: not-allowed;
      }
    }

    .text-input-label {
      display: block;
      position: absolute;
      left: 5px;
      bottom: 50%;

      transform: translateY(50%);

      padding: 0 4px;

      z-index: 1;

      font-size: 14px;
      color: $font-color-secondary-default;
      user-select: none;

      background-color: $color-primary-default;
      border-radius: 7px;

      transition: $transition-default;
      transition-property: bottom, font-size, color;

      &:hover {
        cursor: text;
      }
      &.outside {
        bottom: 100%;

        font-size: 11px;
      }
    }
    .text-input-field:focus ~ .text-input-label {
      bottom: 100%;

      font-size: 11px;
      color: $font-color-accent-default;
    }
  }
</style>
