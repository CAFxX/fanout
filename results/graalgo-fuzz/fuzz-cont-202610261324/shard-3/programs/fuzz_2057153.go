package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 string }

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[string]bool)
    var v3 bool = ((("a" + ("a" + ("\n\t" + ("" + "")))) + "") != "\n\t")
    var v4 uint32 = 0
    v3 = (54 == 255)
    var v5 uint64 = 0
    var v6 uint32 = (v4 + v4)
    for fi7 := 0; fi7 < 17; fi7++ {
        if (true || true) {
            v5 -= (v5 + v5)
            for fi8 := 0; fi8 < 5; fi8++ {
                v5 = 2
                _ = fi8
            }
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int8 = -87
                _ = fi9
                _ = v10
            }
        } else {
            v6 += (^(v4))
            switch (("hello" + "hello") + "a") {
            case "hello":
                if v3 {
                    var v11 int = fi7
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := GenS1{F0: 1, F1: v5, F2: ""}
                        switch "\n\t" {
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v14 int64 = -9223372036854775808
                    _ = v11
                    _ = v14
                }
            case "a":
                v15 := GenS1{F0: 225, F1: v5, F2: ("a" + (("" + "a") + "hello"))}
                _ = v15
            case "foo bar":
                var v16 uint32 = (v6 & (^(v4)))
                _ = v16
            default:
                switch "\n\t" {
                case "":
                    v6 = 2147483648
                case "a":
                    fmt.Println(v3)
                }
            }
        }
        _ = fi7
    }
    v17 := GenS2{F0: 0}
    var v18 uint16 = 32767
    v4 += (v4 >> ((42) & 63))
    v2["foo bar"] = true
    var v19 bool = (-1 > 127)
    if false {
        v2["\n\t"] = v3
    }
    if v3 {
        fmt.Println(v4, v19)
    } else {
        if (1 <= 128) {
            var v20 int32 = 2147483647
            _ = v20
        }
        switch ("a" + ("\n\t" + "\n\t")) {
        case "hello":
            v21 := GenS3{F0: 1000, F1: ("" + "foo bar")}
            _ = v21
        case "":
            var v22 float32 = float32(((0.0 / (((0.1 - ((-2.25 - (((0.1 - 1e10) - 1e10) + 1e10)) + 0.1)) / (((1.5 * 0.0) + -2.25) * 1.5)) / float64((0.1 * (0.1 + -2.25))))) * 1.5))
            _ = v22
        case "\n\t":
            switch -1 {
            case -2147483648:
                switch ("" + "\n\t") {
                case "hello":
                    var v23 bool = v3
                    _ = v23
                default:
                    v19 = v3
                }
            default:
                if false {
                    var v24 int64 = -9223372036854775808
                    _ = v24
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
