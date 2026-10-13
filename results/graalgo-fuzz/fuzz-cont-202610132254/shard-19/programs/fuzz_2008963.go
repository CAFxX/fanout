package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int; F1 int; F2 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 int16; F2 int32 }
func genFn0(p0 int8) uint8 {
    v7 := GenS1{F0: 127, F1: (-(p0)), F2: 65535}
    var v8 int = -209
    _ = p0
    _ = v7
    _ = v8
    return 176
}

func main() {
    var v1 float32 = (((0.0 * 0.0) + 1.5) - float32(1e10))
    var v2 uint32 = 42
    var v3 int64 = 0
    v4 := []uint32{v2, (v2 ^ 2147483648)}
    var v5 string = "hello"
    var v6 int32 = 42
    var v9 int16 = 42
    var v10 uint32 = (v2 + (v2 / ((1) | 1)))
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS2{F0: "hello"}
        _ = fi11
        _ = v12
    }
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            if ((!false) && false) {
                if false {
                    var v16 uint = 65535
                    if (((v9 - -912) < (v9 ^ int16(v16))) || true) {
                        if true {
                            var v17 uint32 = (v10 / ((315) | 1))
                            _ = v17
                        } else {
                            v3 += v3
                            switch fi14 {
                            case 824:
                                v1 = v1
                            case -465:
                                var v18 uint32 = v2
                                _ = v18
                            }
                        }
                    } else {
                        v19 := GenS3{F0: true}
                        if (!(v3 > (v3 % (((v3 - (v3 + (v3 % ((v3) | 1))))) | 1)))) {
                            if ((v3 & -9223372036854775808) >= -330) {
                                v20 := GenS4{F0: "foo bar"}
                                var v21 int32 = v6
                                _ = v20
                                _ = v21
                            }
                            v10 = (v2 - (v10 + (v10 + v2)))
                            fmt.Println(v6)
                        }
                        fmt.Println(v10, v5, v6)
                        if (928 != 65535) {
                            var v22 int64 = -1
                            if (-128 == -1) {
                                v22 = (-(v3))
                            }
                            _ = v22
                        }
                        _ = v19
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v16
                }
                v23 := GenS4{F0: ""}
                if false {
                    var v24 uint64 = 4294967295
                    var v25 uint8 = 255
                    _ = v24
                    _ = v25
                } else {
                    v13 = v13
                }
                _ = v23
            } else {
                v4 = append(v4, (^(v2)))
                fmt.Println(fi14, v5)
            }
            _ = fi15
        }
        var v26 uint64 = 417
        if (false && true) {
            var v27 int = (fi14 | (fi14 / ((23) | 1)))
            var v28 int64 = int64(v10)
            for fi29 := 0; fi29 < 5; fi29++ {
                if (false || false) {
                    v30 := GenS5{F0: "\n\t", F1: (^(v6))}
                    switch v27 {
                    case -9223372036854775808:
                        if false {
                            var v31 uint32 = uint32(v26)
                            var v32 int32 = 653
                            _ = len(v4)
                            v33 := GenS6{F0: fi29, F1: 127, F2: 1}
                            switch ("a" + ((v5 + ("\n\t" + "\n\t")) + v5)) {
                            case "a":
                                _ = len(v4)
                            }
                            v13 += float64(((v1 + 1.5) - v1))
                            var v34 int64 = -229
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            switch (v27 * 295) {
                            case 707:
                                v5 = "a"
                            }
                        }
                    }
                    _ = v30
                } else {
                    v26 = (v26 * 18446744073709551615)
                    if (-110 > 127) {
                        if (false || (true && true)) {
                            switch v5 {
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "hello":
                                v35 := GenS1{F0: (fi14 >> ((42) & 63)), F1: 42, F2: v10}
                                _ = v35
                            }
                            v9 = (^(v9))
                            var v36 int8 = -63
                            _ = v36
                        }
                    } else {
                        var v37 int = (v27 & 127)
                        v38 := GenS2{F0: (v5 + (((v5 + ((v5 + (v5 + "a")) + v5)) + v5) + ((v5 + "\n\t") + "\n\t")))}
                        _ = v37
                        _ = v38
                    }
                }
                v4 = append(v4, v10)
                _ = fi29
            }
            _ = v27
            _ = v28
        }
        _ = fi14
        _ = v26
    }
    v39 := GenS1{F0: -727, F1: 100, F2: (^(v10))}
    var v40 uint = 1
    var v41 bool = (!true)
    _ = len(v4)
    v10 = (^(v2))
    v1 += 1e10
    switch (v5 + "a") {
    case "\n\t":
        for fi42 := 0; fi42 < 10; fi42++ {
            if (32767 == 2) {
                if (((int16(v40) >= -318) || true) && v41) {
                    for fi43 := 0; fi43 < 11; fi43++ {
                        _ = len(v4)
                        if len(v4) > 0 {
                            v4[0] = v10
                        }
                        if len(v4) > 0 {
                            v4[0] = 703
                        }
                        _ = fi43
                    }
                    v44 := GenS5{F0: (v5 + (v5 + ((v5 + v5) + "\n\t"))), F1: v6}
                    _ = v44
                } else {
                    v4 = append(v4, (v10 & 560))
                    v45 := GenS7{F0: v2}
                    _ = v45
                }
            }
            v46 := GenS8{F0: v3, F1: (v9 / ((v9) | 1)), F2: v6}
            _ = fi42
            _ = v46
        }
    }
    v10 -= (v10 ^ 2147483647)
    fmt.Println(v41, v13)
    fmt.Println(v1, v40)
    var v47 bool = ((!false) || (!false))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
