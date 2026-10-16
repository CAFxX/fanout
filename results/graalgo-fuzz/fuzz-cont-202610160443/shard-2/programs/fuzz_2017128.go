package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 int16; F1 int }
func genFn0(p0 string, p1 uint32) int32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            if (true && (p0 >= (("" + ("" + (p0 + p0))) + (p0 + p0)))) {
                var v6 bool = false
                if v6 {
                    v7 := GenS1{F0: 4294967295, F1: 2}
                    v8 := GenS2{F0: 183, F1: 418}
                    _ = v7
                    _ = v8
                }
                var v9 string = p0
                _ = v6
                _ = v9
            }
            var v10 uint8 = 239
            var v11 uint8 = 154
            _ = fi5
            _ = v10
            _ = v11
        }
        _ = fi4
    }
    switch 463 {
    case 42:
        v12 := []int{1, -214}
        _ = v12
    case -741:
        v13 := []int{1, 32767, -9223372036854775808}
        _ = v13
    case -1:
        var v14 uint32 = p1
        _ = v14
    default:
        if false {
            v15 := GenS2{F0: 4294967295, F1: int64(p1)}
            v16 := []int{-612, 2147483647, -89, 128}
            v17 := GenS1{F0: 534, F1: (^(p1))}
            _ = v15
            _ = v16
            _ = v17
        } else {
            var v18 int32 = -898
            v19 := GenS3{F0: 253, F1: (!(false || true))}
            switch "a" {
            case "foo bar":
                p1 = p1
            }
            _ = v18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v20 int = -725
    v21 := v2[46]
    v1 = (^(v1))
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 uint16 = 42
        v21 = v21
        _ = fi22
        _ = v23
    }
    if v21 {
        switch (v20 | v20) {
        case 94:
            v24 := GenS4{F0: -989}
            _ = v24
        case 42:
            if (42 >= 32767) {
                if (false && v21) {
                    var v25 bool = v21
                    _ = v25
                }
                var v26 int = v20
                _ = v26
            }
        case 9223372036854775807:
            if false {
                v27 := v2[4294967295]
                var v28 int8 = 47
                v20 = v20
                _ = v27
                _ = v28
            } else {
                if ((v20 << ((355) & 63)) > v20) {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        var v30 int32 = 42
                        for fi31 := 0; fi31 < 6; fi31++ {
                            fmt.Println(v20, v30, fi29)
                            var v32 float64 = 0.1
                            if (-885 != 578) {
                                switch ("" + ("" + "\n\t")) {
                                case "a":
                                    v30 = genFn0("foo bar", 2)
                                case "foo bar":
                                    v33 := GenS5{F0: 1, F1: "a", F2: 0}
                                    _ = v33
                                case "":
                                    if false {
                                        v34 := []int{-818, (fi29 << ((258) & 63))}
                                        _ = v34
                                    } else {
                                        v20 = fi31
                                        v21 = false
                                        fmt.Println(v20, v32)
                                        var v35 uint = 881
                                        _ = v35
                                    }
                                }
                                if false {
                                    v30 = 42
                                    var v36 string = "foo bar"
                                    _ = v36
                                }
                                v37 := GenS4{F0: 0}
                                _ = v37
                            }
                            v38 := GenS2{F0: 1, F1: -647}
                            v21 = (255 >= v1)
                            if (v21 || true) {
                                v1 = 127
                                v39 := GenS6{F0: 32767, F1: (fi29 + 128)}
                                v40 := v2[2147483647]
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v1, fi31)
                                case "\n\t":
                                    v41 := GenS5{F0: 0, F1: ("foo bar" + "a"), F2: 132}
                                    _ = v41
                                case "hello":
                                    fmt.Println(v21)
                                }
                                if v21 {
                                    var v42 float64 = v32
                                    v20 = 42
                                    _ = v42
                                }
                                _ = v39
                                _ = v40
                            } else {
                                var v43 float64 = v32
                                _ = v43
                            }
                            _ = fi31
                            _ = v32
                            _ = v38
                        }
                        _ = fi29
                        _ = v30
                    }
                } else {
                    switch v20 {
                    case 778:
                        if ((((3.14159 * (0.1 / 1.0)) + ((float64(0.1) / (1.5 * -2.25)) + -2.25)) / 0.1) > 0.0) {
                            v20 = v20
                            v44 := GenS2{F0: 18446744073709551615, F1: 0}
                            _ = v44
                        }
                    default:
                        var v45 string = ""
                        _ = v45
                    }
                }
            }
        default:
            v1 = v1
        }
    }
    var v46 int64 = 0
    v1 += (v1 >> ((uint(v1)) & 63))
    v20 += 127
    switch "\n\t" {
    case "foo bar":
        v20 = v20
    }
    v46 -= v46
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v46
}
