<template>
  <div class="manual" />
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "Manual",
  props: {
    destroyID: {
      type: Number,
      required: true,
    },
  },
  watch: {
    destroyID(current: number): void {
      if (String(current) === this.$vnode.key) {
        this.$emit("destroyed", current)
        this.$destroy()
      }
    },
  },
})
</script>
