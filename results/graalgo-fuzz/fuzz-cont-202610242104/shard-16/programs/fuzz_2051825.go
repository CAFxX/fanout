package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int16; F1 string }
func genFn0(p0 bool) bool {
    v6 := GenS1{F0: 255, F1: 404}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint32 = 36
            _ = fi8
            _ = v9
        }
        p0 = (1 != 2147483647)
        _ = fi7
    }
    _ = p0
    _ = v6
    return (42 > 65535)
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    v5 := []uint32{2147483648, 0}
    if false {
        switch 127 {
        case -73:
            var v10 float64 = float64(1.5)
            _ = v10
        case -9223372036854775808:
            fmt.Println(v1)
        case 4294967295:
            var v11 int32 = 244
            _ = v11
        }
        var v12 int16 = -32768
        fmt.Println(v12)
        _ = v12
    } else {
        var v13 int64 = -1
        _ = v13
    }
    delete(v3, 42)
    var v14 float64 = float64(0.0)
    fmt.Println(v1, v14)
    if true {
        if (!true) {
            v1 = v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            v15 := GenS2{F0: 214, F1: -72}
            if (1 <= 909) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint32 = 2
                    _ = fi16
                    _ = v17
                }
                switch 4294967295 {
                case 42:
                    v1 = ""
                default:
                    fmt.Println(v14, v1)
                }
                if true {
                    switch -9223372036854775808 {
                    case 4294967295:
                        v18 := GenS3{F0: 42}
                        _ = v18
                    case 128:
                        v1 = ((v1 + (v1 + "")) + (v1 + "hello"))
                    default:
                        v14 -= ((v14 * -2.25) + ((((3.14159 - float64(0.1)) + (1.5 / v14)) * float64((float32(-2.25) * 0.0))) - v14))
                    }
                    var v19 int32 = 758
                    _ = v19
                }
            }
            _ = v15
        }
        v1 = v1
    }
    var v20 uint = 4294967295
    var v21 int8 = 127
    fmt.Println(v20, v20, v20)
    if (4294967295 > 0) {
        v22 := GenS4{F0: 42}
        _ = v22
    } else {
        fmt.Println(v21, v1)
    }
    if (!(1 <= 32767)) {
        var v23 float32 = ((-2.25 / 1e10) - -2.25)
        fmt.Println(v23, v20)
        _ = v23
    } else {
        var v24 string = (v1 + v1)
        _ = v24
    }
    if ((128 != 140) && (false || (!((42 > 32767) || (!true))))) {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint16 = 2
            v27 := v2[2]
            if true {
                var v28 uint8 = 30
                switch "foo bar" {
                case "a":
                    v21 -= v21
                default:
                    if true {
                        fmt.Println(v26)
                        if (!true) {
                            if (("hello" != "foo bar") && false) {
                                v28 = (v28 % (((v28 << (((v20 & v20)) & 63))) | 1))
                            }
                            switch (fi25 >> (((v20 >> ((v20) & 63))) & 63)) {
                            case 32768:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case -1:
                                v29 := GenS2{F0: (v28 & v28), F1: v21}
                                _ = v29
                            case 959:
                                v4[408] = ((!("hello" != ((("\n\t" + v1) + (v1 + "foo bar")) + "foo bar"))) && false)
                            default:
                                if (false || (false || (2147483647 > int32(v21)))) {
                                    var v30 int32 = 42
                                    _ = v30
                                } else {
                                    fmt.Println(v14, v1, fi25)
                                    for fi31 := 0; fi31 < 2; fi31++ {
                                        var v32 int32 = 1
                                        switch fi25 {
                                        case 42:
                                            v27 = v27
                                        default:
                                            if true {
                                                if (1000 >= 42) {
                                                    if (!false) {
                                                        var v33 int16 = 0
                                                        fmt.Println(v1, fi25)
                                                        _ = v33
                                                    } else {
                                                        if true {
                                                            v27 = v27
                                                            var v34 uint8 = (v28 ^ v28)
                                                            v35 := GenS5{F0: (fi25 | fi31), F1: 42, F2: v26}
                                                            _ = v34
                                                            _ = v35
                                                        } else {
                                                            v21 += v21
                                                            fmt.Println(v32)
                                                            v26 = v26
                                                        }
                                                        var v36 uint = 2
                                                        var v37 uint16 = v26
                                                        var v38 uint32 = 2147483648
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                                delete(v4, 65535)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v39 := GenS6{F0: (v21 >> ((2) & 63)), F1: (v20 ^ (^(v20))), F2: 1}
                                                    _ = v39
                                                }
                                            }
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                    fmt.Println(v21, v14, v14)
                                }
                            }
                        }
                        if true {
                            v40 := GenS5{F0: 32768, F1: (v21 - (v21 % ((0) | 1))), F2: 42}
                            v20 = v20
                            v41 := GenS7{F0: 963, F1: 538}
                            _ = v40
                            _ = v41
                        } else {
                            if (v20 < 1) {
                                if ((2 >= 35) || true) {
                                    fmt.Println(v27, v21)
                                    v20 = (v20 & (v20 * v20))
                                    var v42 uint32 = uint32(v28)
                                    _ = v42
                                } else {
                                    var v43 int16 = 1000
                                    var v44 uint32 = 42
                                    _ = v43
                                    _ = v44
                                }
                                v21 -= 1
                            } else {
                                var v45 bool = true
                                if (791 < 0) {
                                    v46 := GenS6{F0: 100, F1: (v20 / (((^(v20))) | 1)), F2: 1}
                                    for fi47 := 0; fi47 < 2; fi47++ {
                                        switch (fi25 % ((fi25) | 1)) {
                                        case 32767:
                                            var v48 uint64 = 18446744073709551615
                                            _ = v48
                                        case 9223372036854775807:
                                            var v49 uint = (v20 & 255)
                                            _ = v49
                                        case 2147483647:
                                            if ((-32768 <= 1) || ((2147483648 == 1) || v45)) {
                                                if len(v5) > 0 {
                                                    v5[0] = 381
                                                }
                                                var v50 uint32 = 864
                                                fmt.Println(v1, v20)
                                                _ = v50
                                            } else {
                                                v14 += float64(0.0)
                                            }
                                        }
                                        v27 = (v1 + "a")
                                        var v51 string = "foo bar"
                                        _ = fi47
                                        _ = v51
                                    }
                                    v28 -= 2
                                    _ = v46
                                } else {
                                    if false {
                                        v27 = v27
                                        for fi52 := 0; fi52 < 20; fi52++ {
                                            v2[2147483648] = v1
                                            _ = fi52
                                        }
                                        v1 = (v27 + "")
                                    } else {
                                        switch (fi25 - -489) {
                                        case 127:
                                            v20 = v20
                                        default:
                                            _ = len(v5)
                                        }
                                        v27 = v27
                                        v53 := GenS7{F0: 52, F1: 2147483647}
                                        _ = v53
                                    }
                                    fmt.Println(v1)
                                }
                                v26 += v26
                                _ = v45
                            }
                            if true {
                                v27 = v27
                                fmt.Println(v20, v1, v26)
                            } else {
                                var v54 float64 = v14
                                v1 = "\n\t"
                                _ = v54
                            }
                            fmt.Println(v20)
                        }
                    } else {
                        var v55 string = v27
                        if len(v5) > 0 {
                            v5[0] = 90
                        }
                        _ = v55
                    }
                }
                v2[474] = (("\n\t" + v27) + ((((v27 + v27) + v1) + v1) + (("hello" + v27) + v1)))
                _ = v28
            } else {
                fmt.Println(v20, v1, v20)
                v4[185] = (v26 < v26)
                for fi56 := 0; fi56 < 7; fi56++ {
                    v5 = append(v5, 176)
                    if false {
                        v57 := GenS4{F0: 127}
                        v20 = v20
                        v58 := GenS1{F0: 245, F1: 42}
                        v59 := GenS5{F0: (fi25 / ((128) | 1)), F1: (v21 ^ v21), F2: v26}
                        delete(v3, fi25)
                        _ = v57
                        _ = v58
                        _ = v59
                    } else {
                        fmt.Println(v1, fi25, fi56)
                    }
                    if false {
                        v20 += (^(v20))
                        if false {
                            v1 = v27
                        }
                        v27 = ""
                    } else {
                        v60 := GenS8{F0: 32767, F1: (v1 + (("foo bar" + v1) + (v27 + (v27 + v1))))}
                        _ = v60
                    }
                    _ = fi56
                }
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        switch 9223372036854775807 {
        case 2147483647:
            v21 = v21
        case 0:
            v61 := GenS1{F0: 51, F1: 65535}
            _ = v61
        case -2147483648:
            if false {
                for fi62 := 0; fi62 < 18; fi62++ {
                    _ = len(v5)
                    v21 -= v21
                    _ = fi62
                }
            } else {
                v21 = v21
            }
        }
        v14 -= ((v14 + v14) - 0.1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
