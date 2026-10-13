package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := []int{-2147483648, 42}
    v4 := make(map[uint32]string)
    v2["a"] = true
    var v5 uint16 = 355
    v1 -= (^(v1))
    fmt.Println(v5, v5)
    v5 = 255
    var v6 bool = true
    for fi7 := 0; fi7 < 9; fi7++ {
        switch ("a" + "\n\t") {
        case "hello":
            for fi8 := 0; fi8 < 2; fi8++ {
                delete(v2, ("hello" + "hello"))
                var v9 uint16 = (v5 * (v5 << ((2) & 63)))
                v1 = 677
                if false {
                    fmt.Println(v6)
                    fmt.Println(fi7, v6)
                    fmt.Println(fi7)
                } else {
                    v10 := GenS1{F0: fi7, F1: -1}
                    v11 := GenS2{F0: (fi8 << ((42) & 63)), F1: 1}
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        default:
            switch "" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    var v12 int8 = 70
                    _ = v12
                case "\n\t":
                    v13 := GenS1{F0: (^(fi7)), F1: fi7}
                    _ = v13
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v14 := GenS1{F0: (fi7 * fi7), F1: fi7}
        _ = fi7
        _ = v14
    }
    v5 += v5
    fmt.Println(v6, v6, v6)
    v6 = v6
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(fi16, fi15, v6)
            fmt.Println(v5, fi15)
            v6 = false
            v5 = (v5 % (((v5 + (v5 >> ((0) & 63)))) | 1))
            v6 = v6
            _ = fi16
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v2[((("a" + "a") + ("foo bar" + ("foo bar" + "hello"))) + "a")] = (1 <= 42)
        v1 = v1
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
