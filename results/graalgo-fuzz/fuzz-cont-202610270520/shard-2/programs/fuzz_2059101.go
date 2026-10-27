package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) uint8 {
    p0 = 42
    p0 -= (^(p0))
    _ = p0
    return 2
}

func main() {
    var v1 int64 = -9
    var v2 int32 = 42
    var v3 float64 = 1e10
    v4 := make(map[uint32]uint32)
    var v5 int64 = v1
    v6 := GenS1{F0: 1000, F1: 255, F2: 1}
    var v7 int8 = -78
    fmt.Println(v5)
    switch -834 {
    case 2147483647:
        var v8 int16 = 1000
        _ = v8
    case 1:
        var v9 uint = 932
        _ = v9
    case -638:
        v10 := GenS1{F0: 498, F1: 127, F2: 575}
        _ = v10
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            v1 = (v5 * 9223372036854775807)
            _ = fi11
        }
    }
    if (1 > 641) {
        switch ("\n\t" + ("" + "")) {
        case "\n\t":
            var v12 uint64 = 2
            _ = v12
        case "":
            v13 := []int{32768}
            _ = v13
        case "a":
            switch int(v1) {
            case 127:
                var v14 int16 = 258
                _ = v14
            }
        }
    }
    var v15 float64 = v3
    switch "a" {
    case "":
        var v16 bool = true
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v15)
            var v18 uint8 = 42
            v2 = v2
            _ = fi17
            _ = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint8 = 19
            var v21 bool = ((v3 * -2.25) < (((float64(0.0) + ((v3 - (v3 / (v3 * ((1.5 - 0.0) - v15)))) + v15)) * v3) * ((0.0 / (float64(0.1) / (1.5 - float64(1e10)))) * v15)))
            var v22 int32 = (-(v2))
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    default:
        var v23 int16 = 0
        _ = v23
    }
    v15 = 3.14159
    var v24 float64 = v3
    var v25 int32 = (v2 >> ((284) & 63))
    for fi26 := 0; fi26 < 8; fi26++ {
        v27 := []int{(^(fi26)), fi26}
        v25 += (v25 % ((-1) | 1))
        _ = fi26
        _ = v27
    }
    if (0.1 != 0.1) {
        for fi28 := 0; fi28 < 10; fi28++ {
            if ((v1 << ((2) & 63)) >= (v5 / (((v1 / ((v1) | 1))) | 1))) {
                v2 = 0
            }
            var v29 float64 = v3
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS2{F0: 42}
                v7 -= v7
                var v32 uint32 = 2147483647
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi28
            _ = v29
        }
        var v33 string = "hello"
        _ = v33
    } else {
        v15 += -2.25
        fmt.Println(v5)
        v34 := GenS2{F0: 2147483648}
        _ = v34
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        v25 = v25
        fmt.Println(v3, v5, v24)
        v36 := GenS1{F0: 1, F1: 2, F2: 2}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
