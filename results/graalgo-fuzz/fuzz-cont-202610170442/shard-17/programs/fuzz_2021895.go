package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 42
    var v3 uint16 = 492
    v4 := []string{"foo bar", "\n\t", "\n\t"}
    v5 := make(map[int]int)
    switch (("foo bar" + "\n\t") + "a") {
    case "a":
        if (128 != -198) {
            if false {
                switch 24 {
                case 699:
                    var v6 uint32 = 2147483647
                    _ = v6
                default:
                    v4 = append(v4, "a")
                }
            }
        } else {
            switch "\n\t" {
            case "hello":
                switch "" {
                case "a":
                    if (("hello" + "") > (("a" + "") + "\n\t")) {
                        var v7 int = 4294967295
                        v8 := GenS1{F0: -220, F1: 939}
                        _ = v7
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 10; fi9++ {
                            v3 = v3
                            var v10 uint64 = 1
                            _ = fi9
                            _ = v10
                        }
                        fmt.Println(v2)
                    }
                }
            }
        }
    case "hello":
        v11 := GenS2{F0: 255, F1: v2}
        _ = v11
    default:
        if len(v4) > 0 {
            v4[0] = ("foo bar" + (("" + "a") + "hello"))
        }
    }
    fmt.Println(v3, v2)
    _ = len(v4)
    fmt.Println(v2, v3)
    delete(v5, -138)
    v1[324] = 127
    fmt.Println(v3, v3)
    switch (("foo bar" + "a") + "a") {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        if (0 <= 1) {
            v2 += v2
            var v12 int16 = 255
            _ = v12
        }
    case "\n\t":
        var v13 int32 = -1
        _ = v13
    default:
        if (85 > -128) {
            v2 -= -530
            v3 += v3
            v14 := GenS1{F0: 416, F1: 886}
            _ = v14
        } else {
            var v15 bool = true
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 = (v2 - v2)
        for fi17 := 0; fi17 < 9; fi17++ {
            _ = len(v4)
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            _ = fi17
        }
        fmt.Println(v2)
        _ = fi16
    }
    v3 -= 255
    fmt.Println(v3, v2)
    fmt.Println(v2)
    for fi18 := 0; fi18 < 11; fi18++ {
        v3 += v3
        _ = fi18
    }
    var v19 int64 = (v2 + 585)
    if false {
        var v20 string = (("foo bar" + ("\n\t" + "")) + "\n\t")
        fmt.Println(v3)
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: 0, F1: 255}
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    v19 -= (v2 << ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
