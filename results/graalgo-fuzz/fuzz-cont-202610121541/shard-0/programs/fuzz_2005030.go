package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 int8, p2 int8) uint16 {
    v6 := []int{9223372036854775807, 9223372036854775807, -2147483648}
    var v7 int8 = (p1 ^ (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 754
}
func genFn1(p0 bool) uint16 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int8 = -102
        v10 := []int{}
        var v11 string = ("a" + "a")
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    p0 = p0
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int32 = 2147483647
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    p0 = p0
    _ = p0
    return 873
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 255
    var v3 float32 = 1e10
    v4 := make(map[int]string)
    v5 := []string{"a", "hello", (("\n\t" + "\n\t") + "")}
    var v15 int32 = 232
    switch -1 {
    case 1:
        v1 = v1
    case -284:
        var v16 int16 = 247
        _ = v16
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(v3)
            _ = fi17
        }
    }
    v2 = (v2 | v2)
    for fi18 := 0; fi18 < 16; fi18++ {
        switch fi18 {
        case 198:
            v1 = v1
        case 127:
            for fi19 := 0; fi19 < 9; fi19++ {
                v15 = -608
                fmt.Println(fi19, v1)
                _ = fi19
            }
        }
        v20 := GenS1{F0: 2, F1: 1, F2: uint16(v1)}
        if (!false) {
            fmt.Println(v15)
            var v21 string = "a"
            _ = len(v5)
            _ = v21
        }
        _ = fi18
        _ = v20
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if false {
            if true {
                for fi23 := 0; fi23 < 10; fi23++ {
                    switch ("\n\t" + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        var v24 float32 = -2.25
                        _ = v24
                    case "hello":
                        v25 := GenS2{F0: 488, F1: (fi22 - -9223372036854775808), F2: (v1 % (((v1 | (v1 / ((v1) | 1)))) | 1))}
                        _ = v25
                    case "":
                        var v26 uint = (v2 + v2)
                        _ = v26
                    }
                    v27 := v4[fi22]
                    _ = fi23
                    _ = v27
                }
                v5 = append(v5, "foo bar")
                if true {
                    v3 = float32(((1.5 / 3.14159) - 3.14159))
                    v3 -= v3
                }
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                v15 -= 0
                fmt.Println(v1, fi28, v3)
                _ = fi28
            }
        }
        v3 -= (float32(3.14159) - v3)
        _ = fi22
    }
    v3 -= (v3 + 1e10)
    v29 := GenS3{F0: v1, F1: 1000}
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 int8 = 60
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
