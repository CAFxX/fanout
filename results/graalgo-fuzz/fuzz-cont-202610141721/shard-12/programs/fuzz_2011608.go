package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 int8 }
func genFn0() bool {
    var v6 bool = (1 <= 2147483647)
    switch 1 {
    case 888:
        v7 := []int{}
        _ = v7
    case 1:
        v6 = true
    default:
        v8 := GenS1{F0: ("hello" + ("" + "foo bar")), F1: 445}
        _ = v8
    }
    switch "hello" {
    case "a":
        if ((255 >= 2) && (!(3.14159 >= (float32(-2.25) / 0.1)))) {
            switch "" {
            case "hello":
                switch "\n\t" {
                case "hello":
                    if (!(v6 || (true || v6))) {
                        var v9 float64 = 0.0
                        v9 += (float64(0.0) - 1e10)
                        _ = v9
                    }
                case "":
                    var v10 int64 = -1
                    _ = v10
                case "a":
                    v11 := []int{-427, 0}
                    _ = v11
                }
            case "\n\t":
                var v12 float32 = float32(3.14159)
                _ = v12
            default:
                v13 := []int{32768, 42, 32767, 747}
                _ = v13
            }
            v6 = (633 == 892)
            for fi14 := 0; fi14 < 5; fi14++ {
                switch -2147483648 {
                case 128:
                    v6 = (100 >= 127)
                default:
                    fmt.Println(v6, v6)
                }
                v15 := GenS1{F0: "hello", F1: (fi14 << ((42) & 63))}
                _ = fi14
                _ = v15
            }
        } else {
            if v6 {
                if (1e10 < 3.14159) {
                    v16 := []int{-534, 42}
                    _ = v16
                }
                var v17 float64 = (0.0 * 0.1)
                _ = v17
            }
            if (true && (!true)) {
                fmt.Println(v6)
            } else {
                v6 = (2 > 18446744073709551615)
                if (!v6) {
                    v6 = v6
                } else {
                    if (v6 && v6) {
                        v6 = v6
                        if v6 {
                            v18 := GenS1{F0: ("hello" + "\n\t"), F1: -1}
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v20 := GenS1{F0: (("a" + "\n\t") + ("hello" + ("\n\t" + (("\n\t" + "\n\t") + "a")))), F1: fi19}
                                _ = fi19
                                _ = v20
                            }
                            if v6 {
                                var v21 uint64 = 18446744073709551615
                                var v22 bool = ((106 >= 32767) || (!false))
                                _ = v21
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    for fi24 := 0; fi24 < 19; fi24++ {
                                        switch ("hello" + ("" + "a")) {
                                        case "":
                                            v6 = v6
                                        }
                                        _ = fi24
                                    }
                                    v6 = ((("foo bar" + "a") + (("\n\t" + ("" + "hello")) + "\n\t")) < ("\n\t" + "foo bar"))
                                    _ = fi23
                                }
                                fmt.Println(v6)
                            }
                            _ = v18
                        } else {
                            if (0 != 12) {
                                v6 = v6
                                v25 := GenS2{F0: 787, F1: 2147483647}
                                _ = v25
                            }
                            v26 := GenS2{F0: 402, F1: 4294967295}
                            if true {
                                v6 = (18446744073709551615 >= 4294967295)
                                var v27 uint64 = 738
                                if (!false) {
                                    v28 := GenS1{F0: "hello", F1: 128}
                                    v29 := []int{-9223372036854775808}
                                    var v30 float32 = 0.1
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                _ = v27
                            } else {
                                switch (((("\n\t" + ((("foo bar" + "a") + "hello") + ("" + "a"))) + "a") + (("foo bar" + ("" + "")) + "a")) + ("hello" + ("foo bar" + "hello"))) {
                                case "foo bar":
                                    for fi31 := 0; fi31 < 4; fi31++ {
                                        v32 := GenS2{F0: 255, F1: 356}
                                        fmt.Println(fi31, v6)
                                        _ = fi31
                                        _ = v32
                                    }
                                default:
                                    for fi33 := 0; fi33 < 15; fi33++ {
                                        fmt.Println(v6, fi33)
                                        v34 := GenS3{F0: (fi33 >> ((194) & 63)), F1: 592, F2: ("a" + "foo bar")}
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                            }
                            _ = v26
                        }
                        v6 = v6
                    } else {
                        for fi35 := 0; fi35 < 4; fi35++ {
                            var v36 int32 = 42
                            _ = fi35
                            _ = v36
                        }
                        var v37 uint16 = 1
                        for fi38 := 0; fi38 < 12; fi38++ {
                            v39 := []int{fi38}
                            var v40 int64 = -1
                            for fi41 := 0; fi41 < 1; fi41++ {
                                v37 -= v37
                                v42 := GenS1{F0: ("hello" + "hello"), F1: (-(fi38))}
                                _ = fi41
                                _ = v42
                            }
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                    }
                    v43 := []int{-2147483648, -853, 4294967295, -808}
                    var v44 int = 0
                    _ = v43
                    _ = v44
                }
            }
        }
    }
    if false {
        v6 = (true && (127 < -57))
    }
    fmt.Println(v6)
    _ = v6
    return v6
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"hello", "", "\n\t"}
    var v3 uint64 = 1
    v4 := []string{"hello"}
    var v5 uint16 = 65535
    v45 := GenS1{F0: "a", F1: 639}
    v46 := GenS1{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: 32768}
    var v47 int8 = 1
    var v48 uint8 = 2
    if ((v48 % ((v48) | 1)) < (^(v48))) {
        var v49 float32 = 0.0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v47, v1, v5)
        _ = v49
    }
    var v50 uint32 = (v1 << ((65535) & 63))
    for fi51 := 0; fi51 < 1; fi51++ {
        for fi52 := 0; fi52 < 9; fi52++ {
            v53 := GenS1{F0: "\n\t", F1: fi52}
            v47 = int8(v50)
            var v54 int16 = -1
            _ = fi52
            _ = v53
            _ = v54
        }
        v55 := GenS1{F0: "foo bar", F1: fi51}
        var v56 uint32 = v1
        _ = fi51
        _ = v55
        _ = v56
    }
    if true {
        if ((v47 + (v47 & v47)) > 127) {
            if (false && (-1 < 1)) {
                v50 += (v50 - v50)
            } else {
                v50 -= 65535
                v57 := GenS1{F0: ("foo bar" + "hello"), F1: 42}
                var v58 int8 = -1
                _ = v57
                _ = v58
            }
        }
        if ("a" == ("" + "a")) {
            v48 = (^(v48))
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v59 bool = true
    v60 := GenS4{F0: v47}
    v1 += (v1 * uint32(v5))
    fmt.Println(v48)
    _ = len(v2)
    if (true || (!false)) {
        var v61 int32 = -1
        _ = v61
    } else {
        if len(v2) > 0 {
            v2[0] = ("\n\t" + (("foo bar" + "a") + ("a" + "hello")))
        }
        switch 1 {
        case 32768:
            var v62 int32 = 977
            _ = v62
        case 4294967295:
            var v63 uint32 = (v1 % ((v1) | 1))
            _ = v63
        case -1:
            v48 = uint8(v5)
        default:
            for fi64 := 0; fi64 < 20; fi64++ {
                _ = len(v4)
                v5 = (v5 & (v5 << ((380) & 63)))
                _ = fi64
            }
        }
        v47 = -1
    }
    v50 += v50
    switch "foo bar" {
    case "hello":
        for fi65 := 0; fi65 < 10; fi65++ {
            v50 -= (v50 * v1)
            if true {
                if (756 != 255) {
                    v47 = v47
                    var v66 uint64 = (v3 << ((1) & 63))
                    _ = v66
                } else {
                    var v67 uint = 42
                    var v68 int64 = 245
                    var v69 uint8 = (v48 << ((uint(v1)) & 63))
                    _ = v67
                    _ = v68
                    _ = v69
                }
                var v70 uint16 = v5
                _ = v70
            }
            _ = fi65
        }
    case "a":
        _ = len(v4)
    case "\n\t":
        var v71 int16 = int16(v48)
        _ = v71
    default:
        v48 -= 1
    }
    switch 2147483647 {
    case 32767:
        v5 += (v5 >> ((615) & 63))
    case 128:
        _ = len(v4)
    case 2147483647:
        var v72 int32 = -1
        _ = v72
    default:
        for fi73 := 0; fi73 < 13; fi73++ {
            v1 += (v1 >> ((329) & 63))
            _ = fi73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v59
    _ = v60
}
