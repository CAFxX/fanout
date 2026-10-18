package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint = 994
    var v2 bool = true
    v3 := []uint32{65535}
    v4 := []int{}
    var v5 int32 = 2147483647
    v6 := GenS1{F0: 0, F1: 65535}
    var v7 uint = v1
    var v8 string = "hello"
    var v9 bool = v2
    var v10 float32 = 3.14159
    var v11 float64 = (1.5 + 0.1)
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                if (!false) {
                    var v14 uint8 = 255
                    _ = v14
                }
                if false {
                    switch (fi12 & int(v5)) {
                    case 355:
                        if false {
                            switch "a" {
                            case "":
                                var v15 uint64 = 42
                                _ = v15
                            }
                            var v16 int32 = v5
                            v10 = ((float32(((0.1 * -2.25) + v11)) + 1.5) - (v10 + -2.25))
                            v10 -= (3.14159 - (float32((v11 + float64((v10 * v10)))) + -2.25))
                            _ = v16
                        } else {
                            v17 := GenS2{F0: ((("foo bar" + v8) + v8) + "hello"), F1: 297, F2: 42}
                            _ = v17
                        }
                        v3 = append(v3, uint32(v7))
                    }
                    var v18 uint16 = 274
                    _ = v18
                } else {
                    v10 = (1e10 - v10)
                    v5 = v5
                }
                fmt.Println(v1)
            } else {
                var v19 int = (fi12 | fi13)
                v20 := GenS3{F0: 42}
                var v21 int8 = 115
                _ = v19
                _ = v20
                _ = v21
            }
            v22 := GenS2{F0: v8, F1: -32768, F2: 18446744073709551615}
            _ = fi13
            _ = v22
        }
        v11 -= (3.14159 * 3.14159)
        for fi23 := 0; fi23 < 12; fi23++ {
            v24 := GenS4{F0: 42, F1: (-1 != v5)}
            var v25 uint64 = 0
            var v26 int = (fi12 ^ fi12)
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi12
    }
    v8 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
