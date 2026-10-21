package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
func genFn0(p0 uint, p1 int32, p2 int16) int {
    v7 := GenS1{F0: -1, F1: 0, F2: 2147483647}
    if true {
        var v8 uint16 = 821
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -703
}

func main() {
    v1 := []uint32{87}
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 uint = 65535
    v5 := make(map[int]uint32)
    var v6 int = 32767
    if false {
        if true {
            if (1 == 431) {
                v2["a"] = "hello"
                v9 := GenS2{F0: v4, F1: -1, F2: -123}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 17; fi10++ {
                    v6 = int(v4)
                    fmt.Println(v6, v6, fi10)
                    _ = fi10
                }
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(v4)
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS2{F0: v4, F1: 32767, F2: 52}
                    v14 := GenS2{F0: v4, F1: 1, F2: 1}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
            v6 = 32767
        }
    }
    v1 = append(v1, 4294967295)
    var v15 int32 = 42
    var v16 int64 = 1
    var v17 string = ("a" + ("" + (((("" + "hello") + "a") + (("hello" + "a") + "\n\t")) + "\n\t")))
    var v18 int8 = int8(v6)
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int32 = 668
        if ((153 > uint8(v4)) || false) {
            v1 = append(v1, 111)
        } else {
            var v21 float64 = (1e10 * -2.25)
            if true {
                _ = len(v2)
                _ = len(v1)
            } else {
                v21 -= 0.0
                v6 -= -810
            }
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int64 = v16
        v24 := GenS2{F0: (^(v4)), F1: 179, F2: (v18 | (-(v18)))}
        var v25 uint32 = 655
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 int = -9223372036854775808
    var v27 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
