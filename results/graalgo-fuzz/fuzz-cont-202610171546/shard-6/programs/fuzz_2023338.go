package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint; F1 int; F2 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 bool; F2 int8 }
type GenS7 struct { F0 string; F1 string; F2 uint }
func genFn0(p0 bool, p1 uint8, p2 uint64) uint {
    var v6 int8 = -128
    var v7 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string) uint {
    var v8 int64 = 127
    if (42 != -1) {
        if (-1 < 1) {
            if true {
                if true {
                    if true {
                        var v9 int64 = v8
                        fmt.Println(v9, p0)
                        _ = v9
                    }
                    p0 = "\n\t"
                    v10 := []int{0, 127}
                    _ = v10
                } else {
                    v8 = (v8 | (-(v8)))
                    p0 = (p0 + (p0 + p0))
                }
            } else {
                fmt.Println(p0, p0)
                var v11 int32 = -55
                switch ((((("\n\t" + p0) + (p0 + "\n\t")) + (p0 + "")) + "foo bar") + p0) {
                case "\n\t":
                    v12 := []int{-9223372036854775808, -774}
                    _ = v12
                case "hello":
                    if true {
                        v13 := GenS1{F0: 28, F1: 255, F2: 0}
                        var v14 uint32 = 2147483647
                        if (1000 != -1) {
                            p0 = p0
                            v15 := GenS1{F0: -1, F1: 2, F2: -768}
                            _ = v15
                        } else {
                            v11 -= -2147483648
                            v16 := []int{}
                            v14 = (^(v14))
                            _ = v16
                        }
                        _ = v13
                        _ = v14
                    } else {
                        v17 := GenS2{F0: 261, F1: 0, F2: v11}
                        var v18 uint8 = 127
                        fmt.Println(v18)
                        _ = v17
                        _ = v18
                    }
                case "foo bar":
                    v19 := []int{67}
                    _ = v19
                default:
                    v8 += 9223372036854775807
                }
                _ = v11
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint16 = 65535
                _ = fi20
                _ = v21
            }
            if true {
                v22 := GenS2{F0: 4294967295, F1: 127, F2: 1}
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := []int{fi23, fi23, (fi23 >> ((244) & 63)), fi23}
                p0 = "hello"
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 19; fi25++ {
                var v26 uint8 = 127
                p0 = "foo bar"
                var v27 uint64 = 0
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS3{F0: 2, F1: 844}
            _ = fi28
            _ = v29
        }
        v30 := []int{-944, 32767, 2147483647}
        _ = v30
    } else {
        v31 := GenS2{F0: 1, F1: -277, F2: 1}
        switch 9223372036854775807 {
        case 9223372036854775807:
            v8 -= (v8 >> ((819) & 63))
        case 1:
            if false {
                var v32 int64 = v8
                var v33 int = -294
                _ = v32
                _ = v33
            } else {
                var v34 string = "a"
                var v35 uint = 65535
                _ = v34
                _ = v35
            }
        case 2147483647:
            p0 = p0
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            v37 := []int{128}
            var v38 int64 = 1
            var v39 int8 = -2
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = v31
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        for fi41 := 0; fi41 < 20; fi41++ {
            switch (fi41 << ((42) & 63)) {
            case 0:
                v42 := []int{fi40, fi41, fi41}
                _ = v42
            case 125:
                var v43 bool = ((!true) && (241 >= 4294967295))
                _ = v43
            }
            v44 := GenS2{F0: 216, F1: fi40, F2: 2147483647}
            _ = fi41
            _ = v44
        }
        v45 := GenS2{F0: 1, F1: (fi40 & 32768), F2: 2147483647}
        _ = fi40
        _ = v45
    }
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 string = p0
        _ = fi46
        _ = v47
    }
    if true {
        switch 32767 {
        case 1:
            v8 = 392
        case 32768:
            fmt.Println(p0)
        }
        p0 = (p0 + p0)
    }
    _ = p0
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int32 = -1
    var v2 int8 = 82
    var v3 string = "foo bar"
    v4 := make(map[int]int)
    var v5 uint8 = 96
    var v48 int64 = 0
    v1 += (v1 ^ v1)
    fmt.Println(v48, v1, v5)
    v3 = ""
    if false {
        if (true && false) {
            if (-319 > -2147483648) {
                v49 := []int{42, 0, -864, 42}
                _ = v49
            }
            v50 := GenS4{F0: ((v3 + ((v3 + v3) + v3)) + (v3 + v3))}
            for fi51 := 0; fi51 < 10; fi51++ {
                v52 := GenS5{F0: -9223372036854775808}
                v2 = -128
                v3 = ""
                _ = fi51
                _ = v52
            }
            _ = v50
        } else {
            v53 := []int{}
            v54 := GenS4{F0: (v3 + v3)}
            fmt.Println(v48, v5)
            _ = v53
            _ = v54
        }
        v5 = v5
        if (127 != 100) {
            var v55 int = -1
            _ = v55
        }
    } else {
        switch (((v3 + v3) + "") + ((v3 + v3) + v3)) {
        case "":
            for fi56 := 0; fi56 < 12; fi56++ {
                if true {
                    if (1.5 < float64(1.5)) {
                        fmt.Println(v5, fi56, v2)
                        if false {
                            v2 = (v2 | (v2 * (v2 - 100)))
                        }
                        if (v1 >= int32(v5)) {
                            v1 = 2147483647
                        } else {
                            v1 -= 1
                        }
                    }
                }
                if true {
                    v57 := GenS1{F0: v2, F1: (^(v5)), F2: 453}
                    var v58 uint32 = 65535
                    var v59 int64 = v48
                    _ = v57
                    _ = v58
                    _ = v59
                }
                _ = fi56
            }
        default:
            v60 := GenS6{F0: (v1 / (((v1 ^ v1)) | 1)), F1: (1.5 > 0.0), F2: (v2 / (((v2 >> ((genFn1((("\n\t" + "foo bar") + ""))) & 63))) | 1))}
            _ = v60
        }
        v48 -= -1
        switch ((("hello" + v3) + v3) + "hello") {
        case "foo bar":
            for fi61 := 0; fi61 < 12; fi61++ {
                switch (v3 + "foo bar") {
                case "":
                    if true {
                        fmt.Println(fi61)
                        switch ("hello" + "hello") {
                        case "\n\t":
                            if ((false || (!false)) || (!false)) {
                                var v62 uint32 = 0
                                _ = v62
                            }
                        case "hello":
                            if (!((true && false) && true)) {
                                if false {
                                    v63 := v4[-1]
                                    v64 := GenS4{F0: ((v3 + ("" + v3)) + v3)}
                                    _ = v63
                                    _ = v64
                                }
                            } else {
                                v65 := GenS1{F0: (^(v2)), F1: (v5 + (v5 ^ 1)), F2: -32768}
                                fmt.Println(fi61, v2)
                                var v66 int32 = 814
                                _ = v65
                                _ = v66
                            }
                        default:
                            if (!true) {
                                var v67 uint16 = 0
                                v68 := []int{}
                                if true {
                                    if (fi61 != (fi61 | (fi61 - fi61))) {
                                        var v69 string = ((v3 + "foo bar") + ("hello" + ""))
                                        _ = v69
                                    }
                                }
                                _ = v67
                                _ = v68
                            } else {
                                v70 := GenS7{F0: (v3 + v3), F1: v3, F2: 427}
                                var v71 uint32 = 65535
                                _ = v70
                                _ = v71
                            }
                        }
                    }
                }
                v5 = v5
                _ = fi61
            }
        case "a":
            v72 := GenS4{F0: v3}
            _ = v72
        case "hello":
            switch (v3 + "hello") {
            case "foo bar":
                delete(v4, 676)
            }
        }
    }
    for fi73 := 0; fi73 < 15; fi73++ {
        v48 = (v48 & -1)
        _ = fi73
    }
    v74 := GenS7{F0: ("hello" + ("\n\t" + "")), F1: "", F2: 454}
    v5 += v5
    if true {
        fmt.Println(v1, v5)
        fmt.Println(v48)
        v3 = "foo bar"
    }
    switch 1 {
    case 127:
        var v75 int32 = v1
        _ = v75
    case -2147483648:
        for fi76 := 0; fi76 < 6; fi76++ {
            _ = len(v4)
            var v77 int16 = 32767
            v1 += (v1 & (v1 ^ int32(v48)))
            _ = fi76
            _ = v77
        }
    case 539:
        fmt.Println(v48, v2)
    }
    for fi78 := 0; fi78 < 11; fi78++ {
        v79 := []int{fi78, int(v1)}
        v3 = ("\n\t" + "hello")
        _ = fi78
        _ = v79
    }
    var v80 string = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v74
    _ = v80
}
