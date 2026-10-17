package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint64, p1 string) uint8 {
    var v7 uint32 = 2
    fmt.Println(p0, p1, v7)
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 = (v7 >> ((255) & 63))
        _ = fi8
    }
    v7 -= (v7 % ((2147483647) | 1))
    v9 := GenS1{F0: 1, F1: 1, F2: uint16(p0)}
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 2
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    v5 := make(map[uint32]bool)
    var v6 float32 = v1
    v6 -= v1
    v1 += v6
    fmt.Println(v1, v6)
    if (((0.1 * float32((0.1 * 1.5))) - v6) > v1) {
        v3[2147483648] = 2147483647
        v10 := GenS2{F0: -684, F1: 255, F2: "foo bar"}
        v1 -= v6
        _ = v10
    } else {
        v1 -= v6
        fmt.Println(v1, v4, v4)
        if (!true) {
            if (4294967295 <= 42) {
                v11 := GenS3{F0: 74, F1: v4, F2: 128}
                _ = v11
            } else {
                v6 += (float32(0.0) + ((float32(0.1) * (v6 / v6)) * ((v6 / 1.5) + (v6 + (v1 * (1e10 * ((float32(-2.25) * 0.0) - float32((((1.5 * ((3.14159 * 0.1) + -2.25)) / -2.25) / -2.25)))))))))
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := v5[0]
                v14 := []int{fi12}
                if v13 {
                    fmt.Println(v13, v1, v4)
                } else {
                    delete(v2, 788)
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := GenS3{F0: 0, F1: (v4 / ((v4) | 1)), F2: 128}
            _ = v15
        } else {
            v16 := GenS2{F0: -1, F1: 898, F2: "\n\t"}
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v6, v1, v4)
                var v18 int16 = -32768
                v18 -= (v18 - 42)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v1 += ((v1 / v6) * (float32(1e10) * 0.0))
                fmt.Println(v1, v1, v1)
                if (!false) {
                    if (65535 >= 0) {
                        fmt.Println(v6)
                        var v20 int32 = v4
                        if true {
                            v21 := GenS4{F0: -334, F1: 0}
                            var v22 uint16 = 65535
                            v20 -= v4
                            _ = v21
                            _ = v22
                        } else {
                            v4 += (v4 << ((42) & 63))
                        }
                        _ = v20
                    } else {
                        switch fi19 {
                        case 506:
                            switch ("a" + "hello") {
                            case "foo bar":
                                if (255 >= 2) {
                                    if true {
                                        v4 += (v4 / ((v4) | 1))
                                        v6 = v1
                                        fmt.Println(v6, v4)
                                    }
                                }
                            case "a":
                                var v23 int = fi19
                                _ = v23
                            case "\n\t":
                                v5[0] = (!true)
                            }
                        case 127:
                            v24 := GenS1{F0: 255, F1: 128, F2: 691}
                            _ = v24
                        default:
                            var v25 uint32 = 53
                            _ = v25
                        }
                        var v26 float32 = v1
                        _ = v26
                    }
                }
                _ = fi19
            }
            _ = v16
        }
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS3{F0: 0, F1: (v4 / (((v4 / (((v4 >> ((42) & 63))) | 1))) | 1)), F2: 60}
        _ = fi27
        _ = v28
    }
    fmt.Println(v6, v4, v4)
    v29 := v5[65535]
    v4 = v4
    if v29 {
        v30 := GenS4{F0: 1000, F1: 4294967295}
        switch "a" {
        case "":
            switch 218 {
            case 861:
                v4 = v4
            case -561:
                v6 -= (1e10 * (1.5 / (0.0 - float32(1e10))))
            case 128:
                var v31 int64 = 1
                _ = v31
            default:
                var v32 int16 = 42
                _ = v32
            }
        case "foo bar":
            var v33 int8 = 127
            _ = v33
        case "hello":
            v29 = false
        default:
            var v34 int16 = -1
            _ = v34
        }
        _ = v30
    }
    fmt.Println(v4, v29, v29)
    fmt.Println(v6)
    v35 := GenS4{F0: -32768, F1: 2147483648}
    if false {
        var v36 string = "foo bar"
        for fi37 := 0; fi37 < 4; fi37++ {
            v29 = v29
            var v38 string = "hello"
            if v29 {
                var v39 uint64 = 197
                v40 := GenS1{F0: 65535, F1: 39, F2: 65535}
                var v41 uint16 = 32767
                _ = v39
                _ = v40
                _ = v41
            } else {
                var v42 int64 = 42
                _ = v42
            }
            _ = fi37
            _ = v38
        }
        if false {
            for fi43 := 0; fi43 < 17; fi43++ {
                var v44 int64 = -1
                fmt.Println(v6, v36)
                _ = fi43
                _ = v44
            }
            var v45 float32 = (0.1 - v1)
            var v46 int64 = 42
            _ = v45
            _ = v46
        } else {
            v47 := v2[32768]
            var v48 uint = 2
            _ = v47
            _ = v48
        }
        _ = v36
    }
    v4 = (v4 * 42)
    v6 += 1e10
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
