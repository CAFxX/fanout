package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 369
    var v3 int32 = v2
    var v4 int8 = 100
    var v5 uint32 = 65535
    v6 := GenS1{F0: "a", F1: (v2 - int32(v4))}
    if true {
        var v7 float32 = 1.5
        if ((!(!false)) && (false || true)) {
            v7 = (1.5 + ((float32(3.14159) + (v7 * v7)) + v7))
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := v1[(v5 / (((v5 << ((4294967295) & 63))) | 1))]
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 8; fi10++ {
                v7 = v7
                v2 = (v3 * (v3 + -902))
                _ = fi10
            }
        }
        _ = v7
    }
    var v11 uint = 42
    if true {
        delete(v1, v5)
    }
    var v12 int = -154
    var v13 uint64 = 1
    var v14 uint8 = 127
    if false {
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS1{F0: "a", F1: v3}
            if (true && false) {
                v17 := []int{}
                _ = v17
            }
            v18 := GenS2{F0: 32767, F1: v13}
            _ = fi15
            _ = v16
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 bool = ("a" >= ("foo bar" + "a"))
            v21 := GenS3{F0: (v4 <= v4)}
            if (v13 <= v13) {
                var v22 uint32 = v5
                _ = v22
            } else {
                if (((((0.1 - 0.0) - 0.0) + -2.25) * float64(0.0)) >= 0.0) {
                    var v23 float32 = 0.0
                    fmt.Println(v23, v11)
                    v23 -= ((0.0 / 1e10) + (0.0 - v23))
                    _ = v23
                }
                if v20 {
                    v24 := GenS4{F0: 255}
                    _ = v24
                } else {
                    var v25 float32 = 0.1
                    _ = v25
                }
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                switch ("hello" + ("" + "\n\t")) {
                case "\n\t":
                    v12 = 127
                case "foo bar":
                    v3 = v2
                default:
                    var v28 int = 165
                    _ = v28
                }
                v29 := GenS3{F0: (((fi26 * (fi27 | 32768)) <= fi27) || false)}
                _ = fi27
                _ = v29
            }
            v14 -= v14
            v2 -= v2
            _ = fi26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
