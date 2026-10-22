package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int8; F2 bool }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint8, p2 uint64) int16 {
    v7 := GenS1{F0: 610, F1: 42, F2: 2147483647}
    v8 := GenS2{F0: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1() int16 {
    v9 := []int{2147483647, 296, 32768}
    var v10 int16 = 1000
    if false {
        for fi11 := 0; fi11 < 4; fi11++ {
            v10 -= genFn0(0, 128, 212)
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int32 = -159
                _ = len(v9)
                var v14 float64 = 3.14159
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    } else {
        v9 = append(v9, 9223372036854775807)
        v15 := GenS3{F0: 42, F1: 0, F2: 4294967295}
        if false {
            if len(v9) > 0 {
                v9[0] = 42
            }
            if len(v9) > 0 {
                v9[0] = -2147483648
            }
        } else {
            fmt.Println(v10)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if (0 <= -1) {
                var v16 uint16 = 332
                var v17 string = "a"
                _ = v16
                _ = v17
            }
        }
        _ = v15
    }
    v10 = v10
    v10 = v10
    _ = v10
    _ = v9
    return 1000
}

func main() {
    v1 := []uint32{}
    var v2 bool = (!true)
    v3 := []string{"a", "a", "\n\t"}
    var v4 int16 = 985
    var v5 uint32 = 2147483648
    var v6 bool = (342 == 1)
    var v18 int64 = 911
    var v19 int64 = v18
    var v20 uint8 = 128
    fmt.Println(v19, v5)
    var v21 float64 = (-2.25 - 1.5)
    if (1 == (v20 - (^(v20)))) {
        if v2 {
            v18 = (v19 + 0)
            if true {
                v5 = (v5 | (v5 & (v5 | (v5 | 42))))
            } else {
                if ((v18 + 9223372036854775807) == v19) {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v1)
                    case "":
                        v22 := GenS4{F0: 127}
                        _ = v22
                    case "a":
                        if v6 {
                            if false {
                                switch (("a" + "") + ("a" + ((("foo bar" + "\n\t") + "\n\t") + "foo bar"))) {
                                case "a":
                                    switch "a" {
                                    case "":
                                        var v23 float64 = v21
                                        _ = v23
                                    default:
                                        for fi24 := 0; fi24 < 6; fi24++ {
                                            v25 := GenS1{F0: 1, F1: v18, F2: -918}
                                            var v26 bool = true
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                case "hello":
                                    switch 9223372036854775807 {
                                    case 128:
                                        v27 := GenS3{F0: 80, F1: 1, F2: 2}
                                        _ = v27
                                    default:
                                        _ = len(v3)
                                    }
                                case "\n\t":
                                    var v28 int32 = 0
                                    _ = v28
                                default:
                                    fmt.Println(v5, v6)
                                }
                            }
                            var v29 bool = (0 != uint64(v5))
                            fmt.Println(v29)
                            _ = v29
                        }
                    default:
                        if true {
                            var v30 int16 = v4
                            fmt.Println(v6, v2, v30)
                            v6 = true
                            _ = v30
                        } else {
                            var v31 int = -504
                            _ = len(v1)
                            for fi32 := 0; fi32 < 10; fi32++ {
                                var v33 uint = 1
                                var v34 uint64 = 0
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            _ = v31
                        }
                    }
                    if (true && (!(false || true))) {
                        v35 := GenS2{F0: 2147483647}
                        _ = v35
                    }
                    var v36 bool = v6
                    _ = v36
                } else {
                    v21 = 0.1
                    for fi37 := 0; fi37 < 9; fi37++ {
                        var v38 float64 = 3.14159
                        _ = fi37
                        _ = v38
                    }
                    if (42 == 424) {
                        v2 = v6
                        switch "" {
                        case "":
                            var v39 uint = 255
                            _ = v39
                        default:
                            switch 360 {
                            case -813:
                                if v2 {
                                    v40 := GenS5{F0: 42, F1: -9223372036854775808}
                                    var v41 float64 = float64(1.5)
                                    _ = v40
                                    _ = v41
                                }
                            case -9223372036854775808:
                                switch int(v18) {
                                case 2147483647:
                                    _ = len(v3)
                                default:
                                    if (((!v6) || v6) || v2) {
                                        for fi42 := 0; fi42 < 17; fi42++ {
                                            fmt.Println(v20)
                                            v2 = (0 > 42)
                                            _ = fi42
                                        }
                                    }
                                }
                            case 4294967295:
                                v43 := GenS6{F0: 65535}
                                _ = v43
                            default:
                                if (true && v6) {
                                    var v44 int16 = genFn1()
                                    var v45 uint8 = 48
                                    _ = v44
                                    _ = v45
                                } else {
                                    if (-413 < 1) {
                                        v4 -= genFn0(v20, (v20 | v20), 18446744073709551615)
                                        switch ((("" + "\n\t") + (("" + "") + (("\n\t" + "hello") + ("a" + "")))) + ("foo bar" + "foo bar")) {
                                        case "a":
                                            if v6 {
                                                v6 = ((1e10 - ((0.0 / 1.5) * (0.1 + 0.0))) < 0.0)
                                            }
                                        case "foo bar":
                                            if (-120 <= 100) {
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                            }
                                        default:
                                            v4 -= (v4 - (v4 ^ 32767))
                                        }
                                    } else {
                                        if (1 == 42) {
                                            _ = len(v3)
                                        } else {
                                            var v46 int16 = v4
                                            if v6 {
                                                fmt.Println(v46, v19)
                                                v3 = append(v3, "hello")
                                                for fi47 := 0; fi47 < 4; fi47++ {
                                                    v18 = (v18 >> ((123) & 63))
                                                    _ = fi47
                                                }
                                            }
                                            switch ("\n\t" + "a") {
                                            case "\n\t":
                                                var v48 float32 = 0.1
                                                _ = v48
                                            case "":
                                                for fi49 := 0; fi49 < 20; fi49++ {
                                                    fmt.Println(v20, v21)
                                                    _ = fi49
                                                }
                                            case "a":
                                                for fi50 := 0; fi50 < 2; fi50++ {
                                                    v46 += genFn0((^(v20)), v20, 1)
                                                    v46 -= v4
                                                    v6 = v6
                                                    _ = fi50
                                                }
                                            default:
                                                var v51 int = 128
                                                _ = v51
                                            }
                                            _ = v46
                                        }
                                        if (9223372036854775807 <= v19) {
                                            var v52 uint32 = (^(v5))
                                            for fi53 := 0; fi53 < 4; fi53++ {
                                                var v54 float32 = 1.5
                                                var v55 uint = 4294967295
                                                _ = len(v1)
                                                _ = fi53
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = v52
                                        }
                                    }
                                    v2 = (-1 == (v4 * (v4 - v4)))
                                    v5 += 42
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println(v6)
        } else {
            v56 := GenS5{F0: 1, F1: v19}
            var v57 int = -224
            _ = v56
            _ = v57
        }
    }
    var v58 int64 = int64(v20)
    if len(v1) > 0 {
        v1[0] = 674
    }
    switch ("hello" + "\n\t") {
    case "hello":
        switch 213 {
        case 127:
            v59 := GenS7{F0: (1 < 614), F1: 1, F2: true}
            _ = v59
        case 149:
            v20 -= 0
        default:
            var v60 uint32 = v5
            _ = v60
        }
    default:
        var v61 bool = ((v6 && false) || (!(18446744073709551615 <= 42)))
        _ = v61
    }
    if (!v2) {
        switch ("\n\t" + "a") {
        case "a":
            for fi62 := 0; fi62 < 9; fi62++ {
                v63 := GenS3{F0: -1, F1: (v5 * (v5 / ((v5) | 1))), F2: 84}
                v21 += (1.5 + v21)
                _ = len(v1)
                _ = fi62
                _ = v63
            }
        case "\n\t":
            v64 := GenS8{F0: (-(v58))}
            _ = v64
        case "hello":
            fmt.Println(v19, v19)
        default:
            switch 127 {
            case -2147483648:
                var v65 float32 = (1e10 / 1.0)
                _ = v65
            case -283:
                for fi66 := 0; fi66 < 8; fi66++ {
                    fmt.Println(v19, v21, v6)
                    _ = fi66
                }
            }
        }
    } else {
        var v67 int64 = v19
        var v68 uint16 = 1
        v69 := GenS3{F0: int8(v68), F1: 1, F2: 0}
        _ = v67
        _ = v68
        _ = v69
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v70 bool = true
    var v71 string = "hello"
    switch 840 {
    case 128:
        var v72 int16 = v4
        _ = v72
    case 562:
        var v73 string = (("" + "\n\t") + "foo bar")
        _ = v73
    }
    var v74 uint64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v70
    _ = v71
    _ = v74
}
