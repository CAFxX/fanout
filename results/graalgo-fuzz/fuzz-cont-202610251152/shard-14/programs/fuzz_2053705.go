package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint; F2 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint64; F1 int; F2 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []uint32{1, 42}
    v4 := GenS1{F0: 856, F1: 650, F2: 0}
    switch "a" {
    case "foo bar":
        var v5 int = -612
        _ = v5
    }
    if true {
        var v6 int16 = -288
        for fi7 := 0; fi7 < 12; fi7++ {
            v6 = v6
            switch "\n\t" {
            case "a":
                for fi8 := 0; fi8 < 9; fi8++ {
                    v6 -= (v6 & (v6 >> ((65535) & 63)))
                    if false {
                        var v9 int = 32767
                        var v10 uint16 = 255
                        switch "" {
                        case "hello":
                            var v11 int32 = 674
                            _ = v11
                        case "foo bar":
                            var v12 int8 = 33
                            _ = v12
                        default:
                            v13 := GenS1{F0: (-(fi7)), F1: 18446744073709551615, F2: -29}
                            _ = v13
                        }
                        _ = v10
                        _ = v9
                    }
                    _ = fi8
                }
            default:
                var v14 uint = 65535
                _ = v14
            }
            _ = fi7
        }
        _ = v6
    }
    if (42 < 120) {
        switch "a" {
        case "\n\t":
            v15 := GenS1{F0: -748, F1: 962, F2: -64}
            _ = v15
        default:
            if false {
                switch "foo bar" {
                case "hello":
                    var v16 uint16 = 2
                    _ = v16
                case "\n\t":
                    switch "foo bar" {
                    case "foo bar":
                        if (false && false) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "a":
                        var v17 int8 = 126
                        _ = v17
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                v18 := GenS2{F0: 1}
                _ = v18
            } else {
                var v19 uint = 65535
                if true {
                    fmt.Println(v19)
                }
                _ = v19
            }
        }
    } else {
        var v20 int = 1
        _ = v20
    }
    var v21 float32 = 1.5
    v21 = (3.14159 * v21)
    fmt.Println(v21)
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 bool = (0 >= 781)
        if (v23 && (v23 && (2 < 2147483647))) {
            if v23 {
                if v23 {
                    v21 = v21
                    if true {
                        switch (fi22 - 128) {
                        case 4294967295:
                            for fi24 := 0; fi24 < 9; fi24++ {
                                if (!(710 < 2)) {
                                    if (42 == 1000) {
                                        v25 := GenS3{F0: ("a" + "a"), F1: 255, F2: 539}
                                        v21 = 0.0
                                        var v26 uint8 = 2
                                        _ = v25
                                        _ = v26
                                    } else {
                                        _ = len(v3)
                                    }
                                    if v23 {
                                        v27 := GenS3{F0: "hello", F1: 65535, F2: 42}
                                        fmt.Println(v23, fi24)
                                        var v28 float64 = 3.14159
                                        v23 = ((1.5 + float32(-2.25)) <= float32(v28))
                                        if (((!true) || ((-2.25 + -2.25) <= float32(0.0))) && false) {
                                            v28 += v28
                                            var v29 uint8 = 42
                                            var v30 float64 = -2.25
                                            _ = v29
                                            _ = v30
                                        } else {
                                            v23 = false
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if ("foo bar" <= ((((("\n\t" + "\n\t") + "a") + "hello") + "foo bar") + "foo bar")) {
                                            var v31 string = ("a" + "a")
                                            switch (fi22 | fi24) {
                                            case 4294967295:
                                                fmt.Println(v28, fi24)
                                            case 748:
                                                v23 = v23
                                            case 127:
                                                v23 = (4294967295 >= 2147483648)
                                            default:
                                                v32 := GenS2{F0: 32767}
                                                _ = v32
                                            }
                                            v33 := v1[-9223372036854775808]
                                            _ = v31
                                            _ = v33
                                        } else {
                                            fmt.Println(fi22, v23)
                                            _ = len(v2)
                                            var v34 int32 = 0
                                            if (v23 && v23) {
                                                v35 := GenS4{F0: -48, F1: "\n\t"}
                                                switch "hello" {
                                                case "hello":
                                                    v36 := GenS5{F0: 210, F1: 1, F2: 0}
                                                    _ = v36
                                                }
                                                _ = v35
                                            } else {
                                                switch 32767 {
                                                case 1:
                                                    var v37 int8 = 42
                                                    _ = v37
                                                case 690:
                                                    var v38 uint8 = 1
                                                    _ = v38
                                                default:
                                                    v39 := GenS4{F0: 97, F1: "\n\t"}
                                                    _ = v39
                                                }
                                                v23 = v23
                                            }
                                            var v40 uint64 = 0
                                            _ = v34
                                            _ = v40
                                        }
                                        _ = len(v3)
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v41 int64 = -796
                                        v42 := GenS1{F0: (fi22 * 42), F1: 18446744073709551615, F2: 127}
                                        v23 = ("" >= "foo bar")
                                        _ = v41
                                        _ = v42
                                    }
                                } else {
                                    _ = len(v1)
                                }
                                _ = fi24
                            }
                        default:
                            v23 = (127 <= 2)
                        }
                        for fi43 := 0; fi43 < 6; fi43++ {
                            v21 = v21
                            v44 := GenS3{F0: "", F1: 18446744073709551615, F2: 4294967295}
                            _ = fi43
                            _ = v44
                        }
                    }
                    for fi45 := 0; fi45 < 2; fi45++ {
                        var v46 string = ("foo bar" + "hello")
                        _ = fi45
                        _ = v46
                    }
                } else {
                    v23 = ((true || false) || (false && false))
                }
            } else {
                v2[fi22] = fi22
            }
            v47 := v2[(fi22 % ((fi22) | 1))]
            _ = v47
        }
        _ = fi22
        _ = v23
    }
    var v48 bool = false
    for fi49 := 0; fi49 < 5; fi49++ {
        for fi50 := 0; fi50 < 8; fi50++ {
            v51 := GenS1{F0: (fi50 | 42), F1: 962, F2: -1}
            if (!v48) {
                v48 = ((!(("\n\t" + "\n\t") >= "hello")) && v48)
                fmt.Println(v48, fi49)
                var v52 int8 = 126
                _ = v52
            } else {
                v3 = append(v3, 255)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v21 = ((v21 - v21) + -2.25)
            }
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    var v53 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v53
}
