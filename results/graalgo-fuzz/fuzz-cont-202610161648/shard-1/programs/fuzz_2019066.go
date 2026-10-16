package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 255
    v4 := []string{}
    var v5 uint32 = 65535
    var v6 string = ""
    var v7 int = -2147483648
    if false {
        switch 128 {
        case -274:
            var v8 int16 = -272
            _ = v8
        case -447:
            delete(v2, v5)
        default:
            var v9 uint8 = 127
            _ = v9
        }
    } else {
        switch "a" {
        case "a":
            v3 += v3
        default:
            delete(v1, v5)
        }
        var v10 int16 = 1000
        var v11 int64 = -672
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v6 = ("" + ("\n\t" + v6))
            v5 += v5
            var v14 uint = 18446744073709551615
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS1{F0: false, F1: 42}
    if false {
        if (746 <= 0) {
            v16 := GenS1{F0: false, F1: 32767}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                v18 := GenS1{F0: ((-2.25 / -2.25) == 0.0), F1: 1000}
                _ = fi17
                _ = v18
            }
            if true {
                v19 := v2[340]
                v6 = ((v6 + ("a" + v19)) + ((((v19 + "") + "") + v19) + "foo bar"))
                var v20 uint64 = 42
                _ = v19
                _ = v20
            } else {
                fmt.Println(v6, v5, v3)
                v7 += (v7 | v7)
            }
            fmt.Println(v3)
        }
    }
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v7 {
    case 138:
        if len(v4) > 0 {
            v4[0] = (v6 + (("" + (v6 + "foo bar")) + "a"))
        }
    case 982:
        var v21 float32 = (1e10 * (1.5 / 1.5))
        _ = v21
    case 946:
        if true {
            var v22 uint16 = (v3 - 0)
            _ = len(v2)
            v2[v5] = "hello"
            _ = v22
        }
    default:
        v4 = append(v4, v6)
    }
    v3 -= (v3 ^ v3)
    for fi23 := 0; fi23 < 16; fi23++ {
        v3 -= (v3 % ((136) | 1))
        v24 := v2[(v5 | 1)]
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 int16 = 32767
        _ = fi25
        _ = v26
    }
    if true {
        v6 = (v6 + v6)
        fmt.Println(v5, v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
