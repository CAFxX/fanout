package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint16; F1 bool }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]string)
    v3 := []string{("a" + ""), "\n\t"}
    var v4 uint8 = 103
    v1 -= (-(v1))
    _ = len(v2)
    if (true && true) {
        if false {
            if (42 != 42) {
                var v5 uint = 4294967295
                var v6 uint16 = 65535
                _ = v5
                _ = v6
            } else {
                var v7 uint16 = 255
                v1 += -336
                var v8 uint = 4294967295
                _ = v7
                _ = v8
            }
        }
        fmt.Println(v1, v4)
        if (1e10 != 1.5) {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                for fi9 := 0; fi9 < 18; fi9++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2[141] = ("" + "hello")
                    _ = fi9
                }
            }
        }
    }
    var v10 int64 = -9223372036854775808
    if (4294967295 <= 85) {
        v1 -= (v1 ^ (v1 & (v1 % (((v1 / (((v1 - v1)) | 1))) | 1))))
        v1 = (v1 * 448)
        v3 = append(v3, "a")
    } else {
        if len(v3) > 0 {
            v3[0] = ""
        }
        fmt.Println(v1, v4)
    }
    var v11 string = "foo bar"
    fmt.Println(v10, v1, v4)
    var v12 uint8 = (v4 << ((0) & 63))
    v12 += (v4 | 223)
    var v13 bool = (23 > 178)
    var v14 int64 = v10
    switch v11 {
    case "hello":
        switch 32767 {
        case 362:
            if v13 {
                switch "hello" {
                case "hello":
                    v15 := GenS1{F0: 598}
                    _ = v15
                case "":
                    v16 := GenS2{F0: (v1 - 0), F1: 225, F2: 18446744073709551615}
                    _ = v16
                case "a":
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 int32 = -1
                        _ = fi17
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        delete(v2, 2)
                        _ = fi19
                    }
                }
                switch (v11 + v11) {
                case "":
                    v20 := GenS3{F0: "\n\t", F1: v14}
                    _ = v20
                case "\n\t":
                    fmt.Println(v11, v10, v4)
                case "foo bar":
                    v21 := v2[689]
                    _ = v21
                }
                var v22 float64 = (0.0 / float64(3.14159))
                _ = v22
            } else {
                var v23 bool = v13
                if (42 == 689) {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v11 = v11
                            var v26 int8 = int8(v1)
                            _ = fi25
                            _ = v26
                        }
                        v27 := GenS1{F0: 0}
                        for fi28 := 0; fi28 < 4; fi28++ {
                            v29 := GenS1{F0: uint32(v12)}
                            fmt.Println(v14, v14)
                            _ = fi28
                            _ = v29
                        }
                        _ = fi24
                        _ = v27
                    }
                } else {
                    v30 := GenS2{F0: (v1 >> ((1) & 63)), F1: v4, F2: 42}
                    var v31 bool = v23
                    _ = v30
                    _ = v31
                }
                fmt.Println(v11, v10)
                _ = v23
            }
        case 760:
            var v32 int64 = v14
            _ = v32
        }
    case "foo bar":
        if (2147483648 <= 4294967295) {
            v33 := GenS4{F0: 0, F1: 337, F2: (!((v13 || v13) || (42 < 724)))}
            v10 += v10
            _ = v33
        } else {
            v34 := GenS4{F0: -998, F1: 4294967295, F2: (true || ((v10 % ((35) | 1)) <= v14))}
            _ = v34
        }
    case "\n\t":
        switch -2147483648 {
        case 42:
            fmt.Println(v1, v13, v12)
        case -2147483648:
            var v35 int32 = 42
            _ = v35
        case 9223372036854775807:
            v36 := GenS3{F0: v11, F1: (v14 & (v10 & v10))}
            _ = v36
        default:
            for fi37 := 0; fi37 < 13; fi37++ {
                if (127 > 127) {
                    v38 := GenS5{F0: -33, F1: (v14 / ((v14) | 1))}
                    _ = v38
                } else {
                    switch ((v11 + (v11 + (((v11 + ("hello" + (("\n\t" + ("hello" + "a")) + v11))) + (v11 + "hello")) + v11))) + v11) {
                    case "hello":
                        v39 := GenS6{F0: 42, F1: (!(!(((1e10 - -2.25) >= 1.5) || v13)))}
                        _ = v39
                    case "a":
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v14 += (v14 & v14)
                            _ = fi40
                        }
                    default:
                        for fi41 := 0; fi41 < 7; fi41++ {
                            v3 = append(v3, "hello")
                            _ = len(v2)
                            fmt.Println(fi41)
                            _ = fi41
                        }
                    }
                }
                switch fi37 {
                case 385:
                    fmt.Println(v11, v12)
                case 345:
                    fmt.Println(v12, v10)
                case 1:
                    fmt.Println(v11)
                default:
                    var v42 bool = true
                    _ = v42
                }
                fmt.Println(v4, v14)
                _ = fi37
            }
        }
    default:
        for fi43 := 0; fi43 < 2; fi43++ {
            _ = len(v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi43
        }
    }
    var v44 int32 = 482
    _ = len(v2)
    v3 = append(v3, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
