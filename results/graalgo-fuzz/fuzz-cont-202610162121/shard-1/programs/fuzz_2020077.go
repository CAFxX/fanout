package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 int, p2 uint16) int8 {
    p2 -= 1
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int8, p1 uint) string {
    if false {
        v4 := GenS1{F0: 42}
        _ = v4
    } else {
        var v5 float64 = (-2.25 / ((-2.25 / 0.1) - -2.25))
        _ = v5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        p1 += p1
        if true {
            p1 = (^(p1))
            p1 -= (p1 ^ 50)
        }
        _ = fi6
    }
    switch 2147483647 {
    case 2147483647:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int = (fi7 >> (((p1 & p1)) & 63))
            _ = fi7
            _ = v8
        }
    }
    if false {
        v9 := GenS2{F0: 51, F1: 722}
        p0 -= -1
        _ = v9
    } else {
        v10 := GenS3{F0: 2, F1: 147, F2: 42}
        if false {
            p0 -= (-(p0))
            v11 := []int{-1, -2147483648}
            p1 = (p1 - p1)
            _ = v11
        } else {
            if ((128 < 576) && false) {
                if (!true) {
                    v12 := []int{}
                    p0 += -120
                    _ = v12
                }
            }
            v13 := GenS4{F0: 220, F1: 1, F2: -469}
            _ = v13
        }
        if (("foo bar" + "foo bar") < ("" + ((("foo bar" + ("a" + "\n\t")) + ("\n\t" + ("a" + ""))) + ""))) {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int16 = 137
                fmt.Println(v15, v15, p0)
                fmt.Println(v15, p1)
                _ = fi14
                _ = v15
            }
        }
        _ = v10
    }
    var v16 uint32 = 100
    _ = p0
    _ = p1
    _ = v16
    return "\n\t"
}

func main() {
    var v1 int32 = 1
    var v2 int8 = 127
    var v3 uint8 = 182
    var v17 float64 = 3.14159
    v18 := []int{4294967295}
    v19 := GenS3{F0: 2, F1: 373, F2: 0}
    v20 := GenS3{F0: 18446744073709551615, F1: 4294967295, F2: -949}
    fmt.Println(v2)
    if ("a" == "foo bar") {
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 uint = 226
                _ = len(v18)
                v17 = v17
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    if (!(42 > 36)) {
        if true {
            v17 = ((-2.25 + ((v17 / 1.5) + -2.25)) / (-2.25 * v17))
        }
    }
    if (false && (-2.25 == (-2.25 - (0.1 / (1e10 / 1e10))))) {
        if false {
            v1 -= (v1 * int32(v3))
            if len(v18) > 0 {
                v18[0] = 2147483647
            }
        } else {
            v3 = (v3 % (((v3 / ((127) | 1))) | 1))
            fmt.Println(v3, v2, v3)
        }
        var v24 uint = 0
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS2{F0: 193, F1: 42}
            v3 = v3
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    switch "\n\t" {
    case "foo bar":
        var v27 uint16 = 32767
        _ = v27
    case "hello":
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS2{F0: 324, F1: 216}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                v2 += v2
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
    }
    v1 -= (v1 * (v1 - 2147483647))
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
