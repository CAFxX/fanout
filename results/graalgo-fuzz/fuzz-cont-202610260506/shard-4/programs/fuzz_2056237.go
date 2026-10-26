package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 int8 = -77
    v2 := []string{}
    var v3 float32 = 1e10
    var v4 int8 = (^(v1))
    v5 := GenS1{F0: 446, F1: -971}
    for fi6 := 0; fi6 < 8; fi6++ {
        if (!(-2.25 <= v3)) {
            if (-2147483648 < 2147483647) {
                var v7 int32 = 2147483647
                v2 = append(v2, "a")
                _ = v7
            } else {
                switch fi6 {
                case 596:
                    fmt.Println(v4, fi6, v1)
                case 560:
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 uint8 = 1
                        switch fi8 {
                        case 595:
                            fmt.Println(v1, v3, v1)
                        }
                        _ = fi8
                        _ = v9
                    }
                case 9223372036854775807:
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                }
            }
            _ = len(v2)
        } else {
            v2 = append(v2, "a")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            _ = len(v2)
            fmt.Println(v3, fi6)
            _ = fi10
        }
        _ = fi6
    }
    v1 -= v1
    fmt.Println(v4)
    if (-2147483648 <= 113) {
        v1 -= (v1 - (v4 - (v1 * 43)))
        for fi11 := 0; fi11 < 16; fi11++ {
            _ = len(v2)
            var v12 int = (fi11 << ((1) & 63))
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: -697, F1: 32767}
        _ = v13
    } else {
        v1 += v4
        v2 = append(v2, "foo bar")
        v4 -= (v4 % (((v1 | 127)) | 1))
    }
    switch 580 {
    case 2147483647:
        var v14 uint16 = 773
        _ = v14
    default:
        v3 -= v3
    }
    switch ((("hello" + (("\n\t" + "") + (("a" + ("hello" + ("foo bar" + ((("foo bar" + "a") + (("a" + "a") + "hello")) + "")))) + "foo bar"))) + "foo bar") + "a") {
    case "":
        switch ("a" + ("foo bar" + (("a" + "") + "a"))) {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            for fi15 := 0; fi15 < 10; fi15++ {
                if true {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v17 := GenS1{F0: 2147483647, F1: -32768}
                        v18 := GenS2{F0: 128, F1: ((true || (-2.25 >= 0.0)) && false)}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v4)
                    switch "foo bar" {
                    case "hello":
                        v3 = v3
                    default:
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    }
                } else {
                    v3 = v3
                    var v20 int = int(v1)
                    _ = v20
                }
                _ = fi15
            }
        }
    case "a":
        var v21 string = (("hello" + "foo bar") + "")
        _ = v21
    case "\n\t":
        var v22 bool = false
        _ = v22
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
