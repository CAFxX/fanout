package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 int64; F1 uint; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 uint64) uint64 {
    p0 += (p0 * (p0 % (((p0 % (((^(p0))) | 1))) | 1)))
    p0 = (p0 - p0)
    _ = p0
    return (p0 >> ((761) & 63))
}
func genFn1(p0 int, p1 uint8) uint {
    var v5 uint = uint(p1)
    var v6 int16 = -525
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]string)
    v3 := []string{("foo bar" + "a"), "a", "\n\t"}
    var v4 uint = 794
    v4 += (v4 - (v4 ^ genFn1(651, 194)))
    var v7 uint64 = (v1 - genFn0(v1))
    switch -811 {
    case -872:
        if false {
            if false {
                if len(v3) > 0 {
                    v3[0] = ""
                }
            } else {
                v8 := GenS1{F0: "\n\t", F1: 1}
                _ = v8
            }
        }
    }
    v3 = append(v3, (("a" + "\n\t") + ("hello" + "hello")))
    var v9 uint64 = (v7 << (((^(v4))) & 63))
    v7 = v9
    if false {
        if (!(true && ((423 == 0) && ("a" < "\n\t")))) {
            v4 = v4
            var v10 uint8 = 127
            switch 32768 {
            case 32768:
                v9 += genFn0((v7 * v1))
            }
            _ = v10
        }
        v9 = genFn0((v1 << (((v4 - 255)) & 63)))
        var v11 uint8 = uint8(v4)
        _ = v11
    }
    if (!(true || true)) {
        v12 := v2[4294967295]
        if true {
            switch v12 {
            case "hello":
                for fi13 := 0; fi13 < 13; fi13++ {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        switch (fi13 / ((int(v7)) | 1)) {
                        case 360:
                            _ = len(v2)
                        case 620:
                            v15 := GenS1{F0: "foo bar", F1: 58}
                            _ = v15
                        default:
                            if true {
                                v16 := GenS1{F0: "hello", F1: int64(v9)}
                                _ = v16
                            }
                        }
                        _ = fi14
                    }
                    _ = fi13
                }
            default:
                switch ("foo bar" + v12) {
                case "foo bar":
                    if true {
                        v17 := GenS2{F0: -1, F1: 217, F2: -939}
                        var v18 int64 = 9223372036854775807
                        _ = v17
                        _ = v18
                    } else {
                        v4 = uint(v1)
                        v2[784] = "a"
                    }
                case "\n\t":
                    v7 += (v7 << ((v4) & 63))
                }
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                v1 -= 2
                var v20 string = "foo bar"
                _ = fi19
                _ = v20
            }
            switch "a" {
            case "a":
                v12 = ("\n\t" + v12)
            default:
                v1 = v1
            }
        } else {
            if (false && false) {
                switch "\n\t" {
                case "foo bar":
                    var v21 int16 = 32767
                    _ = v21
                case "":
                    if (false && (false || true)) {
                        v22 := GenS3{F0: -9223372036854775808, F1: (v4 | v4), F2: 32767}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 16; fi23++ {
                            fmt.Println(v12, fi23)
                            _ = fi23
                        }
                    }
                default:
                    switch v12 {
                    case "hello":
                        for fi24 := 0; fi24 < 8; fi24++ {
                            var v25 int16 = -603
                            _ = fi24
                            _ = v25
                        }
                    case "\n\t":
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    default:
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            var v27 int = 678
                            _ = v27
                        case 1:
                            v12 = v12
                        default:
                            if true {
                                v1 += 2
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch 127 {
                                case 9223372036854775807:
                                    delete(v2, 65535)
                                case -9223372036854775808:
                                    var v28 uint64 = (v9 * v9)
                                    _ = v28
                                case -1:
                                    var v29 string = v12
                                    _ = v29
                                }
                            }
                        }
                    }
                }
                fmt.Println(v12)
            } else {
                v4 = 255
            }
            var v30 uint32 = 2147483648
            _ = v30
        }
        _ = v12
    } else {
        var v31 bool = true
        _ = len(v3)
        _ = v31
    }
    v4 -= (v4 | (^(v4)))
    v32 := GenS4{F0: (v9 <= 18446744073709551615), F1: 2147483648, F2: 1}
    switch ("\n\t" + ("a" + "foo bar")) {
    case "foo bar":
        var v33 int64 = -1
        _ = v33
    case "":
        switch 267 {
        case -9223372036854775808:
            _ = len(v2)
        case 127:
            if true {
                var v34 uint8 = 236
                fmt.Println(v1)
                var v35 int8 = 1
                _ = v34
                _ = v35
            }
        case 373:
            var v36 float32 = 0.1
            _ = v36
        default:
            v1 += (v9 >> ((18446744073709551615) & 63))
        }
    }
    if false {
        var v37 bool = (!(0 < 998))
        _ = v37
    }
    _ = len(v2)
    var v38 uint8 = 13
    var v39 float32 = 0.1
    switch (("\n\t" + "") + "") {
    case "":
        _ = len(v3)
    case "foo bar":
        var v40 uint8 = v38
        _ = v40
    default:
        v41 := GenS1{F0: "a", F1: 42}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v7
    _ = v9
}
