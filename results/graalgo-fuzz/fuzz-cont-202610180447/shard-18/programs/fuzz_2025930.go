package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 bool }

func main() {
    var v1 uint = 626
    var v2 float32 = 1e10
    var v3 int8 = 1
    var v4 int8 = v3
    var v5 int = 127
    v6 := make(map[string]bool)
    var v7 bool = ((3.14159 > (3.14159 / (0.1 - -2.25))) && (true || ((v3 / ((v3) | 1)) <= 127)))
    v2 = float32(1.5)
    var v8 int16 = 32767
    v3 = 127
    for fi9 := 0; fi9 < 19; fi9++ {
        if (v2 > v2) {
            for fi10 := 0; fi10 < 13; fi10++ {
                switch "foo bar" {
                case "a":
                    delete(v6, "foo bar")
                case "":
                    v11 := GenS1{F0: -555, F1: 866, F2: (v8 ^ (v8 ^ (-(v8))))}
                    _ = v11
                }
                var v12 float32 = (1e10 / 0.1)
                v3 = (v3 & (v4 + -1))
                _ = fi10
                _ = v12
            }
        } else {
            v8 += (v8 - (v8 * (-(v8))))
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 9223372036854775807, F1: 489, F2: (^(v8))}
                if v7 {
                    v15 := GenS1{F0: -1, F1: uint16(v1), F2: v8}
                    var v16 uint16 = 255
                    v1 = 1
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint8 = 18
                    if true {
                        var v18 string = "foo bar"
                        _ = v18
                    }
                    _ = v17
                }
                var v19 bool = false
                _ = fi13
                _ = v14
                _ = v19
            }
            var v20 uint16 = 0
            _ = v20
        }
        v2 += v2
        v1 = (v1 + (v1 << ((v1) & 63)))
        _ = fi9
    }
    switch (("" + "") + (("hello" + "a") + "a")) {
    case "a":
        switch (v5 - 9223372036854775807) {
        case 0:
            v5 = (v5 - 1)
        default:
            switch "a" {
            case "":
                v4 = v3
            case "hello":
                v21 := GenS2{F0: (237 <= 145), F1: 42, F2: (42 != (v1 ^ (v1 >> ((v1) & 63))))}
                _ = v21
            }
        }
    case "\n\t":
        v5 -= (v5 >> (((v1 & (v1 / ((v1) | 1)))) & 63))
    case "hello":
        var v22 uint = v1
        _ = v22
    default:
        var v23 float64 = float64(float32((0.1 / 0.1)))
        _ = v23
    }
    v4 = v4
    var v24 int32 = 0
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := []int{v5, (v5 * (v5 & 2147483647)), v5, int(v24)}
        var v27 int16 = v8
        v4 -= v4
        _ = fi25
        _ = v26
        _ = v27
    }
    v4 = 127
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
