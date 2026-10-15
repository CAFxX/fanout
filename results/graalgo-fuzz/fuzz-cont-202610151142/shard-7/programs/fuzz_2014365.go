package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]string)
    v3 := []int{32768, -9223372036854775808, 1, 4294967295}
    var v4 uint = 42
    var v5 uint16 = 32767
    v6 := GenS1{F0: "", F1: -32768, F2: (v1 * -1)}
    switch 127 {
    case 354:
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                if (true && (((("\n\t" + ("a" + (("foo bar" + "hello") + "hello"))) + "") <= "a") && (true && false))) {
                    switch (^(fi7)) {
                    case -555:
                        _ = len(v2)
                    case 58:
                        switch (fi7 * -760) {
                        case 128:
                            v1 = v1
                            var v9 uint64 = 501
                            var v10 float32 = float32(1.5)
                            _ = v10
                            _ = v9
                        case -274:
                            var v11 int64 = (^(v1))
                            _ = v11
                        case 991:
                            switch (fi7 + fi8) {
                            case 32767:
                                var v12 uint16 = (v5 >> ((v4) & 63))
                                _ = v12
                            case -9223372036854775808:
                                if (4294967295 >= v4) {
                                    var v13 bool = true
                                    _ = v13
                                }
                            }
                        }
                    case -586:
                        fmt.Println(v5, fi8)
                        switch ("a" + "a") {
                        case "":
                            v14 := GenS1{F0: (("" + "foo bar") + "a"), F1: -1, F2: (v1 * (v1 + 534))}
                            fmt.Println(fi7, fi8, v1)
                            _ = v14
                        }
                        if (true && false) {
                            var v15 int64 = v1
                            _ = v15
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v16 uint32 = 42
                            _ = v16
                        case "hello":
                            var v17 int8 = 71
                            _ = v17
                        default:
                            v18 := v2[""]
                            _ = v18
                        }
                    default:
                        v4 -= 4294967295
                    }
                    var v19 uint8 = uint8(v4)
                    _ = v19
                } else {
                    var v20 uint16 = v5
                    _ = v20
                }
                v4 = v4
                var v21 int16 = 168
                _ = fi8
                _ = v21
            }
            v4 -= 255
            v3 = append(v3, -50)
            _ = fi7
        }
    }
    if (42 != 9223372036854775807) {
        var v22 uint32 = 0
        var v23 int64 = 0
        if (!true) {
            if ((true || ((^(v5)) <= v5)) || true) {
                v2["hello"] = (("\n\t" + ("a" + "a")) + "a")
                if true {
                    var v24 int16 = 1
                    _ = v24
                } else {
                    fmt.Println(v22, v22, v22)
                    var v25 int64 = (^(v23))
                    _ = v25
                }
                var v26 int32 = 42
                _ = v26
            } else {
                var v27 uint64 = 982
                _ = len(v2)
                _ = v27
            }
            fmt.Println(v4, v22, v22)
            var v28 int16 = 0
            _ = v28
        }
        _ = v22
        _ = v23
    } else {
        fmt.Println(v5)
        if false {
            var v29 uint = (v4 % ((255) | 1))
            _ = v29
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 int16 = 42
                var v32 float64 = -2.25
                _ = fi30
                _ = v31
                _ = v32
            }
            if true {
                fmt.Println(v1, v4, v5)
            } else {
                v33 := GenS1{F0: "\n\t", F1: 32767, F2: (^(v1))}
                v5 -= (v5 << ((v4) & 63))
                _ = v33
            }
        }
    }
    if true {
        for fi34 := 0; fi34 < 12; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                v4 += 189
                v36 := GenS2{F0: 825, F1: 18446744073709551615}
                var v37 int8 = 42
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
        if ((!(true && false)) && true) {
            var v38 float64 = ((-2.25 + 1.5) / 1.5)
            var v39 bool = ((true && (2147483647 != 1)) || (!(42 == -128)))
            v5 = v5
            _ = v38
            _ = v39
        }
        fmt.Println(v5)
    } else {
        v5 -= 255
        v40 := GenS2{F0: 0, F1: 2}
        var v41 int64 = (v1 & (v1 * 9223372036854775807))
        _ = v40
        _ = v41
    }
    v42 := GenS2{F0: 545, F1: 1}
    v5 = (v5 / (((v5 - 32767)) | 1))
    v2["foo bar"] = ""
    v43 := GenS2{F0: 972, F1: 273}
    for fi44 := 0; fi44 < 19; fi44++ {
        v2[("foo bar" + ("\n\t" + "\n\t"))] = "foo bar"
        if false {
            var v45 uint32 = 216
            for fi46 := 0; fi46 < 1; fi46++ {
                if len(v3) > 0 {
                    v3[0] = fi44
                }
                v1 -= (^(v1))
                var v47 float64 = -2.25
                _ = fi46
                _ = v47
            }
            v48 := GenS3{F0: -577}
            _ = v45
            _ = v48
        }
        _ = fi44
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        v1 = (v1 % ((-9223372036854775808) | 1))
        if true {
            fmt.Println(v1)
        }
        v5 -= (^(v5))
        _ = fi49
    }
    var v50 int64 = -1
    var v51 float64 = ((-2.25 - (3.14159 / ((1e10 / (3.14159 - (1.5 - ((1e10 / 3.14159) - -2.25)))) / (1e10 - (float64(((-2.25 * 1.5) / 3.14159)) - 0.0))))) / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
