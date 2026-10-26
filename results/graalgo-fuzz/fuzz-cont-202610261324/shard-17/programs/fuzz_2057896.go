package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint }

func main() {
    var v1 uint8 = 127
    var v2 int = 128
    var v3 bool = (false || false)
    var v4 uint32 = 669
    v5 := GenS1{F0: (^(v4)), F1: -32768}
    var v6 uint16 = 31
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{fi7, fi7, v2}
        var v9 int8 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v3, v3)
    v2 -= v2
    v10 := GenS2{F0: 2147483647, F1: 2147483647, F2: v4}
    var v11 float64 = 1e10
    if (v6 <= v6) {
        v12 := GenS2{F0: 806, F1: 1, F2: (^(v4))}
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint8 = (v1 / (((v1 + (^(v1)))) | 1))
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            switch "hello" {
            case "a":
                var v17 string = "foo bar"
                _ = v17
            case "foo bar":
                var v18 string = "foo bar"
                _ = v18
            default:
                v19 := GenS1{F0: (v4 | v4), F1: 0}
                _ = v19
            }
            var v20 uint = 0
            v1 = (^(v1))
            _ = fi16
            _ = v20
        }
        _ = v12
    }
    v11 += (3.14159 * v11)
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v6, v4, fi21)
        var v22 bool = false
        _ = fi21
        _ = v22
    }
    if true {
        if ((3 == 127) || true) {
            v6 += 282
            v11 -= v11
            var v23 int16 = 658
            _ = v23
        } else {
            v24 := []int{}
            for fi25 := 0; fi25 < 15; fi25++ {
                if v3 {
                    fmt.Println(v3, v2)
                    var v26 uint8 = v1
                    _ = v26
                } else {
                    if (998 >= (-(fi25))) {
                        fmt.Println(fi25, fi25)
                    }
                    switch "foo bar" {
                    case "a":
                        var v27 float32 = 0.0
                        _ = v27
                    case "hello":
                        fmt.Println(v2, v3, v1)
                    case "foo bar":
                        v6 = (v6 << ((65535) & 63))
                    default:
                        fmt.Println(v1)
                    }
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 uint = 1
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        v30 := GenS3{F0: -9223372036854775808, F1: (("foo bar" + "hello") > "hello")}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                if v3 {
                    var v31 float64 = v11
                    _ = v31
                }
                _ = fi25
            }
            switch -9223372036854775808 {
            case 0:
                v32 := GenS1{F0: v4, F1: 42}
                _ = v32
            default:
                var v33 uint16 = v6
                _ = v33
            }
            _ = v24
        }
        fmt.Println(v3, v1, v1)
        switch "a" {
        case "\n\t":
            v34 := []int{(v2 + 127), (v2 + v2), (v2 | (v2 & (v2 - v2)))}
            _ = v34
        case "foo bar":
            fmt.Println(v1, v1, v4)
        }
    }
    if (42 > 0) {
        for fi35 := 0; fi35 < 19; fi35++ {
            if (!(!v3)) {
                v1 += (^(v1))
            }
            v36 := GenS4{F0: v3}
            var v37 uint64 = 42
            _ = fi35
            _ = v36
            _ = v37
        }
    } else {
        v38 := GenS4{F0: ((v6 << ((826) & 63)) != v6)}
        if (!(-9223372036854775808 > 989)) {
            v4 += v4
        } else {
            v39 := []int{v2, v2, (v2 * (v2 * 9223372036854775807)), -9223372036854775808}
            _ = v39
        }
        _ = v38
    }
    switch (v2 & (v2 ^ v2)) {
    case 1:
        v4 += (v4 ^ (^(v4)))
    }
    var v40 uint8 = 2
    switch (v2 << ((412) & 63)) {
    case 127:
        v41 := []int{(v2 % (((v2 - (v2 & v2))) | 1)), v2, (v2 >> ((uint(v4)) & 63)), 32767}
        _ = v41
    case -842:
        v42 := GenS3{F0: 1, F1: (!((((v4 | (^(v4))) >= 0) && (!v3)) && v3))}
        _ = v42
    default:
        v43 := []int{}
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := []int{fi44, v2, 128}
        if (!v3) {
            fmt.Println(v4, v3)
        } else {
            if false {
                if true {
                    switch fi44 {
                    case -930:
                        if (4294967295 == 4294967295) {
                            v46 := GenS5{F0: uint16(v4)}
                            _ = v46
                        }
                    default:
                        v40 = 255
                    }
                } else {
                    if v3 {
                        var v47 int64 = 0
                        v48 := GenS6{F0: v3, F1: 18446744073709551615}
                        if len(v45) > 0 {
                            v45[0] = 9223372036854775807
                        }
                        _ = v47
                        _ = v48
                    } else {
                        v49 := GenS4{F0: (v6 > (^(v6)))}
                        _ = v49
                    }
                    v2 = v2
                }
                _ = len(v45)
                v50 := GenS4{F0: ("hello" > (((("" + "hello") + "\n\t") + ("foo bar" + (("a" + "") + ""))) + "a"))}
                _ = v50
            }
            var v51 int16 = -1
            _ = v51
        }
        v3 = v3
        _ = fi44
        _ = v45
    }
    v2 = (v2 << ((18446744073709551615) & 63))
    var v52 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v6
}
