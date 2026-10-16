package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 uint64) uint8 {
    if false {
        v5 := []int{9223372036854775807, -2147483648, -90}
        _ = v5
    } else {
        fmt.Println(p0)
        p0 = (p0 + p0)
        var v6 uint16 = 2
        _ = v6
    }
    v7 := []int{-1, -290, 42, -54}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint = 4294967295
    if len(v7) > 0 {
        v7[0] = 32767
    }
    _ = p0
    _ = v7
    _ = v8
    return 30
}
func genFn1(p0 uint16, p1 uint16, p2 uint16) uint16 {
    p2 = (^(p1))
    fmt.Println(p0, p2, p0)
    fmt.Println(p0, p0, p0)
    if false {
        p0 = p1
        var v9 string = ""
        _ = v9
    } else {
        v10 := GenS1{F0: (p0 - p1), F1: 9223372036854775807, F2: -921}
        _ = v10
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]string)
    var v3 float64 = (float64(float32(-2.25)) - (0.1 + (1e10 - (0.0 / 1.5))))
    var v4 bool = (v1 > v1)
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 = 22
        switch fi11 {
        case 128:
            var v12 int64 = 0
            _ = v12
        }
        _ = fi11
    }
    v13 := v2[375]
    for fi14 := 0; fi14 < 16; fi14++ {
        if (!v4) {
            v4 = v4
            switch fi14 {
            case 128:
                var v15 uint32 = 2
                _ = v15
            case 127:
                delete(v2, fi14)
            }
        } else {
            var v16 int = 288
            var v17 uint8 = 2
            if v4 {
                fmt.Println(v1)
                switch v13 {
                case "\n\t":
                    var v18 uint8 = 1
                    _ = v18
                case "foo bar":
                    v17 = genFn0(4294967295)
                }
            }
            _ = v16
            _ = v17
        }
        v1 = (v1 << ((42) & 63))
        _ = fi14
    }
    switch v13 {
    case "a":
        switch "" {
        case "hello":
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 uint16 = 1
                _ = fi19
                _ = v20
            }
        case "":
            fmt.Println(v13, v4, v13)
        case "a":
            var v21 int32 = 2147483647
            _ = v21
        }
    default:
        v22 := GenS2{F0: -197, F1: 272}
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v3 += v3
        var v24 int = (fi23 * 0)
        _ = fi23
        _ = v24
    }
    v4 = true
    if ((1 >= -769) || v4) {
        v4 = v4
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := []int{(fi25 ^ fi25), fi25, (fi25 << ((4294967295) & 63)), (fi25 - -537)}
            _ = fi25
            _ = v26
        }
    } else {
        v27 := GenS2{F0: 0, F1: 305}
        _ = v27
    }
    var v28 uint8 = 231
    v29 := GenS2{F0: -2147483648, F1: 255}
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 uint = 255
        v4 = ((105 == -9223372036854775808) && v4)
        _ = fi30
        _ = v31
    }
    v32 := GenS1{F0: 0, F1: 42, F2: 9223372036854775807}
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int16 = 604
        _ = fi33
        _ = v34
    }
    var v35 uint32 = 2
    if false {
        v36 := v2[128]
        _ = v36
    }
    v3 += float64(1e10)
    if (v35 > (v35 & (v35 >> ((371) & 63)))) {
        for fi37 := 0; fi37 < 11; fi37++ {
            if (42 < -878) {
                var v38 bool = false
                switch -316 {
                case -392:
                    for fi39 := 0; fi39 < 9; fi39++ {
                        v40 := GenS1{F0: 65535, F1: 0, F2: -1}
                        _ = fi39
                        _ = v40
                    }
                case -536:
                    if (0.0 >= -2.25) {
                        for fi41 := 0; fi41 < 11; fi41++ {
                            if (true && false) {
                                fmt.Println(v38, v28)
                                if ((v13 + v13) == (v13 + ("" + v13))) {
                                    v4 = v4
                                    var v42 uint64 = 42
                                    v3 -= 0.1
                                    fmt.Println(fi41)
                                    var v43 uint8 = 128
                                    _ = v42
                                    _ = v43
                                } else {
                                    delete(v2, 642)
                                    var v44 uint16 = 2
                                    var v45 float64 = v3
                                    _ = v44
                                    _ = v45
                                }
                                if v4 {
                                    var v46 float32 = (1e10 / (-2.25 * (1.5 - 1e10)))
                                    _ = v46
                                } else {
                                    if ((!false) || false) {
                                        var v47 int = (fi37 + fi37)
                                        _ = v47
                                    } else {
                                        v2[0] = "a"
                                        v13 = v13
                                        v28 = 155
                                    }
                                    var v48 bool = v38
                                    _ = v48
                                }
                            }
                            v49 := GenS1{F0: 2, F1: 42, F2: -547}
                            _ = fi41
                            _ = v49
                        }
                        var v50 uint64 = 0
                        _ = v50
                    } else {
                        switch (fi37 ^ fi37) {
                        case 199:
                            if true {
                                v3 -= v3
                            } else {
                                switch v13 {
                                case "\n\t":
                                    fmt.Println(v35, v38)
                                case "foo bar":
                                    for fi51 := 0; fi51 < 9; fi51++ {
                                        switch v13 {
                                        case "\n\t":
                                            fmt.Println(v35, v4, v13)
                                        default:
                                            var v52 uint32 = 870
                                            _ = v52
                                        }
                                        _ = fi51
                                    }
                                case "a":
                                    switch (v13 + (v13 + "")) {
                                    case "\n\t":
                                        v3 = v3
                                    default:
                                        var v53 bool = true
                                        _ = v53
                                    }
                                }
                                var v54 bool = v4
                                var v55 bool = (32767 != 2)
                                _ = v54
                                _ = v55
                            }
                        }
                    }
                default:
                    v38 = (v35 >= (v35 % (((v35 & 2147483647)) | 1)))
                }
                _ = v38
            }
            for fi56 := 0; fi56 < 20; fi56++ {
                fmt.Println(v35)
                _ = fi56
            }
            _ = fi37
        }
        switch -642 {
        case -979:
            v57 := GenS1{F0: 32767, F1: 448, F2: 277}
            _ = v57
        case 42:
            v13 = (v13 + ("hello" + (v13 + v13)))
        default:
            delete(v2, -2147483648)
        }
        fmt.Println(v28)
    } else {
        fmt.Println(v28, v35, v13)
        for fi58 := 0; fi58 < 15; fi58++ {
            for fi59 := 0; fi59 < 16; fi59++ {
                v2[1] = "foo bar"
                v60 := []int{(fi59 | 32768), (fi58 >> ((155) & 63)), (fi58 * 621)}
                switch fi59 {
                case 42:
                    var v61 int32 = 42
                    _ = v61
                case 4294967295:
                    delete(v2, fi58)
                }
                _ = fi59
                _ = v60
            }
            _ = fi58
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v35
    _ = v4
}
