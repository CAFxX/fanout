package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 int, p1 int64) uint16 {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 -= (fi6 + -302)
        _ = fi6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        p1 = -896
        _ = fi7
    }
    var v8 uint16 = 65535
    p0 -= p0
    _ = p0
    _ = p1
    _ = v8
    return (v8 & v8)
}
func genFn1(p0 string, p1 int16, p2 int64) uint8 {
    v9 := GenS1{F0: -437, F1: 2}
    fmt.Println(p1)
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := []int{(fi10 / (((fi10 & -345)) | 1)), int(p1)}
        var v12 float64 = 0.0
        if false {
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(p2, fi10, fi13)
                if (true && true) {
                    var v14 float32 = (1e10 * -2.25)
                    _ = v14
                }
                _ = len(v11)
                _ = fi13
            }
            v12 = v12
            var v15 int8 = 42
            _ = v15
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 235
}

func main() {
    v1 := []string{(("a" + "\n\t") + "foo bar"), "foo bar", "\n\t", "a"}
    var v2 uint8 = 255
    var v3 uint16 = 42
    v4 := []int{int(v2)}
    v5 := []uint32{1, 42}
    v3 += v3
    v16 := GenS2{F0: 466, F1: -24}
    v1 = append(v1, "a")
    var v17 bool = (!false)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v18 := GenS2{F0: (^(v3)), F1: -2147483648}
    v5 = append(v5, 2147483647)
    if true {
        v19 := GenS3{F0: -703, F1: "\n\t"}
        var v20 int64 = -430
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v19
        _ = v20
    } else {
        var v21 float64 = 1.5
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v21
    }
    v22 := GenS2{F0: v3, F1: -9223372036854775808}
    var v23 bool = (v17 || false)
    var v24 float64 = 0.0
    fmt.Println(v23, v24)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v25 := GenS4{F0: -31, F1: 65535}
    var v26 float32 = float32((1e10 + (float64(float32(3.14159)) + (float64(-2.25) * 0.0))))
    v27 := GenS3{F0: 42, F1: (("\n\t" + ("" + "")) + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
