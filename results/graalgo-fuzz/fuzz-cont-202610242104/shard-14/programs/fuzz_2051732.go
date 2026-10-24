package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 string; F2 int64 }
func genFn0(p0 int64, p1 int) int64 {
    if (-646 >= -32768) {
        v6 := []int{p1, (p1 | (p1 >> ((65535) & 63)))}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v7 float32 = 3.14159
        _ = v6
        _ = v7
    }
    var v8 uint = 869
    _ = p0
    _ = p1
    _ = v8
    return (^(p0))
}
func genFn1(p0 uint8, p1 uint64) int64 {
    for fi9 := 0; fi9 < 12; fi9++ {
        switch int(p1) {
        case 0:
            v10 := []int{}
            _ = v10
        case -1:
            v11 := GenS1{F0: -1, F1: true}
            _ = v11
        case -9223372036854775808:
            v12 := []int{fi9, (fi9 ^ 492)}
            _ = v12
        }
        _ = fi9
    }
    v13 := GenS1{F0: 42, F1: (1 == -1)}
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := []int{(fi14 ^ fi14), (-(fi14)), fi14, (fi14 >> ((934) & 63))}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return genFn0(1, -240)
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int64 = 262
    var v3 float64 = -2.25
    var v4 int8 = 100
    var v5 bool = true
    switch "hello" {
    case "hello":
        v16 := []int{}
        _ = v16
    case "\n\t":
        if v5 {
            if false {
                v3 -= v3
                for fi17 := 0; fi17 < 13; fi17++ {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if (-1 <= 42) {
                            fmt.Println(v3, v2)
                            v5 = ((0 != 62) && v5)
                        } else {
                            v5 = (458 <= 1)
                            fmt.Println(v1)
                            var v19 int16 = 0
                            _ = v19
                        }
                        _ = fi18
                    }
                    var v20 bool = (false || (false && false))
                    v1 = float32((1.5 + float64((0.0 / v1))))
                    _ = fi17
                    _ = v20
                }
            }
        }
    default:
        v1 = ((((1.5 - ((v1 * (float32(0.1) + float32(v3))) - float32((3.14159 + v3)))) / 1e10) - v1) * 1e10)
    }
    switch 127 {
    case 42:
        switch "a" {
        case "":
            fmt.Println(v1, v4, v3)
        }
    case -1:
        var v21 string = ""
        _ = v21
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v3 = v3
                _ = fi23
            }
            if true {
                for fi24 := 0; fi24 < 6; fi24++ {
                    fmt.Println(v1)
                    v1 = ((0.1 + v1) * (v1 + (float32(3.14159) - 0.1)))
                    var v25 uint64 = 1
                    fmt.Println(v1, v4, v3)
                    _ = fi24
                    _ = v25
                }
            }
            _ = fi22
        }
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        switch ((("hello" + "foo bar") + "") + "foo bar") {
        case "\n\t":
            v4 = v4
        case "a":
            var v27 float32 = 1e10
            _ = v27
        }
        _ = fi26
    }
    v3 -= v3
    if ((v2 <= genFn1(42, 687)) || v5) {
        switch 32768 {
        case 9223372036854775807:
            var v28 string = "a"
            _ = v28
        case 32768:
            if (!((!v5) || v5)) {
                v29 := GenS1{F0: 2147483647, F1: true}
                var v30 int = 127
                v31 := []int{(v30 - v30), v30}
                _ = v29
                _ = v30
                _ = v31
            }
        default:
            v32 := []int{}
            _ = v32
        }
        v33 := GenS1{F0: -218, F1: (genFn1(22, 18446744073709551615) >= v2)}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch "" {
            case "":
                if v5 {
                    v35 := GenS2{F0: 465, F1: (!v5), F2: 42}
                    if v5 {
                        if v5 {
                            v2 = v2
                        }
                        for fi36 := 0; fi36 < 15; fi36++ {
                            v37 := []int{9223372036854775807, fi34, 4294967295}
                            switch (^(fi34)) {
                            case 127:
                                v38 := GenS3{F0: 532}
                                _ = v38
                            case 128:
                                var v39 uint = 4294967295
                                _ = v39
                            case 9223372036854775807:
                                switch ("foo bar" + "") {
                                case "":
                                    switch fi36 {
                                    case 0:
                                        var v40 string = "hello"
                                        _ = v40
                                    case 42:
                                        var v41 int16 = 919
                                        _ = v41
                                    case 1:
                                        var v42 string = ((("" + ("a" + ("hello" + "\n\t"))) + "foo bar") + "hello")
                                        fmt.Println(v4, v1, v4)
                                        _ = v42
                                    }
                                case "\n\t":
                                    v43 := GenS2{F0: 2, F1: false, F2: 255}
                                    _ = v43
                                }
                                v2 = (v2 * (-(v2)))
                                var v44 int64 = -9223372036854775808
                                _ = v44
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                    _ = v35
                }
            default:
                v45 := GenS4{F0: 4294967295, F1: 243}
                _ = v45
            }
            _ = fi34
        }
        v5 = false
    }
    var v46 uint = 65535
    v3 -= v3
    if ((!false) || (3.14159 > float32(float64(-2.25)))) {
        v1 = float32((v3 * 0.0))
        for fi47 := 0; fi47 < 1; fi47++ {
            switch ("" + ("" + ("hello" + (("hello" + "a") + (("\n\t" + "\n\t") + ("foo bar" + (((((("\n\t" + (("foo bar" + "") + "foo bar")) + "foo bar") + (("foo bar" + "a") + "hello")) + "a") + "") + "\n\t"))))))) {
            case "a":
                if (65535 > 2147483647) {
                    var v48 int16 = 1
                    _ = v48
                }
            case "hello":
                var v49 uint32 = 2147483648
                _ = v49
            }
            for fi50 := 0; fi50 < 16; fi50++ {
                var v51 float32 = v1
                v52 := GenS3{F0: v46}
                if ((v5 || (false || v5)) || (3.14159 < v3)) {
                    v1 = ((((float32(v3) / 3.14159) * (v1 + (0.0 - (((v1 / v51) * v51) / 1e10)))) * 1e10) - 0.1)
                } else {
                    v2 = v2
                }
                if v5 {
                    v53 := GenS5{F0: (v4 * v4)}
                    v54 := []int{(fi47 & (fi47 >> (((^(v46))) & 63))), (fi47 - 4294967295), (-(fi50)), fi47}
                    fmt.Println(v3, v51)
                    _ = v53
                    _ = v54
                } else {
                    v55 := []int{fi47, fi50, (-(fi47)), fi47}
                    _ = v55
                }
                _ = fi50
                _ = v51
                _ = v52
            }
            if (981 >= 42) {
                for fi56 := 0; fi56 < 17; fi56++ {
                    switch ("foo bar" + ("\n\t" + ((("hello" + "hello") + "hello") + "a"))) {
                    case "foo bar":
                        v57 := GenS6{F0: 255}
                        _ = v57
                    case "hello":
                        v2 += (v2 | genFn0(v2, 2147483647))
                    case "a":
                        var v58 float64 = ((1.5 * v3) + v3)
                        _ = v58
                    }
                    v1 -= 0.0
                    v59 := []int{802, -1, fi47, (fi47 << (((^(v46))) & 63))}
                    _ = fi56
                    _ = v59
                }
                for fi60 := 0; fi60 < 5; fi60++ {
                    v5 = false
                    _ = fi60
                }
                v46 += (^(v46))
            }
            _ = fi47
        }
    }
    v61 := GenS4{F0: 65535, F1: 1}
    v62 := GenS7{F0: 32767, F1: "a", F2: -534}
    var v63 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v61
    _ = v62
    _ = v63
}
