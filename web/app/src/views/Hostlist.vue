<template>
  <div class="hostlist">
    <div class="hostlist-configuration">
      <span class="hostlist-title">Configuration</span>
      <div class="hostlist-switch">
        <label><i class="mdi mdi-account-search" /> Username is Host</label>
        <ToggleSwitch />
      </div>
      <div class="hostlist-switch">
        <label><i class="mdi mdi-plus-box-multiple" /> Multi-Threaded</label>
        <ToggleSwitch />
      </div>
      <div class="hostlist-threads-port-timeout">
        <TextInput :width="100">
          Threads
        </TextInput>
        <TextInput :width="100">
          Port
        </TextInput>
        <TextInput :width="100">
          Timeout
        </TextInput>
      </div>
      <div class="hostlist-file">
        <label><i class="mdi mdi-book-open-variant" /> Dictionary</label>
        <FileInput :width="220" />
      </div>
      <div class="hostlist-file">
        <label><i class="mdi mdi-view-list" /> Hostlist</label>
        <FileInput :width="220" />
      </div>
      <TextInput :width="320">
        Output File
      </TextInput>
    </div>
    <div class="hostlist-status">
      <span class="hostlist-title">Status</span>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import TextInput from "../components/TextInput.vue"
import ToggleSwitch from "../components/ToggleSwitch.vue"
import FileInput from "../components/FileInput.vue"

export default Vue.extend({
  name: "Hostlist",
  components: {
    TextInput,
    ToggleSwitch,
    FileInput,
  },
  props: {
    destroyID: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      cfg: {
        usrIsHost: Boolean(),
        multiThreaded: Boolean(),
        goroutines: Number(),
        port: Number(),
        timeout: Number(),
        dictPath: String(),
        hostlistPath: String(),
        outputPath: String(),
      },
    }
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

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .hostlist {
    display: grid;
    grid-template-columns: 50% 50%;

    padding: 10px;

    font-size: 14px;

    .hostlist-title {
      display: inline-block;

      margin-bottom: 8px;

      color: $font-color-primary-default;
      text-transform: uppercase;
    }

    .hostlist-configuration {
      .hostlist-switch {
        display: flex;
        align-items: center;

        width: 320px;
        margin-bottom: 12px;

        .toggle-switch {
          margin-left: auto;
        }
      }
      .hostlist-threads-port-timeout {
        display: flex;
        justify-content: space-between;

        width: 320px;
        margin-bottom: 4px;
      }
      .hostlist-file {
        display: flex;
        align-items: center;

        width: 320px;
        margin-bottom: 12px;

        .file-input {
          margin-left: auto;
        }
      }
    }

    .hostlist-status {

    }
  }
</style>
