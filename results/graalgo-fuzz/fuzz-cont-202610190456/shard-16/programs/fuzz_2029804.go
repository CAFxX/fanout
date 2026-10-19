package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 int8, p1 uint64, p2 int8) uint64 {
    if false {
        if true {
            var v5 float32 = 3.14159
            _ = v5
        } else {
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: (-(p0)), F1: 622}
                if (false && true) {
                    p1 -= (p1 + p1)
                }
                p1 = p1
                _ = fi6
                _ = v7
            }
            var v8 string = ("a" + "")
            if (p1 >= (p1 << ((912) & 63))) {
                if true {
                    v8 = (v8 + ((v8 + v8) + v8))
                    if false {
                        fmt.Println(p0, p2)
                        var v9 uint8 = 42
                        _ = v9
                    }
                    if (false && (-128 > p2)) {
                        p2 += p2
                        var v10 int32 = 764
                        _ = v10
                    } else {
                        var v11 uint = 42
                        var v12 uint64 = p1
                        _ = v11
                        _ = v12
                    }
                }
                v13 := GenS2{F0: -206, F1: 32768, F2: 1}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 8; fi14++ {
                    if false {
                        v15 := GenS1{F0: (^(p2)), F1: 1}
                        v16 := GenS2{F0: 916, F1: -38, F2: 128}
                        _ = v15
                        _ = v16
                    }
                    _ = fi14
                }
                v17 := GenS1{F0: p2, F1: 935}
                _ = v17
            }
            _ = v8
        }
        var v18 int64 = 0
        _ = v18
    }
    var v19 int8 = (p0 & p0)
    var v20 int = 792
    v21 := GenS1{F0: v19, F1: 1}
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            v20 += (fi22 + v20)
            p2 += p0
            switch ("" + "hello") {
            case "a":
                fmt.Println(p2, v19, p1)
            case "foo bar":
                fmt.Println(p2, p1, v20)
            case "\n\t":
                var v23 int8 = v19
                _ = v23
            default:
                if ((-(v20)) == (fi22 & v20)) {
                    v20 = 4294967295
                    var v24 float32 = 1.5
                    switch (((("a" + ("a" + "a")) + "a") + "") + "\n\t") {
                    case "hello":
                        var v25 int16 = 952
                        _ = v25
                    case "\n\t":
                        v26 := []int{(v20 % ((fi22) | 1)), 1}
                        _ = v26
                    case "a":
                        var v27 string = "a"
                        _ = v27
                    }
                    _ = v24
                } else {
                    v28 := []int{v20, 9223372036854775807, 2147483647}
                    _ = v28
                }
            }
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return p1
}

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 78
    var v3 uint16 = 65535
    v4 := []int{2147483647, 795, 127}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v29 := GenS1{F0: int8(v3), F1: 42}
        fmt.Println(v2)
        _ = v29
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v30 string = "\n\t"
        _ = v30
    case 127:
        fmt.Println(v2, v2)
    }
    v1 = (v1 * (v1 * 255))
    v31 := GenS3{F0: v3, F1: 4294967295, F2: true}
    var v32 uint16 = (v3 / ((uint16(v1)) | 1))
    var v33 uint32 = 4294967295
    v32 += v32
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
