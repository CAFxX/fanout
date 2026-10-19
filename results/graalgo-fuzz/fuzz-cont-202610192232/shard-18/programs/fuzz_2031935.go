package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 int16) string {
    if false {
        v5 := GenS1{F0: 909, F1: p1}
        _ = v5
    } else {
        v6 := []int{377, 42, -683, 1}
        _ = v6
    }
    if true {
        p1 += p1
        var v7 uint16 = 65535
        fmt.Println(p0)
        _ = v7
    }
    _ = p0
    _ = p1
    return ("hello" + (("hello" + ((("hello" + "") + "") + (("" + "foo bar") + "foo bar"))) + "hello"))
}
func genFn1(p0 int32, p1 uint64) uint {
    p1 += 42
    p0 = (p0 / ((p0) | 1))
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = false
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]string)
    v3 += v3
    v3 -= v3
    v2 = ("hello" <= "hello")
    v2 = v2
    v3 = 540
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS2{F0: 1, F1: -410, F2: 624}
            var v11 float32 = ((float32((3.14159 + 1e10)) + ((0.0 - 1.5) / ((3.14159 - (0.1 - -2.25)) - 1.5))) / -2.25)
            v11 += -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    _ = len(v1)
    v2 = true
    v12 := GenS1{F0: 65535, F1: 1000}
    v2 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
