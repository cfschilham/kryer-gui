<template>
  <div
    :style="{ width: containerWidth }"
    class="file-input"
  >
    <input
      :id="'file-input-field'+uid"
      class="file-input-field"
      type="file"
    >
    <label class="file-input-label">{{ label.filename || "No file selected" }} </label>
    <div
      class="btn"
      @click="browse()"
    >
      Browse
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "FileInput",
  props: {
    width: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      label: {
        filename: String(),
      },
    }
  },
  computed: {
    containerWidth(): string {
      return String(this.width) + "px"
    },
    labelWidth(): string {
      return String(this.width - 40) + "px"
    },
  },
  mounted() {
    document.addEventListener("change", event => {
      let field = document.getElementById("file-input-field" + this.uid) as HTMLInputElement
      if (event.target !== field || field == null) {
        return
      }

      // eslint-disable-next-line no-useless-escape
      let filename = field.value.match(/[\/\\]([\w\d\s\.\-\(\)]+)$/)
      if (filename == null || filename[1] == null) {
        return
      }
      this.label.filename = filename[1]
    })
  },
  methods: {
    browse(): void {
      let field = document.getElementById("file-input-field" + this.uid)
      if (field == null) {
        return
      }
      field.click()
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .file-input {
    display: flex;
    justify-content: flex-end;
    align-items: center;

    .file-input-label {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;

      font-size: 14px;
      color: $font-color-primary-default;
    }
    .file-input-field {
      display: none;

      margin-left: auto;
    }
    .btn {
      margin-left: 10px;
    }
  }
</style>
