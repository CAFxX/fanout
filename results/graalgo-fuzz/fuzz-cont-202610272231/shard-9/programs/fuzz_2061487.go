package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483647}
    var v3 uint16 = 32767
    v4 := []uint32{}
    delete(v1, 825)
    switch (("\n\t" + ("\n\t" + "a")) + ("hello" + "")) {
    case "foo bar":
        if (1 == 100) {
            v5 := v1[367]
            _ = v5
        } else {
            v6 := GenS1{F0: 818}
            _ = v6
        }
    case "":
        delete(v1, 65535)
    case "hello":
        fmt.Println(v3)
    default:
        var v7 uint16 = 674
        _ = v7
    }
    if ((float32((0.1 + 1.5)) - ((float32(0.0) / float32(-2.25)) / (1e10 / 0.1))) < (1.5 / 1.0)) {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 float64 = float64(-2.25)
                _ = fi9
                _ = v10
            }
            if (("" + "hello") <= "a") {
                fmt.Println(fi8)
                v3 -= 42
            } else {
                fmt.Println(v3, v3)
                _ = len(v1)
                v11 := GenS2{F0: 127}
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v3 += (^(v3))
                    v3 += v3
                    _ = fi12
                }
                var v13 int16 = -1
                var v14 uint64 = 2
                _ = v13
                _ = v14
            } else {
                var v15 uint16 = 32767
                _ = v15
            }
            _ = fi8
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint64 = 42
        if false {
            for fi18 := 0; fi18 < 9; fi18++ {
                switch fi16 {
                case 32767:
                    v19 := GenS3{F0: 583, F1: -48, F2: -612}
                    _ = v19
                default:
                    v17 += 0
                }
                _ = fi18
            }
            v4 = append(v4, 24)
        } else {
            if (-32768 == -1) {
                fmt.Println(v3, fi16)
            } else {
                v1[501] = "a"
            }
            switch 4294967295 {
            case 4294967295:
                delete(v1, 205)
            default:
                if (!(!(v3 == 42))) {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v17 += 4294967295
                        var v21 float32 = (1e10 + (1e10 * 3.14159))
                        if len(v4) > 0 {
                            v4[0] = uint32(v17)
                        }
                        v21 += (1e10 + 3.14159)
                        _ = fi20
                        _ = v21
                    }
                    switch int(v17) {
                    case 9223372036854775807:
                        v22 := v1[42]
                        _ = v22
                    case 207:
                        v1[2] = "a"
                    default:
                        var v23 uint32 = 2
                        _ = v23
                    }
                }
            }
        }
        _ = fi16
        _ = v17
    }
    var v24 int = -90
    v25 := GenS4{F0: 4294967295, F1: false}
    var v26 int16 = 1000
    fmt.Println(v24)
    v3 += (v3 * (v3 ^ (^(v3))))
    if false {
        v27 := v1[65535]
        switch ("foo bar" + (v27 + v27)) {
        case "hello":
            if false {
                var v28 bool = true
                var v29 uint8 = 128
                if true {
                    v2 = append(v2, 2)
                    for fi30 := 0; fi30 < 16; fi30++ {
                        var v31 uint8 = (^(v29))
                        v32 := GenS1{F0: 2147483647}
                        v33 := v1[1]
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                } else {
                    delete(v1, 42)
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v35 := GenS3{F0: -1, F1: 42, F2: 0}
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS5{F0: 127}
                    _ = v36
                }
                _ = v28
                _ = v29
            }
        case "a":
            for fi37 := 0; fi37 < 10; fi37++ {
                if (1e10 > (0.1 / 0.1)) {
                    fmt.Println(v27)
                    switch -1 {
                    case -509:
                        switch v24 {
                        case -1:
                            fmt.Println(v24, v27)
                        case 346:
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                        case 742:
                            var v38 uint8 = 114
                            _ = v38
                        default:
                            switch fi37 {
                            case 9223372036854775807:
                                v2 = append(v2, 986)
                            default:
                                fmt.Println(v3)
                            }
                        }
                    case 491:
                        v4 = append(v4, 1)
                    case 1:
                        for fi39 := 0; fi39 < 6; fi39++ {
                            _ = len(v4)
                            v27 = v27
                            _ = fi39
                        }
                    default:
                        var v40 uint32 = 2147483648
                        _ = v40
                    }
                }
                if false {
                    delete(v1, 0)
                }
                var v41 uint64 = 2
                _ = fi37
                _ = v41
            }
        default:
            fmt.Println(v27)
        }
        v42 := GenS6{F0: true}
        _ = v27
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
