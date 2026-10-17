package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint }
func genFn0() int {
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: (fi8 % (((fi8 ^ (fi8 | (fi8 - fi8)))) | 1)), F1: true, F2: "\n\t"}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        switch -9223372036854775808 {
        case 4294967295:
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := GenS1{F0: 1, F1: (true && false), F2: "\n\t"}
                v13 := []int{(fi11 << ((269) & 63))}
                v7 = v7
                _ = fi11
                _ = v12
                _ = v13
            }
        case -2147483648:
            v14 := []int{32768, fi10, fi10}
            _ = v14
        default:
            if false {
                switch (v7 + v7) {
                case "a":
                    v15 := []int{fi10}
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 1; fi16++ {
                        if ((318 >= 32767) && true) {
                            v17 := []int{}
                            var v18 bool = true
                            var v19 int16 = 1
                            switch ("foo bar" + ("hello" + ((v7 + ((v7 + v7) + v7)) + v7))) {
                            case "":
                                var v20 uint16 = 515
                                _ = v20
                            case "\n\t":
                                if (v18 && v18) {
                                    if v18 {
                                        var v21 int8 = 42
                                        _ = v21
                                    } else {
                                        if (!v18) {
                                            fmt.Println(fi16, v19)
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            if (!v18) {
                                                v22 := GenS1{F0: (fi10 - (fi10 >> ((8) & 63))), F1: true, F2: "hello"}
                                                fmt.Println(v18, fi16)
                                                v23 := GenS2{F0: (fi10 ^ -2147483648), F1: 4294967295}
                                                _ = v22
                                                _ = v23
                                            }
                                        default:
                                            v7 = (v7 + v7)
                                        }
                                    }
                                } else {
                                    v24 := GenS1{F0: (fi10 & -1), F1: (802 >= 42), F2: v7}
                                    _ = v24
                                }
                            }
                            _ = v17
                            _ = v18
                            _ = v19
                        } else {
                            v7 = "a"
                            if false {
                                var v25 int32 = 1
                                if false {
                                    if (249 != 255) {
                                        v26 := GenS1{F0: (fi16 & (fi10 * (fi10 / ((127) | 1)))), F1: false, F2: v7}
                                        _ = v26
                                    } else {
                                        fmt.Println(v25, fi10)
                                        var v27 bool = false
                                        if false {
                                            var v28 float32 = ((3.14159 * 1.5) - 1e10)
                                            _ = v28
                                        } else {
                                            v29 := GenS3{F0: 0}
                                            var v30 uint8 = 128
                                            var v31 uint8 = (v30 & (v30 * v30))
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v27
                                    }
                                    var v32 int64 = 1
                                    _ = v32
                                }
                                _ = v25
                            } else {
                                var v33 int = (fi10 << ((528) & 63))
                                _ = v33
                            }
                        }
                        var v34 uint16 = 75
                        _ = fi16
                        _ = v34
                    }
                default:
                    switch 274 {
                    case 9223372036854775807:
                        v7 = "a"
                    default:
                        v7 = v7
                    }
                }
                var v35 uint32 = 0
                v36 := GenS4{F0: 42, F1: (fi10 > (fi10 & fi10))}
                _ = v35
                _ = v36
            } else {
                v37 := []int{(fi10 & 127), fi10, fi10, 128}
                var v38 float64 = 1e10
                var v39 string = ((v7 + v7) + v7)
                _ = v37
                _ = v38
                _ = v39
            }
        }
        var v40 int = fi10
        _ = fi10
        _ = v40
    }
    v41 := []int{32768, 32768, -644}
    _ = v41
    _ = v7
    return -637
}

func main() {
    var v1 float32 = ((0.1 - ((0.1 - 1e10) / float32((((0.1 / 3.14159) * (0.1 + 1e10)) / 1.0)))) * 0.1)
    v2 := make(map[string]bool)
    var v3 float32 = float32(1e10)
    v4 := make(map[uint32]int)
    var v5 int8 = -72
    v6 := make(map[int]bool)
    var v42 float32 = (v1 - v3)
    v42 -= float32(1e10)
    v43 := GenS2{F0: 42, F1: 4294967295}
    v44 := GenS4{F0: 0, F1: (!(0.1 <= 1e10))}
    var v45 bool = (!(-1 >= 0))
    v46 := GenS3{F0: 29}
    v1 = (float32((1e10 / 0.1)) + (0.1 - v3))
    if (65535 >= 2) {
        v42 = 1e10
        v5 = (v5 ^ 1)
    } else {
        fmt.Println(v1, v3)
        var v47 int16 = 817
        _ = v47
    }
    v48 := []int{genFn0(), 4294967295, -9223372036854775808}
    for fi49 := 0; fi49 < 5; fi49++ {
        var v50 string = ("foo bar" + ("hello" + ("a" + "")))
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        for fi52 := 0; fi52 < 15; fi52++ {
            if (0.0 <= float64((1e10 / v42))) {
                fmt.Println(v42)
            } else {
                fmt.Println(v3, fi52, v5)
                _ = len(v48)
            }
            switch "" {
            case "\n\t":
                v45 = v45
            case "":
                v42 = 0.0
            case "a":
                if (v45 || false) {
                    fmt.Println(fi51, fi52)
                    v53 := GenS5{F0: 640}
                    _ = v53
                } else {
                    v2["\n\t"] = (42 < 8)
                    if len(v48) > 0 {
                        v48[0] = (fi52 - fi51)
                    }
                    delete(v4, 2147483647)
                }
            }
            _ = fi52
        }
        _ = fi51
    }
    v54 := GenS3{F0: 451}
    for fi55 := 0; fi55 < 8; fi55++ {
        v45 = v45
        _ = fi55
    }
    var v56 int = 9223372036854775807
    var v57 int32 = -2147483648
    if true {
        var v58 string = ("foo bar" + "foo bar")
        v58 = v58
        v42 += -2.25
        _ = v58
    } else {
        fmt.Println(v45, v42, v3)
        v42 -= float32((3.14159 * 0.1))
    }
    switch genFn0() {
    case 0:
        switch genFn0() {
        case 966:
            if (!true) {
                v56 = (^(v56))
                for fi59 := 0; fi59 < 2; fi59++ {
                    if (-9223372036854775808 > -34) {
                        v56 = 128
                    } else {
                        var v60 int8 = 100
                        _ = v60
                    }
                    _ = fi59
                }
            } else {
                var v61 uint = 255
                _ = v61
            }
        default:
            v48 = append(v48, v56)
        }
    default:
        v3 = 0.1
    }
    var v62 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v54
    _ = v56
    _ = v57
    _ = v6
    _ = v62
}
