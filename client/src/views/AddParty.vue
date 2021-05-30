<template>
  <div class="container is-max-widescreen">
    <p v-if="$auth.isAuthenticated" class="">Welcome, {{ $auth.user }}!</p>
    <div class="field">
      <label class="label">ชื่อปาร์ตี้</label>
      <div class="control">
        <input class="input" type="text" placeholder="" />
      </div>
    </div>

    <div class="field">
      <label class="label">จำนวนคนที่ขาด</label>
      <div class="control">
        <input class="input" type="number" placeholder="" />
      </div>
    </div>


    <div class="field is-grouped">
      <div class="control">
        <button class="button is-link">Submit</button>
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
    };
  },
  created() {
    // this.getEventData();
  },
  methods: {
    async getEventData() {
      // Get the access token from the auth wrapper
      const accessToken = await this.$auth.getTokenSilently();

      // Use the eventService to call the getEventSingle method
      EventService.getEventSingle(this.$route.params.id, accessToken).then(
        ((event) => {
          this.$set(this, "event", event);
        }).bind(this)
      );
    },
  },
};
</script>