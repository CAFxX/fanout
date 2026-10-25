package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 string; F1 int8; F2 string }
func genFn0() uint {
    v5 := []int{}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: "foo bar", F1: "foo bar", F2: 0}
        if false {
            var v8 int32 = 1
            var v9 uint = 607
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 13; fi10++ {
                if true {
                    var v11 uint8 = 1
                    if ((0.0 < float32(((-2.25 - (1e10 + ((-2.25 / 0.1) * -2.25))) * 3.14159))) && false) {
                        var v12 uint64 = 1
                        if false {
                            if (true && (4294967295 == 498)) {
                                v13 := GenS2{F0: (^(v11)), F1: -128}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v11 = (^(v11))
                                _ = v13
                            } else {
                                fmt.Println(v11, v11)
                                v12 = v12
                                v5 = append(v5, fi6)
                            }
                        }
                        _ = v12
                    } else {
                        _ = len(v5)
                        var v14 float64 = ((-2.25 - float64(0.1)) / ((float64(-2.25) * 1e10) - 1.5))
                        v11 += (v11 + v11)
                        _ = v14
                    }
                    _ = v11
                }
                var v15 bool = ((false || false) || ((192 != 2) || true))
                _ = fi10
                _ = v15
            }
            var v16 uint = 153
            var v17 string = "hello"
            _ = v16
            _ = v17
        }
        _ = fi6
        _ = v7
    }
    var v18 int = -9223372036854775808
    _ = v18
    _ = v5
    return 0
}

func main() {
    var v1 int32 = -346
    var v2 bool = false
    var v3 bool = (213 < 2)
    var v4 int16 = 1
    var v19 float64 = (((3.14159 + 0.0) - (float64((0.1 - 0.1)) - 3.14159)) + (float64((0.1 + (-2.25 + -2.25))) + 0.0))
    v20 := GenS2{F0: 1, F1: 127}
    if false {
        v21 := GenS2{F0: 70, F1: -37}
        _ = v21
    } else {
        if true {
            v22 := []int{-1, 893, 0}
            var v23 int8 = 0
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            switch "hello" {
            case "hello":
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS2{F0: 12, F1: 42}
                    _ = fi25
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                switch "" {
                case "\n\t":
                    var v28 int16 = (^(v4))
                    _ = v28
                default:
                    if ((float32(((v19 * v19) + (v19 / v19))) - -2.25) != 1.5) {
                        var v29 uint64 = 4294967295
                        if v3 {
                            if true {
                                var v30 uint32 = 970
                                switch "a" {
                                case "\n\t":
                                    v19 = v19
                                case "a":
                                    fmt.Println(v29, v19, v29)
                                    var v31 int16 = (v4 - v4)
                                    _ = v31
                                }
                                _ = v30
                            } else {
                                v32 := GenS1{F0: "a", F1: "foo bar", F2: (v29 & v29)}
                                var v33 int8 = 1
                                fmt.Println(fi27)
                                _ = v32
                                _ = v33
                            }
                            v19 -= 1.5
                            var v34 uint64 = 804
                            _ = v34
                        } else {
                            v4 = (v4 & v4)
                            var v35 uint64 = v29
                            if v2 {
                                switch fi24 {
                                case 32768:
                                    if ((int32(v29) <= 2147483647) && (!v3)) {
                                        v2 = true
                                        switch (fi27 << ((genFn0()) & 63)) {
                                        case 439:
                                            v36 := []int{-359, 128, -721}
                                            _ = v36
                                        case 1:
                                            var v37 uint64 = 190
                                            _ = v37
                                        default:
                                            v19 = v19
                                        }
                                    }
                                case 1:
                                    fmt.Println(v29)
                                    var v38 float32 = (0.1 + ((float32(v19) * 1e10) - 1e10))
                                    _ = v38
                                case 9223372036854775807:
                                    v39 := GenS3{F0: 127, F1: 204}
                                    _ = v39
                                }
                                v40 := GenS1{F0: (("\n\t" + "a") + "hello"), F1: ("a" + ("foo bar" + "hello")), F2: (v29 | v29)}
                                fmt.Println(fi27, fi24, v29)
                                _ = v40
                            }
                            _ = v35
                        }
                        v41 := GenS4{F0: 629, F1: uint(v29)}
                        var v42 uint32 = 2147483648
                        switch ("foo bar" + ("foo bar" + "hello")) {
                        case "":
                            fmt.Println(fi24, v19)
                        }
                        v4 -= v4
                        _ = v29
                        _ = v41
                        _ = v42
                    }
                }
                v43 := GenS3{F0: -1, F1: 9223372036854775807}
                _ = fi27
                _ = v43
            }
            _ = fi24
        }
        var v44 bool = (!(1 > 443))
        _ = v44
    }
    v4 += (-(v4))
    var v45 float32 = 0.0
    if v3 {
        v46 := GenS1{F0: "foo bar", F1: "\n\t", F2: 679}
        v47 := []int{}
        _ = v46
        _ = v47
    }
    switch -1 {
    case -2147483648:
        v45 = v45
    case 146:
        v48 := []int{}
        _ = v48
    case -957:
        switch -186 {
        case 128:
            switch (("foo bar" + "") + (("\n\t" + "") + ("" + (("" + ((((("a" + (("foo bar" + "foo bar") + "foo bar")) + "hello") + "\n\t") + "foo bar") + ("\n\t" + ("hello" + "")))) + "")))) {
            case "foo bar":
                for fi49 := 0; fi49 < 5; fi49++ {
                    if (!v3) {
                        var v50 string = ""
                        _ = v50
                    }
                    v19 -= 3.14159
                    if (-1 != 9223372036854775807) {
                        for fi51 := 0; fi51 < 4; fi51++ {
                            fmt.Println(fi51, v19, v3)
                            var v52 int16 = v4
                            v53 := []int{-2147483648, 42, 32768}
                            v54 := GenS5{F0: ("hello" + "a"), F1: 42, F2: "a"}
                            _ = fi51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    } else {
                        v4 = (v4 % ((v4) | 1))
                    }
                    _ = fi49
                }
            case "\n\t":
                if v2 {
                    if v2 {
                        var v55 uint64 = 551
                        v1 += v1
                        _ = v55
                    }
                    fmt.Println(v4, v45)
                }
            default:
                v1 = v1
            }
        }
    default:
        if v2 {
            for fi56 := 0; fi56 < 11; fi56++ {
                if (2 == 416) {
                    v1 += (-(v1))
                    v1 += v1
                }
                _ = fi56
            }
            v57 := []int{-1}
            _ = v57
        } else {
            v58 := GenS4{F0: 65535, F1: 42}
            _ = v58
        }
    }
    v4 -= (v4 << ((799) & 63))
    var v59 uint8 = 127
    for fi60 := 0; fi60 < 13; fi60++ {
        if ((-(v4)) >= v4) {
            var v61 int = (fi60 - 4294967295)
            v62 := GenS5{F0: ("foo bar" + (("hello" + "") + "a")), F1: 0, F2: "foo bar"}
            _ = v61
            _ = v62
        }
        var v63 int = fi60
        v59 = 2
        _ = fi60
        _ = v63
    }
    if v3 {
        v64 := GenS5{F0: "", F1: 42, F2: "a"}
        _ = v64
    } else {
        fmt.Println(v2, v2, v4)
        v3 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v59
}
