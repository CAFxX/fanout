package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int) string {
    p0 -= p0
    v7 := GenS1{F0: 18446744073709551615, F1: false, F2: 42}
    p0 += p0
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = fi8
        p0 -= p0
        var v9 uint64 = 0
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return "\n\t"
}
func genFn1(p0 int8, p1 string) uint8 {
    fmt.Println(p0, p1)
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    return 128
}

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    v3 := []uint32{4294967295}
    var v4 bool = false
    var v5 float64 = -2.25
    var v6 uint = 18446744073709551615
    for fi11 := 0; fi11 < 9; fi11++ {
        if ((fi11 * fi11) < (fi11 | fi11)) {
            for fi12 := 0; fi12 < 14; fi12++ {
                if v4 {
                    v5 -= v5
                    switch "foo bar" {
                    case "hello":
                        var v13 uint64 = 119
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = len(v3)
                        _ = v13
                    case "a":
                        var v14 int64 = -1
                        _ = v14
                    }
                    v15 := GenS1{F0: 0, F1: ("foo bar" >= "hello"), F2: -1}
                    _ = v15
                }
                var v16 uint64 = 0
                var v17 int16 = 0
                _ = fi12
                _ = v16
                _ = v17
            }
        }
        v2 = "\n\t"
        _ = fi11
    }
    var v18 string = (v2 + ("a" + (((v2 + (("a" + v1) + "\n\t")) + (v2 + v1)) + "\n\t")))
    switch v18 {
    case "\n\t":
        if (!(("hello" + ("" + v18)) >= v2)) {
            fmt.Println(v4, v2)
            var v19 uint8 = 1
            _ = v19
        }
    case "":
        if v4 {
            var v20 string = ("" + "")
            if (true && (255 == v6)) {
                fmt.Println(v4, v20, v5)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v20
        } else {
            fmt.Println(v5, v4, v6)
            var v21 uint = (v6 / (((v6 - (v6 >> ((v6) & 63)))) | 1))
            _ = v21
        }
    case "foo bar":
        var v22 uint16 = 65535
        _ = v22
    }
    if ("hello" >= "hello") {
        v23 := GenS2{F0: 4294967295, F1: 1}
        _ = v23
    }
    switch v1 {
    case "hello":
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 uint = v6
            _ = fi24
            _ = v25
        }
    }
    v18 = v2
    var v26 float64 = v5
    v2 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
