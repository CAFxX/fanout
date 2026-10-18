package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 string }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16 }

func main() {
    var v1 int16 = 598
    var v2 int = -2147483648
    var v3 uint64 = 2
    v4 := []int{}
    var v5 uint16 = 65535
    v6 := []int{(v2 % ((v2) | 1)), v2}
    if false {
        var v7 uint64 = v3
        var v8 int = (-(v2))
        fmt.Println(v2, v3, v8)
        _ = v7
        _ = v8
    }
    if false {
        var v9 uint64 = (v3 | 4294967295)
        v3 -= (v9 | v3)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            v1 += (v1 - v1)
            var v11 int16 = v1
            _ = fi10
            _ = v11
        }
    }
    var v12 uint16 = (v5 ^ (v5 << ((534) & 63)))
    var v13 int32 = 2147483647
    v14 := GenS1{F0: v13, F1: 155, F2: 42}
    _ = len(v4)
    for fi15 := 0; fi15 < 5; fi15++ {
        switch (fi15 << ((899) & 63)) {
        case 32768:
            if (true && ((1.5 + (0.0 + (0.1 - 0.0))) > 3.14159)) {
                if (0 == 42) {
                    var v16 uint8 = 109
                    if true {
                        fmt.Println(v1, fi15)
                    }
                    v6 = append(v6, 0)
                    _ = v16
                } else {
                    var v17 uint64 = (v3 - (v3 % ((v3) | 1)))
                    switch "" {
                    case "hello":
                        fmt.Println(v1, v12)
                    case "a":
                        v4 = append(v4, -633)
                    case "foo bar":
                        if (9223372036854775807 <= 1) {
                            if (float32((-2.25 + 1e10)) == (3.14159 - 1e10)) {
                                v18 := GenS2{F0: 4294967295, F1: -128, F2: 180}
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    var v20 bool = false
                                    var v21 uint64 = v3
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v13 -= 699
                                    if v20 {
                                        if (-2.25 > ((float32(0.1) / ((float32(0.0) * 0.1) - 1.5)) - ((-2.25 * 0.1) + 0.0))) {
                                            v6 = append(v6, (fi15 << ((402) & 63)))
                                            v3 = 0
                                            if (32767 == v12) {
                                                var v22 uint = 567
                                                var v23 int = fi15
                                                var v24 int64 = 9223372036854775807
                                                if len(v6) > 0 {
                                                    v6[0] = (v2 >> (((v22 | (v22 ^ (v22 % ((v22) | 1))))) & 63))
                                                }
                                                v25 := GenS3{F0: 1, F1: -7, F2: "hello"}
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v5 = (v12 & (v5 / ((65535) | 1)))
                                            }
                                        } else {
                                            var v26 float32 = 1.5
                                            switch ("\n\t" + "\n\t") {
                                            case "foo bar":
                                                switch "a" {
                                                case "":
                                                    v3 = 18446744073709551615
                                                default:
                                                    if v20 {
                                                        v20 = (2 >= v3)
                                                        v2 = 0
                                                    }
                                                }
                                            case "a":
                                                var v27 uint32 = 231
                                                _ = v27
                                            case "\n\t":
                                                v13 += v13
                                                v28 := GenS1{F0: v13, F1: v5, F2: 83}
                                                if (541 > 18446744073709551615) {
                                                    if (316 <= 0) {
                                                        fmt.Println(fi15, v2, v5)
                                                    }
                                                    v29 := GenS3{F0: (^(v17)), F1: 0, F2: "foo bar"}
                                                    var v30 int16 = (-(v1))
                                                    _ = v29
                                                    _ = v30
                                                }
                                                switch (v2 >> ((42) & 63)) {
                                                case -634:
                                                    v31 := GenS4{F0: (v5 * v5), F1: uint8(v17), F2: 9223372036854775807}
                                                    _ = v31
                                                case -9223372036854775808:
                                                    fmt.Println(v17)
                                                case -2147483648:
                                                    switch (^(v2)) {
                                                    case 32767:
                                                        if (v20 && (!(0 > 1))) {
                                                            var v32 uint16 = (v12 | v12)
                                                            fmt.Println(v12, v1)
                                                            _ = v32
                                                        } else {
                                                            v17 = (^(v17))
                                                        }
                                                    case 32768:
                                                        v33 := GenS5{F0: 65535, F1: 426, F2: 121}
                                                        _ = v33
                                                    }
                                                default:
                                                    v34 := GenS1{F0: (v13 + (v13 / ((v13) | 1))), F1: (v5 << ((255) & 63)), F2: 0}
                                                    _ = v34
                                                }
                                                _ = v28
                                            }
                                            _ = v26
                                        }
                                    } else {
                                        if v20 {
                                            v4 = append(v4, fi19)
                                            fmt.Println(v21, v20)
                                        } else {
                                            fmt.Println(v13)
                                            _ = len(v6)
                                            v35 := GenS2{F0: 531, F1: 42, F2: 65535}
                                            _ = v35
                                        }
                                        var v36 int32 = (v13 / ((v13) | 1))
                                        if true {
                                            var v37 uint32 = 335
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                v38 := GenS6{F0: v17, F1: 255, F2: (^(v17))}
                                                _ = v38
                                            }
                                            _ = v37
                                        } else {
                                            v39 := GenS4{F0: (v12 - (v12 - v12)), F1: 255, F2: 42}
                                            _ = v39
                                        }
                                        if v20 {
                                            var v40 uint32 = 2147483647
                                            _ = v40
                                        }
                                        var v41 uint32 = 629
                                        _ = v36
                                        _ = v41
                                    }
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                v42 := GenS7{F0: v12}
                                _ = v42
                            }
                        }
                    }
                    for fi43 := 0; fi43 < 8; fi43++ {
                        var v44 uint = 42
                        _ = fi43
                        _ = v44
                    }
                    _ = v17
                }
                v45 := GenS6{F0: v3, F1: 2, F2: v3}
                _ = v45
            } else {
                switch v2 {
                case 2147483647:
                    v13 += v13
                case -9223372036854775808:
                    var v46 int64 = 0
                    _ = v46
                case 127:
                    var v47 float32 = 0.1
                    _ = v47
                default:
                    v5 -= v12
                }
                v5 = (^(v5))
            }
        default:
            v48 := GenS8{F0: 50}
            _ = v48
        }
        var v49 int64 = -9223372036854775808
        switch (fi15 ^ (fi15 - (fi15 * v2))) {
        case 9223372036854775807:
            v50 := GenS7{F0: v5}
            _ = v50
        case -2147483648:
            var v51 float64 = 1.5
            _ = v51
        case 858:
            if (18446744073709551615 > 1) {
                v6 = append(v6, v2)
            } else {
                v52 := GenS9{F0: (-(v1))}
                var v53 int8 = 42
                var v54 uint32 = 4294967295
                _ = v52
                _ = v53
                _ = v54
            }
        }
        _ = fi15
        _ = v49
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        var v56 float64 = (1.5 * ((0.1 / 0.1) / 1.5))
        for fi57 := 0; fi57 < 19; fi57++ {
            v13 -= v13
            if true {
                var v58 int64 = 42
                v13 -= 0
                _ = v58
            }
            var v59 int = 128
            if false {
                v59 += int(v13)
            }
            var v60 uint64 = (v3 >> ((688) & 63))
            _ = fi57
            _ = v59
            _ = v60
        }
        var v61 uint64 = v3
        _ = fi55
        _ = v56
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
