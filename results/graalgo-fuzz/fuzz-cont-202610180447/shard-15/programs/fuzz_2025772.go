package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int16; F2 int16 }
type GenS7 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 uint = 2
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    var v4 int64 = 844
    var v5 float32 = 1.5
    if true {
        v6 := v3["foo bar"]
        _ = v6
    }
    if false {
        if false {
            v5 = v5
            if true {
                for fi7 := 0; fi7 < 6; fi7++ {
                    if ((0 != -8) && (888 >= 231)) {
                        for fi8 := 0; fi8 < 18; fi8++ {
                            v9 := GenS1{F0: 0, F1: 4294967295}
                            v10 := GenS1{F0: 0, F1: 850}
                            switch (-(fi7)) {
                            case 128:
                                if false {
                                    if false {
                                        v11 := GenS2{F0: v1, F1: v2, F2: -1}
                                        _ = v11
                                    }
                                    v2 = (v2 % (((v2 % (((v2 % ((v2) | 1))) | 1))) | 1))
                                } else {
                                    var v12 uint = (v1 * 4294967295)
                                    v1 = (v12 | 1)
                                    switch (fi7 >> ((18446744073709551615) & 63)) {
                                    case 32768:
                                        v13 := GenS2{F0: v1, F1: (v2 ^ (v2 & (^(v2)))), F2: 462}
                                        _ = v13
                                    case 4294967295:
                                        var v14 float32 = v5
                                        v5 = ((v14 - v14) - float32((1.5 + (0.1 / 1.5))))
                                        _ = v14
                                    case -640:
                                        switch (("foo bar" + "") + "foo bar") {
                                        case "a":
                                            switch (fi8 >> (((v12 | v1)) & 63)) {
                                            case 4294967295:
                                                v15 := GenS1{F0: 71, F1: 984}
                                                _ = v15
                                            case 9223372036854775807:
                                                v16 := GenS1{F0: 127, F1: 2}
                                                _ = v16
                                            default:
                                                v4 = v4
                                            }
                                        }
                                    }
                                    _ = v12
                                }
                            default:
                                var v17 string = "hello"
                                _ = v17
                            }
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    } else {
                        v5 += (1e10 + ((v5 - float32(((1.5 * 1.5) + 0.0))) / 1.0))
                        if (32767 != 255) {
                            for fi18 := 0; fi18 < 20; fi18++ {
                                v19 := GenS1{F0: 42, F1: 636}
                                _ = fi18
                                _ = v19
                            }
                            var v20 float32 = v5
                            fmt.Println(v4, v4)
                            _ = v20
                        } else {
                            v5 = v5
                        }
                    }
                    _ = len(v3)
                    _ = fi7
                }
                if true {
                    if (v5 < v5) {
                        v1 = v1
                        v21 := GenS1{F0: 89, F1: 1}
                        v2 -= 32767
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 20; fi22++ {
                        fmt.Println(fi22, v5)
                        if true {
                            var v23 uint16 = 978
                            v24 := GenS1{F0: 42, F1: 2}
                            _ = v23
                            _ = v24
                        } else {
                            v5 = v5
                            v2 += v2
                            v4 = int64(v1)
                        }
                        if (0 >= 0) {
                            switch (fi22 * (fi22 ^ 127)) {
                            case -1:
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    fmt.Println(v4, v2, fi25)
                                    _ = fi25
                                }
                            default:
                                fmt.Println(v5, v4)
                            }
                        }
                        _ = fi22
                    }
                    v1 += v1
                } else {
                    v3[("" + "foo bar")] = (2147483647 == 634)
                    v26 := GenS3{F0: 378}
                    fmt.Println(v1, v2, v1)
                    _ = v26
                }
            }
        } else {
            var v27 string = (("" + "a") + "\n\t")
            if (false || (!true)) {
                if ((v4 % ((v4) | 1)) >= (v4 << (((v1 % ((v1) | 1))) & 63))) {
                    var v28 float64 = 1.5
                    _ = v28
                }
            }
            _ = v27
        }
        v29 := GenS3{F0: 255}
        _ = v29
    }
    if ((uint16(v1) == 0) && (v2 > 1)) {
        fmt.Println(v2)
        var v30 float32 = 3.14159
        fmt.Println(v2, v1, v2)
        _ = v30
    } else {
        for fi31 := 0; fi31 < 10; fi31++ {
            if true {
                switch fi31 {
                case 1:
                    var v32 uint = v1
                    _ = v32
                case 32767:
                    if false {
                        v1 = v1
                    }
                }
                v4 += 42
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 float32 = (1e10 + (1.5 + 0.1))
                _ = fi33
                _ = v34
            }
            v35 := v3["\n\t"]
            _ = fi31
            _ = v35
        }
        v36 := GenS4{F0: 702, F1: (^(v4)), F2: 0}
        v2 += v2
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        var v38 int = (fi37 ^ (-(fi37)))
        _ = fi37
        _ = v38
    }
    v3[("\n\t" + ("a" + ("foo bar" + (("a" + "\n\t") + ("a" + ("a" + "hello"))))))] = (1 >= uint16(v1))
    fmt.Println(v1)
    for fi39 := 0; fi39 < 5; fi39++ {
        switch -241 {
        case 0:
            v1 = v1
        case 9223372036854775807:
            if false {
                v3["foo bar"] = true
                var v40 bool = false
                _ = v40
            }
        default:
            var v41 string = "\n\t"
            _ = v41
        }
        var v42 uint8 = 0
        _ = fi39
        _ = v42
    }
    var v43 bool = (!true)
    var v44 int8 = 77
    v45 := GenS5{F0: (v43 && v43)}
    var v46 uint16 = 255
    var v47 string = "\n\t"
    switch 2147483647 {
    case 4294967295:
        v48 := []int{2147483647, 2147483647}
        _ = v48
    case -1:
        switch (("hello" + v47) + v47) {
        case "foo bar":
            v49 := GenS6{F0: v1, F1: int16(v46), F2: v2}
            _ = v49
        }
    default:
        v1 = v1
    }
    delete(v3, v47)
    if v43 {
        v50 := []int{-9223372036854775808}
        _ = v50
    }
    v51 := GenS7{F0: (v2 + v2), F1: (^(v46)), F2: v44}
    v3[(v47 + (("\n\t" + v47) + (v47 + "foo bar")))] = v43
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v51
}
