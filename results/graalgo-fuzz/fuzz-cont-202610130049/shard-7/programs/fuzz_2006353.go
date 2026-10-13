package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    v1 := []int{696, 2147483647, -402, 9223372036854775807}
    var v2 uint = 153
    var v3 int = -2147483648
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: -1}
    v3 = 42
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            v3 += 32768
        }
        _ = fi7
    }
    if (1 < 167) {
        if false {
            v8 := GenS2{F0: 114, F1: (!(true && true)), F2: -899}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 10; fi9++ {
                delete(v5, "foo bar")
                _ = fi9
            }
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        _ = len(v5)
        fmt.Println(v3, v3, v3)
        var v11 uint32 = 0
        _ = fi10
        _ = v11
    }
    fmt.Println(v2)
    if ((false || false) && false) {
        switch (^(v3)) {
        case 127:
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v3, v2)
            case "a":
                for fi12 := 0; fi12 < 3; fi12++ {
                    fmt.Println(v2, v3)
                    if true {
                        if true {
                            var v13 uint8 = 0
                            for fi14 := 0; fi14 < 3; fi14++ {
                                var v15 float32 = -2.25
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 16; fi16++ {
                                var v17 int64 = 9223372036854775807
                                if true {
                                    v18 := GenS1{F0: v17}
                                    var v19 uint8 = v13
                                    _ = v18
                                    _ = v19
                                }
                                fmt.Println(v2)
                                _ = fi16
                                _ = v17
                            }
                            _ = v13
                        } else {
                            if true {
                                var v20 float32 = -2.25
                                _ = v20
                            } else {
                                fmt.Println(v2, v2, v2)
                            }
                            v2 = 379
                            var v21 int = 2147483647
                            _ = v21
                        }
                        switch "hello" {
                        case "\n\t":
                            v3 = (v3 | 32767)
                        case "hello":
                            v3 += (v3 | v3)
                        case "":
                            switch (v3 ^ v3) {
                            case 128:
                                v3 = fi12
                            case 32768:
                                v1 = append(v1, int(v2))
                            case -2147483648:
                                v3 = fi12
                            }
                        default:
                            v2 -= v2
                        }
                    } else {
                        fmt.Println(v3, fi12, v3)
                    }
                    if ((((1.5 / 1.5) >= (3.14159 / (0.1 / -2.25))) && false) || true) {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            if (127 != 42) {
                                v3 = (fi22 * v3)
                                fmt.Println(v2, v3, fi12)
                                v3 += -167
                            }
                            var v23 float32 = float32((1.5 - (0.0 - 3.14159)))
                            v2 = v2
                            fmt.Println(v2)
                            v3 += 32767
                            _ = fi22
                            _ = v23
                        }
                        var v24 float64 = 1.5
                        if (716 == (v2 % (((^(v2))) | 1))) {
                            for fi25 := 0; fi25 < 7; fi25++ {
                                fmt.Println(fi12, fi12, fi25)
                                _ = fi25
                            }
                            var v26 int64 = 700
                            _ = v26
                        }
                        _ = v24
                    }
                    _ = fi12
                }
            }
        case 1:
            v27 := GenS2{F0: 0, F1: (42 <= 255), F2: -2147483648}
            _ = v27
        case 2147483647:
            v3 += 32767
        default:
            v2 -= (^(v2))
        }
    } else {
        v2 -= (v2 | (v2 & (v2 + v2)))
        v28 := GenS1{F0: int64(v2)}
        _ = v28
    }
    v1 = append(v1, (^(v3)))
    fmt.Println(v3, v3)
    if false {
        v2 = 1
        var v29 bool = (false && true)
        var v30 int = 571
        _ = v29
        _ = v30
    } else {
        if false {
            v2 -= v2
            v31 := GenS3{F0: -1, F1: 128, F2: v2}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 18; fi32++ {
                v2 = (^(v2))
                fmt.Println(fi32, fi32)
                if false {
                    v2 = 1
                }
                _ = fi32
            }
            var v33 int8 = 16
            _ = v33
        }
        if (false && ((!false) || (18446744073709551615 == 42))) {
            if (!((1 == 874) && false)) {
                for fi34 := 0; fi34 < 6; fi34++ {
                    fmt.Println(v3, v2)
                    _ = fi34
                }
            }
            v2 = v2
        } else {
            var v35 int8 = 100
            if (0.0 < (3.14159 / 0.1)) {
                fmt.Println(v3, v35)
                var v36 uint64 = 152
                _ = v36
            } else {
                delete(v5, "hello")
                v37 := v5[""]
                for fi38 := 0; fi38 < 3; fi38++ {
                    fmt.Println(fi38, v35)
                    v35 -= v35
                    v39 := GenS4{F0: 1000}
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            v3 = -1
            _ = v35
        }
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS5{F0: (v2 | v2), F1: true}
        var v42 bool = (-2 < -128)
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int = (v3 / (((v3 >> ((4294967295) & 63))) | 1))
    for fi44 := 0; fi44 < 4; fi44++ {
        switch "hello" {
        case "a":
            switch fi44 {
            case 32768:
                v2 -= (^(v2))
            case 0:
                v2 = 1
            }
        }
        _ = fi44
    }
    _ = len(v4)
    var v45 bool = true
    for fi46 := 0; fi46 < 8; fi46++ {
        _ = len(v4)
        switch ("" + "\n\t") {
        case "hello":
            v47 := GenS5{F0: (v2 * 255), F1: v45}
            _ = v47
        case "foo bar":
            v48 := GenS3{F0: -1, F1: 133, F2: 1}
            _ = v48
        case "\n\t":
            var v49 bool = (float64(1.5) < float64((1.5 / 1.0)))
            _ = v49
        }
        _ = fi46
    }
    var v50 int = v3
    for fi51 := 0; fi51 < 8; fi51++ {
        var v52 float32 = 0.0
        var v53 uint32 = 65535
        if v45 {
            v2 += 1
            v54 := GenS4{F0: 32767}
            _ = v54
        } else {
            v55 := GenS4{F0: -870}
            for fi56 := 0; fi56 < 2; fi56++ {
                var v57 uint = 1
                if (!v45) {
                    delete(v4, v53)
                } else {
                    v53 -= (v53 << ((v2) & 63))
                }
                if len(v1) > 0 {
                    v1[0] = fi51
                }
                _ = fi56
                _ = v57
            }
            for fi58 := 0; fi58 < 11; fi58++ {
                v50 += (v50 + 9223372036854775807)
                v3 = 4294967295
                var v59 float32 = v52
                _ = fi58
                _ = v59
            }
            _ = v55
        }
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v50
    _ = v6
}
