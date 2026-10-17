package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int32 = -713
    var v2 uint32 = 65535
    var v3 uint8 = 128
    v4 := []string{("" + (("" + ("foo bar" + "\n\t")) + ("hello" + "\n\t")))}
    _ = len(v4)
    switch "hello" {
    case "foo bar":
        v1 += v1
    case "hello":
        for fi5 := 0; fi5 < 13; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := GenS1{F0: v2, F1: 42}
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    }
    if ((v3 & (v3 % ((v3) | 1))) != v3) {
        if (("" + "foo bar") < "\n\t") {
            v3 = (^(v3))
        } else {
            var v8 float64 = 3.14159
            v1 = 42
            var v9 uint8 = v3
            _ = v8
            _ = v9
        }
        fmt.Println(v3, v1)
        switch "" {
        case "a":
            switch -9223372036854775808 {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 5; fi10++ {
                    v2 += v2
                    _ = fi10
                }
            case 32768:
                for fi11 := 0; fi11 < 3; fi11++ {
                    v1 -= v1
                    _ = fi11
                }
            case -369:
                switch -553 {
                case 32767:
                    if (!true) {
                        var v12 uint = 42
                        _ = v12
                    }
                }
            }
        default:
            var v13 bool = (("foo bar" + "\n\t") > "foo bar")
            _ = v13
        }
    }
    if true {
        v14 := GenS2{F0: "", F1: 188}
        var v15 uint64 = 2
        _ = v14
        _ = v15
    }
    var v16 int8 = 42
    v17 := GenS3{F0: 1, F1: 0}
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: (v2 | v2), F1: 18446744073709551615}
        _ = fi18
        _ = v19
    }
    var v20 uint32 = (v2 >> ((1) & 63))
    var v21 float64 = ((0.0 + 0.0) * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
