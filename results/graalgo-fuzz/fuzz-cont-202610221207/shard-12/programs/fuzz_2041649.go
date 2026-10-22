package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int32; F1 uint64; F2 uint64 }
func genFn0() int8 {
    if true {
        var v6 uint64 = 4294967295
        _ = v6
    }
    switch -414 {
    case 32767:
        var v7 float64 = float64(1.5)
        _ = v7
    case 2147483647:
        v8 := GenS1{F0: 1, F1: false}
        _ = v8
    default:
        v9 := []int{-9223372036854775808, 42, 826, 9223372036854775807}
        _ = v9
    }
    if (0 < 1) {
        v10 := []int{2147483647, 2147483647, 32767}
        _ = v10
    } else {
        v11 := []int{32768}
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: false}
    _ = v12
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 2147483647
    var v3 float64 = 0.1
    v4 := make(map[string]bool)
    var v5 bool = (!((!(true && (true || (v3 > ((1e10 / (v3 + float64(0.0))) + v3))))) && false))
    if (4294967295 <= 18446744073709551615) {
        var v13 int64 = 9223372036854775807
        _ = v13
    } else {
        fmt.Println(v5, v3)
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 bool = ("foo bar" > "\n\t")
            v16 := GenS2{F0: 255, F1: 0, F2: 2147483647}
            _ = fi14
            _ = v15
            _ = v16
        }
        v5 = v5
    }
    if (42 != 0) {
        v5 = (65535 <= 322)
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint = 65535
            var v19 int8 = -35
            v20 := GenS3{F0: ("a" + "hello"), F1: 127, F2: (v18 >> (((v18 << (((v18 & 255)) & 63))) & 63))}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        switch 42 {
        case -470:
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS2{F0: 18446744073709551615, F1: 32767, F2: 0}
                _ = fi21
                _ = v22
            }
        case 9223372036854775807:
            for fi23 := 0; fi23 < 6; fi23++ {
                v5 = true
                v5 = false
                _ = fi23
            }
        case -590:
            var v24 int8 = 127
            _ = v24
        default:
            delete(v1, "")
        }
    }
    fmt.Println(v2, v3, v3)
    if (!((!(-1 > 1000)) || true)) {
        if false {
            if true {
                v5 = false
            } else {
                var v25 int16 = 0
                _ = v25
            }
            var v26 float64 = float64((1e10 / -2.25))
            v27 := GenS1{F0: 466, F1: true}
            _ = v26
            _ = v27
        }
        v28 := GenS3{F0: (("a" + "a") + ("hello" + "\n\t")), F1: 0, F2: 65535}
        _ = v28
    }
    v2 = (v2 + (v2 * v2))
    v29 := []int{(v2 << ((65535) & 63)), v2, 128, v2}
    if len(v29) > 0 {
        v29[0] = v2
    }
    switch "a" {
    case "\n\t":
        _ = len(v29)
    default:
        v30 := GenS2{F0: 18446744073709551615, F1: 65535, F2: 65535}
        _ = v30
    }
    v1[""] = true
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS3{F0: "foo bar", F1: 158, F2: 65535}
        _ = fi31
        _ = v32
    }
    v3 = v3
    if (1 <= -1) {
        v3 -= v3
        switch (v2 * (v2 & 32767)) {
        case -807:
            if len(v29) > 0 {
                v29[0] = (-(v2))
            }
        case 213:
            fmt.Println(v2, v2)
        default:
            for fi33 := 0; fi33 < 10; fi33++ {
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                _ = fi33
            }
        }
    } else {
        var v34 uint16 = 65535
        var v35 float32 = (((1e10 + (1e10 / 1.5)) / float32(((v3 * (1.5 - ((v3 / 1.0) / 1e10))) - float64(float32(float64(((-2.25 / 0.1) + (1.5 + (float32(float64(-2.25)) * 0.0))))))))) * 3.14159)
        for fi36 := 0; fi36 < 10; fi36++ {
            v35 = (float32(v3) - -2.25)
            switch v2 {
            case 127:
                switch (("" + "hello") + "") {
                case "hello":
                    var v37 bool = true
                    _ = v37
                }
            case 32768:
                if true {
                    delete(v1, "foo bar")
                }
            case 9223372036854775807:
                var v38 float32 = v35
                _ = v38
            default:
                v29 = append(v29, (v2 << ((2) & 63)))
            }
            if (0 == 1000) {
                switch "\n\t" {
                case "foo bar":
                    for fi39 := 0; fi39 < 19; fi39++ {
                        v5 = v5
                        var v40 uint8 = 1
                        v41 := GenS3{F0: ("" + "\n\t"), F1: v40, F2: 255}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                case "a":
                    var v42 float64 = (float64(((-2.25 / v35) - v35)) + ((float64(float32(v3)) - v3) * (v3 - v3)))
                    _ = v42
                case "hello":
                    v35 = (-2.25 - -2.25)
                default:
                    fmt.Println(v3)
                }
                if v5 {
                    v29 = append(v29, (fi36 % ((32768) | 1)))
                    for fi43 := 0; fi43 < 13; fi43++ {
                        v44 := GenS3{F0: "", F1: 127, F2: 1}
                        if ((1 != 716) || v5) {
                            v3 -= ((((3.14159 + v3) - v3) / (float64(v35) + v3)) / 1.5)
                            v45 := GenS4{F0: -994, F1: 0, F2: 199}
                            var v46 string = "foo bar"
                            _ = v45
                            _ = v46
                        } else {
                            switch (v2 ^ v2) {
                            case 32767:
                                var v47 string = ("foo bar" + "a")
                                _ = v47
                            case 2147483647:
                                v29 = append(v29, fi36)
                            case -733:
                                v29 = append(v29, fi43)
                            }
                            var v48 int16 = 42
                            v49 := GenS2{F0: 0, F1: (v34 & (v34 & (v34 & (^(v34))))), F2: 385}
                            _ = v48
                            _ = v49
                        }
                        fmt.Println(v35)
                        _ = fi43
                        _ = v44
                    }
                    var v50 int64 = -9223372036854775808
                    _ = v50
                }
            } else {
                var v51 int = (fi36 & fi36)
                _ = v51
            }
            _ = fi36
        }
        _ = v34
        _ = v35
    }
    var v52 int8 = 0
    v3 -= 1.5
    v29 = append(v29, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
