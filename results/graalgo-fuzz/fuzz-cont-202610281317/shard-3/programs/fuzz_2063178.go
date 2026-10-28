package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
func genFn0() string {
    v7 := []int{127, -2147483648, -76, -11}
    var v8 uint64 = 42
    v7 = append(v7, 450)
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            var v10 uint32 = 2147483648
            switch ("a" + "") {
            case "foo bar":
                _ = len(v7)
            case "hello":
                var v11 int64 = 9223372036854775807
                _ = v11
            }
            _ = v10
        } else {
            var v12 int16 = 42
            v13 := GenS1{F0: (v12 << ((255) & 63)), F1: 1, F2: 65535}
            v8 += v8
            _ = v12
            _ = v13
        }
        if (!true) {
            switch "a" {
            case "hello":
                var v14 uint8 = 255
                _ = v14
            }
        }
        v8 = (v8 & 771)
        _ = fi9
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 bool = true
    var v2 string = "a"
    v3 := []uint32{547}
    var v4 uint8 = 0
    var v5 float32 = (1e10 - float32(0.1))
    v6 := make(map[string]int)
    v4 -= (v4 * v4)
    var v15 int16 = 0
    var v16 uint8 = v4
    v2 = (v2 + v2)
    var v17 uint64 = 42
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v18 string = "a"
    v19 := GenS1{F0: 32767, F1: (v4 | (^(v16))), F2: 1}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
