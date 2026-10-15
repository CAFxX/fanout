package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int16; F1 bool; F2 uint32 }
type GenS6 struct { F0 int64; F1 int8 }
func genFn0(p0 uint64, p1 uint8, p2 int) int32 {
    if (-529 > 402) {
        var v4 float64 = 1.5
        var v5 int64 = 9223372036854775807
        for fi6 := 0; fi6 < 14; fi6++ {
            v5 += -9223372036854775808
            _ = fi6
        }
        _ = v4
        _ = v5
    } else {
        fmt.Println(p2, p0, p1)
        var v7 int64 = -945
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 string = (("\n\t" + ("" + "\n\t")) + "a")
        switch "hello" {
        case "a":
            fmt.Println(v9)
        case "":
            var v10 uint64 = (^(p0))
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 int8 = -79
    v2 := make(map[int]uint32)
    var v3 uint = 2
    v3 = (v3 % ((1) | 1))
    if false {
        v1 = (v1 - 42)
    }
    v1 = v1
    if true {
        if true {
            v11 := []int{-2147483648, -335}
            v12 := v2[127]
            _ = len(v2)
            _ = v11
            _ = v12
        } else {
            v13 := v2[127]
            _ = v13
        }
        var v14 uint = 2
        _ = v14
    }
    fmt.Println(v3, v3)
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(fi15)
        if ((^(v3)) > 470) {
            switch 301 {
            case -197:
                for fi16 := 0; fi16 < 17; fi16++ {
                    v3 += 2
                    v17 := GenS1{F0: 42, F1: v3}
                    _ = fi16
                    _ = v17
                }
            case -222:
                v1 = int8(fi15)
            case -365:
                var v18 uint16 = 0
                _ = v18
            default:
                v19 := []int{}
                _ = v19
            }
            switch ("foo bar" + "") {
            case "\n\t":
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 string = "foo bar"
                    v22 := GenS1{F0: 1000, F1: v3}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                switch "" {
                case "\n\t":
                    v23 := v2[(fi15 << ((v3) & 63))]
                    _ = v23
                case "a":
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := GenS1{F0: 1000, F1: (v3 ^ 42)}
                        switch (fi15 * fi15) {
                        case -243:
                            v3 = 300
                        case -963:
                            v1 += 100
                        default:
                            v26 := GenS2{F0: "", F1: "a"}
                            _ = v26
                        }
                        _ = fi24
                        _ = v25
                    }
                case "foo bar":
                    v27 := GenS1{F0: 1000, F1: 4294967295}
                    _ = v27
                default:
                    v3 -= v3
                }
            }
        } else {
            v3 -= v3
            for fi28 := 0; fi28 < 7; fi28++ {
                v3 -= 65535
                var v29 int = fi28
                _ = fi28
                _ = v29
            }
            v30 := GenS3{F0: 606}
            _ = v30
        }
        _ = fi15
    }
    var v31 uint32 = 2
    v3 = v3
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS4{F0: "", F1: 255}
        fmt.Println(v31, v3)
        _ = fi32
        _ = v33
    }
    v34 := GenS5{F0: 20, F1: ((^(v31)) >= (v31 - 2)), F2: 4294967295}
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v3, v31, fi35)
        _ = fi35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v3 -= v3
        var v37 float64 = 1e10
        _ = fi36
        _ = v37
    }
    switch ("" + ((("hello" + "a") + "") + "foo bar")) {
    case "\n\t":
        if true {
            v3 -= 42
            var v38 uint32 = 65535
            v39 := GenS2{F0: ((("a" + "a") + (((("\n\t" + ("" + "")) + "") + "") + ("\n\t" + "a"))) + ""), F1: (("\n\t" + "foo bar") + "\n\t")}
            _ = v38
            _ = v39
        } else {
            v1 += v1
            v40 := GenS5{F0: 42, F1: false, F2: v31}
            _ = v40
        }
    case "hello":
        v41 := GenS3{F0: v31}
        _ = v41
    case "foo bar":
        v42 := []int{32768, 128}
        _ = v42
    }
    v3 = 657
    for fi43 := 0; fi43 < 8; fi43++ {
        v44 := GenS6{F0: 1, F1: (v1 % ((1) | 1))}
        _ = fi43
        _ = v44
    }
    var v45 int16 = 0
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v45
}
