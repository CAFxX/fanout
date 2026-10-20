package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint32; F2 int16 }
type GenS5 struct { F0 int; F1 int32; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int; F2 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 uint32 }

func main() {
    v1 := []int{32768, 2147483647, 2147483647}
    var v2 float64 = float64(3.14159)
    var v3 uint64 = 2
    var v4 uint32 = 2
    v5 := make(map[uint32]uint32)
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint8 = 127
        if true {
            fmt.Println(v4, v8)
            for fi9 := 0; fi9 < 4; fi9++ {
                v2 += v2
                var v10 int8 = 94
                var v11 uint = 106
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = fi7
        _ = v8
    }
    v6 -= v6
    v6 += (v6 / 1.0)
    switch 127 {
    case -1:
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                v13 := GenS1{F0: -994}
                var v14 bool = false
                _ = v13
                _ = v14
            } else {
                if (v2 == (v2 / 1.0)) {
                    if false {
                        v15 := GenS2{F0: 0, F1: 196, F2: -1}
                        var v16 int64 = 9223372036854775807
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v6 += ((0.0 + v6) + 0.1)
                            v18 := GenS3{F0: 42}
                            _ = fi17
                            _ = v18
                        }
                        _ = v15
                        _ = v16
                    } else {
                        var v19 bool = (((170 != (^(v3))) && false) && ((v4 ^ (^(v4))) == (v4 * v4)))
                        if (v19 && (!v19)) {
                            for fi20 := 0; fi20 < 14; fi20++ {
                                var v21 string = ""
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            if (-9223372036854775808 == 970) {
                                var v22 int64 = 0
                                if v19 {
                                    var v23 uint = 0
                                    _ = v23
                                } else {
                                    var v24 int64 = -368
                                    _ = v24
                                }
                                v25 := GenS1{F0: 42}
                                _ = v22
                                _ = v25
                            }
                        }
                        v3 = (v3 | v3)
                        _ = v19
                    }
                    v5[2147483648] = v4
                    v3 = (v3 << ((0) & 63))
                } else {
                    switch (("\n\t" + ("hello" + "")) + "hello") {
                    case "":
                        if (false && (535 > 0)) {
                            v2 -= float64(v6)
                            v1 = append(v1, -9223372036854775808)
                            if false {
                                if (float32(1e10) <= v6) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    fmt.Println(v3, v2, v3)
                                }
                                if false {
                                    switch (fi12 ^ (^(fi12))) {
                                    case 0:
                                        v26 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 42}
                                        _ = v26
                                    case 2147483647:
                                        if len(v1) > 0 {
                                            v1[0] = fi12
                                        }
                                    case 606:
                                        v2 = (v2 - float64(v6))
                                    default:
                                        var v27 int64 = 1
                                        _ = v27
                                    }
                                    v4 -= (^(v4))
                                }
                                for fi28 := 0; fi28 < 3; fi28++ {
                                    switch fi12 {
                                    case 42:
                                        var v29 uint64 = v3
                                        _ = v29
                                    case 127:
                                        switch (-(fi28)) {
                                        case 7:
                                            v6 -= float32(v2)
                                        }
                                    }
                                    v30 := GenS4{F0: -128, F1: (v4 % ((2147483647) | 1)), F2: -564}
                                    v3 = v3
                                    _ = fi28
                                    _ = v30
                                }
                            }
                        }
                    case "\n\t":
                        v3 = 1
                    default:
                        fmt.Println(v6, fi12, v3)
                    }
                    _ = len(v5)
                }
                var v31 uint8 = 1
                if (0 < int8(v4)) {
                    v3 = (v3 | 2)
                    v4 = (v4 - v4)
                    _ = len(v1)
                }
                _ = v31
            }
            _ = fi12
        }
    case -9223372036854775808:
        for fi32 := 0; fi32 < 18; fi32++ {
            fmt.Println(fi32, v3)
            fmt.Println(v6)
            var v33 bool = true
            _ = fi32
            _ = v33
        }
    case 4294967295:
        var v34 int = 4294967295
        _ = v34
    }
    v35 := GenS5{F0: -990, F1: int32(v4), F2: 255}
    v3 += 467
    var v36 int64 = 1
    var v37 bool = (true || (true || (false && (!(false || (-576 > 42))))))
    var v38 uint = 2
    fmt.Println(v37)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v39 := GenS3{F0: -1}
    v40 := v5[2147483647]
    switch "" {
    case "hello":
        switch "\n\t" {
        case "hello":
            v41 := GenS1{F0: 0}
            _ = v41
        }
    case "":
        for fi42 := 0; fi42 < 12; fi42++ {
            _ = len(v5)
            for fi43 := 0; fi43 < 2; fi43++ {
                _ = len(v1)
                _ = fi43
            }
            if v37 {
                var v44 int = (fi42 + 482)
                var v45 bool = true
                for fi46 := 0; fi46 < 12; fi46++ {
                    var v47 float64 = v2
                    switch 4294967295 {
                    case 0:
                        if true {
                            var v48 uint32 = v4
                            switch "hello" {
                            case "\n\t":
                                v3 = (v3 >> ((4294967295) & 63))
                            case "":
                                v49 := GenS6{F0: "hello"}
                                _ = v49
                            default:
                                var v50 int16 = 148
                                if false {
                                    if v37 {
                                        fmt.Println(v40, v6)
                                        v6 -= v6
                                        fmt.Println(fi42, v50, v36)
                                    } else {
                                        delete(v5, (v4 << (((^(v38))) & 63)))
                                        v2 -= 1e10
                                    }
                                    if false {
                                        var v51 uint64 = (v3 + (v3 ^ (v3 | (v3 >> (((^(v38))) & 63)))))
                                        if v45 {
                                            v1 = append(v1, (v44 - fi46))
                                            _ = len(v5)
                                        }
                                        _ = v51
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 = v3
                                    }
                                    v3 = v3
                                    if ((-128 == 127) && v45) {
                                        if (!v45) {
                                            switch "foo bar" {
                                            case "a":
                                                v44 += (^(fi46))
                                            case "":
                                                var v52 float32 = ((0.0 - (v6 / v6)) / ((float32(v47) - 3.14159) * v6))
                                                var v53 string = ("" + ("a" + ""))
                                                v54 := GenS3{F0: 2147483647}
                                                v50 -= -262
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            case "hello":
                                                switch "" {
                                                case "":
                                                    v37 = (!(true || (v45 && (v37 && v45))))
                                                case "\n\t":
                                                    v2 = (float64((-2.25 - (0.1 - 1e10))) + 0.1)
                                                }
                                                v55 := GenS5{F0: fi42, F1: -472, F2: 2}
                                                _ = v55
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = v44
                                                }
                                            }
                                            var v56 uint = v38
                                            v57 := GenS1{F0: v50}
                                            v40 = 52
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 uint32 = v4
                                        _ = v58
                                    } else {
                                        fmt.Println(fi42, v36, v2)
                                        v50 += int16(v36)
                                    }
                                }
                                var v59 uint64 = (v3 * v3)
                                _ = v50
                                _ = v59
                            }
                            v60 := GenS1{F0: 680}
                            _ = v48
                            _ = v60
                        } else {
                            v61 := v5[(v4 % ((v4) | 1))]
                            _ = v61
                        }
                    case 978:
                        v62 := GenS7{F0: 727, F1: (^(fi46)), F2: 0}
                        _ = v62
                    case 4294967295:
                        v2 += float64(1e10)
                    }
                    switch int(v36) {
                    case 127:
                        var v63 string = "a"
                        _ = v63
                    case 0:
                        var v64 uint16 = 2
                        _ = v64
                    default:
                        switch (fi42 & fi42) {
                        case 9223372036854775807:
                            if v45 {
                                v65 := GenS8{F0: v36}
                                _ = v65
                            }
                        case 42:
                            v66 := GenS9{F0: 255}
                            _ = v66
                        default:
                            v67 := GenS4{F0: 100, F1: v4, F2: 42}
                            _ = v67
                        }
                    }
                    _ = fi46
                    _ = v47
                }
                _ = v44
                _ = v45
            }
            _ = fi42
        }
    case "foo bar":
        v68 := GenS10{F0: -1, F1: v4}
        _ = v68
    default:
        v3 -= v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
