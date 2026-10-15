package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }
func genFn0() bool {
    v7 := GenS1{F0: (!(!false))}
    var v8 uint16 = 205
    _ = v7
    _ = v8
    return (620 < 65535)
}

func main() {
    v1 := []int{550, 32767}
    var v2 int16 = 520
    var v3 uint64 = 2
    var v4 uint16 = 329
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    switch 127 {
    case 127:
        fmt.Println(v3)
    }
    fmt.Println(v2, v3, v2)
    v3 -= v3
    if false {
        v2 = (v2 & (v2 << ((4294967295) & 63)))
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            v1 = append(v1, (fi9 % ((232) | 1)))
            _ = fi9
        }
    }
    if true {
        var v10 uint16 = (v4 + (v4 >> ((0) & 63)))
        for fi11 := 0; fi11 < 5; fi11++ {
            switch (fi11 << ((760) & 63)) {
            case 32767:
                var v12 float32 = (-2.25 * (1e10 / ((1.5 * (1e10 - float32(1e10))) - 3.14159)))
                _ = v12
            default:
                switch "hello" {
                case "":
                    var v13 uint16 = v4
                    _ = v13
                }
            }
            _ = fi11
        }
        var v14 float64 = 0.0
        _ = v10
        _ = v14
    }
    var v15 int64 = 0
    v2 = (v2 ^ v2)
    fmt.Println(v15, v15)
    v2 = 1
    switch 32767 {
    case 128:
        if false {
            switch 32768 {
            case -1:
                fmt.Println(v3, v15, v2)
            case 42:
                var v16 int32 = 874
                _ = v16
            default:
                var v17 float32 = (1e10 - ((0.0 * -2.25) / 1e10))
                _ = v17
            }
            v18 := GenS2{F0: "a", F1: 255}
            v4 -= v4
            _ = v18
        } else {
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(fi19)
                _ = fi19
            }
        }
    }
    v5[("hello" + "foo bar")] = -745
    v6[32767] = ("a" + "")
    if true {
        v15 -= v15
    } else {
        if true {
            v2 += (v2 + (v2 / ((v2) | 1)))
            v20 := GenS3{F0: 100, F1: 848, F2: (v2 & v2)}
            v3 = (v3 % ((18446744073709551615) | 1))
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS1{F0: (!((v4 % (((v4 ^ 255)) | 1)) < (v4 * 69)))}
                _ = fi21
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint16 = (v4 & (v4 & v4))
            _ = fi23
            _ = v24
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
