import axios from "axios"

export default {
  async AddParty() {
    let res = await axios.get("http://localhost:8000/events");
    return res.data;
  },
//   async getEventSingle(eventId, accessToken) {
//     let res = await axios.get("http://localhost:8000/events/" + eventId, {
//         headers: {
//             Authorization: `Bearer ${accessToken}`
//         }
//     });
//     return res.data;
//   }
}