package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint8) uint {
    var v5 uint = 2
    var v6 uint = v5
    var v7 uint = (v6 & v6)
    var v8 uint8 = (p1 + 1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (1000 == 0)
    var v3 int = -1
    v4 := make(map[int]string)
    fmt.Println(v3, v2)
    _ = len(v4)
    v3 = v3
    v9 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: 42}
    v10 := []int{v3, (-(v3)), v3, v3}
    switch "" {
    case "foo bar":
        var v11 float64 = (1e10 - -2.25)
        _ = v11
    case "hello":
        var v12 int16 = -1
        _ = v12
    case "":
        v2 = (2147483647 < -1)
    }
    if (1e10 > 1e10) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS2{F0: -1, F1: -9223372036854775808}
            switch fi13 {
            case 128:
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS1{F0: -406, F1: "", F2: 65535}
                    if v2 {
                        if v2 {
                            var v17 uint16 = 1
                            if true {
                                var v18 string = (("\n\t" + "hello") + ("hello" + "\n\t"))
                                if ("foo bar" == v18) {
                                    v18 = v18
                                    v10 = append(v10, (fi13 % ((1) | 1)))
                                    v17 = 255
                                }
                                if v2 {
                                    v19 := GenS3{F0: 384}
                                    _ = v19
                                }
                                _ = v18
                            } else {
                                v20 := GenS1{F0: 738, F1: "", F2: 2}
                                v17 = 2
                                if (-2.25 == -2.25) {
                                    var v21 int8 = 42
                                    switch -2147483648 {
                                    case 1:
                                        fmt.Println(v3)
                                    case 32767:
                                        v3 = fi15
                                    }
                                    var v22 int8 = (v21 * 1)
                                    _ = v21
                                    _ = v22
                                } else {
                                    fmt.Println(v17, v3)
                                    var v23 uint16 = v17
                                    _ = v23
                                }
                                _ = v20
                            }
                            _ = v17
                        }
                        v1[(fi15 << ((907) & 63))] = v3
                        switch 843 {
                        case 9223372036854775807:
                            if (fi15 == 32767) {
                                v1[1] = -2147483648
                                fmt.Println(fi13)
                            }
                        case 32767:
                            fmt.Println(v3)
                        case -9223372036854775808:
                            v3 += 32768
                        default:
                            if (!v2) {
                                v3 -= fi15
                            }
                        }
                        v10 = append(v10, (v3 / (((-(fi13))) | 1)))
                    } else {
                        if false {
                            _ = len(v10)
                        } else {
                            if true {
                                v24 := GenS4{F0: 42, F1: 759}
                                var v25 int = (-(v3))
                                var v26 uint = 65535
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                if (float32((float64(-2.25) + 3.14159)) <= 0.0) {
                                    var v27 uint32 = 1
                                    _ = v27
                                } else {
                                    v28 := GenS3{F0: -398}
                                    var v29 int8 = 0
                                    v30 := GenS5{F0: 127, F1: 931}
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS1{F0: 0, F1: "", F2: 42}
                                var v32 int8 = 127
                                v3 -= (fi13 % ((42) | 1))
                                _ = v31
                                _ = v32
                            }
                            var v33 float64 = 1.5
                            v2 = (504 == -32768)
                            _ = v33
                        }
                    }
                    var v34 int64 = 104
                    _ = fi15
                    _ = v16
                    _ = v34
                }
            case -1:
                if v2 {
                    var v35 int8 = 1
                    fmt.Println(fi13, v2, fi13)
                    _ = v35
                } else {
                    var v36 string = ""
                    _ = v36
                }
            default:
                var v37 int8 = -1
                _ = v37
            }
            for fi38 := 0; fi38 < 18; fi38++ {
                fmt.Println(fi13, v2)
                v3 = (v3 + 32767)
                _ = fi38
            }
            _ = fi13
            _ = v14
        }
        for fi39 := 0; fi39 < 5; fi39++ {
            v40 := GenS6{F0: -37}
            var v41 int8 = -72
            _ = fi39
            _ = v40
            _ = v41
        }
        v3 = (v3 << ((4294967295) & 63))
    } else {
        switch ("a" + ("\n\t" + "a")) {
        case "a":
            _ = len(v10)
        case "\n\t":
            if (false || v2) {
                v2 = (v2 || (((4294967295 >= 18446744073709551615) || v2) || (!false)))
                v42 := GenS7{F0: 4294967295, F1: v3}
                var v43 uint = 0
                _ = v42
                _ = v43
            }
        case "hello":
            var v44 string = ("foo bar" + ("a" + "\n\t"))
            _ = v44
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            v46 := GenS4{F0: 2147483647, F1: 65535}
            v2 = v2
            var v47 uint = 65535
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        if len(v10) > 0 {
            v10[0] = v3
        }
        if v2 {
            switch "" {
            case "a":
                if len(v10) > 0 {
                    v10[0] = 9223372036854775807
                }
            }
            v2 = (!(!(v2 || ((((v2 || (361 >= 2)) || (-9223372036854775808 <= 42)) && true) || v2))))
        } else {
            for fi49 := 0; fi49 < 7; fi49++ {
                if (!(!("foo bar" > "hello"))) {
                    v2 = (v2 && ((127 <= 100) && v2))
                    fmt.Println(v3)
                    if v2 {
                        v2 = (-508 <= -450)
                    } else {
                        var v50 uint = 152
                        _ = v50
                    }
                }
                if (255 <= 0) {
                    var v51 uint8 = 127
                    var v52 uint = 511
                    if v2 {
                        v3 += fi48
                    } else {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v53 int16 = 345
                        v54 := GenS8{F0: (v51 % ((v51) | 1)), F1: (v51 - v51)}
                        if v2 {
                            switch "" {
                            case "":
                                v55 := GenS9{F0: 319}
                                _ = v55
                            }
                        } else {
                            delete(v4, fi49)
                        }
                        _ = v53
                        _ = v54
                    }
                    _ = v51
                    _ = v52
                }
                var v56 int = fi49
                _ = fi49
                _ = v56
            }
            for fi57 := 0; fi57 < 5; fi57++ {
                if v2 {
                    var v58 int32 = -1
                    _ = v58
                } else {
                    var v59 uint16 = 1
                    var v60 int64 = 1
                    v60 -= -26
                    _ = v59
                    _ = v60
                }
                _ = fi57
            }
            v61 := GenS1{F0: -1, F1: "", F2: 699}
            _ = v61
        }
        _ = fi48
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
