package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int16 = -1
    v2 := make(map[string]bool)
    v3 := []string{"\n\t", "foo bar", "hello"}
    v4 := []uint32{65535, 712, 2, 42}
    v5 := []uint32{422, 42}
    var v6 int32 = 42
    var v7 int = 72
    fmt.Println(v7, v6, v7)
    switch -2147483648 {
    case 9223372036854775807:
        var v8 uint8 = 42
        _ = v8
    case -9223372036854775808:
        _ = len(v4)
    case 4:
        v6 -= 2147483647
    default:
        if false {
            var v9 int = v7
            v9 = v7
            _ = v9
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                if (false && false) {
                    delete(v2, "hello")
                    v11 := GenS1{F0: 42, F1: (^(v7))}
                    _ = v11
                }
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi10
            }
            v12 := GenS1{F0: -1, F1: (^(v7))}
            _ = v12
        }
    }
    if true {
        v13 := GenS1{F0: 42, F1: v7}
        if true {
            var v14 int8 = 127
            var v15 uint8 = 118
            fmt.Println(v6, v1)
            _ = v14
            _ = v15
        }
        _ = v13
    } else {
        var v16 int16 = (v1 & v1)
        _ = v16
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        if ((v6 | (v6 % ((202) | 1))) > v6) {
            switch "foo bar" {
            case "\n\t":
                if (2 > 477) {
                    var v17 float64 = float64(-2.25)
                    v1 = v1
                    for fi18 := 0; fi18 < 10; fi18++ {
                        if true {
                            fmt.Println(fi18, v1, v1)
                        }
                        _ = len(v3)
                        _ = fi18
                    }
                    _ = v17
                }
            case "hello":
                delete(v2, "\n\t")
            }
            v19 := GenS1{F0: -1, F1: (-(v7))}
            var v20 uint64 = 0
            _ = v19
            _ = v20
        }
    default:
        switch 128 {
        case 76:
            for fi21 := 0; fi21 < 2; fi21++ {
                v3 = append(v3, (("hello" + "\n\t") + ("a" + "a")))
                switch "hello" {
                case "\n\t":
                    if true {
                        var v22 float64 = float64(0.1)
                        fmt.Println(v1, v6)
                        _ = v22
                    } else {
                        if true {
                            v23 := GenS1{F0: -1, F1: 128}
                            v5 = append(v5, 817)
                            var v24 bool = (42 < 2)
                            _ = v23
                            _ = v24
                        }
                    }
                case "":
                    v25 := GenS2{F0: 743, F1: (true || (!(1 == -1))), F2: 1}
                    _ = v25
                }
                _ = fi21
            }
        }
    }
    v26 := GenS1{F0: 1, F1: (v7 & v7)}
    v27 := GenS2{F0: 65535, F1: (!(!((42 >= 0) || false))), F2: (v1 >> ((42) & 63))}
    v28 := GenS3{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
