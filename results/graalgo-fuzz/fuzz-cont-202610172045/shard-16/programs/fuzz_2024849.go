package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = 42
    v3 := []int{-2147483648, -2147483648}
    var v4 int64 = 9223372036854775807
    _ = len(v3)
    var v5 uint16 = 915
    v1 = (v1 ^ v1)
    switch "" {
    case "":
        switch 2147483647 {
        case 9223372036854775807:
            var v6 uint8 = 0
            _ = v6
        default:
            v1 -= v1
        }
    case "foo bar":
        var v7 uint = 0
        _ = v7
    case "\n\t":
        var v8 int16 = -32768
        _ = v8
    }
    var v9 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = -68
    }
    var v10 uint64 = 122
    v3 = append(v3, 127)
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println(v1)
    var v11 uint = 255
    var v12 uint16 = 290
    if false {
        v4 += 9223372036854775807
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 163, F1: v12}
            _ = fi13
            _ = v14
        }
        if false {
            if ((true && (int64(v2) <= int64(v10))) && (-541 < 32767)) {
                if ((!(1 <= 519)) || true) {
                    v12 += 42
                    var v15 float32 = 3.14159
                    _ = v15
                }
                v16 := GenS2{F0: 65535, F1: uint8(v12)}
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                v1 = v1
                v18 := GenS1{F0: 255, F1: v12}
                for fi19 := 0; fi19 < 13; fi19++ {
                    if (42 >= 10) {
                        var v20 int16 = 1000
                        _ = v20
                    }
                    if ((v11 ^ (v11 >> ((65535) & 63))) == 962) {
                        var v21 string = "\n\t"
                        if false {
                            var v22 uint32 = 415
                            switch v21 {
                            case "a":
                                v10 = (^(v10))
                            case "hello":
                                v1 = v1
                            case "foo bar":
                                fmt.Println(v4, fi19, v9)
                            default:
                                v10 += 42
                            }
                            _ = v22
                        } else {
                            var v23 bool = true
                            var v24 int16 = 0
                            var v25 int8 = 1
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v11 += v11
                        _ = v21
                    } else {
                        v10 += v10
                    }
                    v10 = v10
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
            v9 = v9
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
