package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64, p1 string) uint8 {
    v7 := GenS1{F0: -1, F1: p0}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 int8) uint64 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]int)
    var v3 uint64 = 4294967295
    var v4 uint = 42
    v5 := make(map[string]uint32)
    var v6 uint8 = 229
    v8 := []int{(v1 + v1)}
    switch "foo bar" {
    case "a":
        v3 = v3
    }
    v9 := GenS1{F0: 42, F1: genFn1(1, 100)}
    var v10 int32 = -2147483648
    switch (("hello" + "hello") + "\n\t") {
    case "hello":
        v10 += int32(v1)
    case "":
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: 32767, F1: 2}
            _ = fi11
            _ = v12
        }
    default:
        if true {
            var v13 int = -414
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v4 -= 0
        _ = fi14
    }
    delete(v2, 854)
    v5["foo bar"] = 0
    if ((2 < 65535) && false) {
        var v15 int = (v1 + (v1 + v1))
        var v16 uint32 = 42
        var v17 float32 = 0.0
        _ = v15
        _ = v16
        _ = v17
    }
    switch -605 {
    case 765:
        for fi18 := 0; fi18 < 18; fi18++ {
            switch v1 {
            case 42:
                v10 += v10
            default:
                var v19 uint8 = 255
                _ = v19
            }
            _ = fi18
        }
    default:
        v5["a"] = 2147483648
    }
    if ((("" + "foo bar") + ("a" + "")) <= "") {
        var v20 int16 = 32767
        _ = v20
    }
    switch "a" {
    case "a":
        v21 := GenS1{F0: 1000, F1: genFn1(-1, int8(v6))}
        _ = v21
    }
    v3 = genFn1(v10, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
