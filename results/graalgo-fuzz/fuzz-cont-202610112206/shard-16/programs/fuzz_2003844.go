package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
func genFn0(p0 uint16) int32 {
    var v5 int64 = 1
    var v6 int16 = -32768
    v7 := []int{32768}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int16, p1 int16) uint8 {
    var v8 int64 = 1
    v9 := GenS1{F0: v8, F1: -772}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 bool = false
        var v12 float64 = 0.1
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 7
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 780
    v3 := make(map[int]bool)
    var v4 uint8 = 203
    fmt.Println(v2, v4)
    var v13 int16 = 619
    v13 -= -747
    v4 -= 194
    for fi14 := 0; fi14 < 10; fi14++ {
        v2 = 799
        v15 := []int{fi14, fi14}
        if (v2 != v2) {
            var v16 int64 = 42
            v17 := GenS1{F0: v16, F1: 127}
            v13 = v13
            _ = v16
            _ = v17
        } else {
            var v18 uint8 = (v4 - v4)
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v18, fi14, fi14)
                _ = fi19
            }
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    switch 127 {
    case -9223372036854775808:
        v1[2] = (("foo bar" + ((("\n\t" + (("" + "") + "\n\t")) + "a") + ("\n\t" + (("" + "\n\t") + "foo bar")))) + "a")
    case 2147483647:
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int = fi20
                _ = fi21
                _ = v22
            }
            fmt.Println(v4, v2)
            v23 := GenS2{F0: 32767, F1: "foo bar", F2: 9223372036854775807}
            _ = fi20
            _ = v23
        }
    case -894:
        var v24 uint32 = 65535
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        v13 = (v13 - v13)
        _ = fi25
    }
    v26 := []int{-938, 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
