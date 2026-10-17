package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
func genFn0(p0 int64) uint8 {
    var v5 int = 690
    v6 := []int{(v5 + 2147483647)}
    v7 := GenS1{F0: 42, F1: ((272 == 272) && (-1 > -1))}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 86
}

func main() {
    var v1 int = -33
    v2 := make(map[uint32]string)
    v3 := []uint32{65535}
    var v4 int = v1
    var v8 uint8 = 179
    _ = len(v3)
    v3 = append(v3, 2)
    if (uint64(v8) < 18446744073709551615) {
        v9 := GenS2{F0: 573, F1: 490}
        _ = v9
    }
    switch v1 {
    case 572:
        var v10 uint8 = v8
        _ = v10
    case 4294967295:
        fmt.Println(v1, v8, v4)
    }
    var v11 uint = 42
    var v12 uint64 = 422
    v4 = (v1 + v4)
    switch 4294967295 {
    case 32767:
        var v13 float64 = 3.14159
        _ = v13
    default:
        if (((0.1 * (1.5 + (3.14159 * 1.5))) / 0.1) == 3.14159) {
            v12 -= v12
            switch ("a" + ("" + ((((("a" + "\n\t") + ("hello" + ("" + "\n\t"))) + ("\n\t" + "hello")) + "a") + ("" + (("\n\t" + (("a" + "foo bar") + "hello")) + "foo bar"))))) {
            case "hello":
                v12 = v12
            case "\n\t":
                for fi14 := 0; fi14 < 2; fi14++ {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 uint16 = 42
                        fmt.Println(fi15, v1)
                        _ = len(v3)
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
            }
            if true {
                _ = len(v3)
                if true {
                    if false {
                        _ = len(v3)
                        v11 = 255
                        v1 = v1
                    } else {
                        delete(v2, 4294967295)
                    }
                    v17 := GenS3{F0: 247, F1: -9223372036854775808, F2: false}
                    _ = v17
                }
            } else {
                v3 = append(v3, 42)
                v12 = v12
                var v18 int = 2147483647
                _ = v18
            }
        }
    }
    fmt.Println(v11)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
