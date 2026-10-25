package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 string }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := []int{609, 42, -9223372036854775808}
    var v4 int64 = 131
    v5 := GenS1{F0: 1}
    if true {
        var v6 float64 = 3.14159
        var v7 int8 = -128
        var v8 int32 = 42
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 9}
        _ = v9
    }
    v10 := GenS2{F0: (^(v4))}
    var v11 int64 = (v4 >> ((18446744073709551615) & 63))
    v12 := GenS3{F0: "a", F1: ((1 != 159) && true)}
    switch -1 {
    case 127:
        v13 := GenS4{F0: 127, F1: (v11 + 9223372036854775807), F2: "foo bar"}
        _ = v13
    case -96:
        v4 += (v11 / ((v11) | 1))
    case -887:
        var v14 float32 = 1e10
        _ = v14
    }
    v4 = (v11 >> ((4294967295) & 63))
    if (1 <= 42) {
        v15 := GenS5{F0: "a", F1: 65535, F2: 2147483647}
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int32 = -2147483648
        v18 := GenS6{F0: false, F1: "a"}
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 int16 = -32768
            v21 := GenS7{F0: (-1 >= 0), F1: 42}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    _ = len(v3)
    if (212 <= 32767) {
        if false {
            v22 := GenS4{F0: 120, F1: v11, F2: "\n\t"}
            v23 := GenS3{F0: "a", F1: false}
            for fi24 := 0; fi24 < 11; fi24++ {
                fmt.Println(v11, v4, fi24)
                _ = fi24
            }
            _ = v22
            _ = v23
        } else {
            for fi25 := 0; fi25 < 12; fi25++ {
                if (1 <= 1) {
                    v26 := GenS6{F0: (1 == (v4 - (v11 * (-(v11))))), F1: "a"}
                    switch fi25 {
                    case 42:
                        var v27 uint = 380
                        _ = v27
                    default:
                        if false {
                            if false {
                                var v28 uint16 = 2
                                _ = v28
                            } else {
                                var v29 int8 = 103
                                _ = v29
                            }
                            _ = len(v3)
                        } else {
                            v4 = (v11 << ((326) & 63))
                            var v30 uint64 = 18446744073709551615
                            v4 = -9223372036854775808
                            _ = v30
                        }
                    }
                    v3 = append(v3, fi25)
                    _ = v26
                } else {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v32 := GenS8{F0: (v4 / ((-1) | 1))}
                        _ = fi31
                        _ = v32
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch (("" + "a") + "\n\t") {
                case "\n\t":
                    if (true || (!(!(-1 == 11)))) {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            var v34 int8 = 1
                            switch fi25 {
                            case -916:
                                var v35 int = fi25
                                _ = v35
                            default:
                                fmt.Println(v4, fi25)
                            }
                            v4 -= -1
                            _ = fi33
                            _ = v34
                        }
                        var v36 float32 = ((1e10 + ((1e10 + 1e10) * 1.5)) * 0.0)
                        for fi37 := 0; fi37 < 10; fi37++ {
                            v36 = (1.5 - v36)
                            _ = fi37
                        }
                        _ = v36
                    }
                case "":
                    delete(v2, 539)
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi25
            }
        }
    } else {
        var v38 uint8 = 0
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        fmt.Println(fi39, fi39)
        if false {
            if (-1 >= -911) {
                fmt.Println(v4)
                var v40 uint8 = 42
                v41 := GenS8{F0: (v4 | -1)}
                _ = v40
                _ = v41
            } else {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v11 += v11
                    v2[-602] = (fi39 << ((2) & 63))
                    _ = fi42
                }
            }
            if true {
                var v43 float32 = 1e10
                v44 := GenS2{F0: (v4 & (v4 + (v11 | 42)))}
                _ = v43
                _ = v44
            }
        }
        var v45 bool = ((-2.25 * 3.14159) == (1e10 / -2.25))
        _ = fi39
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        if false {
            fmt.Println(v4, fi46)
            v47 := GenS4{F0: 177, F1: v11, F2: ""}
            v3 = append(v3, (fi46 / ((-868) | 1)))
            _ = v47
        }
        if (false || false) {
            v11 += 1
        } else {
            for fi48 := 0; fi48 < 4; fi48++ {
                v49 := GenS6{F0: true, F1: (("a" + "hello") + "")}
                fmt.Println(v11, fi48)
                _ = fi48
                _ = v49
            }
            v4 = (v11 ^ v4)
        }
        v50 := GenS2{F0: v11}
        _ = fi46
        _ = v50
    }
    v4 -= v11
    var v51 uint = 42
    if true {
        v4 -= (v4 | (v4 / (((^(v4))) | 1)))
        v2[32768] = 4294967295
        if true {
            v11 -= -684
            if (512 != 4294967295) {
                delete(v1, "")
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 int64 = int64(v51)
                    _ = fi52
                    _ = v53
                }
            }
            if (127 == 127) {
                v54 := GenS8{F0: (v11 / ((-1) | 1))}
                _ = v54
            } else {
                for fi55 := 0; fi55 < 18; fi55++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v56 float32 = 0.0
                    _ = fi55
                    _ = v56
                }
                delete(v2, 198)
            }
        } else {
            if false {
                var v57 float64 = ((-2.25 / float64(3.14159)) / 1.5)
                v58 := GenS4{F0: 121, F1: (v4 - int64(v51)), F2: ("" + ("foo bar" + "a"))}
                v57 -= v57
                _ = v57
                _ = v58
            }
        }
    }
    var v59 float64 = ((0.1 + float64(0.1)) + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v59
}
