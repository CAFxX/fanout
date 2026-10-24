package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint; F2 int16 }

func main() {
    var v1 uint16 = 32767
    var v2 int16 = -233
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    var v5 uint16 = 524
    switch 9223372036854775807 {
    case -24:
        v6 := GenS1{F0: 0}
        _ = v6
    case 42:
        v7 := []int{32768, -9223372036854775808, -1, 0}
        _ = v7
    default:
        v8 := GenS1{F0: v2}
        _ = v8
    }
    v2 = -675
    v1 = v1
    v9 := GenS1{F0: v2}
    v1 = (v1 - 255)
    var v10 uint64 = 4294967295
    delete(v4, "a")
    if true {
        v5 = v5
        v1 = v5
        v5 -= (v1 + 175)
    } else {
        v1 = v5
    }
    var v11 float64 = (1.5 - 1.5)
    v12 := GenS1{F0: (^(v2))}
    v2 = v2
    v5 -= v1
    if false {
        var v13 int64 = 0
        if true {
            if (0 < uint8(v1)) {
                if (1 <= (v13 << ((482) & 63))) {
                    if (v10 < (v10 / ((242) | 1))) {
                        if true {
                            v14 := GenS2{F0: 1}
                            switch ("a" + "\n\t") {
                            case "\n\t":
                                _ = len(v4)
                            case "foo bar":
                                for fi15 := 0; fi15 < 4; fi15++ {
                                    if false {
                                        var v16 uint16 = v5
                                        v17 := []int{(fi15 ^ fi15), fi15, (fi15 & fi15), 4294967295}
                                        v1 += uint16(v10)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v10 = (^(v10))
                                        if true {
                                            var v18 uint32 = 2147483648
                                            if ((3.14159 <= v11) || (!true)) {
                                                var v19 uint8 = 127
                                                var v20 uint64 = (v10 % ((v10) | 1))
                                                _ = v19
                                                _ = v20
                                            }
                                            _ = v18
                                        }
                                        for fi21 := 0; fi21 < 12; fi21++ {
                                            var v22 float64 = (1.5 + ((v11 - (((v11 * v11) * 3.14159) - v11)) + (float64(((1.5 / -2.25) - 1.5)) * (v11 - v11))))
                                            if ((("hello" + "a") + "") > "\n\t") {
                                                v11 -= 0.0
                                            } else {
                                                v23 := GenS3{F0: (v2 - 442)}
                                                v10 -= 1
                                                v22 = (float64(3.14159) * 0.1)
                                                v24 := []int{}
                                                _ = v23
                                                _ = v24
                                            }
                                            v25 := []int{}
                                            _ = fi21
                                            _ = v22
                                            _ = v25
                                        }
                                    }
                                    switch 4294967295 {
                                    case -9223372036854775808:
                                        v5 += (v1 % (((^(v5))) | 1))
                                    case -2147483648:
                                        switch "hello" {
                                        case "":
                                            v26 := GenS3{F0: (v2 % (((-(v2))) | 1))}
                                            _ = v26
                                        default:
                                            for fi27 := 0; fi27 < 2; fi27++ {
                                                switch "\n\t" {
                                                case "a":
                                                    v4["hello"] = true
                                                    switch "\n\t" {
                                                    case "":
                                                        var v28 uint32 = 0
                                                        _ = v28
                                                    default:
                                                        var v29 int32 = -1
                                                        _ = v29
                                                    }
                                                    v30 := GenS4{F0: -1}
                                                    _ = v30
                                                }
                                                v5 = 32767
                                                _ = fi27
                                            }
                                        }
                                    default:
                                        var v31 uint32 = 65535
                                        _ = v31
                                    }
                                    _ = fi15
                                }
                            case "":
                                if false {
                                    var v32 string = "hello"
                                    _ = v32
                                } else {
                                    if (4294967295 != 0) {
                                        fmt.Println(v10, v2)
                                        v13 -= v13
                                        switch 707 {
                                        case 779:
                                            fmt.Println(v11, v2)
                                        }
                                    } else {
                                        if false {
                                            if false {
                                                if (127 < 164) {
                                                    var v33 float32 = 0.1
                                                    v34 := GenS4{F0: int32(v2)}
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    for fi35 := 0; fi35 < 7; fi35++ {
                                                        var v36 uint8 = uint8(v5)
                                                        fmt.Println(v36, v2, v10)
                                                        v1 = 255
                                                        _ = fi35
                                                        _ = v36
                                                    }
                                                }
                                                v37 := v4["foo bar"]
                                                var v38 int32 = -1
                                                _ = v37
                                                _ = v38
                                            }
                                        }
                                    }
                                    if false {
                                        v39 := GenS3{F0: v2}
                                        fmt.Println(v2)
                                        v10 = 637
                                        _ = v39
                                    } else {
                                        v11 -= v11
                                    }
                                    for fi40 := 0; fi40 < 9; fi40++ {
                                        if false {
                                            var v41 int16 = 32767
                                            v2 = (v41 >> ((791) & 63))
                                            _ = v41
                                        } else {
                                            var v42 uint8 = 1
                                            fmt.Println(v5)
                                            for fi43 := 0; fi43 < 11; fi43++ {
                                                fmt.Println(v11, v13)
                                                _ = fi43
                                            }
                                            _ = v42
                                        }
                                        v11 = v11
                                        if true {
                                            var v44 string = "a"
                                            fmt.Println(v10)
                                            v45 := GenS2{F0: 359}
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = fi40
                                    }
                                }
                            }
                            v11 = v11
                            _ = v14
                        } else {
                            switch int(v1) {
                            case 127:
                                v1 = 65535
                            case 4294967295:
                                fmt.Println(v13, v1, v11)
                            case 946:
                                var v46 int = 32767
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 16; fi47++ {
                                var v48 string = "a"
                                v49 := []int{}
                                if len(v49) >= 2 {
                                    v49 = v49[1:len(v49)]
                                }
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                        }
                        v1 += v5
                        for fi50 := 0; fi50 < 12; fi50++ {
                            v51 := GenS1{F0: (v2 * v2)}
                            _ = fi50
                            _ = v51
                        }
                    }
                    v10 = v10
                    fmt.Println(v5)
                } else {
                    v13 += 42
                    for fi52 := 0; fi52 < 17; fi52++ {
                        var v53 float32 = (1e10 / (-2.25 * 1.5))
                        _ = fi52
                        _ = v53
                    }
                    switch 4294967295 {
                    case -2147483648:
                        _ = len(v3)
                    case 1:
                        _ = len(v4)
                    case -440:
                        for fi54 := 0; fi54 < 5; fi54++ {
                            var v55 int = (fi54 >> ((486) & 63))
                            _ = fi54
                            _ = v55
                        }
                    default:
                        var v56 int16 = 1000
                        _ = v56
                    }
                }
                for fi57 := 0; fi57 < 2; fi57++ {
                    switch -855 {
                    case -2147483648:
                        v4["foo bar"] = (!false)
                    case 0:
                        var v58 int64 = -1
                        _ = v58
                    case 32767:
                        v13 -= -1
                    }
                    _ = fi57
                }
            } else {
                for fi59 := 0; fi59 < 5; fi59++ {
                    var v60 string = "\n\t"
                    for fi61 := 0; fi61 < 4; fi61++ {
                        v2 = v2
                        v62 := v4[""]
                        v13 = 0
                        _ = fi61
                        _ = v62
                    }
                    v63 := GenS3{F0: 85}
                    _ = fi59
                    _ = v60
                    _ = v63
                }
            }
            v10 = v10
        }
        v64 := GenS5{F0: 255, F1: 4294967295, F2: v2}
        _ = v13
        _ = v64
    } else {
        for fi65 := 0; fi65 < 18; fi65++ {
            var v66 bool = false
            v67 := GenS2{F0: 1}
            for fi68 := 0; fi68 < 8; fi68++ {
                fmt.Println(fi65, v66, v5)
                v11 = 3.14159
                var v69 uint16 = (v5 | v1)
                _ = fi68
                _ = v69
            }
            _ = fi65
            _ = v66
            _ = v67
        }
        for fi70 := 0; fi70 < 8; fi70++ {
            var v71 float64 = 0.0
            var v72 string = "a"
            v73 := []int{}
            _ = fi70
            _ = v71
            _ = v72
            _ = v73
        }
    }
    var v74 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v9
}
