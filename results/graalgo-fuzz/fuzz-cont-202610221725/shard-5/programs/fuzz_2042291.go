package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
func genFn0(p0 uint, p1 int64, p2 uint64) int8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int32 = 212
        var v6 bool = false
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -1, F1: p0}
    v8 := GenS2{F0: 73, F1: "foo bar", F2: -372}
    fmt.Println(p1, p2)
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18
}
func genFn1(p0 bool, p1 int64, p2 int8) string {
    p0 = (!(0 < 756))
    p0 = (((!(4294967295 != 540)) && false) && p0)
    _ = p0
    _ = p1
    _ = p2
    return ("a" + "hello")
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    var v3 int8 = -128
    if false {
        _ = len(v2)
        delete(v1, ("hello" + ""))
    } else {
        v3 += genFn0(0, 472, 2)
        v9 := GenS1{F0: 32768, F1: 338}
        _ = v9
    }
    switch -2147483648 {
    case 2147483647:
        switch 41 {
        case 9223372036854775807:
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint32 = 275
                switch fi10 {
                case -259:
                    fmt.Println(v11, v11)
                case 9223372036854775807:
                    v1[""] = false
                default:
                    if (("foo bar" + "a") < "a") {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            var v13 int = -599
                            if false {
                                fmt.Println(v3, fi12)
                                v11 = 2147483647
                            } else {
                                var v14 float64 = 1.5
                                _ = v14
                            }
                            fmt.Println(v11)
                            v1["\n\t"] = true
                            _ = fi12
                            _ = v13
                        }
                        v11 = 65535
                        delete(v1, "\n\t")
                    }
                }
                v15 := GenS3{F0: 121, F1: 4294967295, F2: 42}
                _ = fi10
                _ = v11
                _ = v15
            }
        case 269:
            v3 = v3
        default:
            v16 := GenS1{F0: 2147483647, F1: 2}
            _ = v16
        }
    case 60:
        v3 = -100
    case 1:
        if true {
            if true {
                v3 -= v3
                fmt.Println(v3)
            }
            switch 32768 {
            case 4294967295:
                switch 32768 {
                case 32767:
                    fmt.Println(v3)
                case 1:
                    v3 -= v3
                }
            case -905:
                fmt.Println(v3)
            default:
                var v17 uint16 = 1
                _ = v17
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := v1["\n\t"]
                fmt.Println(fi18, v3)
                _ = fi18
                _ = v19
            }
        } else {
            v20 := v1[""]
            v21 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v20
            _ = v21
        }
    }
    if (1 > 1) {
        switch 9223372036854775807 {
        case 128:
            fmt.Println(v3)
        default:
            v3 = genFn0(42, 0, 2)
        }
    }
    fmt.Println(v3)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint16 = 171
        switch "\n\t" {
        case "hello":
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS3{F0: 1, F1: 2, F2: -705}
                var v26 uint64 = 2
                _ = fi24
                _ = v25
                _ = v26
            }
        case "a":
            if ("" < "foo bar") {
                switch (fi22 % (((fi22 | (fi22 * -221))) | 1)) {
                case -1:
                    v23 = (v23 / ((v23) | 1))
                default:
                    v27 := GenS1{F0: (fi22 << ((65535) & 63)), F1: 255}
                    _ = v27
                }
            } else {
                switch "" {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v28 uint64 = 90
                    _ = v28
                }
            }
        }
        _ = fi22
        _ = v23
    }
    v29 := GenS1{F0: -824, F1: 396}
    v3 -= (v3 ^ (v3 - genFn0(517, 42, 2)))
    var v30 int = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
