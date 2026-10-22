package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 string }

func main() {
    v1 := make(map[string]string)
    var v2 string = ""
    var v3 uint = 18446744073709551615
    var v4 float32 = 1e10
    v5 := make(map[uint32]uint32)
    v4 += (0.0 - ((1e10 * ((v4 / v4) + v4)) - -2.25))
    var v6 uint64 = 1
    switch -271 {
    case 994:
        v6 = uint64(v3)
    case 0:
        var v7 float64 = 3.14159
        _ = v7
    case 128:
        v8 := GenS1{F0: 42}
        _ = v8
    }
    v9 := GenS1{F0: 127}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v6 = 822
            var v12 int = fi10
            _ = fi11
            _ = v12
        }
        if true {
            for fi13 := 0; fi13 < 4; fi13++ {
                switch (("a" + ((("hello" + v2) + v2) + ("a" + v2))) + "") {
                case "\n\t":
                    var v14 string = "a"
                    _ = v14
                case "foo bar":
                    if (255 != 42) {
                        var v15 uint16 = 165
                        v3 = (v3 & (^(v3)))
                        var v16 int64 = 1
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int8 = 42
                        v3 = (v3 & 255)
                        _ = v17
                    }
                default:
                    var v18 bool = (!true)
                    _ = v18
                }
                v19 := []int{-2147483648, fi13}
                switch (fi10 / ((fi13) | 1)) {
                case 32767:
                    v6 = v6
                case 57:
                    fmt.Println(v2, v6, fi10)
                case 9223372036854775807:
                    var v20 int64 = -840
                    _ = v20
                default:
                    var v21 float32 = float32(1e10)
                    _ = v21
                }
                _ = fi13
                _ = v19
            }
            fmt.Println(v3, fi10)
            fmt.Println(fi10)
        }
        _ = fi10
    }
    fmt.Println(v2, v3)
    fmt.Println(v4)
    var v22 uint64 = v6
    var v23 uint16 = 56
    v2 = v2
    for fi24 := 0; fi24 < 4; fi24++ {
        switch -2147483648 {
        case 32767:
            v23 -= v23
        default:
            for fi25 := 0; fi25 < 19; fi25++ {
                switch (v2 + v2) {
                case "hello":
                    v3 -= 18446744073709551615
                default:
                    v26 := GenS2{F0: 42, F1: ((-56 != -85) || false), F2: v2}
                    _ = v26
                }
                var v27 uint16 = (v23 << ((uint(v23)) & 63))
                fmt.Println(v6)
                _ = fi25
                _ = v27
            }
        }
        v2 = "\n\t"
        _ = fi24
    }
    v23 -= 402
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
