package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    var v3 uint16 = 42
    var v4 uint = 65535
    v4 -= v4
    v5 := []int{0, 934, 32768, -764}
    var v6 float64 = (1e10 / (-2.25 / -2.25))
    var v7 string = ((("" + (v2 + ((v2 + (v2 + v2)) + v2))) + v2) + "a")
    v3 = v3
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint = (^(v4))
            fmt.Println(v4)
            _ = fi9
            _ = v10
        }
        var v11 uint = 841
        var v12 float32 = 3.14159
        _ = fi8
        _ = v11
        _ = v12
    }
    fmt.Println(v6, v4)
    var v13 bool = (42 != 2147483647)
    var v14 bool = (((42 != 0) && false) || v13)
    for fi15 := 0; fi15 < 9; fi15++ {
        switch 32768 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 11; fi16++ {
                fmt.Println(v14)
                _ = fi16
            }
        case 128:
            var v17 bool = (!(42 < -1))
            _ = v17
        default:
            var v18 uint64 = 2
            _ = v18
        }
        _ = fi15
    }
    v7 = ""
    for fi19 := 0; fi19 < 16; fi19++ {
        if (!v14) {
            var v20 int32 = 1
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v2)
                if v13 {
                    v22 := v1[0]
                    var v23 int64 = -1
                    if (1 != 546) {
                        v4 -= (v4 * (v4 | v4))
                    } else {
                        v6 += (1.5 - 3.14159)
                        v5 = append(v5, (fi19 / ((32767) | 1)))
                        v5 = append(v5, (^(fi21)))
                    }
                    switch ((v2 + ("foo bar" + v2)) + "hello") {
                    case "foo bar":
                        v6 = 1e10
                    case "":
                        v6 += (3.14159 * v6)
                    case "hello":
                        _ = len(v5)
                    }
                    v24 := GenS1{F0: uint(v3), F1: v4}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v25 int32 = 1
                _ = fi21
                _ = v25
            }
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
