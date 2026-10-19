package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 uint16; F2 int }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 255
    v2 := []int{0}
    var v3 int32 = 617
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    var v6 int64 = -9223372036854775808
    var v7 bool = ((18446744073709551615 != 1) && true)
    var v8 int32 = (v3 << ((0) & 63))
    v9 := GenS1{F0: "\n\t", F1: v6, F2: -55}
    fmt.Println(v6, v7, v7)
    delete(v5, 1)
    v6 += v6
    switch -1 {
    case 1:
        if false {
            v8 = -956
            v10 := v5[128]
            v1 = (v1 % ((952) | 1))
            _ = v10
        }
    case 32768:
        fmt.Println(v1, v1, v7)
    case 42:
        v1 += (v1 * (v1 - 255))
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int64 = v6
        v3 += v3
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint = 1
            if ((v6 | (v6 / ((0) | 1))) > v12) {
                var v15 uint64 = 18446744073709551615
                if v7 {
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        v4[(fi13 & fi13)] = (904 > v1)
                    case -2147483648:
                        if true {
                            v16 := GenS2{F0: "a", F1: (v1 << ((959) & 63)), F2: fi11}
                            var v17 int8 = 127
                            _ = v16
                            _ = v17
                        } else {
                            if v7 {
                                switch "a" {
                                case "foo bar":
                                    if (("hello" >= ("a" + ("" + ((("\n\t" + "") + "a") + (("hello" + "a") + "a"))))) || v7) {
                                        fmt.Println(v3, v3)
                                        fmt.Println(v3)
                                    }
                                case "\n\t":
                                    v7 = (!(128 < 127))
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = (-(fi13))
                                    }
                                }
                            }
                            v12 -= (^(v12))
                            var v18 float32 = 1.5
                            _ = v18
                        }
                    case 128:
                        v12 = (v12 | 42)
                    default:
                        switch "" {
                        case "\n\t":
                            v3 -= 0
                        case "hello":
                            delete(v5, (fi11 / (((fi11 + fi13)) | 1)))
                        case "a":
                            v5[-9223372036854775808] = ("hello" + "a")
                        default:
                            v15 = 682
                        }
                    }
                    v12 -= (v12 * (v6 & v12))
                    v19 := GenS2{F0: ("\n\t" + ("" + "foo bar")), F1: (v1 << ((18446744073709551615) & 63)), F2: 42}
                    _ = v19
                } else {
                    v20 := GenS3{F0: "hello"}
                    _ = v20
                }
                var v21 uint8 = 128
                _ = v15
                _ = v21
            }
            v4[(fi13 >> ((v14) & 63))] = true
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v22 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
