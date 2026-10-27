package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int64; F2 int16 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint }
type GenS8 struct { F0 int; F1 uint64; F2 string }
type GenS9 struct { F0 uint16; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 711
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := []uint32{58, 2147483647, 2147483648}
    var v6 bool = (-1 < -128)
    var v7 int16 = 42
    var v8 int64 = -1
    if (65535 >= 902) {
        var v9 float64 = 0.0
        _ = v9
    } else {
        v2 = v2
    }
    var v10 int = v2
    v11 := GenS1{F0: 127}
    v12 := GenS1{F0: 1}
    v13 := GenS2{F0: 156, F1: 378}
    if (42 > 27) {
        if ((3.14159 > 0.1) && (-1 > 0)) {
            v14 := GenS3{F0: 380}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v6 = v6
            _ = v14
        }
    } else {
        v5 = append(v5, 1)
        v7 = (v7 ^ 0)
    }
    switch ("" + "a") {
    case "":
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 string = ""
            _ = fi15
            _ = v16
        }
    case "hello":
        v2 -= (v2 * (v2 * (v2 ^ 128)))
    default:
        v7 = v7
    }
    if (1 >= 42) {
        for fi17 := 0; fi17 < 11; fi17++ {
            if v6 {
                v18 := GenS4{F0: 255}
                _ = v18
            } else {
                switch v10 {
                case -9223372036854775808:
                    var v19 uint8 = 0
                    _ = v19
                case 4294967295:
                    v1[288] = 992
                default:
                    var v20 uint16 = 255
                    _ = v20
                }
                v7 = 803
                for fi21 := 0; fi21 < 7; fi21++ {
                    switch ((("hello" + "a") + "\n\t") + ("hello" + "")) {
                    case "":
                        var v22 uint16 = 255
                        _ = v22
                    }
                    if (42 != 1) {
                        fmt.Println(v2, fi17)
                        if v6 {
                            v23 := GenS5{F0: v7}
                            v6 = v6
                            v7 = -683
                            _ = v23
                        }
                    }
                    var v24 uint = 18446744073709551615
                    delete(v4, 128)
                    _ = fi21
                    _ = v24
                }
            }
            if (v7 > v7) {
                v25 := v3[929]
                switch (v10 | (v2 / (((fi17 - (fi17 * v10))) | 1))) {
                case 127:
                    v8 -= 42
                case 845:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v27 := v1[0]
                        var v28 uint = uint(v27)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                fmt.Println(v25, v7)
                _ = v25
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                if v6 {
                    if ((v7 * 1000) >= -1) {
                        switch (((((("foo bar" + "hello") + ("foo bar" + "foo bar")) + "hello") + ((("" + "hello") + ("" + ("" + "hello"))) + ((("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "\n\t") + "hello"))) + "a") + "\n\t") {
                        case "hello":
                            v5 = append(v5, 65535)
                        default:
                            v4[fi29] = 2
                        }
                        v30 := GenS6{F0: 1, F1: 938, F2: (v7 % (((v7 | v7)) | 1))}
                        _ = v30
                    }
                    if false {
                        var v31 uint32 = 558
                        v8 -= 0
                        _ = v31
                    } else {
                        if (-2.25 <= (-2.25 - -2.25)) {
                            fmt.Println(fi29, v8)
                        }
                    }
                    var v32 int = 0
                    if false {
                        var v33 uint = 65535
                        _ = v33
                    } else {
                        var v34 uint = 1
                        _ = v34
                    }
                    var v35 int64 = v8
                    v36 := GenS2{F0: 839, F1: 814}
                    _ = v32
                    _ = v35
                    _ = v36
                }
                fmt.Println(fi17)
                _ = fi29
            }
            _ = fi17
        }
        var v37 bool = v6
        if false {
            var v38 uint8 = 2
            _ = v38
        } else {
            if v6 {
                switch v2 {
                case -671:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case -1:
                    switch ("\n\t" + "") {
                    case "":
                        v37 = (!v6)
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            for fi39 := 0; fi39 < 19; fi39++ {
                                v40 := v1[2147483647]
                                _ = fi39
                                _ = v40
                            }
                        case "":
                            for fi41 := 0; fi41 < 7; fi41++ {
                                var v42 float64 = 0.0
                                _ = fi41
                                _ = v42
                            }
                        default:
                            if (30 != 0) {
                                v10 = (v10 ^ v10)
                                if v37 {
                                    v6 = (1 > 0)
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    switch v2 {
                                    case -932:
                                        v5 = append(v5, 2147483648)
                                    default:
                                        fmt.Println(v8, v37)
                                    }
                                }
                            } else {
                                if v37 {
                                    v43 := GenS7{F0: 0, F1: 42, F2: 0}
                                    fmt.Println(v7, v37)
                                    _ = v43
                                } else {
                                    var v44 uint64 = 1
                                    _ = v44
                                }
                            }
                        }
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        v10 += v2
                    }
                }
                for fi45 := 0; fi45 < 4; fi45++ {
                    v46 := GenS5{F0: 42}
                    _ = fi45
                    _ = v46
                }
                if ((v6 || false) || v37) {
                    if (1 != 128) {
                        for fi47 := 0; fi47 < 8; fi47++ {
                            if (3.14159 < (0.0 / -2.25)) {
                                _ = len(v5)
                            }
                            _ = fi47
                        }
                        v48 := GenS2{F0: 1, F1: 760}
                        _ = v48
                    }
                    if len(v5) > 0 {
                        v5[0] = 247
                    }
                }
            } else {
                var v49 uint64 = 4294967295
                var v50 bool = (!(!v37))
                _ = v49
                _ = v50
            }
            var v51 float64 = (1e10 + 0.1)
            fmt.Println(v37, v10, v10)
            _ = v51
        }
        _ = v37
    } else {
        for fi52 := 0; fi52 < 12; fi52++ {
            if v6 {
                var v53 float64 = 1.5
                _ = v53
            }
            v2 = v2
            _ = fi52
        }
        if (!(65535 >= 2)) {
            for fi54 := 0; fi54 < 10; fi54++ {
                if v6 {
                    var v55 int16 = 42
                    _ = v55
                } else {
                    switch v10 {
                    case -1:
                        v56 := GenS8{F0: 108, F1: 2, F2: ""}
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 11; fi57++ {
                        var v58 int8 = -94
                        var v59 float32 = float32(0.0)
                        delete(v3, 1)
                        var v60 uint64 = 0
                        fmt.Println(v10)
                        v58 += -1
                        _ = fi57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                }
                var v61 uint64 = 2
                _ = fi54
                _ = v61
            }
            _ = len(v3)
            var v62 uint32 = 0
            _ = v62
        }
    }
    v7 = v7
    var v63 int32 = 2147483647
    v5 = append(v5, 65535)
    for fi64 := 0; fi64 < 19; fi64++ {
        if (18446744073709551615 != 18446744073709551615) {
            _ = len(v5)
        } else {
            fmt.Println(v8, fi64, v8)
        }
        _ = fi64
    }
    if v6 {
        var v65 int32 = (v63 % (((v63 << ((255) & 63))) | 1))
        _ = v65
    }
    v10 -= v2
    if ((v2 << ((65535) & 63)) == 4294967295) {
        v66 := GenS9{F0: 0, F1: false, F2: 65535}
        _ = v66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v7
    _ = v8
}
