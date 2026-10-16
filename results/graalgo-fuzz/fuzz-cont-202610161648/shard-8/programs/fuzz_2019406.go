package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 bool; F1 string }

func main() {
    var v1 int16 = 1
    var v2 uint32 = 0
    v3 := []string{"\n\t", "foo bar", "hello", ("" + ((("foo bar" + "hello") + "\n\t") + "hello"))}
    v2 = v2
    fmt.Println(v1, v2)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v1 = v1
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    var v6 uint32 = v2
    v6 = v2
    var v7 int = 1
    for fi8 := 0; fi8 < 9; fi8++ {
        if true {
            v9 := GenS1{F0: -9223372036854775808, F1: -2147483648, F2: 32768}
            fmt.Println(v2)
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            _ = len(v3)
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            fmt.Println(v1, v7)
            _ = fi10
        }
        _ = fi8
    }
    v11 := GenS2{F0: false, F1: "hello"}
    for fi12 := 0; fi12 < 8; fi12++ {
        v2 = v6
        if true {
            fmt.Println(v1, fi12, fi12)
            var v13 uint16 = 156
            _ = v13
        }
        _ = fi12
    }
    v1 = v1
    var v14 uint32 = 65535
    switch 2147483647 {
    case -9223372036854775808:
        for fi15 := 0; fi15 < 16; fi15++ {
            switch 147 {
            case 32767:
                v2 += v2
            }
            v1 = v1
            v16 := GenS1{F0: -920, F1: -1, F2: (fi15 * (v7 ^ (v7 | -1)))}
            _ = fi15
            _ = v16
        }
    case 42:
        var v17 uint8 = 1
        _ = v17
    }
    var v18 uint = 255
    var v19 uint32 = v6
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
