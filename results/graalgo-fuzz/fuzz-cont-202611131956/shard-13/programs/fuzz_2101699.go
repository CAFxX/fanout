package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = (("" + "foo bar") + ((("a" + ("" + "foo bar")) + ("foo bar" + ((("a" + "") + (("" + "hello") + "foo bar")) + "a"))) + ""))
    v3 := []int{32768}
    v4 := make(map[uint32]int)
    if (540 == 2147483647) {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 uint64 = 1
            for fi7 := 0; fi7 < 3; fi7++ {
                if ((v6 < (v6 | (v6 % (((v6 ^ (v6 / (((v6 | (^(v6)))) | 1)))) | 1)))) && (!false)) {
                    var v8 int16 = 1
                    v9 := GenS1{F0: 255, F1: 255, F2: 127}
                    if false {
                        var v10 int32 = 733
                        switch "a" {
                        case "a":
                            v11 := GenS2{F0: "", F1: 2, F2: 2}
                            _ = v11
                        case "":
                            switch (^(fi7)) {
                            case -589:
                                fmt.Println(fi5, fi7)
                            }
                        default:
                            _ = len(v4)
                        }
                        switch v2 {
                        case "":
                            var v12 uint16 = 42
                            switch (v2 + "\n\t") {
                            case "foo bar":
                                var v13 int16 = (^(v8))
                                _ = v13
                            case "a":
                                var v14 bool = (!(!false))
                                _ = v14
                            case "\n\t":
                                _ = len(v3)
                            }
                            var v15 int8 = 127
                            switch (-(fi5)) {
                            case 32767:
                                var v16 int32 = (v10 | 42)
                                _ = v16
                            case 42:
                                v17 := GenS1{F0: v12, F1: (^(v12)), F2: 91}
                                _ = v17
                            case -1:
                                v6 = v6
                            default:
                                var v18 float64 = 0.1
                                _ = v18
                            }
                            switch int(v8) {
                            case 32767:
                                if (((-(fi7)) == 797) && (fi5 <= 132)) {
                                    if false {
                                        var v19 uint = 4294967295
                                        _ = v19
                                    }
                                }
                            case -1:
                                switch v2 {
                                case "hello":
                                    if (!true) {
                                        if (true && false) {
                                            fmt.Println(v15)
                                        } else {
                                            fmt.Println(v10)
                                            v8 -= v8
                                            v20 := GenS2{F0: "", F1: 185, F2: 131}
                                            v12 += v12
                                            _ = v20
                                        }
                                        var v21 int16 = (v8 + (v8 + (v8 * v8)))
                                        _ = v21
                                    } else {
                                        if true {
                                            v22 := GenS2{F0: "a", F1: 42, F2: 42}
                                            v6 += (v6 / (((^(v6))) | 1))
                                            v10 = (v10 | int32(v6))
                                            var v23 uint16 = (^(v12))
                                            _ = v22
                                            _ = v23
                                        } else {
                                            var v24 int32 = (v10 << ((18446744073709551615) & 63))
                                            v8 += (v8 - v8)
                                            _ = v24
                                        }
                                    }
                                case "foo bar":
                                    var v25 int8 = -128
                                    _ = v25
                                }
                            default:
                                var v26 uint8 = 244
                                _ = v26
                            }
                            _ = v12
                            _ = v15
                        case "foo bar":
                            v27 := GenS1{F0: 42, F1: 677, F2: 42}
                            _ = v27
                        case "hello":
                            v8 += (v8 ^ (v8 << ((2) & 63)))
                        }
                        switch fi7 {
                        case 4294967295:
                            v6 += 4294967295
                        case 32767:
                            v28 := GenS3{F0: (^(v6)), F1: v8, F2: 127}
                            _ = v28
                        case -1:
                            v29 := GenS2{F0: v2, F1: 128, F2: 0}
                            _ = v29
                        default:
                            var v30 uint16 = 872
                            _ = v30
                        }
                        v8 = (-(v8))
                        _ = v10
                    } else {
                        var v31 bool = (!(1 != 2147483647))
                        _ = v31
                    }
                    _ = v8
                    _ = v9
                } else {
                    v32 := GenS4{F0: fi7}
                    if false {
                        switch (v2 + v2) {
                        case "\n\t":
                            fmt.Println(fi5)
                        default:
                            v2 = "hello"
                        }
                        switch "a" {
                        case "":
                            if false {
                                var v33 uint64 = (v6 >> ((0) & 63))
                                _ = v33
                            }
                            v2 = "\n\t"
                            v2 = v2
                        case "hello":
                            if true {
                                v34 := GenS1{F0: 0, F1: 570, F2: 11}
                                v2 = "foo bar"
                                _ = v34
                            } else {
                                var v35 int = fi5
                                _ = v35
                            }
                        }
                        if ((!true) || true) {
                            fmt.Println(v6, v6)
                            fmt.Println(fi5)
                        } else {
                            var v36 uint = 42
                            v37 := GenS4{F0: (fi7 - fi5)}
                            _ = v36
                            _ = v37
                        }
                    }
                    v2 = ((v2 + (v2 + "foo bar")) + "")
                    fmt.Println(fi5, v2)
                    _ = v32
                }
                v2 = ""
                _ = fi7
            }
            _ = fi5
            _ = v6
        }
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        v39 := GenS5{F0: 42}
        _ = fi38
        _ = v39
    }
    v40 := GenS2{F0: ("" + "a"), F1: 42, F2: 1}
    if true {
        for fi41 := 0; fi41 < 3; fi41++ {
            var v42 int64 = -641
            var v43 int16 = -952
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 int = 2147483647
                _ = fi44
                _ = v45
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        if true {
            v46 := GenS2{F0: "hello", F1: 231, F2: 2147483647}
            switch 9223372036854775807 {
            case 532:
                v2 = v2
            case -9223372036854775808:
                v2 = v2
            case -777:
                var v47 int8 = 1
                _ = v47
            }
            _ = v46
        }
        v2 = v2
    } else {
        v48 := GenS6{F0: 941, F1: 32767}
        v49 := v1[1]
        var v50 uint32 = (v49 | 875)
        _ = v48
        _ = v49
        _ = v50
    }
    v2 = "hello"
    v2 = (v2 + (((v2 + v2) + "foo bar") + v2))
    v2 = v2
    if (1 != 32767) {
        if ((0.0 / (1e10 - float32(1.5))) <= 1e10) {
            delete(v4, 2147483648)
            fmt.Println(v2)
            switch 128 {
            case -1:
                if false {
                    var v51 uint8 = 205
                    delete(v4, 4294967295)
                    v51 = v51
                    _ = v51
                } else {
                    v52 := GenS1{F0: 2, F1: 255, F2: 155}
                    v2 = v2
                    v53 := GenS7{F0: 127, F1: 0}
                    _ = v52
                    _ = v53
                }
            default:
                fmt.Println(v2)
            }
        } else {
            v2 = v2
            v2 = v2
            switch ("hello" + (("foo bar" + v2) + "\n\t")) {
            case "a":
                v54 := GenS8{F0: 9223372036854775807}
                _ = v54
            case "":
                v55 := v4[2147483647]
                _ = v55
            case "foo bar":
                v56 := v1[0]
                _ = v56
            }
        }
        if len(v3) > 0 {
            v3[0] = 1
        }
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        var v58 int = 426
        fmt.Println(v2, v2)
        v58 = v58
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        var v60 int = (fi59 / ((fi59) | 1))
        v2 = v2
        _ = fi59
        _ = v60
    }
    v2 = ("foo bar" + "\n\t")
    var v61 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v61
}
