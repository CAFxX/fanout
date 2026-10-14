package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 int; F1 uint }
func genFn0() bool {
    var v7 uint16 = 42
    if false {
        var v8 uint16 = 2
        v9 := []int{4294967295}
        _ = v8
        _ = v9
    } else {
        var v10 uint = 0
        _ = v10
    }
    var v11 float32 = 3.14159
    var v12 float32 = 3.14159
    if (!false) {
        if false {
            v13 := GenS1{F0: (v7 - 42), F1: 1, F2: 4294967295}
            var v14 string = (("\n\t" + "foo bar") + "")
            switch 128 {
            case 128:
                v15 := []int{32768, 442, 1}
                _ = v15
            case 32767:
                v16 := []int{1, 9223372036854775807, 127, -620}
                _ = v16
            case 32768:
                v11 -= v11
            }
            _ = v13
            _ = v14
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := []int{(fi17 & 0)}
                if (-126 == -13) {
                    var v19 uint = 1
                    if len(v18) > 0 {
                        v18[0] = -9223372036854775808
                    }
                    if ((fi17 + fi17) >= (fi17 - (fi17 & (^(fi17))))) {
                        if true {
                            fmt.Println(v7)
                            v12 = 1e10
                            fmt.Println(v12, v12, v7)
                        }
                        if false {
                            if (!(87 == 127)) {
                                if false {
                                    if false {
                                        switch (("a" + "") + "a") {
                                        case "":
                                            v11 += float32((3.14159 * (1.5 - 0.1)))
                                        }
                                        v20 := GenS2{F0: (v19 * 0), F1: 0}
                                        _ = v20
                                    }
                                } else {
                                    var v21 int64 = 9223372036854775807
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        fmt.Println(v19, v11, fi22)
                                        var v23 uint32 = 2147483648
                                        fmt.Println(fi17, fi22, v11)
                                        _ = fi22
                                        _ = v23
                                    }
                                    _ = v21
                                }
                                v7 = (v7 >> ((255) & 63))
                                var v24 string = "foo bar"
                                _ = v24
                            }
                            var v25 uint8 = 2
                            var v26 int16 = -1
                            _ = v25
                            _ = v26
                        } else {
                            var v27 int16 = -42
                            if false {
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v11 -= v11
                                    _ = fi28
                                }
                                _ = len(v18)
                                var v29 int8 = 1
                                _ = v29
                            } else {
                                v30 := GenS3{F0: 42, F1: 71}
                                _ = v30
                            }
                            _ = v27
                        }
                    } else {
                        v31 := GenS4{F0: 551, F1: 11, F2: -1}
                        fmt.Println(fi17, v19, v19)
                        for fi32 := 0; fi32 < 1; fi32++ {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            v18 = append(v18, (fi32 | fi32))
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            var v33 bool = (!true)
                            var v34 uint8 = 99
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    }
                    _ = v19
                } else {
                    var v35 int64 = -9223372036854775808
                    switch "" {
                    case "hello":
                        var v36 string = ("\n\t" + "foo bar")
                        _ = v36
                    default:
                        v18 = append(v18, fi17)
                    }
                    _ = v35
                }
                _ = fi17
                _ = v18
            }
        }
    }
    _ = v11
    _ = v12
    _ = v7
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    v3 := []int{}
    var v4 string = "foo bar"
    v5 := []int{-452}
    v6 := []uint32{42, 130}
    v4 = "a"
    var v37 int64 = 0
    v2 = (v2 % (((v2 * 0)) | 1))
    for fi38 := 0; fi38 < 19; fi38++ {
        if (!true) {
            var v39 int64 = (v37 - v37)
            fmt.Println(v2, v4)
            for fi40 := 0; fi40 < 15; fi40++ {
                switch v4 {
                case "":
                    var v41 int = fi38
                    _ = v41
                }
                _ = fi40
            }
            _ = v39
        } else {
            v2 -= 32767
            if ((v37 * (^(v37))) != v37) {
                if false {
                    v2 += v2
                    v2 += (v2 ^ (v2 - (v2 / ((v2) | 1))))
                }
            } else {
                v2 += v2
                var v42 int64 = (^(v37))
                var v43 uint16 = 65535
                _ = v42
                _ = v43
            }
            var v44 int = int(v2)
            _ = v44
        }
        var v45 int64 = (v37 ^ v37)
        for fi46 := 0; fi46 < 1; fi46++ {
            v47 := GenS5{F0: fi38, F1: 65535}
            fmt.Println(v37)
            _ = fi46
            _ = v47
        }
        _ = fi38
        _ = v45
    }
    switch v4 {
    case "a":
        for fi48 := 0; fi48 < 13; fi48++ {
            var v49 bool = true
            var v50 uint8 = 33
            for fi51 := 0; fi51 < 7; fi51++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v52 uint16 = 65535
                _ = fi51
                _ = v52
            }
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    v37 -= v37
    for fi53 := 0; fi53 < 17; fi53++ {
        var v54 uint8 = 11
        _ = len(v5)
        switch 530 {
        case 1:
            if (!(true || true)) {
                v55 := GenS1{F0: 186, F1: (v37 - -1), F2: 994}
                var v56 float64 = (((1e10 - 3.14159) - (((1.5 + 1e10) * 0.0) * 0.0)) * 1e10)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v55
                _ = v56
            } else {
                var v57 uint = 0
                switch int(v37) {
                case -606:
                    v1[(v4 + (v4 + ""))] = false
                }
                _ = v57
            }
        case -1:
            var v58 float64 = -2.25
            _ = v58
        default:
            v59 := GenS5{F0: fi53, F1: 779}
            _ = v59
        }
        _ = fi53
        _ = v54
    }
    var v60 float64 = 1e10
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
