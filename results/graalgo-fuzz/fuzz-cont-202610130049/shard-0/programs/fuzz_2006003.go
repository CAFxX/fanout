package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]bool)
    var v3 int8 = 127
    v4 := []string{(("" + "foo bar") + "\n\t"), ""}
    v5 := make(map[uint32]bool)
    var v6 string = ""
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                v6 = "a"
                var v9 float64 = 1e10
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                if false {
                    var v11 float64 = 0.0
                    var v12 uint64 = 0
                    var v13 int32 = (v1 + v1)
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    if true {
                        v6 = ""
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi10
            }
            if ((false || (v1 < v1)) || true) {
                var v14 int8 = (^(v3))
                _ = v14
            } else {
                switch v6 {
                case "":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 int8 = (-(v3))
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 int8 = v3
                if true {
                    if false {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            fmt.Println(v1, fi18, fi7)
                            v19 := GenS1{F0: ((v6 + "a") + (v6 + "hello")), F1: 0}
                            v4 = append(v4, (v6 + v6))
                            var v20 int16 = -1
                            v21 := GenS1{F0: v6, F1: -3}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        v3 += 127
                        for fi22 := 0; fi22 < 15; fi22++ {
                            if false {
                                v5[847] = true
                            } else {
                                if true {
                                    switch (v6 + (v6 + ((v6 + v6) + "\n\t"))) {
                                    case "foo bar":
                                        v23 := v2[-225]
                                        _ = v23
                                    }
                                    if (true && ((float32(1.5) + 0.1) == float32(3.14159))) {
                                        v24 := GenS2{F0: 831}
                                        v25 := v2[42]
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 uint32 = 4294967295
                                    if true {
                                        var v27 uint16 = 577
                                        v26 = (^(v26))
                                        fmt.Println(v6, v26, v1)
                                        _ = v27
                                    } else {
                                        var v28 uint32 = (v26 << ((0) & 63))
                                        v2[fi22] = false
                                        var v29 float32 = 0.0
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v26
                                } else {
                                    v30 := GenS1{F0: "\n\t", F1: 1000}
                                    if (((446 >= 843) && ((v3 ^ (v3 * (v17 / (((v17 & v17)) | 1)))) == int8(fi7))) || (false && false)) {
                                        v6 = v6
                                        _ = len(v4)
                                        v1 = (v1 ^ (v1 / (((-(v1))) | 1)))
                                    } else {
                                        v31 := GenS3{F0: v6, F1: 844, F2: (1e10 <= (float32(0.1) + 0.0))}
                                        var v32 int = (^(fi22))
                                        fmt.Println(v17)
                                        _ = v31
                                        _ = v32
                                    }
                                    switch (fi22 * (fi7 | (fi7 << ((0) & 63)))) {
                                    case 1:
                                        fmt.Println(v17)
                                    case 0:
                                        v3 = (^(v3))
                                    case 42:
                                        _ = len(v2)
                                    default:
                                        var v33 int16 = 32767
                                        _ = v33
                                    }
                                    _ = v30
                                }
                                v34 := GenS3{F0: v6, F1: 42, F2: (903 == 100)}
                                v35 := GenS1{F0: "foo bar", F1: 32767}
                                _ = v34
                                _ = v35
                            }
                            v3 += (v3 - v17)
                            v4 = append(v4, "\n\t")
                            _ = fi22
                        }
                    }
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 uint8 = 128
                        switch 958 {
                        case 127:
                            v37 += 66
                        default:
                            v1 = v1
                        }
                        _ = fi36
                        _ = v37
                    }
                } else {
                    var v38 int = (fi7 / ((fi7) | 1))
                    _ = v38
                }
                _ = v17
            }
            _ = fi7
        }
        var v39 uint16 = 325
        _ = v39
    }
    v40 := GenS1{F0: v6, F1: 0}
    for fi41 := 0; fi41 < 15; fi41++ {
        var v42 bool = false
        _ = fi41
        _ = v42
    }
    if true {
        v4 = append(v4, (v6 + v6))
        if true {
            v43 := GenS4{F0: true}
            _ = v43
        }
    } else {
        if (2147483647 == -87) {
            v2[32768] = false
            var v44 uint16 = 65535
            v44 -= (v44 + (v44 % ((v44) | 1)))
            _ = v44
        }
    }
    var v45 int32 = 42
    fmt.Println(v3)
    for fi46 := 0; fi46 < 16; fi46++ {
        v45 += (v45 - (v1 / ((v45) | 1)))
        v47 := GenS5{F0: -6}
        v6 = (v6 + v6)
        _ = fi46
        _ = v47
    }
    fmt.Println(v45, v6, v45)
    var v48 float64 = 0.0
    v49 := GenS3{F0: v6, F1: 229, F2: (!(18446744073709551615 <= 371))}
    v50 := GenS4{F0: (18446744073709551615 <= 18446744073709551615)}
    v51 := GenS5{F0: 0}
    if false {
        v6 = "\n\t"
    }
    var v52 uint16 = 0
    for fi53 := 0; fi53 < 16; fi53++ {
        var v54 int8 = -10
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
