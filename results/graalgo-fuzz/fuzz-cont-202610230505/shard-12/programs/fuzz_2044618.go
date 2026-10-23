package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    var v4 uint8 = 255
    v5 := []uint32{}
    v6 := []uint32{}
    var v7 int32 = 1
    if (("hello" > ("\n\t" + "hello")) || false) {
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                var v9 string = ("hello" + ("foo bar" + ("a" + "foo bar")))
                _ = v9
            }
            _ = fi8
        }
    } else {
        var v10 bool = false
        fmt.Println(v4)
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        delete(v3, fi11)
        _ = fi11
    }
    fmt.Println(v2, v4)
    if len(v6) > 0 {
        v6[0] = 141
    }
    switch ("" + "") {
    case "\n\t":
        if (!(2 < 152)) {
            v7 -= (v7 / (((-(v7))) | 1))
            v12 := GenS1{F0: 4294967295, F1: -368, F2: 132}
            _ = v12
        } else {
            _ = len(v5)
        }
    case "hello":
        v13 := GenS1{F0: uint32(v2), F1: -349, F2: 56}
        _ = v13
    }
    fmt.Println(v2, v7, v2)
    var v14 int = 42
    _ = len(v3)
    v14 += v14
    switch (v14 + 1) {
    case -1:
        v15 := GenS1{F0: 2, F1: 412, F2: 107}
        _ = v15
    case 42:
        v2 -= v2
    }
    var v16 uint64 = (v2 | (^(v2)))
    v3[v14] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
