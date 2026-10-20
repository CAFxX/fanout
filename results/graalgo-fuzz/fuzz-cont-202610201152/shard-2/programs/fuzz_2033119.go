package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
func genFn0(p0 int32, p1 uint16) int64 {
    switch "hello" {
    case "":
        var v5 string = ("" + "\n\t")
        _ = v5
    default:
        p0 += p0
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 ^ int(p1)), fi6, (fi6 / (((fi6 ^ 128)) | 1))}
        v8 := GenS1{F0: 2, F1: false}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}
func genFn1() int8 {
    var v9 int8 = -116
    v10 := GenS1{F0: 1, F1: (-2.25 < (-2.25 * 3.14159))}
    v9 += (v9 / (((v9 << ((2) & 63))) | 1))
    var v11 bool = (255 <= 99)
    v9 = 127
    _ = v10
    _ = v11
    _ = v9
    return (v9 >> ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 int32 = -167
    var v4 int16 = 42
    switch (("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + "") {
    case "a":
        v2 -= (^(v2))
    default:
        v12 := GenS1{F0: 2, F1: false}
        _ = v12
    }
    switch ("" + ("foo bar" + "")) {
    case "":
        var v13 int8 = 127
        _ = v13
    }
    var v14 uint8 = 127
    v2 = (v2 - (v2 * 100))
    for fi15 := 0; fi15 < 20; fi15++ {
        v14 = 232
        var v16 int16 = -1
        _ = fi15
        _ = v16
    }
    if true {
        delete(v1, 65535)
        for fi17 := 0; fi17 < 17; fi17++ {
            delete(v1, 2147483648)
            v14 = 0
            for fi18 := 0; fi18 < 20; fi18++ {
                v3 -= 2147483647
                _ = fi18
            }
            _ = fi17
        }
        var v19 bool = (false && true)
        _ = v19
    }
    v14 = 207
    v20 := []int{9223372036854775807, 303, 128, 128}
    switch "foo bar" {
    case "hello":
        v21 := GenS1{F0: 152, F1: false}
        _ = v21
    case "foo bar":
        delete(v1, 2)
    case "":
        if (698 != 350) {
            v22 := v1[614]
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v14 += v14
        v4 -= v4
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
