package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 int16 }

func main() {
    var v1 string = ("" + ("foo bar" + "\n\t"))
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    v4 := []int{-2147483648}
    for fi5 := 0; fi5 < 1; fi5++ {
        v1 = v1
        v1 = "a"
        if (0 != 255) {
            fmt.Println(v1, fi5)
        }
        _ = fi5
    }
    if true {
        var v6 uint64 = 18446744073709551615
        if true {
            if true {
                v4 = append(v4, 920)
                var v7 int64 = 0
                if (42 > 557) {
                    v8 := GenS1{F0: 4294967295, F1: true}
                    var v9 int32 = -2147483648
                    switch (("hello" + v1) + v1) {
                    case "foo bar":
                        v10 := GenS2{F0: 564, F1: 54, F2: -32768}
                        _ = v10
                    case "\n\t":
                        _ = len(v2)
                    case "hello":
                        v11 := GenS1{F0: 127, F1: false}
                        _ = v11
                    }
                    _ = v8
                    _ = v9
                }
                _ = v7
            }
            fmt.Println(v6)
        }
        _ = v6
    } else {
        if false {
            fmt.Println(v1)
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            switch "\n\t" {
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            fmt.Println(fi12, fi13, fi12)
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v1)
    v1 = v1
    if (1e10 >= (((1e10 + (float64((((-2.25 - -2.25) * 1.5) * -2.25)) * 0.0)) + 0.1) * float64(1e10))) {
        var v14 uint16 = 0
        _ = v14
    } else {
        switch 4294967295 {
        case 702:
            v1 = "hello"
        }
        fmt.Println(v1)
        for fi15 := 0; fi15 < 13; fi15++ {
            if len(v4) > 0 {
                v4[0] = (fi15 - fi15)
            }
            _ = fi15
        }
    }
    if false {
        v1 = ""
        fmt.Println(v1)
    }
    v1 = v1
    v16 := GenS2{F0: 127, F1: 1, F2: -591}
    var v17 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
