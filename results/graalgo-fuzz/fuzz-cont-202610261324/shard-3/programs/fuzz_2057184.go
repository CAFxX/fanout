package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 uint64 }
type GenS4 struct { F0 int16; F1 int64; F2 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 int16, p1 uint64, p2 uint8) uint16 {
    p1 -= 744
    var v6 uint8 = 1
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 19; fi8++ {
        if ((v7 | v7) > (^(v7))) {
            p2 = (v6 >> ((677) & 63))
            p1 += p1
        }
        _ = fi8
    }
    var v9 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return v7
}

func main() {
    var v1 int32 = -1
    v2 := []string{"\n\t", "a", "foo bar"}
    var v3 uint8 = 119
    var v4 int = 32767
    var v5 uint = 18446744073709551615
    switch ("hello" + "") {
    case "":
        var v10 string = "\n\t"
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int = fi11
            var v13 uint64 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    case "\n\t":
        if (4294967295 >= 65535) {
            if (!false) {
                v5 += 255
                var v14 int8 = -1
                v14 += int8(v4)
                _ = v14
            }
        }
    }
    if ((float32(3.14159) * 3.14159) <= (0.1 / 1e10)) {
        v2 = append(v2, (("\n\t" + "") + ""))
        switch "a" {
        case "":
            v4 += v4
        case "foo bar":
            for fi15 := 0; fi15 < 2; fi15++ {
                if (65535 == 0) {
                    switch (("\n\t" + "a") + "\n\t") {
                    case "":
                        if false {
                            var v16 uint = (v5 / (((v5 ^ (v5 ^ 0))) | 1))
                            if ((!false) && true) {
                                v16 = v16
                                v17 := GenS1{F0: v3, F1: 122}
                                _ = v17
                            }
                            var v18 uint = uint(v3)
                            _ = v16
                            _ = v18
                        } else {
                            if (42 == 271) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    default:
                        v19 := GenS2{F0: 1000, F1: int(v3), F2: uint32(v5)}
                        _ = v19
                    }
                    v5 = v5
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if (true && (!(-1 == 1000))) {
                        var v20 uint64 = 2
                        _ = v20
                    }
                    var v21 float64 = -2.25
                    _ = v21
                }
                fmt.Println(fi15, fi15, v3)
                v22 := GenS3{F0: 868, F1: ("foo bar" + "a"), F2: 1}
                _ = fi15
                _ = v22
            }
        }
        v4 -= v4
    }
    v1 = (v1 | (v1 | v1))
    v23 := GenS3{F0: 42, F1: "a", F2: 18446744073709551615}
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int32 = -1
        _ = fi24
        _ = v25
    }
    v1 -= (v1 * v1)
    for fi26 := 0; fi26 < 18; fi26++ {
        if false {
            var v27 uint32 = 42
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 int = fi26
                v5 = 65535
                v30 := GenS4{F0: -32768, F1: 9223372036854775807, F2: (!(!true))}
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
        fmt.Println(v1)
        _ = fi26
    }
    v4 = v4
    v4 += 0
    v31 := GenS5{F0: false}
    if (!((!(!false)) && (false && (1 < 42)))) {
        _ = len(v2)
        switch 4294967295 {
        case 9223372036854775807:
            switch ("\n\t" + "\n\t") {
            case "":
                var v32 uint8 = (v3 - (v3 - 0))
                _ = v32
            default:
                for fi33 := 0; fi33 < 20; fi33++ {
                    if (!true) {
                        for fi34 := 0; fi34 < 19; fi34++ {
                            fmt.Println(v1)
                            v5 -= 832
                            fmt.Println(v1, fi34)
                            _ = fi34
                        }
                    }
                    fmt.Println(v1, v3, v3)
                    _ = fi33
                }
            }
        case 32768:
            v1 -= -974
        }
        fmt.Println(v3)
    } else {
        if (0 >= -1) {
            switch "\n\t" {
            case "\n\t":
                v35 := GenS4{F0: 1, F1: -9223372036854775808, F2: (-627 <= 42)}
                _ = v35
            case "":
                v36 := GenS5{F0: true}
                _ = v36
            case "hello":
                v5 += v5
            default:
                v37 := GenS2{F0: 0, F1: v4, F2: 2147483647}
                _ = v37
            }
            var v38 int = 232
            _ = v38
        } else {
            v3 = v3
            v3 -= (^(v3))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v39 := GenS5{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
}
