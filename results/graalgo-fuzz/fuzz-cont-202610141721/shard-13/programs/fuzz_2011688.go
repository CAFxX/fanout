package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1
    v3 := make(map[uint32]bool)
    fmt.Println(v2)
    v2 = (v2 % (((v2 % (((^(v2))) | 1))) | 1))
    v4 := GenS1{F0: 42, F1: 1, F2: 1}
    fmt.Println(v2)
    v5 := GenS2{F0: 346, F1: 610, F2: 1}
    var v6 int = 488
    for fi7 := 0; fi7 < 2; fi7++ {
        if (!true) {
            for fi8 := 0; fi8 < 9; fi8++ {
                v9 := GenS1{F0: -32, F1: 431, F2: 1}
                v6 = (fi8 * fi8)
                if true {
                    var v10 bool = true
                    var v11 int64 = -9223372036854775808
                    _ = v10
                    _ = v11
                } else {
                    v6 = fi8
                    v6 += 32768
                }
                _ = fi8
                _ = v9
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := GenS2{F0: 0, F1: 2147483647, F2: 211}
                v14 := []int{(fi7 / (((fi7 + 128)) | 1)), fi7, fi12, (-(v6))}
                switch "" {
                case "\n\t":
                    v6 -= v6
                case "":
                    if true {
                        var v15 string = "foo bar"
                        switch ("\n\t" + (("\n\t" + ("a" + v15)) + v15)) {
                        case "foo bar":
                            var v16 uint32 = 782
                            _ = v16
                        }
                        fmt.Println(v15, v2)
                        _ = v15
                    } else {
                        v17 := GenS1{F0: 127, F1: (^(fi7)), F2: 900}
                        var v18 bool = true
                        _ = v17
                        _ = v18
                    }
                case "foo bar":
                    if (0 == 65535) {
                        v2 += 32767
                        var v19 uint8 = 141
                        _ = v19
                    }
                default:
                    _ = len(v1)
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            fmt.Println(fi7)
            v2 = 524
        }
        _ = fi7
    }
    v20 := GenS1{F0: -79, F1: (v6 >> ((42) & 63)), F2: 9223372036854775807}
    v21 := []int{(-(v6)), v6}
    fmt.Println(v2)
    switch -9223372036854775808 {
    case -528:
        v6 -= v6
    case 4294967295:
        fmt.Println(v6)
    case 127:
        switch 32767 {
        case 4294967295:
            switch v6 {
            case 341:
                v6 = v6
            }
        case 686:
            delete(v3, 2)
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                v2 += v2
                var v23 int = 0
                v24 := GenS2{F0: 1, F1: 2, F2: 18446744073709551615}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    var v25 string = (("foo bar" + ("foo bar" + ("" + "a"))) + "hello")
    if ((v25 + v25) <= v25) {
        var v26 uint = 1
        if true {
            for fi27 := 0; fi27 < 4; fi27++ {
                v6 -= fi27
                switch (("foo bar" + "hello") + v25) {
                case "foo bar":
                    v21 = append(v21, v6)
                default:
                    _ = len(v21)
                }
                var v28 uint32 = 232
                _ = fi27
                _ = v28
            }
        }
        switch v6 {
        case 523:
            v25 = (("foo bar" + v25) + "hello")
        }
        _ = v26
    }
    var v29 int16 = v2
    fmt.Println(v2)
    _ = len(v21)
    v30 := GenS2{F0: 9223372036854775807, F1: 2, F2: 0}
    v25 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
