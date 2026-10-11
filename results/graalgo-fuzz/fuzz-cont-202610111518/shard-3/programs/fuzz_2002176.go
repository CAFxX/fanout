package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16) uint32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p0)
        p0 = 206
        _ = fi7
    }
    var v8 uint32 = 2
    v8 = (v8 ^ (v8 ^ 2147483647))
    _ = p0
    _ = v8
    return (^(v8))
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 32767
    v3 := make(map[string]string)
    v4 := []string{"a", ("a" + "foo bar"), (("\n\t" + "foo bar") + "")}
    v5 := []uint32{548}
    v6 := []uint32{42, 195, 614}
    v9 := GenS1{F0: false, F1: -786}
    var v10 int32 = -988
    if false {
        if (!true) {
            _ = len(v3)
        }
        switch "\n\t" {
        case "hello":
            for fi11 := 0; fi11 < 16; fi11++ {
                delete(v3, ((("" + (("a" + "\n\t") + "a")) + "hello") + "\n\t"))
                var v12 int = (fi11 % ((-9223372036854775808) | 1))
                v10 += int32(v2)
                _ = fi11
                _ = v12
            }
        case "foo bar":
            var v13 int32 = -930
            _ = v13
        case "\n\t":
            if true {
                v1 = (v1 >> ((18446744073709551615) & 63))
                switch "\n\t" {
                case "\n\t":
                    var v14 float32 = 1e10
                    _ = v14
                default:
                    v3[("a" + "foo bar")] = "foo bar"
                }
            } else {
                v10 = (^(v10))
            }
        default:
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(fi15)
                var v16 uint64 = (v1 ^ 344)
                _ = fi15
                _ = v16
            }
        }
    } else {
        v1 -= 665
    }
    var v17 string = ""
    fmt.Println(v1, v1)
    var v18 int32 = -2147483648
    v2 += v2
    fmt.Println(v1)
    var v19 float64 = float64(1e10)
    fmt.Println(v17)
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 uint32 = 65535
        _ = fi20
        _ = v21
    }
    v10 = v18
    v19 = v19
    fmt.Println(v1, v19, v2)
    for fi22 := 0; fi22 < 4; fi22++ {
        switch "hello" {
        case "a":
            if (true && false) {
                var v23 uint32 = uint32(v1)
                _ = v23
            }
        case "":
            v17 = v17
        case "\n\t":
            switch v17 {
            case "a":
                for fi24 := 0; fi24 < 1; fi24++ {
                    v1 = v1
                    _ = fi24
                }
            case "":
                var v25 float32 = (((1.5 * 3.14159) - 3.14159) / 1e10)
                _ = v25
            }
        default:
            var v26 int64 = -802
            _ = v26
        }
        _ = fi22
    }
    var v27 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
