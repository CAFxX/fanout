package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
func genFn0() int {
    v5 := GenS1{F0: ""}
    var v6 int64 = -382
    if (1e10 == (1e10 - (float32(1.5) * (float32(1e10) + -2.25)))) {
        fmt.Println(v6)
    }
    _ = v5
    _ = v6
    return 392
}
func genFn1(p0 uint64, p1 bool, p2 int8) uint16 {
    var v7 int64 = -9223372036854775808
    var v8 uint8 = 0
    var v9 int16 = 1000
    for fi10 := 0; fi10 < 18; fi10++ {
        p2 = p2
        var v11 bool = (68 >= 303)
        fmt.Println(v11, p0, p0)
        _ = fi10
        _ = v11
    }
    fmt.Println(p1, v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 255
    var v3 uint64 = 560
    v4 := make(map[int]int)
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int32 = 42
            _ = fi13
            _ = v14
        }
        if false {
            delete(v4, fi12)
            var v15 int32 = (v1 % ((v1) | 1))
            fmt.Println(v3, v15, v2)
            _ = v15
        } else {
            v1 = v1
        }
        _ = fi12
    }
    v4[-429] = 346
    var v16 uint = 784
    v2 -= (v2 * (v2 | (^(v2))))
    v2 = v2
    switch "hello" {
    case "a":
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS1{F0: ""}
            _ = fi17
            _ = v18
        }
    default:
        v1 -= (v1 * v1)
    }
    v19 := GenS1{F0: "foo bar"}
    v20 := v4[0]
    if false {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int = (v20 - v20)
            _ = fi21
            _ = v22
        }
        if ((((((((3.14159 * float32(3.14159)) + 1.5) * 0.1) * 0.1) - 0.0) - 0.0) == float32(0.1)) && (false || false)) {
            var v23 int16 = 748
            _ = v23
        } else {
            fmt.Println(v20, v2)
            v2 -= v2
            if false {
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 uint32 = 685
                    v3 = (v3 % (((^(v3))) | 1))
                    _ = fi24
                    _ = v25
                }
                fmt.Println(v2)
            }
        }
        v16 = v16
    } else {
        var v26 int32 = v1
        for fi27 := 0; fi27 < 6; fi27++ {
            v1 = v1
            v28 := v4[(^(fi27))]
            fmt.Println(fi27)
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    v3 = 918
    var v29 int = genFn0()
    if ((!false) || false) {
        fmt.Println(v20, v1, v1)
    }
    v2 += 42
    v16 -= v16
    v30 := GenS2{F0: (int32(v20) >= (v1 + 49))}
    var v31 uint16 = genFn1(v3, true, -118)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
