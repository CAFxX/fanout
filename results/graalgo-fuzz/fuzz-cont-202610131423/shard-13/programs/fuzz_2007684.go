package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0() uint16 {
    v5 := []int{}
    var v6 uint32 = 65535
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 uint32, p1 string, p2 uint32) uint {
    var v7 int16 = 42
    if (true && ((!false) || (0 < 0))) {
        var v8 uint16 = 42
        fmt.Println(v7, v7)
        _ = v8
    }
    if ((false && (!(!(4294967295 == 2147483648)))) && (("a" < "a") || false)) {
        v9 := []int{127, -1, 42}
        v10 := GenS1{F0: 127, F1: 1}
        v11 := GenS1{F0: 128, F1: 2}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 35
    var v2 float32 = 0.0
    var v3 uint = 843
    v4 := make(map[string]int)
    v3 -= v3
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := v4["a"]
        _ = fi12
        _ = v13
    }
    if (v3 > (v3 | 4294967295)) {
        if (v3 <= 605) {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := v4["\n\t"]
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := GenS1{F0: (fi16 * fi16), F1: 0}
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
            var v18 int8 = 0
            _ = v18
        } else {
            if (!true) {
                if (1.5 != 0.1) {
                    v19 := []int{}
                    var v20 uint8 = 0
                    _ = v19
                    _ = v20
                }
            }
        }
        v2 -= float32((0.1 + 0.0))
    } else {
        v21 := GenS1{F0: 42, F1: (v3 - (v3 * v3))}
        v2 -= v2
        var v22 uint = genFn1((v1 << ((v3) & 63)), ((("a" + "") + "hello") + "\n\t"), v1)
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: -9, F1: genFn1(4294967295, "foo bar", 191)}
    v24 := v4[("\n\t" + "foo bar")]
    fmt.Println(v1)
    fmt.Println(v24, v24, v1)
    var v25 int = v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
