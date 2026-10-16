package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 uint8, p2 int64) string {
    var v6 uint8 = 127
    v7 := GenS1{F0: 41, F1: int8(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("hello" + "a")
}
func genFn1() int16 {
    v8 := []int{9223372036854775807, 32768}
    if (127 > 0) {
        var v9 int8 = 127
        _ = v9
    } else {
        var v10 uint8 = 1
        _ = v10
    }
    var v11 float64 = ((-2.25 + (1e10 + float64(0.0))) / float64((((float32(((3.14159 * 1e10) + 0.0)) - (1.5 - 3.14159)) + (1.5 - float32(1.5))) + 0.0)))
    v11 = -2.25
    v11 = -2.25
    _ = v11
    _ = v8
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint16 = 1
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    v5 := make(map[uint32]int)
    v1 = (602 != 536)
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS2{F0: 410, F1: v4, F2: -1}
        var v14 uint64 = 1
        var v15 int32 = 2147483647
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: 42, F1: 127}
        v18 := GenS3{F0: 31, F1: 42, F2: 0}
        v19 := GenS4{F0: 0, F1: genFn1()}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint8 = 128
    v21 := v5[2147483648]
    v22 := GenS1{F0: 0, F1: 42}
    v23 := []int{}
    v4 -= v4
    fmt.Println(v20, v1)
    v24 := GenS2{F0: 42, F1: int32(v20), F2: -9223372036854775808}
    fmt.Println(v20)
    _ = len(v23)
    var v25 int = 42
    switch ("hello" + "") {
    case "\n\t":
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint64 = 18446744073709551615
            v25 = (-(v25))
            v5[438] = v25
            _ = fi26
            _ = v27
        }
    case "hello":
        for fi28 := 0; fi28 < 12; fi28++ {
            v3[2147483647] = (v25 << ((1) & 63))
            v29 := v5[uint32(v2)]
            _ = fi28
            _ = v29
        }
    default:
        v4 = (v4 | v4)
    }
    v1 = (-105 < -1)
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            v1 = v1
            if (1000 >= 1) {
                v2 -= (v2 ^ (v2 % ((v2) | 1)))
                v21 += (v21 * v21)
            }
            var v31 int64 = -466
            _ = v31
        }
        switch v21 {
        case -1:
            v20 -= v20
        case -558:
            var v32 float64 = 1.5
            _ = v32
        case 703:
            v25 = (v21 * (fi30 & (v25 * v21)))
        }
        _ = fi30
    }
    v33 := GenS5{F0: v4, F1: 4294967295, F2: 4294967295}
    if ((1e10 * (3.14159 / (1e10 + 0.0))) == 1e10) {
        if v1 {
            v23 = append(v23, (v21 + (v25 * (v25 - 0))))
        } else {
            if (!v1) {
                switch "" {
                case "":
                    v4 = (v4 << ((4294967295) & 63))
                case "hello":
                    _ = len(v5)
                }
            }
            v34 := GenS6{F0: -447, F1: -219}
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint32 = 1
            fmt.Println(v1, fi35)
            v21 = (fi35 | (fi35 >> ((960) & 63)))
            _ = fi35
            _ = v36
        }
    } else {
        var v37 float64 = 0.1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
