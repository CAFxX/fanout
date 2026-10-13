package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 int32, p1 int8, p2 int64) uint32 {
    p2 -= p2
    p2 = p2
    var v7 string = ("\n\t" + (("a" + ("" + "foo bar")) + "foo bar"))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint8 = 0
    var v4 int8 = 0
    var v5 int64 = -9223372036854775808
    var v6 bool = ((true && v1) && v1)
    switch "foo bar" {
    case "hello":
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint8 = 2
                if (!(!v6)) {
                    v10 += (^(v3))
                    v11 := []int{127, (fi8 - fi9), (-(fi8)), (^(fi8))}
                    var v12 int32 = 1
                    _ = v11
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
            v6 = (-316 < fi8)
            _ = fi8
        }
    case "foo bar":
        v3 += v3
    }
    if (v5 == (-(v5))) {
        v3 = 47
    }
    switch -332 {
    case -1:
        v3 -= (v3 * (^(v3)))
    case 54:
        v13 := GenS1{F0: 53, F1: 257, F2: -2147483648}
        _ = v13
    case 128:
        v4 = 1
    }
    v1 = ((v5 | -1) < (^(v5)))
    var v14 bool = true
    v15 := GenS1{F0: 4294967295, F1: 1, F2: 401}
    v1 = true
    v16 := GenS2{F0: v5, F1: 127}
    if (!(-9223372036854775808 <= v5)) {
        v17 := []int{-9223372036854775808, -81, -9223372036854775808}
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int8 = (v4 & (v4 & (v4 / ((v4) | 1))))
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        v20 := []int{4294967295, 0, 779}
        v3 -= v3
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
