package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0() int64 {
    var v7 int64 = 9223372036854775807
    fmt.Println(v7)
    v8 := []int{}
    v9 := GenS1{F0: 0}
    v10 := GenS1{F0: 2}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v7 * (v7 << ((739) & 63)))
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a", "a"}
    var v2 int = 32767
    v3 := make(map[int]bool)
    var v4 string = ""
    v5 := []string{"foo bar", (((("" + v4) + v4) + "") + "hello"), ("hello" + v4), (v4 + ("" + v4))}
    var v6 string = "foo bar"
    v4 = (v6 + v4)
    switch v2 {
    case 32768:
        if ((true || true) || false) {
            var v11 string = "foo bar"
            _ = v11
        } else {
            var v12 uint = 18446744073709551615
            if true {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 uint64 = 24
                    _ = fi13
                    _ = v14
                }
                var v15 uint8 = 81
                _ = v15
            } else {
                v4 = "foo bar"
                var v16 uint8 = 0
                v16 = v16
                _ = v16
            }
            _ = v12
        }
    case 4294967295:
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS2{F0: 2, F1: 0, F2: 24}
            v2 = (fi17 ^ -673)
            var v19 uint8 = 128
            _ = fi17
            _ = v18
            _ = v19
        }
    case -9223372036854775808:
        v4 = v6
    default:
        if false {
            v4 = v4
            for fi20 := 0; fi20 < 9; fi20++ {
                v2 = v2
                v3[(-(fi20))] = ((!(v6 != (((("foo bar" + "hello") + ((v6 + v6) + (("hello" + "a") + "\n\t"))) + ((("hello" + v4) + ("hello" + "hello")) + v6)) + (v6 + v6)))) && false)
                v4 = "hello"
                _ = fi20
            }
            fmt.Println(v2, v2)
        }
    }
    switch 42 {
    case 127:
        v21 := GenS2{F0: 128, F1: 2, F2: 1}
        _ = v21
    default:
        if true {
            for fi22 := 0; fi22 < 5; fi22++ {
                v2 = 32768
                switch v2 {
                case 9223372036854775807:
                    v2 = v2
                case 4294967295:
                    fmt.Println(v2, v4)
                case -2147483648:
                    _ = len(v3)
                default:
                    v6 = (v6 + ((v6 + ("\n\t" + v4)) + (v6 + (v6 + "a"))))
                }
                for fi23 := 0; fi23 < 14; fi23++ {
                    v2 = (v2 | -595)
                    switch fi23 {
                    case 331:
                        v2 += -2147483648
                    }
                    _ = fi23
                }
                _ = fi22
            }
        }
    }
    v1 = append(v1, (v6 + "foo bar"))
    v2 = v2
    for fi24 := 0; fi24 < 14; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            if true {
                var v26 int32 = 2147483647
                _ = v26
            }
            switch v4 {
            case "foo bar":
                v6 = "foo bar"
            case "\n\t":
                v2 = (fi24 % (((fi24 | (fi25 >> ((42) & 63)))) | 1))
                v6 = "\n\t"
                var v27 int8 = -16
                _ = v27
            case "a":
                fmt.Println(fi25, v2, v2)
            }
            _ = fi25
        }
        if len(v5) > 0 {
            v5[0] = (((v4 + "\n\t") + v6) + v4)
        }
        _ = fi24
    }
    fmt.Println(v4, v2)
    v28 := GenS3{F0: 4294967295, F1: -1}
    if true {
        var v29 float32 = 1.5
        _ = v29
    }
    var v30 uint16 = 504
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
