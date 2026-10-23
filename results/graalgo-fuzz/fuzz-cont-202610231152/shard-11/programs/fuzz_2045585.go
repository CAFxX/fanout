package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0() uint16 {
    v7 := []int{}
    v7 = append(v7, 535)
    _ = v7
    return 255
}

func main() {
    var v1 bool = true
    var v2 string = "a"
    v3 := []int{548, 128, -508}
    var v4 float32 = float32(3.14159)
    var v5 uint = 0
    v6 := []int{0, 32768}
    if false {
        if v1 {
            var v8 int16 = -13
            _ = v8
        } else {
            fmt.Println(v2)
            _ = len(v3)
        }
        switch 9223372036854775807 {
        case -714:
            var v9 int32 = -1
            _ = v9
        case -849:
            if (v4 <= (float32(0.1) + v4)) {
                switch v2 {
                case "a":
                    v2 = "a"
                }
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 uint8 = 1
                    v11 += 2
                    _ = fi10
                    _ = v11
                }
            } else {
                v2 = v2
            }
        }
        v5 = (v5 & (v5 << ((v5) & 63)))
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 = (v1 || v1)
            v1 = v1
            _ = fi12
        }
        var v13 uint16 = 2
        _ = v13
    }
    v14 := GenS1{F0: 18446744073709551615, F1: 0, F2: 50}
    if ((((float64(0.1) - (0.0 / (1e10 + 0.1))) - 3.14159) - -2.25) != 1.5) {
        var v15 uint64 = 42
        _ = v15
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(v2, v4)
            _ = fi16
        }
        switch (((((v2 + ("foo bar" + v2)) + ("hello" + "")) + "") + "hello") + (("hello" + (((v2 + v2) + "") + v2)) + v2)) {
        case "a":
            v6 = append(v6, 0)
        case "hello":
            v17 := GenS2{F0: 2, F1: -835, F2: 0}
            _ = v17
        case "":
            if false {
                if (-32768 <= 1000) {
                    switch 0 {
                    case 9223372036854775807:
                        v1 = false
                    default:
                        switch 32768 {
                        case 1:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case 128:
                            fmt.Println(v4, v2, v2)
                        default:
                            v18 := GenS1{F0: 42, F1: 4294967295, F2: 42}
                            _ = v18
                        }
                    }
                    _ = len(v6)
                }
                v2 = ("a" + v2)
            } else {
                v19 := GenS1{F0: 18446744073709551615, F1: 1, F2: 110}
                v20 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 42}
                var v21 float64 = (1.5 / -2.25)
                _ = v19
                _ = v20
                _ = v21
            }
        }
        if false {
            fmt.Println(v4)
        }
    }
    v1 = (v1 && v1)
    switch ("foo bar" + v2) {
    case "":
        v22 := GenS3{F0: 1}
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 18; fi23++ {
            v2 = "hello"
            v24 := GenS2{F0: 255, F1: (fi23 ^ (fi23 & (fi23 >> ((0) & 63)))), F2: 0}
            v2 = v2
            _ = fi23
            _ = v24
        }
    case "foo bar":
        if (42 < 607) {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 string = v2
                v5 -= (v5 % ((v5) | 1))
                v2 = (v26 + (v26 + (v2 + v2)))
                _ = fi25
                _ = v26
            }
            v5 = 255
            var v27 int8 = -1
            _ = v27
        }
    }
    var v28 int32 = -2147483648
    v28 = 0
    switch 32767 {
    case 127:
        var v29 float64 = 0.1
        _ = v29
    default:
        var v30 float64 = 0.0
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS4{F0: -32768, F1: 0}
        v2 = (v2 + ((((v2 + v2) + v2) + v2) + "hello"))
        var v33 uint16 = 255
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 bool = v1
    switch v2 {
    case "\n\t":
        var v35 uint16 = 12
        _ = v35
    case "hello":
        v34 = v34
    case "":
        var v36 uint16 = 693
        _ = v36
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 uint16 = 417
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            v28 -= 1
            switch (fi37 | 1) {
            case 2147483647:
                fmt.Println(v5, v2)
            case 1:
                fmt.Println(v2, v2)
            case 298:
                v41 := GenS4{F0: -1, F1: -1}
                fmt.Println(v4)
                v34 = true
                _ = v41
            }
            _ = fi40
        }
        _ = fi37
    }
    if (!v1) {
        v34 = v34
        if v1 {
            fmt.Println(v4)
            v1 = (976 > 241)
        } else {
            v4 = (v4 + float32(float64(3.14159)))
        }
    } else {
        var v42 int64 = 42
        if (true || (true && v34)) {
            for fi43 := 0; fi43 < 15; fi43++ {
                switch v2 {
                case "\n\t":
                    v44 := GenS1{F0: uint64(v5), F1: 1, F2: 119}
                    _ = v44
                case "hello":
                    if false {
                        var v45 uint16 = 830
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v45
                    } else {
                        v42 -= v42
                        if v34 {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v46 := GenS2{F0: 127, F1: -1, F2: 2}
                            if (((v42 | v42) <= v42) || false) {
                                for fi47 := 0; fi47 < 17; fi47++ {
                                    v2 = (((v2 + (("\n\t" + (v2 + v2)) + "a")) + (v2 + v2)) + "")
                                    _ = fi47
                                }
                            }
                            _ = v46
                        } else {
                            switch v2 {
                            case "foo bar":
                                v48 := GenS3{F0: 0}
                                _ = v48
                            }
                        }
                        var v49 string = v2
                        _ = v49
                    }
                default:
                    if ((!v1) && false) {
                        v4 -= -2.25
                        var v50 bool = (1.5 > (-2.25 * 0.0))
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 11; fi51++ {
                            fmt.Println(fi43)
                            _ = fi51
                        }
                        switch fi43 {
                        case 4294967295:
                            var v52 string = "a"
                            _ = v52
                        default:
                            for fi53 := 0; fi53 < 19; fi53++ {
                                var v54 uint = v5
                                _ = fi53
                                _ = v54
                            }
                        }
                    }
                }
                if false {
                    for fi55 := 0; fi55 < 5; fi55++ {
                        v6 = append(v6, fi43)
                        var v56 uint64 = 1
                        _ = fi55
                        _ = v56
                    }
                }
                _ = fi43
            }
            var v57 uint16 = 65535
            _ = v57
        } else {
            var v58 int8 = -2
            v34 = v1
            if false {
                var v59 int8 = v58
                _ = len(v6)
                _ = v59
            } else {
                var v60 float64 = 3.14159
                switch -2147483648 {
                case -818:
                    if ((v28 >> (((v5 - v5)) & 63)) > v28) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v61 int32 = v28
                        if len(v3) > 0 {
                            v3[0] = -819
                        }
                        _ = v61
                    }
                case 863:
                    _ = len(v3)
                default:
                    v4 += 1.5
                }
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                _ = v60
            }
            _ = v58
        }
        var v62 float32 = ((v4 * v4) / (((v4 * -2.25) + (0.1 - v4)) * (v4 / (-2.25 * v4))))
        _ = v42
        _ = v62
    }
    v4 = 0.1
    v6 = append(v6, 197)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
