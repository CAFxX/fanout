package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint32, p1 int64, p2 uint) uint {
    v5 := GenS1{F0: p0, F1: 0, F2: ("hello" + ((("" + ((("" + "\n\t") + "a") + ((("hello" + "\n\t") + "a") + "foo bar"))) + "\n\t") + ("hello" + "")))}
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS2{F0: ("a" + ((("\n\t" + "") + "a") + "\n\t"))}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        switch (((("" + "foo bar") + "a") + "hello") + "foo bar") {
        case "a":
            fmt.Println(p0, p1, fi8)
        case "hello":
            v9 := GenS1{F0: (p0 * (p0 + (^(p0)))), F1: 128, F2: "\n\t"}
            _ = v9
        case "foo bar":
            var v10 int64 = (^(p1))
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            p0 = p0
            var v12 float32 = -2.25
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    var v13 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return p2
}
func genFn1(p0 string, p1 bool, p2 int8) uint8 {
    p0 = (((("a" + (p0 + (p0 + p0))) + p0) + p0) + p0)
    v14 := GenS2{F0: (p0 + p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 206
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -851
    v3 := make(map[string]int)
    v4 := []uint32{879}
    if ((false || false) && true) {
        v2 += v2
        if (1000 <= -1) {
            v15 := GenS3{F0: 1, F1: 9223372036854775807, F2: "hello"}
            switch ("hello" + "\n\t") {
            case "\n\t":
                v16 := GenS2{F0: "hello"}
                _ = v16
            case "a":
                var v17 float32 = (3.14159 - (0.0 / 0.1))
                _ = v17
            case "hello":
                if false {
                    v2 -= v2
                }
            default:
                v2 = (v2 << ((0) & 63))
            }
            _ = v15
        } else {
            var v18 int16 = 32767
            if len(v4) > 0 {
                v4[0] = 1
            }
            v18 = v18
            _ = v18
        }
        var v19 string = ""
        _ = v19
    }
    var v20 uint8 = 255
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v21 := GenS3{F0: 617, F1: 128, F2: ("foo bar" + "hello")}
        fmt.Println(v20, v20)
        if (false || ((("foo bar" + "a") > "") || (!(true || ((-53 <= 1000) || (1 == 2)))))) {
            if (!true) {
                var v22 int16 = 1000
                var v23 int8 = -1
                _ = v22
                _ = v23
            }
            if (1e10 != 0.1) {
                v4 = append(v4, 2147483648)
            } else {
                v4 = append(v4, 65535)
            }
        } else {
            var v24 int = 127
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v24
        }
        _ = v21
    } else {
        for fi25 := 0; fi25 < 16; fi25++ {
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS4{F0: 100}
                switch (fi26 * fi25) {
                case 2147483647:
                    switch 128 {
                    case 9223372036854775807:
                        var v28 int8 = 0
                        _ = v28
                    case -258:
                        if false {
                            if (false || false) {
                                var v29 uint32 = 1
                                v2 += (v2 << ((2) & 63))
                                var v30 uint32 = (v29 | (v29 - 0))
                                _ = v29
                                _ = v30
                            } else {
                                v20 -= v20
                                v2 = (v2 | int32(fi26))
                            }
                            var v31 float32 = -2.25
                            var v32 bool = (75 < 42)
                            _ = v31
                            _ = v32
                        }
                    }
                case 0:
                    _ = len(v4)
                case -575:
                    var v33 uint64 = 18446744073709551615
                    _ = v33
                }
                switch "" {
                case "foo bar":
                    v2 -= v2
                case "hello":
                    if false {
                        v20 = v20
                        v20 += v20
                        if (!true) {
                            var v34 uint = 4294967295
                            v34 -= 704
                            v20 = v20
                            _ = v34
                        }
                    } else {
                        v20 += v20
                        var v35 uint8 = v20
                        _ = v35
                    }
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        var v36 uint16 = 728
        _ = v36
    }
    var v37 int16 = -32768
    v38 := GenS5{F0: 213, F1: (0.1 > 1e10)}
    if (!false) {
        if true {
            switch 127 {
            case 4294967295:
                switch "" {
                case "\n\t":
                    v37 -= (v37 + v37)
                default:
                    var v39 uint16 = 1
                    _ = v39
                }
            case 32767:
                fmt.Println(v2)
            case 500:
                for fi40 := 0; fi40 < 2; fi40++ {
                    if (4294967295 != 1) {
                        _ = len(v4)
                        if true {
                            v4 = append(v4, 1)
                        } else {
                            if false {
                                for fi41 := 0; fi41 < 13; fi41++ {
                                    v42 := GenS6{F0: 280, F1: 911}
                                    v37 = 966
                                    _ = fi41
                                    _ = v42
                                }
                                v37 = v37
                            } else {
                                if (1 != (^(v20))) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v43 int8 = 59
                                    for fi44 := 0; fi44 < 20; fi44++ {
                                        if len(v4) > 0 {
                                            v4[0] = 2147483648
                                        }
                                        var v45 uint = 4
                                        var v46 int = -921
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v43
                                } else {
                                    v37 += (v37 ^ v37)
                                }
                                v37 += (v37 & (v37 ^ 277))
                            }
                            if (v20 < 255) {
                                _ = len(v4)
                                var v47 uint16 = 65535
                                for fi48 := 0; fi48 < 19; fi48++ {
                                    v47 = 2
                                    _ = fi48
                                }
                                _ = v47
                            }
                        }
                        for fi49 := 0; fi49 < 5; fi49++ {
                            v37 -= 32767
                            fmt.Println(v20, v2)
                            _ = fi49
                        }
                    } else {
                        v50 := GenS6{F0: 255, F1: 1}
                        var v51 string = ((("hello" + "hello") + ("\n\t" + "")) + (("" + "a") + "hello"))
                        if (false || false) {
                            delete(v3, ((((v51 + (("hello" + "hello") + (v51 + v51))) + v51) + v51) + (v51 + v51)))
                            fmt.Println(fi40, v37, v37)
                            fmt.Println(v37, v2, v51)
                        }
                        _ = v50
                        _ = v51
                    }
                    if true {
                        for fi52 := 0; fi52 < 8; fi52++ {
                            if len(v4) > 0 {
                                v4[0] = 247
                            }
                            v37 -= v37
                            var v53 int32 = (v2 * int32(v37))
                            _ = fi52
                            _ = v53
                        }
                        v54 := v3[("foo bar" + "")]
                        _ = v54
                    }
                    _ = fi40
                }
            }
            v55 := GenS7{F0: -40}
            v2 -= v2
            _ = v55
        }
    }
    delete(v3, (("hello" + "hello") + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
