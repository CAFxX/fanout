package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0(p0 uint32) uint {
    p0 += p0
    var v4 bool = (3.14159 == 3.14159)
    _ = p0
    _ = v4
    return 299
}
func genFn1(p0 int8, p1 uint64) uint64 {
    var v5 uint32 = 2
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{32768}
        _ = fi7
        _ = v8
    }
    v9 := []int{128, 2147483647, 2147483647}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 int8 = -1
    var v10 string = "a"
    v10 = ((((v10 + "a") + (v10 + (v10 + ("foo bar" + v10)))) + v10) + ((v10 + (v10 + ("foo bar" + ("" + (v10 + "a"))))) + "hello"))
    var v11 int = -1
    v11 += v11
    v11 = (v11 << ((18446744073709551615) & 63))
    var v12 string = (v10 + (v10 + (((v10 + "a") + "") + v10)))
    var v13 uint16 = 2
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint64 = 725
        v17 := []int{fi15}
        v16 -= 42
        _ = fi15
        _ = v16
        _ = v17
    }
    if (722 != 42) {
        if (true && false) {
            v18 := []int{}
            v14 += (v14 | (^(v14)))
            _ = len(v18)
            _ = v18
        } else {
            var v19 bool = true
            v20 := []int{}
            _ = v19
            _ = v20
        }
        if (true || (false && false)) {
            v21 := GenS1{F0: true}
            if (0.0 == 0.0) {
                _ = len(v2)
            }
            var v22 string = (v12 + (v10 + (((v10 + "foo bar") + v12) + v10)))
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 8; fi23++ {
                switch "" {
                case "foo bar":
                    switch 2147483647 {
                    case 1:
                        v24 := v2[""]
                        _ = v24
                    default:
                        fmt.Println(v3, v11)
                    }
                case "":
                    for fi25 := 0; fi25 < 3; fi25++ {
                        fmt.Println(v13)
                        fmt.Println(v12, v13, fi25)
                        v10 = ""
                        _ = fi25
                    }
                case "hello":
                    v11 = (^(fi23))
                default:
                    switch v12 {
                    case "hello":
                        fmt.Println(v13, fi23)
                    }
                }
                v13 = v13
                _ = fi23
            }
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            v12 = (("" + "a") + "foo bar")
            fmt.Println(fi26, v12, v13)
            v13 = v13
            _ = fi26
        }
    }
    delete(v1, "\n\t")
    v27 := GenS2{F0: 42, F1: (v13 / ((625) | 1)), F2: v13}
    fmt.Println(v13, v3)
    v28 := GenS1{F0: (-128 > (v3 - 46))}
    v29 := GenS2{F0: 42, F1: 811, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
