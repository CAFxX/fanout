package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 uint8; F2 int32 }
type GenS7 struct { F0 int32; F1 int; F2 int32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 42
    var v3 int16 = 32767
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: 255}
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 42
    fmt.Println(v1, v3)
    if (v6 != (v6 << ((18446744073709551615) & 63))) {
        fmt.Println(v6)
        if ((42 != -2147483648) && (2 > 65535)) {
            var v7 uint16 = uint16(v2)
            switch (("a" + "") + "foo bar") {
            case "hello":
                v8 := GenS1{F0: 255}
                _ = v8
            case "a":
                fmt.Println(v1, v7)
            }
            _ = v7
        } else {
            fmt.Println(v3)
        }
        v9 := GenS2{F0: -47}
        _ = v9
    } else {
        switch 120 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 3; fi10++ {
                switch 127 {
                case 0:
                    v11 := GenS3{F0: 402}
                    _ = v11
                }
                v6 += (v6 ^ 42)
                var v12 float64 = 1.5
                _ = fi10
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                v2 -= (v2 / (((v2 << ((uint(v6)) & 63))) | 1))
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := []int{fi13, 32768, fi13}
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi14
                    _ = v15
                }
                v2 = (v2 ^ (v2 + 42))
            } else {
                v16 := []int{(fi13 & 128), 128, int(v6), fi13}
                _ = v16
            }
            _ = fi13
        }
        switch ("foo bar" + (("" + "\n\t") + "hello")) {
        case "\n\t":
            v2 += 1
        case "":
            v17 := []int{1, -9223372036854775808, -303}
            _ = v17
        case "a":
            var v18 int8 = -1
            _ = v18
        default:
            switch -525 {
            case 127:
                for fi19 := 0; fi19 < 11; fi19++ {
                    if (!(true || (true && false))) {
                        for fi20 := 0; fi20 < 13; fi20++ {
                            switch (fi20 << ((2) & 63)) {
                            case 2147483647:
                                if false {
                                    v21 := GenS4{F0: v2}
                                    _ = v21
                                }
                            case -2147483648:
                                v1 += (v1 * 42)
                                var v22 float32 = (3.14159 + 3.14159)
                                switch (("foo bar" + ("a" + ("\n\t" + ""))) + (((("a" + "\n\t") + "") + "a") + "a")) {
                                case "hello":
                                    if (-2147483648 != 615) {
                                        var v23 uint64 = v2
                                        v3 = -1
                                        _ = v23
                                    }
                                case "":
                                    v24 := []int{fi19, (fi20 * fi20)}
                                    _ = v24
                                default:
                                    if ((1.5 <= (v22 - 3.14159)) || false) {
                                        switch fi20 {
                                        case 32767:
                                            v1 += (^(v1))
                                        case 311:
                                            v3 -= v3
                                        default:
                                            if (!(v6 != v6)) {
                                                var v25 int64 = (-(v1))
                                                v22 += v22
                                                v1 -= v1
                                                _ = v25
                                            } else {
                                                fmt.Println(fi19)
                                            }
                                        }
                                    } else {
                                        v22 += (1e10 - (1.5 * 1e10))
                                    }
                                }
                                var v26 int16 = v3
                                _ = v22
                                _ = v26
                            }
                            var v27 int8 = 1
                            _ = fi20
                            _ = v27
                        }
                        switch "hello" {
                        case "hello":
                            var v28 uint8 = 42
                            _ = v28
                        }
                        v3 = (v3 | v3)
                    }
                    var v29 float32 = (0.1 / 1.5)
                    fmt.Println(fi19, v6, fi19)
                    _ = fi19
                    _ = v29
                }
            default:
                v3 -= v3
            }
        }
    }
    v30 := GenS5{F0: (true && (0.1 <= 0.1))}
    for fi31 := 0; fi31 < 6; fi31++ {
        v1 += v1
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint = 2
            if false {
                var v34 int16 = (-(v3))
                v1 = 0
                var v35 int32 = -2147483648
                v33 = v33
                v1 = (v1 / ((v1) | 1))
                _ = v34
                _ = v35
            }
            v36 := GenS4{F0: v2}
            _ = fi32
            _ = v33
            _ = v36
        }
        _ = fi31
    }
    var v37 int32 = -1
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 int = int(v6)
        v39 = fi38
        _ = fi38
        _ = v39
    }
    fmt.Println(v6)
    v40 := []int{127, 32767, -9223372036854775808}
    switch 484 {
    case 458:
        fmt.Println(v2, v1)
    case 997:
        v1 -= (v1 % ((v1) | 1))
    case 32767:
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    }
    switch -310 {
    case -111:
        var v41 int64 = (v1 + (v1 / (((v1 >> ((536) & 63))) | 1)))
        _ = v41
    case 265:
        if (-298 == 2147483647) {
            if ("hello" <= "foo bar") {
                var v42 int = 4294967295
                _ = v42
            } else {
                if (int64(v37) > v1) {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        for fi43 := 0; fi43 < 3; fi43++ {
                            if false {
                                fmt.Println(v1, v3)
                                v44 := GenS6{F0: v37, F1: 66, F2: (v37 * -715)}
                                v45 := GenS7{F0: 42, F1: fi43, F2: v37}
                                _ = v44
                                _ = v45
                            }
                            var v46 uint16 = (v6 >> ((823) & 63))
                            _ = fi43
                            _ = v46
                        }
                    case -487:
                        var v47 float32 = 0.0
                        _ = v47
                    case -756:
                        var v48 uint8 = 128
                        _ = v48
                    default:
                        fmt.Println(v3)
                    }
                    fmt.Println(v2, v2)
                    var v49 bool = (!false)
                    _ = v49
                }
                var v50 float32 = 3.14159
                for fi51 := 0; fi51 < 1; fi51++ {
                    var v52 uint32 = 4294967295
                    if ((^(v2)) < 1) {
                        v50 += v50
                        if (int64(v52) == (v1 & 42)) {
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                            if false {
                                var v53 int32 = v37
                                var v54 uint64 = (v2 >> ((65535) & 63))
                                fmt.Println(v1, v54, v3)
                                _ = v53
                                _ = v54
                            } else {
                                var v55 string = "foo bar"
                                _ = v55
                            }
                        }
                    } else {
                        var v56 int8 = 127
                        for fi57 := 0; fi57 < 12; fi57++ {
                            fmt.Println(v50)
                            fmt.Println(fi57, v2, v37)
                            v40 = append(v40, -2147483648)
                            _ = fi57
                        }
                        _ = v56
                    }
                    _ = fi51
                    _ = v52
                }
                _ = v50
            }
            var v58 int = -101
            var v59 float32 = float32(float64(-2.25))
            _ = v58
            _ = v59
        } else {
            var v60 uint = 640
            fmt.Println(v2, v3, v6)
            _ = v60
        }
    case 919:
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v40
    _ = v6
}
