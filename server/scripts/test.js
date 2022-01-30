import http from 'k6/http';
import { check} from 'k6';


export const options = {
  duration: '10s',
  vus: 10
};

export default function () {
  // Get random car
  const mockCar = mockCars[Math.floor(Math.random()* mockCars.length)];

  // Update url if server is not on the same host
  const baseUrl = 'http://localhost:8080';

  // Car creation
  let res = http.post(`${baseUrl}/cars`, JSON.stringify(mockCar));
  check(res, {
    'status was 200': (r) => r.status == 200,
    "car properties have been correctly parsed": (r) => {
      const jsonCar = JSON.parse(r.body)
      return jsonCar.brand === mockCar.brand
        && jsonCar.model === mockCar.model
        && jsonCar.horse_power === mockCar.horse_power
    }
  });
  // Car read
  const carId = JSON.parse(res.body).id;
  res = http.get(`${baseUrl}/cars/${carId}`);
  check(res, {
    'status was 200': (r) => r.status == 200,
    "car properties have been correctly parsed": (r) => {
      const jsonCar = JSON.parse(r.body)
      return jsonCar.brand === mockCar.brand
        && jsonCar.model === mockCar.model
        && jsonCar.horse_power === mockCar.horse_power
        && jsonCar.id === carId
    }
  });

  // Car deletion
  res = http.del(`${baseUrl}/cars/${carId}`);
  check(res, {
    'status was 200': (r) => r.status == 200,
    "car properties have been correctly parsed": (r) => {
      const jsonCar = JSON.parse(r.body)
      return jsonCar.brand === mockCar.brand
        && jsonCar.model === mockCar.model
        && jsonCar.horse_power === mockCar.horse_power
        && jsonCar.id === carId
    }
  });

  // Check car has been deleted
  res = http.get(`${baseUrl}/cars/${carId}`);
  check(res, {
    'status was 404': (r) => r.status == 404,
  });
}

const mockCars = [
  {
    brand: "Volkswagen",
    model: "Golf",
    horse_power: 320
  },
  {
    brand: "Volkswagen",
    model: "T-Roc",
    horse_power: 300
  },
  {
    brand: "Peugeot",
    model: "208",
    horse_power: 130
  },
  {
    brand: "Peugeot",
    model: "2008",
    horse_power: 155
  },
  {
    brand: "Dacia",
    model: "Sandero",
    horse_power: 101
  },
  {
    brand: "Reanult",
    model: "Clio",
    horse_power: 140
  },
  {
    brand: "Toyota",
    model: "Yaris",
    horse_power: 125
  },
  {
    brand: "Opel",
    model: "Corsa",
    horse_power: 130
  },
  {
    brand: "Citroën",
    model: "C3",
    horse_power: 110
  },
];