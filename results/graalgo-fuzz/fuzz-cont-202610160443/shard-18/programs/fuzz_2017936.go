package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS6 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 255
    var v4 uint32 = 4294967295
    if (true && (((-2.25 - 1e10) * (1.5 - 1e10)) >= (0.0 + 0.0))) {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 uint = 192
            v4 += v4
            var v7 float64 = 0.0
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 42, F1: 1, F2: 18446744073709551615}
        _ = v8
    } else {
        var v9 float64 = 3.14159
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int64 = 1
            var v12 int8 = 100
            _ = fi10
            _ = v11
            _ = v12
        }
        switch ("a" + ("" + ("a" + ((("foo bar" + "") + "\n\t") + (("\n\t" + "foo bar") + (("hello" + ("foo bar" + "\n\t")) + "a")))))) {
        case "a":
            v9 = v9
        case "foo bar":
            switch 9223372036854775807 {
            case -1:
                fmt.Println(v3)
            case 32768:
                if false {
                    v13 := GenS2{F0: 2, F1: (true || (-9223372036854775808 <= -926))}
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v3 = 128
                        switch fi14 {
                        case -333:
                            if true {
                                v15 := v1[(v4 - (v4 % ((v4) | 1)))]
                                var v16 float32 = ((0.0 / 1.0) / 1.5)
                                v17 := GenS3{F0: 42, F1: v4}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        case -150:
                            switch (fi14 & fi14) {
                            case 4294967295:
                                fmt.Println(v9)
                            case 9223372036854775807:
                                if (uint(v4) == 0) {
                                    v4 = 2147483648
                                    for fi18 := 0; fi18 < 2; fi18++ {
                                        v3 = v3
                                        var v19 string = "hello"
                                        v20 := []int{fi14, fi14, fi18, (fi18 & 415)}
                                        fmt.Println(v19, v4, v4)
                                        if true {
                                            var v21 bool = false
                                            v22 := GenS2{F0: 1, F1: ((^(v3)) >= v3)}
                                            var v23 bool = false
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    _ = len(v2)
                                }
                            case -9223372036854775808:
                                delete(v2, (v4 | 65535))
                            default:
                                v24 := GenS4{F0: -9223372036854775808, F1: (^(v3)), F2: true}
                                _ = v24
                            }
                        case 2147483647:
                            v4 += (v4 << ((255) & 63))
                        }
                        _ = fi14
                    }
                    v3 = v3
                    _ = v13
                }
            case 2147483647:
                if (true && true) {
                    v25 := v1[v4]
                    _ = v25
                } else {
                    v26 := GenS4{F0: -251, F1: (v3 - 2), F2: (true && (2 >= (^(v4))))}
                    v9 = (((-2.25 * 3.14159) / v9) + v9)
                    v27 := GenS4{F0: 42, F1: v3, F2: (((((1.5 * ((1.5 - -2.25) + 3.14159)) / 0.1) / -2.25) != 1e10) || true)}
                    _ = v26
                    _ = v27
                }
            default:
                v28 := GenS1{F0: 32767, F1: 0, F2: 508}
                _ = v28
            }
        default:
            v29 := GenS2{F0: 535, F1: (32767 == 344)}
            _ = v29
        }
        _ = v9
    }
    v30 := GenS5{F0: 876, F1: int16(v3), F2: 88}
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint8 = v3
        if (255 < 1) {
            var v33 float64 = (((3.14159 * 1.5) * (3.14159 - 3.14159)) / (-2.25 / -2.25))
            var v34 string = ""
            if (0 > 0) {
                if (127 >= 1) {
                    v1[(v4 >> ((0) & 63))] = v34
                } else {
                    v35 := GenS6{F0: (("a" + ("foo bar" + v34)) + v34), F1: 1}
                    _ = v35
                }
                if true {
                    v36 := []int{fi31, (fi31 >> ((660) & 63)), (fi31 / ((fi31) | 1))}
                    switch (-(fi31)) {
                    case 999:
                        for fi37 := 0; fi37 < 7; fi37++ {
                            var v38 uint64 = 0
                            _ = fi37
                            _ = v38
                        }
                    case 15:
                        var v39 uint8 = v3
                        _ = v39
                    case 2147483647:
                        for fi40 := 0; fi40 < 17; fi40++ {
                            _ = len(v36)
                            var v41 string = v34
                            v3 -= v3
                            _ = fi40
                            _ = v41
                        }
                    }
                    var v42 string = v34
                    _ = v36
                    _ = v42
                } else {
                    v43 := v1[511]
                    fmt.Println(v32, v33, v34)
                    _ = v43
                }
            } else {
                fmt.Println(v33, v32)
                if (42 != 372) {
                    v4 += v4
                } else {
                    v34 = ((v34 + ("\n\t" + v34)) + v34)
                    v33 -= v33
                    fmt.Println(v4, v3)
                }
                v33 = float64((0.0 + 0.0))
            }
            _ = v33
            _ = v34
        } else {
            var v44 int16 = 1
            _ = v44
        }
        var v45 uint16 = 65535
        _ = fi31
        _ = v32
        _ = v45
    }
    v4 = (v4 ^ v4)
    var v46 int = 128
    v47 := v2[(v4 % ((2147483647) | 1))]
    v3 = (v3 * 128)
    if true {
        v47 = v4
    } else {
        v3 = (^(v3))
    }
    fmt.Println(v3, v47, v46)
    for fi48 := 0; fi48 < 17; fi48++ {
        v4 -= 2
        v3 = v3
        _ = fi48
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        fmt.Println(v47, v4)
    case "a":
        var v49 float32 = 3.14159
        _ = v49
    }
    v50 := v2[4294967295]
    var v51 int32 = -1
    v50 = (v50 << ((1) & 63))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v47
    _ = v50
    _ = v51
}
