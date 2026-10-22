package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []string{("" + ("\n\t" + "foo bar")), ("a" + "\n\t")}
    var v2 int16 = 1
    var v3 uint16 = 65535
    var v4 int8 = -52
    v3 += (v3 / ((v3) | 1))
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint16 = v3
        _ = fi5
        _ = v6
    }
    v2 = (v2 + (v2 >> ((65535) & 63)))
    _ = len(v1)
    var v7 uint8 = 0
    var v8 int64 = -269
    v9 := GenS1{F0: (true && ((v2 & v2) == v2)), F1: 1}
    var v10 int16 = v2
    fmt.Println(v10, v8, v8)
    if true {
        switch -934 {
        case 753:
            if false {
                var v11 uint16 = v3
                _ = v11
            }
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v8 += (v8 * (v8 + (v8 + (v8 >> ((531) & 63)))))
        switch "foo bar" {
        case "foo bar":
            v4 = (v4 - (^(v4)))
        default:
            fmt.Println(v10, v7, v3)
        }
        _ = fi12
    }
    var v13 uint = 0
    var v14 uint64 = 937
    var v15 int8 = (v4 % (((^(v4))) | 1))
    v13 -= (v13 * v13)
    if ((v13 / ((0) | 1)) < 77) {
        v16 := GenS1{F0: (!false), F1: v14}
        _ = v16
    } else {
        v17 := GenS1{F0: (false && (false || true)), F1: uint64(v3)}
        var v18 float32 = -2.25
        v1 = append(v1, "foo bar")
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if (!true) {
            v2 = v10
            v3 += v3
        }
        fmt.Println(v8)
        _ = fi19
    }
    v3 -= (v3 << (((v13 / ((42) | 1))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
