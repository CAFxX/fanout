package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 int16, p2 bool) uint64 {
    var v5 uint32 = 4294967295
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 252
}
func genFn1(p0 uint16, p1 string, p2 string) uint8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 42
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    p0 += p0
    p1 = (("hello" + (((((p2 + p1) + "a") + p2) + ("a" + p1)) + p2)) + "a")
    v8 = append(v8, 32767)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := []string{""}
    v2 := make(map[int]string)
    var v3 bool = (9223372036854775807 <= 9223372036854775807)
    var v4 int32 = 2147483647
    if ((3.14159 > -2.25) && ((-947 == -9223372036854775808) || v3)) {
        var v9 uint64 = 42
        if v3 {
            v10 := GenS1{F0: v4, F1: 463, F2: 255}
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS2{F0: 127}
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            v1 = append(v1, ("\n\t" + "a"))
            var v13 float32 = ((1e10 + 3.14159) + (0.0 + 0.0))
            if v3 {
                v3 = false
            }
            _ = v13
        }
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 float64 = -2.25
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int64 = -952
        _ = fi15
        _ = v16
    }
    if (781 != 808) {
        v14 += 0.0
    } else {
        if (-9223372036854775808 == 42) {
            var v17 bool = (42 > 1000)
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            delete(v2, 32767)
            _ = v17
        }
        var v18 uint8 = 127
        switch "foo bar" {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = ""
            }
        case "":
            v4 -= (v4 % (((v4 << ((791) & 63))) | 1))
        }
        _ = v18
    }
    var v19 int32 = (^(v4))
    v20 := GenS1{F0: -1, F1: 0, F2: 2}
    var v21 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
