package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int32 }
type GenS7 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 uint32 = 790
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int64 = -9223372036854775808
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 418
    fmt.Println(v6)
    v3 += v3
    v7 := GenS1{F0: 179, F1: 884, F2: 32767}
    v8 := GenS2{F0: ((("a" + "") + "foo bar") + "\n\t"), F1: "\n\t"}
    v9 := GenS3{F0: 42, F1: 570, F2: 288}
    var v10 bool = false
    switch -696 {
    case 128:
        if (!(true || v10)) {
            fmt.Println(v3, v10)
        } else {
            _ = len(v2)
            v3 = 65535
            switch "foo bar" {
            case "hello":
                if true {
                    v11 := GenS3{F0: 9223372036854775807, F1: 160, F2: -1}
                    var v12 float32 = 0.0
                    _ = v11
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v14 := []int{(fi13 << ((42) & 63)), 0}
                        v15 := GenS4{F0: fi13}
                        if (v10 && (0 != 1)) {
                            fmt.Println(v3, v10)
                            for fi16 := 0; fi16 < 4; fi16++ {
                                fmt.Println(v10, fi13)
                                _ = fi16
                            }
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v1[v3] = (v3 - uint32(v6))
                }
            case "\n\t":
                for fi17 := 0; fi17 < 12; fi17++ {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        switch "\n\t" {
                        case "a":
                            v19 := GenS3{F0: 152, F1: 255, F2: 32767}
                            _ = v19
                        case "\n\t":
                            v20 := GenS5{F0: 32767, F1: int64(fi18), F2: 773}
                            _ = v20
                        case "hello":
                            var v21 uint32 = v3
                            _ = v21
                        default:
                            switch 0 {
                            case 32768:
                                switch "hello" {
                                case "foo bar":
                                    var v22 uint32 = v3
                                    _ = v22
                                case "hello":
                                    var v23 int = (fi17 + (fi18 - (fi17 | 32767)))
                                    _ = v23
                                default:
                                    v10 = true
                                }
                            case 1:
                                v6 -= v6
                                fmt.Println(v3, v3, v6)
                            }
                        }
                        _ = fi18
                    }
                    _ = fi17
                }
            case "":
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 int32 = 292
                    _ = fi24
                    _ = v25
                }
            default:
                var v26 uint64 = (v6 ^ (v6 & v6))
                _ = v26
            }
        }
    case -1:
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS6{F0: 2, F1: 255, F2: -1}
            for fi29 := 0; fi29 < 16; fi29++ {
                v10 = false
                var v30 int32 = -1
                _ = fi29
                _ = v30
            }
            v6 = (v6 - 1)
            _ = fi27
            _ = v28
        }
    case -341:
        v31 := []int{1, 127, 191}
        _ = v31
    default:
        v32 := v1[(^(v3))]
        _ = v32
    }
    fmt.Println(v6)
    fmt.Println(v6, v6)
    fmt.Println(v3)
    var v33 int = -1
    v6 -= v6
    v33 -= 127
    switch "\n\t" {
    case "\n\t":
        switch "hello" {
        case "\n\t":
            v34 := GenS7{F0: v10, F1: int64(v33)}
            _ = v34
        case "foo bar":
            v35 := []int{42, (v33 & int(v6))}
            _ = v35
        default:
            if v10 {
                fmt.Println(v6)
                if ((28 >= 2147483647) && (false || v10)) {
                    for fi36 := 0; fi36 < 6; fi36++ {
                        var v37 string = "a"
                        if true {
                            for fi38 := 0; fi38 < 8; fi38++ {
                                var v39 string = ("a" + (v37 + ("\n\t" + v37)))
                                v6 = 2
                                var v40 int16 = -426
                                if (!((v10 && (1e10 <= 1e10)) || (!(v3 > (^(v3)))))) {
                                    var v41 int32 = -2147483648
                                    v37 = v39
                                    _ = v41
                                } else {
                                    var v42 bool = v10
                                    fmt.Println(fi36, v10)
                                    var v43 uint8 = 12
                                    v10 = ((((0.0 / 1.0) / 0.1) > 0.1) || ((^(v6)) > (v6 / (((^(v6))) | 1))))
                                    var v44 int16 = (v40 * (v40 & (v40 + (v40 - v40))))
                                    var v45 uint64 = (v6 * 4294967295)
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        }
                        v46 := GenS1{F0: (v3 % ((v3) | 1)), F1: v3, F2: 2}
                        _ = fi36
                        _ = v37
                        _ = v46
                    }
                } else {
                    var v47 uint8 = 1
                    v6 -= (v6 - 42)
                    var v48 bool = (("a" + "a") < "hello")
                    _ = v47
                    _ = v48
                }
                var v49 string = ("hello" + "foo bar")
                _ = v49
            } else {
                switch 4294967295 {
                case 127:
                    var v50 int = v33
                    _ = v50
                }
            }
        }
    case "":
        v51 := GenS2{F0: "a", F1: "hello"}
        _ = v51
    case "hello":
        v3 = v3
    default:
        v52 := []int{}
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
