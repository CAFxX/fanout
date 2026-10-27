package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 string; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 135
    v3 := []uint32{305, 2147483648, 4294967295}
    var v4 bool = (4294967295 > uint(v2))
    v5 := []uint32{42, 494, 2147483648, 0}
    var v6 int16 = -32768
    var v7 uint8 = v2
    v8 := v1[2147483648]
    fmt.Println(v6)
    var v9 int8 = -57
    if (!(((float32((0.0 - (3.14159 - float64((1.5 - (1.5 * (float32((1e10 / (1e10 * -2.25))) / ((1.5 * 0.0) + (float32((0.0 / 0.1)) - 1e10))))))))) - 1e10) * float32(1.5)) <= 0.0)) {
        v7 = (v7 + (v2 * 65))
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v9, v8, v8)
            var v11 int64 = -9223372036854775808
            v6 = (v6 | (v6 - (v6 | (v6 - 0))))
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            switch "a" {
            case "":
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint8 = uint8(v8)
                    _ = fi13
                    _ = v14
                }
            default:
                v1[v8] = v8
            }
            var v15 uint64 = 0
            _ = fi12
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            v7 = 111
            if ("foo bar" != "foo bar") {
                v17 := v1[v8]
                _ = v17
            } else {
                switch (fi16 ^ (fi16 | (fi16 & 9223372036854775807))) {
                case 32768:
                    var v18 int = -2147483648
                    _ = v18
                case 32767:
                    var v19 uint8 = (v2 - 1)
                    _ = v19
                }
                fmt.Println(fi16, v4)
            }
            v4 = false
            _ = fi16
        }
        v20 := GenS1{F0: 128, F1: 9223372036854775807, F2: 13}
        v6 = (v6 - (^(v6)))
        _ = v20
    }
    v7 -= 128
    var v21 string = ""
    var v22 string = ("" + v21)
    var v23 uint64 = 42
    if (!((661 <= 4294967295) || true)) {
        for fi24 := 0; fi24 < 13; fi24++ {
            if (v7 <= 228) {
                v7 += v7
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                if v4 {
                    switch v21 {
                    case "":
                        v1[42] = (^(v8))
                    default:
                        if false {
                            v21 = v21
                            switch v22 {
                            case "hello":
                                var v26 int8 = (v9 / ((127) | 1))
                                _ = v26
                            case "a":
                                if (v21 > (v21 + v22)) {
                                    v21 = "foo bar"
                                } else {
                                    v9 += v9
                                    var v27 int8 = 100
                                    if false {
                                        fmt.Println(fi25, v2, v8)
                                        v28 := GenS1{F0: (fi24 | (^(fi25))), F1: 0, F2: 0}
                                        _ = v28
                                    }
                                    _ = v27
                                }
                            }
                            v6 -= 1
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v1)
                } else {
                    var v29 uint8 = (v7 * 202)
                    var v30 uint8 = (v2 & v7)
                    v31 := GenS2{F0: 1, F1: (^(v8))}
                    _ = v29
                    _ = v30
                    _ = v31
                }
                v1[v8] = (v8 >> ((65535) & 63))
                _ = fi25
            }
            var v32 uint32 = 2147483647
            _ = fi24
            _ = v32
        }
        if v4 {
            v21 = ("" + v21)
            fmt.Println(v2, v21)
            _ = len(v1)
        } else {
            switch v22 {
            case "":
                var v33 string = ("foo bar" + v22)
                _ = v33
            case "\n\t":
                for fi34 := 0; fi34 < 11; fi34++ {
                    if v4 {
                        var v35 int = int(v6)
                        _ = v35
                    } else {
                        v36 := GenS3{F0: v21}
                        v6 = (^(v6))
                        _ = v36
                    }
                    var v37 uint64 = v23
                    _ = fi34
                    _ = v37
                }
            case "foo bar":
                var v38 uint = 1
                _ = v38
            default:
                if true {
                    v39 := GenS3{F0: v22}
                    if len(v3) > 0 {
                        v3[0] = v8
                    }
                    _ = v39
                } else {
                    v40 := GenS3{F0: ("foo bar" + "foo bar")}
                    var v41 int = 32768
                    v4 = (v6 > (v6 % ((v6) | 1)))
                    _ = v40
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 uint8 = v7
                var v44 uint = 0
                if ((^(v44)) <= 0) {
                    delete(v1, (v8 + v8))
                    v45 := GenS1{F0: -9223372036854775808, F1: int(v23), F2: 1}
                    v7 -= 2
                    _ = v45
                }
                _ = fi42
                _ = v43
                _ = v44
            }
            fmt.Println(v8)
        }
        for fi46 := 0; fi46 < 12; fi46++ {
            switch "\n\t" {
            case "a":
                v47 := GenS4{F0: "foo bar", F1: ("" + v22), F2: (v6 - (-(v6)))}
                _ = v47
            default:
                fmt.Println(v22)
            }
            _ = fi46
        }
    } else {
        delete(v1, 2147483647)
        for fi48 := 0; fi48 < 14; fi48++ {
            v49 := v1[(^(v8))]
            var v50 int8 = -128
            _ = fi48
            _ = v49
            _ = v50
        }
        fmt.Println(v23)
    }
    var v51 bool = (true && (899 <= 0))
    v2 = v7
    if (889 == 128) {
        var v52 uint8 = (v7 - (v2 + (v2 - (v2 & 48))))
        _ = v52
    } else {
        v1[v8] = v8
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 4294967295 {
        case 9223372036854775807:
            fmt.Println(v8, v21)
        case 128:
            var v53 int = 4294967295
            _ = v53
        }
    }
    var v54 uint64 = (v23 & v23)
    fmt.Println(v8, v22)
    if (((v6 + v6) >= -361) || true) {
        v6 = v6
        var v55 uint64 = v23
        fmt.Println(v4, v4, v23)
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
