package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 uint64; F2 uint }
func genFn0(p0 int16, p1 string) int32 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: fi7, F1: -403, F2: 193}
        v9 := GenS1{F0: 576, F1: 1, F2: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: -1, F1: -1, F2: -1}
        v12 := []int{fi10}
        var v13 float64 = 1.5
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    var v3 uint8 = 255
    var v4 uint = 42
    var v5 float32 = (1e10 - 0.0)
    v6 := make(map[int]uint32)
    if true {
        v14 := GenS1{F0: 9223372036854775807, F1: genFn0(1, v2), F2: 42}
        _ = v14
    } else {
        if true {
            var v15 uint16 = 0
            _ = v15
        } else {
            v5 = v5
            if (!true) {
                v4 -= 2
            }
            var v16 uint = (v4 & (v4 & v4))
            _ = v16
        }
        var v17 int8 = 1
        for fi18 := 0; fi18 < 20; fi18++ {
            if (-19 != 32767) {
                v19 := GenS2{F0: (^(v17)), F1: 1}
                var v20 uint = 18446744073709551615
                if true {
                    switch "hello" {
                    case "a":
                        v21 := GenS1{F0: fi18, F1: 0, F2: 42}
                        _ = v21
                    }
                } else {
                    if true {
                        if false {
                            if (463 <= 85) {
                                var v22 int16 = 0
                                _ = v22
                            } else {
                                if (int(v20) < (-(fi18))) {
                                    fmt.Println(v5)
                                    var v23 float32 = (1e10 / 1.0)
                                    _ = v23
                                } else {
                                    switch fi18 {
                                    case 0:
                                        v4 = v20
                                    case -449:
                                        for fi24 := 0; fi24 < 2; fi24++ {
                                            var v25 uint16 = 255
                                            if true {
                                                var v26 float64 = 0.1
                                                _ = v26
                                            } else {
                                                fmt.Println(fi24, fi24)
                                                fmt.Println(v17, v20, fi18)
                                            }
                                            _ = fi24
                                            _ = v25
                                        }
                                    default:
                                        if true {
                                            v3 += 2
                                            fmt.Println(v3, v5)
                                            fmt.Println(v5, v2, v17)
                                        } else {
                                            v27 := v1[1]
                                            if ((v3 < v3) && true) {
                                                v4 -= 42
                                                var v28 int32 = -1
                                                var v29 int8 = (-(v17))
                                                _ = v28
                                                _ = v29
                                            } else {
                                                for fi30 := 0; fi30 < 19; fi30++ {
                                                    v31 := []int{}
                                                    switch -2147483648 {
                                                    case 9223372036854775807:
                                                        v31 = append(v31, fi18)
                                                    case 128:
                                                        v32 := v6[fi30]
                                                        _ = v32
                                                    case 127:
                                                        switch fi30 {
                                                        case 4294967295:
                                                            var v33 int = (-(fi30))
                                                            _ = v33
                                                        case -2147483648:
                                                            _ = len(v31)
                                                        case -1:
                                                            if ((1000 != -1) && false) {
                                                                var v34 uint16 = 1
                                                                if true {
                                                                    var v35 int32 = 0
                                                                    delete(v6, 42)
                                                                    _ = v35
                                                                }
                                                                v5 += v5
                                                                _ = v34
                                                            } else {
                                                                v17 = v17
                                                                switch ((v2 + "hello") + "") {
                                                                case "\n\t":
                                                                    v27 = v27
                                                                case "hello":
                                                                    fmt.Println(fi18, v3)
                                                                case "foo bar":
                                                                    v36 := GenS1{F0: 2147483647, F1: -2147483648, F2: -553}
                                                                    _ = v36
                                                                default:
                                                                    fmt.Println(fi18)
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if len(v31) >= 2 {
                                                            v31 = v31[1:len(v31)]
                                                        }
                                                    }
                                                    v37 := v6[(fi30 ^ -257)]
                                                    _ = fi30
                                                    _ = v31
                                                    _ = v37
                                                }
                                                v38 := []int{590, (fi18 ^ 32767), 42}
                                                v4 = 42
                                                _ = v38
                                            }
                                            if false {
                                                _ = len(v6)
                                            }
                                            _ = v27
                                        }
                                    }
                                    var v39 float64 = (0.0 - (3.14159 / (3.14159 / 1.5)))
                                    var v40 int = 988
                                    _ = v39
                                    _ = v40
                                }
                                v4 -= 65535
                                v17 = (v17 % ((v17) | 1))
                            }
                            var v41 uint32 = 2147483648
                            var v42 uint16 = 32767
                            _ = v41
                            _ = v42
                        } else {
                            fmt.Println(v20)
                            for fi43 := 0; fi43 < 6; fi43++ {
                                v5 = v5
                                var v44 float32 = 0.1
                                _ = fi43
                                _ = v44
                            }
                        }
                        v45 := GenS3{F0: 4294967295, F1: 1000}
                        _ = v45
                    } else {
                        if true {
                            v46 := GenS4{F0: 158, F1: 1, F2: int64(v3)}
                            v47 := GenS5{F0: 491, F1: v3, F2: 255}
                            _ = v46
                            _ = v47
                        } else {
                            var v48 uint64 = 1
                            fmt.Println(fi18, v3, v17)
                            v20 -= (v20 >> ((361) & 63))
                            _ = v48
                        }
                    }
                    v17 = v17
                }
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = v17
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 int8 = 0
        v51 := GenS2{F0: (v50 - (v50 & (v50 + 1))), F1: 1}
        _ = fi49
        _ = v50
        _ = v51
    }
    _ = len(v1)
    var v52 uint32 = 605
    v53 := GenS4{F0: 1, F1: 42, F2: 59}
    var v54 uint16 = 32767
    v4 += 65535
    switch v2 {
    case "":
        v55 := []int{128}
        _ = v55
    }
    v54 += (^(v54))
    v54 -= 65535
    if false {
        for fi56 := 0; fi56 < 18; fi56++ {
            var v57 int = fi56
            v58 := GenS6{F0: 551, F1: uint64(v52)}
            _ = fi56
            _ = v57
            _ = v58
        }
        v4 = (v4 >> ((65535) & 63))
        v59 := GenS7{F0: v3}
        _ = v59
    } else {
        v60 := GenS7{F0: v3}
        _ = v60
    }
    for fi61 := 0; fi61 < 4; fi61++ {
        fmt.Println(v2, v54)
        fmt.Println(v54, v2, v3)
        switch fi61 {
        case 1:
            v62 := []int{9223372036854775807, 299}
            _ = v62
        case -922:
            switch (((v2 + "foo bar") + ("a" + ((v2 + "a") + v2))) + v2) {
            case "\n\t":
                v63 := []int{fi61, 42}
                _ = v63
            case "":
                var v64 uint16 = 32767
                _ = v64
            default:
                fmt.Println(v52, v52, v54)
            }
        default:
            var v65 float32 = v5
            _ = v65
        }
        _ = fi61
    }
    switch 4294967295 {
    case -2147483648:
        var v66 uint = v4
        _ = v66
    default:
        for fi67 := 0; fi67 < 15; fi67++ {
            if (true || (845 >= fi67)) {
                var v68 int16 = -365
                _ = v68
            }
            var v69 string = v2
            v70 := []int{(fi67 >> ((v4) & 63)), fi67}
            _ = fi67
            _ = v69
            _ = v70
        }
    }
    for fi71 := 0; fi71 < 3; fi71++ {
        for fi72 := 0; fi72 < 2; fi72++ {
            v3 = (^(v3))
            var v73 uint32 = v52
            v6[(fi72 >> ((1) & 63))] = v73
            _ = fi72
            _ = v73
        }
        _ = fi71
    }
    switch (v2 + v2) {
    case "":
        var v74 uint8 = (v3 + (v3 << (((v4 | v4)) & 63)))
        _ = v74
    }
    v75 := GenS7{F0: 197}
    v76 := GenS8{F0: false, F1: 0, F2: v4}
    var v77 int = 993
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v75
    _ = v76
    _ = v77
}
