package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 uint8; F1 bool; F2 int32 }
type GenS6 struct { F0 uint32; F1 bool; F2 bool }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int16, p1 string) int {
    v6 := []int{2147483647, -266, 933, 42}
    var v7 uint64 = 1
    if (p0 <= 1) {
        var v8 uint64 = v7
        v9 := GenS1{F0: -9223372036854775808, F1: 1}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            switch p1 {
            case "":
                v11 := GenS2{F0: v7, F1: 1, F2: 1}
                _ = v11
            case "a":
                switch "" {
                case "hello":
                    v6 = append(v6, (-(fi10)))
                }
            case "\n\t":
                if (p0 == (p0 ^ -1)) {
                    switch p1 {
                    case "foo bar":
                        if true {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v13 := GenS3{F0: 42, F1: 65535}
                                fmt.Println(fi10)
                                if (false && true) {
                                    var v14 uint = 65535
                                    var v15 int = fi10
                                    _ = v14
                                    _ = v15
                                }
                                _ = fi12
                                _ = v13
                            }
                            var v16 int16 = (p0 | (^(p0)))
                            _ = v16
                        } else {
                            var v17 string = p1
                            _ = v17
                        }
                    }
                    v7 = 42
                    p1 = (p1 + (("hello" + p1) + ("hello" + p1)))
                }
            default:
                var v18 int64 = 721
                _ = v18
            }
            v19 := GenS4{F0: (149 < 1), F1: 1, F2: 395}
            _ = fi10
            _ = v19
        }
        if false {
            fmt.Println(p1, p0)
        } else {
            if false {
                var v20 uint = 4294967295
                _ = v20
            } else {
                fmt.Println(p0)
                for fi21 := 0; fi21 < 8; fi21++ {
                    v22 := GenS4{F0: ((v7 >> ((2) & 63)) < 51), F1: 42, F2: 4294967295}
                    p1 = p1
                    _ = fi21
                    _ = v22
                }
            }
            var v23 uint16 = 2
            _ = v23
        }
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -101
}

func main() {
    var v1 uint8 = 87
    v2 := []int{-40, 127}
    var v3 uint16 = 462
    var v4 int8 = 100
    v5 := []string{"\n\t"}
    v4 = (v4 >> ((4294967295) & 63))
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint16 = v3
        var v26 string = ("" + "")
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint8 = v1
    fmt.Println(v3, v27, v4)
    v28 := GenS2{F0: 0, F1: 4294967295, F2: -1}
    for fi29 := 0; fi29 < 6; fi29++ {
        switch (fi29 & (fi29 & fi29)) {
        case -1:
            for fi30 := 0; fi30 < 4; fi30++ {
                switch ("a" + (("a" + "hello") + ("\n\t" + ("hello" + "hello")))) {
                case "a":
                    switch fi29 {
                    case 127:
                        if len(v2) > 0 {
                            v2[0] = (fi29 / (((fi30 * -2147483648)) | 1))
                        }
                    case 2147483647:
                        _ = len(v2)
                    default:
                        var v31 string = "hello"
                        _ = v31
                    }
                case "foo bar":
                    fmt.Println(v1)
                }
                var v32 uint8 = 255
                var v33 float64 = -2.25
                _ = fi30
                _ = v32
                _ = v33
            }
        case 2147483647:
            v34 := GenS5{F0: v1, F1: true, F2: 42}
            _ = v34
        case 127:
            v3 = (v3 % ((v3) | 1))
        default:
            fmt.Println(v1, fi29, v27)
        }
        _ = fi29
    }
    switch ("a" + "hello") {
    case "a":
        fmt.Println(v1, v3)
    default:
        if (!true) {
            var v35 int32 = 1
            for fi36 := 0; fi36 < 5; fi36++ {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v38 := GenS2{F0: 4294967295, F1: 42, F2: -990}
                    switch (fi37 | genFn0(1, ("foo bar" + "a"))) {
                    case 32768:
                        fmt.Println(v35, v4)
                    }
                    if true {
                        var v39 uint8 = v27
                        _ = v39
                    } else {
                        fmt.Println(v35, v27, v3)
                        var v40 int32 = 0
                        _ = v40
                    }
                    _ = fi37
                    _ = v38
                }
                if len(v2) > 0 {
                    v2[0] = fi36
                }
                for fi41 := 0; fi41 < 4; fi41++ {
                    v5 = append(v5, "\n\t")
                    _ = fi41
                }
                _ = fi36
            }
            if false {
                fmt.Println(v27, v1)
                if (-561 <= 9223372036854775807) {
                    v42 := GenS6{F0: 65535, F1: false, F2: false}
                    v35 -= v35
                    v4 -= (v4 * (v4 | 127))
                    _ = v42
                } else {
                    v43 := GenS4{F0: (genFn0(-595, ("\n\t" + "hello")) != genFn0(1, "hello")), F1: (v3 << ((42) & 63)), F2: 4294967295}
                    if false {
                        v44 := GenS4{F0: (3.14159 <= 0.0), F1: v3, F2: 473}
                        var v45 string = "a"
                        v4 = (v4 * (v4 & -128))
                        _ = v44
                        _ = v45
                    }
                    if true {
                        if true {
                            for fi46 := 0; fi46 < 3; fi46++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi46
                            }
                        }
                    } else {
                        v5 = append(v5, "\n\t")
                        var v47 int8 = v4
                        v5 = append(v5, (("\n\t" + "a") + "foo bar"))
                        _ = v47
                    }
                    _ = v43
                }
                v35 = (v35 << ((65535) & 63))
            } else {
                if false {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        v49 := GenS1{F0: (^(fi48)), F1: 255}
                        _ = fi48
                        _ = v49
                    }
                    if true {
                        v50 := GenS7{F0: v27}
                        if (42 > (v3 - v3)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v51 int32 = v35
                            v35 = v35
                            _ = v51
                        } else {
                            switch "" {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = int(v35)
                                }
                            }
                            switch 128 {
                            case -2147483648:
                                for fi52 := 0; fi52 < 12; fi52++ {
                                    var v53 uint16 = (^(v3))
                                    v53 += v3
                                    var v54 int16 = 42
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                        }
                        v4 -= -1
                        _ = v50
                    }
                }
                v55 := GenS4{F0: (!false), F1: 704, F2: 4294967295}
                _ = v55
            }
            _ = v35
        }
    }
    var v56 uint16 = v3
    if false {
        if true {
            var v57 uint8 = 42
            if (!false) {
                var v58 bool = false
                if (v58 && (!(v4 <= v4))) {
                    var v59 string = ("" + "hello")
                    for fi60 := 0; fi60 < 13; fi60++ {
                        v58 = true
                        var v61 uint64 = 0
                        var v62 float64 = 0.0
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    _ = v59
                }
                _ = v58
            }
            _ = v57
        }
        for fi63 := 0; fi63 < 18; fi63++ {
            v3 = v3
            _ = fi63
        }
        for fi64 := 0; fi64 < 16; fi64++ {
            var v65 int8 = 100
            v56 -= 2
            _ = fi64
            _ = v65
        }
    } else {
        var v66 uint16 = v3
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
