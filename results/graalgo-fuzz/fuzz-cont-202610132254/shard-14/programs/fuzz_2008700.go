package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 int8, p1 uint64) int8 {
    p0 = p0
    v5 := GenS1{F0: 126, F1: 66}
    if false {
        switch "hello" {
        case "a":
            var v6 float64 = ((float64((float32(1e10) - -2.25)) + 0.1) * ((((1.5 + (3.14159 - 1.5)) * (1.5 * 1e10)) - float64(1.5)) * -2.25))
            _ = v6
        default:
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 uint8 = 39
                v9 := GenS2{F0: 1, F1: fi7, F2: -9223372036854775808}
                switch fi7 {
                case 128:
                    v8 = 147
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    }
    switch -1 {
    case 127:
        switch ("\n\t" + "foo bar") {
        case "a":
            var v10 float32 = 1.5
            _ = v10
        }
    case 42:
        var v11 bool = ((0.0 - (1.5 - (3.14159 - 1.5))) >= ((3.14159 / ((0.1 + ((3.14159 - float64(0.0)) + 0.0)) + 3.14159)) - 1e10))
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 >> ((978) & 63))
}
func genFn1(p0 uint8, p1 string, p2 uint16) int64 {
    var v12 uint32 = 956
    var v13 float32 = (1e10 * ((-2.25 * 3.14159) + 1.5))
    v12 = v12
    v14 := []int{int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 903
}

func main() {
    var v1 uint = 2
    v2 := []uint32{uint32(v1), 19, 492, 65535}
    v3 := make(map[uint32]bool)
    var v4 bool = true
    _ = len(v2)
    if (!((!v4) || v4)) {
        var v15 uint8 = 2
        _ = v15
    } else {
        var v16 uint64 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if true {
            var v18 uint8 = 43
            var v19 int64 = 1
            _ = v18
            _ = v19
        }
        v20 := v3[830]
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint32 = 1
        v4 = (!false)
        v1 = v1
        _ = fi21
        _ = v22
    }
    var v23 bool = ((float32(1e10) <= (-2.25 * 0.1)) && v4)
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 float64 = -2.25
        _ = fi24
        _ = v25
    }
    switch 42 {
    case 0:
        fmt.Println(v1)
    case 33:
        var v26 uint8 = 42
        _ = v26
    }
    fmt.Println(v4, v1, v4)
    var v27 int8 = genFn0(42, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
