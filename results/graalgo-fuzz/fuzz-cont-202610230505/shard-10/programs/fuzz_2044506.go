package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int8) bool {
    v5 := []int{128, -143}
    v5 = append(v5, 42)
    var v6 uint32 = 396
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 int = -58
    v2 := make(map[string]string)
    var v3 float32 = (1e10 - 1e10)
    var v4 bool = true
    fmt.Println(v3, v1)
    if (-32768 < -818) {
        var v7 uint = 585
        _ = v7
    } else {
        v2[""] = "foo bar"
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v3)
        _ = fi8
    }
    v9 := GenS1{F0: 2}
    if false {
        if v4 {
            v3 += 3.14159
            fmt.Println(v1)
        }
        v2["foo bar"] = "a"
    } else {
        if true {
            if (v4 || (0.1 <= 0.0)) {
                fmt.Println(v3, v1, v4)
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                v2["\n\t"] = ("foo bar" + "foo bar")
                _ = fi10
            }
        } else {
            var v11 uint64 = 409
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int32 = 42
                v14 := GenS2{F0: ((("\n\t" + "foo bar") + "\n\t") + ("hello" + ("hello" + "\n\t")))}
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 bool = (false && v4)
            _ = v11
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint16 = 65535
        _ = fi16
        _ = v17
    }
    v3 = float32(1.5)
    fmt.Println(v1)
    switch (v1 << ((340) & 63)) {
    case -1:
        switch 32768 {
        case -9223372036854775808:
            v18 := []int{(v1 + (v1 - v1)), (v1 | (v1 * (-(v1)))), v1}
            _ = v18
        case 92:
            v4 = v4
        default:
            switch ("hello" + ("foo bar" + (("\n\t" + ("a" + (("\n\t" + "") + "\n\t"))) + ("foo bar" + "a")))) {
            case "":
                fmt.Println(v4, v3, v1)
            case "hello":
                v19 := GenS3{F0: 65535, F1: (("a" + "a") + "hello"), F2: 776}
                _ = v19
            default:
                if (1 != 65535) {
                    if true {
                        var v20 string = "hello"
                        v21 := GenS1{F0: 74}
                        v4 = v4
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS1{F0: 0}
                    _ = v22
                } else {
                    var v23 uint16 = 42
                    v24 := []int{32768, v1}
                    if (v4 && (!(true || (("\n\t" + (("" + "\n\t") + "\n\t")) == "foo bar")))) {
                        v1 -= v1
                        delete(v2, "")
                    }
                    _ = v23
                    _ = v24
                }
            }
        }
    case -845:
        if (!true) {
            var v25 uint32 = 65535
            v1 -= v1
            _ = v25
        }
    case 544:
        for fi26 := 0; fi26 < 12; fi26++ {
            if (!(1 > 42)) {
                if false {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if v4 {
                            v4 = v4
                            v2[("hello" + "")] = "foo bar"
                            var v28 float32 = (v3 + v3)
                            _ = v28
                        }
                        _ = fi27
                    }
                }
                var v29 uint8 = 128
                var v30 uint = 4294967295
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 float64 = 3.14159
                    v33 := []int{v1}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v34 int64 = 42
                _ = v34
            }
            var v35 uint32 = 2147483647
            _ = fi26
            _ = v35
        }
    default:
        v3 -= v3
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            if (-1 > 1000) {
                v4 = (v4 || (v4 && (-67 >= -1)))
                v1 = (v1 * (v1 >> ((255) & 63)))
            } else {
                fmt.Println(fi37, fi37)
                v3 += 0.1
                v38 := GenS4{F0: 42, F1: 65535, F2: 127}
                _ = v38
            }
            v39 := GenS3{F0: 1, F1: "", F2: 65535}
            _ = fi37
            _ = v39
        }
        fmt.Println(v3)
        _ = fi36
    }
    var v40 bool = false
    fmt.Println(v4)
    v3 = (((v3 + 1.5) + ((v3 / (1e10 / v3)) + (float32(3.14159) * (v3 - (v3 - ((v3 - v3) / (0.1 / v3))))))) * v3)
    _ = len(v2)
    var v41 uint = 381
    if v40 {
        for fi42 := 0; fi42 < 13; fi42++ {
            fmt.Println(fi42, v3, v40)
            fmt.Println(v1)
            var v43 bool = (!v4)
            _ = fi42
            _ = v43
        }
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint16 = 255
            _ = fi44
            _ = v45
        }
    } else {
        v46 := []int{(^(v1)), v1, (-(v1)), (-(v1))}
        if v40 {
            v40 = false
        } else {
            v47 := GenS2{F0: ((("hello" + "foo bar") + ("foo bar" + "hello")) + ("hello" + ""))}
            for fi48 := 0; fi48 < 3; fi48++ {
                fmt.Println(fi48)
                var v49 bool = (!(419 >= 1))
                _ = fi48
                _ = v49
            }
            _ = v47
        }
        _ = v46
    }
    v2["foo bar"] = "hello"
    if (true || v40) {
        v3 = (1e10 * v3)
    } else {
        fmt.Println(v4, v40)
        v50 := GenS5{F0: 18446744073709551615}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v9
}
