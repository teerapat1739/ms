import axios from "axios"

export default {
  async getEvents() {
    // let res = await axios.get("http://localhost:8000/events");
    let res = await axios.get("http://localhost:8000/api-party/party");
    return res.data.data || [];
  },
  async getEventSingle(eventId, accessToken) {
    let res = await axios.get("http://localhost:8000/events/" + eventId, {
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    });
    return res.data;
  }
}