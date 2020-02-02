<template>
  <div
    class="tab"
    :class="{ active: active }"
    :style="{ width: tabWidth }"
    @click="$emit('click')"
  >
    <span
      class="tab-title"
      :style="{ width: titleWidth }"
    >
      <slot />
    </span>
    <span
      class="tab-close"
      @click="$emit('close')"
    >
      <i class="mdi mdi-close" />
    </span>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "Tab",
  props: {
    width: {
      type: Number,
      required: true,
    },
    active: {
      type: Boolean,
      required: true,
    },
  },
  computed: {
    tabWidth(): string {
      return String(this.width - 1) + "px"
    },
    titleWidth(): string {
      if (this.active) {
        return String(this.width - 42) + "px"
      }
      return String(this.width - 20) + "px"
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .tab {
    position: relative;

    display: flex;
    align-items: center;

    height: 100%;

    border-right: 1px solid $color-primary-default;

    color: $font-color-primary-default;
    user-select: none;

    &:hover {
      cursor: pointer;
    }

    &::before {
      position: absolute;
      top: 0;
      left: 50%;
      transform: translateX(-50%);

      width: 0;
      height: 2px;

      content: "";

      background-color: rgba(255, 255, 255, 0.15);
    }
    &:hover::before {
      width: 100%;
      transition: $transition-default;
      transition-property: width;
    }

    &.active {
      background-color: $color-primary-default;
      .tab-close {
        position: relative;

        display: block;

        z-index: 0;

        &::before {
          position: absolute;
          top: 50%;
          left: 50%;
          transform: translate(-50%, -50%);

          height: 0;
          width: 0;

          z-index: -1;

          content: "";

          background-color: rgba($color: #000000, $alpha: 0.15);
          border-radius: 13px;

          transition: $transition-default;
          transition-property: height, width;
        }
        &:hover::before {
          width: 24px;
          height: 24px;
        }
      }
      &::before {
        width: 100%;
        background-color: $color-accent-default;
      }
    }

    .tab-title {
      padding-left: 10px;

      font-size: 14px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    .tab-close {
      display: none;

      &:hover {
        cursor: pointer;
      }
      i {
        padding: 7px;
        font-size: 18px;
      }
    }
  }
</style>
