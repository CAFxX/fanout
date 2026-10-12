package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 int16) int64 {
    if false {
        p0 = (p0 >> ((687) & 63))
        v6 := GenS1{F0: 127}
        _ = v6
    } else {
        if (0.0 <= (3.14159 * 3.14159)) {
            p0 -= (^(p0))
            if (875 != 0) {
                if false {
                    p0 = (p0 | (p0 * p0))
                    p0 = (p0 & 42)
                } else {
                    var v7 uint16 = 410
                    fmt.Println(v7, p0)
                    for fi8 := 0; fi8 < 20; fi8++ {
                        fmt.Println(fi8, fi8)
                        v7 -= (v7 ^ 0)
                        v9 := []int{0}
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                }
                var v10 uint16 = 938
                p0 += p0
                _ = v10
            } else {
                v11 := GenS2{F0: 0, F1: 1}
                switch -139 {
                case -2147483648:
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                _ = v11
            }
            p0 += (-(p0))
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                p0 += (p0 >> ((4294967295) & 63))
                v14 := GenS2{F0: -949, F1: 69}
                _ = fi13
                _ = v14
            }
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
        var v16 int16 = (p0 + 1000)
        switch "\n\t" {
        case "\n\t":
            p0 = v16
        case "foo bar":
            v17 := []int{1}
            _ = v17
        default:
            v18 := GenS2{F0: v16, F1: 100}
            _ = v18
        }
        _ = v16
    }
    if ((!(1 == 0)) && (!(192 == 2))) {
        var v19 uint32 = 2
        fmt.Println(p0)
        _ = v19
    } else {
        switch 32767 {
        case 128:
            v20 := GenS3{F0: 253}
            _ = v20
        }
        switch "" {
        case "":
            v21 := []int{42, 469, 1}
            _ = v21
        }
        p0 += (p0 + p0)
    }
    v22 := []int{909, 127, 4294967295, -1}
    var v23 uint64 = 18446744073709551615
    var v24 int = -119
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return 774
}
func genFn1() string {
    v25 := GenS3{F0: 1}
    var v26 float64 = (0.1 + (float64((3.14159 * -2.25)) * 0.1))
    _ = v25
    _ = v26
    return "a"
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    var v4 int64 = 42
    var v5 float64 = (3.14159 / 0.1)
    delete(v2, v1)
    if false {
        var v27 uint = 255
        _ = v27
    }
    if (78 != 65535) {
        v28 := v2["hello"]
        switch v3 {
        case "a":
            var v29 bool = false
            _ = v29
        case "foo bar":
            var v30 uint8 = 198
            _ = v30
        }
        switch "" {
        case "a":
            var v31 int8 = -75
            _ = v31
        }
        _ = v28
    }
    v2[v3] = 1
    var v32 uint64 = 0
    if (v5 == (0.0 / 1.0)) {
        switch 128 {
        case 1:
            v33 := GenS1{F0: 42}
            _ = v33
        case -331:
            v34 := GenS1{F0: 1}
            _ = v34
        default:
            var v35 uint8 = 189
            _ = v35
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 3; fi37++ {
                v1 = (v1 + "\n\t")
                _ = fi37
            }
            var v38 uint64 = (v32 + v32)
            _ = fi36
            _ = v38
        }
        var v39 float32 = (((float32(((0.1 / 1.0) * ((1e10 + v5) / v5))) * 1e10) / -2.25) + (3.14159 - (1.5 * ((0.0 + float32(v5)) + -2.25))))
        _ = v39
    } else {
        fmt.Println(v3, v4)
    }
    v2[v3] = int(v4)
    v40 := GenS3{F0: 2147483647}
    v41 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
