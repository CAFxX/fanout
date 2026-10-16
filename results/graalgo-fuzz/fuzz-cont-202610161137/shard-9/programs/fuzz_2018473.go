package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 int8; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int8; F2 uint8 }
type GenS7 struct { F0 bool; F1 uint64; F2 int8 }
type GenS8 struct { F0 uint8; F1 int64 }
func genFn0(p0 int16) int8 {
    p0 += p0
    var v7 uint32 = 370
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 uint8) int64 {
    v8 := GenS1{F0: 978, F1: 1}
    v9 := []int{-68}
    p0 += 204
    _ = len(v9)
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{"", "\n\t", "hello", "a"}
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    var v4 string = "\n\t"
    var v5 float32 = 0.1
    v6 := []string{}
    v10 := GenS1{F0: -2147483648, F1: -1}
    var v11 int = 0
    v12 := v2[0]
    if true {
        switch v4 {
        case "a":
            _ = len(v1)
        case "":
            delete(v2, v12)
        }
        var v13 int32 = 536
        _ = v13
    } else {
        var v14 int = v11
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v12)
        _ = v14
    }
    switch v4 {
    case "hello":
        var v15 float64 = 3.14159
        _ = v15
    case "foo bar":
        switch v4 {
        case "":
            var v16 int8 = (v3 - 127)
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 7; fi17++ {
                if true {
                    _ = len(v1)
                } else {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (18446744073709551615 >= 18446744073709551615) {
                            switch v4 {
                            case "\n\t":
                                v19 := GenS1{F0: 0, F1: -32768}
                                _ = v19
                            }
                            v4 = v4
                        } else {
                            v2[4294967295] = 2147483648
                        }
                        if (!true) {
                            v11 = fi17
                            v12 = v12
                        } else {
                            v4 = ""
                            var v20 float32 = 3.14159
                            v21 := GenS1{F0: 810, F1: 1}
                            _ = v20
                            _ = v21
                        }
                        v4 = "foo bar"
                        v22 := GenS2{F0: v4, F1: v12, F2: -9223372036854775808}
                        var v23 uint16 = 0
                        _ = fi18
                        _ = v22
                        _ = v23
                    }
                    var v24 uint = 799
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v24
                }
                delete(v2, v12)
                for fi25 := 0; fi25 < 2; fi25++ {
                    v26 := GenS1{F0: 0, F1: -32768}
                    _ = fi25
                    _ = v26
                }
                _ = fi17
            }
        }
    default:
        fmt.Println(v3, v5, v3)
    }
    switch v4 {
    case "\n\t":
        v27 := GenS3{F0: v4, F1: false}
        _ = v27
    default:
        v4 = v4
    }
    v3 = v3
    v4 = "foo bar"
    for fi28 := 0; fi28 < 9; fi28++ {
        v1 = append(v1, v4)
        fmt.Println(v11, v11, v4)
        fmt.Println(v4)
        _ = fi28
    }
    if false {
        var v29 uint64 = 47
        if ((0.1 + -2.25) != -2.25) {
            fmt.Println(v3, v12, v3)
            v3 -= v3
            v11 -= 1
        } else {
            v30 := GenS1{F0: 0, F1: 42}
            var v31 int = (v11 | (v11 >> ((255) & 63)))
            _ = v30
            _ = v31
        }
        _ = v29
    } else {
        var v32 uint64 = 0
        _ = v32
    }
    if false {
        v11 -= -1
        for fi33 := 0; fi33 < 18; fi33++ {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 float64 = 0.1
                v35 = (v35 / 1e10)
                if false {
                    v11 = fi34
                } else {
                    if true {
                        v36 := v2[v12]
                        _ = v36
                    } else {
                        if true {
                            v37 := GenS4{F0: "a", F1: v3, F2: (v12 >> ((887) & 63))}
                            fmt.Println(v4, v35)
                            if (!(!true)) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v38 string = "a"
                                var v39 int64 = 575
                                v4 = v4
                                var v40 uint64 = 18446744073709551615
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            v11 -= int(v12)
                            _ = v37
                        } else {
                            if (0 == 2) {
                                v4 = (v4 + (v4 + (v4 + v4)))
                                var v41 int8 = genFn0(-1)
                                switch v4 {
                                case "":
                                    var v42 uint8 = 255
                                    _ = v42
                                case "hello":
                                    var v43 int64 = -9223372036854775808
                                    _ = v43
                                case "a":
                                    v2[v12] = v12
                                }
                                var v44 uint = uint(v12)
                                v45 := v2[v12]
                                _ = v41
                                _ = v44
                                _ = v45
                            } else {
                                if ((v3 & v3) > 100) {
                                    var v46 int32 = 167
                                    switch fi34 {
                                    case 0:
                                        if (!((65535 != 2) && (24 <= 4294967295))) {
                                            v5 = float32(v35)
                                            v47 := v2[v12]
                                            _ = v47
                                        } else {
                                            if (v12 > 2147483647) {
                                                var v48 float64 = v35
                                                v49 := GenS5{F0: -29}
                                                v3 = (v3 + -1)
                                                var v50 int64 = -1
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                    }
                                    v51 := v2[v12]
                                    _ = v46
                                    _ = v51
                                }
                                switch 9223372036854775807 {
                                case 127:
                                    fmt.Println(v3, fi33, fi33)
                                }
                                if ((v3 << ((65535) & 63)) >= (v3 | (v3 >> ((18446744073709551615) & 63)))) {
                                    v52 := GenS6{F0: true, F1: genFn0(42), F2: 42}
                                    v53 := GenS7{F0: (276 != 675), F1: 4294967295, F2: genFn0(698)}
                                    _ = v52
                                    _ = v53
                                }
                            }
                            fmt.Println(v4, v11, v12)
                        }
                    }
                    if false {
                        v35 = (1.5 * (((v35 * 1e10) * (v35 - 0.1)) + (0.0 + v35)))
                    } else {
                        var v54 uint32 = v12
                        v55 := v2[(v12 * v54)]
                        var v56 float64 = v35
                        v5 -= v5
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    _ = len(v6)
                }
                _ = fi34
                _ = v35
            }
            v4 = "foo bar"
            for fi57 := 0; fi57 < 13; fi57++ {
                switch int(v12) {
                case -2147483648:
                    fmt.Println(fi33)
                }
                v58 := GenS8{F0: 1, F1: -9223372036854775808}
                _ = fi57
                _ = v58
            }
            _ = fi33
        }
    }
    for fi59 := 0; fi59 < 20; fi59++ {
        v3 = (v3 + 1)
        if true {
            fmt.Println(v4)
        }
        _ = fi59
    }
    v60 := GenS8{F0: 1, F1: 42}
    var v61 int16 = 32767
    fmt.Println(v11)
    v1 = append(v1, (v4 + v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
