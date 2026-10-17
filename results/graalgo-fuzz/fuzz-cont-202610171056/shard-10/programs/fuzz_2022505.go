package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 int; F2 string }
func genFn0(p0 int16) uint8 {
    if (true && (true || false)) {
        v5 := GenS1{F0: 1, F1: 4294967295}
        var v6 uint8 = 2
        _ = v5
        _ = v6
    } else {
        var v7 bool = false
        fmt.Println(p0, p0)
        var v8 int64 = 1
        _ = v7
        _ = v8
    }
    var v9 uint8 = 203
    _ = p0
    _ = v9
    return v9
}
func genFn1() bool {
    var v10 uint = 692
    v11 := []int{328, 1}
    v10 = 0
    v12 := GenS2{F0: 6, F1: 127, F2: "a"}
    var v13 bool = true
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return (1.5 != float64(1.5))
}

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    v4 := []string{(("a" + "a") + "a")}
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint8 = 255
        var v16 int8 = 100
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if false {
            v1 = v1
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 int = 128
                var v20 int8 = 42
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 bool = (false && (42 < 151))
                v2 -= (v2 * (v2 % ((v2) | 1)))
                _ = fi21
                _ = v22
            }
        }
        _ = fi17
    }
    v23 := GenS1{F0: (v1 >> ((65535) & 63)), F1: 1}
    var v24 uint64 = 42
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint32 = 2
        var v27 uint16 = v1
        _ = fi25
        _ = v26
        _ = v27
    }
    v2 += v2
    var v28 int8 = -128
    v29 := GenS2{F0: 2, F1: 32767, F2: ("" + ("" + "foo bar"))}
    fmt.Println(v24, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
