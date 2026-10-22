package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64) uint8 {
    switch "a" {
    case "":
        v4 := GenS1{F0: ""}
        _ = v4
    case "foo bar":
        v5 := GenS2{F0: 65535}
        _ = v5
    case "\n\t":
        v6 := []int{32768}
        _ = v6
    }
    var v7 uint32 = 42
    var v8 int = 32767
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    var v3 uint16 = 32767
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint = 355
        _ = fi9
        _ = v10
    }
    v11 := []int{582, 1}
    if false {
        v12 := GenS2{F0: 0}
        var v13 string = "hello"
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                fmt.Println(v3, v3)
                v2 = v2
                var v16 uint32 = 650
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi15
                _ = v16
            }
            if false {
                switch "hello" {
                case "foo bar":
                    var v17 int32 = 2147483647
                    _ = v17
                case "\n\t":
                    switch fi14 {
                    case 127:
                        v3 = v3
                    case 1:
                        var v18 uint8 = 255
                        _ = v18
                    default:
                        v19 := GenS1{F0: "\n\t"}
                        _ = v19
                    }
                case "":
                    var v20 uint64 = (v2 + 719)
                    _ = v20
                default:
                    fmt.Println(v3, v2)
                }
                v3 += (v3 ^ (^(v3)))
            }
            _ = fi14
        }
        var v21 string = ("foo bar" + "\n\t")
        v22 := GenS3{F0: 502, F1: 2}
        _ = v21
        _ = v22
    }
    v23 := v1[("a" + "\n\t")]
    if (1 != 1000) {
        switch 128 {
        case -2147483648:
            for fi24 := 0; fi24 < 17; fi24++ {
                v25 := GenS4{F0: 47}
                _ = fi24
                _ = v25
            }
        case 923:
            var v26 int8 = 127
            _ = v26
        default:
            v2 = 2
        }
        v27 := GenS1{F0: ("a" + "foo bar")}
        _ = v27
    } else {
        fmt.Println(v2, v3, v23)
    }
    v23 = v23
    v23 = v23
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            v2 -= v2
            v2 = (v2 % ((971) | 1))
            _ = fi29
        }
        var v30 string = ""
        for fi31 := 0; fi31 < 16; fi31++ {
            v23 = v23
            _ = fi31
        }
        _ = fi28
        _ = v30
    }
    fmt.Println(v23, v3, v23)
    _ = len(v1)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v32 := GenS4{F0: 100}
    fmt.Println(v3, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v32
}
