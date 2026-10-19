package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 int) int16 {
    var v7 int64 = 0
    var v8 uint16 = 255
    switch p1 {
    case -2147483648:
        fmt.Println(v8, v7, v8)
    case 1:
        v9 := GenS1{F0: 317}
        _ = v9
    case -9223372036854775808:
        fmt.Println(v7, p1, v8)
    default:
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 float64 = -2.25
            var v12 bool = ((p0 & (p0 | (p0 << ((42) & 63)))) == (p0 | (p0 + (p0 ^ (p0 * 1)))))
            for fi13 := 0; fi13 < 3; fi13++ {
                switch int(v8) {
                case -1:
                    v14 := GenS2{F0: (v8 * (v8 << ((42) & 63))), F1: p0}
                    _ = v14
                case -2147483648:
                    fmt.Println(v7, v11, fi10)
                }
                fmt.Println(fi10)
                _ = fi13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int8, p1 int32) uint8 {
    v15 := []int{128, 128}
    var v16 int16 = -1
    fmt.Println(p0, p0, v16)
    fmt.Println(v16, v16)
    if false {
        if (((v16 | -32768) == v16) || true) {
            v15 = append(v15, 170)
        }
        if (42 == 4294967295) {
            if (584 > -1) {
                var v17 int16 = 32767
                switch 32768 {
                case 127:
                    var v18 uint64 = 627
                    _ = v18
                case 4294967295:
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        var v19 bool = true
                        _ = v19
                    case -2147483648:
                        v20 := GenS1{F0: 174}
                        _ = v20
                    }
                default:
                    v16 -= v17
                }
                _ = v17
            } else {
                var v21 uint64 = 42
                fmt.Println(v21)
                if true {
                    var v22 uint = 255
                    var v23 uint16 = 65535
                    _ = v22
                    _ = v23
                }
                _ = v21
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint32 = 978
                v16 = (v16 << ((376) & 63))
                p1 -= p1
                _ = fi24
                _ = v25
            }
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 2
}

func main() {
    v1 := []string{"hello", ""}
    v2 := []uint32{0, 2147483647}
    var v3 int = 301
    v4 := make(map[int]uint32)
    var v5 uint8 = 42
    var v6 uint64 = 42
    v5 += (v5 % ((v5) | 1))
    var v26 float64 = float64(0.0)
    v27 := GenS1{F0: 701}
    if true {
        delete(v4, (v3 % ((v3) | 1)))
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS2{F0: 306, F1: -72}
            v30 := GenS1{F0: 65535}
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        v31 := GenS2{F0: 2, F1: 55}
        v3 += v3
        _ = v31
    }
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := GenS3{F0: "foo bar"}
        delete(v4, (fi32 << ((464) & 63)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v26, v26)
    v3 -= 1
    v6 = 4294967295
    fmt.Println(v6, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v34 uint32 = 2
    var v35 int16 = -114
    var v36 int16 = v35
    v37 := GenS2{F0: 65535, F1: -51}
    v26 -= v26
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
