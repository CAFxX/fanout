package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0(p0 int) uint64 {
    var v4 uint16 = 1
    v4 -= (v4 | (v4 + v4))
    fmt.Println(v4, v4)
    _ = p0
    _ = v4
    return 270
}

func main() {
    var v1 int64 = 429
    var v2 uint32 = 0
    var v3 bool = (2 > 1)
    for fi5 := 0; fi5 < 12; fi5++ {
        if (v3 || true) {
            v6 := GenS1{F0: 42, F1: 231, F2: (v2 * v2)}
            _ = v6
        }
        v7 := GenS1{F0: 128, F1: 0, F2: (v2 * v2)}
        var v8 uint64 = 2
        _ = fi5
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint = 18446744073709551615
        v11 := []int{(fi9 * (fi9 | fi9)), (fi9 & fi9), fi9, (fi9 >> (((v10 - v10)) & 63))}
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 11; fi12++ {
            if ((v3 && ((((v3 || (!(!((v1 >> ((2) & 63)) < -9223372036854775808)))) && ((!true) || (v3 || v3))) && v3) || true)) || v3) {
                var v13 float64 = ((-2.25 * (0.1 - (float64(float32(1e10)) / 1e10))) / 1.5)
                _ = v13
            } else {
                fmt.Println(v1)
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(fi14, fi14, v1)
                    _ = fi14
                }
                v15 := []int{9223372036854775807, (-(fi12)), fi12}
                _ = v15
            }
            _ = fi12
        }
    }
    switch 9223372036854775807 {
    case 2147483647:
        v1 = (v1 * v1)
    }
    var v16 int64 = (v1 ^ v1)
    fmt.Println(v1, v3, v3)
    var v17 uint16 = 65535
    for fi18 := 0; fi18 < 4; fi18++ {
        if (0 == -95) {
            if v3 {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v2 = 4294967295
                    _ = fi19
                }
                v1 -= v1
                if ((65535 != 42) && v3) {
                    v3 = true
                    v17 += (v17 | (^(v17)))
                }
            }
            fmt.Println(v16, v16, v17)
            v20 := []int{fi18, (fi18 / ((9223372036854775807) | 1))}
            _ = v20
        } else {
            var v21 uint64 = 42
            _ = v21
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float64 = float64(0.1)
            v24 := GenS1{F0: 44, F1: 2, F2: (^(v2))}
            var v25 int = (fi22 >> ((2) & 63))
            var v26 string = "a"
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi18
    }
    v3 = v3
    for fi27 := 0; fi27 < 6; fi27++ {
        v16 += v1
        v28 := GenS1{F0: 1, F1: 170, F2: (v2 & v2)}
        _ = fi27
        _ = v28
    }
    v29 := GenS1{F0: 86, F1: 42, F2: (^(v2))}
    for fi30 := 0; fi30 < 15; fi30++ {
        if (uint8(v17) <= 2) {
            switch "\n\t" {
            case "":
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (false && v3) {
                        if v3 {
                            v32 := GenS2{F0: (-(v1))}
                            _ = v32
                        }
                    } else {
                        var v33 int8 = 0
                        _ = v33
                    }
                    _ = fi31
                }
            case "\n\t":
                var v34 uint8 = 154
                _ = v34
            case "foo bar":
                switch fi30 {
                case 0:
                    v1 -= v16
                default:
                    v35 := []int{(fi30 * fi30)}
                    _ = v35
                }
            default:
                v36 := GenS2{F0: -1}
                _ = v36
            }
        } else {
            if (-2147483648 <= 82) {
                v2 = 2
                v37 := []int{fi30, 127, fi30, fi30}
                _ = v37
            } else {
                v1 = (v16 % ((v1) | 1))
            }
        }
        fmt.Println(v1)
        if (-404 > (v16 * 9223372036854775807)) {
            for fi38 := 0; fi38 < 9; fi38++ {
                v39 := []int{fi38}
                if true {
                    if true {
                        v39 = append(v39, fi30)
                        if ((("" + "hello") + "a") != "a") {
                            if (1e10 >= 1.5) {
                                v40 := GenS3{F0: "foo bar", F1: 0}
                                _ = v40
                            }
                            if len(v39) > 0 {
                                v39[0] = fi30
                            }
                            switch ("a" + ("a" + "\n\t")) {
                            case "\n\t":
                                switch (fi30 * (fi30 + int(v1))) {
                                case 32767:
                                    var v41 string = ("foo bar" + ("" + ("hello" + ("foo bar" + ""))))
                                    v42 := GenS2{F0: (v1 ^ v16)}
                                    if (-1 <= 9223372036854775807) {
                                        fmt.Println(fi38, v41, v16)
                                        var v43 string = v41
                                        _ = v43
                                    }
                                    _ = v41
                                    _ = v42
                                default:
                                    v16 -= (v1 / ((9223372036854775807) | 1))
                                }
                            }
                            v1 = (-(v16))
                        }
                        var v44 uint8 = 1
                        _ = v44
                    } else {
                        v17 = v17
                    }
                    v16 -= (v16 * (v1 * (v1 ^ (-(v1)))))
                } else {
                    if v3 {
                        var v45 int64 = v1
                        v46 := GenS2{F0: v1}
                        v45 = -1
                        _ = v45
                        _ = v46
                    }
                    v39 = append(v39, (^(fi38)))
                }
                v47 := GenS4{F0: 946, F1: 0, F2: (v2 + (v2 << ((1) & 63)))}
                _ = fi38
                _ = v39
                _ = v47
            }
            var v48 int32 = -2147483648
            _ = v48
        } else {
            var v49 uint16 = (v17 << ((0) & 63))
            _ = v49
        }
        _ = fi30
    }
    var v50 int64 = int64(v2)
    for fi51 := 0; fi51 < 16; fi51++ {
        fmt.Println(v1, v16, v1)
        _ = fi51
    }
    v52 := GenS3{F0: ("foo bar" + "a"), F1: 612}
    for fi53 := 0; fi53 < 15; fi53++ {
        v3 = v3
        var v54 float64 = 3.14159
        for fi55 := 0; fi55 < 19; fi55++ {
            switch (fi55 - fi53) {
            case 32767:
                if v3 {
                    fmt.Println(v1, fi53)
                } else {
                    v56 := []int{fi55, (fi53 >> ((1) & 63)), -871}
                    _ = v56
                }
            }
            _ = fi55
        }
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v50
    _ = v52
}
