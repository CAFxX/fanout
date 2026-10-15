package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint64; F2 bool }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0() uint32 {
    if true {
        v5 := []int{0, -163}
        if true {
            _ = len(v5)
            var v6 int16 = -1
            if len(v5) > 0 {
                v5[0] = -681
            }
            _ = v6
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    _ = len(v5)
                    fmt.Println(fi7)
                    _ = fi8
                }
                _ = fi7
            }
            switch -601 {
            case 9223372036854775807:
                if (2147483648 < 1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            default:
                if (0 <= 1) {
                    v9 := GenS1{F0: 0}
                    _ = v9
                } else {
                    switch "hello" {
                    case "":
                        var v10 int8 = 8
                        _ = v10
                    case "\n\t":
                        var v11 uint64 = 42
                        _ = v11
                    }
                    var v12 bool = true
                    _ = v12
                }
            }
        }
        _ = v5
    }
    var v13 int8 = 100
    v13 -= v13
    _ = v13
    return 42
}
func genFn1(p0 int64) uint {
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: 65535}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 255}
    switch 0 {
    case 32768:
        var v17 uint32 = 2147483647
        _ = v17
    case -9223372036854775808:
        v18 := GenS2{F0: 65535, F1: 551}
        _ = v18
    case -620:
        var v19 int32 = 0
        _ = v19
    }
    _ = p0
    _ = v16
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 875
    var v3 int = 336
    var v4 float32 = 1e10
    if (!false) {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch v3 {
            case 32767:
                v21 := GenS1{F0: 32767}
                _ = v21
            case 42:
                if true {
                    var v22 int16 = 0
                    fmt.Println(v4)
                    _ = v22
                }
            }
            _ = fi20
        }
        if ((true && (true || false)) && (-32768 > -1)) {
            v23 := GenS1{F0: 0}
            v24 := v1["a"]
            v4 = float32(((3.14159 - 0.0) * (((((float64((v4 + v4)) * -2.25) * float64(0.0)) / -2.25) * (-2.25 / 1.5)) / float64(v4))))
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 2; fi25++ {
                v3 -= fi25
                fmt.Println(v3, v2)
                for fi26 := 0; fi26 < 17; fi26++ {
                    v4 = ((v4 + float32(3.14159)) + ((((-2.25 - v4) - 0.0) * (float32(1.5) + 1.5)) - ((0.1 + v4) * 1.5)))
                    v27 := GenS2{F0: 1, F1: v2}
                    switch (v3 % ((-1) | 1)) {
                    case -802:
                        var v28 int32 = 0
                        _ = v28
                    case 9223372036854775807:
                        var v29 uint8 = 2
                        _ = v29
                    case 2147483647:
                        if true {
                            var v30 string = ("hello" + ("" + "\n\t"))
                            var v31 uint64 = v2
                            fmt.Println(v30)
                            _ = v30
                            _ = v31
                        } else {
                            switch fi26 {
                            case 10:
                                if true {
                                    var v32 int64 = -924
                                    switch fi25 {
                                    case -656:
                                        fmt.Println(fi26)
                                    }
                                    _ = v32
                                } else {
                                    v3 -= 1
                                    var v33 uint64 = v2
                                    var v34 uint32 = 2147483647
                                    _ = v33
                                    _ = v34
                                }
                            default:
                                if false {
                                    switch (v3 & fi25) {
                                    case -2147483648:
                                        if ((v4 + 0.0) < float32(float64(v4))) {
                                            v35 := GenS3{F0: 1, F1: v2, F2: (0 <= 0)}
                                            _ = v35
                                        }
                                    case 32768:
                                        var v36 int64 = 1
                                        _ = v36
                                    }
                                    var v37 int16 = -837
                                    _ = v37
                                }
                            }
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            var v38 uint32 = 992
            _ = v38
        }
        var v39 uint64 = (v2 - v2)
        _ = v39
    }
    v1["a"] = (v3 ^ (^(v3)))
    v3 -= -1
    v40 := GenS4{F0: 2147483647, F1: 127}
    v41 := GenS1{F0: 1}
    v42 := GenS5{F0: genFn0(), F1: 2147483647}
    _ = len(v1)
    v2 += (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
