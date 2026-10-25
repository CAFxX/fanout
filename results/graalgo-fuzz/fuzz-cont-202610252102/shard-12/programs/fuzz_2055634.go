package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 bool; F1 int32 }
func genFn0(p0 int64, p1 bool, p2 int64) int16 {
    fmt.Println(p0, p1, p2)
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := []int{-164}
        var v6 bool = ((-(p2)) > p0)
        var v7 float64 = (1.5 + 0.0)
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{453, -1}
    p0 = (p2 - (p2 ^ (p2 << ((2) & 63))))
    switch "foo bar" {
    case "hello":
        if len(v8) > 0 {
            v8[0] = 128
        }
    case "a":
        var v9 uint8 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1000
}

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    v3 := []uint32{4294967295, 0, 2, 4294967295}
    switch 0 {
    case -9223372036854775808:
        v1 = (!false)
    case 2147483647:
        _ = len(v2)
    case -23:
        if v1 {
            var v10 int16 = 1
            _ = v10
        } else {
            var v11 float64 = 1e10
            _ = v11
        }
    default:
        switch 938 {
        case 32768:
            var v12 uint = 255
            _ = v12
        default:
            _ = len(v3)
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if (0 <= 42) {
            var v14 float32 = -2.25
            fmt.Println(v14, fi13, fi13)
            if false {
                for fi15 := 0; fi15 < 11; fi15++ {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    var v16 int32 = 559
                    var v17 uint16 = 255
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    if v1 {
                        v14 = -2.25
                    } else {
                        var v19 int32 = -890
                        v19 += v19
                        fmt.Println(fi18)
                        v1 = false
                        switch "foo bar" {
                        case "":
                            v20 := GenS1{F0: -1}
                            _ = v20
                        }
                        _ = v19
                    }
                    _ = fi18
                }
            }
            _ = v14
        }
        switch 4294967295 {
        case -1:
            v1 = (852 < 13)
        default:
            switch fi13 {
            case 32768:
                switch fi13 {
                case 127:
                    v21 := GenS1{F0: 212}
                    _ = v21
                case -284:
                    switch "hello" {
                    case "":
                        v22 := GenS2{F0: 697}
                        _ = v22
                    default:
                        var v23 uint = 565
                        _ = v23
                    }
                case -2147483648:
                    for fi24 := 0; fi24 < 13; fi24++ {
                        if v1 {
                            if v1 {
                                if (3.14159 < 1.5) {
                                    v2[("foo bar" + "hello")] = ((42 >= -1) || v1)
                                } else {
                                    var v25 float64 = 1e10
                                    if false {
                                        v1 = false
                                        fmt.Println(fi13, v25)
                                        if false {
                                            var v26 float32 = 1.5
                                            v26 = 3.14159
                                            _ = v26
                                        } else {
                                            if (!false) {
                                                var v27 float64 = v25
                                                v1 = (v25 <= v27)
                                                v28 := GenS1{F0: int32(fi13)}
                                                if ((!(103 == 1)) || v1) {
                                                    if len(v3) > 0 {
                                                        v3[0] = 325
                                                    }
                                                } else {
                                                    v1 = false
                                                }
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                    }
                                    _ = v25
                                }
                            } else {
                                if v1 {
                                    v29 := GenS3{F0: 0, F1: 42, F2: 990}
                                    if v1 {
                                        v1 = (!(42 < 0))
                                    } else {
                                        v1 = (0 >= 161)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 699
                                    }
                                    _ = v29
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = 670
                                        }
                                    case "a":
                                        v1 = ((!(733 > 42)) || v1)
                                    }
                                    v1 = (true && v1)
                                    fmt.Println(fi24, fi24, v1)
                                }
                                v1 = v1
                                var v30 int = fi24
                                _ = v30
                            }
                        }
                        v31 := GenS2{F0: 42}
                        _ = fi24
                        _ = v31
                    }
                }
            case 1:
                var v32 uint64 = 1
                _ = v32
            case 128:
                if (-2.25 >= (3.14159 / float32(-2.25))) {
                    fmt.Println(fi13, v1)
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 uint32 = 0
                        v34 += (^(v34))
                        _ = fi33
                        _ = v34
                    }
                }
            default:
                fmt.Println(v1, v1)
            }
        }
        for fi35 := 0; fi35 < 17; fi35++ {
            fmt.Println(fi35)
            _ = fi35
        }
        _ = fi13
    }
    if false {
        v3 = append(v3, 712)
        switch 95 {
        case 128:
            v1 = v1
        case 4294967295:
            if true {
                switch "\n\t" {
                case "\n\t":
                    v2["hello"] = v1
                }
                v36 := v2["a"]
                _ = v36
            } else {
                for fi37 := 0; fi37 < 19; fi37++ {
                    var v38 uint32 = 2147483648
                    v38 += 454
                    fmt.Println(fi37, v1)
                    _ = fi37
                    _ = v38
                }
            }
        case -302:
            for fi39 := 0; fi39 < 6; fi39++ {
                var v40 uint32 = 832
                _ = len(v3)
                for fi41 := 0; fi41 < 12; fi41++ {
                    fmt.Println(v40)
                    _ = fi41
                }
                _ = fi39
                _ = v40
            }
        }
    } else {
        v2["foo bar"] = false
        var v42 string = "\n\t"
        _ = len(v3)
        _ = v42
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 float64 = 3.14159
        switch (fi43 | fi43) {
        case 1:
            fmt.Println(fi43, v1)
        case 2147483647:
            var v45 uint = 65535
            _ = v45
        case 107:
            if (!(1000 >= 42)) {
                var v46 uint = 1
                v1 = (!false)
                _ = v46
            } else {
                switch fi43 {
                case -404:
                    v44 -= v44
                case 42:
                    if false {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            var v48 uint16 = 381
                            v49 := GenS4{F0: fi47}
                            var v50 int8 = 42
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    } else {
                        if (v44 < 1e10) {
                            for fi51 := 0; fi51 < 19; fi51++ {
                                v52 := GenS2{F0: 255}
                                _ = fi51
                                _ = v52
                            }
                        } else {
                            for fi53 := 0; fi53 < 7; fi53++ {
                                v54 := GenS5{F0: "a", F1: 127}
                                v44 += ((1e10 - 0.1) * v44)
                                _ = fi53
                                _ = v54
                            }
                            v1 = v1
                        }
                    }
                case 2147483647:
                    var v55 int64 = 0
                    _ = v55
                }
                v56 := GenS5{F0: "", F1: 52}
                _ = v56
            }
        default:
            var v57 uint = 709
            _ = v57
        }
        _ = fi43
        _ = v44
    }
    fmt.Println(v1)
    var v58 uint16 = 255
    v3 = append(v3, 65535)
    v59 := GenS6{F0: false, F1: 817}
    var v60 uint8 = 1
    var v61 string = "hello"
    v62 := GenS6{F0: v1, F1: -2147483648}
    var v63 int16 = 0
    v63 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
