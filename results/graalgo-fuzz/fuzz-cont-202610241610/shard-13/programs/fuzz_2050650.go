package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint64, p1 int64) int16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(p0, p0)
        v7 := []int{(fi6 << ((2) & 63)), fi6, fi6}
        _ = fi6
        _ = v7
    }
    v8 := []int{4294967295}
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v8 = append(v8, fi9)
            _ = fi10
        }
        p0 -= p0
        if false {
            var v11 uint = 42
            p0 -= 42
            var v12 uint16 = 0
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                if len(v8) > 0 {
                    v8[0] = (fi13 - (fi13 | fi13))
                }
                _ = fi13
            }
            fmt.Println(p0, p0)
            v14 := GenS1{F0: (p1 >> ((226) & 63))}
            _ = v14
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 1000
}
func genFn1() int64 {
    if (0 == 65535) {
        var v15 float32 = (0.1 - -2.25)
        _ = v15
    }
    v16 := GenS2{F0: 551, F1: -1, F2: -464}
    _ = v16
    return 0
}

func main() {
    v1 := []int{127, -9223372036854775808, 42, -376}
    var v2 string = "\n\t"
    var v3 uint8 = 128
    var v4 int16 = 42
    var v5 float64 = (-2.25 + ((0.0 + 1e10) / 1.0))
    v17 := GenS3{F0: 4294967295, F1: v3}
    var v18 string = ("hello" + v2)
    _ = len(v1)
    fmt.Println(v18, v18)
    v3 = (v3 - (v3 | v3))
    fmt.Println(v2, v18)
    if false {
        switch v18 {
        case "hello":
            v18 = ("" + "hello")
        }
        var v19 float64 = v5
        v18 = ((v2 + "\n\t") + (v2 + ("" + "")))
        _ = v19
    }
    if (3.14159 >= (1.5 / (3.14159 + 0.1))) {
        if ((false && ((false && true) && (true || (89 != 42)))) && false) {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            switch -1 {
            case 0:
                var v20 uint8 = 42
                _ = v20
            case -132:
                for fi21 := 0; fi21 < 9; fi21++ {
                    if len(v1) > 0 {
                        v1[0] = -1
                    }
                    _ = fi21
                }
            }
            fmt.Println(v18, v5)
        }
        v3 -= 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
