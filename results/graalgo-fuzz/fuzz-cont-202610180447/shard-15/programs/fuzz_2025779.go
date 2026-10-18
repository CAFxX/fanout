package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint) uint {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = -344
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(fi8)
        var v9 float32 = (3.14159 / (1.5 * 1.5))
        _ = fi8
        _ = v9
    }
    p0 = (p0 >> ((18446744073709551615) & 63))
    _ = p0
    _ = v7
    return (p0 + p0)
}

func main() {
    v1 := []uint32{}
    var v2 bool = ((1e10 / (float32(float64(((-2.25 + (0.1 + (1.5 / 1.0))) * 0.1))) * (float32(1.5) * ((0.0 + (((-2.25 - (0.1 / 1.0)) / (3.14159 - 0.0)) - ((0.0 + -2.25) * 1.5))) / 0.1)))) <= 0.1)
    v3 := []string{"\n\t", ("a" + "a"), ((("foo bar" + ("\n\t" + (("foo bar" + "\n\t") + "a"))) + ("a" + "")) + "")}
    var v4 uint16 = 0
    v5 := []string{}
    v6 := make(map[uint32]int)
    switch 2147483647 {
    case -125:
        fmt.Println(v4, v4)
    default:
        v2 = false
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v6, 829)
    var v10 uint = 1
    switch -9223372036854775808 {
    case 0:
        v11 := GenS1{F0: 9223372036854775807}
        _ = v11
    case 618:
        if (!false) {
            v12 := GenS1{F0: 966}
            var v13 uint8 = uint8(v10)
            v6[509] = 0
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 1}
            _ = v14
        }
    default:
        var v15 uint16 = (v4 ^ v4)
        _ = v15
    }
    _ = len(v6)
    v4 += 0
    v10 += (v10 ^ v10)
    for fi16 := 0; fi16 < 11; fi16++ {
        switch "a" {
        case "a":
            v5 = append(v5, (("a" + "hello") + "hello"))
        case "hello":
            var v17 uint8 = 63
            _ = v17
        default:
            v18 := GenS1{F0: 9223372036854775807}
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS1{F0: -1}
    v10 -= 65535
    v10 += genFn0(v10)
    var v20 int = 4294967295
    v20 += v20
    v2 = (false || ((v10 % (((^(v10))) | 1)) != 42))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
