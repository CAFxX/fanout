package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 uint32, p1 int32) string {
    fmt.Println(p0, p0)
    if (734 > 255) {
        if true {
            var v4 float64 = float64(float32(-2.25))
            _ = v4
        }
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 69, F1: 32767}
            var v7 int = (fi5 * fi5)
            var v8 bool = true
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := []int{}
        _ = v9
    } else {
        switch "hello" {
        case "hello":
            v10 := GenS2{F0: p0, F1: 9223372036854775807, F2: -128}
            _ = v10
        case "":
            if true {
                var v11 uint8 = 1
                if false {
                    v11 = (v11 - v11)
                    var v12 uint8 = 0
                    var v13 int16 = -669
                    _ = v12
                    _ = v13
                }
                fmt.Println(v11, p1)
                _ = v11
            }
        case "\n\t":
            p0 += p0
        }
        var v14 bool = (1 != 27)
        v15 := []int{int(p1)}
        _ = v14
        _ = v15
    }
    var v16 string = ("a" + ("hello" + ""))
    var v17 uint8 = 42
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return "hello"
}
func genFn1(p0 uint32, p1 int) uint8 {
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int32 = -390
        if (42 >= 42) {
            v20 := []int{}
            p0 -= 42
            _ = v20
        } else {
            v19 += -2147483648
            p1 = p1
            p0 = p0
        }
        v21 := []int{9223372036854775807, fi18, (p1 ^ (p1 / ((p1) | 1)))}
        _ = fi18
        _ = v19
        _ = v21
    }
    switch (("a" + (("a" + ("a" + ("hello" + "a"))) + "hello")) + "foo bar") {
    case "hello":
        v22 := []int{(p1 / ((p1) | 1)), (p1 << ((841) & 63)), (-(p1)), p1}
        _ = v22
    default:
        if true {
            v23 := GenS1{F0: 899, F1: 1}
            _ = v23
        } else {
            v24 := GenS1{F0: 0, F1: 255}
            v25 := GenS2{F0: 2, F1: p1, F2: 42}
            _ = v24
            _ = v25
        }
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v26 := GenS2{F0: 65535, F1: -364, F2: -128}
        _ = v26
    default:
        v27 := []int{(p1 >> ((uint(p0)) & 63)), (p1 | p1), 32768, p1}
        _ = v27
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 uint32 = 4294967295
    fmt.Println(v3, v3)
    var v28 uint32 = v3
    v3 = v3
    v29 := []int{-2147483648, 127, -1, -9223372036854775808}
    for fi30 := 0; fi30 < 2; fi30++ {
        switch 0 {
        case -2147483648:
            var v31 int32 = 42
            _ = v31
        }
        v3 = (v3 >> ((18446744073709551615) & 63))
        _ = fi30
    }
    var v32 uint8 = 255
    fmt.Println(v32, v1)
    var v33 int16 = 42
    var v34 int64 = 0
    if ((0.1 - 1.5) == 1.5) {
        var v35 float32 = 3.14159
        v1 = v1
        _ = v35
    } else {
        if v2 {
            if false {
                v36 := GenS3{F0: "a", F1: -881}
                _ = v36
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
