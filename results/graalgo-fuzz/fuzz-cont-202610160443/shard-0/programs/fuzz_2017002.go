package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint; F1 bool; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0(p0 uint64) uint {
    fmt.Println(p0)
    p0 += (p0 - p0)
    var v6 int64 = 477
    _ = p0
    _ = v6
    return 388
}

func main() {
    var v1 string = (("a" + "") + "\n\t")
    var v2 uint64 = 18446744073709551615
    var v3 int = 32767
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    _ = len(v4)
    v7 := GenS1{F0: 50}
    v8 := GenS2{F0: 4294967295}
    var v9 float64 = (((float64((float32(0.0) * ((3.14159 - 0.0) * float32(0.1)))) + -2.25) + (0.1 + ((((((1.5 * 0.0) / (-2.25 * float64(1e10))) / 1.5) / (0.1 - (1.5 / (0.0 - float64(1e10))))) - (1e10 + 1.5)) - 0.1))) / float64((0.1 - 1e10)))
    v10 := GenS3{F0: -1, F1: v2, F2: 127}
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float32 = 1.5
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v1 = ""
        _ = fi13
    }
    if true {
        v14 := GenS1{F0: 1}
        _ = v14
    }
    var v15 uint = 255
    for fi16 := 0; fi16 < 12; fi16++ {
        v15 = (v15 * v15)
        var v17 uint16 = 42
        v18 := GenS4{F0: 0, F1: false, F2: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 = v15
    v15 += (^(v15))
    switch int(v15) {
    case -9223372036854775808:
        switch -568 {
        case -9223372036854775808:
            v19 := GenS5{F0: (v2 + v2), F1: (v15 / ((65535) | 1)), F2: 2}
            _ = v19
        }
    default:
        v20 := GenS2{F0: (v2 >> (((v15 | v15)) & 63))}
        _ = v20
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
