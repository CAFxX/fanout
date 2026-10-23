package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 float32 = ((((((float32(((1e10 - -2.25) / float64((1.5 * 0.1)))) * 0.0) - 3.14159) / (1.5 + -2.25)) / float32(0.1)) + (((-2.25 / 1e10) * float32((0.1 / -2.25))) / float32((1.5 + 1.5)))) + 3.14159)
    v2[0] = 32767
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v3[""] = "\n\t"
            v7 := []int{(fi6 * (fi5 - (fi6 & (fi5 + 42)))), (fi5 | fi6), (^(fi5)), fi5}
            var v8 uint8 = 127
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 255, F1: -1, F2: 255}
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = ("a" + "a")
        _ = fi10
        _ = v11
    }
    v12 := []int{-649, -9223372036854775808}
    var v13 float64 = 0.1
    if true {
        var v14 float64 = (1e10 / (0.0 / float64(v4)))
        _ = v14
    }
    v12 = append(v12, -9223372036854775808)
    fmt.Println(v1, v13, v4)
    if true {
        var v15 uint32 = 675
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS2{F0: 0}
            fmt.Println(fi16, v1)
            for fi18 := 0; fi18 < 1; fi18++ {
                switch "hello" {
                case "hello":
                    v15 = (v15 / (((v15 % ((2147483647) | 1))) | 1))
                case "a":
                    switch (fi18 >> ((65535) & 63)) {
                    case -526:
                        var v19 float64 = v13
                        _ = v19
                    }
                case "foo bar":
                    v1 = v1
                }
                var v20 int64 = 596
                var v21 string = "foo bar"
                _ = fi18
                _ = v20
                _ = v21
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        switch 127 {
        case 906:
            delete(v2, 0)
        case 42:
            switch "hello" {
            case "\n\t":
                var v22 uint32 = 0
                _ = v22
            default:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        }
    case "":
        var v23 int16 = -1
        _ = v23
    }
    v1 = (v1 + 255)
    if false {
        if (!(65535 >= 42)) {
            v3[("foo bar" + "hello")] = "hello"
        } else {
            v13 += v13
        }
        v24 := GenS2{F0: -1}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
