package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 bool, p2 uint64) int64 {
    var v5 uint = (p0 << ((p0) & 63))
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{(fi6 * (fi6 ^ fi6)), 209}
        _ = fi6
        _ = v7
    }
    var v8 float32 = float32((float64(1e10) + float64(((0.0 * 3.14159) / 3.14159))))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v9 := GenS1{F0: ("a" + "")}
    var v10 int64 = -1
    if true {
        v10 += v10
        var v11 float64 = 1e10
        v12 := GenS2{F0: 18446744073709551615, F1: 447, F2: -827}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        switch ("\n\t" + "a") {
        case "hello":
            var v14 int32 = 2147483647
            _ = v14
        case "\n\t":
            for fi15 := 0; fi15 < 1; fi15++ {
                fmt.Println(fi15, v10, fi13)
                var v16 float64 = (float64((3.14159 / 1.0)) + -2.25)
                var v17 string = "hello"
                v18 := []int{-67, fi15}
                var v19 float64 = v16
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
        }
        _ = fi13
    }
    delete(v2, "")
    v4[127] = (!false)
    fmt.Println(v10)
    v20 := GenS2{F0: 695, F1: (v10 % ((v10) | 1)), F2: 1}
    var v21 uint = 2
    var v22 float32 = 0.0
    if ((v21 << (((v21 * (^(v21)))) & 63)) == (v21 - (v21 * (v21 & 65535)))) {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint8 = 161
            v25 := GenS1{F0: "hello"}
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint8 = 2
        v3[1] = (true || true)
        _ = v26
    }
    v27 := GenS1{F0: "foo bar"}
    var v28 bool = true
    v29 := GenS2{F0: 42, F1: genFn0((v21 - (^(v21))), (!v28), 2), F2: 375}
    switch "a" {
    case "a":
        fmt.Println(v21, v28)
    case "\n\t":
        v30 := []int{128, 4294967295, 9223372036854775807}
        _ = v30
    case "foo bar":
        if v28 {
            v31 := []int{694, -989, -569, 127}
            var v32 float64 = 3.14159
            _ = v31
            _ = v32
        } else {
            v22 -= 3.14159
            v22 += -2.25
            v28 = (!(!v28))
        }
    default:
        fmt.Println(v21)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
