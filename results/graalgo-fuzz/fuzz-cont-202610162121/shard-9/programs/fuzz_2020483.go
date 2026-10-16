package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 string, p1 uint, p2 uint8) uint {
    if false {
        for fi4 := 0; fi4 < 13; fi4++ {
            p0 = p0
            switch "hello" {
            case "a":
                v5 := GenS1{F0: 127, F1: -9223372036854775808, F2: -612}
                _ = v5
            default:
                var v6 uint8 = p2
                _ = v6
            }
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: -1, F1: int64(p1), F2: 42}
                _ = fi7
                _ = v8
            }
            _ = fi4
        }
        if true {
            var v9 bool = false
            p1 = (p1 + 18446744073709551615)
            _ = v9
        } else {
            var v10 int8 = -1
            _ = v10
        }
    } else {
        p2 -= (p2 * 2)
        var v11 int8 = 91
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 87, F1: 690, F2: 885}
        var v14 int = 127
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: 1, F1: false, F2: (^(p1))}
        switch (("foo bar" + ("a" + p0)) + p0) {
        case "\n\t":
            fmt.Println(p1, p1, p2)
        case "foo bar":
            fmt.Println(p2, p0, p0)
        }
        _ = fi15
        _ = v16
    }
    p0 = (p0 + p0)
    switch 1 {
    case 128:
        switch (p0 + p0) {
        case "":
            v17 := GenS2{F0: 255, F1: (42 < 0), F2: 4294967295}
            _ = v17
        case "a":
            if true {
                v18 := GenS1{F0: 98, F1: -9223372036854775808, F2: 292}
                v19 := GenS1{F0: 16, F1: -9223372036854775808, F2: -848}
                _ = v18
                _ = v19
            } else {
                p2 = 42
                for fi20 := 0; fi20 < 14; fi20++ {
                    if false {
                        if false {
                            var v21 uint = p1
                            v21 -= v21
                            _ = v21
                        }
                        switch "hello" {
                        case "foo bar":
                            for fi22 := 0; fi22 < 16; fi22++ {
                                var v23 int32 = -512
                                p1 += p1
                                fmt.Println(p0)
                                _ = fi22
                                _ = v23
                            }
                        case "hello":
                            for fi24 := 0; fi24 < 9; fi24++ {
                                v25 := []int{}
                                v26 := GenS2{F0: p2, F1: false, F2: p1}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        case "a":
                            switch 32767 {
                            case -2147483648:
                                v27 := GenS3{F0: 395, F1: 1}
                                _ = v27
                            }
                        }
                        fmt.Println(p2, p0, fi20)
                    }
                    v28 := GenS4{F0: p2, F1: -9223372036854775808, F2: 960}
                    _ = fi20
                    _ = v28
                }
            }
        }
    case 2147483647:
        var v29 uint = p1
        _ = v29
    case 127:
        p0 = (p0 + ("foo bar" + p0))
    default:
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 2147483648
    var v3 int32 = -293
    var v30 int8 = -66
    switch "\n\t" {
    case "foo bar":
        v1 = ("\n\t" + "\n\t")
    case "hello":
        for fi31 := 0; fi31 < 9; fi31++ {
            v30 -= (v30 * 42)
            v30 = 100
            for fi32 := 0; fi32 < 1; fi32++ {
                v1 = "\n\t"
                v33 := []int{(^(fi31))}
                v30 = int8(v3)
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    }
    v30 = v30
    switch -2147483648 {
    case -634:
        if false {
            v34 := GenS4{F0: 42, F1: -9223372036854775808, F2: 2}
            v3 -= 1
            _ = v34
        }
    }
    v30 += (v30 ^ (v30 * v30))
    v2 += v2
    v35 := []int{}
    fmt.Println(v3, v30, v1)
    var v36 int8 = v30
    if len(v35) > 0 {
        v35[0] = -308
    }
    v30 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
}
