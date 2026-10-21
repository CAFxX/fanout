package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint8 = 223
    v2 := make(map[int]string)
    var v3 int16 = -120
    var v4 uint64 = 0
    var v5 uint = 2
    v6 := []uint32{}
    switch ("\n\t" + "a") {
    case "foo bar":
        var v7 int8 = 44
        _ = v7
    case "hello":
        v6 = append(v6, 2147483648)
    default:
        var v8 int16 = v3
        _ = v8
    }
    var v9 uint = (v5 << ((2) & 63))
    v10 := GenS1{F0: (v4 % (((v4 << (((v9 + v5)) & 63))) | 1)), F1: v4, F2: (v1 & 0)}
    v4 += (^(v4))
    var v11 int = 4294967295
    switch ("a" + (("" + "hello") + ("" + ("" + "foo bar")))) {
    case "\n\t":
        var v12 int8 = 127
        _ = v12
    case "":
        v3 = (v3 >> ((18446744073709551615) & 63))
    case "a":
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint16 = 255
            _ = fi13
            _ = v14
        }
    default:
        if true {
            v5 = (v9 / ((v9) | 1))
        } else {
            if true {
                var v15 string = "a"
                v16 := GenS1{F0: uint64(v9), F1: (v4 >> (((^(v9))) & 63)), F2: (v1 << (((v9 << ((v9) & 63))) & 63))}
                _ = len(v2)
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint64 = v4
                if (!(0.0 >= float64(1.5))) {
                    if (!("a" <= ("a" + "hello"))) {
                        v11 = fi17
                        for fi19 := 0; fi19 < 11; fi19++ {
                            fmt.Println(v11)
                            v5 = v9
                            _ = fi19
                        }
                        v5 += (v9 >> ((1) & 63))
                    } else {
                        var v20 string = (("a" + "a") + ("foo bar" + "foo bar"))
                        var v21 uint16 = 1
                        if (true || (1 != int(v3))) {
                            v22 := GenS2{F0: 2147483648, F1: 65535}
                            var v23 int64 = 9223372036854775807
                            var v24 int64 = 9223372036854775807
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            delete(v2, 9223372036854775807)
                        }
                        _ = v20
                        _ = v21
                    }
                    v25 := GenS1{F0: 536, F1: (v18 >> ((v5) & 63)), F2: v1}
                    _ = v25
                }
                _ = fi17
                _ = v18
            }
        }
    }
    v4 += (v4 - v4)
    v26 := GenS3{F0: (v1 + v1), F1: uint16(v1)}
    for fi27 := 0; fi27 < 5; fi27++ {
        switch v11 {
        case -1:
            v4 -= v4
        }
        _ = fi27
    }
    switch "hello" {
    case "":
        v6 = append(v6, uint32(v1))
    default:
        fmt.Println(v11, v9, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
