package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int16, p1 int8, p2 int64) uint32 {
    p1 -= (-(p1))
    var v7 uint8 = 42
    var v8 uint64 = 18446744073709551615
    var v9 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 669
    var v3 int8 = 127
    var v4 uint = 65535
    var v5 uint64 = (v2 << ((v4) & 63))
    var v6 bool = false
    if false {
        v10 := GenS1{F0: ""}
        _ = v10
    }
    var v11 uint64 = v2
    var v12 uint16 = 65535
    v13 := GenS2{F0: v2, F1: genFn0(42, v3, -774), F2: v2}
    var v14 float64 = 1.5
    var v15 uint8 = 188
    var v16 bool = (v6 && (false && false))
    var v17 float64 = (((0.0 + ((v14 * (1e10 / ((v14 - (v14 + float64(3.14159))) - (float64(v1) * v14)))) * 0.0)) + float64(v1)) * (v14 / (float64((((0.0 * 1e10) + 1.5) / v1)) * v14)))
    switch "" {
    case "":
        v18 := []int{9223372036854775807, -9223372036854775808, 128, 1}
        _ = v18
    }
    var v19 string = "hello"
    var v20 int32 = -1
    if (102 < 40) {
        var v21 int = -2147483648
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint64 = 1
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        v20 = v20
        for fi24 := 0; fi24 < 19; fi24++ {
            if ((^(v2)) > v2) {
                var v25 int32 = int32(v4)
                fmt.Println(v11, v14)
                var v26 uint32 = 42
                _ = v25
                _ = v26
            } else {
                v17 = 3.14159
            }
            _ = fi24
        }
        v27 := GenS2{F0: v2, F1: 1, F2: (v11 << ((v4) & 63))}
        _ = v27
    }
    fmt.Println(v20, v20)
    fmt.Println(v20, v15, v15)
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 string = ("" + ((("" + v19) + v19) + v19))
        _ = fi28
        _ = v29
    }
    switch -1 {
    case 0:
        v20 += (v20 + v20)
    case 411:
        var v30 bool = v16
        _ = v30
    default:
        if (128 <= 42) {
            var v31 string = ((v19 + (v19 + v19)) + "\n\t")
            var v32 int64 = -961
            v31 = (((v19 + "a") + v19) + v31)
            _ = v31
            _ = v32
        } else {
            fmt.Println(v17, v16)
            v1 += (((1e10 / v1) * (float32((v14 - v14)) + v1)) * (float32(0.0) / v1))
        }
    }
    switch ("hello" + v19) {
    case "":
        v33 := GenS2{F0: 319, F1: 29, F2: (v2 * 1)}
        _ = v33
    case "a":
        if ((v3 ^ v3) >= 0) {
            var v34 string = (("\n\t" + (("foo bar" + v19) + v19)) + v19)
            v35 := GenS1{F0: (v34 + "")}
            _ = v34
            _ = v35
        } else {
            var v36 bool = (2 == (v12 * v12))
            _ = v36
        }
    case "hello":
        var v37 int32 = 471
        _ = v37
    default:
        switch -966 {
        case 407:
            if v6 {
                var v38 string = "\n\t"
                var v39 float64 = v17
                _ = v38
                _ = v39
            } else {
                var v40 uint64 = (v11 & 1)
                v41 := GenS3{F0: ((v4 - (v4 - (v4 - 364))) < uint(v12))}
                v19 = ((v19 + (v19 + v19)) + (v19 + (((v19 + "foo bar") + "a") + v19)))
                _ = v40
                _ = v41
            }
        }
    }
    v42 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
