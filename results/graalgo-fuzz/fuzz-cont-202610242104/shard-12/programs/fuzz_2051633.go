package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 int32 = -1
    var v2 int64 = 202
    var v3 float32 = 1e10
    v4 := []uint32{65535}
    var v5 string = ((("\n\t" + "foo bar") + ((((("\n\t" + ((("foo bar" + "hello") + "") + (("\n\t" + (("\n\t" + ("hello" + "")) + "\n\t")) + ((("hello" + "a") + ((("\n\t" + (("a" + "a") + "foo bar")) + "hello") + "")) + ("foo bar" + "hello"))))) + "") + "a") + "") + "")) + "\n\t")
    _ = len(v4)
    if (false && (2147483647 < v1)) {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 bool = (2147483648 <= 2147483648)
            if v7 {
                v8 := GenS1{F0: ((v5 + ("foo bar" + v5)) + (v5 + (("\n\t" + "hello") + "hello")))}
                v9 := GenS1{F0: v5}
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS2{F0: "foo bar", F1: -1}
                    _ = fi10
                    _ = v11
                }
                var v12 uint = 2
                _ = v12
            }
            _ = fi6
            _ = v7
        }
        v4 = append(v4, 1)
    }
    switch v5 {
    case "\n\t":
        if (!true) {
            for fi13 := 0; fi13 < 18; fi13++ {
                if true {
                    v14 := GenS2{F0: "foo bar", F1: 100}
                    v15 := GenS2{F0: v5, F1: 1}
                    if false {
                        v16 := GenS1{F0: "foo bar"}
                        fmt.Println(v3, v3, v3)
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        fmt.Println(v2, v3, fi13)
                    } else {
                        if true {
                            v17 := GenS3{F0: true}
                            for fi18 := 0; fi18 < 13; fi18++ {
                                switch (fi18 & fi13) {
                                case 127:
                                    v3 = v3
                                }
                                if (int64(v1) < (v2 - v2)) {
                                    fmt.Println(fi18, v2)
                                    _ = len(v4)
                                }
                                _ = fi18
                            }
                            _ = v17
                        }
                    }
                    v19 := GenS3{F0: false}
                    var v20 uint64 = 1
                    _ = v19
                    _ = v20
                }
                fmt.Println(v2, fi13, v2)
                _ = fi13
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                if false {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if false {
                            if (-188 == fi21) {
                                v23 := GenS2{F0: v5, F1: -73}
                                v24 := GenS4{F0: 0, F1: -128, F2: 507}
                                v2 = -9223372036854775808
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(fi22, v2)
                            v25 := GenS5{F0: 82}
                            var v26 float64 = -2.25
                            _ = v25
                            _ = v26
                        }
                        if false {
                            var v27 uint16 = 940
                            if false {
                                v2 -= 42
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                fmt.Println(v27)
                            }
                            _ = v27
                        }
                        _ = fi22
                    }
                }
                var v28 float32 = v3
                fmt.Println(v5, fi21)
                _ = fi21
                _ = v28
            }
        } else {
            v3 -= v3
            var v29 string = v5
            fmt.Println(v29, v2)
            _ = v29
        }
    default:
        var v30 float64 = 3.14159
        _ = v30
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 float32 = (1.5 - -2.25)
            _ = fi31
            _ = v32
        }
        var v33 int32 = 0
        if (true && (127 != 127)) {
            switch (v5 + (v5 + ((("hello" + (v5 + v5)) + v5) + v5))) {
            case "\n\t":
                var v34 int16 = 849
                _ = v34
            case "foo bar":
                var v35 bool = true
                _ = v35
            default:
                v36 := GenS1{F0: ((v5 + "foo bar") + v5)}
                _ = v36
            }
            v3 = -2.25
            var v37 uint16 = 65535
            _ = v37
        }
        _ = v33
    }
    v3 += v3
    for fi38 := 0; fi38 < 13; fi38++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v39 := GenS2{F0: "\n\t", F1: 42}
        v40 := GenS2{F0: (((v5 + "\n\t") + "a") + (v5 + v5)), F1: 106}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 string = ((v5 + (("a" + v5) + "")) + v5)
    v42 := GenS6{F0: 2147483647}
    switch v41 {
    case "":
        var v43 uint = 273
        _ = v43
    case "hello":
        switch 0 {
        case -228:
            var v44 bool = ((true && true) && (true || ((v1 % (((v1 & (v1 | v1))) | 1)) > 653)))
            _ = v44
        case 490:
            if ((v41 + ("" + (v5 + "foo bar"))) <= (v41 + "\n\t")) {
                var v45 uint32 = 1
                var v46 int64 = v2
                _ = v45
                _ = v46
            } else {
                v47 := GenS4{F0: 164, F1: 87, F2: 0}
                v5 = (v41 + ("foo bar" + v41))
                _ = len(v4)
                _ = v47
            }
        case 42:
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        }
    case "foo bar":
        var v48 int8 = 42
        _ = v48
    default:
        switch 127 {
        case -1:
            switch (("a" + v41) + (("" + "a") + ("hello" + ((v41 + v41) + "\n\t")))) {
            case "":
                if (2 > 170) {
                    v49 := GenS5{F0: 2}
                    _ = v49
                }
            case "\n\t":
                var v50 bool = true
                _ = v50
            case "a":
                fmt.Println(v5)
            }
        case 127:
            switch 32767 {
            case -9223372036854775808:
                v1 -= v1
            }
        case 390:
            var v51 uint32 = 42
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 15; fi52++ {
        if (-128 < -128) {
            v1 -= 42
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v53 uint64 = 0
        _ = fi52
        _ = v53
    }
    if true {
        var v54 uint32 = 668
        for fi55 := 0; fi55 < 6; fi55++ {
            switch ((v5 + v41) + v41) {
            case "":
                if true {
                    if (0.1 < 0.0) {
                        var v56 string = "\n\t"
                        var v57 int = fi55
                        _ = v56
                        _ = v57
                    }
                }
            default:
                v54 = v54
            }
            v41 = v5
            if len(v4) > 0 {
                v4[0] = v54
            }
            _ = fi55
        }
        _ = v54
    } else {
        var v58 uint = 255
        for fi59 := 0; fi59 < 20; fi59++ {
            for fi60 := 0; fi60 < 5; fi60++ {
                fmt.Println(fi60, v58, v58)
                fmt.Println(v2)
                switch int(v58) {
                case -2147483648:
                    _ = len(v4)
                case 0:
                    switch v5 {
                    case "a":
                        v4 = append(v4, 0)
                    case "":
                        fmt.Println(fi60, v3)
                    case "hello":
                        var v61 string = v5
                        _ = v61
                    }
                case 9223372036854775807:
                    v3 += ((float32(1e10) - v3) + ((((v3 / v3) - v3) + 3.14159) / 1.0))
                }
                _ = fi60
            }
            v58 -= (v58 & 4294967295)
            _ = fi59
        }
        if false {
            for fi62 := 0; fi62 < 11; fi62++ {
                for fi63 := 0; fi63 < 2; fi63++ {
                    v58 = (v58 >> ((18446744073709551615) & 63))
                    _ = fi63
                }
                fmt.Println(v41)
                _ = fi62
            }
        }
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
