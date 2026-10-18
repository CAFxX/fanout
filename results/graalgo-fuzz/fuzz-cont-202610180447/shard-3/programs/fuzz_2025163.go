package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64) int8 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 208, F1: fi6, F2: "a"}
        v8 := GenS1{F0: 127, F1: fi6, F2: ("foo bar" + ("a" + ("a" + "hello")))}
        p0 -= p0
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: -2147483648, F1: 2147483647, F2: 589}
    v10 := []int{668, 4294967295}
    _ = p0
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 int = 32768
    var v2 float32 = (((1.5 / -2.25) / 3.14159) + (3.14159 - (1.5 / (0.0 + 1e10))))
    var v3 uint = 545
    v4 := make(map[string]bool)
    v5 := []int{}
    var v11 bool = false
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v3)
        v1 += 2147483647
        _ = fi12
    }
    var v13 float32 = v2
    v3 = 4294967295
    switch (v1 & (-(v1))) {
    case 0:
        var v14 int8 = 100
        _ = v14
    case 445:
        if v11 {
            v15 := GenS3{F0: -1, F1: 0}
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int8 = 127
                if (((1 >= -1) || (1 <= v3)) || v11) {
                    switch (fi16 << (((^(v3))) & 63)) {
                    case -513:
                        if (!false) {
                            switch "a" {
                            case "\n\t":
                                var v18 uint8 = 42
                                _ = v18
                            case "a":
                                var v19 uint16 = 65535
                                _ = v19
                            }
                        }
                    default:
                        v20 := GenS4{F0: fi16}
                        _ = v20
                    }
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                for fi22 := 0; fi22 < 5; fi22++ {
                    v13 += v2
                    var v23 uint16 = 249
                    _ = fi22
                    _ = v23
                }
                v24 := GenS5{F0: uint64(v3)}
                v3 += (v3 << ((18446744073709551615) & 63))
                _ = fi21
                _ = v24
            }
            _ = len(v4)
        }
    case 127:
        v2 += 0.1
    }
    v11 = (556 != 2147483648)
    if (0 == 2147483648) {
        var v25 uint16 = 2
        for fi26 := 0; fi26 < 13; fi26++ {
            v13 = v13
            v2 -= 0.1
            _ = fi26
        }
        var v27 bool = (65 <= -1)
        _ = v25
        _ = v27
    } else {
        if v11 {
            v3 = (v3 + (v3 ^ 771))
            if (true && true) {
                delete(v4, "hello")
            } else {
                fmt.Println(v2, v13)
            }
        } else {
            fmt.Println(v3)
            for fi28 := 0; fi28 < 4; fi28++ {
                v29 := GenS1{F0: 203, F1: (v1 | (-(fi28))), F2: "\n\t"}
                _ = fi28
                _ = v29
            }
            var v30 uint16 = 255
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            v5 = append(v5, (v1 % (((fi31 % (((v1 / ((fi31) | 1))) | 1))) | 1)))
            var v32 int = 4
            _ = fi31
            _ = v32
        }
    }
    fmt.Println(v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
