package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []int{474, 108, 284, 0}
    v2 := make(map[string]int)
    var v3 uint16 = 483
    var v4 uint8 = 2
    if len(v1) > 0 {
        v1[0] = -310
    }
    fmt.Println(v4)
    fmt.Println(v4, v4)
    v5 := GenS1{F0: (v4 | 24), F1: 388}
    var v6 int16 = 32767
    switch "\n\t" {
    case "\n\t":
        var v7 int8 = 19
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: v4, F1: 42}
        _ = v8
    case "hello":
        v6 = v6
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            if (32767 != (-(v6))) {
                fmt.Println(v3, v3)
                v4 = 255
                switch "a" {
                case "a":
                    v3 = (^(v3))
                case "\n\t":
                    var v10 int16 = -32768
                    _ = v10
                case "foo bar":
                    v11 := GenS1{F0: 127, F1: 682}
                    _ = v11
                }
            } else {
                switch "a" {
                case "a":
                    if (false && false) {
                        var v12 uint = 43
                        _ = v12
                    }
                case "\n\t":
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v3 -= 796
                        v14 := GenS2{F0: 18446744073709551615, F1: 1}
                        v15 := GenS3{F0: (fi9 ^ fi13), F1: v3}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                var v16 string = ((("\n\t" + "hello") + "a") + "\n\t")
                if (393 > 1) {
                    v17 := GenS4{F0: 410, F1: v3, F2: v6}
                    delete(v2, v16)
                    fmt.Println(fi9, v3)
                    _ = v17
                } else {
                    if true {
                        switch (fi9 & (^(fi9))) {
                        case 0:
                            fmt.Println(v16, v16, v4)
                        case 9223372036854775807:
                            var v18 float32 = 1e10
                            _ = v18
                        default:
                            var v19 int8 = 127
                            _ = v19
                        }
                        fmt.Println(v4)
                        for fi20 := 0; fi20 < 3; fi20++ {
                            v1 = append(v1, fi20)
                            fmt.Println(fi20)
                            if ((^(v4)) > 42) {
                                var v21 uint16 = (v3 << ((42) & 63))
                                _ = v21
                            }
                            _ = fi20
                        }
                    } else {
                        var v22 uint32 = 2
                        for fi23 := 0; fi23 < 13; fi23++ {
                            var v24 int8 = 100
                            var v25 float64 = 0.1
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    }
                }
                _ = v16
            }
            v26 := GenS1{F0: v4, F1: 18446744073709551615}
            v6 = v6
            _ = fi9
            _ = v26
        }
    }
    switch -2147483648 {
    case 4294967295:
        if (true && (0 != 793)) {
            v4 += v4
            v3 += 1
            v27 := v2[("hello" + "")]
            _ = v27
        } else {
            var v28 int = -2147483648
            _ = len(v2)
            _ = v28
        }
    case 32767:
        v29 := GenS1{F0: (v4 % ((v4) | 1)), F1: 1}
        _ = v29
    case -1:
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS3{F0: 4294967295, F1: 32767}
            var v32 uint32 = 166
            for fi33 := 0; fi33 < 19; fi33++ {
                fmt.Println(v32, fi30)
                switch -2147483648 {
                case -731:
                    fmt.Println(v3, fi33, fi33)
                default:
                    v6 -= (v6 / (((v6 | (v6 * int16(v3)))) | 1))
                    var v34 float32 = float32(1.5)
                    v35 := GenS2{F0: 448, F1: 1}
                    if true {
                        var v36 int8 = int8(v4)
                        var v37 uint8 = v4
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS2{F0: 403, F1: 0}
                        fmt.Println(v4, fi33, v6)
                        _ = v38
                    }
                    _ = v34
                    _ = v35
                }
                _ = fi33
            }
            _ = fi30
            _ = v31
            _ = v32
        }
    default:
        if true {
            var v39 int8 = 127
            var v40 uint64 = 530
            _ = v39
            _ = v40
        } else {
            v3 = 0
            if (-2147483648 <= 2147483647) {
                for fi41 := 0; fi41 < 20; fi41++ {
                    v42 := GenS4{F0: 18446744073709551615, F1: v3, F2: (v6 | 42)}
                    _ = fi41
                    _ = v42
                }
                v6 = v6
                if (924 != 946) {
                    v4 = v4
                    v4 = (^(v4))
                }
            }
        }
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        v4 = v4
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v6, v3)
        case "foo bar":
            if false {
                v1 = append(v1, fi43)
            }
        }
        for fi44 := 0; fi44 < 1; fi44++ {
            _ = len(v1)
            _ = fi44
        }
        _ = fi43
    }
    if false {
        v4 += v4
        v6 = -1
        if true {
            v45 := GenS5{F0: 255}
            var v46 uint8 = v4
            v4 += v4
            _ = v45
            _ = v46
        }
    }
    v47 := GenS2{F0: 999, F1: 42}
    v48 := v2["foo bar"]
    var v49 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
