package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    v2 := []int{725, 0}
    var v3 float32 = float32(0.1)
    var v4 float64 = -2.25
    var v5 int64 = 42
    var v6 bool = (v4 > 3.14159)
    fmt.Println(v4, v6)
    var v7 bool = v6
    if true {
        v4 += 3.14159
        if len(v2) > 0 {
            v2[0] = -478
        }
        var v8 float32 = 0.1
        _ = v8
    }
    if v7 {
        v5 = (v5 >> ((633) & 63))
        switch "\n\t" {
        case "a":
            switch -441 {
            case -2147483648:
                v9 := GenS1{F0: (v4 < 0.0)}
                _ = v9
            case -502:
                var v10 float64 = 3.14159
                _ = v10
            case -9223372036854775808:
                if v7 {
                    if v7 {
                        v3 = 1.5
                    } else {
                        var v11 float64 = 1.5
                        var v12 int = -675
                        fmt.Println(v6)
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS2{F0: -2147483648, F1: 811, F2: 4294967295}
                    _ = v13
                }
            }
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            switch "\n\t" {
            case "\n\t":
                v14 := GenS3{F0: 0, F1: 127}
                _ = v14
            case "foo bar":
                var v15 string = ("foo bar" + ((("a" + "\n\t") + "foo bar") + "foo bar"))
                _ = v15
            case "":
                for fi16 := 0; fi16 < 9; fi16++ {
                    v17 := GenS1{F0: ((2 <= 255) && (!("hello" == "")))}
                    if v6 {
                        v3 = 1e10
                        switch (-(fi16)) {
                        case 878:
                            switch (fi16 | 127) {
                            case 944:
                                v18 := GenS4{F0: 42, F1: 82}
                                _ = v18
                            case 128:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v19 int32 = -2147483648
                                _ = v19
                            }
                        default:
                            if (((v4 != v4) && (!v6)) || v7) {
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v7)
                                }
                            } else {
                                v7 = false
                                var v20 uint8 = 255
                                if true {
                                    v21 := GenS4{F0: 1, F1: 0}
                                    if (v6 || (18446744073709551615 == 42)) {
                                        if (1 <= 758) {
                                            fmt.Println(fi16, v3)
                                        } else {
                                            var v22 uint8 = (^(v20))
                                            _ = v22
                                        }
                                        v23 := GenS5{F0: 1, F1: int(v5), F2: (55 > 479)}
                                        _ = v23
                                    } else {
                                        _ = len(v2)
                                    }
                                    _ = v21
                                } else {
                                    var v24 int8 = -128
                                    _ = v24
                                }
                                _ = v20
                            }
                        }
                        var v25 uint32 = 120
                        _ = v25
                    }
                    _ = fi16
                    _ = v17
                }
            default:
                delete(v1, -310)
            }
        }
    }
    if (v4 < (v4 * v4)) {
        v26 := GenS6{F0: 186, F1: 42}
        v3 = v3
        _ = v26
    } else {
        var v27 float32 = (v3 + 3.14159)
        _ = v27
    }
    switch 9223372036854775807 {
    case 0:
        if (32767 < 42) {
            var v28 int16 = 0
            if len(v2) > 0 {
                v2[0] = 127
            }
            v3 += v3
            _ = v28
        }
    case 80:
        if true {
            for fi29 := 0; fi29 < 7; fi29++ {
                v7 = v7
                v3 -= v3
                _ = fi29
            }
            if (!(!v7)) {
                v3 += 1.5
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 = 1e10
            }
        } else {
            var v30 int8 = -128
            v4 = v4
            v30 -= v30
            _ = v30
        }
    case 128:
        if (1 == (v5 | v5)) {
            var v31 int16 = int16(v5)
            _ = v31
        }
    }
    v6 = ((!((268 != 476) || (1 <= 255))) || ((!v7) && (1 >= 1)))
    var v32 uint8 = 1
    if false {
        v5 = v5
        v5 += 0
        if (!(307 == 1)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v33 uint64 = 930
            _ = v33
        }
    } else {
        v34 := GenS7{F0: (42 != -9223372036854775808)}
        v7 = true
        v35 := GenS5{F0: 42, F1: -796, F2: true}
        _ = v34
        _ = v35
    }
    if v7 {
        v2 = append(v2, 296)
        v36 := GenS5{F0: 783, F1: 9223372036854775807, F2: false}
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
