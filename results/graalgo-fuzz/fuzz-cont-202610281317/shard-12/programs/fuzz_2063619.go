package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32) int32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{9223372036854775807, fi5, (fi5 * fi5)}
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = v7
    return (p0 ^ p0)
}
func genFn1(p0 int64, p1 uint, p2 string) uint32 {
    fmt.Println(p0, p1)
    p0 += (p0 ^ (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[uint32]uint32)
    var v4 int32 = -1
    var v8 int8 = 127
    var v9 int16 = 1
    v4 = (v4 << ((443) & 63))
    v10 := GenS1{F0: v9}
    v4 = int32(v9)
    switch "a" {
    case "\n\t":
        for fi11 := 0; fi11 < 17; fi11++ {
            v4 = genFn0(int32(fi11))
            var v12 bool = (-2.25 != (((-2.25 / 3.14159) / 1.0) * (3.14159 * (1e10 + float64(-2.25)))))
            v13 := []int{fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        v4 = -2147483648
    }
    var v14 int = 679
    v15 := GenS1{F0: v9}
    var v16 uint16 = 245
    v4 = v4
    if false {
        fmt.Println(v8, v4, v2)
        if ((0.1 * -2.25) <= -2.25) {
            var v17 float32 = ((0.1 - (((0.1 / ((3.14159 - (float32(3.14159) / -2.25)) / 3.14159)) / float32(0.1)) / 1.0)) / 0.1)
            switch ("a" + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) {
            case "":
                v16 -= (v16 ^ 2)
            case "\n\t":
                switch (v14 >> ((4294967295) & 63)) {
                case 454:
                    v2 = (v2 % ((v2) | 1))
                case 9223372036854775807:
                    var v18 int = (v14 % (((^(v14))) | 1))
                    _ = v18
                default:
                    v19 := GenS1{F0: 32767}
                    _ = v19
                }
            case "a":
                v16 -= (v16 - uint16(v2))
            }
            _ = v17
        }
    } else {
        v20 := GenS1{F0: (^(v9))}
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
