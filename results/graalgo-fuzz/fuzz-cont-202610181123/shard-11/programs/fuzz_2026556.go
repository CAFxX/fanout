package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 2
    v3 := []string{"\n\t", "", ""}
    var v4 uint8 = v2
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v5 := GenS1{F0: 0, F1: uint(v2), F2: 42}
    fmt.Println(v4, v1)
    var v6 int16 = 1
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = 893
            fmt.Println(v4)
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint32 = 302
            _ = fi9
            _ = v10
        }
        if true {
            v11 := GenS1{F0: 42, F1: 255, F2: 2147483648}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v12 uint16 = v1
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            switch ("a" + "foo bar") {
            case "hello":
                fmt.Println(v2, fi13, fi13)
            }
            _ = fi14
        }
        _ = fi13
    }
    fmt.Println(v4, v6)
    if (4294967295 >= -9223372036854775808) {
        if (!(457 > 18446744073709551615)) {
            _ = len(v3)
            fmt.Println(v6, v6)
            v4 = (^(v2))
        }
        v15 := GenS2{F0: 4294967295, F1: 0}
        _ = v15
    } else {
        v3 = append(v3, "\n\t")
        switch "foo bar" {
        case "":
            switch "foo bar" {
            case "a":
                var v16 int = 964
                _ = v16
            case "\n\t":
                v2 = (v4 + (v2 % ((v2) | 1)))
            case "foo bar":
                var v17 uint16 = (v1 & 1)
                _ = v17
            default:
                var v18 uint64 = 0
                _ = v18
            }
        }
        v19 := GenS2{F0: 0, F1: -1}
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v20 uint8 = (v2 / ((v4) | 1))
    v21 := GenS1{F0: (-(v6)), F1: 506, F2: 297}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
