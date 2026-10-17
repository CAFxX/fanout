package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool; F1 string; F2 int64 }

func main() {
    v1 := []uint32{666, 1, 4294967295, 2147483648}
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    var v4 uint16 = 65535
    _ = len(v1)
    v2 += (v2 >> ((65535) & 63))
    if true {
        v2 = (v2 * 4294967295)
        fmt.Println(v4)
        v2 += 4294967295
    } else {
        v5 := GenS1{F0: 543, F1: 13}
        _ = len(v1)
        _ = v5
    }
    v3[-2147483648] = (v2 | (v2 % ((2) | 1)))
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            if ((v4 ^ (^(v4))) <= 657) {
                fmt.Println(fi7, v2, v2)
                if false {
                    var v8 float32 = float32(-2.25)
                    _ = v8
                } else {
                    var v9 int64 = -1
                    v10 := GenS2{F0: false, F1: "hello", F2: v9}
                    if (v2 == v2) {
                        v9 = 0
                        var v11 uint8 = 2
                        v9 = v9
                        _ = v11
                    }
                    _ = v10
                    _ = v9
                }
            } else {
                if (1e10 >= (1.5 + ((0.0 * 1e10) + float64((0.0 - float32(-2.25)))))) {
                    var v12 float64 = (float64(1.5) + 3.14159)
                    var v13 int8 = -6
                    delete(v3, 127)
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint = 124
                    v4 -= 0
                    var v15 int16 = 868
                    _ = v14
                    _ = v15
                }
            }
            fmt.Println(v4, fi6)
            _ = fi7
        }
        _ = fi6
    }
    var v16 string = ("a" + "")
    v16 = v16
    switch v16 {
    case "\n\t":
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                v2 = v2
                _ = fi18
            }
            switch v16 {
            case "a":
                if ((((v2 | 2147483648) < (^(v2))) || (!(true && true))) || (true || false)) {
                    _ = len(v3)
                } else {
                    fmt.Println(v4)
                }
            case "foo bar":
                delete(v3, (fi17 - (fi17 * (fi17 & (^(fi17))))))
            default:
                var v19 int32 = 42
                _ = v19
            }
            _ = fi17
        }
    case "hello":
        v20 := GenS1{F0: 0, F1: -90}
        _ = v20
    case "":
        var v21 int = -178
        _ = v21
    }
    var v22 uint64 = 18446744073709551615
    fmt.Println(v2, v2, v4)
    var v23 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
