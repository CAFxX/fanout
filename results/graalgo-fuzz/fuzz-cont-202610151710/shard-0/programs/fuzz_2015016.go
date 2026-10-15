package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int8 }
func genFn0(p0 int) uint32 {
    v7 := GenS1{F0: false, F1: (!(true && (!false))), F2: 32767}
    fmt.Println(p0)
    if false {
        if false {
            switch 2147483647 {
            case -1:
                p0 += 127
            }
            v8 := GenS1{F0: false, F1: false, F2: -32768}
            p0 = p0
            _ = v8
        }
        v9 := GenS2{F0: 1, F1: 2}
        switch (p0 - -2147483648) {
        case 639:
            v10 := GenS3{F0: -991, F1: -2147483648, F2: -897}
            _ = v10
        case 9223372036854775807:
            var v11 int64 = 0
            _ = v11
        }
        _ = v9
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 int16) string {
    if (1 == 725) {
        fmt.Println(p0, p1)
    }
    v12 := GenS4{F0: 822}
    p1 -= -780
    switch -2147483648 {
    case -422:
        v13 := GenS1{F0: p0, F1: ("a" == ("\n\t" + "")), F2: p1}
        _ = v13
    default:
        v14 := []int{4294967295}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return "a"
}

func main() {
    var v1 float64 = (0.0 + (((1.5 * 1.5) * 0.0) * float64(1e10)))
    var v2 float64 = v1
    v3 := []uint32{2147483648, 280}
    v4 := make(map[uint32]bool)
    var v5 uint8 = 253
    var v6 int64 = 42
    var v15 int16 = 0
    if (true || (!false)) {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float32 = 0.1
            var v18 uint8 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        switch -2147483648 {
        case 128:
            if false {
                var v19 uint = 4294967295
                _ = v19
            }
        case 42:
            v1 = v2
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v2, v5)
                var v21 int8 = -56
                _ = fi20
                _ = v21
            }
        }
        var v22 bool = (-2.25 >= ((1e10 - 1.5) * 0.1))
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        switch ((("a" + ("a" + "a")) + (("a" + "hello") + "\n\t")) + ("hello" + "hello")) {
        case "":
            var v24 int8 = -69
            _ = v24
        case "hello":
            switch (fi23 & fi23) {
            case 9223372036854775807:
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint32 = 131
                    v27 := GenS4{F0: 0}
                    var v28 int32 = 42
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            case -393:
                v1 = (0.1 * -2.25)
            case 32768:
                if ("foo bar" <= "foo bar") {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    var v29 float32 = 1e10
                    _ = v29
                } else {
                    if true {
                        if (-404 != 0) {
                            var v30 uint8 = (v5 + (v5 << ((982) & 63)))
                            v2 -= float64(0.0)
                            _ = v30
                        }
                        v31 := GenS4{F0: 0}
                        _ = v31
                    }
                }
            }
        case "foo bar":
            var v32 int16 = (v15 - v15)
            _ = v32
        }
        fmt.Println(v1, v6)
        _ = len(v3)
        _ = fi23
    }
    var v33 int8 = 1
    if (!true) {
        switch "\n\t" {
        case "hello":
            v33 -= 42
        default:
            for fi34 := 0; fi34 < 18; fi34++ {
                v35 := GenS2{F0: 42, F1: 1}
                for fi36 := 0; fi36 < 9; fi36++ {
                    v5 = (v5 % ((v5) | 1))
                    _ = fi36
                }
                v37 := GenS5{F0: 858, F1: 207}
                _ = fi34
                _ = v35
                _ = v37
            }
        }
        switch 498 {
        case 2147483647:
            switch -366 {
            case 42:
                var v38 float64 = v2
                _ = v38
            case 280:
                var v39 bool = true
                _ = v39
            case -116:
                var v40 uint32 = 2147483648
                _ = v40
            }
        }
    } else {
        if ((false || false) && ("hello" <= "foo bar")) {
            for fi41 := 0; fi41 < 16; fi41++ {
                fmt.Println(v15, v6)
                fmt.Println(v33, v33)
                var v42 uint64 = 23
                _ = fi41
                _ = v42
            }
            if true {
                v43 := GenS6{F0: v6, F1: ((v15 % ((v15) | 1)) < v15)}
                _ = v43
            } else {
                fmt.Println(v2, v2)
                var v44 int16 = v15
                for fi45 := 0; fi45 < 20; fi45++ {
                    if false {
                        if ((!(v33 <= v33)) && (false && (65535 < 42))) {
                            if (false && (-25 >= v33)) {
                                v46 := GenS6{F0: v6, F1: (!true)}
                                var v47 float64 = (v2 - v2)
                                _ = v46
                                _ = v47
                            }
                            v33 += (v33 | v33)
                            delete(v4, genFn0(int(v15)))
                        } else {
                            switch (fi45 % (((fi45 << ((476) & 63))) | 1)) {
                            case -360:
                                var v48 float64 = (float64(((1.5 + 3.14159) - float32(float64(1.5)))) * 1e10)
                                _ = v48
                            }
                            if (("" != (("" + ("hello" + "hello")) + ("" + "foo bar"))) || (!false)) {
                                switch (("a" + "hello") + ("\n\t" + "")) {
                                case "foo bar":
                                    v49 := GenS3{F0: (v6 + v6), F1: (^(fi45)), F2: (fi45 - (fi45 | (fi45 + (fi45 << ((42) & 63)))))}
                                    _ = v49
                                case "\n\t":
                                    var v50 uint32 = 971
                                    _ = v50
                                case "":
                                    v51 := GenS7{F0: v33}
                                    _ = v51
                                }
                                v3 = append(v3, 65535)
                                v15 = 0
                            } else {
                                fmt.Println(fi45, fi45, v6)
                                v15 = (v44 % ((v44) | 1))
                                fmt.Println(fi45, v44)
                            }
                            v1 -= v1
                        }
                        v5 += (v5 + 114)
                    } else {
                        fmt.Println(v15)
                        var v52 int8 = (v33 & (^(v33)))
                        _ = v52
                    }
                    var v53 int8 = (v33 - -65)
                    _ = fi45
                    _ = v53
                }
                _ = v44
            }
            switch "foo bar" {
            case "":
                var v54 uint8 = 179
                _ = v54
            }
        }
        fmt.Println(v33, v5)
    }
    v15 = 0
    v2 = v1
    var v55 int32 = -579
    if false {
        var v56 int = -241
        _ = v56
    }
    switch -968 {
    case 128:
        v5 += v5
    case 959:
        fmt.Println(v6, v55)
    default:
        var v57 int8 = v33
        _ = v57
    }
    for fi58 := 0; fi58 < 11; fi58++ {
        v59 := v4[genFn0((fi58 - (fi58 & (fi58 ^ fi58))))]
        for fi60 := 0; fi60 < 8; fi60++ {
            fmt.Println(v55, fi58, v5)
            var v61 uint64 = 42
            v33 = v33
            v3 = append(v3, 42)
            _ = fi60
            _ = v61
        }
        fmt.Println(fi58, v33)
        _ = fi58
        _ = v59
    }
    delete(v4, 2147483648)
    if (0 != 4294967295) {
        var v62 int32 = v55
        v2 += (v2 + float64(1e10))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v62
    } else {
        v15 = (-(v15))
        v4[4294967295] = (2 <= 1)
        v63 := GenS4{F0: 659}
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
