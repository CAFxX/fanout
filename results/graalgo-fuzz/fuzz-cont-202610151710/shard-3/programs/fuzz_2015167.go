package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    v1 := []string{"foo bar", "hello", "foo bar"}
    var v2 string = ("foo bar" + ((("\n\t" + "a") + "") + "hello"))
    v3 := []int{4294967295, -879, 4294967295}
    if (255 >= 42) {
        switch v2 {
        case "hello":
            var v4 string = "foo bar"
            _ = v4
        default:
            var v5 string = v2
            _ = v5
        }
    }
    fmt.Println(v2)
    v2 = ("foo bar" + ((("\n\t" + ((v2 + "foo bar") + "a")) + v2) + ("hello" + "a")))
    switch -1 {
    case 32768:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint16 = 65535
            if false {
                v2 = v2
                fmt.Println(v7)
            }
            _ = fi6
            _ = v7
        }
    case -1:
        for fi8 := 0; fi8 < 10; fi8++ {
            if (4294967295 >= 2) {
                var v9 int32 = 2147483647
                fmt.Println(fi8, fi8)
                _ = v9
            } else {
                var v10 float64 = 1e10
                if ((fi8 << ((30) & 63)) >= fi8) {
                    var v11 int8 = 127
                    if false {
                        if false {
                            var v12 uint16 = 209
                            _ = v12
                        } else {
                            _ = len(v3)
                        }
                        v10 -= 0.1
                    }
                    _ = v11
                }
                _ = v10
            }
            _ = fi8
        }
    default:
        v13 := GenS1{F0: 2147483648}
        _ = v13
    }
    switch 342 {
    case 128:
        v2 = v2
    case 32768:
        fmt.Println(v2)
    }
    v14 := GenS2{F0: -1}
    _ = len(v3)
    v15 := GenS3{F0: 323, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
