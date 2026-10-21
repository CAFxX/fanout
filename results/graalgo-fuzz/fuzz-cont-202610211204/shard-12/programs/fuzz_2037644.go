package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int8 }
func genFn0() bool {
    var v6 bool = (1 <= 42)
    var v7 uint = 856
    fmt.Println(v7)
    v7 = (v7 | (v7 << (((^(v7))) & 63)))
    _ = v6
    _ = v7
    return (874 < 975)
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -1
    var v3 uint64 = 42
    var v4 int64 = -9223372036854775808
    var v5 int64 = (v4 / ((9223372036854775807) | 1))
    v2 = 41
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 string = ""
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 192, F1: 0, F2: 409}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    switch 127 {
    case 2147483647:
        v12 := GenS1{F0: 255, F1: 1, F2: 377}
        _ = v12
    case 516:
        v3 -= v3
    }
    if true {
        switch 128 {
        case 991:
            v13 := GenS2{F0: 148, F1: 1, F2: -2147483648}
            _ = v13
        }
    }
    if true {
        if true {
            fmt.Println(v4)
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v3)
                v4 += -557
                _ = fi14
            }
        }
        v2 = v2
    } else {
        if true {
            var v15 uint8 = 2
            fmt.Println(v4, v5)
            if true {
                v16 := GenS2{F0: v15, F1: v15, F2: -9223372036854775808}
                v5 = v4
                v17 := v1["hello"]
                _ = v16
                _ = v17
            } else {
                v15 -= (v15 - 127)
            }
            _ = v15
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                fmt.Println(v5, v3)
                switch (fi18 + (fi18 & (fi18 + fi18))) {
                case 1:
                    v19 := GenS3{F0: 747, F1: 2147483647}
                    _ = v19
                case 2147483647:
                    switch ("a" + "a") {
                    case "\n\t":
                        if true {
                            v4 -= -432
                            var v20 uint = 1
                            _ = v20
                        }
                    case "hello":
                        for fi21 := 0; fi21 < 9; fi21++ {
                            var v22 uint64 = 18446744073709551615
                            var v23 bool = (false && false)
                            switch 234 {
                            case 128:
                                var v24 int16 = -1
                                _ = v24
                            case -9223372036854775808:
                                if v23 {
                                    v25 := GenS3{F0: 42, F1: 1}
                                    _ = v25
                                }
                            default:
                                var v26 float64 = 1e10
                                _ = v26
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    default:
                        if true {
                            v5 = v4
                            switch fi18 {
                            case -664:
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    v28 := GenS4{F0: -415, F1: (2 != (v3 - v3)), F2: false}
                                    var v29 uint16 = 0
                                    var v30 uint = 65535
                                    v31 := GenS5{F0: v4, F1: (v30 & v30)}
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            case 4294967295:
                                v32 := []int{-135, -608}
                                _ = v32
                            }
                            v2 = (v2 / ((v2) | 1))
                        }
                    }
                case 32767:
                    for fi33 := 0; fi33 < 4; fi33++ {
                        if true {
                            v3 -= 18446744073709551615
                        } else {
                            v34 := []int{fi18}
                            if true {
                                _ = len(v1)
                                var v35 uint8 = 2
                                v36 := GenS3{F0: 2, F1: 818}
                                var v37 float32 = -2.25
                                v38 := GenS6{F0: v2, F1: false}
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            v39 := GenS2{F0: 255, F1: 1, F2: (fi18 << ((42) & 63))}
                            _ = v34
                            _ = v39
                        }
                        _ = fi33
                    }
                }
                var v40 int16 = -634
                _ = fi18
                _ = v40
            }
            for fi41 := 0; fi41 < 3; fi41++ {
                v42 := GenS1{F0: 42, F1: 2147483647, F2: 65535}
                _ = fi41
                _ = v42
            }
        }
        var v43 int32 = 42
        v5 = (v4 << ((18446744073709551615) & 63))
        _ = v43
    }
    fmt.Println(v4, v4, v2)
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 float64 = (-2.25 * 1e10)
        v5 -= (^(v5))
        _ = fi44
        _ = v45
    }
    v46 := GenS2{F0: 42, F1: 255, F2: 32767}
    var v47 uint8 = 255
    switch 129 {
    case -2147483648:
        var v48 uint8 = v47
        _ = v48
    case 127:
        fmt.Println(v2, v47)
    case 32768:
        var v49 int = -750
        _ = v49
    }
    v1["foo bar"] = ""
    var v50 uint32 = 592
    fmt.Println(v2, v4, v50)
    for fi51 := 0; fi51 < 8; fi51++ {
        var v52 int8 = v2
        _ = fi51
        _ = v52
    }
    switch -321 {
    case 9223372036854775807:
        if (!false) {
            switch "a" {
            case "hello":
                _ = len(v1)
            case "\n\t":
                var v53 uint16 = 65535
                _ = v53
            case "":
                fmt.Println(v3, v2, v47)
            }
            var v54 int8 = (-(v2))
            v55 := GenS3{F0: 255, F1: 640}
            _ = v54
            _ = v55
        }
    default:
        for fi56 := 0; fi56 < 4; fi56++ {
            fmt.Println(v50)
            if (v47 != 212) {
                for fi57 := 0; fi57 < 9; fi57++ {
                    fmt.Println(v47, v47, v50)
                    v58 := GenS3{F0: 42, F1: 42}
                    v59 := GenS7{F0: v2}
                    var v60 uint16 = 661
                    v61 := GenS2{F0: v47, F1: (^(v47)), F2: (-(fi57))}
                    _ = fi57
                    _ = v58
                    _ = v59
                    _ = v60
                    _ = v61
                }
            }
            for fi62 := 0; fi62 < 13; fi62++ {
                v63 := []int{int(v47), 127, fi62, (fi62 | -2147483648)}
                switch ("foo bar" + "hello") {
                case "hello":
                    v4 += v4
                case "":
                    v2 += (v2 << ((18446744073709551615) & 63))
                default:
                    if (false || false) {
                        fmt.Println(fi62, v2)
                    } else {
                        v1["foo bar"] = "hello"
                        var v64 bool = ((v47 - v47) < v47)
                        v1[""] = ("\n\t" + "")
                        _ = v64
                    }
                }
                fmt.Println(v2)
                _ = fi62
                _ = v63
            }
            _ = fi56
        }
    }
    v5 = -1
    if (true && (255 == 1)) {
        v65 := GenS5{F0: (v4 - v4), F1: 435}
        if false {
            v47 = 42
        } else {
            v4 -= (v5 | v4)
        }
        _ = v65
    } else {
        for fi66 := 0; fi66 < 17; fi66++ {
            for fi67 := 0; fi67 < 8; fi67++ {
                delete(v1, "\n\t")
                fmt.Println(v4, fi67)
                v4 = v4
                _ = fi67
            }
            _ = fi66
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
