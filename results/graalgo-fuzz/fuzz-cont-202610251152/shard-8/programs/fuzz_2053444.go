package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool) int16 {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = (4294967295 != 843)
    _ = p0
    return -32768
}
func genFn1(p0 int, p1 uint16, p2 uint8) bool {
    if (true || ((p1 & (p1 - (^(p1)))) > p1)) {
        v7 := []int{p0, p0}
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch (-(fi8)) {
            case 354:
                v9 := []int{}
                _ = v9
            }
            var v10 float32 = (0.0 + 1e10)
            var v11 uint = 320
            _ = fi8
            _ = v10
            _ = v11
        }
        p2 = (p2 << ((0) & 63))
        if (!(42 <= 2)) {
            switch (p0 + p0) {
            case -9223372036854775808:
                for fi12 := 0; fi12 < 17; fi12++ {
                    p0 = fi12
                    v13 := GenS1{F0: -910, F1: -9223372036854775808, F2: ("a" + "hello")}
                    _ = fi12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                if false {
                    fmt.Println(p0, p1, p1)
                    p2 = p2
                    v15 := []int{1, p0, (fi14 * (p0 >> ((65535) & 63))), p0}
                    _ = v15
                } else {
                    v16 := GenS2{F0: 1}
                    var v17 bool = ((2147483648 != 2147483648) && ((-(p0)) >= 127))
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 uint64 = 42
    var v2 uint8 = 255
    var v3 int8 = 103
    var v4 uint16 = 2
    v5 := make(map[int]bool)
    v6 := []string{}
    delete(v5, int(v1))
    v2 = v2
    var v18 float32 = float32(1.5)
    v19 := v5[4294967295]
    if (1e10 == float64((v18 / (1.5 * (v18 - v18))))) {
        v4 -= (v4 / ((uint16(v2)) | 1))
        var v20 bool = v19
        var v21 uint32 = 4294967295
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 602, F1: 335, F2: "hello"}
    _ = len(v6)
    v1 = 788
    var v23 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
