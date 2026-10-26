package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int64, p1 uint16, p2 string) int16 {
    if true {
        var v7 string = ((p2 + p2) + "a")
        v8 := []int{1, 32767, 9223372036854775807, 9223372036854775807}
        _ = v7
        _ = v8
    }
    switch ((p2 + p2) + "a") {
    case "foo bar":
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = 2147483647
            _ = fi9
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 9; fi11++ {
            if (p1 > p1) {
                p2 = p2
            }
            var v12 int8 = 42
            _ = fi11
            _ = v12
        }
    }
    switch 1 {
    case -2147483648:
        var v13 int64 = (^(p0))
        _ = v13
    case 32768:
        if ((2147483647 < 2147483647) || true) {
            if (3.14159 < 1e10) {
                v14 := GenS1{F0: -385}
                switch -563 {
                case 0:
                    p0 = (-(p0))
                }
                _ = v14
            } else {
                p2 = (p2 + p2)
            }
        }
    }
    switch "" {
    case "foo bar":
        v15 := GenS2{F0: p2}
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 18; fi16++ {
            p2 = (((p2 + ((("a" + "a") + ("a" + p2)) + "a")) + ("a" + "foo bar")) + "hello")
            v17 := []int{4294967295, (fi16 >> ((42) & 63)), fi16, (fi16 | 0)}
            _ = fi16
            _ = v17
        }
    case "hello":
        for fi18 := 0; fi18 < 12; fi18++ {
            switch fi18 {
            case 2147483647:
                p0 -= -1
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS2{F0: "hello"}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []uint32{965}
    var v2 int32 = -2147483648
    v3 := []int{32768, 0, 379, 906}
    var v4 float32 = (0.0 - float32((0.1 - 3.14159)))
    v5 := []uint32{611, 2147483648, 618}
    var v6 uint32 = 4294967295
    var v21 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = 911
    }
    v6 += 4294967295
    if false {
        for fi22 := 0; fi22 < 3; fi22++ {
            if true {
                fmt.Println(v2, v4, v21)
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 uint16 = 2
                    v3 = append(v3, 545)
                    switch 32768 {
                    case 32768:
                        switch ("" + "a") {
                        case "hello":
                            var v25 uint32 = 4294967295
                            _ = v25
                        default:
                            v26 := GenS2{F0: "foo bar"}
                            _ = v26
                        }
                    case -9223372036854775808:
                        var v27 uint = 255
                        _ = v27
                    }
                    v28 := GenS3{F0: 42}
                    switch "foo bar" {
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "hello":
                        var v29 uint8 = 152
                        _ = v29
                    default:
                        if (1 <= 1) {
                            var v30 uint = 386
                            switch "a" {
                            case "a":
                                if (!(2 >= 42)) {
                                    var v31 uint32 = v6
                                    if false {
                                        var v32 int8 = 127
                                        v33 := GenS1{F0: 1}
                                        _ = v32
                                        _ = v33
                                    } else {
                                        switch (-(fi22)) {
                                        case 2147483647:
                                            switch "\n\t" {
                                            case "a":
                                                if (true && ("\n\t" != "\n\t")) {
                                                    var v34 uint = (v30 | v30)
                                                    v35 := GenS2{F0: (((("foo bar" + "a") + "foo bar") + "\n\t") + "\n\t")}
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            default:
                                                v4 += (v4 / (3.14159 - v4))
                                            }
                                        case 658:
                                            fmt.Println(v4, fi22)
                                        case 0:
                                            var v36 float32 = -2.25
                                            _ = v36
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                    _ = v31
                                }
                            }
                            var v37 uint8 = 255
                            _ = v30
                            _ = v37
                        }
                        var v38 string = ("hello" + ("" + "\n\t"))
                        v24 = (v24 ^ (v24 >> ((4294967295) & 63)))
                        _ = v38
                    }
                    _ = fi23
                    _ = v24
                    _ = v28
                }
            } else {
                var v39 uint64 = 2
                switch fi22 {
                case -699:
                    var v40 float32 = v4
                    _ = v40
                }
                v39 += 1
                _ = v39
            }
            v4 = float32(float64((v4 - v4)))
            _ = fi22
        }
    }
    fmt.Println(v4)
    v4 = v4
    var v41 float32 = v4
    if false {
        if false {
            if false {
                v5 = append(v5, (v6 ^ v6))
            }
            for fi42 := 0; fi42 < 1; fi42++ {
                v41 = (0.0 / v4)
                var v43 uint8 = 0
                v41 += (v41 / ((3.14159 - -2.25) / (v4 - (v4 / v41))))
                _ = fi42
                _ = v43
            }
        }
        v44 := GenS4{F0: 2, F1: 100, F2: 2}
        if true {
            if true {
                v5 = append(v5, (v6 ^ (v6 | v6)))
            } else {
                fmt.Println(v6)
                v21 = ((-2.25 / v21) / float64(v41))
            }
            var v45 uint16 = 0
            _ = v45
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v44
    }
    var v46 uint16 = 41
    _ = len(v5)
    for fi47 := 0; fi47 < 2; fi47++ {
        if false {
            var v48 uint16 = 287
            switch "hello" {
            case "":
                v48 -= (^(v46))
            }
            v2 -= (^(v2))
            _ = v48
        } else {
            fmt.Println(v2)
            v21 += v21
            v4 = 1e10
        }
        v49 := GenS5{F0: 42, F1: v6}
        _ = fi47
        _ = v49
    }
    if false {
        if len(v5) > 0 {
            v5[0] = (^(v6))
        }
    } else {
        if (438 <= 739) {
            for fi50 := 0; fi50 < 1; fi50++ {
                if len(v3) > 0 {
                    v3[0] = fi50
                }
                _ = fi50
            }
            v2 -= (v2 << ((65535) & 63))
            if true {
                var v51 int8 = 1
                if len(v1) > 0 {
                    v1[0] = v6
                }
                var v52 float64 = (v21 * v21)
                _ = v51
                _ = v52
            } else {
                if len(v1) > 0 {
                    v1[0] = (^(v6))
                }
                for fi53 := 0; fi53 < 20; fi53++ {
                    v54 := GenS1{F0: (v2 | (-(v2)))}
                    _ = fi53
                    _ = v54
                }
            }
        } else {
            for fi55 := 0; fi55 < 15; fi55++ {
                v56 := GenS1{F0: v2}
                switch fi55 {
                case 0:
                    var v57 uint = 144
                    _ = v57
                case 916:
                    v6 = 0
                }
                v4 -= (v4 * (float32(v21) - (((0.1 * v41) * float32((1e10 - v21))) * 0.0)))
                _ = fi55
                _ = v56
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ("\n\t" + "a") {
    case "a":
        if true {
            for fi58 := 0; fi58 < 15; fi58++ {
                fmt.Println(v46)
                if true {
                    v59 := GenS6{F0: -9223372036854775808}
                    _ = v59
                }
                _ = fi58
            }
            for fi60 := 0; fi60 < 14; fi60++ {
                v2 += v2
                _ = fi60
            }
            v46 -= v46
        } else {
            v61 := GenS4{F0: 65535, F1: 100, F2: v6}
            v62 := GenS3{F0: -128}
            _ = v61
            _ = v62
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v6
}
