package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint8 }

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 0
    var v3 int32 = 2147483647
    var v4 uint = 65535
    var v5 int8 = -1
    v6 := GenS1{F0: 42, F1: v3, F2: 624}
    if true {
        if (!((v2 & (v2 / (((v2 * (v2 << ((v4) & 63)))) | 1))) < v2)) {
            var v7 int = -784
            _ = v7
        }
    } else {
        var v8 uint8 = 2
        _ = v8
    }
    fmt.Println(v2, v3, v3)
    var v9 uint64 = 974
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v4, v1)
            var v11 uint8 = 255
            var v12 int16 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 255, F1: int32(v9), F2: 145}
        _ = v13
    } else {
        fmt.Println(v4, v1)
    }
    fmt.Println(v4, v2, v3)
    if false {
        v14 := []int{int(v4), -284, -894, -1}
        v4 = (v4 / ((18446744073709551615) | 1))
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: 0, F1: v3, F2: v9}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v17 := GenS1{F0: 0, F1: 2147483647, F2: v9}
    for fi18 := 0; fi18 < 7; fi18++ {
        v9 = 2
        _ = fi18
    }
    v19 := []int{254, 127, 42}
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS1{F0: 212, F1: (v3 << ((v4) & 63)), F2: 4294967295}
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 720
    for fi23 := 0; fi23 < 1; fi23++ {
        v9 -= 0
        _ = len(v19)
        _ = fi23
    }
    v19 = append(v19, 1)
    v5 -= (^(v5))
    if (v2 <= (v2 / ((v2) | 1))) {
        if (false && true) {
            if true {
                v24 := GenS1{F0: 0, F1: (v3 + (v3 / ((v3) | 1))), F2: 0}
                switch -1 {
                case 486:
                    fmt.Println(v2, v22)
                }
                v22 = v22
                _ = v24
            } else {
                v4 = 0
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 float64 = 0.0
                    if true {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 uint16 = v2
                            var v29 int8 = int8(v22)
                            if false {
                                v9 = 707
                                var v30 uint8 = 1
                                var v31 int32 = 977
                                _ = v30
                                _ = v31
                            } else {
                                var v32 float64 = -2.25
                                v1 += 3.14159
                                if (42 == -866) {
                                    var v33 float32 = v1
                                    fmt.Println(v1, v28)
                                    v34 := GenS2{F0: 378, F1: 545, F2: (v5 / ((v5) | 1))}
                                    _ = v33
                                    _ = v34
                                }
                                v2 += uint16(v9)
                                _ = v32
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        for fi35 := 0; fi35 < 5; fi35++ {
                            v1 += -2.25
                            if false {
                                var v36 int16 = 1
                                var v37 string = ""
                                v2 -= (v2 >> ((v4) & 63))
                                _ = v36
                                _ = v37
                            }
                            v26 -= (0.1 + v26)
                            _ = fi35
                        }
                        if (((^(fi25)) == (-(fi25))) || true) {
                            fmt.Println(v2, v1, v1)
                        }
                    } else {
                        v1 = 0.1
                        for fi38 := 0; fi38 < 11; fi38++ {
                            switch "\n\t" {
                            case "a":
                                v2 = (v2 - 214)
                            case "":
                                v22 += 65535
                            case "hello":
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            }
                            var v39 uint16 = (v2 & (v2 + v2))
                            v19 = append(v19, fi38)
                            _ = fi38
                            _ = v39
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                var v40 int8 = v5
                _ = v40
            }
            for fi41 := 0; fi41 < 12; fi41++ {
                v2 = 2
                v42 := GenS2{F0: 0, F1: -929, F2: v5}
                if false {
                    v19 = append(v19, (fi41 << (((v4 >> ((v4) & 63))) & 63)))
                } else {
                    var v43 uint64 = (^(v9))
                    v44 := GenS3{F0: -4, F1: 185, F2: 243}
                    v4 -= 42
                    _ = v43
                    _ = v44
                }
                _ = fi41
                _ = v42
            }
        }
        var v45 uint = v4
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
