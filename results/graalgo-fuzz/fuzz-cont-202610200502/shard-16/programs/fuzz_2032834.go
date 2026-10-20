package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 float32 = 1.5
    v2 := []string{"foo bar", ""}
    var v3 uint8 = 138
    v4 := make(map[int]uint32)
    _ = len(v4)
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: "\n\t", F1: 255}
        v1 += float32(0.1)
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                switch "a" {
                case "":
                    v3 -= 1
                    v3 -= (v3 % ((v3) | 1))
                    v3 = 190
                case "a":
                    v2 = append(v2, ((("" + "\n\t") + "a") + (("a" + "foo bar") + "a")))
                case "\n\t":
                    _ = len(v2)
                }
                if (v3 < v3) {
                    v8 := GenS2{F0: -1}
                    _ = v8
                }
                v3 = v3
                _ = fi7
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int64 = -713
                _ = fi9
                _ = v10
            }
        } else {
            v11 := GenS3{F0: false, F1: (v3 ^ v3)}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    if true {
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            if (((0.0 / 0.1) >= 1.5) && true) {
                var v12 float32 = (v1 / v1)
                var v13 int32 = -2147483648
                _ = v12
                _ = v13
            } else {
                v14 := GenS3{F0: (!(1 < 4294967295)), F1: 40}
                _ = len(v4)
                _ = v14
            }
        default:
            v3 = 0
        }
    }
    var v15 uint = 65535
    v1 -= 3.14159
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(fi16)
            fmt.Println(fi17, fi16)
            fmt.Println(v1)
            fmt.Println(fi16, v15)
            fmt.Println(v3, fi16, v3)
            _ = fi17
        }
        v1 += v1
        var v18 int = (fi16 + 127)
        _ = fi16
        _ = v18
    }
    if (0 == -1) {
        var v19 int16 = 0
        var v20 int = 42
        _ = v19
        _ = v20
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
