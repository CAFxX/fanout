package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int; F2 string }
func genFn0() bool {
    var v4 int64 = 9223372036854775807
    fmt.Println(v4)
    v5 := GenS1{F0: 65535, F1: 127, F2: 1}
    if (12 > 4294967295) {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 252, F1: 1, F2: 557}
            v8 := []int{}
            v9 := GenS1{F0: 4294967295, F1: 100, F2: 255}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 int64 = (v4 | v4)
        if (-1 >= 2147483647) {
            v11 := []int{-9223372036854775808, 286, 2147483647, 161}
            _ = len(v11)
            _ = v11
        } else {
            v4 += 9223372036854775807
        }
        _ = v10
    }
    if true {
        var v12 string = "\n\t"
        _ = v12
    }
    _ = v4
    _ = v5
    return (!((1.5 + 1.5) <= -2.25))
}
func genFn1(p0 int8, p1 int, p2 int16) bool {
    var v13 uint32 = 4294967295
    var v14 uint32 = (v13 ^ v13)
    if false {
        v15 := []int{0, (^(p1))}
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                if false {
                    switch int(p2) {
                    case 144:
                        switch "hello" {
                        case "\n\t":
                            v14 -= (^(v14))
                            p2 -= p2
                        case "hello":
                            v18 := GenS1{F0: 42, F1: p0, F2: 0}
                            _ = v18
                        case "":
                            v19 := GenS1{F0: 0, F1: -128, F2: 1}
                            _ = v19
                        }
                    }
                    if len(v15) > 0 {
                        v15[0] = 32767
                    }
                }
                fmt.Println(v13)
                _ = fi17
            }
            _ = fi16
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -9223372036854775808
    var v3 float64 = 1e10
    v2 += v2
    v20 := []int{-9223372036854775808, -538, 32768, -1}
    if false {
        v21 := GenS1{F0: 255, F1: 42, F2: 18446744073709551615}
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v2 = v2
        if false {
            v2 -= -1
        }
        var v23 float64 = ((v3 + (v3 * v3)) + 0.0)
        _ = fi22
        _ = v23
    }
    var v24 int64 = (v2 | 9223372036854775807)
    v25 := GenS2{F0: 0}
    var v26 int32 = 42
    var v27 int = -589
    var v28 bool = true
    v27 -= (v27 * v27)
    var v29 int64 = (-(v24))
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS2{F0: (-(fi30))}
            switch ("\n\t" + (("hello" + "hello") + (("a" + "hello") + (("hello" + "") + (("foo bar" + "a") + "\n\t"))))) {
            case "foo bar":
                v33 := GenS2{F0: (v27 >> ((65535) & 63))}
                v2 = (v24 | (-(v24)))
                var v34 float32 = 0.0
                _ = v33
                _ = v34
            case "\n\t":
                v2 -= v29
            }
            v27 = -9223372036854775808
            var v35 float64 = v3
            fmt.Println(v27, fi31, v24)
            v28 = (557 >= 893)
            v36 := v1[fi30]
            _ = fi31
            _ = v32
            _ = v35
            _ = v36
        }
        _ = fi30
    }
    if true {
        v3 += 3.14159
        v37 := GenS3{F0: 991, F1: v27, F2: "hello"}
        _ = len(v1)
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
