package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 bool; F2 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 int = 2147483647
    var v3 int8 = 1
    var v4 int8 = (v3 >> ((4294967295) & 63))
    v5 := []string{}
    v6 := make(map[int]string)
    switch (("foo bar" + (("a" + "foo bar") + "a")) + "foo bar") {
    case "hello":
        var v7 uint8 = 87
        _ = v7
    case "":
        for fi8 := 0; fi8 < 4; fi8++ {
            v3 += (^(v3))
            v5 = append(v5, (("foo bar" + (("" + "foo bar") + "hello")) + "foo bar"))
            if len(v5) > 0 {
                v5[0] = ("a" + "hello")
            }
            _ = fi8
        }
    case "\n\t":
        var v9 float64 = (0.0 + 0.1)
        _ = v9
    default:
        v4 += int8(v2)
    }
    var v10 uint8 = 0
    v11 := v6[v2]
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: -13, F1: 0, F2: false}
        _ = len(v5)
        _ = fi12
        _ = v13
    }
    var v14 bool = (42 > 200)
    for fi15 := 0; fi15 < 8; fi15++ {
        v2 = (fi15 >> ((4294967295) & 63))
        if v14 {
            _ = len(v6)
            var v16 uint32 = 513
            _ = v16
        }
        _ = fi15
    }
    v14 = v14
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 -= (fi17 | (v2 | v2))
            _ = fi18
        }
        _ = fi17
    }
    if (!(839 >= 2)) {
        fmt.Println(v10, v11)
        var v19 uint16 = v1
        fmt.Println(v11)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 string = v11
            for fi22 := 0; fi22 < 7; fi22++ {
                if false {
                    var v23 int32 = 128
                    _ = v23
                }
                v24 := GenS2{F0: 1, F1: 42, F2: 109}
                _ = fi22
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        v25 := GenS3{F0: 1}
        _ = v25
    }
    fmt.Println(v1)
    _ = len(v5)
    for fi26 := 0; fi26 < 19; fi26++ {
        switch fi26 {
        case 32767:
            fmt.Println(v10, v10)
        case 4294967295:
            v6[(fi26 * v2)] = ""
        case 42:
            v27 := GenS2{F0: 0, F1: 642, F2: 42}
            _ = v27
        }
        _ = fi26
    }
    v28 := GenS4{F0: v2, F1: ((!v14) || false), F2: (v3 ^ v4)}
    switch ((v11 + (v11 + "foo bar")) + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v4, v14)
    default:
        var v29 float32 = 3.14159
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
