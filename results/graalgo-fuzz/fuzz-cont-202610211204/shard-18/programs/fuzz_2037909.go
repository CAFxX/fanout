package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []int{42, 128, 0, 127}
    v4 := make(map[uint32]int)
    var v5 int32 = 549
    v5 = (v5 + v5)
    v6 := v2[("hello" + ("a" + ("a" + "hello")))]
    v7 := GenS1{F0: 2, F1: 9223372036854775807, F2: "hello"}
    delete(v1, "a")
    v8 := GenS1{F0: 0, F1: 127, F2: (("a" + ("foo bar" + ((("a" + "\n\t") + "foo bar") + ("hello" + ("\n\t" + ""))))) + (("a" + ("a" + (("hello" + ("" + "a")) + ""))) + ""))}
    switch "foo bar" {
    case "":
        v9 := GenS1{F0: 255, F1: 549, F2: "\n\t"}
        _ = v9
    case "\n\t":
        v3 = append(v3, -450)
    default:
        delete(v2, "hello")
    }
    fmt.Println(v6, v5)
    v6 = true
    var v10 bool = (32767 > -757)
    v11 := GenS1{F0: 29, F1: 32767, F2: "foo bar"}
    v6 = ((65535 >= 42) && v10)
    var v12 bool = false
    v2["\n\t"] = true
    switch "hello" {
    case "a":
        for fi13 := 0; fi13 < 4; fi13++ {
            switch 1 {
            case 42:
                v14 := v2[("" + (("foo bar" + ((("\n\t" + ("\n\t" + "\n\t")) + "hello") + ("hello" + ("a" + (("hello" + "") + ("\n\t" + "\n\t")))))) + "a"))]
                _ = v14
            }
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int16 = -32768
                var v17 bool = v10
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
        }
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v12, v6)
        default:
            var v18 uint8 = 1
            _ = v18
        }
    case "hello":
        var v19 uint8 = 255
        _ = v19
    default:
        v12 = v10
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v5 += v5
        fmt.Println(v6)
        var v21 uint16 = 255
        _ = fi20
        _ = v21
    }
    fmt.Println(v10, v6, v6)
    fmt.Println(v10)
    v12 = ((1e10 + 3.14159) >= (1e10 * float32((0.1 * 1.5))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
