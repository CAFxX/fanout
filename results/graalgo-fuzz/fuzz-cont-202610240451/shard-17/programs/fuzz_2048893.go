package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
func genFn0(p0 int64) bool {
    p0 = p0
    v5 := GenS1{F0: -1, F1: 275, F2: true}
    _ = p0
    _ = v5
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1000
    var v3 bool = (-2147483648 <= 0)
    var v4 uint64 = 18446744073709551615
    if v3 {
        var v6 int64 = 549
        v7 := GenS1{F0: -128, F1: (v4 & 18446744073709551615), F2: ("a" >= "foo bar")}
        switch -1 {
        case 42:
            switch 0 {
            case 30:
                v6 = v6
            default:
                v8 := GenS1{F0: -128, F1: (v4 >> ((495) & 63)), F2: (255 >= 4294967295)}
                _ = v8
            }
        case 0:
            for fi9 := 0; fi9 < 3; fi9++ {
                if v3 {
                    v4 = 1
                    v3 = (!true)
                }
                fmt.Println(fi9)
                var v10 uint64 = 91
                _ = fi9
                _ = v10
            }
        case 778:
            v3 = ((v3 || ("a" == "hello")) && (32767 > 713))
        default:
            v11 := GenS1{F0: 42, F1: v4, F2: v3}
            _ = v11
        }
        _ = v6
        _ = v7
    }
    var v12 int8 = -128
    if v3 {
        var v13 int32 = 1
        switch "hello" {
        case "a":
            v13 = 42
        case "\n\t":
            switch 128 {
            case -341:
                fmt.Println(v13, v3)
            case 1:
                fmt.Println(v12, v12)
            default:
                delete(v1, 32768)
            }
        default:
            fmt.Println(v2, v12)
        }
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(v4, v4, v12)
        _ = fi14
    }
    var v15 int = int(v12)
    v1[v15] = ("hello" + "")
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v3)
        _ = fi16
    }
    v17 := []int{int(v4), v15}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
