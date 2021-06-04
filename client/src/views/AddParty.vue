<template>
  <div class="container is-max-widescreen">
    <p v-if="$auth.isAuthenticated" class="">
      Welcome, {{ $auth.user.email }}!
    </p>
    <div class="field">
      <label class="label">ชื่อปาร์ตี้</label>
      <div class="control">
        <input class="input" type="text" v-model="partyName" placeholder="" />
      </div>
    </div>

    <div class="field">
      <label class="label">จำนวนคนที่ขาด</label>
      <div class="control">
        <input class="input" v-model="members" type="number" placeholder="" />
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button @click="addEventData()" class="button is-link">Submit</button>
      </div>
      <div class="control">
        <button class="button is-link is-light">Cancel</button>
      </div>
    </div>
  </div>
</template>
<script>
// import EventService
import EventService from "@/services/EventService.js";
export default {
  name: "AddParty",
  data() {
    // initialize the event object
    return {
      event: {},
      partyName: "",
      members: 0,
    };
  },
  methods: {
    async addEventData() {
      let obj = {
        name: this.partyName,
        size: parseInt(this.members),
      };
      const accessToken = await this.$auth.getTokenSilently();

      try {
        await EventService.addParty(obj, accessToken);
        this.$router.push({ path: "/" });
      } catch (error) {
        throw error;
      }
    },
  },
};
</script>
<style scoped>
.container {
  padding: 0 20px;
}
</style>