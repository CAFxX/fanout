package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    var v1 int16 = -1
    var v2 float32 = 1.5
    var v3 uint16 = 42
    var v4 int32 = -551
    var v5 uint16 = (v3 | 693)
    if (386 != 913) {
        var v6 float64 = ((0.1 - 1.5) / -2.25)
        v2 = v2
        var v7 int32 = -680
        _ = v6
        _ = v7
    }
    v3 += v5
    var v8 int8 = 1
    v9 := []int{-151, 9223372036854775807, 32767}
    var v10 int = 127
    v5 -= v5
    v9 = append(v9, v10)
    v11 := GenS1{F0: 4294967295, F1: 4294967295, F2: 1}
    if ((v8 ^ v8) > v8) {
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
        var v12 float64 = 0.1
        switch (v10 ^ v10) {
        case 127:
            fmt.Println(v2, v8, v1)
        case -85:
            switch "" {
            case "foo bar":
                switch ("foo bar" + ("" + "")) {
                case "a":
                    var v13 uint32 = 2
                    _ = v13
                }
            }
        }
        _ = v12
    } else {
        v5 -= v3
        var v14 int16 = v1
        v14 += 42
        _ = v14
    }
    var v15 int = (v10 / (((v10 | (v10 & v10))) | 1))
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if ((v1 - (v1 & v1)) != (v1 | int16(v4))) {
                v3 -= (v5 & 2)
            }
            _ = fi17
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
