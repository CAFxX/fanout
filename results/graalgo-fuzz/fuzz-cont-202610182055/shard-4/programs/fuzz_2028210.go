package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, p0)
        fmt.Println(fi7, fi7)
        if (!false) {
            if false {
                p0 = p0
                p0 = "a"
                switch "\n\t" {
                case "foo bar":
                    switch -1 {
                    case 42:
                        v8 := []int{(fi7 % ((42) | 1)), (fi7 >> ((627) & 63)), 221}
                        _ = v8
                    case 773:
                        for fi9 := 0; fi9 < 7; fi9++ {
                            p0 = p0
                            p0 = ""
                            _ = fi9
                        }
                    default:
                        var v10 int32 = 554
                        _ = v10
                    }
                case "":
                    var v11 int16 = -32768
                    _ = v11
                case "hello":
                    v12 := []int{(fi7 * fi7)}
                    _ = v12
                default:
                    p0 = (p0 + ("\n\t" + p0))
                }
            } else {
                fmt.Println(p0, p0)
            }
            var v13 uint64 = 18446744073709551615
            _ = v13
        }
        _ = fi7
    }
    v14 := []int{}
    fmt.Println(p0)
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "hello"
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 bool = (("a" + ("hello" + ((("\n\t" + "foo bar") + "foo bar") + ("a" + "foo bar")))) >= ("a" + "foo bar"))
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    v4 := []int{0, -985, 128, -409}
    v5 := []int{4294967295, 9223372036854775807, 32767, 4294967295}
    v6 := []string{}
    for fi17 := 0; fi17 < 7; fi17++ {
        switch 80 {
        case 2147483647:
            fmt.Println(v1, v1)
        }
        fmt.Println(v1, v1)
        _ = fi17
    }
    var v18 uint = 65535
    v19 := GenS1{F0: 0, F1: v18, F2: (("\n\t" + "a") + "a")}
    v18 -= v18
    if (255 >= 65535) {
        v18 += (v18 | v18)
    } else {
        v1 = false
    }
    var v20 float32 = 1e10
    v18 -= 2
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 int8 = 0
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 0, F1: 939, F2: "hello"}
    v18 = (^(v18))
    if (true || (30 != -128)) {
        switch 32767 {
        case -407:
            v1 = v1
        }
    } else {
        var v24 int8 = 42
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := GenS1{F0: -2147483648, F1: (v18 << (((v18 >> (((v18 | (v18 ^ 0))) & 63))) & 63)), F2: ("" + "")}
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    var v27 uint16 = 0
    v3[941] = ("a" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
