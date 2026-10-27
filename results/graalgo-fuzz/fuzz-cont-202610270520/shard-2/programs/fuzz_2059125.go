package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8; F1 bool; F2 string }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint; F1 int8; F2 int8 }
type GenS8 struct { F0 bool; F1 int32 }
func genFn0(p0 string, p1 uint64, p2 int8) uint32 {
    var v7 string = ("\n\t" + (((p0 + ((p0 + ("foo bar" + p0)) + ("a" + "hello"))) + p0) + "hello"))
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 147
}
func genFn1(p0 uint64, p1 int32, p2 int64) uint {
    p1 = (p1 + (p1 % ((-2147483648) | 1)))
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 65535
    var v3 string = ("" + "")
    var v4 bool = ("" > v3)
    v5 := make(map[int]bool)
    var v6 string = v3
    if v4 {
        v1 = -83
    }
    switch -844 {
    case -820:
        v9 := v5[2147483647]
        _ = v9
    default:
        fmt.Println(v2, v6, v2)
    }
    var v10 float64 = 3.14159
    switch -9223372036854775808 {
    case 4294967295:
        v11 := GenS1{F0: 42, F1: v6}
        _ = v11
    case 1:
        v1 = v1
    }
    var v12 int = int(v1)
    var v13 string = v3
    for fi14 := 0; fi14 < 13; fi14++ {
        v1 -= v1
        if v4 {
            v15 := GenS1{F0: 981, F1: v6}
            fmt.Println(v13, v13)
            switch (v12 | v12) {
            case -9223372036854775808:
                v16 := GenS1{F0: 385, F1: v3}
                _ = v16
            default:
                fmt.Println(v4, v4)
            }
            _ = v15
        }
        v5[(fi14 >> ((831) & 63))] = v4
        _ = fi14
    }
    if ((1e10 < v10) || v4) {
        v13 = v13
        v3 = v6
        fmt.Println(v4, v10, v3)
    }
    var v17 uint32 = genFn0(v13, 18446744073709551615, v1)
    if (-1 > -2147483648) {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (249 == 1) {
                switch (-(fi18)) {
                case 4294967295:
                    v19 := GenS1{F0: 1, F1: v13}
                    _ = v19
                case 128:
                    v13 = v6
                }
            } else {
                var v20 float64 = v10
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 int32 = -672
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: 2, F1: "hello"}
                _ = v20
                _ = v23
            }
            v13 = v13
            v10 = (v10 + ((v10 - float64((((-2.25 * (((1e10 + 1.5) / 0.1) - (3.14159 - (-2.25 / 1.5)))) + 3.14159) / (3.14159 * 0.1)))) / (1.5 / v10)))
            _ = fi18
        }
        v24 := GenS1{F0: 87, F1: "foo bar"}
        _ = v24
    }
    switch (v12 - v12) {
    case 4294967295:
        fmt.Println(v6, v3)
    case 127:
        if (!v4) {
            var v25 bool = (0.1 >= (((((0.0 / 0.1) * 0.0) - 1e10) - (float32(float64(1.5)) - (-2.25 / float32(v10)))) - 3.14159))
            for fi26 := 0; fi26 < 17; fi26++ {
                if false {
                    switch v12 {
                    case 4294967295:
                        var v27 int32 = -804
                        _ = v27
                    case -967:
                        switch 9223372036854775807 {
                        case 4294967295:
                            delete(v5, fi26)
                        case -781:
                            switch v12 {
                            case -916:
                                v28 := GenS2{F0: 0, F1: 1}
                                _ = v28
                            }
                        case 217:
                            if (v1 >= v1) {
                                var v29 bool = v4
                                switch v3 {
                                case "hello":
                                    delete(v5, 625)
                                case "foo bar":
                                    if true {
                                        if true {
                                            v1 -= (v1 % (((v1 / ((v1) | 1))) | 1))
                                            v30 := GenS3{F0: 166}
                                            for fi31 := 0; fi31 < 10; fi31++ {
                                                var v32 int32 = 1
                                                v33 := GenS4{F0: 0, F1: -765}
                                                _ = fi31
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = v30
                                        }
                                        switch v6 {
                                        case "foo bar":
                                            switch -1 {
                                            case 0:
                                                for fi34 := 0; fi34 < 19; fi34++ {
                                                    v35 := GenS5{F0: (v1 << ((genFn1(2, 680, int64(fi34))) & 63)), F1: ("hello" < v13), F2: ""}
                                                    v29 = (!(true || false))
                                                    v36 := GenS5{F0: int8(v17), F1: false, F2: v13}
                                                    _ = fi34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                        case "\n\t":
                                            v4 = (255 == 288)
                                        case "hello":
                                            v37 := GenS6{F0: 0, F1: 42}
                                            _ = v37
                                        }
                                    }
                                default:
                                    v6 = v6
                                }
                                _ = v29
                            } else {
                                v12 += (fi26 % ((fi26) | 1))
                                v38 := GenS1{F0: 2, F1: "a"}
                                _ = v38
                            }
                        default:
                            if ((1 > -9223372036854775808) && v25) {
                                var v39 uint64 = 42
                                _ = v39
                            }
                        }
                    }
                    fmt.Println(v17, v4)
                } else {
                    var v40 string = (v13 + v6)
                    switch v6 {
                    case "hello":
                        switch (v12 + fi26) {
                        case -199:
                            fmt.Println(v3, v3, v1)
                        case 640:
                            var v41 uint8 = 1
                            _ = v41
                        }
                    case "a":
                        v42 := GenS7{F0: genFn1(805, 42, -242), F1: v1, F2: -1}
                        _ = v42
                    case "":
                        var v43 float64 = v10
                        _ = v43
                    }
                    v44 := GenS8{F0: true, F1: -759}
                    _ = v40
                    _ = v44
                }
                _ = fi26
            }
            _ = v25
        }
    default:
        v4 = true
    }
    var v45 string = ((("hello" + v13) + v6) + "foo bar")
    for fi46 := 0; fi46 < 8; fi46++ {
        if false {
            v3 = v45
            for fi47 := 0; fi47 < 20; fi47++ {
                fmt.Println(v3, v6, v17)
                v6 = v13
                if false {
                    var v48 uint16 = 1
                    var v49 int = (fi46 - 32767)
                    v5[(fi46 / ((fi47) | 1))] = v4
                    var v50 uint64 = 18446744073709551615
                    v51 := []int{v12, v49, (fi46 * int(v50))}
                    _ = v48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = fi47
            }
            for fi52 := 0; fi52 < 16; fi52++ {
                v53 := v5[(v12 & 311)]
                _ = fi52
                _ = v53
            }
        } else {
            var v54 uint = 835
            _ = v54
        }
        v55 := []int{-2147483648}
        _ = fi46
        _ = v55
    }
    if ((73 != 785) && false) {
        var v56 uint32 = 715
        var v57 float32 = 0.0
        _ = v56
        _ = v57
    }
    switch (v45 + (("hello" + (("" + "foo bar") + "foo bar")) + v45)) {
    case "\n\t":
        if false {
            fmt.Println(v13, v45)
            fmt.Println(v1, v13, v3)
            v58 := []int{(v12 | v12), v12}
            _ = v58
        }
    }
    v59 := GenS7{F0: 255, F1: (v1 >> ((2) & 63)), F2: (v1 & v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v59
    _ = v6
}
