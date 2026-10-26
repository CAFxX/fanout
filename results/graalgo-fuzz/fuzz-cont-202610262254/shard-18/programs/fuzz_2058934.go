package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 uint8, p1 uint32, p2 int16) int64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 ^ fi7)}
        _ = fi7
        _ = v8
    }
    if false {
        switch "" {
        case "a":
            p0 = p0
        case "foo bar":
            v9 := GenS1{F0: 440}
            _ = v9
        case "\n\t":
            if true {
                var v10 float32 = (1.5 / 0.1)
                p2 -= p2
                p1 -= (p1 & 1)
                _ = v10
            }
        }
    } else {
        var v11 uint16 = 91
        switch (((("a" + ("foo bar" + "\n\t")) + ("a" + ("\n\t" + "\n\t"))) + "hello") + "foo bar") {
        case "foo bar":
            p0 = (p0 % (((p0 << ((4294967295) & 63))) | 1))
        case "":
            for fi12 := 0; fi12 < 16; fi12++ {
                switch fi12 {
                case 127:
                    switch "hello" {
                    case "a":
                        fmt.Println(p2)
                    case "":
                        fmt.Println(p0)
                    case "hello":
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 int64 = 0
                            _ = fi13
                            _ = v14
                        }
                    default:
                        var v15 uint = 2
                        _ = v15
                    }
                case 4294967295:
                    fmt.Println(v11)
                case -2147483648:
                    v16 := GenS2{F0: 0}
                    _ = v16
                }
                v17 := GenS2{F0: 4294967295}
                p2 = p2
                _ = fi12
                _ = v17
            }
        }
        _ = v11
    }
    p2 = (p2 * p2)
    var v18 bool = (-2147483648 > 1)
    v19 := GenS1{F0: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return int64(p0)
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := []string{("" + ("\n\t" + "foo bar")), (("foo bar" + "hello") + (((("\n\t" + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar")) + ("" + "a")) + "")), ("" + ("a" + "hello"))}
    var v6 float64 = 3.14159
    fmt.Println(v6)
    v20 := GenS1{F0: 65535}
    switch -41 {
    case 9223372036854775807:
        var v21 int8 = -82
        _ = v21
    case 2147483647:
        var v22 uint32 = 1
        _ = v22
    }
    v6 += (float64(3.14159) + -2.25)
    v6 += (v6 * v6)
    fmt.Println(v6)
    v23 := GenS3{F0: 451, F1: true, F2: 125}
    fmt.Println(v6)
    switch 4294967295 {
    case 4294967295:
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint = 65535
            if true {
                v26 := GenS2{F0: 18446744073709551615}
                _ = v26
            }
            v6 -= 0.0
            _ = fi24
            _ = v25
        }
    case 480:
        v27 := GenS4{F0: -32, F1: -46}
        _ = v27
    default:
        for fi28 := 0; fi28 < 6; fi28++ {
            if false {
                switch ("\n\t" + "hello") {
                case "a":
                    v29 := GenS1{F0: 65535}
                    _ = v29
                default:
                    v30 := GenS2{F0: 4294967295}
                    _ = v30
                }
                v6 = ((v6 - v6) / 1.0)
            } else {
                _ = len(v4)
                v31 := GenS4{F0: 535, F1: 42}
                v32 := v1["foo bar"]
                _ = v31
                _ = v32
            }
            v33 := GenS4{F0: -326, F1: -24}
            _ = fi28
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
