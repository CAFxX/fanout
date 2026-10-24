package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 string }
func genFn0() uint8 {
    var v4 int8 = 0
    v4 -= (v4 | (v4 + v4))
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 uint8 = 42
            var v7 int64 = -596
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 2147483647, F1: v4}
        if (968 >= -714) {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := []int{}
                _ = fi9
                _ = v10
            }
            v11 := []int{-9223372036854775808}
            _ = v11
        } else {
            var v12 int32 = 42
            var v13 uint32 = 65535
            _ = v12
            _ = v13
        }
        _ = v8
    } else {
        v14 := []int{-2147483648, -505, 128, 42}
        switch "\n\t" {
        case "foo bar":
            if true {
                var v15 uint8 = 207
                switch "foo bar" {
                case "\n\t":
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v17 := GenS1{F0: fi16, F1: (v4 << ((65535) & 63))}
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v15
            } else {
                v14 = append(v14, 379)
                var v18 int16 = 276
                if false {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        fmt.Println(v4, fi19, v4)
                        for fi20 := 0; fi20 < 13; fi20++ {
                            v21 := GenS1{F0: fi20, F1: v4}
                            _ = fi20
                            _ = v21
                        }
                        v4 = v4
                        _ = fi19
                    }
                }
                _ = v18
            }
        case "":
            fmt.Println(v4)
        case "a":
            v22 := GenS2{F0: (^(v4)), F1: "\n\t"}
            _ = v22
        }
        _ = v14
    }
    switch 4294967295 {
    case 888:
        var v23 float32 = 0.1
        _ = v23
    }
    _ = v4
    return 27
}

func main() {
    var v1 bool = true
    var v2 int16 = 1000
    var v3 int64 = -1
    switch -378 {
    case -503:
        if (v1 || v1) {
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 int8 = -57
                v26 := GenS2{F0: v25, F1: "foo bar"}
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v2 -= v2
                fmt.Println(v2, v3)
                fmt.Println(fi27, v1, v1)
                _ = fi27
            }
        } else {
            v2 = v2
        }
    default:
        switch -1 {
        case 128:
            var v28 int16 = v2
            _ = v28
        case -1:
            v2 = v2
        }
    }
    var v29 int = 359
    var v30 uint64 = 42
    for fi31 := 0; fi31 < 2; fi31++ {
        if false {
            for fi32 := 0; fi32 < 12; fi32++ {
                v3 += 1
                v2 += (v2 + (v2 * (v2 >> ((65535) & 63))))
                switch "hello" {
                case "a":
                    var v33 uint8 = 86
                    _ = v33
                }
                _ = fi32
            }
        }
        _ = fi31
    }
    if v1 {
        for fi34 := 0; fi34 < 3; fi34++ {
            if (0 == 255) {
                v2 = 1000
                v2 = 42
                var v35 float64 = (float64(0.0) / ((0.1 * 1.5) - 3.14159))
                _ = v35
            } else {
                for fi36 := 0; fi36 < 15; fi36++ {
                    v37 := GenS3{F0: 2, F1: 930, F2: "a"}
                    fmt.Println(v30, v30)
                    _ = fi36
                    _ = v37
                }
                var v38 uint32 = 1
                _ = v38
            }
            _ = fi34
        }
    } else {
        v30 += (^(v30))
        var v39 bool = (!(1.5 <= 1.5))
        _ = v39
    }
    v40 := GenS3{F0: 42, F1: (v30 + v30), F2: ("hello" + ("" + "foo bar"))}
    for fi41 := 0; fi41 < 1; fi41++ {
        for fi42 := 0; fi42 < 9; fi42++ {
            v43 := []int{128, (v29 ^ -503), -689, (^(fi41))}
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 10; fi44++ {
            v3 -= v3
            _ = fi44
        }
        _ = fi41
    }
    var v45 int64 = 9223372036854775807
    v45 += int64(v29)
    v29 += (v29 & (v29 >> ((4294967295) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v40
    _ = v45
}
