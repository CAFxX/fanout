package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int64 = 693
    v2 := make(map[uint32]bool)
    var v3 int64 = 730
    var v4 float32 = 1.5
    var v5 int32 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int32 = v5
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + ("\n\t" + "foo bar")) {
    case "\n\t":
        if false {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: (v1 & v3)}
                _ = fi8
                _ = v9
            }
            var v10 int16 = 150
            _ = v10
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v1, v1)
                var v12 int16 = 320
                _ = fi11
                _ = v12
            }
            var v13 bool = true
            _ = v13
        }
    case "a":
        fmt.Println(v4, v5)
    default:
        if true {
            for fi14 := 0; fi14 < 1; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    v5 = v5
                    v4 = 1e10
                    _ = fi15
                }
                var v16 bool = false
                v2[161] = v16
                _ = fi14
                _ = v16
            }
        } else {
            var v17 uint8 = 127
            fmt.Println(v4, v1, v5)
            delete(v2, 42)
            _ = v17
        }
    }
    var v18 float32 = -2.25
    v1 = 403
    switch -9223372036854775808 {
    case -2147483648:
        var v19 int = -500
        _ = v19
    case 32767:
        v20 := []int{-276}
        _ = v20
    case 0:
        v3 = v3
    default:
        v4 += v18
    }
    if (!false) {
        v21 := GenS1{F0: (-(v1))}
        v5 = v5
        _ = v21
    } else {
        v3 += 9223372036854775807
    }
    switch "" {
    case "":
        v22 := []int{-19}
        _ = v22
    case "hello":
        v23 := GenS1{F0: -166}
        _ = v23
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS1{F0: v3}
            fmt.Println(fi24, v18, v1)
            switch "foo bar" {
            case "a":
                switch (fi24 << ((601) & 63)) {
                case -1:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        v27 := v2[0]
                        var v28 uint16 = 0
                        var v29 uint8 = 0
                        switch fi24 {
                        case 42:
                            v3 -= (v3 | v1)
                        }
                        var v30 uint16 = 42
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                case 128:
                    v5 += v5
                case 516:
                    v4 -= (float32(0.0) / (0.0 / (1e10 + v4)))
                default:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        if (true || (false && true)) {
                            fmt.Println(v3, v4)
                            switch fi31 {
                            case 32768:
                                if true {
                                    var v32 float64 = -2.25
                                    if (0 >= -719) {
                                        _ = len(v2)
                                    } else {
                                        var v33 uint64 = 0
                                        _ = v33
                                    }
                                    var v34 bool = (true || true)
                                    _ = v32
                                    _ = v34
                                }
                            }
                        } else {
                            v3 = 1
                            if true {
                                v35 := v2[524]
                                _ = v35
                            } else {
                                if true {
                                    var v36 float64 = float64((v4 / v18))
                                    _ = v36
                                } else {
                                    var v37 string = "hello"
                                    switch "a" {
                                    case "foo bar":
                                        v38 := GenS2{F0: 386}
                                        _ = v38
                                    default:
                                        var v39 float32 = -2.25
                                        _ = v39
                                    }
                                    _ = v37
                                }
                            }
                        }
                        _ = fi31
                    }
                }
            }
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v1, v4)
    switch 9223372036854775807 {
    case -1:
        v3 -= 1
    case 128:
        if (!true) {
            var v40 bool = (18446744073709551615 >= 817)
            var v41 int8 = 100
            v5 -= (^(v5))
            _ = v40
            _ = v41
        } else {
            var v42 uint = 599
            switch 753 {
            case -1:
                var v43 uint64 = 0
                _ = v43
            case 1:
                switch "\n\t" {
                case "hello":
                    var v44 uint8 = 1
                    _ = v44
                case "foo bar":
                    if ((186 <= 1) || true) {
                        if false {
                            v45 := GenS1{F0: (^(v1))}
                            v3 -= (v1 * -9223372036854775808)
                            _ = v45
                        }
                        var v46 uint = 336
                        var v47 uint64 = 42
                        _ = v46
                        _ = v47
                    }
                case "\n\t":
                    v5 = v5
                }
            default:
                v48 := []int{}
                _ = v48
            }
            if false {
                var v49 uint64 = 507
                if true {
                    fmt.Println(v49, v42)
                    var v50 int = -2147483648
                    var v51 bool = ((true || ((v42 << ((650) & 63)) != v42)) && true)
                    _ = v50
                    _ = v51
                } else {
                    var v52 int8 = 100
                    var v53 uint64 = (v49 * v49)
                    delete(v2, 0)
                    _ = v52
                    _ = v53
                }
                delete(v2, 445)
                _ = v49
            } else {
                if (v4 != (-2.25 * (3.14159 / 0.1))) {
                    v54 := GenS1{F0: v1}
                    v55 := []int{725, 127, 984, 42}
                    for fi56 := 0; fi56 < 14; fi56++ {
                        var v57 bool = false
                        _ = fi56
                        _ = v57
                    }
                    _ = v54
                    _ = v55
                } else {
                    var v58 int16 = 32767
                    var v59 int = int(v1)
                    var v60 uint = 1
                    _ = v58
                    _ = v59
                    _ = v60
                }
                if true {
                    v61 := GenS3{F0: 1, F1: (v18 <= float32((3.14159 - -2.25)))}
                    _ = v61
                } else {
                    fmt.Println(v1, v4)
                    if (!true) {
                        delete(v2, 2147483647)
                    } else {
                        var v62 uint8 = 1
                        v42 = v42
                        for fi63 := 0; fi63 < 20; fi63++ {
                            var v64 int16 = -191
                            if true {
                                v2[224] = true
                                if true {
                                    delete(v2, 910)
                                }
                            }
                            var v65 int = (^(fi63))
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        _ = v62
                    }
                }
            }
            _ = v42
        }
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v4, v4)
    case "foo bar":
        v66 := []int{}
        _ = v66
    case "hello":
        if false {
            var v67 int = -229
            v68 := GenS4{F0: 1}
            _ = v67
            _ = v68
        } else {
            v3 = -315
        }
    default:
        if (211 <= 115) {
            if true {
                var v69 int64 = (v3 & (v3 << ((4294967295) & 63)))
                _ = v69
            } else {
                fmt.Println(v18, v5)
                v70 := []int{384}
                for fi71 := 0; fi71 < 15; fi71++ {
                    var v72 bool = false
                    _ = fi71
                    _ = v72
                }
                _ = v70
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
