package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
func genFn0(p0 int16, p1 uint16, p2 uint) uint {
    if false {
        var v7 uint32 = 965
        v8 := GenS1{F0: (!(false || false))}
        p0 = (p0 >> (((p2 % ((uint(p1)) | 1))) & 63))
        _ = v7
        _ = v8
    } else {
        if true {
            switch "a" {
            case "a":
                p1 += 294
            case "foo bar":
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: (int16(p1) < (^(p0)))}
                    _ = fi9
                    _ = v10
                }
            default:
                var v11 float64 = 1e10
                _ = v11
            }
            if (uint64(p2) > 252) {
                var v12 int16 = 1000
                if true {
                    v13 := GenS1{F0: true}
                    _ = v13
                } else {
                    v14 := GenS1{F0: true}
                    v15 := GenS1{F0: false}
                    _ = v14
                    _ = v15
                }
                v16 := []int{-590, 9223372036854775807, -9223372036854775808, -720}
                _ = v12
                _ = v16
            } else {
                p2 += p2
                fmt.Println(p0)
            }
        } else {
            switch ("" + "\n\t") {
            case "hello":
                if true {
                    var v17 float32 = (0.0 * 3.14159)
                    _ = v17
                }
            default:
                if false {
                    var v18 int = 32768
                    _ = v18
                }
            }
            fmt.Println(p2)
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := GenS1{F0: true}
                var v21 float32 = 1.5
                v22 := []int{}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        fmt.Println(p1, p1)
        p2 = p2
    }
    var v23 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return p2
}
func genFn1(p0 int8) uint {
    var v24 uint32 = 2147483648
    switch -144 {
    case 32767:
        var v25 float32 = 0.1
        _ = v25
    }
    _ = p0
    _ = v24
    return genFn0(42, 832, 435)
}

func main() {
    v1 := []int{505, -895}
    var v2 int16 = 1000
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v5 uint64 = 42
    var v6 float64 = -2.25
    if true {
        delete(v3, "")
    } else {
        v4[-9223372036854775808] = -437
        v2 = v2
    }
    var v26 uint64 = (^(v5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 float32 = float32(v6)
    var v28 int = -2147483648
    v29 := v3[("foo bar" + "foo bar")]
    var v30 uint32 = 4294967295
    switch int(v2) {
    case 285:
        v30 += (v30 << ((65535) & 63))
    }
    v31 := GenS2{F0: v30, F1: 1, F2: -482}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
