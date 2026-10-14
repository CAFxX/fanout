package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint64 = 4294967295
    var v3 uint8 = 0
    v4 := []int{9223372036854775807}
    v5 := make(map[int]uint32)
    fmt.Println(v1, v3, v1)
    fmt.Println(v2, v2)
    fmt.Println(v3)
    fmt.Println(v1)
    v6 := GenS1{F0: -1, F1: "\n\t"}
    v5[v1] = 371
    if (!(475 > 718)) {
        v7 := GenS1{F0: 1, F1: (("hello" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t")}
        v2 = v2
        _ = v7
    } else {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                switch int(v2) {
                case 4294967295:
                    v1 = fi8
                case -410:
                    v9 := v5[(^(fi8))]
                    _ = v9
                }
                var v10 float32 = 0.1
                var v11 uint32 = 4294967295
                _ = fi8
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 1, F1: "foo bar"}
            _ = v12
        } else {
            v13 := GenS2{F0: 1, F1: 4294967295, F2: 0}
            _ = v13
        }
        var v14 float32 = 1e10
        _ = v14
    }
    switch (v1 << ((65535) & 63)) {
    case -195:
        v2 += v2
    case 537:
        var v15 float64 = (float64(0.1) + 1.5)
        _ = v15
    case 32768:
        v3 = (v3 - v3)
    default:
        v16 := GenS3{F0: 571}
        _ = v16
    }
    v3 = v3
    if true {
        v2 = (^(v2))
        for fi17 := 0; fi17 < 1; fi17++ {
            switch (fi17 << ((uint(v2)) & 63)) {
            case -1:
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v2, fi18)
                    _ = fi18
                }
            default:
                fmt.Println(v1, v2)
            }
            _ = fi17
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1)
    if (((0.0 + ((float32(float64(((1e10 - ((0.1 / (0.1 / 1.0)) - (1e10 / (3.14159 / ((1e10 * -2.25) / -2.25))))) * 1e10))) + (1.5 / 3.14159)) * 3.14159)) * (3.14159 * ((((0.0 / 1.0) - (float32(((-2.25 - (-2.25 + (((1e10 + 3.14159) / (0.1 * 1.5)) * (-2.25 * 0.0)))) + (1e10 / (0.1 + 0.1)))) + (0.1 * 1.5))) / 1e10) * (3.14159 + 3.14159)))) >= float32((0.0 / 1.5))) {
        var v19 string = ("" + "foo bar")
        v20 := GenS4{F0: 0}
        _ = v19
        _ = v20
    } else {
        v1 -= -621
    }
    var v21 int = (v1 * v1)
    v22 := GenS4{F0: -2147483648}
    v2 += 1
    fmt.Println(v3)
    var v23 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
