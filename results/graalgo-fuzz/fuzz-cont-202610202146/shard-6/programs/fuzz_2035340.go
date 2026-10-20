package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    if (false || true) {
        var v6 int64 = 354
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(v6, v6)
            fmt.Println(v6)
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v9 string = ""
        fmt.Println(v9)
        v10 := []int{9223372036854775807, 719, 4294967295}
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    _ = v11
    return ""
}
func genFn1(p0 bool) int16 {
    v12 := []int{-1, 981}
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int = (fi13 | (fi13 & -842))
        v14 -= (v14 % ((fi13) | 1))
        switch (fi13 | v14) {
        case 9223372036854775807:
            switch "a" {
            case "":
                switch (v14 ^ 32768) {
                case -907:
                    var v15 bool = (!true)
                    _ = v15
                default:
                    if p0 {
                        v16 := GenS1{F0: 84}
                        for fi17 := 0; fi17 < 7; fi17++ {
                            if p0 {
                                var v18 int8 = -1
                                v18 = v18
                                _ = v18
                            }
                            _ = fi17
                        }
                        _ = v16
                    }
                }
            case "a":
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 uint8 = 67
                    if (((2 != 255) && p0) && (!("foo bar" >= "\n\t"))) {
                        v14 -= fi19
                        fmt.Println(fi19, p0, p0)
                    }
                    _ = fi19
                    _ = v20
                }
            default:
                switch "hello" {
                case "":
                    fmt.Println(v14, fi13, v14)
                case "foo bar":
                    fmt.Println(fi13, v14)
                }
            }
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := []int{-1}
    var v2 uint16 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 int = 9223372036854775807
    v5 += (v5 + (v5 % ((v5) | 1)))
    switch v5 {
    case -386:
        v21 := GenS1{F0: -110}
        _ = v21
    case 2147483647:
        if true {
            if len(v1) > 0 {
                v1[0] = v5
            }
            v2 += 636
        }
    case 32768:
        v2 -= (v2 | (v2 * (v2 & 2)))
    }
    var v22 uint32 = 4294967295
    var v23 int = int(v2)
    v24 := v3[(v22 + (v22 * v22))]
    v5 = v24
    v25 := GenS1{F0: 0}
    var v26 bool = (806 > 4294967295)
    v27 := GenS1{F0: 6}
    v24 = (v5 >> ((255) & 63))
    if (!((true || (2 != 65535)) || false)) {
        var v28 float32 = (1e10 + 0.0)
        v23 = 32767
        v4[v23] = v22
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
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
