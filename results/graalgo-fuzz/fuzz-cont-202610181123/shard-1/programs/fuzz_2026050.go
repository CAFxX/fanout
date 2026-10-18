package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 uint16, p1 string) int8 {
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 17; fi7++ {
        p1 = (p1 + p1)
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 string = p1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int16 = -556
    var v11 uint32 = 4294967295
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v10, p0, p1)
        v11 = v11
        switch p1 {
        case "":
            fmt.Println(v10, p0)
        default:
            for fi13 := 0; fi13 < 12; fi13++ {
                if (true || (!((!false) || false))) {
                    if (-1 >= (fi12 >> ((18446744073709551615) & 63))) {
                        v14 := GenS1{F0: (!false), F1: fi12}
                        _ = v14
                    } else {
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                    fmt.Println(v6, fi12, fi12)
                    switch p1 {
                    case "hello":
                        p0 = 32767
                    }
                }
                v16 := GenS1{F0: (!(39 != -121)), F1: 4294967295}
                _ = fi13
                _ = v16
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return 1
}

func main() {
    v1 := []int{0, -9223372036854775808}
    var v2 int16 = -414
    v3 := make(map[string]string)
    var v4 int64 = -1
    var v5 int64 = v4
    v5 -= 1
    v4 -= (v4 | 0)
    v17 := GenS1{F0: (!true), F1: -1}
    if true {
        fmt.Println(v4, v4)
        if true {
            v18 := GenS2{F0: 651, F1: 0, F2: 0}
            var v19 uint64 = 2
            _ = v18
            _ = v19
        }
        var v20 int = -2147483648
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 string = ("hello" + "foo bar")
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: 9223372036854775807, F1: 1, F2: v2}
    v4 = -614
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
