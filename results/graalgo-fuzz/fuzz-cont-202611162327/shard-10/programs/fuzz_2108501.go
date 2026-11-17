package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int16 = 1000
    var v2 uint = 911
    var v3 uint = v2
    var v4 int8 = -128
    var v5 uint64 = 84
    v1 = 1
    v6 := []int{128, 32768, 730, 488}
    if (!(0.0 != float64(1e10))) {
        if false {
            for fi7 := 0; fi7 < 4; fi7++ {
                v4 = (v4 - 100)
                v5 = 42
                var v8 int16 = (v1 >> (((v2 / ((42) | 1))) & 63))
                _ = fi7
                _ = v8
            }
            var v9 int64 = 42
            v6 = append(v6, 4294967295)
            _ = v9
        }
        fmt.Println(v5, v3)
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 string = "a"
            if false {
                fmt.Println(v11, v1, v2)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 = v3
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int = (fi10 << (((v2 / (((v3 ^ (v3 << (((v2 & 4294967295)) & 63)))) | 1))) & 63))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v4)
    }
    var v14 uint64 = v5
    if (!true) {
        var v15 uint32 = 42
        if (-1 >= -9223372036854775808) {
            v6 = append(v6, 0)
            switch -2147483648 {
            case 119:
                v15 -= (v15 | v15)
            case 32768:
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: (("hello" + "hello") + "\n\t")}
                    v14 += (v5 & v14)
                    v14 = v5
                    _ = fi16
                    _ = v17
                }
            default:
                var v18 uint32 = (^(v15))
                _ = v18
            }
        }
        v2 += (^(v2))
        _ = v15
    } else {
        var v19 int32 = 1
        var v20 int8 = (v4 | int8(v5))
        if false {
            var v21 uint = 42
            _ = v21
        }
        _ = v19
        _ = v20
    }
    var v22 float64 = ((((0.1 + -2.25) / -2.25) / (0.1 * 3.14159)) * (0.0 + 0.1))
    v6 = append(v6, -500)
    var v23 uint = v3
    v2 -= uint(v14)
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    if true {
        var v24 float64 = 0.1
        v25 := GenS2{F0: 2147483647, F1: 950}
        if (!(false || (true || false))) {
            v6 = append(v6, -770)
        } else {
            switch ("hello" + ("foo bar" + "a")) {
            case "a":
                var v26 bool = (!(true && true))
                _ = v26
            case "":
                var v27 float32 = float32(float64(1.5))
                _ = v27
            }
        }
        _ = v24
        _ = v25
    }
    v6 = append(v6, 128)
    var v28 float32 = -2.25
    v29 := GenS2{F0: 976, F1: 65535}
    switch 552 {
    case -1:
        var v30 uint64 = v14
        _ = v30
    case 9223372036854775807:
        var v31 uint32 = 0
        _ = v31
    }
    v32 := GenS2{F0: 1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
