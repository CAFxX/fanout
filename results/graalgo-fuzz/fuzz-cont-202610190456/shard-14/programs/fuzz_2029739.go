package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int }
func genFn0() uint {
    var v5 uint32 = 2147483647
    if (72 >= 127) {
        v6 := []int{}
        switch "" {
        case "":
            v5 -= (^(v5))
        }
        _ = v6
    } else {
        if (((0.0 + 1.5) > 0.1) && (true || (((float32(1e10) + -2.25) * 1e10) >= 3.14159))) {
            switch (("\n\t" + "\n\t") + "foo bar") {
            case "":
                if (32 <= 779) {
                    for fi7 := 0; fi7 < 18; fi7++ {
                        if true {
                            var v8 float32 = 0.0
                            if (2 >= 394) {
                                switch ("a" + "\n\t") {
                                case "hello":
                                    v5 -= (v5 % (((v5 & v5)) | 1))
                                case "foo bar":
                                    v9 := GenS1{F0: 1, F1: -548, F2: 219}
                                    _ = v9
                                }
                            }
                            v5 = v5
                            _ = v8
                        }
                        _ = fi7
                    }
                } else {
                    v10 := []int{4294967295}
                    switch "foo bar" {
                    case "foo bar":
                        v5 = 2147483647
                    case "hello":
                        var v11 float64 = 0.1
                        _ = v11
                    case "":
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    default:
                        for fi12 := 0; fi12 < 7; fi12++ {
                            v13 := GenS1{F0: v5, F1: 1, F2: 0}
                            for fi14 := 0; fi14 < 9; fi14++ {
                                _ = len(v10)
                                _ = fi14
                            }
                            var v15 uint64 = uint64(v5)
                            _ = fi12
                            _ = v13
                            _ = v15
                        }
                    }
                    var v16 uint8 = 255
                    _ = v10
                    _ = v16
                }
            case "foo bar":
                if true {
                    v17 := []int{int(v5), 42}
                    _ = v17
                } else {
                    var v18 uint8 = 1
                    v18 += (v18 >> ((42) & 63))
                    _ = v18
                }
            }
            v19 := GenS1{F0: 2, F1: 405, F2: 29}
            _ = v19
        } else {
            v20 := []int{-2147483648, 32768, 32768, -2147483648}
            fmt.Println(v5)
            _ = v20
        }
    }
    switch 2147483647 {
    case 2147483647:
        switch -9223372036854775808 {
        case -2147483648:
            v21 := []int{9223372036854775807}
            _ = v21
        case -1:
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 float64 = 3.14159
                _ = fi22
                _ = v23
            }
        }
    case 198:
        v24 := []int{}
        _ = v24
    case 1:
        var v25 uint32 = v5
        _ = v25
    default:
        if (!true) {
            v26 := GenS2{F0: v5}
            v5 -= (v5 << ((117) & 63))
            if true {
                v27 := GenS2{F0: (v5 ^ 2147483648)}
                _ = v27
            }
            _ = v26
        } else {
            for fi28 := 0; fi28 < 6; fi28++ {
                if (127 < 145) {
                    var v29 int = 4294967295
                    v30 := []int{(^(v29)), v29, (fi28 >> ((65535) & 63))}
                    _ = v29
                    _ = v30
                } else {
                    if false {
                        var v31 int = -377
                        v32 := GenS3{F0: 18446744073709551615, F1: 1000, F2: ("foo bar" + (("hello" + "") + ("hello" + "hello")))}
                        _ = v31
                        _ = v32
                    }
                    v5 -= 807
                }
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 float32 = (3.14159 + float32((1e10 - 0.0)))
                    _ = fi33
                    _ = v34
                }
                _ = fi28
            }
            switch 531 {
            case 9223372036854775807:
                if false {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        for fi36 := 0; fi36 < 8; fi36++ {
                            switch (^(fi36)) {
                            case 128:
                                v37 := []int{}
                                _ = v37
                            }
                            v5 = v5
                            v5 -= 1
                            v38 := []int{fi36}
                            v38 = append(v38, (fi36 | fi36))
                            _ = fi36
                            _ = v38
                        }
                        var v39 uint64 = 0
                        _ = fi35
                        _ = v39
                    }
                    if true {
                        var v40 int16 = 42
                        var v41 int8 = 127
                        v42 := []int{806, 4294967295}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
            case -1:
                var v43 float32 = 1.5
                _ = v43
            }
            var v44 string = "hello"
            _ = v44
        }
    }
    _ = v5
    return 791
}
func genFn1(p0 int8, p1 uint) uint64 {
    if false {
        p1 -= (p1 - 880)
        var v45 int8 = (p0 * (p0 | 0))
        p1 -= p1
        _ = v45
    }
    v46 := []int{}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v46
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 float64 = (1e10 - ((3.14159 + (0.1 / -2.25)) * ((float64(float32(1.5)) / 1.5) * (((1e10 + -2.25) * 3.14159) / -2.25))))
    var v4 int64 = 42
    if (false || ((-1 <= -69) && false)) {
        v47 := GenS1{F0: 4294967295, F1: -1, F2: 4294967295}
        v1["\n\t"] = ("a" + "hello")
        _ = v47
    } else {
        v48 := GenS4{F0: 255, F1: 42, F2: 2}
        if (255 >= 209) {
            v3 = v3
            var v49 int8 = -37
            switch -9223372036854775808 {
            case 42:
                if (!true) {
                    v50 := GenS3{F0: 524, F1: -267, F2: "hello"}
                    _ = v50
                } else {
                    v51 := GenS5{F0: genFn1(v49, 18446744073709551615)}
                    switch "" {
                    case "a":
                        v3 = 0.1
                    case "":
                        v52 := GenS6{F0: -1}
                        _ = v52
                    case "foo bar":
                        if false {
                            v53 := GenS7{F0: -2147483648, F1: v49}
                            _ = v53
                        } else {
                            v54 := []int{2147483647}
                            fmt.Println(v3, v3, v3)
                            var v55 float32 = float32((float64(0.1) / v3))
                            _ = v54
                            _ = v55
                        }
                    }
                    _ = v51
                }
            case 32767:
                if ((false || (!(("a" + "hello") == (("hello" + "\n\t") + "hello")))) && ((true || (1 != 0)) && false)) {
                    var v56 int8 = (v49 ^ (v49 % ((v49) | 1)))
                    v57 := []int{-808, 128, 1}
                    _ = v56
                    _ = v57
                } else {
                    v49 += (v49 * v49)
                }
            }
            _ = v49
        } else {
            if false {
                v3 = ((3.14159 + v3) / float64(0.1))
                v58 := GenS6{F0: 0}
                _ = v58
            } else {
                v59 := []int{2147483647, 4294967295}
                v3 -= v3
                var v60 uint32 = 0
                _ = v59
                _ = v60
            }
        }
        var v61 uint32 = 477
        _ = v48
        _ = v61
    }
    v3 = -2.25
    v62 := GenS8{F0: v4}
    v4 = -9223372036854775808
    delete(v1, ("" + "a"))
    switch "\n\t" {
    case "hello":
        for fi63 := 0; fi63 < 10; fi63++ {
            v64 := GenS9{F0: (^(fi63))}
            var v65 int64 = (v4 & v4)
            v65 += (v4 * (v4 * 42))
            _ = fi63
            _ = v64
            _ = v65
        }
    default:
        _ = len(v2)
    }
    var v66 int8 = 100
    v1["\n\t"] = (((("\n\t" + "foo bar") + "") + "a") + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v62
    _ = v66
}
