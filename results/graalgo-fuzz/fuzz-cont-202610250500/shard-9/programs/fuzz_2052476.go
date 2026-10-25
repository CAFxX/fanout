package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint8; F1 string; F2 int8 }

func main() {
    var v1 int32 = 0
    v2 := make(map[int]int)
    v3 := []string{"hello", "", "hello"}
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if true {
        if true {
            var v4 string = "hello"
            var v5 int = 128
            _ = v4
            _ = v5
        }
        for fi6 := 0; fi6 < 6; fi6++ {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 uint32 = 65535
                _ = fi7
                _ = v8
            }
            fmt.Println(fi6, fi6)
            var v9 uint64 = 4294967295
            _ = fi6
            _ = v9
        }
        var v10 int64 = 0
        _ = v10
    }
    fmt.Println(v1)
    v11 := GenS1{F0: (("hello" + "foo bar") + ""), F1: 4294967295}
    var v12 uint64 = 656
    var v13 int8 = 99
    switch ("a" + "foo bar") {
    case "a":
        v14 := GenS1{F0: "hello", F1: (v12 * v12)}
        _ = v14
    case "hello":
        if (v13 >= v13) {
            v13 -= v13
            var v15 uint64 = 242
            _ = v15
        } else {
            if false {
                v13 += (v13 % ((v13) | 1))
                if ("\n\t" > ("foo bar" + "")) {
                    var v16 int8 = v13
                    var v17 uint = 42
                    var v18 uint32 = 2147483648
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    if (("\n\t" > "\n\t") && ((false && true) || (65535 >= 2147483647))) {
                        v19 := GenS2{F0: 255, F1: (v1 & v1)}
                        _ = v19
                    } else {
                        var v20 uint64 = v12
                        _ = v20
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v13 = -1
                    case "a":
                        v21 := GenS1{F0: ("\n\t" + "hello"), F1: v12}
                        _ = v21
                    }
                    if false {
                        if (!(true || (109 >= 2))) {
                            switch 4294967295 {
                            case 4294967295:
                                var v22 uint16 = 0
                                _ = v22
                            case 32767:
                                for fi23 := 0; fi23 < 7; fi23++ {
                                    for fi24 := 0; fi24 < 3; fi24++ {
                                        switch "hello" {
                                        case "foo bar":
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = ("a" + (((("\n\t" + "hello") + "hello") + "foo bar") + "hello"))
                                                }
                                            }
                                        default:
                                            fmt.Println(fi23, fi24, v1)
                                        }
                                        var v25 float32 = (0.0 * (0.0 * 1.5))
                                        if true {
                                            v13 += (v13 / ((v13) | 1))
                                        } else {
                                            if false {
                                                v1 = -1
                                            } else {
                                                switch (^(fi23)) {
                                                case 2147483647:
                                                    var v26 bool = false
                                                    fmt.Println(v12, v12)
                                                    var v27 float64 = 1.5
                                                    _ = v26
                                                    _ = v27
                                                default:
                                                    v28 := GenS1{F0: ("" + "\n\t"), F1: (v12 % (((v12 + 4294967295)) | 1))}
                                                    _ = v28
                                                }
                                                v13 -= (^(v13))
                                                v29 := GenS3{F0: -9223372036854775808, F1: 42}
                                                _ = v29
                                            }
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                    v1 -= v1
                                    var v30 int = (-(fi23))
                                    _ = fi23
                                    _ = v30
                                }
                            default:
                                var v31 string = ("" + "\n\t")
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 14; fi32++ {
                                var v33 int = -920
                                _ = fi32
                                _ = v33
                            }
                        }
                        v34 := GenS4{F0: 1, F1: -619}
                        _ = v34
                    } else {
                        fmt.Println(v12)
                        switch "foo bar" {
                        case "\n\t":
                            v13 = (v13 / ((-128) | 1))
                        }
                        fmt.Println(v13, v12)
                    }
                }
            }
            delete(v2, 42)
        }
    case "foo bar":
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS5{F0: "hello", F1: (fi35 & (^(fi35)))}
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 string = ""
                var v39 float32 = 0.0
                delete(v2, fi37)
                var v40 float32 = 3.14159
                switch "\n\t" {
                case "hello":
                    switch 340 {
                    case 32768:
                        switch (fi37 << ((536) & 63)) {
                        case 128:
                            if false {
                                v41 := GenS6{F0: 42, F1: -32768}
                                fmt.Println(v13, v1, v1)
                                _ = v41
                            } else {
                                var v42 uint8 = 2
                                v13 += v13
                                v39 = (v39 - -2.25)
                                _ = v42
                            }
                        case 32767:
                            if (((!(fi37 <= -9223372036854775808)) || (v1 >= (^(v1)))) && ((362 == 65535) && (94 > 42))) {
                                if (!(!true)) {
                                    fmt.Println(v13, fi37, v40)
                                    _ = len(v3)
                                    var v43 uint8 = 1
                                    _ = v43
                                }
                                v44 := v2[fi35]
                                _ = v44
                            } else {
                                switch v38 {
                                case "a":
                                    v2[32767] = fi37
                                default:
                                    var v45 int64 = 0
                                    _ = v45
                                }
                                if (-572 >= (v1 % ((42) | 1))) {
                                    if true {
                                        var v46 string = ((v38 + v38) + (v38 + "foo bar"))
                                        _ = v46
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        var v47 int8 = v13
                                        _ = v47
                                    }
                                }
                                var v48 uint8 = 255
                                var v49 uint32 = 1
                                v38 = v38
                                _ = v48
                                _ = v49
                            }
                        }
                    case 4294967295:
                        if false {
                            var v50 uint8 = 128
                            var v51 bool = (((^(v50)) > v50) && ((v1 <= (v1 - v1)) || ((42 != 41) && ((789 <= 0) && true))))
                            var v52 uint32 = uint32(v12)
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    default:
                        if (255 <= 0) {
                            var v53 bool = ((3.14159 * 0.1) == 1.5)
                            var v54 float64 = 1.5
                            _ = v53
                            _ = v54
                        }
                        var v55 float64 = float64(v40)
                        _ = v55
                    }
                default:
                    fmt.Println(v40)
                }
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            v56 := GenS7{F0: 127, F1: (("\n\t" + "\n\t") + "hello"), F2: (v13 / ((v13) | 1))}
            _ = fi35
            _ = v36
            _ = v56
        }
    default:
        var v57 int32 = (v1 % ((125) | 1))
        _ = v57
    }
    v58 := GenS2{F0: 4294967295, F1: v1}
    v12 += (v12 + (v12 - (v12 - 2)))
    v1 = 42
    var v59 int = 32768
    var v60 uint8 = 2
    v60 += v60
    fmt.Println(v60, v1, v13)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v58
    _ = v59
    _ = v60
}
