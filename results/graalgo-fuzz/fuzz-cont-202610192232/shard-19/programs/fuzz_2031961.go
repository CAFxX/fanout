package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{2147483647, 42}
    v3 := make(map[uint32]uint32)
    v1 += 2
    for fi4 := 0; fi4 < 2; fi4++ {
        v1 = (v1 % (((v1 - 925)) | 1))
        var v5 uint8 = 2
        fmt.Println(fi4, v5, fi4)
        _ = fi4
        _ = v5
    }
    var v6 uint16 = (v1 | (v1 / (((v1 - (v1 & 1))) | 1)))
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 += (^(v6))
        _ = fi7
    }
    fmt.Println(v1)
    fmt.Println(v6, v6)
    v1 += v6
    fmt.Println(v6)
    v6 -= 1
    v3[393] = 943
    if (42 <= 157) {
        var v8 int = 1
        if false {
            if len(v2) > 0 {
                v2[0] = 2
            }
            v6 += (v1 >> ((65535) & 63))
        } else {
            fmt.Println(v6, v8, v1)
        }
        switch "\n\t" {
        case "":
            fmt.Println(v6)
        case "foo bar":
            v9 := GenS1{F0: (!((0 < 4294967295) && true)), F1: -278, F2: "hello"}
            _ = v9
        }
        _ = v8
    } else {
        v6 -= 2
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                if true {
                    fmt.Println(fi11)
                    if false {
                        v12 := GenS1{F0: (42 > uint32(v1)), F1: 224, F2: ("foo bar" + "hello")}
                        var v13 bool = true
                        _ = v12
                        _ = v13
                    }
                    if true {
                        switch ("\n\t" + ("foo bar" + "")) {
                        case "\n\t":
                            v1 = 1
                        case "":
                            if false {
                                v14 := GenS1{F0: true, F1: -1, F2: (("a" + ("hello" + "a")) + "")}
                                _ = v14
                            }
                        case "hello":
                            var v15 int32 = -1
                            v6 -= 2
                            var v16 int8 = 100
                            if false {
                                var v17 uint64 = 204
                                var v18 int8 = 127
                                var v19 float64 = 0.0
                                _ = v17
                                _ = v18
                                _ = v19
                            } else {
                                if (-9223372036854775808 > (fi11 % (((fi11 ^ int(v6))) | 1))) {
                                    v15 -= v15
                                    v20 := GenS1{F0: true, F1: 1, F2: "foo bar"}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v20
                                } else {
                                    v6 -= (v6 | (v6 >> ((0) & 63)))
                                }
                            }
                            var v21 int = 9223372036854775807
                            _ = v15
                            _ = v16
                            _ = v21
                        default:
                            if false {
                                switch (fi10 / (((fi11 + (fi11 | fi11))) | 1)) {
                                case -376:
                                    v22 := GenS1{F0: true, F1: 867, F2: "foo bar"}
                                    _ = v22
                                case 32768:
                                    switch "a" {
                                    case "\n\t":
                                        var v23 uint32 = 2147483647
                                        _ = v23
                                    case "a":
                                        if (120 >= 100) {
                                            v1 = 615
                                        } else {
                                            v6 -= 32767
                                            v24 := v3[65535]
                                            var v25 float64 = 1e10
                                            var v26 uint8 = 3
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                case -973:
                                    v1 -= v1
                                default:
                                    var v27 uint64 = 0
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        v28 := v3[4294967295]
                                        _ = v28
                                    default:
                                        var v29 float32 = 0.0
                                        _ = v29
                                    }
                                    v6 -= v1
                                    var v30 uint = 447
                                    v2 = append(v2, 2147483647)
                                    fmt.Println(v27)
                                    _ = v27
                                    _ = v30
                                }
                                v31 := GenS1{F0: true, F1: 486, F2: "\n\t"}
                                v1 = v1
                                _ = v31
                            } else {
                                if ((1 > 42) || ((false || true) || (969 == 2147483647))) {
                                    var v32 bool = ((((v6 % (((v1 % ((622) | 1))) | 1)) == 65535) || (!false)) || false)
                                    if (0.0 != 1e10) {
                                        v6 -= 255
                                        fmt.Println(v6)
                                        var v33 float32 = (0.1 - ((0.1 * (-2.25 - 3.14159)) + 3.14159))
                                        v34 := GenS2{F0: "a"}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := GenS3{F0: -845, F1: 365}
                                        _ = v35
                                    }
                                    _ = v32
                                } else {
                                    if true {
                                        switch ("hello" + "") {
                                        case "hello":
                                            fmt.Println(v1, fi10)
                                        default:
                                            var v36 int8 = -1
                                            _ = v36
                                        }
                                    }
                                }
                                v6 += v6
                                v6 = (^(v6))
                            }
                        }
                    }
                } else {
                    fmt.Println(v1, v6)
                }
                var v37 int32 = 2147483647
                v37 = 175
                _ = fi11
                _ = v37
            }
            v38 := GenS4{F0: 0, F1: 32767}
            _ = fi10
            _ = v38
        }
    }
    var v39 int64 = 203
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v6
}
