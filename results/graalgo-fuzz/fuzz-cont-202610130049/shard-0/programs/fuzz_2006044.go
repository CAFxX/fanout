package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 uint }
type GenS4 struct { F0 int32; F1 bool; F2 int }
type GenS5 struct { F0 int16; F1 int16; F2 uint16 }
type GenS6 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := []uint32{2, 42}
    v4 := make(map[int]bool)
    v5 := []string{"a"}
    v6 := make(map[uint32]uint32)
    var v7 bool = ("a" <= "hello")
    if (773 > 520) {
        fmt.Println(v7)
    } else {
        switch "" {
        case "a":
            switch -9223372036854775808 {
            case 128:
                v7 = v7
            }
        case "\n\t":
            for fi8 := 0; fi8 < 5; fi8++ {
                v7 = (197 >= 65535)
                fmt.Println(fi8)
                _ = fi8
            }
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 989, F1: (-(fi9))}
            _ = fi9
            _ = v10
        }
    }
    var v11 int32 = 42
    fmt.Println(v11, v11)
    if (v7 || (v7 && (1e10 != (1.5 - (-2.25 * float64(((1.5 - 1e10) + 1.5))))))) {
        fmt.Println(v7)
    } else {
        v12 := GenS2{F0: (v11 * -2147483648), F1: -1, F2: (v11 * (v11 - v11))}
        var v13 int32 = v11
        v7 = (0 == 2)
        _ = v12
        _ = v13
    }
    fmt.Println(v11, v7)
    if (v7 || false) {
        v14 := GenS3{F0: 0, F1: 744, F2: 65535}
        v15 := GenS3{F0: 16, F1: 312, F2: 42}
        delete(v4, 32767)
        _ = v14
        _ = v15
    }
    if false {
        v7 = true
        var v16 uint8 = 17
        _ = v16
    }
    var v17 float32 = (0.1 / 1e10)
    v7 = false
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 int16 = 1000
        var v20 string = ""
        fmt.Println(v11)
        _ = fi18
        _ = v19
        _ = v20
    }
    if (32767 > -1) {
        switch 2147483647 {
        case 128:
            switch "a" {
            case "":
                var v21 uint64 = 2
                _ = v21
            }
        case -13:
            var v22 uint8 = 128
            _ = v22
        default:
            var v23 uint64 = 4294967295
            _ = v23
        }
        _ = len(v3)
        switch ("a" + "foo bar") {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = "hello"
            }
        default:
            var v24 uint32 = 0
            _ = v24
        }
    }
    if (2 != 2) {
        v25 := GenS4{F0: v11, F1: v7, F2: -1}
        for fi26 := 0; fi26 < 5; fi26++ {
            v11 -= 2147483647
            v17 += v17
            _ = fi26
        }
        _ = v25
    } else {
        v27 := GenS2{F0: v11, F1: -108, F2: (v11 % ((v11) | 1))}
        switch "hello" {
        case "hello":
            var v28 int32 = v11
            _ = v28
        case "\n\t":
            if (false && (2147483648 <= 654)) {
                var v29 int16 = 0
                if (!(!(!true))) {
                    v30 := GenS2{F0: v11, F1: 127, F2: v11}
                    var v31 float32 = (-2.25 / 1.5)
                    var v32 int64 = int64(v11)
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v11, v11)
                    fmt.Println(v7)
                    fmt.Println(v7, v29)
                }
                var v33 float64 = 1.5
                _ = v29
                _ = v33
            } else {
                switch 2147483647 {
                case 32768:
                    v11 = v11
                case -1:
                    var v34 bool = ((!((v17 != 0.0) || (((!v7) || v7) || (!(("" + "foo bar") > ("foo bar" + "\n\t")))))) || (float64(v17) != (1e10 + (1e10 / (1e10 - (((((0.0 - 3.14159) * 3.14159) + (float64(v17) + (0.0 * (0.1 / -2.25)))) / -2.25) + 3.14159))))))
                    _ = v34
                default:
                    var v35 int64 = -1
                    _ = v35
                }
                var v36 uint = 730
                _ = v36
            }
        }
        _ = v27
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            if v7 {
                switch fi37 {
                case 159:
                    v2[fi38] = (v17 < (v17 - v17))
                }
            }
            _ = fi38
        }
        _ = fi37
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        v17 = 3.14159
        switch 256 {
        case 42:
            if (v7 && v7) {
                var v40 uint64 = 168
                v41 := GenS4{F0: (^(v11)), F1: (-516 < 253), F2: (fi39 % ((fi39) | 1))}
                _ = v40
                _ = v41
            }
        }
        if ("a" < "\n\t") {
            switch fi39 {
            case 757:
                switch -1 {
                case 9223372036854775807:
                    if len(v3) > 0 {
                        v3[0] = 853
                    }
                default:
                    v6[395] = 0
                }
            default:
                var v42 uint = 682
                _ = v42
            }
            v43 := GenS5{F0: 986, F1: 1, F2: 2}
            _ = v43
        } else {
            if len(v3) > 0 {
                v3[0] = 42
            }
            v11 -= (v11 + v11)
        }
        _ = fi39
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        _ = len(v6)
        _ = fi44
    }
    var v45 uint16 = 65535
    if true {
        var v46 uint64 = 42
        _ = v46
    } else {
        if true {
            var v47 int16 = 807
            if (v7 || v7) {
                _ = len(v3)
                v48 := GenS6{F0: v7, F1: 406}
                v49 := GenS4{F0: (v11 - -1), F1: (v7 || (false || true)), F2: -873}
                _ = v48
                _ = v49
            }
            _ = v47
        }
        v11 = (v11 ^ (v11 >> ((42) & 63)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
