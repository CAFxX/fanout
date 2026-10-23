package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint; F1 uint16; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }

func main() {
    var v1 int16 = -658
    var v2 uint8 = 13
    var v3 int32 = -2147483648
    var v4 int16 = v1
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            if (!(0.0 == 1.5)) {
                var v7 float64 = (-2.25 / -2.25)
                fmt.Println(fi5, v3)
                _ = v7
            } else {
                v1 += 0
            }
            _ = fi6
        }
        var v8 int32 = int32(v1)
        _ = fi5
        _ = v8
    }
    if false {
        switch "hello" {
        case "\n\t":
            fmt.Println(v4, v2)
        default:
            v9 := GenS1{F0: -1, F1: "\n\t", F2: 2}
            _ = v9
        }
        var v10 float64 = 3.14159
        _ = v10
    }
    var v11 uint32 = 249
    if true {
        var v12 uint8 = v2
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                fmt.Println(fi13, v11, v12)
                _ = fi14
            }
            _ = fi13
        }
        _ = v12
    }
    var v15 int32 = 458
    if true {
        var v16 uint32 = (v11 * 0)
        fmt.Println(v4)
        v17 := []int{32767, 119, 1, 9223372036854775807}
        _ = v16
        _ = v17
    } else {
        v18 := []int{4294967295}
        _ = v18
    }
    fmt.Println(v11, v3)
    v19 := GenS2{F0: v2}
    v20 := GenS2{F0: v2}
    v21 := GenS1{F0: 1, F1: "\n\t", F2: 103}
    switch 128 {
    case 2147483647:
        v11 += 382
    case 9223372036854775807:
        v22 := GenS3{F0: 9223372036854775807, F1: 907, F2: -1}
        _ = v22
    default:
        if (v11 >= 2147483648) {
            v11 = (v11 << ((65535) & 63))
            fmt.Println(v3)
            v4 = v1
        }
    }
    var v23 uint32 = v11
    if false {
        v24 := []int{2147483647, 32768, 127}
        _ = v24
    }
    v11 = (v11 * (v23 >> ((42) & 63)))
    fmt.Println(v11, v1)
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 14; fi26++ {
            fmt.Println(fi26, fi25)
            _ = fi26
        }
        var v27 uint32 = v23
        if (!true) {
            if true {
                v27 = 0
                var v28 float32 = (3.14159 + 1e10)
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 int = fi29
                    v4 = -32768
                    _ = fi29
                    _ = v30
                }
                _ = v28
            } else {
                fmt.Println(v2, v4, v23)
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v23, v3, v27)
                    _ = fi31
                }
                if false {
                    switch (fi25 >> ((18446744073709551615) & 63)) {
                    case 685:
                        var v32 uint16 = 1
                        _ = v32
                    case 32768:
                        v33 := GenS4{F0: 294, F1: (v2 + v2), F2: 0}
                        _ = v33
                    case 9223372036854775807:
                        var v34 int16 = v4
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 3; fi35++ {
                        v27 = 552
                        switch fi35 {
                        case 42:
                            v36 := []int{32767, 540}
                            _ = v36
                        default:
                            var v37 float32 = 1.5
                            _ = v37
                        }
                        if (!(629 <= 65535)) {
                            var v38 uint32 = v11
                            _ = v38
                        } else {
                            v23 = v11
                            if (!false) {
                                v39 := GenS5{F0: 467, F1: 0, F2: (v1 / (((v1 % ((32767) | 1))) | 1))}
                                fmt.Println(v1, v2, v23)
                                v40 := []int{fi25, (fi35 / ((fi25) | 1)), 0, 78}
                                _ = v39
                                _ = v40
                            } else {
                                var v41 string = ""
                                fmt.Println(v41, v23, v3)
                                v42 := GenS6{F0: (v4 % ((32767) | 1))}
                                _ = v41
                                _ = v42
                            }
                        }
                        _ = fi35
                    }
                } else {
                    if true {
                        v43 := GenS7{F0: "a"}
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 15; fi44++ {
                            v15 = v15
                            v45 := []int{fi25, (fi25 ^ fi44), 32767, fi25}
                            var v46 float32 = (-2.25 / -2.25)
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    }
                }
            }
            if (678 >= 532) {
                fmt.Println(v15, v1)
                var v47 uint16 = 65535
                _ = v47
            }
            v15 -= v15
        }
        _ = fi25
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
