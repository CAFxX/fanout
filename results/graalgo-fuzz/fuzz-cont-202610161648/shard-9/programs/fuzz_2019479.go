package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int32; F1 int8; F2 int }
type GenS4 struct { F0 int64; F1 uint16; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint }

func main() {
    var v1 uint = 1
    var v2 int = 9223372036854775807
    var v3 int = (^(v2))
    var v4 uint = v1
    fmt.Println(v1, v3, v3)
    if false {
        for fi5 := 0; fi5 < 14; fi5++ {
            for fi6 := 0; fi6 < 12; fi6++ {
                v3 = v3
                _ = fi6
            }
            _ = fi5
        }
        if true {
            if false {
                if true {
                    if false {
                        for fi7 := 0; fi7 < 12; fi7++ {
                            var v8 uint = v4
                            _ = fi7
                            _ = v8
                        }
                    }
                    var v9 float32 = 0.1
                    for fi10 := 0; fi10 < 20; fi10++ {
                        var v11 int32 = -580
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
                var v12 int8 = -128
                v3 -= (v3 % (((v3 * (v3 / (((v3 / ((330) | 1))) | 1)))) | 1))
                _ = v12
            }
            var v13 int64 = 415
            _ = v13
        } else {
            var v14 string = ((("foo bar" + "\n\t") + ("a" + "\n\t")) + ("a" + (("foo bar" + "a") + "a")))
            fmt.Println(v1)
            var v15 int64 = 947
            _ = v14
            _ = v15
        }
        v2 += v2
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v2 -= (v2 + 32768)
        _ = fi16
    }
    var v17 string = ((("hello" + (("" + "") + "")) + "hello") + "a")
    switch v17 {
    case "foo bar":
        var v18 int64 = 1
        _ = v18
    case "hello":
        fmt.Println(v1, v1)
    }
    fmt.Println(v1, v2, v4)
    switch -762 {
    case -2147483648:
        var v19 float32 = (0.0 * 1.5)
        _ = v19
    case 32768:
        if (3.14159 > 3.14159) {
            switch "" {
            case "hello":
                v20 := GenS1{F0: 42, F1: 114}
                _ = v20
            case "\n\t":
                switch (v2 - v3) {
                case 32768:
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v4 = (^(v4))
                        _ = fi21
                    }
                }
            }
            switch 128 {
            case 4294967295:
                for fi22 := 0; fi22 < 14; fi22++ {
                    switch 9223372036854775807 {
                    case 32767:
                        if true {
                            switch v2 {
                            case -136:
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    fmt.Println(v3, fi23)
                                    _ = fi23
                                }
                            case 127:
                                v24 := GenS1{F0: 2, F1: 65535}
                                _ = v24
                            }
                            v25 := []int{(v3 % ((v2) | 1))}
                            v17 = v17
                            _ = v25
                        } else {
                            var v26 bool = (true || true)
                            var v27 bool = v26
                            switch (v17 + "a") {
                            case "foo bar":
                                v28 := []int{}
                                _ = v28
                            case "\n\t":
                                v4 = v4
                            default:
                                v29 := GenS1{F0: 2, F1: 529}
                                _ = v29
                            }
                            _ = v26
                            _ = v27
                        }
                    case 128:
                        if (true || true) {
                            v4 -= 18446744073709551615
                            var v30 int8 = 42
                            switch -9223372036854775808 {
                            case 42:
                                v3 += (^(v2))
                            case -216:
                                switch (^(fi22)) {
                                case 0:
                                    v31 := []int{v2, fi22, fi22}
                                    _ = v31
                                case 128:
                                    for fi32 := 0; fi32 < 4; fi32++ {
                                        if false {
                                            v2 = (^(fi22))
                                        }
                                        if (!((1.5 - 0.0) == float64((((float32((1e10 * 3.14159)) + (((-2.25 * 1e10) - 3.14159) / 1e10)) * 0.1) * ((1.5 * 0.1) * (1.5 / 1.0)))))) {
                                            var v33 uint8 = 1
                                            v30 -= v30
                                            v34 := []int{fi22}
                                            _ = v33
                                            _ = v34
                                        }
                                        _ = fi32
                                    }
                                case 1:
                                    v35 := GenS2{F0: 422, F1: v4}
                                    _ = v35
                                }
                            case -2147483648:
                                v17 = "foo bar"
                            default:
                                switch ((v17 + v17) + "hello") {
                                case "foo bar":
                                    var v36 uint = 2
                                    _ = v36
                                case "\n\t":
                                    v17 = "hello"
                                }
                            }
                            _ = v30
                        }
                    case 127:
                        v1 = v1
                    default:
                        v3 = fi22
                    }
                    switch "foo bar" {
                    case "a":
                        for fi37 := 0; fi37 < 14; fi37++ {
                            switch (v17 + v17) {
                            case "":
                                if (2 < 2147483647) {
                                    v38 := []int{(v2 >> ((v1) & 63)), (v2 * (v3 + fi22)), (v2 ^ fi37), v3}
                                    _ = v38
                                } else {
                                    v39 := GenS3{F0: 242, F1: int8(v3), F2: v3}
                                    v40 := GenS3{F0: -1, F1: -128, F2: 169}
                                    switch v2 {
                                    case -407:
                                        v41 := GenS1{F0: 2, F1: 259}
                                        _ = v41
                                    default:
                                        if (-272 < 0) {
                                            v17 = "hello"
                                            var v42 string = v17
                                            var v43 bool = (false || (v17 != (v17 + "foo bar")))
                                            v44 := GenS4{F0: -1, F1: 0, F2: -1}
                                            var v45 uint = (v1 % ((v4) | 1))
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    _ = v39
                                    _ = v40
                                }
                            case "a":
                                v46 := []int{}
                                _ = v46
                            default:
                                v47 := []int{(v2 & (v3 * (^(v3)))), fi37, v2}
                                v1 += v1
                                _ = v47
                            }
                            v3 = 0
                            v3 = (^(v3))
                            _ = fi37
                        }
                    case "\n\t":
                        v48 := []int{fi22, fi22, 4294967295}
                        _ = v48
                    default:
                        v49 := []int{(fi22 / ((v3) | 1))}
                        _ = v49
                    }
                    _ = fi22
                }
            }
        } else {
            if (1 != 4294967295) {
                v3 += (v3 | 4294967295)
                var v50 bool = true
                v51 := GenS5{F0: 42}
                _ = v50
                _ = v51
            }
            v52 := GenS6{F0: 0, F1: v4}
            _ = v52
        }
    case 127:
        if false {
            var v53 int8 = -21
            var v54 float64 = 3.14159
            _ = v53
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        var v56 bool = (!(true && (false || false)))
        fmt.Println(v4, v1, v1)
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 1; fi57++ {
        v58 := []int{}
        _ = fi57
        _ = v58
    }
    var v59 uint64 = 4294967295
    v60 := []int{127, -725, v2, v2}
    var v61 string = ((v17 + v17) + "\n\t")
    for fi62 := 0; fi62 < 17; fi62++ {
        v2 += -1
        _ = fi62
    }
    fmt.Println(v4, v3, v4)
    var v63 int8 = 1
    if (89 < 573) {
        v60 = append(v60, -9223372036854775808)
        if len(v60) > 0 {
            v60[0] = 890
        }
    }
    v60 = append(v60, int(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v63
}
