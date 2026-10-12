package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 int32; F2 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{""}
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int8 = 58
        fmt.Println(fi6, v7)
        switch fi6 {
        case 118:
            v8 := GenS1{F0: 4294967295, F1: -401, F2: 42}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if (true && false) {
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: 928, F1: 0, F2: 1}
                _ = fi10
                _ = v11
            }
            var v12 float32 = float32(1.5)
            v12 = 0.0
            _ = fi9
            _ = v12
        }
        var v13 float64 = (3.14159 * 0.1)
        _ = v13
    } else {
        v14 := GenS2{F0: 13, F1: 1, F2: 0}
        _ = v14
    }
    var v15 int16 = 0
    v16 := GenS2{F0: 127, F1: -805, F2: 783}
    if (!false) {
        switch 4294967295 {
        case 215:
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 11; fi18++ {
                    if false {
                        var v19 uint32 = 703
                        var v20 int = 0
                        _ = len(v2)
                        switch "a" {
                        case "\n\t":
                            _ = len(v3)
                        case "foo bar":
                            if (100 != 100) {
                                v19 = v19
                                var v21 uint16 = 9
                                var v22 uint = 2
                                if ((v19 >= 0) && (float64(1e10) >= (float64(0.0) + float64(1e10)))) {
                                    var v23 uint = (v22 & v22)
                                    _ = v23
                                }
                                var v24 float64 = float64(-2.25)
                                _ = v21
                                _ = v22
                                _ = v24
                            } else {
                                var v25 float32 = 0.0
                                _ = v25
                            }
                        default:
                            var v26 bool = true
                            _ = v26
                        }
                        v20 -= (v20 - (v20 % ((127) | 1)))
                        v27 := v1[v20]
                        _ = v19
                        _ = v20
                        _ = v27
                    } else {
                        fmt.Println(v15)
                        if (!true) {
                            switch ("" + "foo bar") {
                            case "foo bar":
                                var v28 bool = (v15 != -863)
                                _ = v28
                            case "a":
                                if false {
                                    fmt.Println(fi18, v15)
                                    v29 := GenS1{F0: 98, F1: 541, F2: (v15 % ((v15) | 1))}
                                    if ((42 != 290) && (0 != 42)) {
                                        var v30 int64 = -88
                                        _ = v30
                                    }
                                    _ = v29
                                }
                            case "":
                                v31 := GenS2{F0: 42, F1: 2147483647, F2: (fi18 << ((1) & 63))}
                                _ = v31
                            }
                            fmt.Println(fi18, fi17)
                            v15 -= v15
                        }
                        if ((0.1 * -2.25) > (0.1 + (float32((-2.25 * 1e10)) - (float32(1e10) + (0.0 - (float32(1.5) / 1.5)))))) {
                            fmt.Println(v15)
                            if true {
                                var v32 uint32 = 42
                                if (-342 < 0) {
                                    if false {
                                        fmt.Println(fi18, fi17)
                                        v33 := GenS3{F0: false, F1: 42}
                                        _ = v33
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v32)
                                } else {
                                    if (!false) {
                                        if false {
                                            v34 := GenS1{F0: 18446744073709551615, F1: -462, F2: 687}
                                            _ = v34
                                        } else {
                                            if false {
                                                fmt.Println(fi18, fi18)
                                                if false {
                                                    var v35 string = ""
                                                    _ = v35
                                                } else {
                                                    var v36 int16 = (v15 * -32768)
                                                    _ = v36
                                                }
                                                var v37 string = "a"
                                                _ = v37
                                            }
                                            fmt.Println(fi18, fi18)
                                            var v38 int64 = 9223372036854775807
                                            _ = v38
                                        }
                                    } else {
                                        if true {
                                            v32 = v32
                                        }
                                        v32 += (v32 / ((v32) | 1))
                                        var v39 uint32 = 65535
                                        _ = v39
                                    }
                                    v32 -= (v32 | v32)
                                    var v40 uint32 = v32
                                    _ = v40
                                }
                                var v41 int = (^(fi18))
                                _ = v32
                                _ = v41
                            }
                            v15 += -32768
                        }
                    }
                    if false {
                        v15 = v15
                        var v42 int16 = int16(fi18)
                        var v43 bool = (!true)
                        _ = v42
                        _ = v43
                    } else {
                        v15 = v15
                    }
                    fmt.Println(fi17, v15, fi18)
                    switch 149 {
                    case 4294967295:
                        var v44 bool = false
                        _ = v44
                    case 2147483647:
                        var v45 bool = (!false)
                        _ = v45
                    default:
                        _ = len(v3)
                    }
                    var v46 int64 = 9223372036854775807
                    _ = fi18
                    _ = v46
                }
                v15 = v15
                _ = fi17
            }
        default:
            switch -81 {
            case -1:
                if true {
                    v15 -= v15
                    for fi47 := 0; fi47 < 17; fi47++ {
                        fmt.Println(v15)
                        v15 += v15
                        _ = fi47
                    }
                    v15 += 1000
                }
            default:
                v48 := GenS4{F0: v15, F1: 42}
                _ = v48
            }
        }
    } else {
        var v49 int32 = -1
        _ = v49
    }
    if true {
        v15 -= v15
        for fi50 := 0; fi50 < 7; fi50++ {
            if false {
                delete(v3, fi50)
                _ = len(v2)
            } else {
                for fi51 := 0; fi51 < 9; fi51++ {
                    v15 += (v15 / (((^(v15))) | 1))
                    var v52 uint32 = 1
                    if (0 > 1) {
                        var v53 string = "a"
                        _ = v53
                    }
                    _ = fi51
                    _ = v52
                }
            }
            fmt.Println(fi50)
            fmt.Println(v15, v15)
            _ = fi50
        }
        var v54 string = (("\n\t" + ("a" + "a")) + "")
        _ = v54
    }
    v55 := GenS5{F0: 2147483647, F1: v15}
    v15 += v15
    var v56 uint64 = 2
    for fi57 := 0; fi57 < 20; fi57++ {
        var v58 int16 = -1
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        var v60 int16 = v15
        v61 := v1[(fi59 | fi59)]
        var v62 bool = true
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    if true {
        var v63 bool = (true || true)
        v64 := GenS1{F0: (v56 ^ 999), F1: 1, F2: (v15 >> ((0) & 63))}
        var v65 int64 = 9223372036854775807
        _ = v63
        _ = v64
        _ = v65
    }
    if true {
        v15 = v15
        if false {
            var v66 int32 = -2147483648
            _ = v66
        }
    }
    v15 -= v15
    v56 = (v56 << ((513) & 63))
    v67 := GenS2{F0: 181, F1: 0, F2: 42}
    var v68 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v67
    _ = v68
}
