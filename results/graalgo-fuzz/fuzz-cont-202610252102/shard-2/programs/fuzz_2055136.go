package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("\n\t" + "a")
    v3 := []string{""}
    v4 := []uint32{}
    var v5 uint32 = 2147483648
    switch 32768 {
    case 903:
        var v6 int8 = 100
        _ = v6
    case -162:
        if (true && ((!false) && true)) {
            v2 = "foo bar"
        } else {
            v2 = ("" + "")
        }
    case 2147483647:
        switch "hello" {
        case "foo bar":
            for fi7 := 0; fi7 < 15; fi7++ {
                fmt.Println(v2, fi7)
                _ = fi7
            }
        default:
            fmt.Println(v2, v5)
        }
    default:
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    switch 128 {
    case 9223372036854775807:
        var v9 int64 = 661
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        switch 2147483647 {
        case 902:
            v5 -= (v5 ^ v5)
        case 2147483647:
            v2 = v2
        default:
            v5 = v5
        }
        _ = fi10
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v11 := v1[-9223372036854775808]
        var v12 uint = 621
        _ = v11
        _ = v12
    } else {
        v1[128] = true
    }
    fmt.Println(v2, v5)
    if false {
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        switch fi13 {
        case -1:
            var v14 bool = (v2 < v2)
            _ = v14
        case -793:
            v15 := GenS1{F0: 0}
            _ = v15
        default:
            if false {
                fmt.Println(v5, v5, v5)
                v2 = v2
            }
        }
        _ = fi13
    }
    if (v2 > v2) {
        var v16 int32 = -97
        var v17 uint = 2
        v16 = 148
        _ = v16
        _ = v17
    }
    delete(v1, 149)
    v18 := GenS2{F0: -751, F1: 65535}
    v19 := GenS1{F0: 42}
    var v20 uint32 = v5
    if (true || (32767 == 65535)) {
        if true {
            var v21 int8 = 0
            _ = v21
        } else {
            var v22 int = -951
            _ = v22
        }
    } else {
        v23 := GenS2{F0: -2147483648, F1: 0}
        var v24 uint16 = 65535
        var v25 int8 = -1
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
