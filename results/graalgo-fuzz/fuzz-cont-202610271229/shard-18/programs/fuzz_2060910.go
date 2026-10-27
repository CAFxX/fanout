package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int; F1 uint64; F2 uint32 }

func main() {
    v1 := []uint32{4294967295}
    var v2 int32 = -2147483648
    var v3 uint32 = 2147483648
    var v4 int16 = -301
    switch -567 {
    case 42:
        v2 -= (v2 >> ((18446744073709551615) & 63))
    }
    if false {
        v5 := GenS1{F0: (^(v2)), F1: false}
        v4 = (v4 * int16(v3))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v5
    } else {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch (^(fi6)) {
            case -430:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v7 int = (fi6 / ((fi6) | 1))
                _ = v7
            }
            v2 += 1
            _ = fi6
        }
        v4 = (v4 * (^(v4)))
        if ((true || false) || (1 < 127)) {
            v3 -= 0
            v8 := GenS1{F0: v2, F1: false}
            for fi9 := 0; fi9 < 8; fi9++ {
                _ = len(v1)
                fmt.Println(fi9)
                _ = fi9
            }
            _ = v8
        } else {
            v2 = -1
            v10 := GenS2{F0: 546}
            v4 += 1000
            _ = v10
        }
    }
    fmt.Println(v4)
    if true {
        v3 = (v3 >> ((1) & 63))
        v11 := GenS3{F0: 255, F1: (v2 / ((42) | 1))}
        _ = v11
    }
    v12 := GenS4{F0: 127, F1: 65535}
    v13 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: (1.5 > ((0.0 / ((1.5 + (((0.1 * -2.25) - (3.14159 * 1e10)) - ((float64(-2.25) / 0.1) * (float64((3.14159 * 0.1)) - -2.25)))) - 1.5)) + (1.5 + (0.0 / -2.25))))}
    switch 195 {
    case 32767:
        var v14 uint8 = 255
        _ = v14
    case 42:
        fmt.Println(v4, v3)
    }
    var v15 uint64 = 4294967295
    var v16 uint8 = 255
    var v17 float64 = (((0.1 * ((0.0 - 0.1) / 1e10)) + 1e10) * 0.1)
    if true {
        switch 0 {
        case 771:
            for fi18 := 0; fi18 < 11; fi18++ {
                v17 -= (-2.25 + 0.1)
                v1 = append(v1, (v3 + (v3 ^ (v3 + 2147483648))))
                _ = fi18
            }
        case 82:
            _ = len(v1)
        default:
            switch 32767 {
            case -742:
                if true {
                    var v19 uint = 65535
                    switch 32767 {
                    case 1:
                        _ = len(v1)
                    case 2147483647:
                        switch (("hello" + "a") + "\n\t") {
                        case "hello":
                            v20 := GenS4{F0: v16, F1: v3}
                            _ = v20
                        default:
                            if (false && (v3 > (^(v3)))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v16)
                            }
                        }
                    }
                    var v21 float32 = 0.1
                    _ = v19
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        if (v3 >= 42) {
                            v2 = (v2 ^ (v2 & v2))
                        }
                        _ = fi22
                    }
                    fmt.Println(v16)
                }
            case -1:
                var v23 float32 = 1e10
                _ = v23
            case 1:
                switch 0 {
                case -56:
                    fmt.Println(v4, v15, v2)
                case -654:
                    if true {
                        var v24 int64 = 42
                        _ = v24
                    }
                case 32768:
                    if true {
                        if (757 < 838) {
                            v4 += 32767
                        }
                    } else {
                        var v25 uint64 = (v15 * uint64(v16))
                        v15 -= v15
                        _ = len(v1)
                        _ = v25
                    }
                }
            default:
                var v26 float64 = v17
                _ = v26
            }
        }
        if len(v1) > 0 {
            v1[0] = v3
        }
    } else {
        v17 = (float64(3.14159) + (1e10 / 1.0))
        var v27 float32 = (3.14159 + ((float32((float64(0.0) - (v17 + v17))) + (0.1 * float32(3.14159))) + (0.0 + ((-2.25 * 1.5) - ((float32(float64(1.5)) + 3.14159) / -2.25)))))
        v4 += v4
        _ = v27
    }
    switch 127 {
    case 2147483647:
        switch 128 {
        case -745:
            switch (("hello" + "a") + "foo bar") {
            case "":
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(v15, v4, v16)
                    if false {
                        v1 = append(v1, 2147483648)
                    } else {
                        switch "a" {
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v1 = append(v1, v3)
                    }
                    v16 = 127
                    _ = fi28
                }
            case "hello":
                v29 := GenS5{F0: -230, F1: (v15 >> ((0) & 63))}
                _ = v29
            }
        }
    case 9223372036854775807:
        if ("" >= "foo bar") {
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 int = fi30
                switch "hello" {
                case "":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 uint64 = 0
                        _ = fi32
                        _ = v33
                    }
                case "foo bar":
                    v16 -= 127
                }
                _ = fi30
                _ = v31
            }
        } else {
            v16 = uint8(v3)
            if (v17 > v17) {
                if ((-128 != 1) && (true || ((!true) && ((1e10 - -2.25) <= 1.5)))) {
                    if true {
                        var v34 uint64 = 0
                        var v35 float64 = (v17 * 1.5)
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int = 2147483647
                        var v37 int = v36
                        _ = v36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 16; fi38++ {
                        switch fi38 {
                        case 0:
                            v39 := GenS6{F0: (fi38 / (((fi38 & fi38)) | 1)), F1: (v15 & v15), F2: (v3 >> ((0) & 63))}
                            _ = v39
                        case 863:
                            if (true || true) {
                                v3 = 65535
                            }
                        default:
                            if false {
                                v1 = append(v1, 65535)
                            }
                        }
                        _ = fi38
                    }
                } else {
                    fmt.Println(v3)
                    v16 = 128
                }
                v40 := GenS3{F0: 2, F1: (v2 >> ((255) & 63))}
                v15 = v15
                _ = v40
            } else {
                v2 = v2
                var v41 float32 = ((1.5 / 0.1) / 1.5)
                var v42 int8 = 0
                _ = v41
                _ = v42
            }
        }
    case 127:
        for fi43 := 0; fi43 < 18; fi43++ {
            v1 = append(v1, (v3 % ((v3) | 1)))
            _ = fi43
        }
    }
    v44 := GenS1{F0: v2, F1: false}
    v2 += 2147483647
    fmt.Println(v16, v17, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
