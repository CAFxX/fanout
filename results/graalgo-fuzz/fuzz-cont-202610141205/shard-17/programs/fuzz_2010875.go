package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 uint }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 uint64 = 1
    v5 := []int{127, -1}
    var v6 bool = false
    fmt.Println(v6, v6, v4)
    fmt.Println(v1)
    var v7 bool = ("hello" >= "foo bar")
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint = 255
        _ = fi8
        _ = v9
    }
    switch 0 {
    case 886:
        switch ("a" + "foo bar") {
        case "":
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := GenS1{F0: -1, F1: (v1 % ((198) | 1))}
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 uint = 255
                    fmt.Println(v4)
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
        case "hello":
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(v6, v7)
                v15 := GenS1{F0: 87, F1: v1}
                _ = fi14
                _ = v15
            }
        default:
            v1 -= v1
        }
    case -9223372036854775808:
        var v16 bool = v7
        _ = v16
    case 2147483647:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v17 := GenS2{F0: 65535, F1: (v3 % (((v3 & 1)) | 1)), F2: 1}
    var v18 int = -9223372036854775808
    var v19 int64 = 1
    if v7 {
        var v20 int8 = 25
        _ = v20
    } else {
        v18 -= v18
        _ = len(v5)
        v21 := GenS2{F0: 894, F1: v3, F2: 4294967295}
        _ = v21
    }
    var v22 uint64 = 1
    var v23 float32 = (3.14159 - 0.1)
    for fi24 := 0; fi24 < 11; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            v23 += (v23 + (1.5 - v23))
            switch -73 {
            case -9223372036854775808:
                if (("hello" > ("foo bar" + "a")) && (v3 <= (^(v3)))) {
                    var v26 int8 = int8(v19)
                    _ = v26
                }
            }
            _ = fi25
        }
        var v27 int32 = v3
        _ = fi24
        _ = v27
    }
    if v7 {
        v3 = v3
        var v28 int64 = v19
        fmt.Println(v3, v4, v1)
        _ = v28
    }
    var v29 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
