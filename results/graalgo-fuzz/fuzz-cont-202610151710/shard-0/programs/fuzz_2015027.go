package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    v1 := []string{"\n\t", ("foo bar" + ("foo bar" + ("hello" + (("hello" + "\n\t") + "")))), "foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    v1 = append(v1, "")
    var v4 bool = (("" + "foo bar") >= "\n\t")
    v4 = (!v4)
    if false {
        fmt.Println(v4)
        for fi5 := 0; fi5 < 15; fi5++ {
            v4 = (!(-2.25 >= -2.25))
            for fi6 := 0; fi6 < 18; fi6++ {
                fmt.Println(fi6)
                _ = fi6
            }
            _ = fi5
        }
        v7 := GenS1{F0: 4294967295, F1: 51, F2: 0}
        _ = v7
    } else {
        fmt.Println(v4)
    }
    fmt.Println(v4)
    for fi8 := 0; fi8 < 3; fi8++ {
        v3[("hello" + "")] = true
        _ = fi8
    }
    var v9 uint64 = 2
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v10 uint64 = (v9 + v9)
    v11 := GenS1{F0: 42, F1: 34, F2: 255}
    switch ("a" + ("foo bar" + "\n\t")) {
    case "foo bar":
        for fi12 := 0; fi12 < 19; fi12++ {
            if (0.1 < 1.5) {
                var v13 float32 = ((((-2.25 * 1e10) - (((3.14159 / 1e10) + 1.5) + (1.5 * 0.1))) * float32(1.5)) * 0.1)
                fmt.Println(v4)
                if v4 {
                    var v14 uint8 = 19
                    v13 -= (0.1 + v13)
                    _ = v14
                } else {
                    var v15 uint64 = 42
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v9 += (v15 << ((730) & 63))
                        v15 = v15
                        _ = fi16
                    }
                    _ = v15
                }
                _ = v13
            }
            _ = fi12
        }
    case "\n\t":
        _ = len(v3)
    case "":
        var v17 int32 = -1
        _ = v17
    default:
        v4 = v4
    }
    var v18 int8 = 42
    fmt.Println(v10, v18)
    var v19 uint64 = (v9 + 4294967295)
    var v20 uint32 = 132
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
