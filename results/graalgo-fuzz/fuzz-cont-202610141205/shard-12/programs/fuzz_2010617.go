package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint64, p1 int16) int {
    v4 := GenS1{F0: 562}
    v5 := GenS1{F0: 2147483647}
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 455}
            v9 := []int{fi6, fi6, fi7, (fi6 ^ 42)}
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 float64 = float64(-2.25)
        _ = fi6
        _ = v10
    }
    p0 += p0
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 int = -2147483648
    v2 := []uint32{1, 110, 131, 0}
    v3 := []string{"foo bar", "hello"}
    switch (v1 ^ v1) {
    case 974:
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                var v12 uint16 = 255
                _ = v12
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                if (42 != 2) {
                    v1 = 4294967295
                } else {
                    var v14 uint16 = 178
                    v15 := GenS1{F0: (fi11 - 127)}
                    switch (("a" + "a") + "foo bar") {
                    case "\n\t":
                        var v16 uint = 2
                        _ = v16
                    case "":
                        fmt.Println(fi13, fi11, v1)
                    default:
                        v1 -= (v1 >> ((65535) & 63))
                        var v17 string = "foo bar"
                        switch ((v17 + v17) + ((((v17 + "foo bar") + ((v17 + v17) + v17)) + (((v17 + ("hello" + (v17 + v17))) + "") + (v17 + "a"))) + ("foo bar" + ("hello" + (v17 + ((("" + (v17 + v17)) + v17) + ("\n\t" + "hello"))))))) {
                        case "":
                            v1 += (fi13 * v1)
                        case "\n\t":
                            v18 := GenS2{F0: 0, F1: 42, F2: 1}
                            _ = v18
                        case "hello":
                            var v19 uint32 = 2
                            _ = v19
                        }
                        v14 -= v14
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
            if false {
                v1 += (v1 % ((42) | 1))
                v3 = append(v3, "hello")
            } else {
                switch (v1 << ((4294967295) & 63)) {
                case 42:
                    var v20 uint8 = 42
                    _ = v20
                default:
                    if true {
                        var v21 uint = 65535
                        for fi22 := 0; fi22 < 13; fi22++ {
                            var v23 int64 = 745
                            if false {
                                v21 -= v21
                                v24 := GenS2{F0: v21, F1: 976, F2: -128}
                                _ = v24
                            }
                            var v25 uint64 = 2
                            v21 -= (v21 * (v21 * 4294967295))
                            _ = fi22
                            _ = v23
                            _ = v25
                        }
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = v21
                    } else {
                        v2 = append(v2, 65535)
                        if true {
                            if true {
                                var v26 string = "a"
                                _ = v26
                            } else {
                                v1 = 1
                            }
                            var v27 int64 = 657
                            v1 -= v1
                            _ = v27
                        }
                    }
                }
                for fi28 := 0; fi28 < 15; fi28++ {
                    switch "a" {
                    case "\n\t":
                        v1 = fi11
                    case "hello":
                        if true {
                            if false {
                                v1 = (^(fi28))
                                v3 = append(v3, "foo bar")
                                var v29 uint = 2
                                _ = v29
                            } else {
                                fmt.Println(fi11, fi11, fi11)
                                var v30 int = (v1 << ((875) & 63))
                                if false {
                                    v1 += v1
                                    v30 -= v30
                                    var v31 bool = (32767 > 430)
                                    _ = v31
                                }
                                _ = v30
                            }
                            if false {
                                v3 = append(v3, "foo bar")
                                var v32 int = (-(v1))
                                _ = v32
                            }
                        }
                    default:
                        var v33 float64 = (float64(((0.1 + 0.1) * 3.14159)) / 1e10)
                        _ = v33
                    }
                    var v34 int16 = -164
                    _ = fi28
                    _ = v34
                }
            }
            _ = fi11
        }
    case 9223372036854775807:
        switch "a" {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        case "":
            v35 := GenS2{F0: 316, F1: 1000, F2: 68}
            _ = v35
        case "\n\t":
            v36 := GenS2{F0: 4294967295, F1: 283, F2: -92}
            _ = v36
        }
    default:
        var v37 int8 = 94
        _ = v37
    }
    v1 -= -1
    var v38 int32 = 1
    v1 -= genFn0(0, 806)
    for fi39 := 0; fi39 < 17; fi39++ {
        for fi40 := 0; fi40 < 20; fi40++ {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            v41 := GenS3{F0: 42, F1: 32767, F2: 12}
            fmt.Println(v38, fi40)
            _ = fi40
            _ = v41
        }
        v42 := GenS4{F0: -32768, F1: (!true)}
        v43 := GenS5{F0: 314}
        _ = fi39
        _ = v42
        _ = v43
    }
    var v44 uint = 932
    v45 := GenS2{F0: v44, F1: 42, F2: -128}
    var v46 float64 = float64(1.5)
    v3 = append(v3, "\n\t")
    if (false || (((v1 & (v1 * (v1 ^ 343))) == genFn0(824, -628)) && (false && (v46 > 1.5)))) {
        fmt.Println(v38, v1, v46)
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        fmt.Println(v44, v1, v46)
        if false {
            var v48 int16 = 1000
            _ = v48
        } else {
            v46 += (((0.1 - v46) / v46) * v46)
            _ = len(v3)
            switch "foo bar" {
            case "":
                fmt.Println(v44, v44)
            case "foo bar":
                for fi49 := 0; fi49 < 17; fi49++ {
                    v1 += (fi47 ^ (fi49 + 2147483647))
                    _ = fi49
                }
            case "a":
                var v50 uint64 = 2
                _ = v50
            default:
                var v51 uint32 = 2147483648
                _ = v51
            }
        }
        _ = fi47
    }
    fmt.Println(v44, v44, v46)
    v1 = 32768
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    for fi52 := 0; fi52 < 8; fi52++ {
        v53 := GenS6{F0: -128}
        for fi54 := 0; fi54 < 9; fi54++ {
            var v55 string = ("" + "\n\t")
            v44 = v44
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v44
    _ = v45
    _ = v46
}
