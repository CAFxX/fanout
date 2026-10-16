package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 int32 }
type GenS4 struct { F0 bool; F1 uint32; F2 int }
func genFn0(p0 uint16, p1 uint8) int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p1 -= p1
        _ = fi6
    }
    p1 -= (p1 << ((1) & 63))
    p0 = 255
    p0 = (p0 & p0)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int16, p1 uint32) int16 {
    p1 = 2
    fmt.Println(p1)
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: -197}
        var v10 int16 = p0
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!false) {
        fmt.Println(p1)
        switch -1 {
        case 32767:
            v7 = (v7 | v7)
        }
        fmt.Println(p0, p0)
    } else {
        if false {
            if true {
                switch 596 {
                case 248:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int64 = -127
                        _ = fi11
                        _ = v12
                    }
                }
                p1 = 0
            } else {
                var v13 int16 = genFn0(569, v7)
                _ = v13
            }
            if false {
                v14 := GenS1{F0: 442}
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: -9223372036854775808}
                    var v17 string = "foo bar"
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 uint32 = 2
                _ = v14
                _ = v18
            } else {
                var v19 int32 = 844
                _ = v19
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    v1 := []string{("foo bar" + ""), "", "foo bar"}
    v2 := make(map[uint32]int)
    var v3 uint8 = 139
    var v4 uint32 = 2
    v5 := []string{"", ("\n\t" + "foo bar")}
    var v20 uint32 = (v4 % ((v4) | 1))
    if true {
        var v21 float32 = ((3.14159 * (-2.25 + 0.1)) + ((0.1 / 0.1) + (0.0 / 1.0)))
        _ = v21
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v22 := GenS1{F0: -9223372036854775808}
    if (4294967295 < 782) {
        v3 = (v3 % ((v3) | 1))
        for fi23 := 0; fi23 < 14; fi23++ {
            fmt.Println(fi23)
            v4 += 65535
            var v24 uint32 = 293
            _ = fi23
            _ = v24
        }
        var v25 uint32 = v4
        _ = v25
    } else {
        switch "a" {
        case "":
            v26 := GenS2{F0: 65535}
            _ = v26
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            _ = len(v2)
            var v28 uint64 = 350
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            var v30 int16 = 0
            var v31 uint32 = 479
            v3 = (v3 % ((v3) | 1))
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    var v32 uint = 0
    if len(v5) > 0 {
        v5[0] = "a"
    }
    if true {
        if (true || false) {
            var v33 uint8 = uint8(v32)
            _ = v33
        }
    } else {
        v1 = append(v1, ("a" + ("\n\t" + "foo bar")))
        _ = len(v1)
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 int16 = genFn1(genFn1(42, (v4 | 42)), 2)
        v36 := GenS3{F0: -108, F1: "\n\t", F2: 2147483647}
        _ = fi34
        _ = v35
        _ = v36
    }
    if false {
        var v37 uint64 = uint64(v32)
        _ = v37
    } else {
        if ((!true) || false) {
            v1 = append(v1, "hello")
        } else {
            if false {
                v38 := GenS3{F0: 0, F1: "\n\t", F2: 42}
                if (-1 == 845) {
                    v20 += v20
                    v3 = uint8(v4)
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v20, v20)
                    case "":
                        _ = len(v1)
                    }
                } else {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v2[(v20 | (v20 ^ v4))] = fi39
                        switch int(v20) {
                        case 9223372036854775807:
                            if true {
                                fmt.Println(v4)
                                switch fi39 {
                                case 42:
                                    var v40 int16 = -714
                                    _ = v40
                                }
                                _ = len(v2)
                            }
                        case -1:
                            for fi41 := 0; fi41 < 5; fi41++ {
                                if (false || false) {
                                    v42 := GenS4{F0: (false && false), F1: v20, F2: (^(fi39))}
                                    v4 += v4
                                    _ = v42
                                }
                                _ = fi41
                            }
                        case 127:
                            for fi43 := 0; fi43 < 6; fi43++ {
                                fmt.Println(v20, fi39, v32)
                                if false {
                                    v2[(v4 << (((v32 ^ (^(v32)))) & 63))] = fi39
                                    if (fi39 >= fi39) {
                                        v3 -= (v3 / ((v3) | 1))
                                    } else {
                                        fmt.Println(v4)
                                        v44 := GenS3{F0: 807, F1: "foo bar", F2: 2147483647}
                                        v45 := v2[(v20 % ((4294967295) | 1))]
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                _ = fi43
                            }
                        }
                        var v46 uint32 = (v20 >> ((v32) & 63))
                        _ = fi39
                        _ = v46
                    }
                }
                _ = v38
            }
            for fi47 := 0; fi47 < 15; fi47++ {
                v32 -= v32
                _ = fi47
            }
            v48 := GenS2{F0: 2}
            _ = v48
        }
        v20 = (v4 / ((38) | 1))
        var v49 uint32 = 2147483647
        _ = v49
    }
    v4 = 2147483647
    if ((!true) && false) {
        if (-2.25 < float32(float64(-2.25))) {
            v50 := v2[uint32(v3)]
            _ = v50
        }
        fmt.Println(v32)
        var v51 uint8 = 2
        _ = v51
    }
    if false {
        _ = len(v1)
    } else {
        fmt.Println(v4, v4)
    }
    v20 -= (^(v20))
    _ = len(v2)
    fmt.Println(v20, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
