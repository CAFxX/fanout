package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 uint16, p1 uint8) uint32 {
    p0 -= p0
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int64 = -480
        fmt.Println(p0, fi7)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 += 1
        _ = fi9
    }
    _ = p0
    _ = p1
    return 606
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]string)
    v3 := []int{32767}
    var v4 int = -315
    var v5 uint16 = 840
    var v6 float64 = 1.5
    v2[405] = ("hello" + "\n\t")
    if false {
        switch -2147483648 {
        case 85:
            var v10 uint64 = 42
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 2; fi11++ {
                switch v4 {
                case -811:
                    for fi12 := 0; fi12 < 12; fi12++ {
                        fmt.Println(v5)
                        var v13 uint = 141
                        v14 := GenS1{F0: ("" + "\n\t"), F1: 42}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 2147483647:
                    v5 += (v5 + (v5 % ((v5) | 1)))
                case -350:
                    for fi15 := 0; fi15 < 2; fi15++ {
                        fmt.Println(v6, fi15)
                        _ = fi15
                    }
                default:
                    v16 := GenS1{F0: "a", F1: -128}
                    _ = v16
                }
                switch "\n\t" {
                case "a":
                    v5 = v5
                case "foo bar":
                    var v17 int = (v4 * fi11)
                    _ = v17
                }
                fmt.Println(v6)
                _ = fi11
            }
        case 4294967295:
            if true {
                v6 = v6
                if true {
                    switch (v4 << ((0) & 63)) {
                    case 1:
                        if true {
                            var v18 bool = (42 == -123)
                            _ = v18
                        } else {
                            _ = len(v2)
                            fmt.Println(v5)
                        }
                    case 9223372036854775807:
                        v19 := GenS1{F0: "", F1: -128}
                        _ = v19
                    }
                    v20 := GenS1{F0: "", F1: 0}
                    fmt.Println(v1)
                    _ = v20
                }
                _ = len(v3)
            } else {
                switch ("\n\t" + "hello") {
                case "foo bar":
                    var v21 uint = 255
                    _ = v21
                case "":
                    v22 := v2[genFn0(v5, 42)]
                    _ = v22
                }
            }
        default:
            if (620 != 2) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS1{F0: "a", F1: 127}
                    _ = fi23
                    _ = v24
                }
            }
        }
        fmt.Println(v1, v4)
    } else {
        fmt.Println(v6)
    }
    v4 = 2147483647
    var v25 int = v4
    v26 := GenS1{F0: "hello", F1: 42}
    var v27 float32 = 0.0
    switch ("foo bar" + "\n\t") {
    case "":
        if (0 == 42) {
            v4 = 0
        } else {
            v5 = v5
            v3 = append(v3, v25)
        }
    case "a":
        fmt.Println(v4, v25)
    }
    var v28 float64 = float64((v27 / 1.5))
    var v29 uint = 4294967295
    v4 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
