package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int; F1 int64; F2 uint32 }
func genFn0() bool {
    v6 := GenS1{F0: 0, F1: (1 == 142), F2: 42}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int8 = 0
        _ = fi7
        _ = v8
    }
    _ = v6
    return false
}
func genFn1() int64 {
    var v9 string = ("foo bar" + "a")
    var v10 uint32 = 1
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 float64 = (3.14159 - ((3.14159 * 1.5) + 0.0))
    var v2 int16 = 32767
    var v3 uint64 = 4294967295
    var v4 uint16 = 75
    var v5 uint8 = 128
    var v11 uint16 = (v4 & (v4 - v4))
    if (370 < 32768) {
        switch ("hello" + "") {
        case "hello":
            v11 = (v4 << ((2) & 63))
        case "a":
            v12 := GenS1{F0: 127, F1: false, F2: genFn1()}
            _ = v12
        case "foo bar":
            fmt.Println(v11)
        default:
            fmt.Println(v4)
        }
    }
    var v13 uint = 42
    fmt.Println(v13, v1, v11)
    v14 := GenS1{F0: 100, F1: (false || (42 != 2)), F2: genFn1()}
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            switch "a" {
            case "\n\t":
                v17 := GenS2{F0: -940, F1: "\n\t", F2: (v3 - (v3 << ((v13) & 63)))}
                _ = v17
            case "hello":
                var v18 int8 = 42
                _ = v18
            }
            _ = fi16
        }
        v19 := GenS1{F0: -1, F1: true, F2: 177}
        _ = fi15
        _ = v19
    }
    v13 -= 802
    switch -2147483648 {
    case -540:
        if (true && (0 >= (v5 | 131))) {
            v2 += (v2 ^ -32768)
        } else {
            if (v11 < v11) {
                var v20 float64 = 3.14159
                v21 := []int{2147483647}
                fmt.Println(v4, v5)
                _ = v20
                _ = v21
            } else {
                v22 := GenS1{F0: 100, F1: (!(-1 != -198)), F2: 1}
                var v23 uint = v13
                if false {
                    var v24 uint8 = v5
                    _ = v24
                } else {
                    v2 += 42
                    v11 += v4
                }
                _ = v22
                _ = v23
            }
            if true {
                if ("a" <= "") {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        if ((96 >= 42) || ((!((true && ((^(v3)) == 1)) && (!true))) && false)) {
                            fmt.Println(fi25, v2)
                            v26 := GenS3{F0: int(v5), F1: fi25, F2: 42}
                            var v27 float32 = (0.1 * 1.5)
                            _ = v26
                            _ = v27
                        }
                        _ = fi25
                    }
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 int = (-(fi28))
                        _ = fi28
                        _ = v29
                    }
                } else {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        fmt.Println(v13, v5, v13)
                        var v31 uint32 = 127
                        _ = fi30
                        _ = v31
                    }
                    v32 := GenS1{F0: 1, F1: (0.0 == float32(float64((3.14159 * (((-2.25 / -2.25) - 0.1) / 0.1))))), F2: -401}
                    _ = v32
                }
            } else {
                v2 += 32767
                v4 -= (v4 >> ((0) & 63))
                fmt.Println(v13, v2)
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                for fi34 := 0; fi34 < 10; fi34++ {
                    fmt.Println(fi33, v1)
                    _ = fi34
                }
                v35 := GenS4{F0: true, F1: "a"}
                _ = fi33
                _ = v35
            }
        }
    case -9223372036854775808:
        v2 -= (v2 ^ (v2 % ((-1) | 1)))
    case 9223372036854775807:
        for fi36 := 0; fi36 < 11; fi36++ {
            v37 := GenS4{F0: ((v11 > v11) || false), F1: ("" + (("foo bar" + "foo bar") + "hello"))}
            if false {
                var v38 int16 = v2
                switch fi36 {
                case -671:
                    v39 := GenS2{F0: int32(v4), F1: (("foo bar" + "foo bar") + "a"), F2: (v3 * (^(v3)))}
                    _ = v39
                case -9223372036854775808:
                    v5 = 37
                default:
                    if (v11 >= v11) {
                        var v40 int8 = 1
                        v41 := GenS2{F0: -2147483648, F1: ("" + ""), F2: v3}
                        _ = v40
                        _ = v41
                    }
                }
                _ = v38
            } else {
                for fi42 := 0; fi42 < 12; fi42++ {
                    if true {
                        var v43 int16 = v2
                        v5 = v5
                        v44 := GenS1{F0: 0, F1: (((("\n\t" + "") + ("\n\t" + "a")) + "") <= ("" + "foo bar")), F2: 42}
                        _ = v43
                        _ = v44
                    }
                    switch "a" {
                    case "a":
                        v13 = v13
                    case "\n\t":
                        fmt.Println(v3, v4)
                    case "hello":
                        fmt.Println(v1, v5, fi36)
                    default:
                        v45 := GenS1{F0: -102, F1: true, F2: 986}
                        _ = v45
                    }
                    v46 := []int{(-(fi36)), (fi36 * (-(fi36))), (fi36 >> ((4294967295) & 63)), fi42}
                    v13 += (v13 - uint(v3))
                    var v47 string = "a"
                    _ = fi42
                    _ = v46
                    _ = v47
                }
                v48 := GenS5{F0: 128, F1: 1, F2: 2147483647}
                _ = v48
            }
            v49 := []int{(fi36 / ((4294967295) | 1))}
            _ = fi36
            _ = v37
            _ = v49
        }
    }
    v13 += 4294967295
    v5 = (v5 << ((v13) & 63))
    switch "hello" {
    case "a":
        v50 := GenS2{F0: -2147483648, F1: "foo bar", F2: (v3 | 18446744073709551615)}
        _ = v50
    }
    var v51 int64 = 9223372036854775807
    var v52 int64 = 0
    v51 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
