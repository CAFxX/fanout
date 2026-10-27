package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 bool }

func main() {
    v1 := []uint32{4294967295, 42, 2147483648}
    v2 := []string{"", "", "a", ("\n\t" + "\n\t")}
    var v3 bool = false
    var v4 int64 = -9223372036854775808
    var v5 int = 32767
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: fi7, F1: 935, F2: (v6 ^ (v6 % ((0) | 1)))}
        _ = fi7
        _ = v8
    }
    switch v5 {
    case -119:
        _ = len(v2)
    case -313:
        _ = len(v2)
    default:
        v3 = (!v3)
    }
    v5 -= (v5 % ((128) | 1))
    var v9 uint = 65535
    if len(v1) > 0 {
        v1[0] = 65535
    }
    switch "foo bar" {
    case "":
        v4 -= (^(v4))
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float32 = 3.14159
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch v5 {
        case 128:
            switch 9223372036854775807 {
            case 26:
                v4 -= v4
            case -1:
                switch (fi12 & 9223372036854775807) {
                case 0:
                    var v13 uint32 = uint32(v9)
                    _ = v13
                }
            }
        case -2147483648:
            v3 = v3
        default:
            v6 += 0
        }
        switch ("" + (("a" + "hello") + ("hello" + "\n\t"))) {
        case "foo bar":
            var v14 uint8 = 177
            _ = v14
        case "\n\t":
            var v15 string = "foo bar"
            _ = v15
        case "a":
            switch v5 {
            case -2147483648:
                fmt.Println(v3)
            case 42:
                var v16 uint = (v9 * (v9 ^ v9))
                _ = v16
            }
        }
        _ = fi12
    }
    var v17 float32 = 0.1
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if true {
        v18 := GenS1{F0: v5, F1: -495, F2: (^(v6))}
        v19 := GenS2{F0: "foo bar", F1: false, F2: v3}
        _ = v18
        _ = v19
    }
    var v20 bool = (!(32767 >= v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
