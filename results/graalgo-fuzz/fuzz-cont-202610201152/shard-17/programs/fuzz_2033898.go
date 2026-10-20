package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        switch -2147483648 {
        case -406:
            v6 := GenS1{F0: "\n\t", F1: 57, F2: "foo bar"}
            _ = v6
        case -197:
            var v7 float64 = (-2.25 * float64((0.0 / -2.25)))
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = ""
        if (!(!(!(-1 != 15)))) {
            var v10 float64 = 1.5
            _ = v10
        }
        var v11 int = (fi8 | (fi8 >> ((1) & 63)))
        _ = fi8
        _ = v11
        _ = v9
    }
    var v12 int8 = -123
    _ = v12
    return (v12 + 42)
}

func main() {
    v1 := []string{"foo bar", "\n\t", "\n\t", "a"}
    v2 := []uint32{65535, 1}
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    delete(v4, 9223372036854775807)
    var v13 bool = true
    v14 := GenS1{F0: "foo bar", F1: 660, F2: ("hello" + "foo bar")}
    var v15 uint8 = 128
    v16 := GenS2{F0: v15, F1: -9223372036854775808}
    var v17 float32 = -2.25
    var v18 uint8 = (v15 | v15)
    v19 := GenS1{F0: ("a" + ""), F1: 100, F2: "\n\t"}
    switch 1 {
    case 2147483647:
        v20 := GenS1{F0: "", F1: 69, F2: (("" + (("a" + ("" + "\n\t")) + "foo bar")) + ("" + ("" + ("foo bar" + ""))))}
        _ = v20
    case 9223372036854775807:
        switch "" {
        case "\n\t":
            v1 = append(v1, (("\n\t" + "hello") + "foo bar"))
        }
    }
    v13 = ((0 < 2147483648) && (!false))
    v15 -= 1
    v21 := GenS3{F0: 42, F1: 115}
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 float32 = float32(0.1)
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
