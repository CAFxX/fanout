package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint16) string {
    if (316 > 42) {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 100, F1: 920}
            _ = fi6
            _ = v7
        }
        fmt.Println(p0)
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                var v9 uint64 = 18446744073709551615
                var v10 int32 = 2147483647
                _ = v10
                _ = v9
            }
            switch fi8 {
            case -1:
                var v11 uint64 = 798
                _ = v11
            case 172:
                var v12 string = ("\n\t" + "hello")
                _ = v12
            default:
                fmt.Println(fi8, p0)
            }
            _ = fi8
        }
        fmt.Println(p0)
        var v13 int16 = 42
        _ = v13
    }
    if (243 > 255) {
        var v14 int32 = -1
        p0 = (^(p0))
        _ = v14
    } else {
        var v15 bool = (1 >= 788)
        _ = v15
    }
    _ = p0
    return "foo bar"
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 string = ("foo bar" + "foo bar")
    var v5 int64 = 0
    v16 := []int{127, 1, 32768, -1}
    v5 = (v5 % ((0) | 1))
    v5 = (-(v5))
    var v17 uint64 = 724
    for fi18 := 0; fi18 < 20; fi18++ {
        v17 = 18446744073709551615
        _ = fi18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int8 = 68
        v1 = v1
        v21 := GenS2{F0: 2, F1: 60, F2: 2}
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v5, v17)
    v17 = v17
    v22 := GenS1{F0: 100, F1: 18446744073709551615}
    if true {
        var v23 string = (v4 + ("hello" + "\n\t"))
        fmt.Println(v17, v17, v5)
        v17 += v17
        _ = v23
    } else {
        switch "\n\t" {
        case "\n\t":
            v4 = ("foo bar" + "\n\t")
        case "":
            if (4294967295 <= 18446744073709551615) {
                for fi24 := 0; fi24 < 20; fi24++ {
                    var v25 uint16 = 0
                    v5 -= (-(v5))
                    _ = fi24
                    _ = v25
                }
                var v26 int64 = v5
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := GenS3{F0: fi27}
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v5 += (v26 - (-(v26)))
                        switch v4 {
                        case "":
                            v30 := GenS2{F0: 42, F1: 227, F2: 512}
                            _ = v30
                        case "a":
                            switch "\n\t" {
                            case "\n\t":
                                v5 += v26
                            case "":
                                v4 = ((v4 + v4) + v4)
                            case "hello":
                                if false {
                                    v1 = v1
                                    v16 = append(v16, fi27)
                                    var v31 int = fi27
                                    v16 = append(v16, v31)
                                    v3[999] = v4
                                    _ = v31
                                }
                            }
                        case "foo bar":
                            var v32 int8 = 0
                            _ = v32
                        default:
                            switch (fi27 * fi27) {
                            case -9223372036854775808:
                                v33 := GenS4{F0: 131}
                                _ = v33
                            case 9223372036854775807:
                                v5 = -836
                            }
                        }
                        if (!false) {
                            v34 := v2[2147483647]
                            var v35 int = (fi29 / ((fi29) | 1))
                            var v36 float64 = 1e10
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = fi29
                    }
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
        case "foo bar":
            var v37 uint32 = 1
            _ = v37
        default:
            for fi38 := 0; fi38 < 6; fi38++ {
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        for fi39 := 0; fi39 < 17; fi39++ {
                            var v40 uint64 = 177
                            _ = fi39
                            _ = v40
                        }
                    case "foo bar":
                        v17 += (v17 - 1)
                    case "":
                        v16 = append(v16, (fi38 >> ((1) & 63)))
                    }
                } else {
                    v41 := GenS5{F0: -402, F1: 42, F2: v17}
                    var v42 uint8 = 127
                    v43 := GenS5{F0: -685, F1: 2, F2: (v17 << ((65535) & 63))}
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi38
            }
        }
    }
    if (v4 >= v4) {
        v4 = ("a" + v4)
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        v1 = v1
        v4 = v4
        _ = fi44
    }
    v45 := GenS2{F0: 0, F1: 26, F2: 65535}
    v16 = append(v16, 1)
    var v46 int16 = 42
    switch (((v4 + ("foo bar" + "foo bar")) + ("hello" + v4)) + "foo bar") {
    case "":
        if false {
            switch (v4 + (("hello" + "hello") + v4)) {
            case "foo bar":
                v47 := GenS4{F0: 128}
                _ = v47
            default:
                var v48 float32 = 0.0
                _ = v48
            }
            if ((255 == 1) && ((v5 | 1) > (v5 + (v5 - (v5 >> ((4294967295) & 63)))))) {
                v49 := GenS6{F0: (-6 < 100), F1: -128, F2: 47}
                fmt.Println(v46, v17)
                var v50 uint8 = 42
                _ = v49
                _ = v50
            }
        } else {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            if len(v16) > 0 {
                v16[0] = 248
            }
            v51 := GenS7{F0: (!false)}
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 uint = 42
        v4 = v4
        v53 = v53
        _ = fi52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
