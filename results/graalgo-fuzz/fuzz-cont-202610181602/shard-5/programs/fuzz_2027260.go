package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1.5
    var v3 bool = (1 == 42)
    v4 := make(map[int]uint32)
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 2; fi6++ {
        v2 -= v2
        _ = fi6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v2 += v2
        v8 := GenS1{F0: 1, F1: "foo bar"}
        v9 := []int{fi7, fi7, (fi7 * fi7), -2147483648}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "a" {
    case "":
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                v2 -= v2
                v3 = true
                if v3 {
                    if (1.5 <= v5) {
                        var v12 uint32 = 2147483647
                        var v13 int16 = 617
                        var v14 bool = (int16(fi11) > (v13 ^ -32768))
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        v15 := v1[42]
                        if v15 {
                            fmt.Println(fi10, fi10)
                            fmt.Println(v15)
                            var v16 bool = false
                            _ = v16
                        } else {
                            v1[469] = v3
                            v17 := v4[fi10]
                            var v18 int = (fi10 >> ((1) & 63))
                            v18 += (fi10 / (((^(fi10))) | 1))
                            _ = v17
                            _ = v18
                        }
                        _ = v15
                    }
                }
                _ = fi11
            }
            if (366 < 42) {
                var v19 bool = false
                v20 := GenS1{F0: 1, F1: "hello"}
                switch ("" + "a") {
                case "hello":
                    if (((0.1 - -2.25) + v2) >= 0.1) {
                        switch "hello" {
                        case "\n\t":
                            for fi21 := 0; fi21 < 5; fi21++ {
                                var v22 string = "hello"
                                _ = fi21
                                _ = v22
                            }
                        }
                    }
                case "\n\t":
                    for fi23 := 0; fi23 < 10; fi23++ {
                        if true {
                            if v3 {
                                _ = len(v4)
                                v24 := []int{0, -95, -892, fi23}
                                v19 = (((fi23 * (fi10 ^ fi23)) == (fi23 << ((157) & 63))) && (42 == 127))
                                v25 := GenS1{F0: 1, F1: "\n\t"}
                                _ = v24
                                _ = v25
                            }
                            var v26 uint8 = 157
                            _ = v26
                        } else {
                            v2 += v2
                            fmt.Println(fi10, v3)
                        }
                        if (!false) {
                            v27 := GenS1{F0: 0, F1: "\n\t"}
                            _ = v27
                        }
                        if v3 {
                            if (false && v19) {
                                var v28 uint16 = 65535
                                v19 = v3
                                v5 -= (v5 - (((0.0 * 0.0) - float64((v2 - float32((1.5 + v5))))) * v5))
                                _ = v28
                            } else {
                                v2 = (v2 + v2)
                                v19 = (65535 >= 440)
                            }
                            var v29 uint16 = 32767
                            var v30 uint8 = 127
                            _ = v29
                            _ = v30
                        }
                        _ = fi23
                    }
                }
                _ = v19
                _ = v20
            } else {
                var v31 int8 = -1
                if (((((3.14159 / 0.1) * (v2 / (((float32(v5) - v2) + (v2 / ((v2 + v2) + (v2 + v2)))) - v2))) + v2) / v2) > v2) {
                    v5 = (0.0 - v5)
                    var v32 int64 = 814
                    _ = v32
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    _ = len(v1)
                    _ = fi33
                }
                _ = v31
            }
            _ = fi10
        }
    case "\n\t":
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 uint16 = 32767
            if v3 {
                switch "hello" {
                case "foo bar":
                    var v36 bool = v3
                    _ = v36
                default:
                    var v37 uint16 = v35
                    _ = v37
                }
                v38 := GenS1{F0: 65535, F1: "a"}
                for fi39 := 0; fi39 < 11; fi39++ {
                    v40 := GenS1{F0: 42, F1: ("" + "\n\t")}
                    if (42 != fi34) {
                        v5 = (float64((0.1 / 1e10)) / ((v5 + v5) + float64((((3.14159 - (-2.25 * (0.1 * float32(0.1)))) / v2) - 3.14159))))
                        switch "\n\t" {
                        case "hello":
                            v35 += 861
                        case "foo bar":
                            v41 := GenS1{F0: uint32(v35), F1: ("" + ("a" + "foo bar"))}
                            _ = v41
                        case "":
                            fmt.Println(fi39)
                        }
                        var v42 int64 = -704
                        _ = v42
                    }
                    var v43 int16 = 32767
                    if v3 {
                        var v44 uint8 = 2
                        var v45 int32 = -2147483648
                        delete(v4, 0)
                        _ = len(v4)
                        _ = v44
                        _ = v45
                    }
                    _ = fi39
                    _ = v40
                    _ = v43
                }
                _ = v38
            }
            _ = fi34
            _ = v35
        }
    default:
        v46 := GenS2{F0: ("a" + "foo bar"), F1: 1}
        _ = v46
    }
    v5 += v5
    var v47 float64 = float64((1e10 + v2))
    if ("\n\t" > "hello") {
        var v48 uint64 = 2
        _ = v48
    } else {
        var v49 float64 = (0.1 - (v47 + 3.14159))
        if (100 != 127) {
            if (0 == -1) {
                v50 := GenS2{F0: "\n\t", F1: 2147483648}
                if v3 {
                    v51 := []int{}
                    _ = v51
                } else {
                    var v52 uint16 = 470
                    _ = v52
                }
                _ = v50
            }
            var v53 float64 = 1e10
            v54 := GenS2{F0: "foo bar", F1: 42}
            _ = v53
            _ = v54
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            v47 += v5
            _ = fi55
        }
        _ = v49
    }
    v56 := GenS3{F0: (466 > 473), F1: 811}
    switch "\n\t" {
    case "hello":
        fmt.Println(v47, v47)
    default:
        fmt.Println(v3)
    }
    var v57 int16 = -754
    v58 := GenS2{F0: ("foo bar" + "\n\t"), F1: 42}
    var v59 int64 = 299
    switch "" {
    case "\n\t":
        v47 = v5
    case "":
        for fi60 := 0; fi60 < 9; fi60++ {
            var v61 uint64 = 0
            var v62 uint64 = v61
            _ = fi60
            _ = v61
            _ = v62
        }
    case "hello":
        v57 -= v57
    }
    var v63 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v63
}
