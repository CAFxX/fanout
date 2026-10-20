package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int8; F1 uint16 }
func genFn0() int16 {
    var v6 int = 128
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 = (-(fi7))
        for fi8 := 0; fi8 < 9; fi8++ {
            v6 = (fi8 | fi8)
            v6 += fi8
            v6 = v6
            _ = fi8
        }
        _ = fi7
    }
    v9 := []int{}
    if ((!(((3.14159 * (0.0 / 3.14159)) + float32(-2.25)) <= 0.1)) || false) {
        fmt.Println(v6)
        fmt.Println(v6)
        v6 += 2147483647
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if true {
                if true {
                    if (0 >= 2147483647) {
                        if (-2147483648 <= 42) {
                            v11 := GenS1{F0: -9223372036854775808, F1: 2147483647}
                            _ = v11
                        } else {
                            v6 -= fi10
                            v9 = append(v9, -2147483648)
                        }
                        fmt.Println(v6)
                        if false {
                            var v12 bool = false
                            _ = v12
                        }
                    }
                    fmt.Println(fi10)
                }
                var v13 bool = (1.5 > 0.1)
                _ = v13
            } else {
                fmt.Println(v6, fi10)
                var v14 bool = (0 <= 42)
                _ = len(v9)
                _ = v14
            }
            _ = fi10
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if true {
            v6 = v6
            var v15 int32 = 1
            _ = v15
        }
    }
    v6 += (v6 << ((9) & 63))
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 int64, p1 int8, p2 uint16) bool {
    var v16 string = ("hello" + "hello")
    v17 := []int{-1}
    if (2147483647 <= 2) {
        if ((p0 >> ((18446744073709551615) & 63)) > (p0 | p0)) {
            var v18 float64 = 0.0
            v17 = append(v17, -1)
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            p1 -= p1
            var v20 string = v16
            v21 := GenS2{F0: 628, F1: p2}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if true {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS2{F0: 550, F1: p2}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            switch -9223372036854775808 {
            case 4294967295:
                switch -486 {
                case 1:
                    if len(v17) > 0 {
                        v17[0] = 127
                    }
                case 148:
                    if ((0.0 > float64(3.14159)) && (p1 > -128)) {
                        v25 := GenS3{F0: 0}
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 18; fi26++ {
                            v27 := GenS3{F0: 786}
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            }
            var v28 uint16 = p2
            switch v16 {
            case "foo bar":
                for fi29 := 0; fi29 < 17; fi29++ {
                    fmt.Println(v16, fi24)
                    _ = fi29
                }
            case "\n\t":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            _ = fi24
            _ = v28
        }
    }
    var v30 float64 = (1e10 * (1.5 * (1e10 - 0.0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v30
    return (false && false)
}

func main() {
    v1 := []int{32767, 98, 608}
    var v2 bool = true
    var v3 int16 = 42
    var v4 uint = 0
    var v5 int16 = (v3 - (v3 | v3))
    if v2 {
        v1 = append(v1, 2147483647)
        v31 := GenS3{F0: v5}
        _ = v31
    } else {
        switch "\n\t" {
        case "a":
            v2 = false
        case "\n\t":
            for fi32 := 0; fi32 < 19; fi32++ {
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 int = (fi33 | (fi32 << (((v4 + v4)) & 63)))
                    _ = fi33
                    _ = v34
                }
                var v35 int = (^(fi32))
                _ = fi32
                _ = v35
            }
        case "hello":
            var v36 string = ("" + "\n\t")
            _ = v36
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi37 := 0; fi37 < 5; fi37++ {
            fmt.Println(v3, v5)
            v3 -= 42
            fmt.Println(fi37, v4)
            _ = fi37
        }
    case "":
        if v2 {
            fmt.Println(v2, v4, v3)
            for fi38 := 0; fi38 < 13; fi38++ {
                v39 := GenS4{F0: v3, F1: "foo bar"}
                v40 := GenS2{F0: (v4 ^ 255), F1: 231}
                _ = fi38
                _ = v39
                _ = v40
            }
            for fi41 := 0; fi41 < 19; fi41++ {
                var v42 uint = v4
                _ = fi41
                _ = v42
            }
        } else {
            v43 := GenS3{F0: 1000}
            _ = v43
        }
    default:
        var v44 float32 = 0.1
        _ = v44
    }
    switch -9223372036854775808 {
    case 42:
        var v45 int64 = 1
        _ = v45
    }
    v46 := GenS5{F0: 127, F1: 27}
    var v47 int = -658
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 uint64 = 823
        _ = fi48
        _ = v49
    }
    var v50 float64 = (1e10 + 1e10)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
