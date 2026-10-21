package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    var v4 bool = (-2.25 != 0.1)
    v5 := []uint32{65535, 65535, 753, 626}
    var v6 int8 = -1
    v1 = ((("foo bar" + v1) + v1) + (v1 + ("a" + (v1 + v1))))
    fmt.Println(v1, v1, v6)
    var v7 uint64 = v2
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v1, v1)
        var v9 int = fi8
        switch (((v1 + (v1 + v1)) + "\n\t") + "hello") {
        case "\n\t":
            if v4 {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 int64 = 0
                    var v12 float32 = 3.14159
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                switch v1 {
                case "":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "a":
                    if (true || (!true)) {
                        if v4 {
                            v13 := GenS1{F0: 32767, F1: "foo bar"}
                            fmt.Println(v7, v4)
                            _ = v13
                        } else {
                            var v14 uint8 = 0
                            v2 -= 304
                            _ = v14
                        }
                        if v4 {
                            for fi15 := 0; fi15 < 2; fi15++ {
                                var v16 uint = 674
                                _ = fi15
                                _ = v16
                            }
                        }
                        fmt.Println(fi8, v2)
                    } else {
                        var v17 float32 = 1.5
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v1 = ("a" + v1)
                            _ = fi18
                        }
                        for fi19 := 0; fi19 < 9; fi19++ {
                            var v20 int64 = -9223372036854775808
                            v7 = v2
                            switch (fi8 ^ (v9 & fi19)) {
                            case -358:
                                switch v1 {
                                case "":
                                    v21 := GenS2{F0: v4, F1: 146}
                                    switch ("hello" + (((v1 + v1) + "foo bar") + "")) {
                                    case "\n\t":
                                        v9 = 128
                                    default:
                                        if v4 {
                                            v9 += 1
                                            switch 865 {
                                            case 2147483647:
                                                switch v1 {
                                                case "foo bar":
                                                    v6 = (v6 | v6)
                                                case "a":
                                                    fmt.Println(v7, v7, v9)
                                                case "hello":
                                                    v6 = 69
                                                }
                                            case -9223372036854775808:
                                                if ((42 == (v7 << ((65535) & 63))) || v4) {
                                                    var v22 string = ("" + v1)
                                                    _ = v22
                                                } else {
                                                    var v23 float64 = 1.5
                                                    var v24 uint64 = (v2 / ((v7) | 1))
                                                    _ = v23
                                                    _ = v24
                                                }
                                            default:
                                                switch 4294967295 {
                                                case -1:
                                                    v2 += 2
                                                default:
                                                    if (!(!false)) {
                                                        v4 = ((128 == 127) && (0.1 != 3.14159))
                                                    }
                                                }
                                            }
                                            v1 = (v1 + ("\n\t" + v1))
                                        }
                                        delete(v3, 4294967295)
                                    }
                                    _ = v21
                                }
                            case -1:
                                if false {
                                    v2 = 1
                                    var v25 uint8 = 42
                                    if (0 > 65535) {
                                        v7 = v7
                                    } else {
                                        v26 := GenS2{F0: (!false), F1: v25}
                                        if false {
                                            var v27 string = v1
                                            var v28 int32 = -2147483648
                                            v3[v9] = fi8
                                            _ = v27
                                            _ = v28
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v1 = (v1 + "hello")
                                            if true {
                                                v6 = v6
                                                fmt.Println(v1, v7, fi19)
                                            } else {
                                                if (v4 || true) {
                                                    var v29 uint16 = 389
                                                    var v30 float32 = 0.1
                                                    switch fi8 {
                                                    case 42:
                                                        var v31 int32 = -465
                                                        _ = v31
                                                    case 1:
                                                        fmt.Println(fi8, v29, v17)
                                                    }
                                                    delete(v3, fi19)
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    if true {
                                                        v7 = 1
                                                        v6 = (v6 & (v6 * 42))
                                                        switch "" {
                                                        case "":
                                                            v5 = append(v5, uint32(v25))
                                                        case "a":
                                                            v25 += v25
                                                        case "\n\t":
                                                            v4 = v4
                                                        default:
                                                            delete(v3, (fi8 / ((fi19) | 1)))
                                                        }
                                                    } else {
                                                        v7 = v2
                                                        var v32 float32 = (float32(1.5) - (-2.25 * v17))
                                                        _ = v32
                                                    }
                                                }
                                                var v33 uint = 171
                                                switch (fi19 * (fi8 % ((1) | 1))) {
                                                case 458:
                                                    v33 = v33
                                                case 2147483647:
                                                    v34 := GenS1{F0: 633, F1: (v1 + v1)}
                                                    _ = v34
                                                case -1:
                                                    fmt.Println(v17, v25, v17)
                                                }
                                                var v35 int8 = -1
                                                _ = v33
                                                _ = v35
                                            }
                                        }
                                        if v4 {
                                            if true {
                                                v2 += 1
                                                if (float32(3.14159) >= v17) {
                                                    v25 += v25
                                                    v36 := GenS2{F0: (((1.5 * 1.5) / 0.1) <= (-2.25 / 3.14159)), F1: v25}
                                                    _ = v36
                                                } else {
                                                    var v37 int32 = 42
                                                    var v38 uint32 = 0
                                                    _ = v37
                                                    _ = v38
                                                }
                                            } else {
                                                switch 42 {
                                                case 42:
                                                    var v39 string = v1
                                                    _ = v39
                                                }
                                            }
                                            v5 = append(v5, 2147483647)
                                        }
                                        _ = v26
                                    }
                                    var v40 bool = v4
                                    _ = v25
                                    _ = v40
                                } else {
                                    v1 = (v1 + (v1 + v1))
                                    var v41 float64 = (float64((float32(-2.25) / -2.25)) - 0.1)
                                    if ((v7 ^ (v2 - v7)) == 18446744073709551615) {
                                        fmt.Println(fi8)
                                        v42 := GenS3{F0: 1, F1: fi8}
                                        v2 = (v7 << ((42) & 63))
                                        v41 = -2.25
                                        _ = v42
                                    } else {
                                        var v43 uint32 = 65535
                                        var v44 float32 = 1e10
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v41
                                }
                                var v45 uint8 = 175
                                _ = v45
                            case 2147483647:
                                var v46 uint32 = 0
                                _ = v46
                            default:
                                v47 := GenS4{F0: 2}
                                _ = v47
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = v17
                    }
                case "hello":
                    fmt.Println(v1, fi8)
                default:
                    var v48 float64 = ((float64(-2.25) / (1e10 * float64(-2.25))) + 1.5)
                    _ = v48
                }
            }
        case "":
            var v49 uint = 1
            _ = v49
        case "foo bar":
            v50 := GenS5{F0: int(v2)}
            _ = v50
        }
        _ = fi8
        _ = v9
    }
    if (417 >= 450) {
        v1 = v1
        v4 = (32767 <= 0)
    } else {
        v6 = -90
        v6 -= v6
        v3[2147483647] = -9223372036854775808
    }
    var v51 string = ("foo bar" + v1)
    fmt.Println(v4)
    fmt.Println(v51, v6)
    var v52 bool = v4
    v53 := GenS1{F0: 1, F1: "\n\t"}
    v7 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v7
}
