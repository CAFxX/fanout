package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 uint8) uint64 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 float64 = 3.14159
        if false {
            var v8 int16 = -32768
            _ = v8
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: ((-746 > 375) && true), F1: (fi9 << ((794) & 63))}
                switch "a" {
                case "\n\t":
                    var v11 int32 = 0
                    _ = v11
                case "hello":
                    if (!((p0 > p0) && (false || ((2147483647 >= 2147483647) && false)))) {
                        if false {
                            v12 := GenS1{F0: false, F1: (fi9 - (fi9 | (fi6 - 1)))}
                            v13 := GenS2{F0: 0, F1: 573, F2: 42}
                            _ = v12
                            _ = v13
                        }
                        p1 -= (p1 + p1)
                    }
                }
                _ = fi9
                _ = v10
            }
            v14 := []int{(fi6 % ((fi6) | 1)), 901, 127}
            if len(v14) > 0 {
                v14[0] = (fi6 * 32768)
            }
            _ = v14
        }
        var v15 uint8 = (p1 * p1)
        _ = fi6
        _ = v15
        _ = v7
    }
    v16 := GenS2{F0: 2147483648, F1: 2, F2: 1}
    v17 := []int{-9223372036854775808, 42, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := []int{1, 9223372036854775807, 752}
    var v4 int32 = 857
    v5 := []string{"a", "\n\t", ("\n\t" + ("\n\t" + ("" + ("foo bar" + "foo bar"))))}
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS3{F0: -886}
        v4 = (v4 << ((42) & 63))
        _ = fi18
        _ = v19
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + "foo bar")
    }
    fmt.Println(v4)
    v4 = -739
    var v20 float32 = 3.14159
    if len(v3) > 0 {
        v3[0] = 32768
    }
    v4 += v4
    v2[32] = 9223372036854775807
    var v21 int8 = 127
    var v22 bool = false
    v22 = ((((-2.25 * 1e10) / v20) - (float32(0.0) / 3.14159)) >= (0.0 + (3.14159 + (v20 + (v20 - (1e10 - v20))))))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        _ = len(v3)
        _ = fi23
    }
    v3 = append(v3, 42)
    v4 += -1
    if (!(!(true && (32767 >= 313)))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    }
    if v22 {
        if v22 {
            v1[2] = 998
            v24 := GenS3{F0: 1}
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
