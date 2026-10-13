package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint8 = 171
    var v2 uint64 = 1
    v3 := []int{}
    v4 := make(map[uint32]string)
    var v5 bool = false
    delete(v4, 42)
    v6 := GenS1{F0: (^(v1)), F1: -268}
    for fi7 := 0; fi7 < 17; fi7++ {
        delete(v4, 65535)
        fmt.Println(fi7, v5)
        _ = fi7
    }
    fmt.Println(v2)
    var v8 bool = v5
    fmt.Println(v2)
    var v9 int16 = 32767
    v5 = true
    if v8 {
        if v5 {
            if true {
                v2 = v2
                v10 := GenS1{F0: (v1 + (v1 | v1)), F1: 2147483647}
                var v11 uint32 = 274
                _ = v10
                _ = v11
            }
            _ = len(v3)
        } else {
            switch ("hello" + "hello") {
            case "\n\t":
                var v12 float64 = 1.5
                _ = v12
            default:
                v13 := GenS2{F0: 9223372036854775807, F1: ("\n\t" + ""), F2: (v2 - uint64(v1))}
                _ = v13
            }
            v1 = (^(v1))
            for fi14 := 0; fi14 < 1; fi14++ {
                if v5 {
                    switch ("a" + "foo bar") {
                    case "hello":
                        var v15 uint16 = 1
                        _ = v15
                    case "":
                        _ = len(v3)
                    case "\n\t":
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v17 := v4[uint32(v1)]
                            var v18 uint = 2
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    v5 = (1 != 65535)
                } else {
                    v4[716] = ((("foo bar" + "\n\t") + ("\n\t" + "")) + ("\n\t" + "\n\t"))
                    v19 := GenS3{F0: (v1 <= v1), F1: (fi14 + (fi14 - (fi14 / (((fi14 - (-(fi14)))) | 1))))}
                    if true {
                        if len(v3) > 0 {
                            v3[0] = fi14
                        }
                        var v20 int8 = 1
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v22 := GenS1{F0: v1, F1: 42}
                            if false {
                                var v23 float64 = 1.5
                                switch 42 {
                                case 2147483647:
                                    if ((float64(((((((-2.25 + (((float32(0.0) - -2.25) + 0.1) - 1e10)) * float32((v23 - float64(0.1)))) / 1e10) / 0.1) * 0.0) * 1.5)) != 0.0) || v5) {
                                        var v24 uint32 = 42
                                        v20 = (^(v20))
                                        _ = v24
                                    }
                                case 4294967295:
                                    var v25 string = "\n\t"
                                    _ = v25
                                }
                                _ = v23
                            } else {
                                fmt.Println(v20, v8, v1)
                                switch "foo bar" {
                                case "hello":
                                    var v26 uint32 = 0
                                    _ = v26
                                case "\n\t":
                                    v2 += (v2 - 2)
                                default:
                                    v27 := GenS3{F0: (!v5), F1: (fi14 % (((fi21 << ((1) & 63))) | 1))}
                                    _ = v27
                                }
                            }
                            fmt.Println(fi21)
                            var v28 float32 = 0.1
                            _ = fi21
                            _ = v22
                            _ = v28
                        }
                        _ = v20
                    }
                    _ = v19
                }
                var v29 uint = 65
                v5 = false
                _ = fi14
                _ = v29
            }
        }
    }
    switch 32767 {
    case 32768:
        v30 := GenS3{F0: v8, F1: -1}
        _ = v30
    case 128:
        v5 = v5
    case 1:
        var v31 bool = ((42 == 2147483647) && (0 >= 42))
        _ = v31
    }
    if v8 {
        var v32 string = "\n\t"
        for fi33 := 0; fi33 < 15; fi33++ {
            _ = len(v3)
            var v34 int32 = 498
            _ = fi33
            _ = v34
        }
        var v35 uint32 = 0
        _ = v32
        _ = v35
    } else {
        var v36 uint8 = (^(v1))
        v37 := GenS4{F0: -413}
        _ = v36
        _ = v37
    }
    if (255 == 65535) {
        v2 = (v2 - (^(v2)))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v8 = v8
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        if (!v8) {
            var v39 uint64 = v2
            var v40 int64 = -1
            if (v8 || (true || true)) {
                for fi41 := 0; fi41 < 7; fi41++ {
                    if ((v1 % ((v1) | 1)) != (v1 & (v1 + v1))) {
                        if len(v3) > 0 {
                            v3[0] = (fi41 - (fi41 >> ((42) & 63)))
                        }
                    }
                    var v42 string = "hello"
                    fmt.Println(v2, v40, v1)
                    _ = fi41
                    _ = v42
                }
                _ = len(v3)
                fmt.Println(v9, fi38)
            } else {
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    for fi43 := 0; fi43 < 18; fi43++ {
                        if (0 == 651) {
                            v44 := GenS2{F0: (v40 * (v40 * v40)), F1: (("\n\t" + "\n\t") + "hello"), F2: uint64(v1)}
                            v9 -= (v9 ^ 772)
                            _ = v44
                        }
                        _ = fi43
                    }
                    var v45 uint = 0
                    fmt.Println(v45)
                    _ = v45
                }
                var v46 uint32 = uint32(v39)
                _ = v46
            }
            _ = v39
            _ = v40
        } else {
            var v47 int16 = -820
            if false {
                var v48 int8 = 100
                _ = v48
            } else {
                v49 := GenS2{F0: -1, F1: ("\n\t" + ""), F2: (v2 | v2)}
                v50 := GenS2{F0: 614, F1: "\n\t", F2: (v2 | 4294967295)}
                v51 := GenS4{F0: -1}
                _ = v49
                _ = v50
                _ = v51
            }
            _ = v47
        }
        v3 = append(v3, 1)
        for fi52 := 0; fi52 < 13; fi52++ {
            var v53 bool = false
            _ = fi52
            _ = v53
        }
        _ = fi38
    }
    v54 := GenS1{F0: v1, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v8
    _ = v9
}
