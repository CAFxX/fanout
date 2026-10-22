package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }
type GenS5 struct { F0 bool; F1 string; F2 int }
type GenS6 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 uint = 1
    v3 := []uint32{363}
    var v4 uint32 = 936
    var v5 int32 = -1
    var v6 uint64 = 1
    if ((1e10 - -2.25) > 3.14159) {
        if (!true) {
            var v7 uint64 = v6
            var v8 int64 = 0
            if (!false) {
                fmt.Println(v2, v7)
                v9 := GenS1{F0: -1, F1: 32767, F2: v5}
                _ = v9
            } else {
                if false {
                    v10 := GenS2{F0: v4, F1: -1}
                    for fi11 := 0; fi11 < 14; fi11++ {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            _ = len(v1)
                            _ = fi12
                        }
                        for fi13 := 0; fi13 < 14; fi13++ {
                            switch (fi13 - (fi11 * (^(fi13)))) {
                            case 160:
                                if true {
                                    var v14 int32 = v5
                                    v6 = 2
                                    _ = v14
                                } else {
                                    fmt.Println(v4, v6)
                                    switch (fi11 & fi13) {
                                    case 0:
                                        fmt.Println(fi11, v8)
                                    }
                                    v1 = append(v1, "")
                                    v2 -= (v2 | 980)
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = (v4 * 19)
                                        }
                                        v2 = (^(v2))
                                    }
                                }
                            default:
                                v15 := GenS3{F0: (("\n\t" + "foo bar") + ("\n\t" + ("\n\t" + "hello"))), F1: 127}
                                _ = v15
                            }
                            _ = len(v1)
                            v16 := GenS2{F0: 42, F1: v5}
                            _ = fi13
                            _ = v16
                        }
                        _ = fi11
                    }
                    _ = v10
                }
                v4 -= v4
                var v17 uint64 = (v6 - (^(v6)))
                _ = v17
            }
            _ = v7
            _ = v8
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := GenS4{F0: (v2 << (((v2 ^ v2)) & 63)), F1: 2, F2: 756}
            v20 := GenS5{F0: ((1.5 - 0.0) != 1e10), F1: "a", F2: fi18}
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint8 = 2
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v23 uint32 = v4
    if true {
        fmt.Println(v5)
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 3; fi25++ {
                if len(v1) > 0 {
                    v1[0] = (("foo bar" + ((("\n\t" + (((("" + "foo bar") + "a") + ("foo bar" + "foo bar")) + ((("\n\t" + ("" + ("hello" + (("foo bar" + ("\n\t" + "\n\t")) + "")))) + "") + "foo bar"))) + "a") + "hello")) + "\n\t")
                }
                switch fi25 {
                case -538:
                    if false {
                        v4 = (v4 & v4)
                        fmt.Println(v4, fi24, v2)
                        fmt.Println(v23)
                        switch fi25 {
                        case 1:
                            v23 += (^(v23))
                        case -884:
                            switch (fi25 % (((fi24 * 689)) | 1)) {
                            case 188:
                                fmt.Println(v23, v5)
                            }
                        }
                    } else {
                        v23 += (v23 / (((v23 & v23)) | 1))
                        fmt.Println(v2, v5)
                        switch ("\n\t" + ("a" + ("" + ""))) {
                        case "a":
                            _ = len(v3)
                        }
                        v1 = append(v1, (("" + ("a" + "hello")) + ("" + ("\n\t" + "foo bar"))))
                    }
                case -9223372036854775808:
                    _ = len(v1)
                case 128:
                    var v26 uint = 0
                    _ = v26
                default:
                    var v27 int16 = -1
                    _ = v27
                }
                _ = fi25
            }
            switch ("a" + "") {
            case "foo bar":
                var v28 int16 = 0
                _ = v28
            case "":
                var v29 float32 = 0.0
                _ = v29
            case "a":
                if true {
                    switch fi24 {
                    case -404:
                        var v30 int8 = 1
                        _ = v30
                    }
                    v2 -= v2
                    v31 := GenS5{F0: true, F1: "foo bar", F2: (fi24 / ((int(v2)) | 1))}
                    _ = v31
                }
            default:
                if ((^(fi24)) > int(v2)) {
                    v3 = append(v3, 383)
                    var v32 int32 = (^(v5))
                    _ = v32
                }
            }
            v2 -= (v2 >> ((v2) & 63))
            _ = fi24
        }
    }
    var v33 int32 = (v5 * v5)
    v5 = -1
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 uint = (v2 & (v2 - (v2 / ((v2) | 1))))
        v36 := GenS2{F0: v23, F1: -2147483648}
        _ = fi34
        _ = v35
        _ = v36
    }
    if true {
        var v37 bool = ((255 < 127) || false)
        v38 := GenS2{F0: 42, F1: 42}
        _ = v37
        _ = v38
    } else {
        v6 -= (v6 ^ (v6 & (v6 - v6)))
        fmt.Println(v4, v4)
    }
    v2 += (v2 >> ((v2) & 63))
    v39 := GenS5{F0: (255 >= 127), F1: "foo bar", F2: -28}
    var v40 int64 = 9223372036854775807
    var v41 uint16 = 42
    fmt.Println(v6, v4)
    v23 += (^(v23))
    _ = len(v1)
    switch -152 {
    case 32768:
        fmt.Println(v2, v2, v41)
    }
    v42 := GenS4{F0: v2, F1: v6, F2: (v40 & v40)}
    switch 2147483647 {
    case 127:
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi43 := 0; fi43 < 2; fi43++ {
                if ("\n\t" >= "foo bar") {
                    v44 := GenS3{F0: ("" + "foo bar"), F1: 0}
                    var v45 uint = 2
                    _ = v44
                    _ = v45
                }
                if ((fi43 >> ((v2) & 63)) == 32768) {
                    var v46 float64 = (((0.0 * (((float64(1e10) / ((0.0 + (0.0 * 1.5)) - 0.1)) / 0.1) - 3.14159)) / 1.5) / (((1.5 + (3.14159 / 3.14159)) - 1e10) + 0.1))
                    var v47 uint32 = (v23 >> ((2) & 63))
                    var v48 uint16 = v41
                    _ = v46
                    _ = v47
                    _ = v48
                }
                v49 := GenS6{F0: 648, F1: 104}
                _ = fi43
                _ = v49
            }
        }
    default:
        fmt.Println(v41, v33, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
