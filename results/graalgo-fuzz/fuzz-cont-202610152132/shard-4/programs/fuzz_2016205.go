package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    var v1 float64 = 3.14159
    var v2 float32 = 0.0
    v3 := []int{}
    var v4 int8 = 23
    var v5 int32 = 214
    v6 := GenS1{F0: 546, F1: 872, F2: 42}
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                fmt.Println(v5, v1)
                fmt.Println(v5, v4)
            } else {
                v5 = int32(fi8)
            }
            _ = fi8
        }
        v9 := GenS1{F0: (fi7 / ((fi7) | 1)), F1: 4294967295, F2: 0}
        _ = fi7
        _ = v9
    }
    v2 = v2
    var v10 int16 = -32768
    if false {
        v11 := GenS1{F0: 42, F1: 0, F2: -32768}
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v5, v4)
            if true {
                fmt.Println(fi12, v1)
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v5 += (v5 * -1)
                v5 = v5
                _ = fi13
            }
            _ = fi12
        }
        _ = v11
    }
    _ = len(v3)
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 float64 = 0.0
            fmt.Println(fi15, v4, fi14)
            switch (fi14 << ((255) & 63)) {
            case 391:
                var v17 uint8 = 107
                _ = v17
            case 127:
                switch "hello" {
                case "hello":
                    var v18 int32 = v5
                    _ = v18
                }
            }
            _ = fi15
            _ = v16
        }
        switch "" {
        case "a":
            _ = len(v3)
        case "":
            v2 = (float32(v1) * (-2.25 + v2))
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
