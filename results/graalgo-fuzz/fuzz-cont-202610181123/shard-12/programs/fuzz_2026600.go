package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 string; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int32 = 256
    v4 := []int{127, 42, -1}
    v5 := make(map[string]bool)
    delete(v2, "foo bar")
    if (42 >= 811) {
        v3 = v3
    }
    var v6 uint = 4294967295
    var v7 string = "\n\t"
    var v8 int = -9223372036854775808
    for fi9 := 0; fi9 < 2; fi9++ {
        v7 = v7
        v10 := GenS1{F0: "foo bar"}
        _ = fi9
        _ = v10
    }
    switch v8 {
    case -888:
        var v11 uint32 = 65535
        _ = v11
    }
    if false {
        switch v7 {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            v6 -= 403
        }
        v6 -= v6
    }
    v7 = v7
    switch v8 {
    case 2147483647:
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "a":
            var v12 string = (v7 + v7)
            _ = v12
        default:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: v7}
                _ = fi13
                _ = v14
            }
        }
    case -715:
        var v15 int16 = 32767
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS1{F0: v7}
        _ = fi16
        _ = v17
    }
    if false {
        v8 -= (v8 & (v8 % (((v8 << ((v6) & 63))) | 1)))
        var v18 float32 = float32(1e10)
        v4 = append(v4, (-(v8)))
        _ = v18
    } else {
        if (v6 != (^(v6))) {
            v4 = append(v4, v8)
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            delete(v1, (fi19 + -307))
            v4 = append(v4, (-(v8)))
            fmt.Println(v3, v6, v7)
            _ = fi19
        }
    }
    var v20 int8 = 72
    fmt.Println(v20)
    v21 := GenS2{F0: (((("foo bar" + (v7 + (v7 + "hello"))) + ("hello" + (v7 + "foo bar"))) + v7) + ((((((v7 + "") + (v7 + v7)) + v7) + v7) + (((v7 + (("a" + v7) + "a")) + v7) + (("foo bar" + v7) + "a"))) + ("a" + (v7 + ("hello" + "foo bar"))))), F1: v7, F2: 2}
    v4 = append(v4, v8)
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 bool = true
        if (194 >= 1) {
            var v24 float32 = 0.0
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 float32 = v24
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            delete(v5, v7)
            var v27 int32 = (v3 / (((^(v3))) | 1))
            v28 := GenS2{F0: v7, F1: v7, F2: 4294967295}
            _ = v27
            _ = v28
        }
        v20 = (v20 & (v20 % ((v20) | 1)))
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
