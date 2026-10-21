package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
func genFn0(p0 int8, p1 int, p2 uint) uint64 {
    p0 -= p0
    var v5 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint64) int32 {
    if (0 > -1) {
        p0 = 2
        var v6 bool = (1 == 1)
        _ = v6
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 61
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = -654
    v3 := make(map[string]bool)
    var v4 string = ""
    for fi7 := 0; fi7 < 13; fi7++ {
        if (18446744073709551615 != 169) {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 float64 = v1
                v10 := GenS1{F0: v4}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi7
    }
    fmt.Println(v2, v2)
    delete(v3, "")
    v1 -= float64(0.1)
    v11 := v3[v4]
    if true {
        v12 := GenS2{F0: 27}
        v4 = ("" + (((v4 + "hello") + ("hello" + "hello")) + (("hello" + ("\n\t" + (v4 + (("a" + (v4 + "a")) + "")))) + v4)))
        _ = v12
    }
    var v13 uint16 = 65535
    var v14 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
