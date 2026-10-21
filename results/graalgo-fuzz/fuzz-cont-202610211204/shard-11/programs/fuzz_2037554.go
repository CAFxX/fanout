package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
type GenS6 struct { F0 int64; F1 string; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 65535
    v3 := []uint32{2147483648}
    for fi4 := 0; fi4 < 13; fi4++ {
        switch (fi4 ^ (fi4 >> ((65535) & 63))) {
        case 0:
            if (!false) {
                if ("foo bar" < "a") {
                    v2 = v2
                    var v5 float32 = -2.25
                    _ = v5
                }
                var v6 uint = 976
                _ = v6
            }
        case 128:
            var v7 uint16 = v2
            _ = v7
        default:
            _ = len(v1)
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            switch (fi8 | -275) {
            case 1:
                var v9 uint16 = (v2 << ((886) & 63))
                _ = v9
            case 183:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch (fi8 * fi4) {
            case 1:
                var v10 int32 = -2147483648
                _ = v10
            case -9223372036854775808:
                switch ("hello" + (("" + "") + "a")) {
                case "hello":
                    fmt.Println(fi4, fi4, fi8)
                    if ((0.0 < 3.14159) && (false && (0.0 > (0.1 / 1.5)))) {
                        fmt.Println(fi8, fi8, v2)
                        var v11 uint64 = 0
                        _ = v11
                    } else {
                        v12 := GenS1{F0: 1000, F1: false, F2: 948}
                        _ = v12
                    }
                    var v13 uint16 = 535
                    _ = v13
                default:
                    v2 -= v2
                }
                if true {
                    switch "\n\t" {
                    case "foo bar":
                        var v14 int = (fi8 - (fi8 + 32768))
                        _ = v14
                    case "a":
                        v2 = v2
                    default:
                        var v15 uint16 = (v2 * v2)
                        _ = v15
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v16 := GenS1{F0: 737, F1: false, F2: (-(fi4))}
                    _ = v16
                }
                var v17 int64 = -1
                _ = v17
            case 392:
                if true {
                    if false {
                        v18 := GenS2{F0: 838}
                        var v19 uint64 = 1
                        switch int(v2) {
                        case 32768:
                            v19 = uint64(v2)
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                var v20 uint32 = 286
                                _ = v20
                            case "hello":
                                v2 -= 2
                            }
                        }
                        _ = v18
                        _ = v19
                    }
                    v2 = v2
                    v2 = v2
                }
                if true {
                    v1[803] = -93
                    delete(v1, 65535)
                } else {
                    var v21 string = "hello"
                    if (v21 < v21) {
                        if (221 == 65535) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v22 uint64 = 4294967295
                            v3 = append(v3, 269)
                            _ = v22
                        } else {
                            switch fi4 {
                            case -2147483648:
                                v21 = v21
                            case 32768:
                                v23 := GenS1{F0: 32767, F1: false, F2: (^(fi8))}
                                _ = v23
                            }
                            var v24 uint8 = 128
                            var v25 int16 = -32768
                            _ = v24
                            _ = v25
                        }
                    }
                    v2 -= (v2 ^ v2)
                    _ = v21
                }
                if false {
                    v2 -= v2
                } else {
                    v26 := GenS1{F0: 42, F1: (!(255 >= 2)), F2: fi4}
                    _ = len(v3)
                    _ = v26
                }
            }
            var v27 string = "\n\t"
            _ = fi8
            _ = v27
        }
        v28 := GenS3{F0: v2}
        _ = fi4
        _ = v28
    }
    fmt.Println(v2)
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            if true {
                v31 := GenS2{F0: 42}
                var v32 int = fi29
                fmt.Println(fi30, v2, v2)
                _ = v31
                _ = v32
            }
            fmt.Println(fi30)
            var v33 int32 = 1
            _ = fi30
            _ = v33
        }
        var v34 int8 = 112
        _ = fi29
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v2 -= (v2 + v2)
        v2 += v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi35
    }
    if (("" + "foo bar") == "") {
        v1[2] = -998
        var v36 string = ("hello" + "\n\t")
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 float64 = 3.14159
            var v39 float32 = 1.5
            var v40 int64 = 1
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        _ = v36
    } else {
        v41 := GenS4{F0: v2}
        fmt.Println(v2)
        switch "" {
        case "hello":
            switch "" {
            case "\n\t":
                switch 2147483647 {
                case 32767:
                    var v42 int32 = 1
                    _ = v42
                }
            case "hello":
                var v43 uint8 = 127
                _ = v43
            }
        case "\n\t":
            delete(v1, 640)
        default:
            v2 = 32767
        }
        _ = v41
    }
    var v44 int16 = 32767
    var v45 uint16 = 653
    var v46 float32 = ((1e10 - 1.5) * float32((1.5 / 1e10)))
    if len(v3) > 0 {
        v3[0] = 268
    }
    fmt.Println(v44, v44, v46)
    if (520 < 2147483648) {
        fmt.Println(v44, v44)
        v47 := GenS3{F0: (v2 ^ 32767)}
        _ = v47
    } else {
        v46 += (1.5 - (-2.25 * float32(3.14159)))
        fmt.Println(v2, v44, v44)
    }
    var v48 int = -568
    v44 += v44
    var v49 float64 = (0.0 / float64((0.1 + ((1.5 * (v46 - v46)) * 1e10))))
    for fi50 := 0; fi50 < 3; fi50++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v51 uint = 227
        _ = fi50
        _ = v51
    }
    v48 = v48
    switch int(v2) {
    case -132:
        v52 := GenS5{F0: -1, F1: v48, F2: 1}
        _ = v52
    case 854:
        _ = len(v1)
    case 0:
        for fi53 := 0; fi53 < 19; fi53++ {
            var v54 uint8 = 0
            v55 := GenS6{F0: 42, F1: "foo bar", F2: -32768}
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v56 uint8 = 42
        if (false || (18446744073709551615 > 942)) {
            var v57 float32 = 1.5
            var v58 float32 = (3.14159 * v57)
            _ = v57
            _ = v58
        } else {
            for fi59 := 0; fi59 < 1; fi59++ {
                v60 := GenS2{F0: 5}
                fmt.Println(v48)
                v61 := GenS6{F0: -68, F1: "", F2: v44}
                _ = fi59
                _ = v60
                _ = v61
            }
        }
        _ = v56
    } else {
        v62 := GenS3{F0: v2}
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v49
}
