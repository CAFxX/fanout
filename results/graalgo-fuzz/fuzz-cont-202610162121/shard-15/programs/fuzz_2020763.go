package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(fi4)
        var v5 string = ""
        if false {
            v6 := GenS1{F0: 18446744073709551615, F1: 635}
            _ = v6
        } else {
            var v7 uint8 = 3
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: 42, F1: 807}
                var v10 float64 = (0.0 + -2.25)
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v11 := GenS1{F0: 728, F1: 276}
    _ = v11
    return 92
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int64 = -1
    v3 := make(map[uint32]uint32)
    if false {
        var v12 uint8 = 33
        if true {
            fmt.Println(v12, v12)
        }
        _ = v12
    } else {
        v2 += 42
        var v13 uint8 = 42
        v14 := GenS1{F0: 255, F1: 37}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 bool = true
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    v2 += 640
    var v17 uint8 = 2
    v2 = v2
    v2 = v2
    v2 = (^(v2))
    v2 += genFn0()
    var v18 bool = true
    if false {
        if v18 {
            if true {
                var v19 uint8 = v17
                _ = v19
            } else {
                var v20 int16 = -714
                _ = v20
            }
        } else {
            if v18 {
                for fi21 := 0; fi21 < 3; fi21++ {
                    v18 = v18
                    var v22 float32 = 3.14159
                    _ = fi21
                    _ = v22
                }
            }
            v2 = genFn0()
            switch "a" {
            case "foo bar":
                v2 = (v2 * 0)
            case "hello":
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 float32 = 3.14159
                    _ = fi23
                    _ = v24
                }
            case "\n\t":
                v2 = (v2 >> ((2) & 63))
            default:
                v18 = true
            }
        }
        v25 := GenS2{F0: (42 > 42)}
        _ = v25
    } else {
        v17 += (v17 * (v17 - 220))
        var v26 int8 = -49
        _ = v26
    }
    var v27 uint8 = v17
    var v28 string = ("" + ("a" + ""))
    var v29 bool = v18
    var v30 int32 = 298
    v31 := GenS3{F0: false, F1: 1, F2: (v17 / (((v27 / ((v17) | 1))) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
