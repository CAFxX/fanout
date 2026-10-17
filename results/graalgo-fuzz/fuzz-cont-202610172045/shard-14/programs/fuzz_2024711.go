package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 int32, p1 uint64, p2 int) uint {
    v5 := GenS1{F0: 42, F1: 255}
    var v6 uint64 = p1
    p2 = -9223372036854775808
    v7 := []int{(^(p2)), (p2 * (p2 >> ((42) & 63))), (p2 * 160)}
    if (!false) {
        var v8 int32 = (-(p0))
        var v9 uint8 = 1
        var v10 int8 = int8(v6)
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 46
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32767
    var v3 uint32 = 2
    v4 := make(map[string]bool)
    var v11 int = (v2 << ((220) & 63))
    v12 := GenS2{F0: -2147483648, F1: 2147483647}
    var v13 int = (^(v2))
    var v14 uint8 = 37
    if true {
        if false {
            if ((("" + (("" + "foo bar") + "hello")) + "\n\t") < "a") {
                var v15 int16 = 42
                _ = v15
            }
            var v16 int = v13
            if ((("foo bar" > "") && (false && true)) || true) {
                v2 = (-(v2))
                v14 = v14
                if true {
                    if false {
                        if (1.5 < ((float32((3.14159 * -2.25)) / 1.5) + (1e10 + 3.14159))) {
                            v17 := GenS2{F0: 122, F1: v3}
                            var v18 int16 = -1
                            var v19 bool = false
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        var v20 uint8 = v14
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 int64 = -9223372036854775808
                        switch "foo bar" {
                        case "\n\t":
                            v11 = (v16 - v16)
                        }
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v13)
                } else {
                    fmt.Println(v2, v16, v11)
                }
            } else {
                for fi23 := 0; fi23 < 17; fi23++ {
                    v13 -= (v13 << ((4294967295) & 63))
                    _ = fi23
                }
                v24 := GenS2{F0: -357, F1: v3}
                _ = v24
            }
            _ = v16
        } else {
            v11 = v2
            v13 -= (v11 * (v2 + v2))
            for fi25 := 0; fi25 < 8; fi25++ {
                v11 = (v13 >> ((546) & 63))
                v13 -= v13
                v26 := GenS3{F0: v14, F1: 42}
                _ = fi25
                _ = v26
            }
        }
    }
    var v27 float64 = float64((1e10 / (1.5 - 1e10)))
    fmt.Println(v11, v11)
    fmt.Println(v2)
    v28 := GenS1{F0: 65535, F1: 127}
    v13 = 32768
    v29 := GenS3{F0: (^(v14)), F1: 0}
    v1[4294967295] = v11
    var v30 string = "hello"
    v31 := GenS4{F0: 18446744073709551615, F1: (v30 + v30)}
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
