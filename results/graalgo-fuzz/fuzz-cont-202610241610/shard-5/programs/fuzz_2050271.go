package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    var v3 string = ("" + ("hello" + ("" + (("\n\t" + "\n\t") + "a"))))
    var v4 int64 = 1
    var v5 uint64 = 324
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 9223372036854775807}
        switch v3 {
        case "foo bar":
            var v8 float64 = -2.25
            _ = v8
        }
        fmt.Println(v5, fi6)
        _ = fi6
        _ = v7
    }
    v3 = "foo bar"
    var v9 float32 = -2.25
    switch v3 {
    case "\n\t":
        v4 = -413
    case "foo bar":
        var v10 uint32 = 2
        _ = v10
    case "a":
        fmt.Println(v4)
    default:
        switch v3 {
        case "\n\t":
            v3 = ("foo bar" + "foo bar")
        }
    }
    _ = len(v2)
    fmt.Println(v3, v5, v1)
    var v11 int32 = -1
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS2{F0: 949, F1: ((("foo bar" + ("foo bar" + v3)) + "foo bar") + v3)}
        if (1000 > -204) {
            var v14 int64 = v4
            v11 = (v11 << ((2) & 63))
            fmt.Println(v3, v11)
            _ = v14
        } else {
            if (true || (((true || ((v11 % ((42) | 1)) <= v11)) || (true && false)) || false)) {
                if false {
                    switch "hello" {
                    case "hello":
                        for fi15 := 0; fi15 < 4; fi15++ {
                            switch v3 {
                            case "\n\t":
                                v16 := GenS3{F0: 4294967295, F1: uint32(v5), F2: v5}
                                v17 := GenS3{F0: 65535, F1: 65535, F2: 2}
                                _ = v16
                                _ = v17
                            case "a":
                                fmt.Println(v3)
                            case "":
                                var v18 float64 = (0.1 - float64(v9))
                                _ = v18
                            default:
                                v11 = 1000
                            }
                            v19 := []int{(fi15 ^ fi12)}
                            _ = fi15
                            _ = v19
                        }
                    }
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v3 = v3
                        var v21 int16 = 32767
                        v22 := []int{128}
                        _ = len(v2)
                        v4 += -406
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                } else {
                    v2[(v3 + "a")] = 2
                    var v23 uint8 = v1
                    _ = v23
                }
                v1 -= 127
            }
        }
        _ = fi12
        _ = v13
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 uint16 = 32767
        _ = fi24
        _ = v25
    }
    var v26 float64 = ((((0.1 + (3.14159 * 0.1)) - 3.14159) - (3.14159 / 1.5)) * (float64(v9) - float64(v9)))
    v4 = v4
    var v27 uint64 = (^(v5))
    v1 = 127
    if false {
        var v28 string = v3
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
