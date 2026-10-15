package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int, p1 int64) string {
    fmt.Println(p1, p0)
    if true {
        var v5 int16 = -1
        _ = v5
    }
    v6 := []int{(-(p0))}
    v7 := GenS1{F0: -1, F1: 9223372036854775807}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ("foo bar" + (("\n\t" + ("\n\t" + "")) + ""))
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 uint8 = 82
    v4 := []uint32{}
    _ = len(v4)
    switch 4294967295 {
    case -121:
        v8 := GenS1{F0: 853, F1: -9223372036854775808}
        _ = v8
    case -509:
        var v9 int = -398
        _ = v9
    default:
        v3 = uint8(v2)
    }
    switch 4294967295 {
    case 918:
        v3 -= 0
    case 9223372036854775807:
        fmt.Println(v3, v2)
    case 701:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 bool = false
            v3 = (^(v3))
            v12 := GenS2{F0: v11, F1: 4294967295}
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS2{F0: (18446744073709551615 == 4294967295), F1: 65535}
            _ = fi13
            _ = v14
        }
    }
    v2 = v2
    var v15 int8 = -57
    var v16 string = "foo bar"
    if (0.1 >= ((1e10 / (-2.25 - (0.1 + 1e10))) / 1.0)) {
        delete(v1, ("hello" + (v16 + ((v16 + "foo bar") + ("\n\t" + v16)))))
    } else {
        var v17 string = ("hello" + v16)
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
