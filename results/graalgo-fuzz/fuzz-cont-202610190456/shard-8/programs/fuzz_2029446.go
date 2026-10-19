package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0() int {
    v4 := []int{42}
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: -165, F1: 970}
        v4 = append(v4, (fi5 - (fi5 ^ (fi5 >> ((42) & 63)))))
        _ = fi5
        _ = v6
    }
    if false {
        if (2 < 0) {
            if true {
                for fi7 := 0; fi7 < 3; fi7++ {
                    v8 := GenS1{F0: -1, F1: 42}
                    var v9 uint16 = 1
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                switch 128 {
                case -9223372036854775808:
                    v10 := GenS1{F0: -1, F1: 4294967295}
                    _ = v10
                case 32768:
                    switch 902 {
                    case 42:
                        var v11 string = ("" + "hello")
                        _ = v11
                    case -494:
                        var v12 float32 = ((1.5 / float32(3.14159)) * (0.1 / ((1.5 - ((1.5 * -2.25) + 1.5)) / 1e10)))
                        _ = v12
                    }
                default:
                    var v13 int8 = 127
                    _ = v13
                }
            }
        } else {
            switch "a" {
            case "hello":
                var v14 uint64 = 0
                _ = v14
            case "a":
                v4 = append(v4, 9223372036854775807)
            }
        }
    } else {
        v15 := GenS1{F0: 1000, F1: 153}
        var v16 int64 = 0
        switch (("a" + "hello") + ("a" + "a")) {
        case "foo bar":
            for fi17 := 0; fi17 < 9; fi17++ {
                if true {
                    v16 = v16
                } else {
                    var v18 uint32 = 2
                    switch (fi17 - -873) {
                    case 1:
                        for fi19 := 0; fi19 < 14; fi19++ {
                            if (42 >= -2147483648) {
                                v4 = append(v4, fi17)
                                v18 = v18
                                v18 = 2147483647
                            }
                            v18 = 2
                            fmt.Println(fi17, v16)
                            _ = fi19
                        }
                    case -2147483648:
                        for fi20 := 0; fi20 < 10; fi20++ {
                            var v21 uint64 = 18446744073709551615
                            _ = fi20
                            _ = v21
                        }
                    case 2147483647:
                        switch fi17 {
                        case 4294967295:
                            var v22 int = (fi17 & (fi17 ^ 0))
                            _ = v22
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = fi17
                        }
                    }
                    v4 = append(v4, -2147483648)
                    _ = v18
                }
                _ = fi17
            }
        case "a":
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v23 := GenS2{F0: 42, F1: 0, F2: 2}
                var v24 uint16 = 83
                _ = v23
                _ = v24
            } else {
                _ = len(v4)
                for fi25 := 0; fi25 < 19; fi25++ {
                    v16 -= v16
                    _ = fi25
                }
            }
        case "\n\t":
            switch "hello" {
            case "hello":
                fmt.Println(v16)
            default:
                switch "foo bar" {
                case "\n\t":
                    switch (("\n\t" + "a") + (("hello" + ("" + "foo bar")) + ("foo bar" + "hello"))) {
                    case "\n\t":
                        if false {
                            var v26 int64 = (v16 << ((2) & 63))
                            _ = v26
                        } else {
                            var v27 uint16 = 255
                            v27 = (v27 & 32767)
                            _ = v27
                        }
                    case "":
                        v4 = append(v4, 1)
                    case "foo bar":
                        _ = len(v4)
                    }
                default:
                    var v28 int32 = 291
                    _ = v28
                }
            }
        }
        _ = v15
        _ = v16
    }
    _ = v4
    return -994
}
func genFn1(p0 uint32, p1 int, p2 uint) uint8 {
    var v29 uint = (p2 & 42)
    v29 -= 42
    switch (p1 | -875) {
    case 42:
        if false {
            v30 := GenS1{F0: 514, F1: p0}
            _ = v30
        }
    case -1:
        switch ("hello" + "a") {
        case "hello":
            var v31 int16 = 1000
            _ = v31
        default:
            fmt.Println(p1, p0, p0)
        }
    case 4294967295:
        switch (^(p1)) {
        case 127:
            for fi32 := 0; fi32 < 14; fi32++ {
                switch (p1 ^ fi32) {
                case -1:
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 int8 = 0
                        if (!false) {
                            p0 = p0
                            fmt.Println(p0, v34)
                            fmt.Println(v34)
                        }
                        v29 += p2
                        _ = fi33
                        _ = v34
                    }
                case 127:
                    var v35 int = p1
                    _ = v35
                }
                if (false && true) {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        v37 := []int{-1}
                        _ = fi36
                        _ = v37
                    }
                    switch "\n\t" {
                    case "a":
                        p0 = p0
                    default:
                        v29 = (v29 ^ (p2 << ((255) & 63)))
                    }
                } else {
                    v29 = 1
                    v38 := []int{32768}
                    _ = v38
                }
                _ = fi32
            }
        default:
            v39 := []int{42, p1, (p1 << ((p2) & 63)), (p1 & (p1 - p1))}
            _ = v39
        }
    default:
        v40 := []int{2147483647, p1, (p1 * (-(p1))), (p1 - p1)}
        _ = v40
    }
    var v41 uint16 = 0
    for fi42 := 0; fi42 < 5; fi42++ {
        var v43 float32 = 0.0
        if true {
            v44 := GenS3{F0: int(v29)}
            p0 += (^(p0))
            for fi45 := 0; fi45 < 8; fi45++ {
                v41 -= (v41 >> (((p2 * (v29 % ((255) | 1)))) & 63))
                var v46 float32 = (3.14159 * v43)
                _ = fi45
                _ = v46
            }
            _ = v44
        }
        _ = fi42
        _ = v43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v41
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 32768
    v3 := make(map[uint32]uint32)
    v2 -= (v2 + -9223372036854775808)
    v2 -= v2
    v2 -= 482
    switch (v2 % (((v2 & v2)) | 1)) {
    case 753:
        v47 := GenS2{F0: 933, F1: 0, F2: 42}
        _ = v47
    case 32768:
        v2 = v2
    default:
        v2 = v2
    }
    v2 += v2
    var v48 string = ""
    v49 := GenS4{F0: -460}
    var v50 uint = 94
    var v51 uint16 = 0
    var v52 int64 = 0
    var v53 string = (v48 + v48)
    for fi54 := 0; fi54 < 3; fi54++ {
        for fi55 := 0; fi55 < 13; fi55++ {
            v51 -= v51
            switch fi55 {
            case -2147483648:
                v56 := GenS3{F0: fi55}
                _ = v56
            case 1:
                v57 := GenS5{F0: 0, F1: -423}
                _ = v57
            case 128:
                fmt.Println(v51)
            default:
                var v58 uint64 = 239
                _ = v58
            }
            _ = fi55
        }
        var v59 int = fi54
        _ = fi54
        _ = v59
    }
    fmt.Println(v52, v53)
    var v60 uint16 = v51
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v60
}
