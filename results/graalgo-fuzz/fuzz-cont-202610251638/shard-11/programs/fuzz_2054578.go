package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []int{4294967295, 143, 0}
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    var v4 int8 = 42
    var v5 int8 = v4
    var v6 int8 = v4
    fmt.Println(v5, v5, v6)
    if ((1e10 * 1.5) == 1e10) {
        if (-644 <= -896) {
            v6 -= -35
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 string = ("a" + "")
                v5 += -1
                _ = fi7
                _ = v8
            }
            v6 = -128
        }
    }
    fmt.Println(v6, v2, v5)
    fmt.Println(v6, v2)
    _ = len(v3)
    v9 := GenS1{F0: 0}
    var v10 int64 = 1
    var v11 int16 = 1
    var v12 int8 = v5
    var v13 uint32 = v2
    fmt.Println(v5)
    var v14 int = 128
    if (!((-909 <= -478) && (!(127 <= 2)))) {
        v12 = v12
        v2 += 2147483648
    } else {
        if (!false) {
            if true {
                var v15 int64 = (v10 | (v10 - v10))
                _ = v15
            } else {
                var v16 int64 = v10
                for fi17 := 0; fi17 < 4; fi17++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v11, v12, v12)
                    var v18 bool = (4294967295 != 457)
                    _ = fi17
                    _ = v18
                }
                v19 := GenS1{F0: v16}
                _ = v16
                _ = v19
            }
            v20 := GenS2{F0: ((v10 % ((v10) | 1)) == v10), F1: v13}
            for fi21 := 0; fi21 < 17; fi21++ {
                switch ("foo bar" + ("a" + "a")) {
                case "a":
                    var v22 bool = false
                    _ = v22
                case "":
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := GenS1{F0: v10}
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi21
            }
            _ = v20
        } else {
            v11 = (v11 << ((2) & 63))
            v25 := GenS3{F0: (v11 * v11), F1: 864}
            for fi26 := 0; fi26 < 14; fi26++ {
                v13 += (v13 ^ 451)
                _ = fi26
            }
            _ = v25
        }
        var v27 uint16 = 42
        _ = v27
    }
    v28 := GenS4{F0: v13, F1: -2147483648}
    v4 = v12
    v6 -= v6
    fmt.Println(v6, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
