package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]int)
    var v3 int = -2147483648
    switch "hello" {
    case "foo bar":
        var v4 uint = 4294967295
        _ = v4
    case "":
        var v5 int64 = (-(v1))
        _ = v5
    default:
        delete(v2, 1)
    }
    if (!true) {
        if ((1e10 > 0.1) && false) {
            if false {
                var v6 string = "hello"
                _ = v6
            }
        }
        var v7 int8 = -1
        switch "a" {
        case "a":
            v8 := GenS1{F0: 494}
            _ = v8
        }
        _ = v7
    } else {
        v9 := GenS1{F0: -813}
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        var v10 uint32 = 42
        _ = v10
    case "":
        v3 = v3
    default:
        switch v3 {
        case 127:
            fmt.Println(v3)
        case 32767:
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int16 = 415
                _ = fi11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: 32767}
    v14 := []int{(v3 / ((0) | 1))}
    switch (v3 * (v3 + 1)) {
    case 1:
        v3 -= (v3 | (v3 / (((-(v3))) | 1)))
    case -1:
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: (v1 & -619), F1: 43, F2: fi15}
            _ = fi15
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 uint8 = 127
                v19 -= v19
                if false {
                    v1 = v1
                    v20 := GenS1{F0: -1}
                    if true {
                        fmt.Println(v1)
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            if (2147483647 > 128) {
                                fmt.Println(v19, v1)
                            } else {
                                switch -2147483648 {
                                case 0:
                                    if true {
                                        var v21 int16 = 177
                                        v3 = (fi17 | fi17)
                                        var v22 uint64 = 4294967295
                                        _ = v21
                                        _ = v22
                                    }
                                case 764:
                                    v19 += v19
                                default:
                                    var v23 int8 = 5
                                    _ = v23
                                }
                                if (true && true) {
                                    var v24 uint8 = (v19 % (((^(v19))) | 1))
                                    var v25 int16 = -32768
                                    _ = v24
                                    _ = v25
                                } else {
                                    if len(v14) > 0 {
                                        v14[0] = -496
                                    }
                                    fmt.Println(fi17, v3)
                                }
                                if (((true && true) && false) || true) {
                                    v3 += (v3 / (((fi17 / ((2147483647) | 1))) | 1))
                                    v19 = (v19 ^ (v19 & (v19 ^ 0)))
                                    var v26 uint = 65535
                                    _ = v26
                                }
                            }
                        case "a":
                            var v27 float32 = (1e10 / 1.0)
                            _ = v27
                        default:
                            var v28 float64 = -2.25
                            _ = v28
                        }
                        v29 := GenS2{F0: v1, F1: (v19 % ((v19) | 1)), F2: 42}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v29
                    }
                    _ = v20
                } else {
                    switch 551 {
                    case 2147483647:
                        var v30 uint64 = 4294967295
                        _ = v30
                    case -2147483648:
                        if (544 == 591) {
                            switch fi17 {
                            case 32768:
                                v31 := GenS1{F0: 513}
                                _ = v31
                            case -9223372036854775808:
                                var v32 int = v3
                                var v33 float64 = -2.25
                                v3 -= v3
                                _ = v32
                                _ = v33
                            case -2147483648:
                                var v34 uint = 255
                                _ = v34
                            }
                            if (false || (true && true)) {
                                var v35 int8 = -1
                                _ = v35
                            } else {
                                var v36 string = "a"
                                v37 := v2[0]
                                v38 := GenS2{F0: 0, F1: v19, F2: (fi17 >> ((459) & 63))}
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            v39 := GenS3{F0: (v3 & v3), F1: (^(v19)), F2: "a"}
                            _ = v39
                        }
                    default:
                        fmt.Println(v1, v1)
                    }
                    var v40 int32 = -516
                    _ = v40
                }
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    if (int16(v3) < 32767) {
        v3 = (v3 | (^(v3)))
        if true {
            _ = len(v2)
            var v41 uint = 1
            var v42 string = "foo bar"
            _ = v41
            _ = v42
        }
        v43 := GenS4{F0: (v1 | v1)}
        _ = v43
    }
    v14 = append(v14, 127)
    var v44 int8 = 0
    for fi45 := 0; fi45 < 7; fi45++ {
        for fi46 := 0; fi46 < 17; fi46++ {
            var v47 float64 = 0.0
            _ = fi46
            _ = v47
        }
        var v48 uint8 = 2
        if (("\n\t" != "hello") && false) {
            switch (fi45 & int(v48)) {
            case 32767:
                if (!false) {
                    for fi49 := 0; fi49 < 15; fi49++ {
                        v3 = (v3 << ((2) & 63))
                        _ = fi49
                    }
                    var v50 uint16 = 42
                    var v51 int8 = v44
                    _ = v50
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 18; fi52++ {
                        v53 := GenS5{F0: (^(fi52))}
                        v3 = 388
                        v54 := GenS6{F0: (v1 ^ (v1 + (v1 | (-(v1)))))}
                        _ = len(v2)
                        _ = len(v14)
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                }
            default:
                fmt.Println(v44, v1)
            }
        } else {
            v3 += v3
            switch "hello" {
            case "hello":
                if false {
                    v55 := GenS7{F0: v1, F1: 2147483647}
                    _ = v55
                }
            default:
                v3 -= fi45
            }
        }
        _ = fi45
        _ = v48
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        for fi57 := 0; fi57 < 2; fi57++ {
            var v58 int = fi56
            _ = fi57
            _ = v58
        }
        v44 -= v44
        var v59 uint64 = 42
        _ = fi56
        _ = v59
    }
    v1 = v1
    var v60 uint64 = 796
    if true {
        delete(v2, uint32(v60))
        var v61 float64 = 0.1
        var v62 bool = true
        _ = v61
        _ = v62
    } else {
        v1 += v1
        if false {
            var v63 int = (v3 | v3)
            var v64 int16 = 1
            _ = v63
            _ = v64
        }
    }
    v65 := GenS5{F0: (v3 | (^(v3)))}
    v66 := GenS5{F0: (v3 ^ 32768)}
    v2[2] = (v3 ^ (v3 << ((453) & 63)))
    v3 -= 111
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v44
    _ = v60
    _ = v65
    _ = v66
}
