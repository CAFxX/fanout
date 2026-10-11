package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 uint8, p1 int64, p2 uint32) uint64 {
    v4 := GenS1{F0: 381, F1: uint8(p2)}
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 int = fi5
        var v7 int64 = p1
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    v8 := GenS2{F0: "hello"}
    var v9 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    _ = v9
    return 563
}

func main() {
    var v1 int64 = 190
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v10 float32 = 3.14159
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    switch 9223372036854775807 {
    case 310:
        v11 := GenS3{F0: 445, F1: "foo bar"}
        _ = v11
    case 127:
        v1 = 0
    case 9223372036854775807:
        switch 32768 {
        case -1:
            if len(v2) > 0 {
                v2[0] = 783
            }
        }
    default:
        v2 = append(v2, 65535)
    }
    switch "foo bar" {
    case "\n\t":
        v3["hello"] = (!false)
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch -2147483648 {
        case 127:
            v10 = float32(0.1)
        case 1:
            var v12 uint64 = 4294967295
            _ = v12
        default:
            v10 += 3.14159
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            v3["hello"] = (0 != 32767)
            var v15 bool = (42 < 18446744073709551615)
            v16 := GenS2{F0: ("\n\t" + ("hello" + "a"))}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    if false {
        v17 := GenS4{F0: -1}
        if false {
            var v18 int64 = v1
            var v19 int8 = 127
            _ = v18
            _ = v19
        }
        v1 += v1
        _ = v17
    }
    v20 := v3[("hello" + "foo bar")]
    v1 = v1
    var v21 uint8 = 127
    v22 := GenS5{F0: 2, F1: 1, F2: 32767}
    var v23 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
