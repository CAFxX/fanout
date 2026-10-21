package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint64; F1 int8; F2 string }
func genFn0(p0 int8, p1 int8) int {
    p1 += (p1 + p1)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 1
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := []int{32768, (fi7 | fi7)}
            if (4294967295 >= 0) {
                var v10 float32 = 1.5
                var v11 int = 42
                _ = v10
                _ = v11
            }
            var v12 uint8 = 166
            _ = fi8
            _ = v12
            _ = v9
        }
        _ = fi7
    }
    var v13 uint = 157
    v13 += v13
    switch 32767 {
    case -221:
        if false {
            var v14 uint32 = 42
            v14 += v14
            _ = v14
        }
    case 459:
        if false {
            v13 = 2
            v13 -= v13
            if false {
                v15 := GenS1{F0: (v13 >> ((v13) & 63)), F1: 255}
                v16 := []int{genFn0(0, 1), 32767}
                v13 = (v13 << (((v13 - (v13 | v13))) & 63))
                _ = v15
                _ = v16
            }
        } else {
            fmt.Println(v13)
            fmt.Println(v13)
        }
    case 128:
        v13 = (v13 ^ v13)
    default:
        switch ("hello" + "\n\t") {
        case "a":
            v13 -= (v13 | (^(v13)))
        default:
            switch 1 {
            case 0:
                v13 = v13
            case 9223372036854775807:
                v17 := []int{genFn0(100, 0), 32768, -2147483648, genFn0(42, 100)}
                _ = v17
            case 4294967295:
                if true {
                    v18 := []int{1}
                    switch -2147483648 {
                    case 9223372036854775807:
                        for fi19 := 0; fi19 < 15; fi19++ {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                var v21 int16 = 516
                                _ = fi20
                                _ = v21
                            }
                            _ = fi19
                        }
                    case 2147483647:
                        if true {
                            if len(v18) > 0 {
                                v18[0] = 673
                            }
                            fmt.Println(v13)
                            v13 = v13
                        }
                    case 464:
                        if ((float64(3.14159) < -2.25) || false) {
                            v22 := GenS1{F0: v13, F1: 113}
                            var v23 string = ""
                            var v24 uint64 = 249
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            fmt.Println(v13)
                            var v25 int8 = 96
                            if (0 == 128) {
                                v13 += (v13 ^ 18446744073709551615)
                                v18 = append(v18, 0)
                            } else {
                                v26 := GenS2{F0: 156, F1: (v13 & 65535)}
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    v25 += (v25 >> ((v13) & 63))
                                    v28 := GenS1{F0: v13, F1: 128}
                                    _ = fi27
                                    _ = v28
                                }
                                _ = v26
                            }
                            _ = v25
                        }
                    }
                    fmt.Println(v13)
                    _ = v18
                } else {
                    v29 := []int{}
                    for fi30 := 0; fi30 < 20; fi30++ {
                        if false {
                            v13 = (^(v13))
                        }
                        if true {
                            var v31 uint8 = 0
                            var v32 float64 = (-2.25 * 1e10)
                            for fi33 := 0; fi33 < 9; fi33++ {
                                v34 := GenS1{F0: uint(v31), F1: 1}
                                _ = fi33
                                _ = v34
                            }
                            _ = v31
                            _ = v32
                        }
                        _ = fi30
                    }
                    _ = v29
                }
            default:
                v13 += v13
            }
        }
    }
    _ = v13
    return 1
}

func main() {
    v1 := []int{428, -404, -432, 42}
    var v2 uint32 = 1
    v3 := []string{("foo bar" + "a")}
    var v4 int32 = 42
    var v5 bool = (!(!false))
    var v6 int16 = 1
    v3 = append(v3, ((((("\n\t" + "") + "hello") + "") + "") + ""))
    if (65535 <= 1) {
        if ((^(v4)) > -1) {
            var v35 uint32 = (^(v2))
            v6 = (v6 & 32767)
            _ = v35
        } else {
            if (0 < 127) {
                v1 = append(v1, -1)
                fmt.Println(v5, v5)
            }
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = len(v1)
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 bool = (!v5)
        switch "foo bar" {
        case "":
            fmt.Println(v5)
        case "a":
            switch ("a" + "\n\t") {
            case "foo bar":
                fmt.Println(v2, v6)
            }
        case "hello":
            v5 = v5
        }
        v37 = false
        _ = fi36
        _ = v37
    }
    var v38 uint16 = 0
    fmt.Println(v4, v2, v6)
    fmt.Println(v38)
    v38 = (v38 - (v38 + (^(v38))))
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS2{F0: 2, F1: 0}
        _ = fi39
        _ = v40
    }
    v6 += -1
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS3{F0: -1, F1: -690}
        switch (-(fi41)) {
        case 728:
            var v43 string = ("hello" + "")
            _ = v43
        case -515:
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 int16 = 32767
                var v46 int8 = -90
                _ = fi44
                _ = v45
                _ = v46
            }
        default:
            for fi47 := 0; fi47 < 20; fi47++ {
                v48 := GenS2{F0: 2, F1: 18446744073709551615}
                _ = fi47
                _ = v48
            }
        }
        var v49 int = fi41
        _ = fi41
        _ = v42
        _ = v49
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        v38 = 255
        _ = fi50
    }
    var v51 float64 = (3.14159 / 1.5)
    var v52 int16 = v6
    v53 := GenS4{F0: v52, F1: (v52 - (-(v6)))}
    fmt.Println(v6, v38)
    v54 := GenS5{F0: 4294967295, F1: int8(v38), F2: "foo bar"}
    v5 = true
    for fi55 := 0; fi55 < 14; fi55++ {
        v56 := GenS2{F0: 0, F1: 255}
        if len(v3) > 0 {
            v3[0] = "a"
        }
        var v57 int16 = v52
        _ = fi55
        _ = v56
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
