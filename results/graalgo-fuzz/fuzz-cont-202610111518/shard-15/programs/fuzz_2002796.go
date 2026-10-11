package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int16 = 1
    var v2 int = 32768
    var v3 string = (("\n\t" + "\n\t") + ("a" + "a"))
    v4 := []int{32768, (v2 & (-(v2))), 127, 1}
    v5 := make(map[uint32]int)
    v3 = ""
    v2 += 262
    var v6 int8 = 127
    var v7 int8 = -128
    var v8 float32 = 1.5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6)
    v3 = ""
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            v8 = 3.14159
            _ = fi9
        }
    } else {
        var v10 string = ((v3 + v3) + (v3 + (v3 + (("" + v3) + (("hello" + v3) + v3)))))
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: -556}
    fmt.Println(v2)
    if false {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = 2
            if false {
                v7 = v6
                v4 = append(v4, fi12)
                for fi14 := 0; fi14 < 5; fi14++ {
                    fmt.Println(v8, v2)
                    switch "\n\t" {
                    case "a":
                        if (false && (0 >= 1)) {
                            v1 = (v1 ^ 800)
                        } else {
                            v13 = v13
                        }
                    case "hello":
                        v8 = ((v8 / (v8 / (v8 + v8))) - (v8 * 1e10))
                    }
                    v15 := GenS2{F0: 0, F1: 176}
                    _ = fi14
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 16; fi16++ {
                    switch fi16 {
                    case 4294967295:
                        var v17 uint32 = 2147483647
                        _ = v17
                    case -2147483648:
                        v8 -= (-2.25 / 0.1)
                    case 127:
                        var v18 float32 = v8
                        _ = v18
                    }
                    _ = fi16
                }
                var v19 int32 = 0
                _ = v19
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
