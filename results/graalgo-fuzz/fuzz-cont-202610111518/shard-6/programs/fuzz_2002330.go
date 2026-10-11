package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 int16) int8 {
    p0 -= (p0 ^ (p0 % ((p0) | 1)))
    v4 := GenS1{F0: 2147483647}
    _ = p0
    _ = v4
    return 92
}

func main() {
    var v1 int32 = -621
    v2 := make(map[string]string)
    var v3 uint = 618
    if (-2.25 <= 0.1) {
        v5 := v2[("\n\t" + ("a" + "foo bar"))]
        _ = v5
    }
    var v6 int64 = -1
    fmt.Println(v6, v1, v1)
    fmt.Println(v1, v6)
    if true {
        v3 = v3
        v7 := []int{32767, 9223372036854775807}
        _ = v7
    } else {
        v8 := GenS1{F0: 2147483647}
        for fi9 := 0; fi9 < 14; fi9++ {
            v3 -= v3
            _ = fi9
        }
        var v10 int64 = v6
        _ = v10
        _ = v8
    }
    v1 = 2147483647
    v3 -= (v3 & v3)
    fmt.Println(v6, v1, v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS2{F0: int8(fi11), F1: fi11, F2: 0}
        _ = len(v2)
        if false {
            v13 := GenS3{F0: 255}
            if (!false) {
                if ((v1 << ((v3) & 63)) >= 0) {
                    v14 := GenS4{F0: ("\n\t" + "hello"), F1: 65535, F2: v6}
                    _ = v14
                }
            }
            _ = v13
        } else {
            for fi15 := 0; fi15 < 8; fi15++ {
                v1 = (v1 << ((v3) & 63))
                v3 = v3
                _ = fi15
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 = (^(v3))
                var v17 float64 = 1.5
                v3 = 533
                _ = fi16
                _ = v17
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
