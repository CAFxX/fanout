package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    p0 = (p0 / ((4294967295) | 1))
    _ = p0
    return 147
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int32 = -658
    fmt.Println(v4, v4)
    fmt.Println(v4)
    for fi5 := 0; fi5 < 10; fi5++ {
        if (0 >= 65535) {
            v6 := []int{(fi5 ^ (-(fi5))), (fi5 >> ((0) & 63)), 9223372036854775807, fi5}
            fmt.Println(v4, v1)
            _ = v6
        } else {
            fmt.Println(fi5, v1, v4)
            var v7 float32 = 0.1
            v8 := GenS1{F0: 18446744073709551615, F1: false}
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    v4 += 403
    v4 = (v4 & v4)
    for fi9 := 0; fi9 < 4; fi9++ {
        v4 += (v4 * 2147483647)
        _ = fi9
    }
    fmt.Println(v1, v1)
    v1 += v1
    switch 902 {
    case 0:
        v4 -= v4
    case 32768:
        var v10 string = ("" + ((("foo bar" + ("\n\t" + "\n\t")) + "a") + "hello"))
        _ = v10
    default:
        v4 += v4
    }
    fmt.Println(v4, v4)
    v11 := GenS2{F0: (v1 - v1), F1: (v4 + (v4 / (((v4 / ((v4) | 1))) | 1))), F2: ("\n\t" + ("foo bar" + ""))}
    var v12 int16 = 865
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
