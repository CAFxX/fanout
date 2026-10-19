package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int32 = -220
    var v2 int16 = 1000
    var v3 int64 = -805
    v4 := []int{-2147483648, 128}
    var v5 string = ("foo bar" + ("hello" + "hello"))
    switch ((v5 + "\n\t") + "hello") {
    case "hello":
        v1 += v1
    case "a":
        if len(v4) > 0 {
            v4[0] = -209
        }
    case "":
        v2 += (v2 << ((45) & 63))
    default:
        v5 = "hello"
    }
    var v6 float32 = float32((0.1 * 1e10))
    v1 += v1
    var v7 int8 = 0
    if false {
        var v8 bool = false
        var v9 int64 = (v3 % ((v3) | 1))
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 uint64 = 272
            v5 = v5
            var v12 int = 2147483647
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
        _ = v9
    }
    v6 = float32(float64(v6))
    if (((v6 >= v6) && (1.5 < (1e10 * (1e10 / 3.14159)))) && true) {
        v13 := GenS1{F0: 127, F1: 483, F2: ("a" + (((((v5 + v5) + (("\n\t" + (v5 + "hello")) + (v5 + (v5 + "foo bar")))) + "hello") + "") + v5))}
        if true {
            switch v5 {
            case "":
                if true {
                    var v14 uint8 = 2
                    var v15 uint8 = (v14 | 183)
                    _ = v14
                    _ = v15
                }
            case "hello":
                v4 = append(v4, 9223372036854775807)
            case "\n\t":
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 int16 = 32767
                    switch fi16 {
                    case 193:
                        v18 := GenS1{F0: 1, F1: 643, F2: ("foo bar" + v5)}
                        _ = v18
                    default:
                        v19 := GenS1{F0: 42, F1: 42, F2: ""}
                        _ = v19
                    }
                    if true {
                        v20 := GenS1{F0: 2, F1: 1, F2: v5}
                        v7 -= v7
                        if true {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v6 = v6
                                var v22 uint16 = 65535
                                switch (^(fi16)) {
                                case -1:
                                    var v23 uint = 702
                                    _ = len(v4)
                                    _ = v23
                                default:
                                    v24 := GenS1{F0: 2, F1: 42, F2: ((("hello" + (("foo bar" + ((v5 + "") + (v5 + (v5 + ("a" + "hello"))))) + "foo bar")) + v5) + v5)}
                                    _ = v24
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi21
                                _ = v22
                            }
                            if ((v2 + v2) >= (v17 << ((4294967295) & 63))) {
                                var v25 int64 = -1
                                _ = v25
                            } else {
                                var v26 bool = (677 == 255)
                                v4 = append(v4, (fi16 << ((1) & 63)))
                                _ = v26
                            }
                        }
                        _ = v20
                    } else {
                        v1 += v1
                        switch "" {
                        case "\n\t":
                            if true {
                                var v27 string = v5
                                if (0 <= 42) {
                                    fmt.Println(v17, v5)
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    var v28 float64 = ((1e10 * 1e10) + 0.1)
                                    _ = v28
                                }
                                v1 = -210
                                _ = v27
                            } else {
                                var v29 uint64 = 4294967295
                                if true {
                                    if (!(!(false && (1 != 255)))) {
                                        fmt.Println(v3, v3)
                                        v2 = (v17 >> ((0) & 63))
                                        v30 := GenS1{F0: 255, F1: 18446744073709551615, F2: "a"}
                                        _ = v30
                                    } else {
                                        v3 = (v3 - (v3 * int64(fi16)))
                                    }
                                }
                                v3 = (^(v3))
                                _ = v29
                            }
                        case "":
                            v31 := GenS1{F0: 0, F1: 2, F2: (v5 + v5)}
                            _ = v31
                        }
                        v32 := GenS2{F0: 65535, F1: 2}
                        _ = v32
                    }
                    _ = fi16
                    _ = v17
                }
            default:
                for fi33 := 0; fi33 < 17; fi33++ {
                    v7 -= int8(fi33)
                    fmt.Println(v7)
                    if (1 <= 1) {
                        var v34 float32 = v6
                        switch fi33 {
                        case 127:
                            var v35 bool = (!false)
                            _ = v35
                        default:
                            v34 = (float32((1.5 + float64((v34 / v6)))) + (float32(0.0) + (v34 - float32(0.1))))
                        }
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v36 uint8 = 127
                            _ = v36
                        } else {
                            v7 = v7
                            switch ("foo bar" + v5) {
                            case "foo bar":
                                v4 = append(v4, fi33)
                            }
                            for fi37 := 0; fi37 < 4; fi37++ {
                                v7 = v7
                                var v38 float64 = 3.14159
                                var v39 uint8 = 127
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = v34
                    }
                    _ = fi33
                }
            }
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS1{F0: 127, F1: 0, F2: v5}
                var v42 float64 = 0.0
                if (!true) {
                    _ = len(v4)
                    v7 -= (v7 & -128)
                    v42 += v42
                } else {
                    for fi43 := 0; fi43 < 16; fi43++ {
                        var v44 bool = (6 >= 0)
                        if len(v4) > 0 {
                            v4[0] = (fi40 % ((128) | 1))
                        }
                        _ = fi43
                        _ = v44
                    }
                }
                _ = fi40
                _ = v41
                _ = v42
            }
        } else {
            var v45 int8 = v7
            switch v5 {
            case "a":
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            case "foo bar":
                v46 := GenS3{F0: 482, F1: (v1 - v1)}
                _ = v46
            case "":
                for fi47 := 0; fi47 < 1; fi47++ {
                    v2 = v2
                    var v48 uint = 1
                    for fi49 := 0; fi49 < 6; fi49++ {
                        fmt.Println(v2, v48)
                        fmt.Println(v48, fi47, v1)
                        _ = fi49
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                v50 := GenS4{F0: 4294967295}
                _ = v50
            }
            _ = v45
        }
        _ = v13
    }
    fmt.Println(v6, v6)
    v51 := GenS5{F0: 32767}
    switch -9223372036854775808 {
    case 0:
        v52 := GenS5{F0: 9223372036854775807}
        _ = v52
    default:
        v3 = v3
    }
    v53 := GenS6{F0: 4294967295}
    v54 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v53
    _ = v54
    _ = v6
    _ = v7
}
