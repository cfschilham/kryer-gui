<template>
  <div id="app">
    <Tabs @add="addTab">
      <Tab
        v-for="tab in tabs"
        :key="tab.id"
        :width="tabWidth"
        :active="tab.active"
        @close="closeTabIDView(tab.id)"
        @click="setActiveTabID(tab.id)"
      >
        {{ tab.title }} {{ tab.id }}
      </Tab>
    </Tabs>
    <keep-alive include="Hostlist,Manual">
      <router-view
        :key="$route.params.id"
        :destroyID="destroyID"
        @destroyed="closeTabID"
      />
    </keep-alive>
  </div>
</template>

<script lang="ts">
import Vue from "vue"
import Tab from "./components/Tab.vue"
import Tabs from "./components/Tabs.vue"
import "./assets/css/fonts.css"
import "./assets/css/normalize.css"
import "@mdi/font/css/materialdesignicons.css"

export default Vue.extend({
  name: "App",
  components: {
    Tab,
    Tabs,
  },
  data() {
    return {
      tabs: Array<TabObject>(),
      destroyID: Number(),
    }
  },
  computed: {
    tabWidth(): number {
      if (this.tabs.length < 5) {
        return 192
      }
      return 768 / this.tabs.length
    },
  },
  watch: {
    $route(current, previous): void {
      let currentIndex = this.resolveTabIDIndex(Number(current.params.id))
      let previousIndex = this.resolveTabIDIndex(Number(previous.params.id))

      if (currentIndex == null && previousIndex == null) {
        return
      }
      if (currentIndex == null && previousIndex != null) {
        this.tabs[previousIndex].active = false
        return
      }
      if (currentIndex != null && previousIndex == null) {
        this.tabs[currentIndex].active = true
        return
      }
      if (currentIndex != null && previousIndex != null) { // Not actually needed, just to satisfy typescript compiler...
        this.tabs[previousIndex].active = false
        this.tabs[currentIndex].active = true
      }
    },
  },
  created() {
    this.$router.replace("/")
  },
  methods: {
    addTab(mode: string): void {
      if (mode !== "hostlist" && mode !== "manual") {
        return
      }

      let newTab = new TabObject({
        title: `AutoSSH - ${mode[0].toUpperCase() +
            mode.slice(1, mode.length)}`, // Capitalize the first letter
        mode: mode,
        id: this.generateTabID(),
        active: false,
      })
      this.tabs.push(newTab)
      this.setActiveTabID(newTab.id)
    },
    closeTabIDView(id: number): void {
      let index = this.resolveTabIDIndex(id)
      if (index === null) {
        return
      }

      this.destroyID = id
    },
    closeTabID(id: number): void {
      let index = this.resolveTabIDIndex(id)
      if (index === null) {
        return
      }

      this.tabs.splice(index, 1)

      if (this.tabs.length === 0) {
        this.$router.replace("/")
        return
      }
      if (index === 0) {
        this.setActiveTabID(this.tabs[0].id)
        return
      }
      this.setActiveTabID(this.tabs[index - 1].id)
    },
    setActiveTabID(id: number): void {
      let index = this.resolveTabIDIndex(id)
      if (index === null) {
        return
      }
      let tab = this.tabs[index]
      this.$router.replace(`/${tab.mode}/${tab.id}`)
    },
    resolveTabIDIndex(id: number): number | null {
      for (let i = 0; i < this.tabs.length; i++) {
        if (this.tabs[i].id === id) {
          return i
        }
      }
      return null
    },
    generateTabID(): number {
      let id = Math.floor(Math.random() * 10 ** 8)
      for (let i = 0; i < this.tabs.length; i++) {
        if (this.tabs[i].id === id) {
          return this.generateTabID()
        }
      }
      return id
    },
  },
})

class TabObject {
  constructor(options: object) {
    let tab: TabObject = options as TabObject
    this.title = tab.title
    this.mode = tab.mode
    this.id = tab.id
    this.active = tab.active
  }
    title: string;
    mode: string;
    id: number;
    active: boolean;
}
</script>

<style lang="scss" scoped>
  @import "./assets/css/main.scss";

  #app {
    width: 100vw;
    height: 100vh;

    font-family: $font-family-default;
    color: $font-color-secondary-default;

    background-color: $color-primary-default;
  }
</style>
