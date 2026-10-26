package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
func genFn0(p0 int16) int8 {
    p0 -= -32768
    fmt.Println(p0)
    var v4 int16 = (p0 & p0)
    if false {
        fmt.Println(v4, p0)
    } else {
        v4 -= 1000
    }
    _ = p0
    _ = v4
    return 1
}

func main() {
    v1 := []string{}
    var v2 int32 = 175
    v3 := []uint32{0}
    v2 -= -1
    switch "a" {
    case "hello":
        switch ("hello" + "hello") {
        case "foo bar":
            for fi5 := 0; fi5 < 2; fi5++ {
                if len(v3) > 0 {
                    v3[0] = 360
                }
                _ = fi5
            }
        default:
            switch -49 {
            case 42:
                switch 9223372036854775807 {
                case 933:
                    v2 -= 42
                case 2147483647:
                    v2 = (v2 & v2)
                }
            }
        }
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 191
        }
    default:
        fmt.Println(v2)
    }
    var v6 string = ""
    var v7 uint32 = 2147483648
    switch ("hello" + "") {
    case "foo bar":
        fmt.Println(v7)
    case "hello":
        var v8 float64 = -2.25
        _ = v8
    }
    fmt.Println(v7, v2)
    v3 = append(v3, v7)
    v6 = v6
    if true {
        v9 := GenS1{F0: -117, F1: 42}
        v10 := GenS1{F0: 32768, F1: 124}
        v11 := GenS1{F0: 2147483647, F1: 0}
        _ = v10
        _ = v11
        _ = v9
    } else {
        v2 += v2
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi12
    }
    var v13 int16 = 42
    fmt.Println(v6, v2)
    fmt.Println(v6, v13, v13)
    var v14 int64 = -1
    v1 = append(v1, (v6 + v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
