package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 uint64) uint64 {
    p0 = (^(p0))
    for fi7 := 0; fi7 < 17; fi7++ {
        switch (^(fi7)) {
        case -996:
            p0 = p0
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(fi7)
                var v9 uint64 = (p0 << ((65535) & 63))
                _ = fi8
                _ = v9
            }
        }
        v10 := GenS1{F0: "foo bar"}
        _ = fi7
        _ = v10
    }
    _ = p0
    return p0
}
func genFn1(p0 int16) int32 {
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint16 = 2
            v13 -= v13
            v13 = (v13 - 255)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch "" {
    case "foo bar":
        p0 = (p0 / ((1000) | 1))
    }
    _ = p0
    return -487
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 uint32 = 4294967295
    v4 := make(map[string]string)
    v5 := []uint32{}
    v6 := make(map[string]string)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    delete(v4, "foo bar")
    switch 0 {
    case 1:
        var v14 int64 = 1
        _ = v14
    case 32767:
        if false {
            fmt.Println(v3)
        } else {
            fmt.Println(v3)
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: ("a" + ("a" + "a"))}
                _ = fi15
                _ = v16
            }
            if false {
                v3 += (^(v3))
                switch 4294967295 {
                case -676:
                    switch 42 {
                    case 326:
                        var v17 uint16 = 42
                        _ = v17
                    case -9223372036854775808:
                        if ("" < "a") {
                            var v18 uint = 42
                            _ = v18
                        }
                    case 0:
                        fmt.Println(v3)
                    }
                case -274:
                    var v19 float32 = float32((1.5 / float64(((float32((1e10 / (3.14159 - 1e10))) / (((1e10 + 3.14159) + (0.0 + (0.0 - -2.25))) * 1e10)) * (1e10 - (1.5 + (1.5 / -2.25)))))))
                    _ = v19
                case 423:
                    v20 := GenS2{F0: 255, F1: 201}
                    _ = v20
                default:
                    if (!(870 <= -2147483648)) {
                        fmt.Println(v3)
                        if (float32((0.0 - 0.1)) != ((1.5 - 3.14159) / 1.5)) {
                            fmt.Println(v3)
                        } else {
                            v3 = (v3 | 537)
                        }
                    }
                }
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v3 = (v3 + v3)
                    _ = fi21
                }
                v3 = (v3 * v3)
                if true {
                    if true {
                        if true {
                            var v22 bool = (false || false)
                            if (!v22) {
                                var v23 bool = v22
                                _ = v23
                            }
                            fmt.Println(v3)
                            _ = v22
                        }
                        v3 = (^(v3))
                    }
                    if (-653 <= 1000) {
                        if true {
                            var v24 bool = true
                            fmt.Println(v24, v3)
                            for fi25 := 0; fi25 < 12; fi25++ {
                                switch (fi25 % (((fi25 >> ((841) & 63))) | 1)) {
                                case 32767:
                                    v3 -= 4294967295
                                }
                                v24 = (!((("a" + "foo bar") + "") == ""))
                                var v26 uint = 255
                                _ = fi25
                                _ = v26
                            }
                            _ = v24
                        } else {
                            v27 := GenS2{F0: 32767, F1: 127}
                            var v28 int64 = 1
                            for fi29 := 0; fi29 < 10; fi29++ {
                                v30 := GenS2{F0: 0, F1: 0}
                                var v31 uint16 = 1
                                if ((v28 << ((712) & 63)) == v28) {
                                    v32 := GenS1{F0: "hello"}
                                    if (true && (18446744073709551615 != 2)) {
                                        if (-32768 < 32767) {
                                            v33 := v6["hello"]
                                            _ = v33
                                        } else {
                                            if (0 < v31) {
                                                var v34 int64 = (^(v28))
                                                _ = v34
                                            } else {
                                                _ = len(v5)
                                                v31 = (v31 | v31)
                                            }
                                            v31 = (^(v31))
                                            switch (("hello" + "a") + (("\n\t" + "foo bar") + "hello")) {
                                            case "foo bar":
                                                v35 := GenS3{F0: v28, F1: v3}
                                                _ = v35
                                            default:
                                                v31 -= v31
                                            }
                                        }
                                        v36 := GenS4{F0: 127}
                                        v37 := GenS5{F0: 42, F1: 2147483647}
                                        _ = v36
                                        _ = v37
                                    }
                                    v31 = 802
                                    _ = v32
                                }
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = v27
                            _ = v28
                        }
                        var v38 float32 = 0.0
                        _ = len(v1)
                        _ = v38
                    } else {
                        v3 += v3
                        v39 := GenS6{F0: int64(v3), F1: 42}
                        _ = v39
                    }
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                } else {
                    _ = len(v6)
                    var v40 uint32 = (v3 >> ((0) & 63))
                    _ = v40
                }
            }
        }
    default:
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 float32 = 0.0
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v3)
    switch ("" + ((("\n\t" + "hello") + "a") + "hello")) {
    case "a":
        v43 := GenS6{F0: 9223372036854775807, F1: -128}
        _ = v43
    default:
        v2["hello"] = v3
    }
    v44 := GenS1{F0: ""}
    var v45 int = 63
    for fi46 := 0; fi46 < 20; fi46++ {
        for fi47 := 0; fi47 < 10; fi47++ {
            v3 -= v3
            _ = fi47
        }
        switch "foo bar" {
        case "a":
            fmt.Println(fi46, v45)
        case "foo bar":
            v48 := GenS3{F0: 42, F1: 308}
            _ = v48
        default:
            v3 -= v3
        }
        _ = fi46
    }
    for fi49 := 0; fi49 < 6; fi49++ {
        var v50 int = (-(fi49))
        _ = fi49
        _ = v50
    }
    var v51 uint8 = 1
    switch ((((("a" + "") + (("" + "hello") + "\n\t")) + "") + (("\n\t" + "") + ("hello" + "hello"))) + "hello") {
    case "\n\t":
        fmt.Println(v51, v51, v3)
    case "a":
        if len(v5) > 0 {
            v5[0] = v3
        }
    case "hello":
        if false {
            for fi52 := 0; fi52 < 17; fi52++ {
                if (146 >= v3) {
                    v3 = (v3 + v3)
                    v53 := GenS5{F0: genFn0(249), F1: 42}
                    v51 = v51
                    _ = v53
                }
                v54 := GenS1{F0: ("a" + "foo bar")}
                for fi55 := 0; fi55 < 6; fi55++ {
                    fmt.Println(fi55, v51)
                    v51 = v51
                    if (0 < -9223372036854775808) {
                        var v56 float32 = 0.0
                        _ = v56
                    }
                    _ = fi55
                }
                _ = fi52
                _ = v54
            }
            var v57 float64 = 3.14159
            if false {
                for fi58 := 0; fi58 < 15; fi58++ {
                    var v59 uint32 = v3
                    v45 = -9223372036854775808
                    for fi60 := 0; fi60 < 3; fi60++ {
                        v6[("hello" + "a")] = "a"
                        _ = fi60
                    }
                    _ = fi58
                    _ = v59
                }
                v45 = v45
                fmt.Println(v45)
            } else {
                v61 := GenS7{F0: 42, F1: (v51 | v51), F2: 4294967295}
                _ = v61
            }
            _ = v57
        } else {
            v45 += (v45 % (((v45 + v45)) | 1))
            fmt.Println(v3)
        }
    default:
        fmt.Println(v45, v3, v45)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v51
    _ = v6
}
