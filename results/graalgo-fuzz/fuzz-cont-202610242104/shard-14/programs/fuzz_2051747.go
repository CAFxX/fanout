package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    var v3 uint64 = 1
    var v4 int64 = -1
    var v5 int16 = 775
    var v6 uint8 = 112
    var v7 int64 = v4
    if (false || (true || (int(v2) <= -485))) {
        v1 -= v2
        var v8 int64 = (v4 * int64(v2))
        v9 := GenS1{F0: 2147483647, F1: 0}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: int(v6), F1: 113}
    var v11 uint16 = 255
    if false {
        v12 := []int{628}
        _ = v12
    }
    v13 := GenS1{F0: 878, F1: -1}
    fmt.Println(v11, v11, v5)
    if (!false) {
        for fi14 := 0; fi14 < 12; fi14++ {
            if (!(false && false)) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v2 = v1
                    switch "" {
                    case "a":
                        var v16 float64 = -2.25
                        _ = v16
                    case "foo bar":
                        var v17 uint64 = uint64(v11)
                        _ = v17
                    }
                    _ = fi15
                }
                switch ("" + "hello") {
                case "":
                    var v18 int = (fi14 * int(v11))
                    _ = v18
                case "foo bar":
                    if (!(v4 == v7)) {
                        v19 := []int{}
                        switch "hello" {
                        case "foo bar":
                            var v20 int32 = (v1 << ((18446744073709551615) & 63))
                            _ = v20
                        case "a":
                            v21 := GenS1{F0: (-(fi14)), F1: 127}
                            _ = v21
                        }
                        _ = len(v19)
                        _ = v19
                    }
                case "hello":
                    v22 := []int{-2147483648, fi14, 32767}
                    _ = v22
                }
                v23 := []int{}
                _ = v23
            } else {
                var v24 uint = 281
                v6 = (^(v6))
                switch ("foo bar" + (("foo bar" + "\n\t") + "foo bar")) {
                case "hello":
                    fmt.Println(v2, v24, v6)
                case "":
                    fmt.Println(v3, v11)
                default:
                    fmt.Println(v5, v6)
                }
                _ = v24
            }
            _ = fi14
        }
    }
    if (false || false) {
        v2 -= v2
    } else {
        v25 := GenS1{F0: 4294967295, F1: 1}
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 float64 = ((-2.25 + 0.0) * -2.25)
            for fi28 := 0; fi28 < 8; fi28++ {
                v27 += v27
                _ = fi28
            }
            v29 := []int{(^(fi26))}
            _ = fi26
            _ = v27
            _ = v29
        }
        _ = v25
    }
    if (-1 != (-(v5))) {
        fmt.Println(v1)
        v4 -= v7
        if (!(1.5 > 3.14159)) {
            v30 := []int{211, 0}
            _ = v30
        }
    } else {
        v5 = (v5 << ((50) & 63))
        var v31 uint = uint(v11)
        var v32 int64 = v7
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := GenS1{F0: (fi33 << ((42) & 63)), F1: -98}
        _ = fi33
        _ = v34
    }
    if true {
        if (!false) {
            fmt.Println(v6, v2)
            switch 4294967295 {
            case 4294967295:
                if (false || (((!((!false) && (!false))) && false) && true)) {
                    fmt.Println(v1)
                    v35 := []int{32768, -1}
                    _ = v35
                } else {
                    v36 := GenS2{F0: 720, F1: v6, F2: v7}
                    _ = v36
                }
            case 127:
                var v37 uint = 458
                _ = v37
            case 935:
                v6 += v6
            default:
                switch "\n\t" {
                case "a":
                    if ((65535 == 522) || false) {
                        fmt.Println(v3)
                    }
                }
            }
        } else {
            v7 = (v4 + v4)
            v38 := GenS3{F0: -1}
            v39 := GenS3{F0: 1}
            _ = v38
            _ = v39
        }
    } else {
        if (((v6 ^ v6) <= v6) || true) {
            v40 := GenS4{F0: (v5 << ((4294967295) & 63)), F1: (v3 * v3)}
            v41 := GenS4{F0: 42, F1: uint64(v11)}
            if false {
                var v42 float64 = 1.5
                var v43 string = "foo bar"
                v44 := GenS5{F0: 65535, F1: 255, F2: v5}
                _ = v42
                _ = v43
                _ = v44
            } else {
                if true {
                    var v45 uint16 = (v11 & v11)
                    v2 = (-(v2))
                    v46 := GenS6{F0: 470}
                    _ = v45
                    _ = v46
                } else {
                    v1 -= int32(v4)
                    for fi47 := 0; fi47 < 3; fi47++ {
                        fmt.Println(v1, v3, v11)
                        _ = fi47
                    }
                }
            }
            _ = v40
            _ = v41
        } else {
            switch "\n\t" {
            case "\n\t":
                var v48 string = (("foo bar" + "foo bar") + "")
                _ = v48
            case "hello":
                for fi49 := 0; fi49 < 17; fi49++ {
                    v5 = (v5 + int16(v2))
                    _ = fi49
                }
            case "":
                switch int(v3) {
                case -13:
                    var v50 int8 = -5
                    _ = v50
                }
            default:
                fmt.Println(v5, v4, v4)
            }
            var v51 uint16 = uint16(v6)
            v7 += (^(v4))
            _ = v51
        }
    }
    var v52 uint32 = 4294967295
    var v53 uint = 1
    for fi54 := 0; fi54 < 1; fi54++ {
        var v55 int16 = v5
        var v56 bool = false
        var v57 uint8 = v6
        _ = fi54
        _ = v55
        _ = v56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v7
}
