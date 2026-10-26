package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string }

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[string]int)
    var v3 uint8 = 42
    var v4 uint = 18446744073709551615
    var v5 int32 = -52
    v6 := []string{"a", ("a" + "a")}
    v3 -= 0
    v5 += v5
    if false {
        v1 = append(v1, 2147483647)
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int8 = int8(v5)
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v10 int16 = -506
        _ = v10
    }
    v11 := GenS1{F0: 65}
    v3 -= uint8(v4)
    v2[(("" + ("foo bar" + "\n\t")) + "")] = -2147483648
    v5 = v5
    v12 := GenS1{F0: 137}
    var v13 string = "a"
    switch v13 {
    case "foo bar":
        switch "foo bar" {
        case "a":
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(v3, v4)
                fmt.Println(fi14, v13, v3)
                _ = fi14
            }
        case "foo bar":
            var v15 uint8 = (v3 ^ 255)
            _ = v15
        default:
            for fi16 := 0; fi16 < 2; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v13 = v13
                    _ = fi17
                }
                _ = fi16
            }
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v4 -= v4
        var v19 uint32 = 975
        v20 := GenS2{F0: 1, F1: "hello"}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
