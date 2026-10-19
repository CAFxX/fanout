package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int8 }

func main() {
    v1 := []uint32{967}
    var v2 float32 = ((-2.25 / (float32(3.14159) / 1.5)) * float32(0.1))
    var v3 uint64 = 4294967295
    var v4 float32 = v2
    v5 := make(map[uint32]bool)
    v6 := []int{0}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 4294967295}
        var v9 int32 = 2147483647
        v10 := GenS2{F0: 0, F1: false, F2: 757}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v2 = ((0.0 - v4) + float32(float64((v4 + float32(3.14159)))))
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v4)
            v13 := GenS3{F0: 774, F1: 589, F2: (v3 | 0)}
            v4 -= v2
            _ = fi12
            _ = v13
        }
        v3 += v3
        _ = fi11
    }
    v14 := GenS4{F0: 893}
    if false {
        var v15 float64 = (((0.0 + ((((0.0 + (0.1 - float64(v2))) - 1.5) / 3.14159) + ((-2.25 - -2.25) + -2.25))) - ((-2.25 / (((3.14159 / 1.0) / (1e10 / 1.5)) - float64(((float32((0.1 / (3.14159 * (float64(v2) + ((-2.25 - 3.14159) - (1e10 - -2.25)))))) * ((0.0 - v4) + v4)) / (v4 * v4))))) * -2.25)) + (-2.25 / 1.5))
        v16 := GenS5{F0: "foo bar"}
        for fi17 := 0; fi17 < 11; fi17++ {
            v2 += float32(float64((v4 - (((v2 + v2) + ((v2 - float32(v15)) - -2.25)) * 1e10))))
            var v18 bool = false
            var v19 int32 = 0
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v15
        _ = v16
    } else {
        fmt.Println(v3, v4)
        var v20 uint = 18446744073709551615
        if true {
            switch int(v3) {
            case -747:
                var v21 int64 = 22
                _ = v21
            case 127:
                v3 = 191
            default:
                v22 := GenS5{F0: (("" + "") + "\n\t")}
                _ = v22
            }
        }
        _ = v20
    }
    switch "foo bar" {
    case "foo bar":
        var v23 int64 = -1
        _ = v23
    case "hello":
        fmt.Println(v3)
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint8 = 183
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 float32 = v4
            if len(v6) > 0 {
                v6[0] = (fi26 * fi26)
            }
            var v28 float64 = 0.0
            var v29 float32 = v2
            v30 := GenS6{F0: true, F1: 0}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi24
        _ = v25
    }
    switch 0 {
    case 251:
        v3 = 42
    case -2147483648:
        var v31 uint16 = 65535
        _ = v31
    }
    switch "\n\t" {
    case "hello":
        v32 := GenS2{F0: -316, F1: (-204 < -774), F2: 155}
        _ = v32
    default:
        v2 -= 1.5
    }
    var v33 uint16 = 32767
    v33 += v33
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
