package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 string; F2 uint64 }
type GenS7 struct { F0 int; F1 int; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 uint32 = 261
    var v6 uint = 2
    v7 := GenS1{F0: 862}
    v5 -= (v5 << ((v6) & 63))
    var v8 string = ("a" + ("a" + ""))
    v9 := GenS2{F0: 1000}
    fmt.Println(v6, v8)
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v8, v5)
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        if false {
            v6 += 14
        } else {
            if true {
                v8 = "foo bar"
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                if true {
                    switch (v8 + (v8 + "a")) {
                    case "":
                        v6 = v6
                    default:
                        v2[65535] = v8
                    }
                    if (false || true) {
                        v5 = v5
                        var v12 uint64 = 1
                        switch fi10 {
                        case 32768:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if (149 != 42) {
                                v13 := GenS3{F0: "foo bar", F1: 296}
                                v6 += 255
                                _ = v13
                            } else {
                                switch ((v8 + v8) + v8) {
                                case "hello":
                                    var v14 uint64 = v12
                                    _ = v14
                                default:
                                    var v15 uint = v6
                                    _ = v15
                                }
                                var v16 uint32 = v5
                                _ = v16
                            }
                            v8 = ((("" + v8) + v8) + v8)
                            if (!(v6 <= 42)) {
                                v8 = (v8 + v8)
                                var v17 int32 = 21
                                var v18 uint32 = v5
                                _ = v17
                                _ = v18
                            } else {
                                _ = len(v3)
                                v19 := GenS1{F0: 167}
                                var v20 uint64 = 996
                                _ = v19
                                _ = v20
                            }
                            if true {
                                var v21 bool = true
                                var v22 float32 = -2.25
                                fmt.Println(v8, v6)
                                v21 = v21
                                _ = v21
                                _ = v22
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                v12 = (v12 % (((v12 + (v12 << (((^(v6))) & 63)))) | 1))
                            }
                        }
                        _ = v12
                    }
                    var v23 uint16 = 42
                    _ = v23
                } else {
                    if (73 >= -19) {
                        v24 := GenS1{F0: 0}
                        _ = v24
                    } else {
                        switch fi11 {
                        case 42:
                            v6 -= v6
                        }
                    }
                    switch v8 {
                    case "":
                        switch v8 {
                        case "foo bar":
                            switch 0 {
                            case 654:
                                v25 := GenS4{F0: 32767, F1: false, F2: -9223372036854775808}
                                _ = v25
                            }
                        case "\n\t":
                            v5 -= v5
                        case "":
                            v8 = v8
                        default:
                            var v26 float32 = 3.14159
                            fmt.Println(fi11, v26)
                            var v27 int32 = 673
                            if (false && false) {
                                v27 = v27
                                delete(v1, v8)
                            }
                            _ = v26
                            _ = v27
                        }
                    case "\n\t":
                        v28 := GenS5{F0: v5, F1: 434, F2: 397}
                        _ = v28
                    }
                    fmt.Println(v5)
                }
                var v29 float64 = ((((((((1e10 - 0.0) / (-2.25 / 1.0)) - 3.14159) / 1.5) / -2.25) / (0.1 + 0.0)) * 0.0) - 0.1)
                fmt.Println(v5, fi11, v8)
                _ = fi11
                _ = v29
            }
        }
        _ = fi10
    }
    v30 := GenS6{F0: 0, F1: v8, F2: 4294967295}
    v31 := GenS3{F0: v8, F1: 1}
    if false {
        v32 := v3[v8]
        fmt.Println(v5)
        if (64 > 0) {
            fmt.Println(v8)
            v32 = (v5 | v32)
        } else {
            fmt.Println(v5, v5, v6)
            var v33 int64 = 0
            var v34 int32 = 0
            _ = v33
            _ = v34
        }
        _ = v32
    }
    var v35 int = 128
    fmt.Println(v35)
    for fi36 := 0; fi36 < 8; fi36++ {
        if len(v4) > 0 {
            v4[0] = (v8 + "hello")
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            fmt.Println(v5)
            _ = fi37
        }
        _ = fi36
    }
    if len(v4) > 0 {
        v4[0] = v8
    }
    v38 := GenS1{F0: 2}
    v35 += int(v6)
    if ((!false) && ("a" <= "\n\t")) {
        var v39 uint16 = 42
        for fi40 := 0; fi40 < 17; fi40++ {
            v41 := GenS6{F0: (v39 / (((v39 + v39)) | 1)), F1: ("foo bar" + ((v8 + ((v8 + "foo bar") + v8)) + "foo bar")), F2: 18446744073709551615}
            var v42 uint64 = 20
            if true {
                fmt.Println(v6, fi40, v8)
                var v43 int16 = 32767
                _ = v43
            } else {
                switch (fi40 + int(v39)) {
                case -767:
                    switch fi40 {
                    case 0:
                        if true {
                            v5 -= 374
                            if false {
                                if true {
                                    v44 := GenS7{F0: (fi40 - v35), F1: 42, F2: (v5 * 1)}
                                    switch (fi40 - v35) {
                                    case 4294967295:
                                        v39 = v39
                                    case 2147483647:
                                        v45 := v1[v8]
                                        _ = v45
                                    }
                                    var v46 float64 = (1e10 / 1.5)
                                    _ = v44
                                    _ = v46
                                } else {
                                    v47 := GenS6{F0: v39, F1: "hello", F2: uint64(v5)}
                                    _ = v47
                                }
                                v48 := GenS5{F0: v5, F1: 42, F2: (^(v39))}
                                _ = v48
                            }
                            var v49 string = v8
                            _ = v49
                        }
                    default:
                        v50 := GenS7{F0: (v35 - (v35 ^ 128)), F1: (fi40 >> (((^(v6))) & 63)), F2: v5}
                        _ = v50
                    }
                case 32767:
                    var v51 float64 = (3.14159 - ((1.5 / 1.0) * (-2.25 - (float64(((float32(3.14159) * -2.25) / -2.25)) + 1.5))))
                    _ = v51
                default:
                    v5 = v5
                }
                var v52 uint8 = 127
                _ = v52
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = v39
    }
    var v53 string = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
