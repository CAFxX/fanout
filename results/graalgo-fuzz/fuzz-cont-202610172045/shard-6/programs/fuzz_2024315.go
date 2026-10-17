package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 string; F1 int; F2 bool }
type GenS8 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []int{0, 42, -584}
    var v2 uint64 = 713
    var v3 float32 = 1.5
    v4 := make(map[int]string)
    var v5 int8 = -114
    for fi6 := 0; fi6 < 15; fi6++ {
        if (!((fi6 & fi6) <= (fi6 / (((fi6 | -135)) | 1)))) {
            if (-1 >= 32767) {
                v5 -= (^(v5))
            } else {
                v7 := GenS1{F0: 1, F1: int32(v2), F2: 100}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v7
            }
            for fi8 := 0; fi8 < 16; fi8++ {
                if ((v5 ^ v5) != (v5 | (v5 + (v5 % ((v5) | 1))))) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v9 uint16 = 65535
                    _ = v9
                } else {
                    delete(v4, 4294967295)
                }
                _ = fi8
            }
        } else {
            v10 := GenS1{F0: 65535, F1: -1, F2: (v5 | (v5 << ((255) & 63)))}
            v3 += -2.25
            var v11 uint8 = 162
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: "\n\t"}
        _ = fi12
        _ = v13
    }
    fmt.Println(v3, v3, v2)
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint32 = 0
        var v16 int8 = 100
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v5 += 127
        _ = fi17
    }
    var v18 uint16 = 65535
    var v19 float32 = 1e10
    _ = len(v4)
    switch -635 {
    case -790:
        for fi20 := 0; fi20 < 3; fi20++ {
            if true {
                if true {
                    v21 := GenS3{F0: 18446744073709551615, F1: 4294967295}
                    _ = v21
                }
                for fi22 := 0; fi22 < 1; fi22++ {
                    switch 127 {
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v23 := GenS4{F0: -1, F1: 417}
                        _ = v23
                    }
                    var v24 uint = 1
                    _ = fi22
                    _ = v24
                }
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 string = "a"
                    v4[9223372036854775807] = v26
                    delete(v4, fi20)
                    _ = fi25
                    _ = v26
                }
            } else {
                if true {
                    var v27 int = 976
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v3 = v19
                        delete(v4, fi20)
                        _ = fi28
                    }
                    _ = v27
                } else {
                    var v29 bool = (v19 < v3)
                    var v30 bool = v29
                    delete(v4, fi20)
                    _ = v29
                    _ = v30
                }
                v3 = (float32(1e10) / 3.14159)
            }
            v3 += float32(((0.1 + 1e10) - ((0.0 - -2.25) - (-2.25 + 3.14159))))
            _ = fi20
        }
    case 9223372036854775807:
        v1 = append(v1, 315)
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        if ((v3 * (1e10 / v19)) != ((1e10 / (v3 - ((0.0 - v19) - v19))) / v19)) {
            var v32 int64 = 1
            if (-1 >= (v5 * (-(v5)))) {
                for fi33 := 0; fi33 < 18; fi33++ {
                    _ = len(v4)
                    fmt.Println(v32)
                    _ = fi33
                }
                v18 = v18
            }
            _ = v32
        } else {
            var v34 int16 = -42
            v35 := GenS4{F0: 0, F1: v18}
            var v36 int16 = v34
            _ = v34
            _ = v35
            _ = v36
        }
        v37 := GenS5{F0: 15, F1: 833, F2: 0}
        _ = fi31
        _ = v37
    }
    var v38 int16 = 885
    v39 := GenS6{F0: (v18 & uint16(v2)), F1: -2147483648}
    v19 = (1.5 + float32(((3.14159 / 0.1) - 1.5)))
    for fi40 := 0; fi40 < 7; fi40++ {
        var v41 float64 = (1.5 - (1.5 / 1.0))
        for fi42 := 0; fi42 < 1; fi42++ {
            fmt.Println(v3, v5)
            switch (fi40 + fi42) {
            case 2147483647:
                var v43 uint32 = 98
                v18 -= (^(v18))
                v44 := GenS2{F0: ((("\n\t" + "") + "\n\t") + "foo bar")}
                _ = v43
                _ = v44
            }
            _ = fi42
        }
        switch "" {
        case "foo bar":
            v45 := GenS7{F0: "hello", F1: (fi40 ^ int(v18)), F2: (v18 >= 937)}
            _ = v45
        }
        _ = fi40
        _ = v41
    }
    var v46 float64 = 1.5
    if true {
        v5 += v5
        for fi47 := 0; fi47 < 15; fi47++ {
            var v48 float64 = float64(v3)
            for fi49 := 0; fi49 < 16; fi49++ {
                var v50 uint64 = (v2 / (((v2 * 1)) | 1))
                _ = fi49
                _ = v50
            }
            var v51 float64 = v48
            _ = fi47
            _ = v48
            _ = v51
        }
    } else {
        switch -2147483648 {
        case 0:
            v52 := GenS8{F0: -1, F1: v18}
            _ = v52
        default:
            v46 = v46
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v46
    _ = v5
}
