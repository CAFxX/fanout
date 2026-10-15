package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 int64) int8 {
    if (!(!(255 != 42))) {
        for fi4 := 0; fi4 < 8; fi4++ {
            for fi5 := 0; fi5 < 17; fi5++ {
                var v6 uint = 255
                _ = fi5
                _ = v6
            }
            for fi7 := 0; fi7 < 12; fi7++ {
                p0 = (p0 >> ((812) & 63))
                _ = fi7
            }
            _ = fi4
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: -469, F1: 896, F2: 691}
            _ = fi8
            _ = v9
        }
        var v10 uint8 = 180
        _ = v10
    }
    var v11 uint = 461
    v12 := GenS1{F0: 517, F1: 511, F2: 2}
    for fi13 := 0; fi13 < 18; fi13++ {
        v11 = (v11 - v11)
        v14 := GenS1{F0: 0, F1: 4294967295, F2: 4294967295}
        _ = fi13
        _ = v14
    }
    v15 := []int{128, 128}
    _ = p0
    _ = v11
    _ = v12
    _ = v15
    return 100
}
func genFn1(p0 string) string {
    if true {
        if (-1 == 32767) {
            p0 = "hello"
            v16 := GenS1{F0: 454, F1: 2147483647, F2: 539}
            _ = v16
        }
    }
    p0 = ""
    _ = p0
    return "\n\t"
}

func main() {
    v1 := []int{128, -2147483648, 32768}
    var v2 int16 = 32767
    v3 := make(map[string]bool)
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 int8 = int8(v2)
        _ = fi17
        _ = v18
    }
    var v19 int32 = 42
    _ = len(v3)
    v3[("" + ("" + ((("a" + "hello") + "foo bar") + ("hello" + ("hello" + "a")))))] = false
    if true {
        delete(v3, ((((("" + "") + "\n\t") + (("foo bar" + "a") + "hello")) + ("\n\t" + "a")) + ("" + "hello")))
        v2 = (v2 | v2)
    }
    if (!true) {
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint32 = 63
                var v23 uint16 = 813
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v2 -= (v2 >> ((42) & 63))
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 uint32 = 0
            v19 = (-(v19))
            fmt.Println(v19, v19, v2)
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS1{F0: (v19 & (v19 + v19)), F1: 2147483648, F2: 42}
    v2 = v2
    var v27 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
