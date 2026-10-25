package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint64) int {
    var v7 int16 = 42
    if true {
        v8 := []int{32767}
        for fi9 := 0; fi9 < 16; fi9++ {
            if (!false) {
                switch fi9 {
                case 42:
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v11 := GenS1{F0: 1}
                        v12 := GenS1{F0: -1}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
            }
            _ = fi9
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return -731
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"foo bar", "\n\t", "a", ""}
    var v3 uint32 = (v1 & (v1 ^ (v1 + (v1 | (^(v1))))))
    var v4 int8 = -1
    var v5 int = -2147483648
    var v6 float32 = float32(0.0)
    var v13 uint = 404
    var v14 int8 = v4
    v1 -= (v1 - v1)
    switch genFn0(608) {
    case 42:
        v13 = (v13 * v13)
    }
    fmt.Println(v5, v4, v13)
    var v15 uint16 = 255
    var v16 int32 = -98
    _ = len(v2)
    if false {
        if (v16 != v16) {
            var v17 string = (("a" + ("foo bar" + "hello")) + "\n\t")
            for fi18 := 0; fi18 < 17; fi18++ {
                v13 += (v13 / ((2) | 1))
                switch v17 {
                case "":
                    var v19 uint16 = v15
                    _ = v19
                case "a":
                    var v20 uint32 = (v3 >> ((65535) & 63))
                    _ = v20
                case "hello":
                    if (657 < (v3 >> ((494) & 63))) {
                        fmt.Println(v4, v14, v1)
                        v2 = append(v2, v17)
                        for fi21 := 0; fi21 < 4; fi21++ {
                            v22 := GenS1{F0: (^(v4))}
                            var v23 int32 = -2147483648
                            v23 += v16
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    }
                default:
                    v14 = 100
                }
                _ = fi18
            }
            fmt.Println(v13)
            _ = v17
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v14 -= (v4 / ((v14) | 1))
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
