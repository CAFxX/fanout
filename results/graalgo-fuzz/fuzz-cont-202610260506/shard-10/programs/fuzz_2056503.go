package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 string }

func main() {
    v1 := []string{}
    var v2 uint64 = 0
    var v3 int8 = -1
    var v4 int16 = 42
    v2 = v2
    v5 := GenS1{F0: 1, F1: 9223372036854775807}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS2{F0: 396}
        _ = fi6
        _ = v7
    }
    var v8 int8 = v3
    var v9 int = 2147483647
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int = fi10
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    switch (v9 + 1) {
    case 1:
        v1 = append(v1, (("" + ("a" + "")) + "foo bar"))
    case -873:
        v13 := GenS2{F0: 2}
        _ = v13
    case 128:
        var v14 bool = false
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        switch fi15 {
        case 32768:
            switch "hello" {
            case "\n\t":
                v16 := GenS3{F0: 42, F1: (false || (-949 != 1)), F2: "foo bar"}
                _ = v16
            }
        default:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        _ = fi15
    }
    v4 = (v4 * (v4 << ((4294967295) & 63)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
