package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 int32, p2 uint8) int {
    var v4 uint32 = 0
    if p0 {
        if false {
            var v5 float64 = 0.0
            var v6 uint8 = 211
            fmt.Println(p0)
            _ = v5
            _ = v6
        } else {
            if p0 {
                switch "hello" {
                case "hello":
                    if (!(-1 < 1)) {
                        v7 := []int{-897, 42}
                        _ = v7
                    } else {
                        v8 := []int{}
                        _ = v8
                    }
                case "":
                    switch 9223372036854775807 {
                    case -33:
                        var v9 uint32 = (v4 / ((v4) | 1))
                        _ = v9
                    case -389:
                        switch 127 {
                        case 128:
                            p2 = (p2 * (p2 & (p2 >> ((0) & 63))))
                        default:
                            fmt.Println(p1)
                        }
                    case 42:
                        v4 = 65535
                    }
                case "\n\t":
                    var v10 float64 = (float64((1e10 + 3.14159)) + 0.0)
                    _ = v10
                }
                v4 += (v4 % ((2147483648) | 1))
            }
            p2 -= p2
        }
    } else {
        var v11 uint64 = 4294967295
        v12 := GenS1{F0: "hello", F1: (^(v4)), F2: v11}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 int = 396
    v2 := []string{"\n\t", ("hello" + (("\n\t" + (("\n\t" + "\n\t") + "\n\t")) + ("a" + ((("\n\t" + (("hello" + "foo bar") + "foo bar")) + "a") + ("a" + "\n\t")))))}
    v3 := make(map[uint32]bool)
    var v13 int = (-(v1))
    v14 := GenS2{F0: -798}
    v1 += v1
    if false {
        if false {
            fmt.Println(v13)
            switch (v1 + 42) {
            case -9223372036854775808:
                v15 := GenS3{F0: 992}
                _ = v15
            }
            switch 32768 {
            case 9223372036854775807:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 32767:
                v16 := GenS2{F0: 519}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            if (388 != 42) {
                var v18 uint8 = 1
                var v19 int32 = 42
                delete(v3, 4294967295)
                _ = v18
                _ = v19
            }
            var v20 uint32 = 65535
            _ = fi17
            _ = v20
        }
    } else {
        fmt.Println(v1)
    }
    var v21 int16 = 1
    var v22 int64 = 118
    switch -577 {
    case 4294967295:
        fmt.Println(v1, v21, v22)
    case -900:
        v23 := GenS1{F0: "foo bar", F1: 2, F2: 863}
        _ = v23
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int32 = -586
            fmt.Println(v25, v13, v25)
            _ = fi24
            _ = v25
        }
    }
    _ = len(v3)
    _ = len(v3)
    v13 += (v13 * 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
