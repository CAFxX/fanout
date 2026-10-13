package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    var v4 uint64 = 18446744073709551615
    if false {
        var v5 uint32 = 65535
        _ = v5
    } else {
        if (42 == 250) {
            var v6 int16 = (v2 * (v2 ^ (v2 | v2)))
            _ = len(v3)
            _ = v6
        }
    }
    if (2147483647 != 1) {
        switch "\n\t" {
        case "a":
            v7 := GenS1{F0: (263 <= 911), F1: true, F2: 410}
            _ = v7
        case "hello":
            v8 := GenS2{F0: v4, F1: (v4 & 18446744073709551615), F2: false}
            _ = v8
        case "\n\t":
            if true {
                switch 9223372036854775807 {
                case 32768:
                    var v9 float32 = 1e10
                    _ = v9
                case -2147483648:
                    v4 = v4
                default:
                    v4 += (^(v4))
                }
            } else {
                fmt.Println(v2, v4)
            }
        }
    }
    var v10 uint8 = 174
    if true {
        var v11 bool = (!true)
        _ = v11
    }
    delete(v3, -1)
    var v12 uint64 = 4294967295
    switch "hello" {
    case "a":
        v2 = v2
    case "":
        var v13 uint = 0
        _ = v13
    case "foo bar":
        var v14 uint32 = 4294967295
        _ = v14
    default:
        if (737 != 1) {
            var v15 uint16 = 42
            v16 := GenS1{F0: true, F1: true, F2: 42}
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v4, v2, v2)
    var v17 int16 = 1000
    v4 -= (v4 / ((718) | 1))
    delete(v3, -2147483648)
    if (true || true) {
        for fi18 := 0; fi18 < 12; fi18++ {
            switch (fi18 << ((65535) & 63)) {
            case 827:
                if false {
                    fmt.Println(v4, fi18, v2)
                    if ((!(577 == 255)) || (0 != 1)) {
                        v4 += v4
                        v19 := GenS1{F0: (!("a" != "foo bar")), F1: true, F2: 42}
                        _ = v19
                    }
                } else {
                    switch ((("\n\t" + "foo bar") + "hello") + "foo bar") {
                    case "a":
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v12 = (v12 | v4)
                            v21 := v3[2147483647]
                            _ = fi20
                            _ = v21
                        }
                    case "\n\t":
                        v22 := GenS2{F0: 1, F1: v12, F2: (-517 > 2147483647)}
                        _ = v22
                    case "foo bar":
                        switch "a" {
                        case "":
                            v10 -= (v10 * (v10 >> ((567) & 63)))
                        default:
                            v23 := GenS1{F0: (!true), F1: false, F2: 1}
                            _ = v23
                        }
                    }
                    var v24 uint16 = 2
                    v25 := GenS3{F0: -9223372036854775808}
                    _ = v24
                    _ = v25
                }
            case 1:
                switch fi18 {
                case -9223372036854775808:
                    fmt.Println(v4, fi18, v17)
                default:
                    var v26 uint32 = 2147483648
                    _ = v26
                }
            case -2147483648:
                var v27 bool = false
                _ = v27
            }
            v2 -= 1000
            _ = fi18
        }
    }
    if (!true) {
        fmt.Println(v2)
        var v28 float64 = (0.0 - -2.25)
        _ = v28
    }
    var v29 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
