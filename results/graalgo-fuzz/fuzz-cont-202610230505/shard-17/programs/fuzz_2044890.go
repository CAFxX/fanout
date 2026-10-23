package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := []int{127, 127, 4294967295, 127}
    v2 := []int{9223372036854775807, -1, 2147483647}
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 int8 = 100
    var v6 uint16 = 999
    fmt.Println(v5, v4)
    var v7 int64 = -1
    switch "a" {
    case "":
        var v8 uint = 421
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v5, v4, v7)
            var v10 uint8 = 182
            _ = fi9
            _ = v10
        }
    case "foo bar":
        switch "\n\t" {
        case "hello":
            var v11 uint = 2
            _ = v11
        default:
            if len(v2) > 0 {
                v2[0] = 32767
            }
        }
    }
    switch 863 {
    case 358:
        fmt.Println(v7)
    default:
        var v12 int32 = 1
        _ = v12
    }
    if (100 > 0) {
        var v13 uint64 = 42
        for fi14 := 0; fi14 < 6; fi14++ {
            delete(v3, 665)
            _ = fi14
        }
        v15 := GenS1{F0: 995}
        _ = v13
        _ = v15
    } else {
        if (v7 <= 435) {
            v6 = (v6 >> ((18446744073709551615) & 63))
            if false {
                v4 = v5
                v16 := GenS1{F0: 1}
                v6 = (^(v6))
                _ = v16
            }
        } else {
            var v17 string = ""
            _ = v17
        }
        if (true && true) {
            switch ("a" + "foo bar") {
            case "":
                for fi18 := 0; fi18 < 15; fi18++ {
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v5 = (v4 << ((802) & 63))
                        _ = fi19
                    }
                    v20 := GenS1{F0: 1}
                    var v21 string = ""
                    _ = fi18
                    _ = v20
                    _ = v21
                }
            }
            switch 32767 {
            case -1:
                v1 = append(v1, 872)
            default:
                for fi22 := 0; fi22 < 1; fi22++ {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v4 = (-(v5))
                        _ = fi23
                    }
                    v6 = (v6 % ((65535) | 1))
                    _ = fi22
                }
            }
            if (true || false) {
                var v24 uint32 = 12
                _ = v24
            }
        } else {
            var v25 float64 = 0.0
            v6 = (v6 | 3)
            if len(v2) > 0 {
                v2[0] = 128
            }
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            _ = len(v3)
            for fi27 := 0; fi27 < 13; fi27++ {
                if false {
                    v28 := GenS1{F0: 68}
                    var v29 bool = false
                    v4 += -45
                    _ = v28
                    _ = v29
                }
                var v30 uint16 = v6
                _ = fi27
                _ = v30
            }
            _ = fi26
        }
    }
    switch "foo bar" {
    case "hello":
        var v31 string = "foo bar"
        _ = v31
    case "":
        var v32 uint64 = 1
        _ = v32
    default:
        fmt.Println(v4, v6, v4)
    }
    var v33 string = ("a" + ("\n\t" + ""))
    if (648 > 412) {
        if (4294967295 > 1) {
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := GenS2{F0: 134, F1: (^(v6)), F2: false}
                var v36 uint8 = 2
                fmt.Println(v7, v36)
                _ = fi34
                _ = v35
                _ = v36
            }
            if (1 == v4) {
                switch "a" {
                case "foo bar":
                    var v37 uint32 = 2147483647
                    _ = v37
                case "\n\t":
                    v2 = append(v2, -1)
                default:
                    v6 -= 1
                }
                var v38 string = (v33 + v33)
                _ = v38
            } else {
                v39 := GenS2{F0: 42, F1: (v6 & (^(v6))), F2: (2 >= 288)}
                var v40 uint8 = 93
                var v41 uint64 = 2
                _ = v39
                _ = v40
                _ = v41
            }
        } else {
            if true {
                v33 = ((v33 + (("foo bar" + "a") + "a")) + "")
                fmt.Println(v33, v4)
                var v42 int = 0
                _ = v42
            } else {
                var v43 uint = 627
                _ = v43
            }
            fmt.Println(v5)
            v33 = ("hello" + (((("foo bar" + v33) + ((v33 + "a") + (("hello" + v33) + "hello"))) + v33) + (v33 + v33)))
        }
    } else {
        v5 = v5
    }
    switch ("" + (v33 + v33)) {
    case "foo bar":
        for fi44 := 0; fi44 < 17; fi44++ {
            var v45 uint = 255
            delete(v3, 459)
            _ = fi44
            _ = v45
        }
    case "hello":
        v46 := GenS1{F0: 2147483648}
        _ = v46
    case "":
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if true {
                var v47 uint16 = (v6 & (v6 | v6))
                var v48 int8 = v4
                _ = v47
                _ = v48
            }
            var v49 int32 = 42
            if true {
                v4 -= (v4 / ((v4) | 1))
                var v50 int64 = v7
                v51 := v3[2147483647]
                _ = v50
                _ = v51
            }
            _ = v49
        }
    default:
        var v52 int8 = -1
        _ = v52
    }
    var v53 int = 4294967295
    for fi54 := 0; fi54 < 13; fi54++ {
        v55 := GenS3{F0: 598, F1: 2}
        _ = fi54
        _ = v55
    }
    v56 := GenS2{F0: 1, F1: v6, F2: false}
    if (true || (("" != v33) || (((1e10 * 3.14159) >= ((0.0 * float32(0.1)) + float32(3.14159))) || true))) {
        var v57 int8 = v4
        v53 = v53
        _ = v57
    } else {
        v53 = (v53 + -2147483648)
        if (411 == -32768) {
            fmt.Println(v33, v7, v7)
            v33 = v33
        } else {
            for fi58 := 0; fi58 < 5; fi58++ {
                for fi59 := 0; fi59 < 8; fi59++ {
                    fmt.Println(v33, v5)
                    if len(v1) > 0 {
                        v1[0] = v53
                    }
                    _ = fi59
                }
                v60 := v3[424]
                _ = fi58
                _ = v60
            }
            var v61 float32 = -2.25
            _ = v61
        }
        switch (v53 << ((536) & 63)) {
        case 9223372036854775807:
            var v62 int16 = 32767
            _ = v62
        default:
            var v63 int32 = 42
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 6; fi64++ {
        v4 = (v4 >> ((252) & 63))
        switch (^(v53)) {
        case 902:
            v4 += (v4 + v5)
        case 127:
            delete(v3, 2147483647)
        case 29:
            var v65 int16 = 939
            _ = v65
        }
        _ = fi64
    }
    var v66 float64 = (1e10 - ((1e10 + float64((3.14159 * ((1.5 / 1.0) * -2.25)))) / 1.0))
    if true {
        delete(v3, 1)
        var v67 uint8 = 0
        v68 := GenS3{F0: 255, F1: 42}
        _ = v67
        _ = v68
    }
    v69 := GenS4{F0: 219}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v53
    _ = v56
    _ = v6
    _ = v66
    _ = v69
    _ = v7
}
