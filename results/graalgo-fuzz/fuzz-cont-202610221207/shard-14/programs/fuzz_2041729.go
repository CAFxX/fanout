package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint32; F1 int64; F2 string }
type GenS7 struct { F0 uint32; F1 uint; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint8 = 218
    var v2 float64 = 1e10
    v3 := []uint32{88, 314, 481}
    var v4 float64 = 1.5
    v1 = v1
    if (32767 <= 255) {
        var v5 int = 127
        for fi6 := 0; fi6 < 5; fi6++ {
            v4 += (v4 - 0.1)
            var v7 string = "a"
            _ = fi6
            _ = v7
        }
        if false {
            if (!(v1 > (v1 % ((v1) | 1)))) {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                fmt.Println(v1, v2, v1)
            } else {
                v5 = 4294967295
            }
            v8 := GenS1{F0: -1, F1: (379 != 32767), F2: 0}
            _ = v8
        }
        _ = v5
    }
    if ((!(!true)) || (509 > 413)) {
        v9 := GenS2{F0: 862, F1: 2147483647, F2: -23}
        if len(v3) > 0 {
            v3[0] = 0
        }
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 bool = (false && false)
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS3{F0: 18446744073709551615}
            v3 = append(v3, 2147483648)
            fmt.Println(fi12)
            _ = fi12
            _ = v13
        }
        var v14 int16 = 392
        _ = v14
    }
    var v15 uint = 18446744073709551615
    if true {
        if true {
            fmt.Println(v4, v1)
            switch "\n\t" {
            case "":
                var v16 string = ""
                _ = v16
            case "foo bar":
                var v17 uint8 = v1
                _ = v17
            case "a":
                var v18 int8 = 1
                _ = v18
            }
            switch 883 {
            case 397:
                fmt.Println(v2, v4)
            case 42:
                switch "\n\t" {
                case "":
                    fmt.Println(v4, v2)
                }
            case 2147483647:
                if false {
                    var v19 string = "\n\t"
                    _ = v19
                }
            }
        }
        switch -1 {
        case 42:
            v4 = (0.1 * (v2 / v2))
        case 32767:
            switch "foo bar" {
            case "a":
                var v20 int64 = 153
                _ = v20
            case "foo bar":
                switch -9223372036854775808 {
                case 127:
                    switch 128 {
                    case 865:
                        var v21 uint32 = 42
                        _ = v21
                    }
                case 4294967295:
                    if false {
                        var v22 string = "\n\t"
                        _ = v22
                    } else {
                        switch -425 {
                        case 127:
                            v23 := GenS4{F0: 42}
                            _ = v23
                        case 32767:
                            v15 = (^(v15))
                        }
                    }
                case 32768:
                    v24 := GenS3{F0: 42}
                    _ = v24
                }
            }
        default:
            var v25 string = "a"
            _ = v25
        }
    } else {
        var v26 string = "a"
        v1 = (v1 << ((v15) & 63))
        _ = v26
    }
    v27 := GenS5{F0: 1, F1: -981}
    switch (("hello" + "a") + "\n\t") {
    case "":
        fmt.Println(v4, v15)
    case "hello":
        v3 = append(v3, 2147483648)
    case "foo bar":
        v1 -= v1
    default:
        if (1 < -9223372036854775808) {
            switch ("a" + "") {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 825
                }
            case "":
                for fi28 := 0; fi28 < 14; fi28++ {
                    if (-584 != -140) {
                        v2 -= (v4 * (0.0 + (float64((1.5 / (0.0 + 3.14159))) + (0.0 / (3.14159 - v4)))))
                        v29 := GenS6{F0: 2147483647, F1: -965, F2: (("hello" + "") + ("foo bar" + "hello"))}
                        v30 := GenS7{F0: 2147483648, F1: (v15 + (v15 % (((v15 - (v15 + 885))) | 1))), F2: (fi28 ^ (fi28 ^ (^(fi28))))}
                        _ = v29
                        _ = v30
                    }
                    _ = fi28
                }
            case "foo bar":
                fmt.Println(v15, v2, v15)
            default:
                v2 = (v2 - (v4 + v4))
            }
        }
    }
    var v31 uint64 = 18446744073709551615
    v32 := GenS8{F0: 42}
    if (42 >= 0) {
        switch -1 {
        case -1:
            var v33 uint32 = 774
            _ = v33
        case -107:
            if (18446744073709551615 < v15) {
                v34 := GenS6{F0: 4294967295, F1: 0, F2: ((("foo bar" + ("\n\t" + "hello")) + (("" + ("foo bar" + "hello")) + "foo bar")) + ("\n\t" + (("a" + "a") + "a")))}
                _ = v34
            }
        }
    }
    if (1000 < 0) {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint64 = 0
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int8 = 0
                v39 := GenS7{F0: 1, F1: (v15 * (^(v15))), F2: fi37}
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi35
            _ = v36
        }
        v4 -= (0.1 - v2)
    }
    switch 32767 {
    case 0:
        fmt.Println(v31)
    case 128:
        v1 = v1
    case 32768:
        v15 = v15
    }
    fmt.Println(v15)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "\n\t" {
        case "":
            if (0 > 42) {
                if (-9223372036854775808 == 4294967295) {
                    v40 := GenS9{F0: -83, F1: 667}
                    _ = v40
                } else {
                    var v41 uint16 = 58
                    v42 := GenS1{F0: int16(v1), F1: (v15 > (v15 & v15)), F2: 100}
                    _ = v41
                    _ = v42
                }
            } else {
                v1 -= v1
                v1 -= uint8(v15)
                if true {
                    v2 = (float64(1.5) + 3.14159)
                    var v43 uint16 = uint16(v31)
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 1; fi44++ {
                        if true {
                            v2 -= 0.0
                        }
                        var v45 uint8 = v1
                        var v46 uint16 = 0
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    v47 := GenS10{F0: (v15 * v15), F1: 255}
                    _ = v47
                }
            }
        case "a":
            for fi48 := 0; fi48 < 2; fi48++ {
                for fi49 := 0; fi49 < 17; fi49++ {
                    var v50 uint = v15
                    switch (fi48 - int(v31)) {
                    case 2147483647:
                        var v51 int = fi48
                        _ = v51
                    }
                    _ = fi49
                    _ = v50
                }
                _ = fi48
            }
        case "foo bar":
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                fmt.Println(v15)
            case "":
                switch "foo bar" {
                case "hello":
                    var v52 uint8 = v1
                    _ = v52
                case "\n\t":
                    v53 := GenS6{F0: 637, F1: -9223372036854775808, F2: ("a" + "hello")}
                    _ = v53
                default:
                    if (false || false) {
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v15, v2, v1)
                        }
                        var v54 int64 = 42
                        if (v2 <= ((1.5 / (((0.0 + v2) + v4) * 1e10)) - 3.14159)) {
                            var v55 uint32 = 2147483647
                            _ = v55
                        }
                        _ = v54
                    } else {
                        v56 := GenS3{F0: (v31 | 0)}
                        _ = v56
                    }
                }
            case "hello":
                var v57 int8 = 100
                _ = v57
            default:
                var v58 float64 = 3.14159
                _ = v58
            }
        }
    }
    var v59 bool = false
    v1 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v60 := GenS3{F0: (^(v31))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v59
    _ = v60
}
