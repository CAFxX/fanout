package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0(p0 uint16) string {
    fmt.Println(p0)
    v7 := GenS1{F0: (^(p0))}
    if (true && ((0.1 < float64(1.5)) && true)) {
        v8 := GenS1{F0: p0}
        switch "a" {
        case "a":
            var v9 uint = 851
            _ = v9
        default:
            v10 := []int{1, 128}
            _ = v10
        }
        p0 = p0
        _ = v8
    }
    _ = p0
    _ = v7
    return ("foo bar" + (("foo bar" + "a") + "a"))
}

func main() {
    var v1 int8 = 0
    var v2 int = 1
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    var v5 int64 = 42
    var v6 int64 = (v5 - v5)
    v11 := []int{(-(v2)), (^(v2)), (v2 * 4294967295), (v2 / ((506) | 1))}
    var v12 float64 = 1e10
    var v13 bool = (0 <= 23)
    v14 := v3[0]
    var v15 int64 = v6
    for fi16 := 0; fi16 < 13; fi16++ {
        v1 = v1
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    var v18 int32 = int32(v6)
    for fi19 := 0; fi19 < 11; fi19++ {
        v2 += fi19
        _ = fi19
    }
    if v13 {
        switch "" {
        case "":
            _ = len(v3)
        default:
            delete(v3, -767)
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        v3[9223372036854775807] = 736
    } else {
        switch "foo bar" {
        case "hello":
            if (v5 < -152) {
                var v20 string = "foo bar"
                _ = v20
            }
        case "\n\t":
            var v21 int8 = -58
            _ = v21
        case "foo bar":
            v22 := GenS2{F0: ("hello" != (("foo bar" + "") + "\n\t")), F1: v1, F2: 255}
            _ = v22
        }
        var v23 int8 = 36
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
