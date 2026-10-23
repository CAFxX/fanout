package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8, p1 bool, p2 int) string {
    var v4 uint8 = (p0 >> ((829) & 63))
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (("" + "hello") + "foo bar")
}

func main() {
    var v1 int = 450
    var v2 uint8 = 216
    var v3 uint32 = 2147483648
    var v6 int16 = -1
    v7 := GenS1{F0: v2, F1: 42, F2: v2}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint32 = v3
        v10 := GenS2{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = v3
    if (!false) {
        v2 += 73
        var v11 uint8 = v2
        v12 := GenS1{F0: (v11 >> ((65535) & 63)), F1: 54, F2: (v11 - v11)}
        _ = v11
        _ = v12
    } else {
        v13 := []int{}
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: 0}
            var v16 uint32 = (v3 / ((v3) | 1))
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    var v17 int16 = (v6 * 1000)
    var v18 uint = 2
    if false {
        var v19 int16 = int16(v3)
        v20 := GenS2{F0: (v18 * 0)}
        var v21 float32 = 0.0
        _ = v19
        _ = v20
        _ = v21
    } else {
        var v22 uint32 = (v3 % ((uint32(v2)) | 1))
        switch 577 {
        case -690:
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := []int{(v1 ^ (v1 & fi23)), (v1 ^ (fi23 ^ fi23)), v1, 2147483647}
                _ = fi23
                _ = v24
            }
        case -2147483648:
            var v25 uint = v18
            _ = v25
        default:
            var v26 uint32 = v22
            _ = v26
        }
        _ = v22
    }
    v27 := []int{(v1 << (((v18 >> ((529) & 63))) & 63)), v1}
    if len(v27) > 0 {
        v27[0] = (v1 & 1)
    }
    if (42 == 0) {
        switch "\n\t" {
        case "foo bar":
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 uint = v18
                _ = fi28
                _ = v29
            }
        case "\n\t":
            v17 += (v6 << ((v18) & 63))
        default:
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
        }
    } else {
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 int64 = 557
            _ = fi30
            _ = v31
        }
        switch (("\n\t" + "hello") + "hello") {
        case "a":
            var v32 string = "foo bar"
            _ = v32
        case "foo bar":
            if (v1 != 667) {
                if false {
                    if false {
                        for fi33 := 0; fi33 < 8; fi33++ {
                            switch v1 {
                            case 927:
                                switch "foo bar" {
                                case "a":
                                    if (("a" + "a") >= "hello") {
                                        if false {
                                            var v34 string = ""
                                            _ = v34
                                        } else {
                                            for fi35 := 0; fi35 < 15; fi35++ {
                                                fmt.Println(v18, v2, fi33)
                                                var v36 float64 = 1.5
                                                _ = fi35
                                                _ = v36
                                            }
                                            v37 := GenS1{F0: v2, F1: 2, F2: v2}
                                            v18 -= 1
                                            _ = v37
                                        }
                                        v38 := GenS3{F0: v1, F1: (v18 << (((v18 << ((uint(v2)) & 63))) & 63)), F2: 58}
                                        _ = v38
                                    }
                                }
                            }
                            _ = fi33
                        }
                    }
                    switch (v1 + (^(v1))) {
                    case 32768:
                        var v39 uint8 = v2
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 12; fi40++ {
                        fmt.Println(fi40)
                        if (uint(v3) <= 384) {
                            var v41 uint64 = uint64(v18)
                            switch "foo bar" {
                            case "hello":
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            default:
                                v1 = (^(v1))
                            }
                            _ = v41
                        }
                        var v42 uint32 = v3
                        _ = fi40
                        _ = v42
                    }
                    v43 := GenS3{F0: (v1 / (((v1 - (v1 * v1))) | 1)), F1: (v18 * v18), F2: 42}
                    _ = v43
                }
                switch "" {
                case "\n\t":
                    fmt.Println(v1, v2, v3)
                case "foo bar":
                    var v44 int64 = 0
                    _ = v44
                case "a":
                    fmt.Println(v6, v2, v6)
                }
            } else {
                var v45 uint16 = 42
                _ = v45
            }
        case "hello":
            v3 -= (v3 * v3)
        }
    }
    v46 := GenS4{F0: 133, F1: ("\n\t" + "\n\t")}
    fmt.Println(v17, v1, v1)
    var v47 int16 = v6
    switch "" {
    case "hello":
        switch (v1 % ((1) | 1)) {
        case -496:
            for fi48 := 0; fi48 < 2; fi48++ {
                fmt.Println(v6)
                for fi49 := 0; fi49 < 5; fi49++ {
                    var v50 string = ""
                    _ = fi49
                    _ = v50
                }
                v51 := GenS4{F0: -1, F1: "foo bar"}
                _ = fi48
                _ = v51
            }
        default:
            var v52 float64 = 1e10
            _ = v52
        }
    case "\n\t":
        v27 = append(v27, int(v2))
    default:
        if (((!true) && (((false || false) || (42 == -2147483648)) || (("a" != "") && true))) || (0 == 503)) {
            switch (-(v1)) {
            case 4294967295:
                v53 := GenS5{F0: "", F1: 255}
                _ = v53
            }
            v54 := GenS4{F0: 256, F1: "hello"}
            fmt.Println(v18)
            _ = v54
        }
    }
    if len(v27) > 0 {
        v27[0] = (v1 * v1)
    }
    var v55 bool = ((v3 - v3) >= 626)
    v56 := GenS4{F0: 2147483647, F1: ("foo bar" + (("hello" + "hello") + ("\n\t" + "")))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v46
    _ = v47
    _ = v55
    _ = v56
    _ = v6
    _ = v7
}
