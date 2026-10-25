package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 uint8, p1 uint) int64 {
    p1 -= (p1 >> ((p1) & 63))
    v7 := []int{128, 32767, 9223372036854775807, -2147483648}
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1(p0 int16, p1 int) int32 {
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := []int{}
            var v11 int = fi9
            v11 += (fi9 + 32767)
            fmt.Println(v11, p1, v11)
            p0 += 32767
            var v12 uint16 = 569
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    p1 = (p1 / ((p1) | 1))
    v13 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v13
    return 2147483647
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = float64(float32((0.0 - 1e10)))
    var v3 uint = 65535
    var v4 uint32 = 4294967295
    v5 := []int{1, 184}
    var v6 int8 = -128
    v3 -= (^(v3))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 += int8(v1)
    switch (("\n\t" + ("" + "a")) + "hello") {
    case "a":
        v4 += v4
    case "":
        v6 -= (v6 ^ (v6 << (((v3 & v3)) & 63)))
    default:
        v3 += v3
    }
    if (!((("hello" + ("foo bar" + "a")) + ("a" + "\n\t")) < ((("hello" + ("\n\t" + "\n\t")) + "") + (("a" + "") + ("foo bar" + "\n\t"))))) {
        for fi14 := 0; fi14 < 7; fi14++ {
            v3 = 1
            var v15 uint8 = v1
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: genFn1(-1, 128)}
        var v17 uint = (^(v3))
        _ = v16
        _ = v17
    } else {
        if false {
            v5 = append(v5, 128)
            v18 := GenS2{F0: "foo bar", F1: 2, F2: true}
            var v19 float64 = v2
            _ = v18
            _ = v19
        }
    }
    v5 = append(v5, 32767)
    var v20 int32 = genFn1(1000, 889)
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v2, v3, fi21)
        for fi22 := 0; fi22 < 18; fi22++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v23 bool = true
            _ = fi22
            _ = v23
        }
        v24 := GenS3{F0: (v20 - v20), F1: v6}
        _ = fi21
        _ = v24
    }
    fmt.Println(v2, v4, v1)
    v5 = append(v5, -9223372036854775808)
    v6 -= v6
    fmt.Println(v2, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
