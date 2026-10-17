package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 uint, p2 uint16) int64 {
    p1 += 255
    v4 := GenS1{F0: 42, F1: (p2 >> ((p1) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -286
}
func genFn1(p0 uint32) uint8 {
    v5 := GenS1{F0: p0, F1: 1}
    v6 := GenS1{F0: 4294967295, F1: 853}
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!(0 < 18446744073709551615))
    v3 := []string{"hello", ("\n\t" + "hello"), "\n\t", "hello"}
    for fi7 := 0; fi7 < 18; fi7++ {
        switch fi7 {
        case -812:
            fmt.Println(fi7, fi7)
        }
        v8 := GenS1{F0: 2147483647, F1: 42}
        var v9 uint8 = 121
        _ = fi7
        _ = v8
        _ = v9
    }
    v2 = v2
    var v10 uint16 = 862
    var v11 uint32 = 42
    v10 = v10
    _ = len(v1)
    if (0.0 >= float32(0.1)) {
        for fi12 := 0; fi12 < 2; fi12++ {
            _ = len(v1)
            var v13 int64 = -9223372036854775808
            if ((((("" + ("" + ("\n\t" + "\n\t"))) + "a") + ("\n\t" + "\n\t")) + "\n\t") == "foo bar") {
                delete(v1, (v11 / ((v11) | 1)))
                v2 = (775 > 2)
            } else {
                var v14 uint8 = 127
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        switch 128 {
        case 2147483647:
            v15 := GenS2{F0: -41}
            _ = v15
        case 1:
            var v16 float32 = (1e10 * 0.1)
            _ = v16
        default:
            v2 = ((-1 >= 83) || ((32767 >= 1000) && true))
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int8 = 26
            _ = fi17
            _ = v18
        }
    } else {
        v2 = v2
        v11 = v11
        v2 = ((72 == 1) && ((4294967295 < 65535) && ((!v2) || (-413 <= -2147483648))))
    }
    fmt.Println(v10)
    v11 = (v11 % (((v11 / ((v11) | 1))) | 1))
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int32 = 1
        v2 = (!v2)
        for fi21 := 0; fi21 < 6; fi21++ {
            v20 += (v20 | (v20 << ((255) & 63)))
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    switch "foo bar" {
    case "foo bar":
        v11 = (v11 ^ (v11 + (v11 % ((18) | 1))))
    default:
        var v22 int32 = -212
        _ = v22
    }
    fmt.Println(v10, v10)
    v11 += (v11 | (v11 % ((v11) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
