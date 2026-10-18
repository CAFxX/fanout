package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint {
    fmt.Println(p0)
    p0 = (127 >= -1)
    p0 = (42 >= 127)
    var v5 string = ("hello" + "hello")
    _ = p0
    _ = v5
    return 223
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    v3 := make(map[string]string)
    var v4 int64 = 42
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -649
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint = 261
            if false {
                v9 = (v9 << ((2) & 63))
            } else {
                v9 = (^(v9))
                v10 := []int{fi6, (fi8 / (((fi8 >> (((v9 + 1)) & 63))) | 1))}
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        v4 = v4
        _ = fi6
        _ = v7
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v1[fi11] = (fi11 & -2147483648)
        v2 += (v2 % ((v2) | 1))
        v12 := GenS1{F0: 964, F1: fi11}
        _ = fi11
        _ = v12
    }
    if (1 > 1) {
        for fi13 := 0; fi13 < 20; fi13++ {
            if true {
                delete(v3, "foo bar")
            }
            v14 := GenS2{F0: 42}
            for fi15 := 0; fi15 < 4; fi15++ {
                if (true || false) {
                    v2 += (v2 / (((v2 % ((v2) | 1))) | 1))
                    delete(v1, (-(fi15)))
                    v2 -= v2
                }
                v16 := []int{}
                if (-2.25 != ((0.0 / -2.25) / 1.5)) {
                    var v17 uint64 = 2
                    _ = v17
                } else {
                    if (121 > -27) {
                        fmt.Println(v2, v4)
                    }
                }
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        var v18 uint64 = 162
        var v19 float32 = 1e10
        _ = v18
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        if (!true) {
            if (!((42 >= 2) || true)) {
                v20 := []int{-509, -721}
                _ = v20
            } else {
                v21 := GenS2{F0: 0}
                _ = v21
            }
        } else {
            v22 := []int{}
            if true {
                var v23 string = ("\n\t" + "hello")
                if false {
                    v23 = v23
                    v4 += (v4 / (((v4 % ((v4) | 1))) | 1))
                } else {
                    delete(v3, (v23 + v23))
                    var v24 float32 = float32(float64(1e10))
                    _ = v24
                }
                _ = v23
            } else {
                v22 = append(v22, 0)
            }
            var v25 uint32 = 2147483647
            _ = v22
            _ = v25
        }
    case "a":
        v26 := []int{127, 127, 32768, 4294967295}
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v4 = v4
        var v28 float64 = (1.5 / 3.14159)
        var v29 int = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 42
    v30 += (v30 % (((v30 ^ v30)) | 1))
    delete(v3, ("foo bar" + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
