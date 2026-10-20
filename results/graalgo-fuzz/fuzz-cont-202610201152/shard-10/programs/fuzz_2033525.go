package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int16, p1 int16, p2 uint64) uint64 {
    v4 := []int{4294967295}
    var v5 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]string)
    var v3 int16 = -32768
    v2[("hello" + "")] = "\n\t"
    fmt.Println(v3, v3)
    delete(v2, ("hello" + "hello"))
    fmt.Println(v1, v3)
    var v6 int = 42
    fmt.Println(v3, v6)
    v7 := GenS1{F0: 119, F1: v1}
    if false {
        if (0 == -2147483648) {
            v6 += v6
            var v8 uint8 = 255
            _ = v8
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                v3 -= (v3 ^ int16(v1))
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 uint16 = 1
                    if false {
                        var v12 float32 = -2.25
                        v13 := []int{2147483647, (fi9 * 9223372036854775807)}
                        _ = len(v2)
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            var v14 uint16 = 0
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            v2["hello"] = (("\n\t" + "a") + "")
            var v16 int64 = 174
            switch "foo bar" {
            case "":
                fmt.Println(v6, v6, v6)
            case "hello":
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 int8 = -7
                    v6 = (fi17 & (v6 ^ (fi17 / ((fi17) | 1))))
                    var v19 int8 = (v18 >> ((0) & 63))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi15
            _ = v16
        }
        v3 = v3
    }
    if false {
        v20 := []int{2147483647, 4294967295, (v6 & 32768)}
        var v21 bool = false
        _ = v20
        _ = v21
    }
    fmt.Println(v6, v1, v3)
    v22 := []int{4294967295, v6}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    _ = len(v2)
    for fi23 := 0; fi23 < 3; fi23++ {
        v3 += v3
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v24 uint16 = 0
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: (v1 / (((v1 % ((2) | 1))) | 1)), F1: 255}
    v1 += (v1 % (((v1 << ((0) & 63))) | 1))
    for fi26 := 0; fi26 < 16; fi26++ {
        v3 = (v3 >> ((755) & 63))
        var v27 int8 = 127
        _ = len(v22)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
