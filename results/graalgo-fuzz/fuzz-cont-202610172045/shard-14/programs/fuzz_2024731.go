package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }

func main() {
    v1 := []int{-548, -2147483648, 128, 9223372036854775807}
    var v2 uint16 = 1
    v3 := make(map[uint32]string)
    var v4 bool = true
    _ = len(v3)
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(v4)
        if (!(1e10 >= 3.14159)) {
            var v6 float32 = 1.5
            if len(v1) > 0 {
                v1[0] = 1
            }
            var v7 int64 = 42
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if (!v4) {
            var v8 string = "\n\t"
            _ = v8
        }
    }
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 14; fi10++ {
        v4 = (!false)
        _ = fi10
    }
    switch "\n\t" {
    case "":
        v11 := GenS1{F0: -430}
        _ = v11
    case "\n\t":
        v3[898] = "hello"
    default:
        if v4 {
            var v12 float32 = 0.0
            v12 += v12
            _ = v12
        }
    }
    _ = len(v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v2, v4)
        if true {
            v9 = -786
            switch fi13 {
            case 54:
                v4 = (v4 || (v9 > v9))
            case -2147483648:
                switch (-(fi13)) {
                case 665:
                    v2 = (^(v2))
                case 1:
                    switch fi13 {
                    case 9223372036854775807:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v15 := GenS1{F0: -32768}
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            }
        } else {
            v16 := GenS2{F0: 1, F1: ("a" + "foo bar"), F2: 42}
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: -1}
                _ = fi17
                _ = v18
            }
            if len(v1) > 0 {
                v1[0] = fi13
            }
            _ = v16
        }
        switch (fi13 - (fi13 + (fi13 + fi13))) {
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -72:
            var v19 bool = (!v4)
            _ = v19
        }
        _ = fi13
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if false {
            var v21 int64 = 0
            fmt.Println(v21)
            _ = v21
        }
        fmt.Println(v4, v4, v4)
        delete(v3, 2147483647)
        _ = fi20
    }
    var v22 uint = 545
    fmt.Println(v9, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
