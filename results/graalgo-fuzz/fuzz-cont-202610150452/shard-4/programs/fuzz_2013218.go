package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0(p0 int32, p1 uint, p2 bool) int16 {
    var v6 float64 = 0.0
    p0 = -152
    v7 := []int{32768, 128, -301, 32768}
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if ((^(p0)) > p0) {
                fmt.Println(fi9)
                if p2 {
                    if (p0 > int32(p1)) {
                        p0 -= (p0 - p0)
                        var v10 int16 = -886
                        _ = v10
                    } else {
                        p1 += 65535
                        var v11 uint8 = 2
                        _ = v11
                    }
                    p1 += (^(p1))
                    switch (fi8 | (fi9 % ((fi9) | 1))) {
                    case 0:
                        v12 := GenS1{F0: 42, F1: fi9}
                        _ = v12
                    case 968:
                        p1 = 255
                    case 2147483647:
                        v7 = append(v7, fi8)
                    }
                    v13 := GenS1{F0: -1, F1: (fi8 ^ fi9)}
                    p1 = p1
                    _ = v13
                }
                p0 = 442
            }
            _ = fi9
        }
        fmt.Println(v6, fi8, p1)
        var v14 uint8 = 255
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 float64 = (float64((-2.25 * -2.25)) * (1e10 + 0.1))
    var v2 int = -2147483648
    var v3 uint8 = 255
    var v4 int16 = -32768
    var v5 string = ""
    if (1e10 >= 3.14159) {
        var v15 int8 = 42
        _ = v15
    } else {
        var v16 int16 = (v4 / ((int16(v2)) | 1))
        _ = v16
    }
    v17 := GenS1{F0: 100, F1: 32767}
    v4 = (v4 % (((v4 & -1)) | 1))
    v2 = (v2 ^ (v2 - (v2 * (v2 << ((42) & 63)))))
    v18 := []int{(v2 / ((v2) | 1)), 4294967295, (v2 >> ((255) & 63)), 127}
    var v19 uint8 = (^(v3))
    v19 += (^(v3))
    v19 = v19
    fmt.Println(v4)
    v4 += genFn0(-124, 825, true)
    var v20 uint16 = 1
    var v21 float32 = 3.14159
    for fi22 := 0; fi22 < 18; fi22++ {
        v21 -= 3.14159
        for fi23 := 0; fi23 < 18; fi23++ {
            if ((9223372036854775807 != 1) && (!(!true))) {
                v2 = fi22
                var v24 uint64 = 234
                v2 -= fi23
                _ = v24
            } else {
                var v25 float64 = (v1 * (1e10 + v1))
                _ = v25
            }
            if len(v18) > 0 {
                v18[0] = (-(fi22))
            }
            var v26 uint16 = 32767
            var v27 int8 = 100
            fmt.Println(v1, fi23)
            _ = fi23
            _ = v26
            _ = v27
        }
        _ = fi22
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        fmt.Println(v1, v19)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
