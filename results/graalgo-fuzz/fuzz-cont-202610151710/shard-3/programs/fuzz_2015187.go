package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 uint64) int64 {
    var v7 int32 = 42
    var v8 uint32 = (p0 | 4294967295)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32, p1 uint16) int32 {
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 18446744073709551615, F1: 247, F2: (^(fi9))}
        v11 := GenS1{F0: 18446744073709551615, F1: 20, F2: (fi9 + fi9)}
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 -= p0
    _ = p0
    _ = p1
    return (p0 % (((-(p0))) | 1))
}

func main() {
    var v1 int8 = -102
    var v2 float32 = 1e10
    var v3 string = ("\n\t" + "a")
    var v4 string = v3
    v5 := make(map[int]int)
    var v6 int32 = -1
    if false {
        if true {
            var v12 float64 = (-2.25 + ((0.1 - (1e10 / -2.25)) / (1.5 - 0.0)))
            if ((v6 + -27) == (v6 | genFn1((v6 >> ((0) & 63)), 742))) {
                var v13 uint16 = 0
                if (((0.0 * ((-2.25 / 1.0) / v2)) < (v2 * v2)) || false) {
                    switch 127 {
                    case 32768:
                        v3 = ("hello" + v3)
                    case -764:
                        v6 = -2147483648
                    default:
                        switch v4 {
                        case "foo bar":
                            for fi14 := 0; fi14 < 14; fi14++ {
                                switch ("" + ("foo bar" + "")) {
                                case "hello":
                                    var v15 float64 = float64(0.0)
                                    _ = v15
                                case "":
                                    v16 := GenS2{F0: 0}
                                    _ = v16
                                }
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    v13 = v13
                                    v18 := GenS1{F0: 454, F1: -1, F2: (fi14 - fi17)}
                                    if (true || (((18446744073709551615 == 42) && true) && false)) {
                                        var v19 string = ("a" + v4)
                                        _ = v19
                                    }
                                    _ = fi17
                                    _ = v18
                                }
                                _ = fi14
                            }
                        case "\n\t":
                            for fi20 := 0; fi20 < 11; fi20++ {
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    if (!(v3 == (v3 + v4))) {
                                        v22 := []int{}
                                        var v23 uint8 = 0
                                        _ = v22
                                        _ = v23
                                    }
                                    v24 := GenS1{F0: 677, F1: 659, F2: (fi20 % ((fi21) | 1))}
                                    _ = fi21
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 4; fi25++ {
                                    var v26 int32 = (v6 * v6)
                                    var v27 uint = 1
                                    if (42 >= (v1 << ((987) & 63))) {
                                        var v28 float64 = v12
                                        v2 = (v2 + v2)
                                        if false {
                                            fmt.Println(v28, v4, v4)
                                        }
                                        _ = v28
                                    } else {
                                        v29 := GenS1{F0: 1, F1: 1, F2: (fi20 ^ fi25)}
                                        _ = v29
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi20
                            }
                        case "a":
                            switch 4294967295 {
                            case 32768:
                                v13 += (v13 - 32767)
                            case 0:
                                var v30 bool = false
                                _ = v30
                            case 760:
                                switch 1 {
                                case 32768:
                                    fmt.Println(v2, v12)
                                case 32767:
                                    for fi31 := 0; fi31 < 6; fi31++ {
                                        v12 = 1e10
                                        delete(v5, (fi31 & int(v1)))
                                        var v32 int8 = v1
                                        _ = fi31
                                        _ = v32
                                    }
                                }
                            }
                        }
                    }
                    v1 = (v1 % ((v1) | 1))
                } else {
                    var v33 uint16 = (v13 & (v13 | v13))
                    v5[-9223372036854775808] = 42
                    if false {
                        _ = len(v5)
                    } else {
                        var v34 uint32 = 0
                        _ = v34
                    }
                    _ = v33
                }
                _ = v13
            }
            var v35 uint = 4294967295
            _ = v12
            _ = v35
        }
        switch v3 {
        case "":
            for fi36 := 0; fi36 < 7; fi36++ {
                v37 := []int{fi36, (fi36 * (-(fi36))), 32767}
                _ = fi36
                _ = v37
            }
        case "\n\t":
            if (!(v3 == v3)) {
                v2 += (v2 - (v2 / v2))
                v3 = ("hello" + "")
                v38 := []int{}
                _ = v38
            } else {
                v1 = 1
            }
        case "hello":
            v4 = v4
        }
        v5[4294967295] = 42
    }
    var v39 int8 = (v1 & v1)
    var v40 float32 = v2
    var v41 int = 16
    v2 -= ((v2 + v40) - v40)
    v42 := GenS2{F0: 2147483648}
    for fi43 := 0; fi43 < 16; fi43++ {
        fmt.Println(v2, v1)
        _ = fi43
    }
    v6 -= genFn1((v6 - (v6 + int32(v39))), 42)
    switch ((v4 + v4) + (v3 + v4)) {
    case "\n\t":
        fmt.Println(v3, v39, v3)
    case "a":
        v6 = v6
    case "hello":
        var v44 int = 32768
        _ = v44
    }
    var v45 uint16 = 255
    fmt.Println(v40)
    switch (v41 ^ 0) {
    case 9223372036854775807:
        if ((v3 + (v4 + v3)) != ("hello" + ("\n\t" + ((v3 + (v3 + (v4 + v3))) + (((v3 + "hello") + "\n\t") + v4))))) {
            v46 := GenS1{F0: 426, F1: -251, F2: (v41 ^ -1)}
            var v47 uint8 = 1
            var v48 int16 = -32768
            _ = v46
            _ = v47
            _ = v48
        }
    case 0:
        for fi49 := 0; fi49 < 14; fi49++ {
            for fi50 := 0; fi50 < 12; fi50++ {
                var v51 int = 127
                v4 = v3
                v52 := GenS1{F0: 302, F1: 0, F2: 573}
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = fi49
        }
    default:
        v53 := v5[v41]
        _ = v53
    }
    for fi54 := 0; fi54 < 18; fi54++ {
        var v55 int16 = 0
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        v57 := GenS3{F0: 1, F1: 255}
        v4 = "foo bar"
        for fi58 := 0; fi58 < 1; fi58++ {
            if (-50 <= (v39 ^ (v1 >> ((272) & 63)))) {
                var v59 int16 = 1
                delete(v5, 42)
                _ = v59
            } else {
                v60 := []int{fi58, (-(fi56)), (v41 % (((fi56 + (-(v41)))) | 1)), -165}
                if len(v60) >= 2 {
                    v60 = v60[1:len(v60)]
                }
                _ = v60
            }
            v61 := v5[fi56]
            var v62 float32 = v40
            _ = fi58
            _ = v61
            _ = v62
        }
        _ = fi56
        _ = v57
    }
    switch "" {
    case "a":
        v63 := GenS3{F0: 845, F1: 151}
        _ = v63
    case "":
        v39 += (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
