package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint8, p1 int16) int64 {
    p1 = -26
    p1 = -892
    var v5 float32 = 0.1
    p0 = (p0 - 127)
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 int8, p1 string) uint64 {
    fmt.Println(p1, p1)
    p0 = (^(p0))
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 += (p0 | p0)
        p1 = (p1 + "foo bar")
        v7 := GenS1{F0: 42, F1: 1}
        _ = fi6
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = p1
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483648, 87, 737, 244}
    var v4 int = 45
    v9 := GenS2{F0: (v1 >> ((287) & 63)), F1: 321}
    v10 := GenS2{F0: v1, F1: 0}
    v3 = append(v3, 382)
    var v11 int32 = v1
    v11 += v1
    switch v4 {
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case 2147483647:
        if (42 != 229) {
            var v12 bool = (genFn0(1, 645) > 42)
            switch v4 {
            case -2147483648:
                var v13 bool = true
                _ = v13
            case 479:
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 float64 = 0.1
                    _ = len(v2)
                    _ = fi14
                    _ = v15
                }
            case 32767:
                var v16 int32 = (-(v1))
                _ = v16
            }
            _ = v12
        }
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 = v11
            _ = fi17
        }
    }
    fmt.Println(v11, v11, v1)
    v1 = (v11 << ((65535) & 63))
    for fi18 := 0; fi18 < 12; fi18++ {
        if ((!(0 == 65535)) || (true && (!((false && true) && (0 == 4294967295))))) {
            v1 = 1
        }
        _ = fi18
    }
    switch ("hello" + "") {
    case "a":
        v3 = append(v3, 2)
    case "hello":
        v1 -= (v11 * v11)
    default:
        switch "\n\t" {
        case "hello":
            switch ("" + "hello") {
            case "":
                if true {
                    fmt.Println(v4, v1, v11)
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                }
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    fmt.Println(v11, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
