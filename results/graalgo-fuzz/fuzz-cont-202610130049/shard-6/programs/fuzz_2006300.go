package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 uint16) uint8 {
    v5 := []int{42, 32767}
    var v6 uint = 255
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 int32) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p1, p0, p1)
        v8 := GenS1{F0: 4294967295, F1: 1, F2: 32767}
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 3; fi10++ {
        _ = len(v9)
        var v11 string = ("" + ("hello" + "hello"))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            if (false || (42 < 18446744073709551615)) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 bool = (p0 < p0)
                    v14 = (!v14)
                    var v15 uint16 = 781
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                if true {
                    p0 += p0
                    v16 := GenS1{F0: 1, F1: -1, F2: 2}
                    _ = v16
                } else {
                    v17 := GenS1{F0: 2, F1: -1, F2: 255}
                    p0 = (p0 << ((410) & 63))
                    switch "\n\t" {
                    case "a":
                        if (1 == 2147483648) {
                            var v18 float32 = 3.14159
                            if true {
                                v19 := GenS1{F0: 2, F1: 29, F2: 468}
                                p0 = p0
                                v18 += ((float32(3.14159) / 3.14159) / (1.5 / float32(3.14159)))
                                _ = v19
                            }
                            switch 2147483647 {
                            case 32767:
                                switch (fi12 + fi12) {
                                case 4294967295:
                                    var v20 string = ("" + "a")
                                    _ = v20
                                default:
                                    fmt.Println(v18)
                                }
                            case -2147483648:
                                v21 := GenS2{F0: false, F1: 234, F2: 2}
                                _ = v21
                            }
                            _ = v18
                        } else {
                            p0 += -966
                        }
                    case "":
                        v22 := GenS3{F0: ("a" + ("hello" + "a")), F1: 2, F2: 0}
                        _ = v22
                    case "hello":
                        fmt.Println(p0, fi12, p1)
                    default:
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := GenS3{F0: "a", F1: 574, F2: 18446744073709551615}
                            p1 = 1
                            fmt.Println(fi23)
                            if true {
                                if (42 > 42) {
                                    var v25 uint64 = 42
                                    _ = len(v9)
                                    _ = len(v9)
                                    _ = v25
                                }
                                if false {
                                    if (!(603 != 1)) {
                                        var v26 bool = (true && (false && false))
                                        _ = v26
                                    } else {
                                        _ = len(v9)
                                        var v27 int32 = int32(p0)
                                        _ = v27
                                    }
                                    var v28 uint16 = 790
                                    var v29 string = (((("a" + "foo bar") + "\n\t") + ("" + "\n\t")) + "a")
                                    _ = v28
                                    _ = v29
                                } else {
                                    if len(v9) > 0 {
                                        v9[0] = (fi12 * fi12)
                                    }
                                    p0 -= (p0 & p0)
                                }
                            }
                            p1 = (p1 | p1)
                            _ = fi23
                            _ = v24
                        }
                    }
                    _ = v17
                }
            } else {
                var v30 bool = false
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS4{F0: 2147483647, F1: (^(p1))}
            if true {
                v33 := GenS4{F0: 970, F1: p1}
                _ = v33
            } else {
                v9 = append(v9, (^(fi12)))
            }
            _ = fi31
            _ = v32
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(p0)
            if false {
                _ = len(v9)
                var v35 float64 = (1.5 + 1e10)
                switch "a" {
                case "\n\t":
                    v36 := GenS5{F0: 987, F1: 18446744073709551615, F2: 128}
                    _ = v36
                case "foo bar":
                    var v37 string = ""
                    _ = v37
                default:
                    fmt.Println(fi12, p0, p1)
                }
                _ = v35
            }
            _ = fi34
        }
        _ = fi12
    }
    var v38 bool = false
    _ = p0
    _ = p1
    _ = v38
    _ = v9
    return -120
}

func main() {
    var v1 int32 = 0
    var v2 uint = 65535
    v3 := make(map[int]uint32)
    var v4 float32 = float32(0.0)
    v4 += (v4 / v4)
    v39 := GenS3{F0: "a", F1: 42, F2: 4294967295}
    v2 += (v2 - v2)
    _ = len(v3)
    var v40 int = 32767
    v41 := GenS2{F0: false, F1: v1, F2: 127}
    v42 := GenS2{F0: (!(!true)), F1: (v1 >> (((v2 | 2)) & 63)), F2: genFn0(974)}
    switch "foo bar" {
    case "foo bar":
        for fi43 := 0; fi43 < 20; fi43++ {
            switch ("\n\t" + (("hello" + "\n\t") + "hello")) {
            case "hello":
                v44 := []int{32767}
                _ = v44
            case "a":
                for fi45 := 0; fi45 < 4; fi45++ {
                    var v46 uint32 = 532
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v1)
                    case "hello":
                        v3[42] = v46
                    case "a":
                        var v47 int8 = int8(fi45)
                        _ = v47
                    }
                    v48 := []int{}
                    _ = fi45
                    _ = v46
                    _ = v48
                }
            default:
                switch 1 {
                case -1:
                    v49 := GenS5{F0: 0, F1: 65535, F2: genFn0(176)}
                    _ = v49
                default:
                    v40 = v40
                }
            }
            if ("\n\t" > ("" + "a")) {
                var v50 uint32 = 374
                for fi51 := 0; fi51 < 12; fi51++ {
                    v1 += v1
                    v50 += v50
                    _ = fi51
                }
                _ = v50
            } else {
                switch ("" + "\n\t") {
                case "\n\t":
                    var v52 int64 = 11
                    _ = v52
                case "hello":
                    var v53 uint16 = 1
                    _ = v53
                case "foo bar":
                    v4 -= float32((1e10 * 0.1))
                }
                v54 := GenS5{F0: 860, F1: v2, F2: 127}
                _ = v54
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi43, v2, v1)
            case "":
                v40 = v40
            case "a":
                var v55 int16 = 32767
                _ = v55
            default:
                var v56 float64 = -2.25
                _ = v56
            }
            _ = fi43
        }
    case "hello":
        v1 = v1
    case "\n\t":
        var v57 int64 = 228
        _ = v57
    default:
        var v58 int16 = 1
        _ = v58
    }
    fmt.Println(v4, v40)
    if false {
        if (false || true) {
            var v59 float64 = 0.0
            var v60 uint = 2
            v3[32767] = 1
            _ = v59
            _ = v60
        } else {
            var v61 uint64 = 1
            _ = v61
        }
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
