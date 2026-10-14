package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 uint16 = 218
    var v2 int64 = -494
    v3 := make(map[uint32]bool)
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint32 = 958
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v6)
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    var v9 int32 = -2147483648
    v1 += (v1 >> ((65535) & 63))
    var v10 int32 = (^(v9))
    var v11 uint32 = 42
    v4 = 4294967295
    v10 -= (v10 & v10)
    v12 := v3[v11]
    var v13 string = ("\n\t" + "hello")
    if ((float32((0.0 / 3.14159)) <= 1.5) && (209 >= 0)) {
        if v12 {
            v1 = v1
        } else {
            if v12 {
                var v14 float32 = (3.14159 - (0.0 * 3.14159))
                var v15 uint16 = (v1 + (v1 & uint16(v4)))
                v11 = (v11 << ((764) & 63))
                _ = v14
                _ = v15
            }
        }
        v12 = v12
    }
    v16 := GenS1{F0: v11, F1: (v2 & v2), F2: v1}
    var v17 string = v13
    v18 := GenS2{F0: v12, F1: (v10 ^ 0)}
    v1 = (v1 >> ((18446744073709551615) & 63))
    if (int8(v9) <= 1) {
        v10 -= v9
    } else {
        switch v17 {
        case "hello":
            if (((float64(3.14159) / 1.0) - 1.5) <= 3.14159) {
                v19 := GenS2{F0: v12, F1: (v10 ^ v10)}
                v20 := []int{9223372036854775807, 949, -1, -301}
                _ = v19
                _ = v20
            }
        case "a":
            var v21 uint32 = 65535
            _ = v21
        case "foo bar":
            if v12 {
                fmt.Println(v9, v11, v1)
            } else {
                fmt.Println(v12, v13, v13)
                if false {
                    switch -1 {
                    case 167:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            for fi23 := 0; fi23 < 10; fi23++ {
                                v13 = v17
                                _ = fi23
                            }
                            for fi24 := 0; fi24 < 3; fi24++ {
                                var v25 uint = 255
                                _ = fi24
                                _ = v25
                            }
                            v4 = v4
                            _ = fi22
                        }
                    }
                    v3[v11] = (1.5 == 0.0)
                    delete(v3, (v11 ^ v11))
                }
            }
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            v12 = (255 != 417)
            if (2147483648 == v11) {
                delete(v3, (v11 * v11))
                if (v12 || (v10 > 1)) {
                    if v12 {
                        switch v13 {
                        case "hello":
                            v17 = (((v17 + v13) + v17) + (((v13 + "a") + v13) + "a"))
                        }
                        v27 := []int{(fi26 ^ (fi26 - 184)), int(v4)}
                        if (v1 != (v1 * (^(v1)))) {
                            switch (fi26 - int(v9)) {
                            case -9223372036854775808:
                                v13 = v17
                            case 760:
                                var v28 uint8 = 210
                                _ = v28
                            }
                        }
                        _ = v27
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if (((float32(3.14159) < (0.1 * 3.14159)) || (0.0 == 3.14159)) || (1e10 >= ((0.0 * 1.5) + float64((3.14159 + (float32(((((3.14159 + 1e10) / 0.1) + 0.0) * 0.0)) + 3.14159)))))) {
                                var v30 string = v13
                                if true {
                                    if (!v12) {
                                        switch fi26 {
                                        case -545:
                                            v11 = (v11 + v11)
                                        }
                                        v12 = ((v12 && (v12 && (0.1 >= 3.14159))) && (-1 <= 1))
                                        if (1 >= 127) {
                                            v31 := GenS3{F0: -286}
                                            v11 = (v11 & v11)
                                            switch int(v4) {
                                            case -1:
                                                v11 -= v11
                                            case 4294967295:
                                                if ((v11 <= 402) || ((42 <= 42) || (309 >= 32767))) {
                                                    v4 -= 0
                                                } else {
                                                    delete(v3, v11)
                                                    var v32 uint8 = 1
                                                    v2 = (v2 - v2)
                                                    _ = v32
                                                }
                                            case 32768:
                                                v12 = v12
                                            }
                                            _ = v31
                                        }
                                        v12 = true
                                    } else {
                                        v33 := GenS3{F0: -714}
                                        _ = v33
                                    }
                                    var v34 int = (^(fi26))
                                    v9 += v9
                                    _ = v34
                                } else {
                                    switch v13 {
                                    case "":
                                        var v35 uint32 = v11
                                        _ = v35
                                    case "foo bar":
                                        fmt.Println(fi26, fi26, v17)
                                    case "hello":
                                        switch v17 {
                                        case "":
                                            if (42 >= -72) {
                                                fmt.Println(v9, v2, v17)
                                            }
                                        case "foo bar":
                                            var v36 float64 = float64((0.1 - float32(1e10)))
                                            _ = v36
                                        case "hello":
                                            v30 = (v13 + v17)
                                        }
                                    }
                                    v30 = (v13 + v30)
                                }
                                if (v10 > v9) {
                                    fmt.Println(v4, v9, v9)
                                    fmt.Println(v13, v12)
                                } else {
                                    _ = len(v3)
                                    var v37 uint8 = uint8(v4)
                                    _ = v37
                                }
                                _ = v30
                            }
                            _ = fi29
                        }
                        v2 -= (-(v2))
                        var v38 int64 = v2
                        _ = v38
                    }
                    v13 = (v17 + ((v17 + "\n\t") + v13))
                }
                if v12 {
                    if false {
                        for fi39 := 0; fi39 < 2; fi39++ {
                            var v40 int16 = -1
                            switch (-(fi26)) {
                            case -9223372036854775808:
                                var v41 int = -2147483648
                                _ = v41
                            case -629:
                                v42 := v3[(v11 ^ 4294967295)]
                                _ = v42
                            }
                            var v43 int32 = (v9 / ((42) | 1))
                            _ = fi39
                            _ = v40
                            _ = v43
                        }
                        v44 := []int{128, -2147483648, fi26, (fi26 & int(v1))}
                        _ = v44
                    }
                    if v12 {
                        for fi45 := 0; fi45 < 3; fi45++ {
                            var v46 int16 = 0
                            delete(v3, v11)
                            var v47 string = v13
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    } else {
                        fmt.Println(fi26, v1)
                    }
                }
            }
            _ = fi26
        }
        v48 := GenS4{F0: 0, F1: 231}
        _ = v48
    }
    if v12 {
        var v49 uint16 = (v1 >> ((1) & 63))
        for fi50 := 0; fi50 < 2; fi50++ {
            v51 := GenS2{F0: v12, F1: v9}
            for fi52 := 0; fi52 < 6; fi52++ {
                var v53 int8 = -71
                fmt.Println(v10)
                var v54 int32 = (v9 * -1)
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = fi50
            _ = v51
        }
        _ = v49
    } else {
        if (!false) {
            v55 := GenS5{F0: 2}
            _ = v55
        } else {
            if false {
                v3[(v11 - v11)] = true
                if (v12 && true) {
                    fmt.Println(v2, v12, v12)
                } else {
                    var v56 int32 = (v9 / ((42) | 1))
                    _ = v56
                }
                v57 := []int{}
                _ = v57
            } else {
                v58 := []int{}
                v2 = (v2 << ((4294967295) & 63))
                _ = v58
            }
            v1 = (v1 % ((v1) | 1))
        }
        if true {
            v12 = false
            for fi59 := 0; fi59 < 13; fi59++ {
                v60 := GenS6{F0: 132}
                _ = fi59
                _ = v60
            }
        } else {
            switch "\n\t" {
            case "hello":
                fmt.Println(v13)
            default:
                v4 -= v4
            }
        }
    }
    var v61 int16 = -27
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v9
}
