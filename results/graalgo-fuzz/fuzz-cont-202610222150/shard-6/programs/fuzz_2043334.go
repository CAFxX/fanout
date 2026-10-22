package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }
type GenS3 struct { F0 int32; F1 int8; F2 uint16 }
func genFn0(p0 int16) int64 {
    switch "\n\t" {
    case "":
        var v4 bool = ("a" > ("" + "\n\t"))
        _ = v4
    case "foo bar":
        v5 := GenS1{F0: 0}
        _ = v5
    case "a":
        fmt.Println(p0)
    }
    v6 := []int{1}
    p0 += (p0 * p0)
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint8) string {
    var v7 uint16 = 255
    var v8 uint64 = 938
    _ = p0
    _ = v7
    _ = v8
    return ("\n\t" + "")
}

func main() {
    v1 := []string{"foo bar", "", "\n\t"}
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    var v9 uint = 255
    if false {
        if (!false) {
            fmt.Println(v2, v2)
            v10 := GenS2{F0: -32768, F1: 20, F2: (v9 & 4294967295)}
            _ = v10
        } else {
            v11 := GenS2{F0: -1, F1: 2, F2: 0}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: -725}
        v14 := GenS1{F0: fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    v9 = v9
    var v15 int = -1
    delete(v3, ("hello" + "foo bar"))
    v16 := GenS3{F0: 2147483647, F1: 82, F2: 0}
    v15 += v15
    v1 = append(v1, (("\n\t" + "") + "\n\t"))
    v9 = v9
    if false {
        var v17 uint8 = 42
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
