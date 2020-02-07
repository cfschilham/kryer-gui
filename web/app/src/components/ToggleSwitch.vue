<template>
  <div class="toggle-switch">
    <span
      :class="{ 'active': active }"
      class="toggle-switch-field"
      @click="toggle()"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "ToggleSwitch",
  data() {
    return {
      active: false,
    }
  },
  methods: {
    toggle(): void {
      this.active = !this.active
      this.$emit("input", this.active)
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .toggle-switch {
    display: inline-block;
    height: 22px;

    .toggle-switch-field {
      position: relative;

      display: inline-block;
      width: 50px;
      height: 22px;

      background-color: rgba(204, 209, 224, 0.15);
      border-radius: 11px;

      transition: $transition-default;
      transition-property: background-color;

      &:hover {
        cursor: pointer;
      }

      &::before {
        position: absolute;
        left: 0;
        top: 50%;
        transform: translateY(-50%);

        width: 22px;
        height: 22px;

        content: "";

        border-radius: 11px;
        background-color: rgba(204, 209, 224, 0.8);

        transition: $transition-default;
        transition-property: left;
      }

      &.active {
        background-color: $color-accent-default;
      }

      &.active::before {
        left: 29px;
      }
    }
  }
</style>
