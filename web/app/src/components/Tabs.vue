<template>
  <div class="tabs">
    <slot />
    <div
      :id="'tabs-add' + uid"
      ref="tabsAdd"
      class="tabs-add"
      @click="dropdown.toggle()"
    >
      <i
        :id="'tabs-add-mdi' + uid"
        class="mdi mdi-plus"
      />
    </div>
    <div
      v-if="dropdown.active"
      :id="'tabs-dropdown' + uid"
      class="tabs-dropdown"
      :style="{ left: dropdownOffsetLeft() }"
    >
      <div
        class="tabs-dropdown-item"
        @click="add('manual')"
      >
        Manual
      </div>
      <div class="separator" />
      <div
        class="tabs-dropdown-item"
        @click="add('hostlist')"
      >
        Hostlist
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue"

export default Vue.extend({
  name: "Tabs",
  data() {
    return {
      dropdown: {
        active: Boolean(),
        toggle(): void {
          this.active = !this.active
        },
      },
    }
  },
  created(): void {
    document.addEventListener("click", event => {
      if (event.target === document.getElementById("tabs-dropdown" + this.uid)) {
        return
      }
      if (event.target === document.getElementById("tabs-add" + this.uid)) {
        return
      }
      if (event.target === document.getElementById("tabs-add-mdi" + this.uid)) {
        return
      }
      this.dropdown.active = false
    })
  },
  methods: {
    add(mode: string): void {
      this.dropdown.active = false
      this.$emit("add", mode);

      (window as any).astilectron.sendMessage("add", () => {})
    },
    dropdownOffsetLeft(): string {
      let tabsAdd = this.$refs.tabsAdd as HTMLElement
      if (tabsAdd.offsetLeft === 0) {
        return "-1px"
      }
      if (tabsAdd.offsetLeft > 730) {
        return "731px"
      }
      return String(tabsAdd.offsetLeft - 1) + "px"
    },
  },
})
</script>

<style lang="scss" scoped>
  @import "../assets/css/main.scss";

  .tabs {
    position: relative;

    display: flex;
    align-items: center;

    height: 32px;

    color: $font-color-primary-default;
    background-color: $color-secondary-default;

    .tabs-add {
      position: relative;

      padding: 0 4px;

      z-index: 0;

      user-select: none;
      font-size: 24px;

      &:hover {
        cursor: pointer;
      }

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
        width: 26px;
        height: 26px;
      }
    }
    .tabs-dropdown {
      position: absolute;
      top: 32px;

      width: 70px;

      z-index: 10;

      border: 1px solid $color-primary-default;

      background-color: $color-secondary-default;

      .tabs-dropdown-item {
        padding: 6px 0;

        font-size: 14px;
        text-align: center;

        &:hover {
          cursor: pointer;

          user-select: none;

          background-color: rgba($color: #000000, $alpha: 0.15);
        }
      }

      .separator {
        width: 100%;
        height: 1px;
        padding: 0;
        margin: auto;

        background-color: $color-primary-default;
      }
    }
  }
</style>
