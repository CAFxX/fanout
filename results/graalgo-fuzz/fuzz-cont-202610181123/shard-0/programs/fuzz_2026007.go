package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 uint, p2 int64) uint64 {
    if true {
        var v6 uint64 = 4294967295
        var v7 int64 = (p2 * p2)
        p0 += 0
        _ = v6
        _ = v7
    } else {
        var v8 int32 = -2147483648
        var v9 int16 = 42
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}
func genFn1(p0 uint32, p1 bool, p2 uint64) int16 {
    var v10 uint32 = 732
    v11 := GenS1{F0: "hello", F1: 0}
    fmt.Println(p1, v10)
    var v12 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint16 = 0
    var v2 uint16 = 540
    var v3 int32 = -533
    var v4 int16 = -901
    var v5 float32 = (1e10 + ((((3.14159 * ((1.5 * 3.14159) / (float32((-2.25 / -2.25)) - -2.25))) - (float32(1e10) / 1.5)) * 0.1) + 3.14159))
    var v13 uint16 = v1
    v3 += v3
    v14 := GenS2{F0: 2}
    v15 := []int{127}
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v4, v3)
        if false {
            v17 := GenS2{F0: 4294967295}
            _ = v17
        } else {
            if len(v15) > 0 {
                v15[0] = 237
            }
            fmt.Println(v13, v4, v5)
            v18 := GenS3{F0: 2147483647}
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 float64 = float64(float32(1e10))
        _ = fi19
        _ = v20
    }
    var v21 int16 = v4
    var v22 uint32 = 4294967295
    var v23 int = 234
    v13 = v2
    var v24 float32 = float32((1e10 - 0.1))
    if (-168 < (v3 + (v3 | (v3 >> ((18446744073709551615) & 63))))) {
        var v25 int16 = (v4 << ((0) & 63))
        var v26 float64 = (1e10 / -2.25)
        var v27 uint8 = 2
        _ = v25
        _ = v26
        _ = v27
    } else {
        var v28 string = "\n\t"
        var v29 uint = 18446744073709551615
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int32 = (v3 & (v3 & v3))
        var v32 uint = 739
        if false {
            switch "" {
            case "\n\t":
                v15 = append(v15, (^(fi30)))
            case "hello":
                var v33 uint = v32
                _ = v33
            default:
                var v34 int64 = 42
                _ = v34
            }
            v5 -= (3.14159 - ((float32(-2.25) / ((v24 / (v24 * ((v5 - (1e10 / v5)) + 0.1))) * float32(0.0))) * (float32(1.5) + -2.25)))
            var v35 int8 = 100
            _ = v35
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
