package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint) int32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float32 = 1.5
        v7 = 0.1
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    var v9 int32 = 79
    _ = p0
    _ = v8
    _ = v9
    return 80
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]int)
    var v3 uint16 = 680
    var v4 uint64 = 18446744073709551615
    v5 := []int{561, 1, 342, -926}
    if false {
        v10 := GenS1{F0: 255}
        _ = v10
    }
    fmt.Println(v3, v1)
    v3 += v3
    var v11 uint32 = uint32(v4)
    v12 := GenS2{F0: v3}
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                switch "foo bar" {
                case "hello":
                    _ = len(v5)
                    fmt.Println(fi13, fi14)
                case "a":
                    fmt.Println(v11, v11)
                default:
                    var v15 float32 = 1e10
                    _ = v15
                }
                v16 := v2[2]
                if false {
                    fmt.Println(v11, v1)
                    fmt.Println(v3)
                } else {
                    v5 = append(v5, (fi13 * 127))
                    var v17 uint16 = v3
                    _ = v17
                }
                _ = v16
            }
            _ = fi14
        }
        if false {
            fmt.Println(v11, fi13)
            switch "hello" {
            case "hello":
                var v18 bool = ((((1.5 / 1.5) * -2.25) == 0.1) && (false && false))
                _ = v18
            }
            v3 = (v3 % ((v3) | 1))
        } else {
            var v19 int = int(v11)
            v20 := GenS2{F0: v3}
            _ = v19
            _ = v20
        }
        fmt.Println(fi13, v11, v4)
        _ = fi13
    }
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    var v21 uint32 = 2
    switch 794 {
    case 127:
        if false {
            var v22 int32 = int32(v4)
            v22 += v22
            switch 35 {
            case -251:
                v3 = (^(v3))
            default:
                v23 := GenS3{F0: 0, F1: (v3 / (((v3 >> ((2) & 63))) | 1)), F2: 768}
                _ = v23
            }
            _ = v22
        }
    default:
        if true {
            fmt.Println(v21, v11, v11)
            v2[2] = 707
        }
    }
    var v24 string = "hello"
    v24 = "\n\t"
    v25 := v2[2]
    switch v25 {
    case 232:
        v3 = uint16(v21)
    }
    if ((v24 + (v24 + (v24 + ""))) <= v24) {
        for fi26 := 0; fi26 < 11; fi26++ {
            switch 32767 {
            case -9223372036854775808:
                v27 := v2[385]
                _ = v27
            case -762:
                fmt.Println(v11, v25)
            case -878:
                var v28 uint = 255
                _ = v28
            }
            v4 = uint64(v21)
            fmt.Println(v21, v1, v3)
            _ = fi26
        }
        switch ("a" + (v24 + "a")) {
        case "hello":
            v29 := GenS2{F0: (v3 >> ((65535) & 63))}
            _ = v29
        case "\n\t":
            v5 = append(v5, v25)
        default:
            var v30 string = "foo bar"
            _ = v30
        }
    } else {
        v4 = (v4 & 18446744073709551615)
        var v31 int32 = genFn0(0)
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            fmt.Println(v25, v4)
            _ = fi33
        }
        _ = fi32
    }
    if false {
        var v34 uint8 = 1
        _ = v34
    }
    v35 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
