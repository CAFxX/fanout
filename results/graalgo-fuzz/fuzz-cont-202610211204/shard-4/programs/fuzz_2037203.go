package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 0
    v3 := []string{("a" + ""), "foo bar", "hello"}
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int16 = 749
            v2 = v2
            fmt.Println(v2, fi4, fi5)
            _ = fi5
            _ = v6
        }
        switch fi4 {
        case 2147483647:
            if true {
                switch 1 {
                case 0:
                    if true {
                        var v7 int64 = -639
                        _ = v7
                    } else {
                        if (677 >= -9223372036854775808) {
                            var v8 int = fi4
                            _ = v8
                        }
                    }
                case -9223372036854775808:
                    v2 -= v2
                default:
                    v2 = (v2 << ((0) & 63))
                }
                v2 = (v2 >> ((1) & 63))
            }
        case 9223372036854775807:
            if (true && ("a" > ("" + "hello"))) {
                if true {
                    var v9 uint = 4294967295
                    v10 := GenS1{F0: uint32(v9)}
                    _ = v10
                    _ = v9
                }
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := v1[173]
                    _ = fi11
                    _ = v12
                }
                var v13 int = (fi4 - fi4)
                _ = v13
            }
        case 214:
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := v1[0]
                v16 := GenS1{F0: 2}
                fmt.Println(v2)
                fmt.Println(v2, v15, fi4)
                v15 = (v15 - (v15 | v15))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v2 = v2
        _ = fi4
    }
    fmt.Println(v2)
    v2 += (v2 / (((v2 ^ (v2 - (v2 << ((70) & 63))))) | 1))
    for fi17 := 0; fi17 < 15; fi17++ {
        v18 := v1[42]
        _ = len(v3)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint = 4294967295
        _ = fi19
        _ = v20
    }
    if false {
        fmt.Println(v2)
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint8 = 0
            if ((((1.5 + ((1.5 * (0.1 - float32(0.1))) / float32(((-2.25 - (3.14159 - 0.0)) + (float64((-2.25 - (3.14159 * 1.5))) / (((float64(1.5) - 3.14159) * (3.14159 - 3.14159)) + 1.5)))))) - 3.14159) * 0.1) == (1e10 - 0.0)) {
                for fi23 := 0; fi23 < 7; fi23++ {
                    if true {
                        v24 := GenS1{F0: 2}
                        if false {
                            if false {
                                v25 := GenS1{F0: 275}
                                var v26 float64 = (0.1 * 1.5)
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS2{F0: (fi21 | fi23)}
                                _ = v27
                            }
                            if false {
                                fmt.Println(fi21, fi23)
                                var v28 int = (fi21 & 0)
                                fmt.Println(v22)
                                _ = v28
                            } else {
                                fmt.Println(fi23, v2)
                                if (-9223372036854775808 > 681) {
                                    v2 -= 1
                                }
                                if (32767 <= 32767) {
                                    _ = len(v3)
                                    v22 -= (v22 / (((v22 >> ((1) & 63))) | 1))
                                    if false {
                                        var v29 bool = false
                                        _ = v29
                                    } else {
                                        fmt.Println(fi21, fi23)
                                        v30 := GenS1{F0: uint32(v2)}
                                        v3 = append(v3, (("foo bar" + "hello") + "hello"))
                                        _ = v30
                                    }
                                }
                            }
                            fmt.Println(v2)
                        } else {
                            v31 := GenS3{F0: 42, F1: 0}
                            v2 = uint16(v22)
                            _ = v31
                        }
                        _ = v24
                    }
                    _ = fi23
                }
            }
            if (3.14159 <= 0.0) {
                v22 = 86
                v32 := GenS1{F0: 65535}
                _ = v32
            } else {
                var v33 string = ("\n\t" + "")
                for fi34 := 0; fi34 < 10; fi34++ {
                    fmt.Println(v2)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi34
                }
                var v35 bool = (1 >= 417)
                _ = v33
                _ = v35
            }
            _ = fi21
            _ = v22
        }
    }
    if true {
        var v36 float32 = (1e10 - float32(float64(0.1)))
        switch "foo bar" {
        case "a":
            v36 -= (float32((0.0 * float64((v36 + v36)))) + 3.14159)
        case "\n\t":
            v37 := GenS1{F0: 0}
            _ = v37
        default:
            for fi38 := 0; fi38 < 19; fi38++ {
                if true {
                    v2 += 65535
                    v3 = append(v3, "a")
                    v39 := GenS4{F0: (("hello" + "a") < "foo bar"), F1: (fi38 * (fi38 * (fi38 << ((175) & 63))))}
                    _ = v39
                } else {
                    v40 := GenS2{F0: fi38}
                    _ = v40
                }
                fmt.Println(v36, v2)
                var v41 int8 = -1
                _ = fi38
                _ = v41
            }
        }
        _ = v36
    } else {
        fmt.Println(v2)
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 float32 = (((0.1 - float32((0.0 / ((float64(((float32(3.14159) / (-2.25 - (0.0 - -2.25))) * 0.1)) / 1.0) - 0.1)))) - float32((((1e10 * 0.1) - 1e10) - 0.0))) + (3.14159 + float32((0.1 - 1.5))))
        var v44 uint64 = 18446744073709551615
        for fi45 := 0; fi45 < 19; fi45++ {
            v1[65535] = 52
            var v46 int = (fi45 * fi42)
            fmt.Println(v2)
            _ = fi45
            _ = v46
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println(v2)
    var v47 uint = 4294967295
    v47 = (v47 * v47)
    if true {
        for fi48 := 0; fi48 < 4; fi48++ {
            v2 += v2
            fmt.Println(v2, v47, v2)
            for fi49 := 0; fi49 < 1; fi49++ {
                var v50 int = fi49
                fmt.Println(v50)
                v47 = 65535
                var v51 int8 = 1
                if ((((4294967295 <= 562) || false) && false) && ((fi49 / ((446) | 1)) >= 975)) {
                    var v52 int16 = 378
                    var v53 string = "foo bar"
                    switch 32768 {
                    case 347:
                        v54 := GenS5{F0: (^(v52)), F1: 4294967295}
                        _ = v54
                    case -2147483648:
                        fmt.Println(v52)
                    }
                    _ = v52
                    _ = v53
                }
                v55 := GenS2{F0: fi48}
                _ = fi49
                _ = v50
                _ = v51
                _ = v55
            }
            _ = fi48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
}
