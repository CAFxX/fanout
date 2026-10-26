package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0(p0 uint16, p1 int8) uint8 {
    p0 += (p0 - (p0 & (p0 >> ((578) & 63))))
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: "hello", F1: -2147483648}
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 string = "a"
    var v3 int16 = -1
    fmt.Println(v3)
    v6 := []int{127, 883, 0}
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 += (v3 | v3)
        switch v2 {
        case "hello":
            v2 = v2
        case "a":
            v2 = v2
        case "foo bar":
            if true {
                fmt.Println(v2, v2)
            }
        default:
            if (!(false && true)) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    _ = len(v6)
                    fmt.Println(v3, fi8)
                    v1 = 255
                    _ = fi8
                }
                if true {
                    switch ("foo bar" + v2) {
                    case "foo bar":
                        v1 = v1
                    case "hello":
                        if true {
                            _ = len(v6)
                        } else {
                            _ = len(v6)
                            var v9 int32 = -2147483648
                            var v10 bool = ((0.1 >= 1.5) && true)
                            _ = v10
                            _ = v9
                        }
                    default:
                        if len(v6) > 0 {
                            v6[0] = fi7
                        }
                    }
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v2)
                    case "\n\t":
                        v2 = v2
                    case "a":
                        fmt.Println(v1, v2, v2)
                    default:
                        switch (fi7 ^ (fi7 ^ 4294967295)) {
                        case 1:
                            v3 += 646
                        case -817:
                            var v11 string = v2
                            _ = v11
                        case -9223372036854775808:
                            var v12 string = ((v2 + v2) + v2)
                            _ = v12
                        default:
                            fmt.Println(fi7, fi7, v1)
                        }
                    }
                }
                v6 = append(v6, 32767)
            } else {
                v2 = v2
            }
        }
        _ = fi7
    }
    fmt.Println(v3)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint32 = 65535
        v1 = v1
        _ = fi13
        _ = v14
    }
    fmt.Println(v1, v1, v3)
    for fi15 := 0; fi15 < 7; fi15++ {
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    var v16 int32 = 0
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v3, v3)
        _ = fi17
    }
    v3 = (v3 | -32768)
    v18 := GenS1{F0: v2, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
