package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint; F2 int }
type GenS8 struct { F0 int16; F1 uint32; F2 string }
type GenS9 struct { F0 uint8; F1 uint8 }
func genFn0(p0 string, p1 uint64, p2 uint64) uint32 {
    var v7 uint = 18446744073709551615
    if false {
        v8 := []int{733, 32768}
        fmt.Println(p1, v7)
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: fi9, F1: fi9}
            _ = fi9
            _ = v10
        }
        if true {
            v11 := GenS1{F0: 42, F1: 0}
            var v12 uint32 = 2147483647
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 706}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 bool, p2 int32) uint32 {
    switch "" {
    case "a":
        if (-1 <= 977) {
            v14 := GenS3{F0: 601, F1: 42}
            v15 := GenS4{F0: 42, F1: p0, F2: (p0 << ((4294967295) & 63))}
            _ = v14
            _ = v15
        } else {
            var v16 int = -642
            if false {
                v17 := []int{v16}
                _ = v17
            } else {
                v18 := []int{}
                _ = v18
            }
            _ = v16
        }
    case "foo bar":
        var v19 uint32 = 302
        _ = v19
    case "":
        if p1 {
            if ("hello" == ("hello" + ("foo bar" + "a"))) {
                p1 = (!true)
            }
        }
    }
    if true {
        for fi20 := 0; fi20 < 15; fi20++ {
            switch "a" {
            case "foo bar":
                p2 -= (p2 * p2)
            }
            v21 := []int{32767, (fi20 - -2147483648), (fi20 % ((fi20) | 1)), 9223372036854775807}
            _ = fi20
            _ = v21
        }
    } else {
        p1 = (!p1)
    }
    var v22 float32 = (1.5 * 0.1)
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 float64 = 1.5
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS3{F0: -62, F1: -535}
            v27 := []int{(^(fi23)), -484}
            v27 = append(v27, (fi23 >> ((794) & 63)))
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 2147483648
}

func main() {
    var v1 uint32 = 42
    var v2 string = ("\n\t" + ("hello" + "\n\t"))
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    v5 := make(map[int]uint32)
    var v6 int = 1
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            v2 = (v2 + v2)
            _ = fi29
        }
        _ = fi28
    }
    if ((!true) || (127 < 42)) {
        if false {
            var v30 int8 = 1
            switch (v6 - 244) {
            case 128:
                switch (v6 & v6) {
                case -9223372036854775808:
                    v30 = v30
                }
            }
            _ = v30
        } else {
            delete(v3, (((v2 + v2) + "\n\t") + v2))
        }
        v4 += v4
        fmt.Println(v1)
    }
    if ((!true) && false) {
        var v31 uint16 = 255
        v1 = (v1 / (((^(v1))) | 1))
        v32 := GenS5{F0: (v1 >> ((42) & 63))}
        _ = v31
        _ = v32
    } else {
        var v33 int = v6
        _ = v33
    }
    if (!(false || (true || (!true)))) {
        var v34 int32 = -1
        v35 := GenS6{F0: -128}
        _ = v34
        _ = v35
    } else {
        fmt.Println(v4, v2)
        v36 := GenS2{F0: 65535}
        _ = v36
    }
    if (-32768 >= 0) {
        var v37 uint64 = (^(v4))
        _ = v37
    } else {
        v38 := v3[(v2 + (v2 + (v2 + v2)))]
        _ = v38
    }
    switch v2 {
    case "":
        var v39 int = (v6 * (-(v6)))
        _ = v39
    default:
        var v40 bool = ((32767 != 42) || false)
        _ = v40
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        switch "\n\t" {
        case "hello":
            v42 := GenS3{F0: 42, F1: 562}
            _ = v42
        case "a":
            fmt.Println(v1, fi41, v4)
        case "\n\t":
            switch v6 {
            case -9223372036854775808:
                for fi43 := 0; fi43 < 18; fi43++ {
                    if (42 > 1) {
                        v44 := []int{}
                        _ = v44
                    }
                    switch (-(fi41)) {
                    case -2147483648:
                        _ = len(v3)
                    default:
                        v6 = fi43
                    }
                    _ = fi43
                }
            case 4294967295:
                for fi45 := 0; fi45 < 16; fi45++ {
                    var v46 int16 = -323
                    _ = fi45
                    _ = v46
                }
            }
        default:
            switch ((("\n\t" + v2) + "foo bar") + (v2 + ("\n\t" + v2))) {
            case "hello":
                if (((1e10 >= 0.0) || false) || false) {
                    switch (v6 << ((18446744073709551615) & 63)) {
                    case 746:
                        var v47 uint = 255
                        _ = v47
                    case 128:
                        fmt.Println(v6, v1)
                    case -131:
                        fmt.Println(v2, v1)
                    }
                    v48 := GenS7{F0: (v2 + v2), F1: 42, F2: v6}
                    var v49 string = "hello"
                    _ = v48
                    _ = v49
                } else {
                    switch v6 {
                    case 4294967295:
                        var v50 bool = (v4 < (v4 | (^(v4))))
                        _ = v50
                    case -9223372036854775808:
                        v1 = genFn1(255, true, 1)
                    default:
                        v4 -= 367
                    }
                    if ((!false) || false) {
                        fmt.Println(v6)
                        for fi51 := 0; fi51 < 18; fi51++ {
                            fmt.Println(v4, v4)
                            if false {
                                if false {
                                    v52 := GenS7{F0: "\n\t", F1: 42, F2: -212}
                                    _ = v52
                                } else {
                                    v53 := GenS3{F0: -32768, F1: 42}
                                    v4 = v4
                                    _ = v53
                                }
                                var v54 int64 = 908
                                if ((1000 > -32768) && true) {
                                    switch (v2 + ((("hello" + (v2 + "foo bar")) + v2) + "")) {
                                    case "hello":
                                        fmt.Println(v6, v2, fi51)
                                    }
                                } else {
                                    v6 += 32768
                                }
                                _ = v54
                            } else {
                                v55 := GenS4{F0: 65535, F1: 1, F2: 65535}
                                fmt.Println(v4)
                                _ = v55
                            }
                            var v56 int = (fi51 & 967)
                            _ = fi51
                            _ = v56
                        }
                    }
                    v57 := []int{(fi41 | -2147483648), (v6 + v6), 0}
                    _ = v57
                }
            default:
                v58 := GenS8{F0: 168, F1: (v1 & (v1 & (v1 - (v1 * v1)))), F2: v2}
                _ = v58
            }
        }
        v59 := GenS8{F0: 1, F1: (v1 + v1), F2: ""}
        _ = fi41
        _ = v59
    }
    switch v6 {
    case 0:
        fmt.Println(v6)
    default:
        switch v6 {
        case 231:
            v1 -= 1
        case 123:
            v2 = "hello"
        default:
            v60 := GenS9{F0: 2, F1: 1}
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
