package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }
func genFn0(p0 uint8) string {
    v4 := GenS1{F0: 1}
    p0 = 127
    if false {
        var v5 uint32 = 564
        fmt.Println(p0)
        v5 = 2147483647
        _ = v5
    }
    if true {
        p0 += 42
        if false {
            switch "hello" {
            case "a":
                p0 = p0
            case "":
                v6 := GenS1{F0: -1}
                _ = v6
            }
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = v4
    return ("foo bar" + ("" + ("" + ("" + "a"))))
}

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 4294967295
    v3 := []string{(("\n\t" + ("" + "")) + ("" + ""))}
    var v7 int32 = 2147483647
    if true {
        var v8 uint64 = (^(v2))
        v9 := GenS2{F0: 712, F1: 316}
        _ = v8
        _ = v9
    } else {
        var v10 int = 128
        var v11 int8 = -106
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int64 = -874
        _ = fi12
        _ = v13
    }
    var v14 float64 = 0.1
    var v15 float32 = (3.14159 + 1.5)
    v3 = append(v3, "a")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 += 0.0
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v7 = (v7 + (v7 | v7))
    v16 := GenS3{F0: true, F1: 1, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
