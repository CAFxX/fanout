package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }

func main() {
    var v1 int16 = -100
    v2 := []string{"\n\t"}
    var v3 string = "\n\t"
    v4 := make(map[string]string)
    var v5 uint32 = 42
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 -= v5
        if len(v2) > 0 {
            v2[0] = v3
        }
        _ = fi7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "hello":
        if ((42 >= 1) && (-1 != -128)) {
            v8 := v4["a"]
            v1 = (v6 >> ((4294967295) & 63))
            var v9 uint16 = 717
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: "foo bar"}
            v11 := GenS2{F0: v5}
            v12 := GenS2{F0: v5}
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        switch "a" {
        case "hello":
            fmt.Println(v1, v5)
        case "foo bar":
            v13 := GenS1{F0: v3}
            _ = v13
        case "":
            var v14 uint = 540
            _ = v14
        }
    }
    if false {
        fmt.Println(v3, v3, v1)
        var v15 uint64 = 2
        if true {
            var v16 int32 = 42
            v2 = append(v2, v3)
            v15 -= v15
            _ = v16
        } else {
            v15 = uint64(v5)
            v5 = (v5 + (^(v5)))
        }
        _ = v15
    } else {
        fmt.Println(v6)
        v2 = append(v2, ((v3 + v3) + "\n\t"))
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS3{F0: 1, F1: 32767, F2: 749}
        _ = fi17
        _ = v18
    }
    if true {
        v3 = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
