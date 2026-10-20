package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
func genFn0() uint64 {
    v4 := GenS1{F0: true}
    var v5 string = "hello"
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v5)
        switch "\n\t" {
        case "foo bar":
            switch fi6 {
            case 927:
                fmt.Println(fi6, fi6)
            case 128:
                var v7 uint64 = 0
                _ = v7
            case 1:
                if (569 >= 65535) {
                    fmt.Println(v5, fi6)
                    fmt.Println(fi6, fi6)
                }
            }
        case "a":
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint = 1
                _ = fi8
                _ = v9
            }
        default:
            v5 = ("foo bar" + (v5 + v5))
        }
        v10 := GenS1{F0: (255 >= 4294967295)}
        _ = fi6
        _ = v10
    }
    _ = v4
    _ = v5
    return 756
}
func genFn1(p0 uint16, p1 string) uint16 {
    fmt.Println(p1, p1)
    var v11 uint = 881
    _ = p0
    _ = p1
    _ = v11
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = ((1e10 + 3.14159) - (float32(((float64((1e10 * 0.0)) * 3.14159) * (((0.1 * 3.14159) + -2.25) / (-2.25 / 0.1)))) + 1.5))
    v3 := make(map[int]uint32)
    v1[2] = "a"
    var v12 int = -2147483648
    switch "\n\t" {
    case "hello":
        if true {
            v2 -= 0.0
            v12 = (v12 << ((18446744073709551615) & 63))
        }
    default:
        v12 = (-(v12))
    }
    v13 := v1[2147483648]
    v14 := v3[(v12 - 2147483647)]
    v15 := GenS1{F0: false}
    var v16 uint64 = uint64(v14)
    v17 := []int{v12, v12, (v12 ^ (^(v12)))}
    v18 := GenS1{F0: (32767 < 0)}
    var v19 float64 = float64((v2 + -2.25))
    v20 := GenS2{F0: (v14 & (v14 & v14))}
    var v21 string = (v13 + "")
    if true {
        var v22 int32 = -2147483648
        _ = v22
    }
    v17 = append(v17, (v12 % ((v12) | 1)))
    var v23 float64 = 3.14159
    v21 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
