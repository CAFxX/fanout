package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 string, p2 bool) uint64 {
    p2 = true
    switch 1 {
    case 32768:
        if p2 {
            fmt.Println(p2)
            switch ("\n\t" + "a") {
            case "hello":
                p0 = p0
            }
            v5 := GenS1{F0: 42, F1: 0}
            _ = v5
        } else {
            p2 = p2
            v6 := []int{32767, 9223372036854775807, -1}
            _ = v6
        }
    }
    p2 = (!p2)
    var v7 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint32) int32 {
    for fi8 := 0; fi8 < 2; fi8++ {
        p0 -= p0
        _ = fi8
    }
    p0 += (p0 * 2147483648)
    var v9 uint8 = 127
    v10 := []int{-1, 32768, -9223372036854775808}
    _ = p0
    _ = v10
    _ = v9
    return -673
}

func main() {
    var v1 uint32 = 522
    var v2 bool = (65535 > 1)
    v3 := make(map[int]string)
    v4 := []int{}
    fmt.Println(v1, v2)
    switch 9223372036854775807 {
    case 42:
        for fi11 := 0; fi11 < 19; fi11++ {
            if v2 {
                if v2 {
                    var v12 int64 = -1
                    var v13 int16 = int16(fi11)
                    var v14 int = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            v2 = (true || (!false))
            _ = fi11
        }
    case 2147483647:
        fmt.Println(v2, v1)
    case 306:
        if v2 {
            if (genFn1((v1 | 304)) <= 1) {
                var v15 int32 = genFn1(1)
                var v16 float64 = 1.5
                v17 := GenS1{F0: 720, F1: v15}
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 int64 = 1
                v1 += 310
                _ = v18
            }
        } else {
            v19 := GenS2{F0: genFn1(1), F1: 42, F2: 18446744073709551615}
            _ = v19
        }
    }
    v2 = (v2 && true)
    v2 = (!v2)
    if len(v4) > 0 {
        v4[0] = -625
    }
    var v20 bool = v2
    v21 := GenS2{F0: 1, F1: 0, F2: 255}
    v22 := v3[9223372036854775807]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
