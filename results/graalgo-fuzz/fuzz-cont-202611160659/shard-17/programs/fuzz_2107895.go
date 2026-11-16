package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    var v1 int64 = -445
    var v2 int32 = 42
    var v3 uint8 = 2
    var v4 int16 = -32768
    v3 = 255
    v5 := []int{9223372036854775807, -554, 128, 9223372036854775807}
    switch -9223372036854775808 {
    case 128:
        var v6 float64 = 3.14159
        _ = v6
    }
    switch 0 {
    case -1:
        var v7 int = 1
        _ = v7
    case -977:
        var v8 uint8 = v3
        _ = v8
    case 1:
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint64 = 2
                v4 += 42
                switch "" {
                case "a":
                    switch (fi10 / ((fi10) | 1)) {
                    case 39:
                        v11 = (v11 * (v11 ^ 0))
                    }
                default:
                    if (true || false) {
                        var v12 uint16 = 32767
                        if false {
                            var v13 int = (fi10 + (fi10 ^ -631))
                            if len(v5) > 0 {
                                v5[0] = fi10
                            }
                            _ = v13
                        } else {
                            if ((false && (true || false)) || false) {
                                v14 := GenS1{F0: (^(v1)), F1: (v3 != (v3 / ((v3) | 1))), F2: ("\n\t" + ("foo bar" + "hello"))}
                                v3 = v3
                                _ = v14
                            }
                            var v15 uint32 = 2
                            _ = v15
                        }
                        _ = v12
                    } else {
                        var v16 float64 = (float64(0.0) + 3.14159)
                        _ = v16
                    }
                }
                _ = fi10
                _ = v11
            }
            fmt.Println(v2, v3)
            _ = fi9
        }
    default:
        v17 := GenS2{F0: 241}
        _ = v17
    }
    switch 32768 {
    case -9223372036854775808:
        v4 = (v4 * v4)
    case -1:
        v3 += v3
    }
    v18 := GenS3{F0: -1, F1: 945}
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS2{F0: fi19}
        var v21 uint64 = 287
        _ = fi19
        _ = v20
        _ = v21
    }
    switch -9223372036854775808 {
    case 2147483647:
        v3 = 127
    case -2147483648:
        var v22 float64 = 0.0
        _ = v22
    default:
        var v23 string = ((("hello" + (("" + "hello") + ("\n\t" + (("" + "hello") + "")))) + "") + ("" + (("foo bar" + ("\n\t" + "foo bar")) + ("hello" + ("hello" + "foo bar")))))
        _ = v23
    }
    v24 := GenS1{F0: v1, F1: true, F2: ("hello" + "\n\t")}
    fmt.Println(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
