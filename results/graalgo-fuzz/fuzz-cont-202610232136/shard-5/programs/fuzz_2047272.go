package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 int32, p1 uint32, p2 uint32) int64 {
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 -= p0
        _ = fi7
    }
    v8 := GenS1{F0: 100, F1: 18446744073709551615, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -527
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    var v3 bool = false
    v4 := []int{v2, (v2 - 947), (^(v2)), (v2 & (v2 % ((v2) | 1)))}
    v5 := make(map[string]uint32)
    var v6 int8 = -1
    if (2 > 71) {
        v6 = 42
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v4) > 0 {
            v4[0] = v2
        }
        _ = fi9
    }
    fmt.Println(v2, v2)
    if false {
        var v10 float32 = 1.5
        _ = v10
    }
    v5[(("foo bar" + ("a" + (("a" + "") + "hello"))) + "")] = 0
    var v11 uint = 2
    v12 := GenS1{F0: (v6 & (v6 | (v6 >> (((v11 | v11)) & 63)))), F1: v11, F2: 65535}
    v11 += 255
    if (0.1 == -2.25) {
        if v3 {
            var v13 uint32 = 2
            if (v13 <= v13) {
                v14 := GenS1{F0: (-(v6)), F1: (^(v11)), F2: 255}
                if ((1.5 + 3.14159) == ((((0.0 - 0.0) * 0.0) - 3.14159) - 1e10)) {
                    var v15 uint8 = 0
                    _ = v15
                }
                switch (v2 + v2) {
                case -465:
                    var v16 int8 = (-(v6))
                    _ = v16
                }
                _ = v14
            } else {
                fmt.Println(v2, v13, v3)
                _ = len(v4)
            }
            _ = v13
        } else {
            var v17 float32 = -2.25
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
