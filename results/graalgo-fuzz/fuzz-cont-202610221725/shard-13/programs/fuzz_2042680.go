package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint, p1 int64) int32 {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint64 = 2
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return -693
}
func genFn1(p0 int16, p1 int8, p2 uint) int8 {
    var v7 int8 = p1
    fmt.Println(v7, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 100
}

func main() {
    var v1 uint32 = 170
    var v2 string = ("" + "a")
    v3 := make(map[string]bool)
    var v4 string = v2
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 20; fi9++ {
        v4 = v4
        _ = fi9
    }
    var v10 int8 = 0
    var v11 bool = (-397 > 592)
    if (v11 && v11) {
        fmt.Println(v1)
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                if ((("" + v2) >= v2) && v11) {
                    var v14 int64 = 0
                    _ = v14
                }
                _ = fi13
            }
            var v15 string = v2
            _ = fi12
            _ = v15
        }
    }
    switch "a" {
    case "":
        fmt.Println(v10)
    case "foo bar":
        if v11 {
            var v16 int8 = (v10 & v10)
            _ = v16
        } else {
            switch v4 {
            case "hello":
                v17 := GenS1{F0: 531, F1: 0}
                _ = v17
            case "\n\t":
                var v18 uint64 = 721
                _ = v18
            }
        }
    case "hello":
        v2 = v4
    }
    v10 += genFn1(42, 42, 626)
    v4 = v4
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v11)
        v20 := GenS1{F0: -980, F1: 4294967295}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v4, fi21)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
