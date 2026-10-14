package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 string }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0(p0 uint32, p1 int32, p2 int) uint32 {
    switch ((("" + ("foo bar" + "foo bar")) + "hello") + "\n\t") {
    case "a":
        if (true || (42 != p1)) {
            p2 -= p2
        }
    case "hello":
        v5 := []int{p2}
        _ = v5
    default:
        var v6 int64 = 995
        _ = v6
    }
    p1 = p1
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    return (p0 << ((255) & 63))
}
func genFn1(p0 int) int16 {
    var v7 float64 = 0.1
    v7 += 3.14159
    fmt.Println(v7)
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := []int{4294967295}
    var v2 uint8 = 2
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: v2, F1: 65535, F2: "hello"}
        var v10 uint64 = 1
        for fi11 := 0; fi11 < 7; fi11++ {
            v10 = 18446744073709551615
            v12 := GenS2{F0: uint(v2), F1: 842, F2: 2147483647}
            v1 = append(v1, 4294967295)
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = (v2 & (v2 & v2))
    var v13 float64 = (-2.25 * float64(((1.5 / 0.1) * (0.0 + float32(-2.25)))))
    fmt.Println(v13, v2)
    v13 -= ((((v13 + v13) / 3.14159) * 3.14159) - (v13 - 0.1))
    var v14 int8 = -120
    for fi15 := 0; fi15 < 14; fi15++ {
        v2 -= v2
        var v16 uint16 = 65535
        var v17 uint = 728
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        v14 = (v14 - (^(v14)))
        var v18 float32 = -2.25
        switch "hello" {
        case "hello":
            v19 := GenS2{F0: 225, F1: 2, F2: 1000}
            _ = v19
        case "\n\t":
            v20 := GenS1{F0: (v2 % ((128) | 1)), F1: 4294967295, F2: "a"}
            _ = v20
        case "a":
            v21 := GenS1{F0: v2, F1: 18446744073709551615, F2: "\n\t"}
            _ = v21
        }
        _ = v18
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = (v2 + (v2 ^ (v2 & (v2 ^ v2))))
    }
    v14 = 1
    switch -2147483648 {
    case 128:
        var v22 uint32 = 4294967295
        _ = v22
    case -747:
        var v23 uint64 = 4294967295
        _ = v23
    case 42:
        v24 := GenS1{F0: (^(v2)), F1: 65535, F2: "hello"}
        _ = v24
    default:
        for fi25 := 0; fi25 < 7; fi25++ {
            if (-57 > genFn1(int(v14))) {
                var v26 uint32 = 2147483647
                switch "foo bar" {
                case "\n\t":
                    var v27 float32 = float32(((3.14159 + 0.0) / (0.1 + (v13 / v13))))
                    _ = v27
                case "hello":
                    var v28 float32 = float32(3.14159)
                    _ = v28
                }
                _ = v26
            } else {
                v14 -= v14
                v29 := GenS1{F0: 2, F1: 4294967295, F2: "foo bar"}
                _ = v29
            }
            var v30 float64 = (float64(3.14159) - 0.0)
            var v31 int = 264
            _ = fi25
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 string = ""
        v34 := v4[4294967295]
        _ = fi32
        _ = v33
        _ = v34
    }
    switch "" {
    case "\n\t":
        var v35 uint32 = 2
        _ = v35
    case "a":
        var v36 float32 = 1.5
        _ = v36
    default:
        for fi37 := 0; fi37 < 11; fi37++ {
            if (false && false) {
                v2 = v2
                var v38 int16 = 42
                v13 = float64(3.14159)
                _ = v38
            }
            _ = fi37
        }
    }
    v14 = 127
    v13 += v13
    if (!false) {
        v13 += (1e10 / ((3.14159 - v13) + (float64(0.0) * (1.5 / (v13 + v13)))))
    }
    v39 := GenS1{F0: 42, F1: 255, F2: "foo bar"}
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 float32 = 1.5
        _ = fi40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
