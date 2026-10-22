package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 bool; F2 bool }
type GenS7 struct { F0 int; F1 uint16; F2 uint8 }
type GenS8 struct { F0 string; F1 uint8; F2 uint16 }
type GenS9 struct { F0 bool; F1 int32 }
func genFn0(p0 int, p1 uint64, p2 uint32) uint32 {
    if (!true) {
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := []int{(fi5 % (((p0 ^ (fi5 << ((75) & 63)))) | 1)), -2147483648}
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: false, F1: 81, F2: -32768}
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        var v9 string = "a"
        _ = v9
    }
    v10 := GenS1{F0: ((^(p2)) <= p2), F1: 344, F2: 1}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 bool = true
    for fi11 := 0; fi11 < 1; fi11++ {
        v1["\n\t"] = genFn0(0, 42, 1)
        _ = fi11
    }
    if (float32(1.5) < 1e10) {
        switch 0 {
        case -9223372036854775808:
            v12 := GenS1{F0: v4, F1: 1, F2: 1}
            _ = v12
        }
        v4 = (42 <= 255)
    } else {
        if (0.0 <= (1.5 - -2.25)) {
            if v4 {
                v13 := GenS1{F0: true, F1: 42, F2: -32768}
                v14 := []int{591}
                for fi15 := 0; fi15 < 12; fi15++ {
                    if ("\n\t" >= "hello") {
                        fmt.Println(fi15, v4)
                        fmt.Println(fi15, v4)
                    } else {
                        if (42 != 32767) {
                            for fi16 := 0; fi16 < 2; fi16++ {
                                v17 := GenS1{F0: ((true || false) && (255 == 1)), F1: 65535, F2: 42}
                                v4 = (42 < 127)
                                _ = fi16
                                _ = v17
                            }
                            fmt.Println(v4, v4)
                            if ("foo bar" < "foo bar") {
                                var v18 uint64 = 860
                                _ = v18
                            }
                        } else {
                            v4 = (2 < 1)
                            var v19 int8 = 0
                            var v20 bool = v4
                            _ = v19
                            _ = v20
                        }
                    }
                    if (-9223372036854775808 != 1) {
                        v4 = v4
                        v4 = v4
                    } else {
                        fmt.Println(v4, fi15)
                        v4 = (!(fi15 < fi15))
                        var v21 int64 = 1
                        _ = v21
                    }
                    _ = fi15
                }
                _ = v13
                _ = v14
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    switch "" {
                    case "":
                        var v24 uint16 = 65535
                        _ = v24
                    default:
                        v25 := []int{-2147483648, fi22, fi22}
                        _ = v25
                    }
                    _ = fi23
                }
                _ = fi22
            }
            _ = len(v3)
        } else {
            if true {
                v26 := []int{127, 800}
                v27 := GenS1{F0: true, F1: 716, F2: 28}
                fmt.Println(v4)
                _ = v26
                _ = v27
            }
        }
        var v28 uint64 = 2
        _ = v28
    }
    switch -9223372036854775808 {
    case 1:
        v29 := GenS2{F0: 127, F1: 115}
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        delete(v3, 2)
        switch (("\n\t" + ("" + ("\n\t" + "foo bar"))) + "\n\t") {
        case "hello":
            if v4 {
                v31 := GenS1{F0: v4, F1: 368, F2: 1000}
                if ("" > ("" + ((("a" + "\n\t") + "") + ("" + "hello")))) {
                    var v32 int64 = 0
                    var v33 float32 = ((1e10 / 1e10) * 1.5)
                    _ = v32
                    _ = v33
                }
                _ = v31
            }
        case "\n\t":
            v34 := GenS1{F0: (v4 && (42 > int8(fi30))), F1: 255, F2: -1}
            _ = v34
        case "a":
            v4 = (!false)
        default:
            v4 = (v4 || v4)
        }
        _ = fi30
    }
    v35 := GenS3{F0: (2 < 2), F1: 1}
    v36 := GenS4{F0: 255, F1: 139, F2: (0 <= 86)}
    if v4 {
        v4 = v4
        var v37 uint8 = 1
        var v38 int16 = 42
        _ = v37
        _ = v38
    } else {
        _ = len(v3)
        v39 := GenS4{F0: 0, F1: 255, F2: true}
        var v40 int8 = 127
        _ = v39
        _ = v40
    }
    switch -9223372036854775808 {
    case 1:
        v41 := []int{1, 1, -29, 1}
        _ = v41
    case -764:
        var v42 uint16 = 255
        _ = v42
    case 2147483647:
        v43 := []int{0, -987, -1}
        _ = v43
    }
    v44 := GenS5{F0: -674}
    var v45 float64 = 1.5
    switch "hello" {
    case "\n\t":
        var v46 int8 = -1
        _ = v46
    case "a":
        v47 := GenS6{F0: (0.1 == float64(3.14159)), F1: v4, F2: (!(-115 != 9223372036854775807))}
        _ = v47
    case "":
        var v48 uint = 424
        _ = v48
    default:
        switch 0 {
        case -1:
            if v4 {
                v49 := GenS7{F0: 32768, F1: 2, F2: 255}
                if v4 {
                    v4 = true
                } else {
                    fmt.Println(v45)
                    if v4 {
                        v50 := GenS7{F0: 32767, F1: 850, F2: 112}
                        var v51 uint8 = 128
                        if (0 <= -456) {
                            if (true || (127 <= -1)) {
                                var v52 uint8 = 70
                                v45 -= v45
                                fmt.Println(v52)
                                _ = v52
                            } else {
                                for fi53 := 0; fi53 < 7; fi53++ {
                                    if true {
                                        v45 = -2.25
                                        switch (fi53 - 9223372036854775807) {
                                        case 4294967295:
                                            for fi54 := 0; fi54 < 18; fi54++ {
                                                if (235 != 1) {
                                                    v2[(fi53 + fi53)] = "foo bar"
                                                }
                                                var v55 uint16 = 65535
                                                _ = fi54
                                                _ = v55
                                            }
                                        default:
                                            v56 := v2[(fi53 >> ((829) & 63))]
                                            _ = v56
                                        }
                                        v57 := []int{(fi53 % ((fi53) | 1))}
                                        _ = v57
                                    }
                                    _ = fi53
                                }
                            }
                            fmt.Println(v4, v4)
                            v4 = false
                        }
                        _ = v50
                        _ = v51
                    }
                    v58 := GenS8{F0: ("" + ("\n\t" + (("foo bar" + ("" + "")) + ""))), F1: 22, F2: 144}
                    _ = v58
                }
                _ = v49
            }
        }
    }
    v4 = v4
    v2[-201] = "foo bar"
    v59 := GenS9{F0: v4, F1: 1}
    if (true && v4) {
        var v60 uint64 = 42
        var v61 int = -9223372036854775808
        v4 = v4
        _ = v60
        _ = v61
    }
    if v4 {
        v62 := GenS7{F0: 1, F1: 480, F2: 127}
        v63 := GenS6{F0: (v4 || false), F1: v4, F2: (!(!(v4 && true)))}
        _ = v62
        _ = v63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v45
    _ = v59
}
