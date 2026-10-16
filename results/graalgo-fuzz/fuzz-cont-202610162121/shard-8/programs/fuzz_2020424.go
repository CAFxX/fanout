package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []uint32{0}
    var v2 float32 = 3.14159
    var v3 int64 = 9223372036854775807
    var v4 float32 = 0.1
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 10; fi6++ {
        _ = len(v1)
        _ = fi6
    }
    v5 = v5
    v2 = 1e10
    v7 := GenS1{F0: 731}
    var v8 int8 = 1
    v9 := GenS2{F0: -1, F1: 42, F2: 7}
    v10 := GenS2{F0: 0, F1: 659, F2: 229}
    switch "hello" {
    case "\n\t":
        for fi11 := 0; fi11 < 4; fi11++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
    case "a":
        for fi12 := 0; fi12 < 18; fi12++ {
            v5 = v5
            fmt.Println(v5, v3)
            _ = fi12
        }
    case "foo bar":
        var v13 string = ""
        _ = v13
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            switch (fi14 & (fi14 ^ (fi14 & -9223372036854775808))) {
            case 2147483647:
                var v15 string = ""
                _ = v15
            case 345:
                if len(v1) > 0 {
                    v1[0] = 2
                }
            case 555:
                _ = len(v1)
            }
            if (false || false) {
                v5 -= (^(v5))
                switch ("hello" + "\n\t") {
                case "hello":
                    v5 = v5
                case "\n\t":
                    v16 := GenS2{F0: -95, F1: 686, F2: -1}
                    _ = v16
                }
                switch (fi14 + fi14) {
                case 1:
                    var v17 int16 = -32768
                    _ = v17
                case 128:
                    v4 -= v4
                case 9223372036854775807:
                    v8 += v8
                default:
                    if (!((v3 % ((1) | 1)) != v3)) {
                        switch (("a" + "") + "foo bar") {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        }
                        v18 := GenS2{F0: 42, F1: -2147483648, F2: 42}
                        v3 = (v3 / ((v3) | 1))
                        _ = v18
                    } else {
                        var v19 string = (((("hello" + ("" + ("" + ("\n\t" + "foo bar")))) + "") + "") + ("a" + "a"))
                        v20 := GenS3{F0: 115}
                        var v21 bool = (true || ((int8(v3) <= 57) && (v4 < v2)))
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
            }
            _ = fi14
        }
    }
    v1 = append(v1, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
