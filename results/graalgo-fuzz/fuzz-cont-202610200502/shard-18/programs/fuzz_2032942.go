package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
type GenS4 struct { F0 int }

func main() {
    var v1 int16 = -58
    v2 := make(map[int]bool)
    var v3 int8 = 127
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    var v6 uint32 = 4294967295
    switch 814 {
    case 127:
        v7 := GenS1{F0: 2, F1: (-(v1)), F2: v6}
        _ = v7
    default:
        fmt.Println(v6)
    }
    fmt.Println(v1, v3)
    fmt.Println(v1, v6, v6)
    v3 += v3
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint = 42
        for fi10 := 0; fi10 < 10; fi10++ {
            switch fi8 {
            case -628:
                if false {
                    v11 := []int{-1, (fi8 + fi10), (fi10 * (fi10 + fi8)), fi8}
                    var v12 int8 = 114
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v11
                    _ = v12
                } else {
                    v1 -= (^(v1))
                    switch "hello" {
                    case "":
                        switch fi8 {
                        case -963:
                            var v13 float32 = 0.1
                            v14 := GenS1{F0: 128, F1: 1000, F2: uint32(v9)}
                            _ = v13
                            _ = v14
                        case 32767:
                            var v15 int = fi8
                            _ = v15
                        case -602:
                            delete(v4, fi10)
                        }
                    case "\n\t":
                        if true {
                            var v16 int64 = 0
                            var v17 int16 = v1
                            v18 := []int{}
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    default:
                        var v19 uint32 = (v6 | 0)
                        _ = v19
                    }
                    v20 := GenS2{F0: 65535, F1: true}
                    _ = v20
                }
            case 9223372036854775807:
                v6 = (v6 % ((2) | 1))
            case 0:
                v21 := GenS2{F0: 65535, F1: true}
                _ = v21
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    delete(v2, 9223372036854775807)
    var v22 int16 = v1
    var v23 uint64 = 42
    v4[304] = v6
    v1 += v1
    v22 += (v22 ^ v1)
    var v24 uint32 = (^(v6))
    fmt.Println(v23)
    for fi25 := 0; fi25 < 10; fi25++ {
        switch fi25 {
        case 2147483647:
            var v26 string = (("a" + "\n\t") + "")
            _ = v26
        case 9223372036854775807:
            v3 = (v3 / ((-128) | 1))
        case -316:
            v24 += v6
        }
        v27 := GenS2{F0: 1, F1: true}
        _ = fi25
        _ = v27
    }
    switch -1 {
    case -2147483648:
        if true {
            v28 := GenS2{F0: 32767, F1: false}
            for fi29 := 0; fi29 < 7; fi29++ {
                v22 -= 32767
                v1 = 788
                switch ("\n\t" + "\n\t") {
                case "hello":
                    v30 := GenS2{F0: 2, F1: true}
                    _ = v30
                case "":
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v32 := []int{(fi29 % ((1) | 1)), (fi29 / ((fi31) | 1))}
                        _ = fi31
                        _ = v32
                    }
                case "foo bar":
                    v3 += v3
                default:
                    fmt.Println(v24)
                }
                _ = fi29
            }
            var v33 uint8 = 42
            _ = v28
            _ = v33
        }
    default:
        switch -9223372036854775808 {
        case 42:
            switch 32767 {
            case 32768:
                for fi34 := 0; fi34 < 4; fi34++ {
                    switch fi34 {
                    case 127:
                        v35 := GenS3{F0: 65535, F1: v3, F2: (255 != 65535)}
                        _ = v35
                    case 29:
                        for fi36 := 0; fi36 < 8; fi36++ {
                            if false {
                                if false {
                                    fmt.Println(fi34, v3, v23)
                                } else {
                                    var v37 int = fi36
                                    v37 = v37
                                    var v38 uint8 = 0
                                    _ = v37
                                    _ = v38
                                }
                                fmt.Println(v3)
                            } else {
                                var v39 uint32 = (v24 ^ 65535)
                                var v40 uint8 = uint8(v23)
                                _ = v39
                                _ = v40
                            }
                            v4[(-(fi34))] = (v24 >> ((425) & 63))
                            var v41 uint16 = 65535
                            _ = fi36
                            _ = v41
                        }
                    }
                    if false {
                        fmt.Println(v22, v23)
                    } else {
                        var v42 bool = (234 == 128)
                        _ = v42
                    }
                    _ = fi34
                }
            case -9223372036854775808:
                for fi43 := 0; fi43 < 13; fi43++ {
                    if false {
                        var v44 uint8 = 124
                        v44 -= 130
                        var v45 bool = (255 != (v44 << ((423) & 63)))
                        _ = v44
                        _ = v45
                    } else {
                        var v46 string = ("a" + "\n\t")
                        if false {
                            if true {
                                v47 := []int{}
                                v47 = append(v47, -1)
                                switch fi43 {
                                case 0:
                                    if (true || ((true || (-473 == 42)) && false)) {
                                        if true {
                                            v23 += v23
                                        }
                                        for fi48 := 0; fi48 < 5; fi48++ {
                                            if len(v47) >= 2 {
                                                v47 = v47[1:len(v47)]
                                            }
                                            _ = fi48
                                        }
                                        var v49 uint32 = (v6 / (((v24 >> ((1) & 63))) | 1))
                                        _ = v49
                                    }
                                default:
                                    if (4294967295 < 1) {
                                        for fi50 := 0; fi50 < 3; fi50++ {
                                            if len(v47) > 0 {
                                                v47[0] = 991
                                            }
                                            if (!false) {
                                                v51 := GenS1{F0: 244, F1: (v1 % ((v1) | 1)), F2: v6}
                                                _ = v51
                                            }
                                            _ = fi50
                                        }
                                    } else {
                                        switch v46 {
                                        case "hello":
                                            if false {
                                                v52 := GenS2{F0: 42, F1: false}
                                                _ = v52
                                            }
                                        case "":
                                            v23 += 628
                                        case "a":
                                            var v53 bool = ((("" + v46) < ("\n\t" + (v46 + "foo bar"))) && true)
                                            _ = v53
                                        default:
                                            v54 := v5[(fi43 % (((fi43 / (((fi43 | (-(fi43)))) | 1))) | 1))]
                                            _ = v54
                                        }
                                    }
                                }
                                _ = v47
                            } else {
                                for fi55 := 0; fi55 < 5; fi55++ {
                                    v56 := GenS2{F0: 0, F1: true}
                                    _ = fi55
                                    _ = v56
                                }
                                v57 := v2[(fi43 | fi43)]
                                fmt.Println(v57, v57, v6)
                                _ = v57
                            }
                            var v58 uint = 65535
                            v3 = (v3 & v3)
                            _ = v58
                        }
                        _ = v46
                    }
                    for fi59 := 0; fi59 < 16; fi59++ {
                        switch (("hello" + "\n\t") + "\n\t") {
                        case "foo bar":
                            v5[32767] = false
                        case "hello":
                            if (((127 <= int8(v23)) || false) || false) {
                                var v60 uint = 0
                                _ = v60
                            }
                        case "a":
                            fmt.Println(fi59, v23)
                        }
                        _ = fi59
                    }
                    _ = fi43
                }
            default:
                v61 := GenS1{F0: 127, F1: (v1 | v22), F2: v6}
                _ = v61
            }
        case 0:
            if (false || ((!(v3 == (-(v3)))) || true)) {
                var v62 uint = 2
                _ = v62
            }
        case -627:
            fmt.Println(v22)
        }
    }
    var v63 float64 = 0.0
    v4[32768] = v24
    v64 := GenS4{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
