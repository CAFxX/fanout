package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
func genFn0(p0 uint32, p1 int8, p2 uint16) bool {
    p0 = (p0 >> ((65535) & 63))
    if (float64(1e10) < float64(0.0)) {
        fmt.Println(p0, p0)
    }
    if (false && (!false)) {
        var v4 int16 = -32768
        _ = v4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 += (p0 * p0)
        _ = fi5
    }
    p0 -= 2147483648
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1() int {
    v6 := GenS1{F0: 2, F1: -711}
    v7 := GenS1{F0: 244, F1: 42}
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            if (true && (!(true || (false && true)))) {
                var v9 uint32 = 866
                switch ("foo bar" + "a") {
                case "foo bar":
                    var v10 uint = 304
                    _ = v10
                }
                v11 := GenS2{F0: 1}
                _ = v11
                _ = v9
            }
            switch fi8 {
            case 9223372036854775807:
                fmt.Println(fi8)
            case 1:
                var v12 uint16 = 772
                _ = v12
            }
            _ = fi8
        }
        if false {
            var v13 int = -2147483648
            _ = v13
        } else {
            var v14 float32 = (-2.25 + (float32((1e10 - (float64(0.0) * 0.1))) - 3.14159))
            v14 += 0.1
            _ = v14
        }
        var v15 bool = (0.0 > 1.5)
        _ = v15
    } else {
        var v16 float32 = 0.1
        var v17 float64 = (1e10 / (1e10 / 3.14159))
        _ = v16
        _ = v17
    }
    if (!false) {
        var v18 float64 = (1.5 * (3.14159 / 1.5))
        _ = v18
    } else {
        v19 := GenS1{F0: 255, F1: -32768}
        var v20 uint8 = 127
        var v21 int32 = 332
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 uint = 0
    _ = v22
    _ = v6
    _ = v7
    return -546
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 1
    v23 := GenS2{F0: 42}
    var v24 string = ""
    var v25 int16 = 42
    _ = len(v1)
    v25 -= -32768
    v26 := GenS3{F0: "a"}
    _ = len(v2)
    v3 = 42
    v27 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
