package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int }
func genFn0(p0 uint16, p1 uint, p2 int32) int8 {
    v5 := GenS1{F0: 9223372036854775807, F1: p0, F2: 58}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int16 = -1
    var v2 uint16 = 42
    var v3 uint8 = 1
    var v4 uint8 = v3
    if false {
        if false {
            if false {
                fmt.Println(v1, v4, v2)
                if (18446744073709551615 >= 2) {
                    if (v1 < 42) {
                        v6 := GenS2{F0: 127}
                        fmt.Println(v1, v3, v4)
                        _ = v6
                    }
                    var v7 float32 = (0.0 - 1e10)
                    _ = v7
                }
                v8 := []int{-135, 4294967295}
                _ = v8
            } else {
                v9 := []int{32768, int(v1), 714, 2147483647}
                v1 += v1
                _ = v9
            }
            switch 1 {
            case -838:
                switch 32767 {
                case 680:
                    switch "\n\t" {
                    case "\n\t":
                        v3 = 127
                    }
                case 32768:
                    v2 -= v2
                default:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v11 := []int{(fi10 | fi10), -290, fi10}
                        for fi12 := 0; fi12 < 4; fi12++ {
                            v13 := GenS2{F0: 127}
                            var v14 int16 = v1
                            if (((("\n\t" + "") + "hello") + "") != ("hello" + (((("foo bar" + "\n\t") + "foo bar") + ("a" + ("\n\t" + "hello"))) + "foo bar"))) {
                                v1 = -825
                            }
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = fi10
                        _ = v11
                    }
                }
            case 492:
                fmt.Println(v3, v2)
            case 42:
                v15 := GenS3{F0: -32768, F1: v2}
                _ = v15
            }
        }
        v1 -= (v1 + (v1 / ((v1) | 1)))
        switch -534 {
        case -9223372036854775808:
            switch "a" {
            case "":
                var v16 int64 = 9223372036854775807
                _ = v16
            }
        case 32768:
            fmt.Println(v2, v3)
        case 46:
            var v17 uint32 = 667
            _ = v17
        }
    } else {
        v4 += 127
        v1 = (v1 % (((^(v1))) | 1))
    }
    v18 := GenS4{F0: 1, F1: 32767}
    v3 -= (v3 / ((v3) | 1))
    v4 = (^(v3))
    v19 := GenS3{F0: (v1 / (((v1 >> ((1) & 63))) | 1)), F1: 255}
    fmt.Println(v4, v2)
    var v20 float64 = 0.1
    v3 -= (v4 >> ((4294967295) & 63))
    var v21 uint32 = 65535
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
