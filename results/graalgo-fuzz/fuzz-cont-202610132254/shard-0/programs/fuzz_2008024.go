package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 bool; F2 int64 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    var v4 float64 = 1.5
    var v5 uint64 = 74
    var v6 int64 = -787
    v4 = 3.14159
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v3, v4)
        _ = fi7
    }
    if (false && true) {
        v6 = int64(v1)
        v3 += v3
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 18446744073709551615}
            v10 := v2[v3]
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v11 uint32 = 42
        v5 = (v5 >> ((255) & 63))
        _ = v11
    }
    if true {
        var v12 int = -2147483648
        v13 := GenS2{F0: 65535, F1: 4}
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS3{F0: 127, F1: 607, F2: (v6 >> ((42) & 63))}
            if true {
                var v16 int32 = -350
                for fi17 := 0; fi17 < 10; fi17++ {
                    if ((((3.14159 / 1.0) < (1e10 - 1.5)) && (1 <= 42)) || (42 > 639)) {
                        v18 := []int{(fi14 + (fi17 / ((fi14) | 1))), (fi14 % ((2147483647) | 1))}
                        switch (v12 - (v12 - -809)) {
                        case -28:
                            var v19 int16 = 1000
                            _ = v19
                        }
                        _ = v18
                    } else {
                        v20 := GenS2{F0: 1, F1: 255}
                        v6 = 9223372036854775807
                        _ = v20
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                v21 := GenS4{F0: (v1 | v1)}
                for fi22 := 0; fi22 < 10; fi22++ {
                    v1 = v1
                    v23 := GenS4{F0: -128}
                    var v24 uint64 = v5
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            }
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    } else {
        var v25 string = "\n\t"
        var v26 uint16 = 2
        _ = v25
        _ = v26
    }
    var v27 uint64 = v5
    var v28 int = 127
    if false {
        var v29 uint = 212
        v30 := GenS5{F0: 1000, F1: -1}
        _ = v29
        _ = v30
    } else {
        var v31 uint16 = 0
        _ = v31
    }
    v5 += (^(v5))
    fmt.Println(v3, v4, v3)
    if ((v3 << ((18446744073709551615) & 63)) > 4294967295) {
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 int = fi32
            fmt.Println(v3, v4, v28)
            var v34 int = fi32
            _ = fi32
            _ = v33
            _ = v34
        }
    } else {
        v6 += int64(v5)
        v35 := []int{v28, v28}
        v6 = 42
        _ = v35
    }
    if false {
        if false {
            var v36 uint = 2
            fmt.Println(v28, v1)
            v37 := GenS5{F0: 1000, F1: -32768}
            _ = v36
            _ = v37
        }
        var v38 uint = 315
        v39 := GenS1{F0: 2}
        _ = v38
        _ = v39
    } else {
        var v40 uint16 = 0
        _ = v40
    }
    switch (v28 | 1) {
    case -9223372036854775808:
        if ((v3 ^ (v3 << ((65535) & 63))) <= (v3 * v3)) {
            for fi41 := 0; fi41 < 10; fi41++ {
                switch "\n\t" {
                case "a":
                    if (2 == 0) {
                        var v42 int8 = (-(v1))
                        _ = v42
                    } else {
                        v3 = v3
                    }
                case "hello":
                    v43 := GenS6{F0: 0}
                    _ = v43
                default:
                    v3 -= (v3 / (((v3 ^ v3)) | 1))
                }
                if false {
                    var v44 int = (v28 / ((v28) | 1))
                    var v45 float32 = 0.1
                    if (v28 == v44) {
                        if (!(!(-2147483648 < -1))) {
                            fmt.Println(v44)
                            v46 := GenS7{F0: (0 < 42), F1: (!true), F2: -964}
                            v2[(v3 % ((42) | 1))] = (!(false && ((1 <= -32768) && true)))
                            _ = v46
                        } else {
                            v47 := GenS5{F0: -42, F1: -1}
                            var v48 int16 = 32767
                            v49 := v2[v3]
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                    } else {
                        v50 := GenS1{F0: 753}
                        _ = v50
                    }
                    _ = v44
                    _ = v45
                } else {
                    v51 := GenS4{F0: v1}
                    _ = v51
                }
                v4 = (((float64(3.14159) * v4) / (v4 * (float64(1.5) + v4))) + float64(float32((1e10 + 0.0))))
                _ = fi41
            }
        } else {
            var v52 uint16 = 1
            v2[65535] = true
            _ = v52
        }
    case 9223372036854775807:
        v3 = 25
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
