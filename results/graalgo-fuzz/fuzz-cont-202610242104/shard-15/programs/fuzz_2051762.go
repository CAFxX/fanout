package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    v7 := []int{224, -103, 431}
    _ = p0
    _ = v7
    return 770
}
func genFn1() uint8 {
    if true {
        var v8 int = 127
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: (3.14159 <= 0.1)}
    v10 := []int{-3}
    v11 := GenS2{F0: ("foo bar" + "a")}
    var v12 int16 = 316
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 143
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]uint32)
    var v3 uint8 = 0
    v4 := make(map[uint32]uint32)
    var v5 float32 = 3.14159
    var v6 int = int(v3)
    v6 += (v6 >> ((42) & 63))
    var v13 bool = (!true)
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := []int{2147483647, (v1 % ((v1) | 1)), -438}
        _ = fi14
        _ = v15
    }
    var v16 int8 = 42
    v2[(^(v6))] = 552
    v3 = (^(v3))
    v17 := []int{}
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v16, v5, v13)
            _ = fi19
        }
        v20 := v4[4294967295]
        _ = fi18
        _ = v20
    }
    v17 = append(v17, v1)
    v21 := GenS1{F0: 2, F1: v13}
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint16 = 0
        if ((!(-1 < 1)) || true) {
            switch v1 {
            case 1:
                if v13 {
                    if (2 > 42) {
                        v16 = (v16 >> ((517) & 63))
                        if false {
                            v13 = (177 < 206)
                            for fi24 := 0; fi24 < 13; fi24++ {
                                if (1 < -32768) {
                                    fmt.Println(v13, v1, v23)
                                    fmt.Println(v3, fi24)
                                    var v25 uint64 = 611
                                    switch "\n\t" {
                                    case "hello":
                                        _ = len(v4)
                                    case "\n\t":
                                        var v26 uint16 = v23
                                        _ = v26
                                    default:
                                        var v27 float64 = 3.14159
                                        _ = v27
                                    }
                                    _ = v25
                                } else {
                                    var v28 uint32 = 256
                                    v13 = false
                                    v5 -= 1e10
                                    _ = v28
                                }
                                _ = fi24
                            }
                            v16 += v16
                        } else {
                            switch (v6 % (((v1 | 9223372036854775807)) | 1)) {
                            case -475:
                                switch "a" {
                                case "hello":
                                    var v29 int16 = -32768
                                    _ = v29
                                }
                            }
                        }
                    }
                    if v13 {
                        v1 += 4294967295
                    }
                    v30 := GenS1{F0: 243, F1: v13}
                    _ = v30
                } else {
                    var v31 int32 = 463
                    switch "" {
                    case "":
                        v13 = v13
                    case "\n\t":
                        v32 := GenS1{F0: 0, F1: true}
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 13; fi33++ {
                            if v13 {
                                var v34 int32 = 1
                                var v35 int8 = v16
                                if false {
                                    fmt.Println(v35)
                                    var v36 float32 = v5
                                    _ = v36
                                } else {
                                    fmt.Println(v35, fi33, v23)
                                    switch "foo bar" {
                                    case "a":
                                        v5 = (v5 * -2.25)
                                    case "foo bar":
                                        v13 = v13
                                    case "\n\t":
                                        var v37 float64 = (3.14159 * 0.0)
                                        _ = v37
                                    }
                                    v3 -= (v3 / (((v3 % ((128) | 1))) | 1))
                                    fmt.Println(v16, v23)
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                }
                                fmt.Println(v5, fi22)
                                v38 := v2[v1]
                                v39 := GenS1{F0: 924, F1: v13}
                                fmt.Println(v35, v1)
                                _ = v34
                                _ = v35
                                _ = v38
                                _ = v39
                            } else {
                                var v40 int16 = 32767
                                _ = v40
                            }
                            if (v13 || (42 < 1)) {
                                fmt.Println(fi22)
                                var v41 uint8 = (^(v3))
                                if ((v3 << ((2) & 63)) <= genFn1()) {
                                    v16 -= 13
                                    var v42 uint = 262
                                    fmt.Println(v16)
                                    _ = v42
                                }
                                _ = v41
                            }
                            _ = fi33
                        }
                    }
                    _ = v31
                }
            }
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            if v13 {
                _ = len(v17)
                v44 := GenS2{F0: ("a" + ("" + "a"))}
                _ = v44
            } else {
                v45 := GenS3{F0: -177, F1: true}
                v6 -= -2147483648
                _ = v45
            }
            v46 := GenS4{F0: "a", F1: (v23 % ((v23) | 1))}
            _ = fi43
            _ = v46
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    v47 := GenS4{F0: (("hello" + ("\n\t" + "\n\t")) + ""), F1: 391}
    var v48 float32 = (v5 * float32(((0.0 / float64((float32(0.1) / (float32(((1.5 - float64(v5)) / 1e10)) - 0.0)))) + 1.5)))
    switch v1 {
    case 9223372036854775807:
        v49 := GenS3{F0: 846, F1: v13}
        _ = v49
    case 0:
        v16 += 6
    default:
        switch (^(v6)) {
        case 4294967295:
            for fi50 := 0; fi50 < 10; fi50++ {
                fmt.Println(v1, v1, v1)
                _ = fi50
            }
        case 145:
            var v51 int8 = 100
            _ = v51
        case 752:
            for fi52 := 0; fi52 < 19; fi52++ {
                var v53 int16 = 0
                _ = fi52
                _ = v53
            }
        default:
            v54 := GenS3{F0: int32(v1), F1: false}
            _ = v54
        }
    }
    v48 -= v48
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
