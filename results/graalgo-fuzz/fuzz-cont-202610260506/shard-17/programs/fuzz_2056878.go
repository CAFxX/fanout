package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int16; F1 uint8 }

func main() {
    var v1 uint = 224
    var v2 float64 = -2.25
    v3 := []uint32{65535, 935, 0, 232}
    v1 = (v1 * (v1 + (v1 + v1)))
    fmt.Println(v2, v2)
    var v4 int16 = -1
    fmt.Println(v4)
    var v5 uint = (v1 << (((v1 ^ 199)) & 63))
    v6 := GenS1{F0: false}
    v7 := GenS2{F0: v5, F1: 128}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 6; fi9++ {
        if false {
            var v10 float64 = (float64((3.14159 * 0.0)) - float64(-2.25))
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int16 = v4
                v13 := GenS3{F0: (v4 / ((-32768) | 1)), F1: v1}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS2{F0: (^(v5)), F1: fi9}
                if false {
                    var v16 uint16 = 1
                    var v17 int64 = 42
                    _ = v16
                    _ = v17
                } else {
                    v2 -= v2
                    switch (fi9 >> ((v5) & 63)) {
                    case 4294967295:
                        v5 = (^(v1))
                    }
                }
                _ = fi14
                _ = v15
            }
            if (!false) {
                var v18 float32 = (3.14159 + -2.25)
                v2 = float64(v18)
                _ = v18
            }
            var v19 int32 = -2147483648
            _ = v19
        }
        v20 := GenS4{F0: (^(v4)), F1: 2}
        switch "hello" {
        case "foo bar":
            v5 += 2
        }
        _ = fi9
        _ = v20
    }
    var v21 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
