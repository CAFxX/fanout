package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }
type GenS6 struct { F0 int32; F1 uint; F2 int8 }
func genFn0(p0 uint, p1 uint8, p2 bool) uint16 {
    v6 := []int{-881, 228}
    for fi7 := 0; fi7 < 14; fi7++ {
        p1 -= 128
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}
func genFn1(p0 int8, p1 int, p2 int32) uint8 {
    v8 := GenS1{F0: (^(p1)), F1: 206, F2: 321}
    var v9 uint8 = 0
    fmt.Println(p2, p1, p1)
    for fi10 := 0; fi10 < 10; fi10++ {
        v9 = v9
        _ = fi10
    }
    if false {
        p0 = p0
        var v11 float32 = 0.1
        _ = v11
    } else {
        var v12 bool = (p0 < (^(p0)))
        var v13 int = (^(p1))
        for fi14 := 0; fi14 < 10; fi14++ {
            if v12 {
                var v15 int16 = -836
                p0 = (p0 | int8(p2))
                _ = v15
            } else {
                for fi16 := 0; fi16 < 10; fi16++ {
                    fmt.Println(v13, p2, v9)
                    switch fi14 {
                    case 9223372036854775807:
                        switch (fi14 ^ (v13 + (^(v13)))) {
                        case -1:
                            var v17 int16 = int16(p2)
                            _ = v17
                        }
                    case -893:
                        p0 = (p0 % (((p0 * (p0 / ((-1) | 1)))) | 1))
                    default:
                        v9 = (^(v9))
                    }
                    v18 := GenS2{F0: 42}
                    var v19 int64 = int64(p1)
                    _ = fi16
                    _ = v18
                    _ = v19
                }
                v20 := []int{32768, fi14, (v13 / ((fi14) | 1))}
                _ = v20
            }
            switch ("foo bar" + "a") {
            case "foo bar":
                v21 := []int{fi14, -9223372036854775808}
                _ = v21
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-2147483648, 2147483647}
    var v3 float64 = 0.1
    v4 := []uint32{2147483648, 1, 2147483647}
    v5 := make(map[string]int)
    for fi22 := 0; fi22 < 6; fi22++ {
        _ = len(v5)
        v3 -= (v3 - v3)
        _ = fi22
    }
    fmt.Println(v3, v3)
    switch 9223372036854775807 {
    case -572:
        v1 += (v1 << ((42) & 63))
    case 917:
        for fi23 := 0; fi23 < 16; fi23++ {
            v1 += (v1 + v1)
            _ = fi23
        }
    case 4294967295:
        if true {
            v24 := GenS3{F0: 379, F1: 2}
            if true {
                switch 128 {
                case 42:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 780:
                    v25 := GenS1{F0: 1, F1: 0, F2: 0}
                    _ = v25
                case 9223372036854775807:
                    var v26 int16 = -1
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 6; fi27++ {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            _ = len(v5)
                            if (1 >= 2) {
                                v1 += (v1 | v1)
                            }
                            _ = fi28
                        }
                        v29 := GenS3{F0: 255, F1: 2}
                        _ = fi27
                        _ = v29
                    }
                }
                var v30 uint16 = genFn0(65535, 2, (!false))
                _ = v30
            } else {
                var v31 uint16 = 515
                _ = v31
            }
            _ = v24
        }
    }
    v32 := GenS4{F0: 595}
    v33 := GenS3{F0: 0, F1: 871}
    v34 := GenS5{F0: 21, F1: 0, F2: ("foo bar" + "a")}
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 uint32 = 814
        var v37 int32 = -529
        for fi38 := 0; fi38 < 13; fi38++ {
            var v39 int32 = v37
            var v40 uint16 = 0
            v3 = v3
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v41 uint32 = 1
    if false {
        if true {
            v3 -= (1e10 - float64(1.5))
        } else {
            if true {
                var v42 uint8 = 1
                fmt.Println(v42)
                if ((!true) && (!(42 < -80))) {
                    v41 = (v41 & (^(v41)))
                }
                _ = v42
            }
            for fi43 := 0; fi43 < 19; fi43++ {
                if false {
                    if true {
                        _ = len(v4)
                        var v44 int64 = 42
                        _ = v44
                    }
                    _ = len(v4)
                    switch (fi43 % (((fi43 - fi43)) | 1)) {
                    case 32767:
                        v3 -= v3
                    case -210:
                        v1 = -1
                    case 127:
                        var v45 bool = false
                        _ = v45
                    default:
                        var v46 int8 = 127
                        _ = v46
                    }
                }
                var v47 string = ("hello" + (("\n\t" + ("" + (("foo bar" + ((("a" + "hello") + "hello") + "hello")) + "foo bar"))) + "\n\t"))
                _ = fi43
                _ = v47
            }
            v3 -= 1e10
        }
        switch 32768 {
        case 32768:
            for fi48 := 0; fi48 < 9; fi48++ {
                switch (fi48 | 987) {
                case -859:
                    if len(v2) > 0 {
                        v2[0] = fi48
                    }
                case 9223372036854775807:
                    if false {
                        switch "a" {
                        case "":
                            v3 -= v3
                        case "hello":
                            var v49 uint32 = 2147483648
                            _ = v49
                        }
                    }
                case 992:
                    v41 = v41
                default:
                    v3 = -2.25
                }
                _ = fi48
            }
        }
    }
    var v50 uint64 = uint64(v41)
    var v51 uint16 = 255
    v52 := GenS6{F0: -1, F1: 42, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
