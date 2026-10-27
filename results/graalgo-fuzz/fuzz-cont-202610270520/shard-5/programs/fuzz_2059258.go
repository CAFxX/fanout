package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 uint16, p1 int32) bool {
    v6 := GenS1{F0: 42, F1: 164}
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{}
        v9 := GenS2{F0: "foo bar"}
        fmt.Println(p1, p0, p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: 42}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return true
}

func main() {
    v1 := []string{"", "\n\t", "foo bar", ((("hello" + "foo bar") + "") + "a")}
    v2 := make(map[string]bool)
    var v3 uint32 = 928
    var v4 bool = (!false)
    var v5 float64 = 1.5
    if v4 {
        var v11 uint = 1
        var v12 uint8 = 0
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v3, v5)
    v13 := GenS1{F0: 35, F1: 255}
    fmt.Println(v3)
    var v14 uint32 = 2147483648
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v14)
    }
    v15 := GenS4{F0: 836, F1: 553}
    switch 128 {
    case -918:
        fmt.Println(v5, v4)
    case 9223372036854775807:
        v16 := GenS3{F0: 223}
        _ = v16
    default:
        v17 := GenS3{F0: 0}
        _ = v17
    }
    v18 := v2["foo bar"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
