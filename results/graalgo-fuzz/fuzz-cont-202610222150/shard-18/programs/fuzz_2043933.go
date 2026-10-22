package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int; F1 int64; F2 int64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint8) bool {
    v4 := []int{-9223372036854775808, -697, -9223372036854775808}
    v5 := GenS1{F0: 0, F1: 65535, F2: 4294967295}
    _ = p0
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int, p1 int8, p2 uint) int8 {
    p1 -= 127
    v6 := GenS1{F0: p2, F1: 2, F2: 42}
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "foo bar"
    var v3 int16 = 649
    if (true && (false && ((v3 << ((255) & 63)) >= -516))) {
        v3 = v3
    } else {
        var v7 int8 = 42
        _ = v7
    }
    switch -316 {
    case 9223372036854775807:
        v8 := GenS1{F0: 42, F1: 2, F2: 1}
        _ = v8
    case 42:
        var v9 uint32 = 2147483648
        _ = v9
    case 1:
        if (2 == 757) {
            fmt.Println(v3, v3)
            v3 = (v3 + v3)
            v10 := []int{2147483647, 0, 2147483647, -2147483648}
            _ = v10
        }
    default:
        var v11 int16 = 0
        _ = v11
    }
    switch ("hello" + v2) {
    case "a":
        v3 = (v3 & (v3 / ((423) | 1)))
    case "":
        switch "hello" {
        case "":
            var v12 int8 = -112
            _ = v12
        case "a":
            var v13 bool = false
            _ = v13
        default:
            fmt.Println(v3, v3)
        }
    }
    switch 0 {
    case 402:
        for fi14 := 0; fi14 < 19; fi14++ {
            v3 = 42
            _ = fi14
        }
    case 2147483647:
        switch -1 {
        case 127:
            v15 := []int{}
            _ = v15
        case 128:
            v16 := GenS1{F0: 18446744073709551615, F1: 2, F2: 212}
            _ = v16
        case 42:
            delete(v1, 2)
        }
    default:
        v17 := GenS2{F0: 2, F1: 42, F2: 65535}
        _ = v17
    }
    v18 := []int{2147483647, -2147483648, 383, 4294967295}
    var v19 int32 = -1
    v19 = v19
    if (65535 <= 65535) {
        v20 := v1[2147483648]
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int64 = 1
            v22 = v22
            if (true && v20) {
                if (int16(v22) <= 622) {
                    v20 = (9223372036854775807 >= fi21)
                }
                if (v20 || true) {
                    v23 := GenS1{F0: 749, F1: 367, F2: 2}
                    var v24 uint8 = 127
                    _ = v23
                    _ = v24
                }
            }
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v25 uint = 2
    if false {
        switch v2 {
        case "foo bar":
            if ((1 <= 42) || (1.5 > 0.0)) {
                var v26 uint8 = 42
                v3 = v3
                _ = v26
            }
        case "\n\t":
            for fi27 := 0; fi27 < 9; fi27++ {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := GenS2{F0: 0, F1: 166, F2: 876}
                    var v30 uint = 4294967295
                    if false {
                        fmt.Println(fi27, fi28)
                    } else {
                        v31 := v1[126]
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v2 = "\n\t"
                _ = fi27
            }
        default:
            v19 = v19
        }
    } else {
        switch 127 {
        case 655:
            if false {
                fmt.Println(v19, v25, v3)
                var v32 uint32 = 552
                if true {
                    if (!false) {
                        v19 = v19
                    } else {
                        v33 := GenS3{F0: 478, F1: 76, F2: -994}
                        fmt.Println(v19, v25)
                        _ = v33
                    }
                    if false {
                        var v34 int8 = -66
                        var v35 uint32 = uint32(v25)
                        _ = v34
                        _ = v35
                    } else {
                        v19 = 731
                    }
                } else {
                    v2 = "\n\t"
                    var v36 bool = (false && true)
                    _ = v36
                }
                _ = v32
            }
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            if (!true) {
                if false {
                    v19 = v19
                    var v38 int = 175
                    for fi39 := 0; fi39 < 5; fi39++ {
                        v2 = v2
                        var v40 int16 = (v3 ^ 1000)
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                }
                var v41 uint8 = 1
                _ = v41
            } else {
                for fi42 := 0; fi42 < 17; fi42++ {
                    fmt.Println(fi37, v25)
                    _ = fi42
                }
                if (true || false) {
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v1[2147483647] = (0 > v3)
                        var v44 int = fi43
                        v45 := GenS4{F0: 4294967295, F1: v2}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                }
                switch ((("a" + ((v2 + v2) + (v2 + (v2 + (v2 + (((v2 + v2) + v2) + "a")))))) + "") + "\n\t") {
                case "a":
                    switch v2 {
                    case "\n\t":
                        _ = len(v1)
                    }
                case "":
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                default:
                    if true {
                        _ = len(v18)
                    }
                }
            }
            v46 := GenS5{F0: -1}
            _ = fi37
            _ = v46
        }
        var v47 float64 = -2.25
        _ = v47
    }
    v2 = ("\n\t" + (v2 + "\n\t"))
    switch "foo bar" {
    case "hello":
        var v48 uint64 = 4294967295
        _ = v48
    case "":
        var v49 int8 = 6
        _ = v49
    case "foo bar":
        for fi50 := 0; fi50 < 15; fi50++ {
            v3 += v3
            v3 = -10
            _ = fi50
        }
    default:
        fmt.Println(v2, v2, v2)
    }
    var v51 float32 = 0.0
    var v52 uint8 = 127
    v3 = -1
    v25 = 255
    delete(v1, uint32(v52))
    v51 = float32(-2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v51
    _ = v52
}
