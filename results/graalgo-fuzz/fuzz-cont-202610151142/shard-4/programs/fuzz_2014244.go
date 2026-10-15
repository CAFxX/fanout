package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 bool }
func genFn0(p0 uint) bool {
    switch -1 {
    case 127:
        var v6 float64 = 3.14159
        _ = v6
    case 721:
        var v7 string = "foo bar"
        _ = v7
    case 0:
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                v9 := []int{128, (fi8 - fi8)}
                if len(v9) > 0 {
                    v9[0] = fi8
                }
                if true {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(fi8, p0)
                        switch (fi10 & fi8) {
                        case 128:
                            v11 := GenS1{F0: (true && (1000 < 0)), F1: fi8}
                            _ = v11
                        case 2147483647:
                            p0 = 652
                        case 32767:
                            p0 -= 65535
                        }
                        var v12 int8 = 100
                        _ = fi10
                        _ = v12
                    }
                    var v13 uint8 = 128
                    _ = v13
                } else {
                    p0 = 1
                    v9 = append(v9, 32768)
                    p0 -= p0
                }
                _ = v9
            } else {
                p0 += 0
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS1{F0: true, F1: (fi8 + fi8)}
                    v16 := GenS1{F0: false, F1: 620}
                    v17 := GenS1{F0: (false && false), F1: 281}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                if ((500 != (fi8 + 127)) || (!true)) {
                    fmt.Println(p0, p0)
                    v18 := GenS2{F0: 282, F1: "\n\t", F2: 430}
                    if (-2.25 >= 1e10) {
                        p0 -= (p0 >> (((p0 << (((p0 + (^(p0)))) & 63))) & 63))
                    }
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        fmt.Println(fi19, p0)
                        _ = fi19
                    }
                    v20 := []int{}
                    _ = v20
                }
            }
            v21 := GenS1{F0: false, F1: fi8}
            _ = fi8
            _ = v21
        }
    default:
        var v22 int64 = -535
        _ = v22
    }
    p0 += 65535
    _ = p0
    return (!true)
}

func main() {
    var v1 bool = false
    var v2 int32 = 2147483647
    var v3 bool = true
    var v4 float32 = -2.25
    v5 := make(map[int]uint32)
    v1 = (!true)
    v23 := []int{0}
    v3 = v1
    var v24 float32 = 1.5
    var v25 float64 = 1e10
    var v26 uint64 = 893
    v27 := GenS3{F0: v2, F1: 127}
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS2{F0: 139, F1: (("" + ("hello" + (("a" + (("hello" + "foo bar") + "foo bar")) + (("\n\t" + ("" + "foo bar")) + "")))) + "\n\t"), F2: -1}
        fmt.Println(v25, v25)
        _ = fi28
        _ = v29
    }
    v30 := GenS4{F0: 1, F1: 347, F2: (!false)}
    var v31 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
