import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  // A number specifying the number of VUs to run concurrently.
  vus: 300,
  // A string specifying the total duration of the test run.
  duration: '20s',

}
export default function() {
  const payload = JSON.stringify({
    latitude: 0.0,
    longitude: 0.0,
    nearest: 5
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwidHlwIjoibmdvIiwiZXhwIjoxNzEzNDU4NTEwfQ.LZlUy1zzYsT3WpRzDbxF5Rrtvi5frv4gnMoPed1aDVY'
    },
  };
  http.post("http://localhost:8080/ngo/dashboard", payload, params)
  sleep(0.5);
}
