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
        <TextInput :width="110">
          Threads
        </TextInput>
        <TextInput :width="110">
          Port
        </TextInput>
        <TextInput :width="110">
          Timeout
        </TextInput>
      </div>
      <div class="hostlist-file">
        <label><i class="mdi mdi-book-open-variant" /> Dictionary</label>
        <FileInput :width="260" />
      </div>
      <div class="hostlist-file">
        <label><i class="mdi mdi-view-list" /> Hostlist</label>
        <FileInput :width="260" />
      </div>
      <TextInput :width="360">
        Output File
      </TextInput>
    </div>
    <div class="hostlist-information">
      <span class="hostlist-title">Information</span>
      <div class="hostlist-status">
        <label><i class="mdi mdi-account" /> Current Host</label>
        <span>{{ 'root@somehost.local' }}</span>
      </div>
      <div class="hostlist-status">
        <label><i class="mdi mdi-progress-clock" /> Status</label>
        <span>{{ 'Dialling' }}</span>
      </div>
      <ProgressBar :width="360" :progress="73" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import TextInput from "../components/TextInput.vue"
import ToggleSwitch from "../components/ToggleSwitch.vue"
import FileInput from "../components/FileInput.vue"
import ProgressBar from "../components/ProgressBar.vue"

export default Vue.extend({
  name: "Hostlist",
  components: {
    TextInput,
    ToggleSwitch,
    FileInput,
    ProgressBar,
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
      padding: 4px 20px 4px 10px;
      .hostlist-switch {
        display: flex;
        align-items: center;

        width: 360px;
        margin-bottom: 12px;

        .toggle-switch {
          margin-left: auto;
        }
      }
      .hostlist-threads-port-timeout {
        display: flex;
        justify-content: space-between;

        width: 360px;
        margin-bottom: 4px;
      }
      .hostlist-file {
        display: flex;
        align-items: center;

        width: 360px;
        margin-bottom: 12px;

        .file-input {
          margin-left: auto;
        }
      }
    }

    .hostlist-information {
      padding: 4px 10px 4px 20px;
      .hostlist-status {
        display: flex;
        align-items: center;
        justify-content: space-between;

        margin-bottom: 9px;

        // width: 360px;
        label {

        }
        span {
          width: 180px;
          padding: 6px 6px 7px 6px;

          text-align: center;
          font-size: 14px;
          color: $font-color-primary-default;

          background-color: rgba($color: #000000, $alpha: 0.2);
          border-radius: $border-radius-default;
        }
      }
    }
  }
</style>
