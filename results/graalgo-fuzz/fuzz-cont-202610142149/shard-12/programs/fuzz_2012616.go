package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 int16 }
type GenS8 struct { F0 uint64 }

func main() {
    var v1 string = ("a" + "foo bar")
    v2 := make(map[uint32]uint32)
    var v3 int = 2147483647
    v4 := []int{(v3 * -1), v3, (v3 / ((-1) | 1))}
    v5 := make(map[uint32]string)
    var v6 bool = true
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 65535}
        v6 = v6
        if (-649 == 1000) {
            for fi10 := 0; fi10 < 17; fi10++ {
                if (!v6) {
                    var v11 bool = (1000 >= -1)
                    v11 = (v6 || v6)
                    _ = v11
                } else {
                    var v12 int64 = 9223372036854775807
                    var v13 uint = 4294967295
                    var v14 float32 = (-2.25 * (0.1 * (((0.0 * (1e10 * 0.0)) / (1.5 - 0.0)) + 1e10)))
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
            var v15 uint64 = 18446744073709551615
            v15 = v15
            _ = v15
        } else {
            v16 := GenS1{F0: 255}
            if (!(!true)) {
                if (!v6) {
                    if len(v4) > 0 {
                        v4[0] = 677
                    }
                    v1 = v1
                    v3 = v3
                } else {
                    if (v6 && v6) {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            var v18 uint64 = 0
                            v1 = v1
                            v19 := GenS1{F0: 2}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    } else {
                        var v20 int32 = int32(v3)
                        _ = v20
                    }
                    fmt.Println(v3, v1, v3)
                    v3 = (v3 >> ((793) & 63))
                }
            } else {
                fmt.Println(fi8)
                switch (v1 + "a") {
                case "":
                    v2[1] = 42
                case "foo bar":
                    v21 := GenS2{F0: -1, F1: "a"}
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 14; fi22++ {
                        var v23 float32 = 0.0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            }
            if (255 == 170) {
                if false {
                    switch fi8 {
                    case -2147483648:
                        v24 := GenS3{F0: 445, F1: (("" + v1) + v1)}
                        _ = v24
                    case 1:
                        var v25 uint16 = 0
                        _ = v25
                    }
                } else {
                    switch (v1 + "foo bar") {
                    case "hello":
                        v26 := GenS4{F0: true, F1: 482, F2: 988}
                        _ = v26
                    case "\n\t":
                        var v27 bool = ((!(float64(0.1) < 0.0)) || (!(((("a" + v1) + ((v1 + v1) + v1)) + "a") >= v1)))
                        _ = v27
                    }
                    if v6 {
                        switch (v3 + v3) {
                        case 127:
                            v6 = (0.0 != 0.0)
                        case -568:
                            switch v3 {
                            case 1:
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    switch "a" {
                                    case "\n\t":
                                        switch fi28 {
                                        case 284:
                                            if ((true && v6) || (!(!v6))) {
                                                var v29 int = fi8
                                                var v30 int16 = -32768
                                                fmt.Println(fi8, fi28, v6)
                                                _ = v29
                                                _ = v30
                                            }
                                        default:
                                            v1 = v1
                                        }
                                    case "a":
                                        v6 = true
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = fi28
                                }
                            case -2147483648:
                                v31 := GenS5{F0: -871}
                                _ = v31
                            case -1:
                                if v6 {
                                    if ((("" + "\n\t") + (("" + "foo bar") + (v1 + v1))) < v1) {
                                        var v32 string = v1
                                        _ = v32
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                } else {
                                    if (2 < 2147483648) {
                                        v1 = v1
                                    }
                                    if true {
                                        if v6 {
                                            var v33 bool = ((!(v6 || v6)) || (!(!(-2.25 != 1e10))))
                                            if (!v33) {
                                                var v34 string = ""
                                                _ = v34
                                            }
                                            _ = v33
                                        }
                                        var v35 int8 = 127
                                        _ = v35
                                    }
                                    v36 := GenS6{F0: fi8, F1: -2147483648}
                                    _ = v36
                                }
                            default:
                                v37 := v2[636]
                                _ = v37
                            }
                        case 2147483647:
                            switch ((v1 + "a") + (v1 + v1)) {
                            case "":
                                var v38 uint32 = 65535
                                _ = v38
                            case "foo bar":
                                for fi39 := 0; fi39 < 6; fi39++ {
                                    switch fi8 {
                                    case 128:
                                        v40 := GenS6{F0: v3, F1: 971}
                                        _ = v40
                                    case -1:
                                        v3 += (fi39 & (v3 >> ((0) & 63)))
                                    }
                                    fmt.Println(fi8, fi39)
                                    var v41 bool = v6
                                    delete(v2, 128)
                                    var v42 uint64 = 4294967295
                                    var v43 bool = (0 == 55)
                                    _ = fi39
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            case "hello":
                                v44 := GenS4{F0: v6, F1: 788, F2: 255}
                                _ = v44
                            }
                        default:
                            fmt.Println(fi8)
                        }
                    }
                    if (-1 > -32768) {
                        v3 -= (v3 >> ((4294967295) & 63))
                    } else {
                        if len(v4) > 0 {
                            v4[0] = fi8
                        }
                        for fi45 := 0; fi45 < 16; fi45++ {
                            v46 := GenS3{F0: 918, F1: (((("a" + "foo bar") + ("" + "\n\t")) + v1) + ((v1 + ("foo bar" + v1)) + v1))}
                            v6 = v6
                            var v47 int8 = 42
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                }
            } else {
                _ = len(v4)
            }
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        if (32767 != 32767) {
            _ = len(v4)
            v1 = v1
            for fi49 := 0; fi49 < 6; fi49++ {
                var v50 int64 = 1
                fmt.Println(fi48)
                var v51 int16 = -376
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        v52 := GenS7{F0: 18446744073709551615, F1: -389, F2: -32768}
        v6 = v6
        _ = fi48
        _ = v52
    }
    if (-811 != -473) {
        if v6 {
            v3 = 442
        }
    } else {
        v53 := GenS5{F0: 71}
        if (true && v6) {
            fmt.Println(v3)
            fmt.Println(v1, v6, v6)
        }
        if (false && v6) {
            if (32767 > 1000) {
                for fi54 := 0; fi54 < 4; fi54++ {
                    if (1 > 192) {
                        v6 = v6
                        v55 := GenS3{F0: 735, F1: (v1 + ((v1 + (v1 + v1)) + v1))}
                        _ = v55
                    } else {
                        v1 = "a"
                        v6 = false
                        var v56 bool = v6
                        _ = v56
                    }
                    _ = fi54
                }
                var v57 uint = 4294967295
                _ = v57
            } else {
                switch (v3 | v3) {
                case 42:
                    fmt.Println(v3, v6)
                case 352:
                    v58 := GenS8{F0: 1}
                    _ = v58
                case 583:
                    v1 = "a"
                }
                fmt.Println(v6, v1)
            }
            var v59 uint = 483
            _ = v59
        }
        _ = v53
    }
    v6 = (42 != 0)
    v6 = true
    v3 -= 42
    v6 = v6
    v60 := GenS7{F0: 0, F1: -594, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v7
}
