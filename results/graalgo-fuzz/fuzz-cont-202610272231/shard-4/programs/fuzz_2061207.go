package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }

func main() {
    v1 := []string{("" + "a"), "", ("hello" + "")}
    var v2 uint8 = 53
    var v3 uint = 2
    v4 := make(map[string]int)
    var v5 bool = false
    var v6 int64 = 0
    if v5 {
        switch "\n\t" {
        case "":
            if (-2.25 > 0.0) {
                var v7 uint64 = 978
                _ = v7
            }
        case "\n\t":
            v3 = 65535
        }
    } else {
        if (1.5 == (0.0 - (-2.25 + (3.14159 - -2.25)))) {
            v2 += (v2 / ((114) | 1))
        }
        v6 = v6
        if v5 {
            v2 += v2
            v6 += (^(v6))
        }
    }
    v5 = v5
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint64 = 597
        var v10 uint = (^(v3))
        for fi11 := 0; fi11 < 7; fi11++ {
            v1 = append(v1, ("\n\t" + "\n\t"))
            var v12 uint16 = 61
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v13 := GenS1{F0: 2147483647, F1: -548, F2: 395}
    delete(v4, "foo bar")
    switch 32768 {
    case 1:
        for fi14 := 0; fi14 < 5; fi14++ {
            v6 += v6
            v15 := GenS1{F0: 42, F1: 42, F2: 9223372036854775807}
            _ = fi14
            _ = v15
        }
    default:
        v16 := GenS2{F0: 42, F1: (^(v2)), F2: 577}
        _ = v16
    }
    switch "" {
    case "":
        v6 -= v6
    case "foo bar":
        var v17 int32 = 0
        _ = v17
    case "hello":
        if (0 <= 32767) {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 uint8 = (v2 | v2)
                _ = fi18
                _ = v19
            }
            fmt.Println(v3)
            switch ("a" + "a") {
            case "a":
                delete(v4, "a")
            case "hello":
                fmt.Println(v5)
            default:
                switch "\n\t" {
                case "foo bar":
                    v20 := GenS2{F0: 42, F1: 200, F2: 65535}
                    _ = v20
                case "hello":
                    if true {
                        v5 = (!true)
                        if (false && false) {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 string = (("foo bar" + "\n\t") + ("a" + "hello"))
                                if (2 >= 42) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v22 = "\n\t"
                                } else {
                                    v6 = v6
                                }
                                _ = fi21
                                _ = v22
                            }
                            fmt.Println(v5, v2)
                        }
                    }
                case "\n\t":
                    v3 = v3
                default:
                    var v23 string = "foo bar"
                    _ = v23
                }
            }
        } else {
            for fi24 := 0; fi24 < 1; fi24++ {
                v2 = 42
                if v5 {
                    fmt.Println(v5, v5, v6)
                }
                v25 := GenS3{F0: (v6 << ((v3) & 63)), F1: "", F2: 127}
                _ = fi24
                _ = v25
            }
            var v26 float64 = (0.0 + 1.5)
            _ = v26
        }
    default:
        var v27 int32 = 0
        _ = v27
    }
    var v28 int = int(v3)
    v2 = (v2 ^ 234)
    if false {
        v3 -= v3
        var v29 int32 = -603
        fmt.Println(v3, v3)
        _ = v29
    }
    fmt.Println(v2, v3)
    var v30 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
