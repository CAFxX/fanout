package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 uint32; F1 int; F2 uint }

func main() {
    v1 := []string{"\n\t", ("hello" + ("foo bar" + "")), "a", ("a" + "")}
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 int16 = 1000
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 0}
        v7 := GenS2{F0: 845, F1: 42, F2: ""}
        if false {
            switch "" {
            case "foo bar":
                fmt.Println(fi5, fi5)
            case "\n\t":
                v8 := v3["\n\t"]
                _ = v8
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v9 bool = false
    if v9 {
        v9 = (160 < 255)
        v10 := GenS1{F0: 65535}
        for fi11 := 0; fi11 < 2; fi11++ {
            v4 = 1
            _ = fi11
        }
        _ = v10
    }
    fmt.Println(v4, v4)
    v4 = -361
    v12 := GenS3{F0: 4294967295, F1: 456, F2: 1}
    _ = len(v1)
    v9 = (false && v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
