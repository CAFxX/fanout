package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []string{"a", (("foo bar" + ("hello" + "")) + "foo bar")}
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 uint64 = 4294967295
    v4 = (^(v4))
    v5 := GenS1{F0: (v4 >> ((18446744073709551615) & 63))}
    _ = len(v2)
    var v6 int16 = 1
    v7 := GenS2{F0: 365, F1: -2147483648, F2: 128}
    if false {
        var v8 float32 = 3.14159
        v4 = (v4 % ((v4) | 1))
        v6 = (v6 + (v6 - (v6 + 42)))
        _ = v8
    }
    v1 = append(v1, ("hello" + "\n\t"))
    if false {
        var v9 int8 = 100
        var v10 int64 = 42
        _ = v10
        _ = v9
    }
    fmt.Println(v4)
    switch -534 {
    case -1:
        v4 += v4
    default:
        v4 = v4
    }
    var v11 uint64 = (v4 % (((^(v4))) | 1))
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            switch (("" + ("" + "foo bar")) + "a") {
            case "hello":
                if true {
                    var v14 int16 = (v6 & (v6 & (v6 ^ v6)))
                    if (false || (42 > 255)) {
                        v11 = v11
                    } else {
                        fmt.Println(fi13)
                        switch "a" {
                        case "\n\t":
                            if (42 != v6) {
                                v4 -= v11
                                var v15 float32 = 1e10
                                _ = v15
                            } else {
                                v11 = 4294967295
                            }
                            var v16 int = (fi13 - (fi13 >> ((877) & 63)))
                            if ((!true) && false) {
                                _ = len(v2)
                            } else {
                                fmt.Println(v14)
                            }
                            _ = v16
                        case "a":
                            var v17 uint64 = v4
                            _ = v17
                        case "":
                            fmt.Println(v14, v6, v14)
                        }
                        switch "foo bar" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "\n\t":
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            var v19 uint8 = 127
                            _ = v19
                        }
                    }
                    _ = v14
                }
                var v20 int32 = -529
                var v21 uint = 255
                if false {
                    var v22 uint32 = 42
                    _ = v22
                } else {
                    v23 := GenS3{F0: "foo bar", F1: false, F2: 4294967295}
                    if (v11 < (v4 - v11)) {
                        switch (^(fi13)) {
                        case 118:
                            if (1 <= 255) {
                                var v24 uint64 = v4
                                v25 := v3[2147483648]
                                v26 := GenS1{F0: (v11 << ((v21) & 63))}
                                var v27 uint8 = 1
                                v11 += 1
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                var v28 int16 = 0
                                v6 = (v6 << ((v21) & 63))
                                _ = v28
                            }
                        default:
                            if false {
                                v29 := v3[65535]
                                if false {
                                    v6 = v6
                                    var v30 uint8 = 42
                                    _ = v30
                                }
                                var v31 uint8 = 128
                                _ = v29
                                _ = v31
                            }
                        }
                        if true {
                            v32 := GenS1{F0: v11}
                            _ = v32
                        } else {
                            var v33 uint = v21
                            _ = v33
                        }
                    } else {
                        if ((1e10 != 0.1) && (false && true)) {
                            v34 := GenS1{F0: 1}
                            fmt.Println(v6, v11)
                            _ = v34
                        } else {
                            var v35 string = ""
                            _ = v35
                        }
                        var v36 uint64 = v11
                        v37 := GenS2{F0: fi12, F1: 2147483647, F2: 107}
                        _ = v36
                        _ = v37
                    }
                    _ = v23
                }
                if false {
                    v38 := GenS4{F0: (v20 * v20), F1: 255}
                    switch fi13 {
                    case -15:
                        var v39 uint64 = (v11 + v11)
                        _ = v39
                    case -2147483648:
                        var v40 float32 = (0.0 * 0.0)
                        _ = v40
                    default:
                        var v41 string = ("" + ("\n\t" + ""))
                        _ = v41
                    }
                    _ = v38
                } else {
                    v6 -= (v6 % ((-1) | 1))
                    var v42 float64 = 0.0
                    _ = v42
                }
                _ = v20
                _ = v21
            default:
                v1 = append(v1, "hello")
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi13
        }
        _ = fi12
    }
    _ = len(v2)
    var v43 int8 = 82
    var v44 int16 = (v6 | v6)
    if (2 == 2) {
        v43 = v43
        var v45 int8 = v43
        _ = v45
    }
    v44 -= v44
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
