package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
func genFn0(p0 int16, p1 int32, p2 int64) uint32 {
    p1 = (p1 >> ((180) & 63))
    p0 += p0
    p0 = (p0 - (p0 / ((p0) | 1)))
    v5 := []int{}
    p1 -= (p1 % (((p1 ^ (p1 - p1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 414
}
func genFn1(p0 int64, p1 uint, p2 string) int8 {
    if (42 <= 1) {
        if false {
            v6 := GenS1{F0: -1, F1: 84, F2: p2}
            _ = v6
        } else {
            var v7 int16 = 309
            _ = v7
        }
    } else {
        var v8 uint = 18446744073709551615
        v9 := []int{-1, 0, 4294967295}
        _ = v8
        _ = v9
    }
    p0 = p0
    var v10 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 100
}

func main() {
    v1 := []string{("hello" + "a"), "hello"}
    var v2 uint8 = 255
    v3 := []int{1, 1, 32767, 178}
    var v4 uint64 = uint64(v2)
    v1 = append(v1, ("foo bar" + ("\n\t" + (("a" + ("" + "")) + ("a" + ((("a" + "") + ((("foo bar" + ("hello" + "foo bar")) + "foo bar") + "")) + ("foo bar" + "foo bar")))))))
    var v11 float64 = float64(1.5)
    v4 += (v4 | v4)
    v12 := GenS2{F0: "", F1: 399, F2: 490}
    if ("a" <= "foo bar") {
        v13 := GenS2{F0: "a", F1: -1, F2: 633}
        var v14 float64 = (v11 + (v11 / v11))
        if (v2 > v2) {
            v11 = (v14 / v11)
            for fi15 := 0; fi15 < 6; fi15++ {
                v14 = 1e10
                _ = fi15
            }
            switch -2147483648 {
            case 674:
                if (v14 >= -2.25) {
                    var v16 int = -2147483648
                    v17 := GenS2{F0: "foo bar", F1: 9223372036854775807, F2: 918}
                    _ = v16
                    _ = v17
                }
            case 4294967295:
                v2 = v2
            }
        } else {
            v14 = (float64(-2.25) - (v11 + (((v11 + ((-2.25 - v11) - 1e10)) * 0.1) / (0.0 * (1.5 - v14)))))
        }
        _ = v13
        _ = v14
    } else {
        fmt.Println(v2, v4, v11)
        var v18 int16 = -32768
        v11 += v11
        _ = v18
    }
    v11 = (v11 + (3.14159 + 1.5))
    fmt.Println(v2)
    for fi19 := 0; fi19 < 16; fi19++ {
        _ = len(v3)
        v11 = v11
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
