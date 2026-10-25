package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint16 }
func genFn0(p0 int) int32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 1}
        _ = fi6
        _ = v7
    }
    if ((true || false) || true) {
        p0 = p0
        switch p0 {
        case 0:
            for fi8 := 0; fi8 < 6; fi8++ {
                if false {
                    v9 := GenS2{F0: ("foo bar" + "\n\t"), F1: 0}
                    _ = v9
                }
                p0 = -9223372036854775808
                p0 += -424
                _ = fi8
            }
        default:
            var v10 uint = 1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    _ = p0
    return -629
}
func genFn1(p0 uint32, p1 uint16) string {
    if false {
        if (p0 == p0) {
            var v13 uint32 = p0
            v14 := GenS2{F0: "", F1: -1}
            v15 := GenS3{F0: -1, F1: 100, F2: -538}
            _ = v13
            _ = v14
            _ = v15
        } else {
            var v16 int32 = -1
            v17 := []int{}
            fmt.Println(v16)
            _ = v16
            _ = v17
        }
        if true {
            switch 127 {
            case -247:
                if false {
                    v18 := GenS2{F0: ("a" + "a"), F1: 167}
                    p1 -= 42
                    if true {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            var v20 uint8 = 255
                            v21 := GenS4{F0: 65535, F1: 42, F2: p0}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        p0 -= 65535
                        switch (("hello" + "foo bar") + "foo bar") {
                        case "":
                            var v22 string = "foo bar"
                            _ = v22
                        case "a":
                            for fi23 := 0; fi23 < 5; fi23++ {
                                v24 := GenS5{F0: (fi23 - int(p0))}
                                _ = fi23
                                _ = v24
                            }
                        default:
                            p0 = (p0 | p0)
                        }
                    }
                    _ = v18
                } else {
                    v25 := []int{1, -9223372036854775808, 9223372036854775807}
                    p0 = (^(p0))
                    p0 = p0
                    _ = v25
                }
            case 994:
                var v26 int16 = 32767
                _ = v26
            }
        }
        var v27 string = ("a" + "\n\t")
        _ = v27
    } else {
        v28 := []int{42, -903}
        if false {
            for fi29 := 0; fi29 < 20; fi29++ {
                switch "" {
                case "":
                    if len(v28) > 0 {
                        v28[0] = (fi29 - (fi29 - fi29))
                    }
                case "\n\t":
                    var v30 int = fi29
                    _ = v30
                }
                v31 := GenS3{F0: -1, F1: -128, F2: 42}
                v32 := GenS1{F0: (^(p0))}
                _ = fi29
                _ = v31
                _ = v32
            }
            p0 -= 65535
        }
        if (128 < uint8(p0)) {
            var v33 uint = 18446744073709551615
            for fi34 := 0; fi34 < 5; fi34++ {
                if true {
                    p1 -= p1
                }
                p0 = p0
                if (42 != 18446744073709551615) {
                    var v35 uint16 = p1
                    _ = v35
                }
                _ = fi34
            }
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            _ = v33
        }
        _ = v28
    }
    p1 = 255
    var v36 int = -2147483648
    _ = p0
    _ = p1
    _ = v36
    return "foo bar"
}

func main() {
    var v1 int32 = 191
    v2 := make(map[int]uint32)
    var v3 int32 = (-(v1))
    v4 := make(map[uint32]uint32)
    var v5 int32 = 2147483647
    for fi37 := 0; fi37 < 12; fi37++ {
        if (0 < 42) {
            var v38 uint = 1
            fmt.Println(v5)
            fmt.Println(v38, v1)
            _ = v38
        } else {
            var v39 string = ((("" + "foo bar") + "hello") + "")
            switch fi37 {
            case 4294967295:
                for fi40 := 0; fi40 < 12; fi40++ {
                    switch (^(fi40)) {
                    case 4294967295:
                        switch v39 {
                        case "a":
                            _ = len(v4)
                        case "hello":
                            var v41 uint = 255
                            var v42 int64 = 9223372036854775807
                            v43 := GenS4{F0: (v41 | v41), F1: 2, F2: 2}
                            _ = v41
                            _ = v42
                            _ = v43
                        case "foo bar":
                            switch (fi37 - fi37) {
                            case 127:
                                var v44 int64 = 42
                                _ = v44
                            case 32768:
                                var v45 uint64 = 42
                                _ = v45
                            case 1:
                                var v46 uint = 343
                                _ = v46
                            }
                        default:
                            delete(v2, fi40)
                        }
                    default:
                        var v47 int16 = 0
                        _ = v47
                    }
                    v48 := []int{fi40, fi40, fi37, fi37}
                    _ = fi40
                    _ = v48
                }
            case 685:
                v49 := GenS3{F0: 32767, F1: -1, F2: -9223372036854775808}
                _ = v49
            }
            v3 = (v3 ^ v3)
            _ = v39
        }
        v3 += (^(v5))
        _ = fi37
    }
    fmt.Println(v3, v3, v1)
    fmt.Println(v1, v5)
    if false {
        var v50 int32 = 550
        _ = v50
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        switch "\n\t" {
        case "foo bar":
            v52 := GenS2{F0: "foo bar", F1: 42}
            _ = v52
        case "hello":
            v53 := v4[65535]
            _ = v53
        case "\n\t":
            delete(v2, fi51)
        default:
            fmt.Println(v5, v3)
        }
        switch fi51 {
        case 32767:
            var v54 uint64 = 42
            _ = v54
        }
        _ = fi51
    }
    switch 127 {
    case 4294967295:
        var v55 uint16 = 85
        _ = v55
    case -9223372036854775808:
        v56 := GenS4{F0: 993, F1: 42, F2: 0}
        _ = v56
    }
    v5 = (v3 << ((4294967295) & 63))
    switch "foo bar" {
    case "":
        if (!(42 < 1)) {
            var v57 int = 42
            v5 = v1
            fmt.Println(v5, v57, v57)
            _ = v57
        } else {
            var v58 string = ""
            v59 := GenS6{F0: 426}
            v58 = (v58 + v58)
            _ = v58
            _ = v59
        }
    case "\n\t":
        switch ("hello" + (((("hello" + ("" + "hello")) + ("hello" + "foo bar")) + (((("a" + "a") + "foo bar") + "a") + ("hello" + "\n\t"))) + "foo bar")) {
        case "foo bar":
            v5 = (v1 / ((v5) | 1))
        case "\n\t":
            for fi60 := 0; fi60 < 20; fi60++ {
                var v61 uint16 = 78
                _ = fi60
                _ = v61
            }
        }
    default:
        var v62 int = 42
        _ = v62
    }
    var v63 int = 42
    for fi64 := 0; fi64 < 4; fi64++ {
        for fi65 := 0; fi65 < 11; fi65++ {
            if (18446744073709551615 >= 0) {
                var v66 bool = (946 > 72)
                _ = v66
            }
            v3 = (v3 ^ v5)
            _ = fi65
        }
        var v67 int8 = 1
        _ = fi64
        _ = v67
    }
    if (false || true) {
        for fi68 := 0; fi68 < 18; fi68++ {
            delete(v2, (v63 * (fi68 & (v63 << ((28) & 63)))))
            _ = fi68
        }
        fmt.Println(v3, v3)
    } else {
        v63 += (v63 / (((v63 & v63)) | 1))
    }
    v69 := GenS7{F0: 222, F1: 42, F2: 841}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v69
}
