package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0(p0 int64, p1 string, p2 int64) int {
    var v4 uint16 = 1
    if (-2147483648 < 128) {
        fmt.Println(v4)
        var v5 string = p1
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 95
}

func main() {
    var v1 int8 = 42
    var v2 uint = 1
    var v3 int8 = 100
    var v6 int64 = 1
    var v7 uint = v2
    var v8 int8 = (v3 - 42)
    var v9 bool = true
    for fi10 := 0; fi10 < 15; fi10++ {
        if true {
            v11 := []int{}
            v8 = v8
            v12 := GenS1{F0: 42}
            _ = v11
            _ = v12
        } else {
            v13 := []int{genFn0((-(v6)), "foo bar", (v6 >> (((v7 ^ (v7 / (((v7 >> ((v7) & 63))) | 1)))) & 63))), genFn0(v6, ("foo bar" + "foo bar"), v6), fi10}
            for fi14 := 0; fi14 < 18; fi14++ {
                v15 := GenS2{F0: 116}
                v9 = ((v6 * v6) < (-(v6)))
                var v16 int8 = v8
                _ = fi14
                _ = v15
                _ = v16
            }
            v13 = append(v13, fi10)
            _ = v13
        }
        v2 = (v7 / (((v2 & v7)) | 1))
        v17 := GenS1{F0: 2}
        _ = fi10
        _ = v17
    }
    v8 = 20
    v18 := []int{32768}
    if ("\n\t" <= "a") {
        _ = len(v18)
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 int8 = (v8 >> ((18446744073709551615) & 63))
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            var v21 uint32 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    _ = len(v18)
    v2 -= 18446744073709551615
    if (975 >= 265) {
        if (v7 <= (v7 + v7)) {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int32 = 2147483647
                _ = fi22
                _ = v23
            }
            v24 := GenS2{F0: 1}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                if (836 <= 18446744073709551615) {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        var v27 uint8 = 42
                        if len(v18) > 0 {
                            v18[0] = (^(fi26))
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v28 int16 = -641
                    v29 := GenS3{F0: "\n\t", F1: 2}
                    _ = v28
                    _ = v29
                }
                if ((v1 / ((-51) | 1)) != 42) {
                    v1 -= -128
                }
                _ = fi25
            }
        }
        var v30 bool = (25 < 2147483647)
        v31 := GenS3{F0: (("\n\t" + "a") + "a"), F1: 215}
        _ = v30
        _ = v31
    } else {
        v3 -= 42
        var v32 uint16 = uint16(v2)
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
