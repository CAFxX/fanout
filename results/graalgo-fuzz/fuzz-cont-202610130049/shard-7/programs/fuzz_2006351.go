package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 uint; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint16 }
func genFn0() int32 {
    var v7 float64 = 1e10
    switch 32767 {
    case 127:
        v8 := GenS1{F0: 2, F1: 127, F2: 42}
        _ = v8
    case 1:
        v9 := GenS1{F0: 1, F1: -408, F2: 127}
        _ = v9
    }
    if true {
        if false {
            v10 := GenS1{F0: 18446744073709551615, F1: 128, F2: 1}
            _ = v10
        }
        var v11 float64 = v7
        _ = v11
    } else {
        if (!(true && (-61 >= -128))) {
            v12 := GenS2{F0: "\n\t", F1: 639, F2: (!(504 <= 65535))}
            _ = v12
        }
    }
    fmt.Println(v7)
    _ = v7
    return -317
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 42
    var v3 int8 = 127
    var v4 float32 = (1e10 + 3.14159)
    v5 := make(map[uint32]uint32)
    v6 := []uint32{}
    var v13 uint32 = 743
    delete(v1, "a")
    if (!true) {
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(fi14, fi14)
            v15 := GenS2{F0: "\n\t", F1: 624, F2: true}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v4, v2, v3)
            delete(v1, "a")
            _ = fi16
        }
    }
    var v17 int8 = v3
    switch "hello" {
    case "hello":
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                _ = len(v6)
                var v20 int = (^(fi18))
                _ = fi19
                _ = v20
            }
            if (!true) {
                v21 := GenS3{F0: false}
                _ = v21
            }
            _ = fi18
        }
    }
    var v22 int = 4294967295
    var v23 float32 = v4
    for fi24 := 0; fi24 < 19; fi24++ {
        if true {
            fmt.Println(v23, v17, fi24)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            delete(v1, "hello")
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                if false {
                    delete(v1, "hello")
                }
                v3 -= v17
                _ = fi25
            }
            v26 := v5[(v13 | v13)]
            _ = v26
        }
        var v27 int16 = 1
        _ = fi24
        _ = v27
    }
    v2 = (v2 % (((^(v2))) | 1))
    v22 = (v22 / ((v22) | 1))
    v4 = (v23 - (v4 - v23))
    fmt.Println(v3, v3)
    var v28 int64 = 0
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 bool = ("\n\t" == "\n\t")
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 int64 = v28
            if true {
                v23 += (v23 / (v23 / v4))
            }
            v2 += v2
            _ = fi32
            _ = v33
        }
        _ = fi29
    }
    v23 = v4
    if true {
        var v34 string = "hello"
        switch v34 {
        case "":
            fmt.Println(v13, v2)
        case "foo bar":
            var v35 int = (v22 * (v22 - -981))
            _ = v35
        default:
            fmt.Println(v13, v13, v28)
        }
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 uint32 = v13
            v37 = (v13 * (^(v13)))
            if (!true) {
                v2 = v2
                v38 := v5[(v13 + (v13 | 1))]
                v37 += v37
                _ = v38
            } else {
                switch fi36 {
                case -2147483648:
                    var v39 int16 = -300
                    _ = v39
                case 42:
                    if (false || true) {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            if (1 != -2147483648) {
                                v5[v13] = (v13 >> ((2) & 63))
                                var v41 bool = (!(false && (false || ((v3 < v3) || true))))
                                _ = v41
                            }
                            var v42 string = v34
                            v28 -= (v28 << ((1) & 63))
                            _ = fi40
                            _ = v42
                        }
                    } else {
                        v43 := GenS1{F0: 2, F1: (v22 - (fi36 * (v22 & 32767))), F2: 2}
                        var v44 uint32 = 612
                        if false {
                            var v45 int32 = 2147483647
                            var v46 int32 = v45
                            v47 := GenS2{F0: "hello", F1: 2, F2: true}
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            fmt.Println(v34, fi36, v44)
                            var v48 uint8 = 255
                            var v49 uint8 = uint8(v13)
                            _ = v48
                            _ = v49
                        }
                        _ = v43
                        _ = v44
                    }
                default:
                    var v50 float64 = 3.14159
                    _ = v50
                }
                if false {
                    v51 := GenS4{F0: (v37 ^ 65535), F1: 4294967295, F2: (v2 % ((725) | 1))}
                    switch (-(fi36)) {
                    case 128:
                        v28 = 1
                    default:
                        var v52 int = v22
                        _ = v52
                    }
                    _ = v51
                } else {
                    switch -54 {
                    case 4294967295:
                        var v53 uint64 = 2
                        _ = v53
                    }
                }
            }
            _ = fi36
            _ = v37
        }
        _ = v34
    } else {
        v54 := GenS3{F0: (true || false)}
        if (2 != 127) {
            fmt.Println(v17, v4, v22)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v55 int64 = v28
            if len(v6) > 0 {
                v6[0] = v13
            }
            _ = v55
        }
        v56 := GenS4{F0: (^(v13)), F1: 524, F2: v2}
        _ = v54
        _ = v56
    }
    var v57 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
