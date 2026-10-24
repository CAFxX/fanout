package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 string; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int32; F2 uint }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 uint64 = 42
    var v4 int = 128
    v5 := []int{(v4 / (((v4 & (v4 | -643))) | 1)), (v4 - (v4 - 0)), (v4 & (-(v4))), 42}
    var v6 uint64 = v3
    var v7 float32 = -2.25
    v8 := GenS1{F0: v3, F1: (v4 & (v4 | 589))}
    var v9 uint32 = 730
    v10 := GenS2{F0: 1}
    v11 := GenS1{F0: 581, F1: 32767}
    v3 = v3
    delete(v2, v1)
    var v12 int32 = -767
    v1 = v1
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v4, v4, v6)
        v3 -= (v3 * v3)
        switch (fi13 * (v4 * (fi13 / ((fi13) | 1)))) {
        case 42:
            v6 = v6
        case 292:
            fmt.Println(v12, v12)
        case 599:
            v12 = v12
        default:
            if (v4 == 32767) {
                if false {
                    if (1 < 0) {
                        if false {
                            if ((1 < -128) && false) {
                                v9 = v9
                            } else {
                                var v14 uint = 18446744073709551615
                                _ = v14
                            }
                            if false {
                                for fi15 := 0; fi15 < 12; fi15++ {
                                    v16 := GenS3{F0: (v3 - 308), F1: "hello", F2: 128}
                                    var v17 int64 = 1
                                    var v18 float32 = -2.25
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            }
                            for fi19 := 0; fi19 < 3; fi19++ {
                                if true {
                                    v20 := GenS2{F0: 42}
                                    v12 -= 2147483647
                                    v2[v1] = fi19
                                    var v21 float64 = (1e10 * -2.25)
                                    _ = v20
                                    _ = v21
                                } else {
                                    v22 := GenS2{F0: 127}
                                    _ = v22
                                }
                                var v23 float32 = (v7 * 1e10)
                                _ = fi19
                                _ = v23
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        fmt.Println(v3)
                    }
                } else {
                    v24 := GenS3{F0: 2, F1: ((((v1 + "a") + "") + (((v1 + "") + v1) + v1)) + v1), F2: 237}
                    for fi25 := 0; fi25 < 20; fi25++ {
                        fmt.Println(fi13, v1)
                        v7 = (((v7 + ((3.14159 * 0.0) - v7)) / ((v7 / -2.25) / v7)) + 1.5)
                        v26 := GenS3{F0: v6, F1: ("" + ""), F2: 169}
                        v9 -= 990
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                if false {
                    v9 = (v9 >> ((255) & 63))
                    for fi27 := 0; fi27 < 2; fi27++ {
                        _ = len(v5)
                        var v28 int64 = -1
                        _ = fi27
                        _ = v28
                    }
                    var v29 int8 = 84
                    _ = v29
                } else {
                    if true {
                        if true {
                            v30 := GenS4{F0: (v3 >> ((42) & 63)), F1: 42}
                            v31 := GenS5{F0: 0}
                            _ = v30
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 1; fi32++ {
                            if (false || (638 > 2)) {
                                v33 := GenS6{F0: v3}
                                if true {
                                    var v34 uint32 = v9
                                    _ = v34
                                } else {
                                    delete(v2, v1)
                                    var v35 int8 = 100
                                    delete(v2, "\n\t")
                                    _ = v35
                                }
                                var v36 uint = 2
                                _ = v33
                                _ = v36
                            }
                            v7 = v7
                            _ = fi32
                        }
                    }
                    if false {
                        fmt.Println(v9)
                    }
                    var v37 int32 = v12
                    _ = v37
                }
            }
        }
        _ = fi13
    }
    var v38 uint = 433
    for fi39 := 0; fi39 < 10; fi39++ {
        v5 = append(v5, (-(fi39)))
        if (!((fi39 >= -2147483648) && (!true))) {
            v1 = v1
            if true {
                var v40 uint64 = (v6 - (v6 / ((uint64(v38)) | 1)))
                var v41 int32 = 0
                _ = v40
                _ = v41
            } else {
                delete(v2, "\n\t")
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                var v43 int32 = v12
                v12 += (v43 << ((uint(v9)) & 63))
                switch v4 {
                case 9223372036854775807:
                    fmt.Println(v1, v4)
                case -600:
                    if (-32768 >= 32767) {
                        fmt.Println(v12, fi39, v43)
                    } else {
                        var v44 int = (fi42 ^ (v4 * -1))
                        _ = v44
                    }
                default:
                    if true {
                        if len(v5) > 0 {
                            v5[0] = fi42
                        }
                    }
                }
                _ = fi42
                _ = v43
            }
        } else {
            v45 := GenS1{F0: v6, F1: (v4 + (fi39 % ((200) | 1)))}
            _ = v45
        }
        _ = fi39
    }
    v38 -= uint(v9)
    var v46 int8 = 100
    if (1 <= 1) {
        if false {
            _ = len(v5)
            v47 := GenS7{F0: (^(v46)), F1: v12, F2: (v38 % (((v38 % ((v38) | 1))) | 1))}
            _ = v47
        }
        var v48 int = (^(v4))
        _ = v48
    } else {
        if len(v5) > 0 {
            v5[0] = (-(v4))
        }
    }
    if false {
        if true {
            for fi49 := 0; fi49 < 15; fi49++ {
                fmt.Println(fi49)
                switch (^(v4)) {
                case 155:
                    v38 += (^(v38))
                case 4294967295:
                    v3 = v3
                }
                v50 := GenS3{F0: (v3 / ((v3) | 1)), F1: "", F2: 2}
                _ = fi49
                _ = v50
            }
        }
        if (42 >= 42) {
            fmt.Println(v38, v9)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
