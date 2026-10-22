package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint16 = 2
    var v2 bool = true
    var v3 float32 = 0.0
    v4 := []int{4294967295, 552, -2147483648, 397}
    for fi5 := 0; fi5 < 12; fi5++ {
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := GenS1{F0: 255, F1: 42, F2: -9223372036854775808}
                _ = fi6
                _ = v7
            }
        case "":
            var v8 int = fi5
            _ = v8
        case "hello":
            var v9 string = (("foo bar" + "") + "a")
            _ = v9
        default:
            switch ("hello" + "hello") {
            case "a":
                if v2 {
                    v10 := GenS1{F0: 166, F1: 18446744073709551615, F2: 318}
                    fmt.Println(v2, fi5)
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v12 := GenS2{F0: -1, F1: v1}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch fi11 {
                        case 2147483647:
                            v2 = ((v1 - v1) > (v1 << ((4294967295) & 63)))
                        case 32768:
                            var v13 float32 = v3
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    var v14 int64 = -691
                    _ = v14
                }
            case "hello":
                if false {
                    var v15 uint = 1
                    _ = v15
                } else {
                    v4 = append(v4, (fi5 * (fi5 - fi5)))
                    for fi16 := 0; fi16 < 3; fi16++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v17 bool = v2
                        fmt.Println(v17, v2)
                        var v18 uint32 = 2
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS1{F0: 140, F1: 1, F2: 0}
                    _ = v19
                }
            case "\n\t":
                if (!false) {
                    var v20 float64 = 1.5
                    var v21 string = "a"
                    _ = v20
                    _ = v21
                } else {
                    switch (("hello" + "") + "hello") {
                    case "\n\t":
                        var v22 uint16 = v1
                        _ = v22
                    case "hello":
                        switch "" {
                        case "hello":
                            if v2 {
                                var v23 uint64 = 507
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    v2 = (true && (fi24 < 32767))
                                    _ = fi24
                                }
                                _ = v23
                            } else {
                                v2 = ((v1 * (v1 | (^(v1)))) == 65535)
                            }
                        case "\n\t":
                            _ = len(v4)
                        case "a":
                            v3 = (v3 / ((((1.5 / (v3 - 3.14159)) * (v3 - 1.5)) * v3) * v3))
                        default:
                            if (!v2) {
                                if v2 {
                                    v1 -= v1
                                    var v25 int8 = -128
                                    v26 := GenS2{F0: -1, F1: (v1 / ((1) | 1))}
                                    _ = v25
                                    _ = v26
                                } else {
                                    v27 := GenS1{F0: 42, F1: 0, F2: 42}
                                    v1 = (v1 * v1)
                                    _ = len(v4)
                                    _ = v27
                                }
                            } else {
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    switch (fi28 % ((1) | 1)) {
                                    case -911:
                                        if v2 {
                                            v29 := GenS2{F0: -22, F1: (v1 ^ v1)}
                                            _ = len(v4)
                                            switch "" {
                                            case "foo bar":
                                                switch (fi5 % ((0) | 1)) {
                                                case 851:
                                                    var v30 uint16 = (v1 % ((v1) | 1))
                                                    _ = v30
                                                }
                                            case "a":
                                                if (v2 && v2) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v31 uint64 = 4294967295
                                                    _ = v31
                                                }
                                            default:
                                                if ((32768 == -239) || true) {
                                                    v32 := GenS3{F0: -128, F1: 0, F2: 0}
                                                    var v33 float32 = v3
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    v34 := GenS3{F0: 42, F1: -1, F2: 42}
                                                    var v35 uint16 = (^(v1))
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            var v36 int16 = 32767
                                            if (248 == uint8(v1)) {
                                                var v37 uint8 = 0
                                                fmt.Println(v36)
                                                _ = v37
                                            } else {
                                                fmt.Println(v36, fi28)
                                                v38 := GenS2{F0: -1, F1: v1}
                                                var v39 uint8 = 116
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v29
                                            _ = v36
                                        }
                                    default:
                                        v2 = ((("foo bar" + (("foo bar" + "\n\t") + "foo bar")) + "hello") == "\n\t")
                                    }
                                    fmt.Println(v2)
                                    _ = fi28
                                }
                                for fi40 := 0; fi40 < 3; fi40++ {
                                    var v41 int64 = 0
                                    var v42 int = fi40
                                    var v43 uint8 = 255
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            }
                        }
                    case "foo bar":
                        v44 := GenS1{F0: 127, F1: 18446744073709551615, F2: int64(v1)}
                        _ = v44
                    default:
                        var v45 uint16 = 255
                        _ = v45
                    }
                    var v46 uint = 925
                    v47 := GenS2{F0: 100, F1: (v1 % ((v1) | 1))}
                    _ = v46
                    _ = v47
                }
            }
        }
        _ = len(v4)
        _ = fi5
    }
    v2 = (!(true && ((128 < -1) && true)))
    v48 := GenS3{F0: 0, F1: 1, F2: -1}
    v4 = append(v4, 2147483647)
    var v49 bool = (255 < 128)
    if v49 {
        switch -145 {
        case -2147483648:
            var v50 int8 = 79
            _ = v50
        case -1:
            if (-2.25 >= 0.0) {
                var v51 float32 = v3
                v52 := GenS2{F0: 42, F1: v1}
                for fi53 := 0; fi53 < 17; fi53++ {
                    v2 = (!v49)
                    v3 += 3.14159
                    _ = fi53
                }
                _ = v51
                _ = v52
            }
        case 9223372036854775807:
            var v54 uint64 = 2
            _ = v54
        default:
            var v55 uint16 = v1
            _ = v55
        }
    } else {
        var v56 uint = 18446744073709551615
        for fi57 := 0; fi57 < 10; fi57++ {
            v58 := GenS4{F0: -1, F1: (true || v49)}
            var v59 uint16 = v1
            if false {
                v60 := GenS5{F0: uint8(v56), F1: 32767}
                var v61 uint16 = (v1 << (((v56 + (v56 ^ (^(v56))))) & 63))
                v4 = append(v4, -527)
                _ = v60
                _ = v61
            } else {
                v4 = append(v4, fi57)
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = len(v4)
        _ = v56
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
