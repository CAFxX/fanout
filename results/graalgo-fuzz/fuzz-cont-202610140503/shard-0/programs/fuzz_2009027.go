package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool, p1 int64) int8 {
    v6 := []int{128, 128, -120, 0}
    var v7 uint64 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 bool = true
    var v2 int8 = 0
    v3 := []string{("a" + "\n\t"), "foo bar", "", ("" + "foo bar")}
    v4 := make(map[string]bool)
    var v5 int = 20
    switch "a" {
    case "hello":
        var v8 uint64 = 42
        _ = v8
    case "\n\t":
        v4[""] = (!v1)
    case "foo bar":
        var v9 uint = 18446744073709551615
        _ = v9
    }
    v10 := GenS1{F0: 32768, F1: v2, F2: 42}
    v1 = (0 < 42)
    v3 = append(v3, "")
    v11 := GenS2{F0: 128}
    for fi12 := 0; fi12 < 17; fi12++ {
        if (v2 > v2) {
            var v13 int8 = 42
            var v14 int8 = (v13 % (((v13 | (v13 >> ((550) & 63)))) | 1))
            _ = v13
            _ = v14
        } else {
            v2 -= -128
        }
        _ = fi12
    }
    if v1 {
        var v15 int = -743
        _ = len(v3)
        _ = v15
    } else {
        v1 = (v2 > (v2 << ((313) & 63)))
        if v1 {
            var v16 float64 = 1.5
            var v17 uint8 = 127
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v5)
    var v18 uint8 = 128
    if (2147483647 == 65535) {
        var v19 uint16 = 2
        _ = v19
    }
    if (!false) {
        for fi20 := 0; fi20 < 19; fi20++ {
            switch (v5 + (fi20 * fi20)) {
            case 32767:
                switch 1 {
                case -43:
                    if v1 {
                        switch "foo bar" {
                        case "a":
                            for fi21 := 0; fi21 < 1; fi21++ {
                                var v22 int8 = v2
                                v23 := GenS1{F0: 2147483647, F1: v2, F2: 107}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        case "foo bar":
                            var v24 uint32 = 42
                            _ = v24
                        }
                        v18 = v18
                    }
                case 42:
                    v25 := v4[(("a" + ("foo bar" + "")) + "a")]
                    _ = v25
                }
            case 127:
                v2 -= (v2 << ((1) & 63))
            }
            switch "\n\t" {
            case "hello":
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 int = fi20
                    v28 := GenS3{F0: 18446744073709551615, F1: 65535}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            var v29 uint32 = 2147483647
            _ = fi20
            _ = v29
        }
    }
    v30 := GenS4{F0: 32767}
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
