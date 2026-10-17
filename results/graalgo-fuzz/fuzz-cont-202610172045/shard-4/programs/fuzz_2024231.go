package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int32, p1 uint) uint16 {
    if (9223372036854775807 == -1) {
        for fi5 := 0; fi5 < 7; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                if ((3.14159 * 0.0) != (float32(3.14159) + float32((0.1 * (0.0 + 1.5))))) {
                    v7 := GenS1{F0: 704}
                    _ = v7
                }
                if (2 >= 255) {
                    v8 := GenS2{F0: 0, F1: 19, F2: p0}
                    var v9 uint8 = 2
                    fmt.Println(p1)
                    _ = v8
                    _ = v9
                }
                v10 := []int{fi6, -1}
                _ = fi6
                _ = v10
            }
            var v11 int64 = 42
            v12 := []int{(fi5 & -987)}
            _ = fi5
            _ = v11
            _ = v12
        }
        p0 = (p0 - 0)
    } else {
        switch "hello" {
        case "a":
            var v13 uint8 = 210
            _ = v13
        default:
            var v14 float32 = 1e10
            _ = v14
        }
        if false {
            v15 := GenS2{F0: 2, F1: 177, F2: p0}
            var v16 uint16 = 255
            var v17 uint8 = 255
            _ = v15
            _ = v16
            _ = v17
        }
        fmt.Println(p0, p0)
    }
    p1 += (^(p1))
    fmt.Println(p1, p1)
    p1 += 255
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 int = -379
    var v2 int32 = -886
    var v3 float32 = 0.0
    var v4 uint64 = 4294967295
    v1 -= (v1 ^ v1)
    var v18 string = (("\n\t" + "") + "\n\t")
    var v19 bool = false
    v20 := GenS1{F0: v4}
    var v21 int8 = 1
    v1 = (v1 / (((v1 + (v1 >> ((65535) & 63)))) | 1))
    var v22 int64 = 0
    v23 := []int{v1, 4294967295, (v1 & v1)}
    v22 = (-(v22))
    if v19 {
        if v19 {
            if v19 {
                var v24 int64 = (^(v22))
                _ = v24
            } else {
                v25 := GenS3{F0: v1}
                _ = v25
            }
            switch (-(v1)) {
            case 451:
                v18 = v18
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                fmt.Println(v2, v22)
                if false {
                    if len(v23) > 0 {
                        v23[0] = v1
                    }
                } else {
                    v27 := GenS1{F0: (v4 * 2)}
                    _ = v27
                }
                _ = fi26
            }
        } else {
            v18 = v18
            _ = len(v23)
        }
        var v28 float64 = 0.0
        var v29 int64 = (v22 & (v22 << ((0) & 63)))
        _ = v28
        _ = v29
    }
    if false {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v30 string = ("" + v18)
        _ = v30
    } else {
        var v31 bool = v19
        _ = v31
    }
    var v32 float64 = 1.5
    var v33 int64 = 1
    if false {
        var v34 int64 = (v33 * (v33 * (-(v33))))
        var v35 uint8 = 1
        var v36 int32 = v2
        _ = v34
        _ = v35
        _ = v36
    }
    v37 := GenS2{F0: 171, F1: 1, F2: v2}
    var v38 int32 = (v2 / (((v2 / (((v2 >> ((0) & 63))) | 1))) | 1))
    fmt.Println(v3)
    var v39 int64 = (v33 >> ((924) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
