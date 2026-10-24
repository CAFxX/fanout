package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 string; F2 uint64 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint8; F1 int8; F2 string }
type GenS10 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    v6 := GenS1{F0: 966, F1: 1}
    var v7 int8 = 70
    switch 128 {
    case 4294967295:
        var v8 uint16 = 819
        _ = v8
    default:
        if (1 > -2147483648) {
            var v9 uint32 = 1
            _ = v9
        } else {
            v10 := GenS2{F0: (v7 & v7)}
            fmt.Println(v7)
            _ = v10
        }
    }
    _ = v6
    _ = v7
    return 533
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    v3 := []string{("foo bar" + "hello")}
    var v4 uint16 = 286
    v5 := make(map[string]string)
    v11 := GenS3{F0: ("foo bar" + "a")}
    v4 += (v4 + v4)
    var v12 uint16 = v4
    if true {
        var v13 bool = true
        _ = v13
    }
    v14 := GenS4{F0: 0}
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS5{F0: -9223372036854775808}
        var v17 uint = 441
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint32 = 42
            fmt.Println(fi15, fi15, v19)
            fmt.Println(v17)
            _ = fi18
            _ = v19
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v12 = (v12 / ((0) | 1))
    var v20 int32 = -1
    var v21 bool = (false && (0 < 18446744073709551615))
    var v22 uint64 = 781
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint8 = 255
        _ = fi23
        _ = v24
    }
    if (int16(v20) >= 42) {
        switch -770 {
        case 42:
            v2 = append(v2, "hello")
        case 4294967295:
            if true {
                if (v21 || false) {
                    var v25 bool = (42 != v12)
                    fmt.Println(v4)
                    if ((v21 && false) || (float32((0.0 * 0.0)) != 3.14159)) {
                        fmt.Println(v4, v12, v25)
                    }
                    _ = v25
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    fmt.Println(v22, fi26)
                    var v27 int16 = -6
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v20, v12)
            }
        case -461:
            v20 = v20
        default:
            switch 127 {
            case 0:
                v28 := GenS3{F0: "a"}
                _ = v28
            case 32767:
                v29 := GenS6{F0: v21}
                _ = v29
            }
        }
        var v30 string = "\n\t"
        _ = v30
    } else {
        v31 := GenS7{F0: 4294967295, F1: "a", F2: v22}
        if ((1.5 * 3.14159) <= float32(-2.25)) {
            if v21 {
                var v32 uint32 = 2
                v22 = genFn0()
                var v33 uint32 = v32
                _ = v32
                _ = v33
            }
        } else {
            delete(v1, 413)
            for fi34 := 0; fi34 < 16; fi34++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "hello" {
                case "hello":
                    if ((fi34 * (fi34 + 4294967295)) == fi34) {
                        var v35 uint = 0
                        v20 += v20
                        var v36 uint32 = uint32(v4)
                        _ = v35
                        _ = v36
                    }
                case "a":
                    switch "hello" {
                    case "\n\t":
                        if (65535 == v12) {
                            var v37 bool = true
                            v20 += v20
                            _ = v37
                        }
                    case "foo bar":
                        v22 = 18446744073709551615
                    case "hello":
                        v22 = (^(v22))
                    default:
                        v38 := GenS7{F0: 156, F1: "", F2: genFn0()}
                        _ = v38
                    }
                default:
                    v21 = true
                }
                _ = fi34
            }
            for fi39 := 0; fi39 < 13; fi39++ {
                delete(v1, 222)
                var v40 float64 = 3.14159
                _ = fi39
                _ = v40
            }
        }
        _ = v31
    }
    switch "hello" {
    case "":
        if len(v3) > 0 {
            v3[0] = (((((("hello" + ("a" + "a")) + ("foo bar" + ("" + ("hello" + "foo bar")))) + "foo bar") + "") + "\n\t") + "\n\t")
        }
    case "hello":
        v41 := GenS8{F0: 306, F1: ("hello" + "")}
        _ = v41
    }
    fmt.Println(v12, v22)
    v42 := GenS9{F0: 139, F1: -128, F2: "foo bar"}
    switch ("" + "a") {
    case "\n\t":
        for fi43 := 0; fi43 < 10; fi43++ {
            if (0 == 806) {
                fmt.Println(v12, v20)
                var v44 int64 = 311
                _ = v44
            } else {
                switch (fi43 * 128) {
                case -540:
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                case -1:
                    v20 = -1
                case -349:
                    var v45 int64 = 42
                    _ = v45
                }
                v3 = append(v3, "")
                if v21 {
                    v4 -= v12
                }
            }
            if (v21 || (!((549 != v4) || (1 != -1)))) {
                v46 := GenS1{F0: 4294967295, F1: 926}
                _ = v46
            } else {
                var v47 uint64 = v22
                switch "" {
                case "foo bar":
                    var v48 uint16 = 65535
                    _ = v48
                }
                _ = v47
            }
            _ = fi43
        }
    case "":
        for fi49 := 0; fi49 < 9; fi49++ {
            fmt.Println(v4, fi49)
            var v50 int16 = -648
            _ = fi49
            _ = v50
        }
    case "a":
        switch 9223372036854775807 {
        case -9223372036854775808:
            switch -209 {
            case 32768:
                fmt.Println(v4, v12)
            case 90:
                for fi51 := 0; fi51 < 6; fi51++ {
                    for fi52 := 0; fi52 < 17; fi52++ {
                        var v53 float32 = (3.14159 - (-2.25 / ((float32((1e10 - -2.25)) * -2.25) * 0.1)))
                        switch fi52 {
                        case -587:
                            v54 := GenS10{F0: fi51, F1: v12}
                            _ = v54
                        case -9223372036854775808:
                            v20 -= v20
                        }
                        _ = fi52
                        _ = v53
                    }
                    _ = fi51
                }
            default:
                var v55 int16 = 438
                _ = v55
            }
        default:
            switch "foo bar" {
            case "a":
                v22 = 18446744073709551615
            case "\n\t":
                fmt.Println(v22, v12, v4)
            case "hello":
                var v56 int = -508
                _ = v56
            default:
                var v57 int = -215
                _ = v57
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
