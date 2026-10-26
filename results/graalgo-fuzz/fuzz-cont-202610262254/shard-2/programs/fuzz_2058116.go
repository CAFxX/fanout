package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int32) uint32 {
    if true {
        fmt.Println(p0, p0)
    } else {
        v4 := GenS1{F0: 6, F1: true}
        _ = v4
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        if ("a" >= "") {
            var v6 uint32 = 65535
            var v7 int16 = -521
            _ = v6
            _ = v7
        } else {
            if false {
                switch ("foo bar" + "foo bar") {
                case "":
                    var v8 string = ("foo bar" + "")
                    _ = v8
                }
            } else {
                v9 := []int{(fi5 - (fi5 % ((fi5) | 1))), (fi5 + 0)}
                _ = v9
            }
            p1 = p1
            var v10 int32 = (p1 & (p1 & (p1 + (p1 ^ 0))))
            _ = v10
        }
        p0 = (^(p0))
        _ = fi5
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint8, p1 int8, p2 int16) int64 {
    v11 := []int{}
    var v12 uint16 = 0
    for fi13 := 0; fi13 < 13; fi13++ {
        _ = len(v11)
        p1 -= (p1 + (p1 << ((788) & 63)))
        p0 -= (^(p0))
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 754
}

func main() {
    var v1 uint8 = 86
    v2 := []int{-596, 79}
    var v3 uint32 = 27
    v14 := GenS1{F0: 404, F1: (-9223372036854775808 >= 32768)}
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = v3
        fmt.Println(fi15, v1)
        var v16 int32 = 2147483647
        _ = fi15
        _ = v16
    }
    v2 = append(v2, 4294967295)
    v1 -= (v1 >> ((18446744073709551615) & 63))
    v1 = (v1 - v1)
    if (65 <= -22) {
        v17 := GenS2{F0: v3}
        var v18 uint16 = 65535
        _ = v17
        _ = v18
    } else {
        switch -1 {
        case 9223372036854775807:
            v1 -= v1
        case 841:
            var v19 float32 = (-2.25 + -2.25)
            _ = v19
        case -948:
            v1 += v1
        default:
            switch "a" {
            case "hello":
                switch ((("foo bar" + "") + "hello") + "") {
                case "a":
                    var v20 int64 = 1
                    _ = v20
                case "":
                    fmt.Println(v1, v3)
                case "hello":
                    var v21 string = ((((("hello" + "\n\t") + ("\n\t" + "foo bar")) + "\n\t") + "foo bar") + "a")
                    _ = v21
                }
            default:
                if true {
                    var v22 uint64 = 1
                    var v23 float32 = 3.14159
                    var v24 int32 = 42
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
        }
    }
    v25 := GenS2{F0: 200}
    var v26 float32 = (1e10 + (1.5 / 3.14159))
    fmt.Println(v26, v1)
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
