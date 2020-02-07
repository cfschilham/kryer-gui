<template>
  <div
    class="progress-bar"
    :style="{ width: String(width) + 'px' }"
  >
    <span
      class="progress-bar-percentage"
    >{{ progressPercent }}</span>
    <div
      class="progress-bar-inner"
      :style="{ width: progressPercent }"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "ProgressBar",
  props: {
    width: {
      type: Number,
      required: true,
    },
    progress: {
      type: Number,
      required: true,
    },
  },
  computed: {
    progressPercent(): string {
      if (this.progress > 100) {
        return "100%"
      }
      if (this.progress < 0) {
        return "0%"
      }
      return String(this.progress) + "%"
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .progress-bar {
    position: relative;

    background-color: rgba(204, 209, 224, 0.15);
    border-radius: $border-radius-default;

    overflow: hidden;
    .progress-bar-inner {
      height: 26px;
      width: 60%;

      background-color: $color-accent-default;

      transition: $transition-default;
      transition-duration: 1s;
      transition-property: width;
    }
    .progress-bar-percentage {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -55%);

      color: $font-color-primary-default;
    }
  }
</style>
