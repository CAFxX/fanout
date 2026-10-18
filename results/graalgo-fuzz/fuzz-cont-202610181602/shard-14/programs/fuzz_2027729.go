package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 uint16, p1 int16) int32 {
    var v6 bool = ((42 == 9223372036854775807) || (255 <= 65535))
    var v7 uint32 = 42
    v8 := GenS1{F0: 0, F1: "a"}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int, p1 uint16) uint64 {
    if (4294967295 != 444) {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int32 = genFn0(p1, 42)
            v10 = v10
            _ = fi9
            _ = v10
        }
        p0 -= p0
    } else {
        if true {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 float64 = 0.0
                _ = fi11
                _ = v12
            }
            p0 = p0
        } else {
            v13 := GenS1{F0: 742, F1: "\n\t"}
            for fi14 := 0; fi14 < 1; fi14++ {
                fmt.Println(p0, p1, fi14)
                var v15 float64 = 1e10
                p0 = p0
                _ = fi14
                _ = v15
            }
            var v16 uint64 = 1
            _ = v13
            _ = v16
        }
    }
    var v17 string = "hello"
    if false {
        var v18 uint64 = 1
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v17
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]uint32)
    var v3 uint8 = 115
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 945
    v19 := GenS2{F0: v5, F1: -1, F2: 1000}
    var v20 float64 = 1.5
    v1 = (v1 + "")
    v21 := GenS3{F0: ((v1 + ((v1 + v1) + "hello")) + "a"), F1: -1}
    switch 127 {
    case -850:
        v22 := []int{0, 659, 127, -1}
        _ = v22
    default:
        var v23 uint16 = 994
        _ = v23
    }
    var v24 int8 = 42
    if ((3.14159 == (v20 / float64((((float32(v20) * 1e10) - float32(3.14159)) / 0.1)))) && (114 != 42)) {
        var v25 int8 = v24
        _ = v25
    } else {
        var v26 int64 = 319
        _ = v26
    }
    var v27 uint32 = 42
    fmt.Println(v5, v5, v3)
    for fi28 := 0; fi28 < 6; fi28++ {
        switch 32767 {
        case 0:
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := GenS2{F0: (v5 ^ genFn1(fi29, uint16(v5))), F1: 42, F2: -94}
                var v31 uint8 = (v3 << ((186) & 63))
                _ = fi29
                _ = v30
                _ = v31
            }
        case 32768:
            v32 := []int{(fi28 ^ (fi28 - 127)), fi28}
            _ = v32
        default:
            var v33 float64 = (v20 - (v20 + v20))
            _ = v33
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
