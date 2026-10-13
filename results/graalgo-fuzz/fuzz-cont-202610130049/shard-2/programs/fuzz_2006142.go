package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0() int64 {
    v4 := GenS1{F0: 9223372036854775807, F1: 255, F2: 128}
    var v5 float64 = ((float64(3.14159) / 1.0) - 3.14159)
    v6 := GenS1{F0: 0, F1: 2, F2: 2}
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []string{(("\n\t" + "foo bar") + ("\n\t" + "\n\t")), "", ""}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    _ = len(v2)
    v7 := v2[726]
    v8 := GenS1{F0: -9223372036854775808, F1: 42, F2: 128}
    if false {
        var v9 uint = 960
        fmt.Println(v9)
        v9 += 255
        _ = v9
    }
    switch 194 {
    case -998:
        switch 2147483647 {
        case -1:
            delete(v2, v7)
        }
    case 128:
        v7 += (v7 / (((v7 - (v7 * v7))) | 1))
    case 2147483647:
        v7 = v7
    }
    v7 += 42
    v10 := GenS2{F0: 602}
    if false {
        fmt.Println(v7)
        v7 = v7
        delete(v2, v7)
    } else {
        if false {
            v1 = append(v1, "a")
            if (0 == 255) {
                var v11 uint16 = 0
                fmt.Println(v11, v7)
                switch "\n\t" {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                case "foo bar":
                    switch 128 {
                    case -406:
                        v12 := GenS2{F0: 4294967295}
                        _ = v12
                    case 4294967295:
                        v11 += (v11 - 414)
                    case 9223372036854775807:
                        fmt.Println(v7, v11)
                    default:
                        var v13 int = -1
                        _ = v13
                    }
                case "hello":
                    fmt.Println(v11)
                }
                _ = v11
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(fi14, fi14)
                    v7 -= v7
                    _ = fi14
                }
                v2[v7] = 1
                v15 := GenS2{F0: 448}
                _ = v15
            }
        } else {
            v7 -= (v7 + (^(v7)))
            fmt.Println(v7)
        }
        v16 := GenS3{F0: 100, F1: (true || false), F2: 18446744073709551615}
        v1 = append(v1, "")
        _ = v16
    }
    if (false || true) {
        var v17 int32 = 2147483647
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS3{F0: 0, F1: true, F2: 4294967295}
                var v21 int16 = -1
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(fi18, fi18, fi18)
                _ = fi22
            }
            _ = fi18
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((false && ((^(v7)) < v7)) && ((v7 < v7) && (255 == 42))) {
            if true {
                fmt.Println(v7)
                for fi23 := 0; fi23 < 7; fi23++ {
                    v7 -= 42
                    for fi24 := 0; fi24 < 1; fi24++ {
                        _ = len(v2)
                        _ = fi24
                    }
                    _ = fi23
                }
            } else {
                var v25 string = ("a" + "")
                var v26 int = 9223372036854775807
                _ = v25
                _ = v26
            }
        } else {
            var v27 uint16 = 32767
            v7 = (v7 * v7)
            _ = v27
        }
    }
    fmt.Println(v7)
    v28 := GenS3{F0: 19, F1: (!(42 > 2)), F2: 1}
    v7 = 870
    switch (((("a" + "foo bar") + "foo bar") + "\n\t") + ("foo bar" + "\n\t")) {
    case "a":
        var v29 uint8 = 0
        _ = v29
    default:
        v7 -= 2147483647
    }
    var v30 int32 = 0
    v1 = append(v1, ("\n\t" + "\n\t"))
    var v31 uint8 = 34
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v7
    _ = v8
}
