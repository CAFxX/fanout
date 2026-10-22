package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int }

func main() {
    v1 := []uint32{42, 804, 42}
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    var v4 uint64 = 2
    v5 := []string{"foo bar", "hello", ("\n\t" + ("a" + "a")), "\n\t"}
    var v6 int16 = 1000
    v7 := GenS1{F0: 952, F1: (((!(4294967295 >= 2)) && ((!true) || ((0 != -1) && (245 < 2147483648)))) || false), F2: 930}
    var v8 string = "foo bar"
    v6 = (v6 % ((v6) | 1))
    v9 := GenS2{F0: 0, F1: 2147483648, F2: 32767}
    if ((v6 ^ v6) != v6) {
        v6 -= 0
        v8 = v8
    }
    fmt.Println(v2)
    v8 = (v8 + "foo bar")
    v2 += 981
    switch 42 {
    case 1:
        v10 := GenS2{F0: 114, F1: 4294967295, F2: 0}
        _ = v10
    default:
        var v11 float32 = 1.5
        _ = v11
    }
    delete(v3, 4294967295)
    if (844 <= 1) {
        switch -1 {
        case 32768:
            if (!false) {
                v12 := GenS2{F0: 0, F1: 2, F2: 32767}
                var v13 int = int(v2)
                _ = v12
                _ = v13
            }
        case 4294967295:
            var v14 uint64 = v4
            _ = v14
        case -200:
            switch 127 {
            case 127:
                v4 -= v4
            default:
                v15 := GenS2{F0: -128, F1: uint32(v2), F2: -2147483648}
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 int8 = 1
                switch ((v8 + (v8 + v8)) + v8) {
                case "hello":
                    fmt.Println(v8, v17)
                case "a":
                    var v18 uint64 = 89
                    _ = v18
                case "\n\t":
                    var v19 int16 = v6
                    _ = v19
                }
                _ = len(v1)
                _ = fi16
                _ = v17
            }
        }
    } else {
        v2 -= v2
        v20 := GenS2{F0: 0, F1: 2, F2: 128}
        v6 = v6
        _ = v20
    }
    v3[2] = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
