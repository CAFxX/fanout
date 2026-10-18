package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    var v1 int16 = -1
    var v2 uint16 = 302
    var v3 uint32 = 42
    var v4 int16 = (v1 * 560)
    var v5 float32 = 0.1
    v6 := GenS1{F0: 2147483647}
    var v7 int64 = -708
    v8 := GenS1{F0: 2147483647}
    if true {
        v9 := []int{223, 128, -164}
        if false {
            v5 = v5
            v3 = v3
        }
        _ = v9
    } else {
        if false {
            if false {
                if true {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v3 = (v3 * (v3 / ((v3) | 1)))
                        fmt.Println(v3, v4, v1)
                        _ = fi10
                    }
                    v11 := GenS1{F0: -1}
                    _ = v11
                } else {
                    v12 := []int{127}
                    v4 = v1
                    _ = v12
                }
                fmt.Println(v1)
            }
            v13 := GenS2{F0: 2147483647, F1: 2}
            v14 := GenS1{F0: -1}
            _ = v13
            _ = v14
        }
        var v15 int64 = (v7 | 160)
        if true {
            if ((((0.1 + v5) - float32((float64(v5) - 0.0))) / v5) == v5) {
                v4 -= (v4 << ((685) & 63))
                fmt.Println(v15)
                if (true || true) {
                    v16 := []int{880, 79, int(v3)}
                    fmt.Println(v1, v4)
                    if false {
                        v17 := GenS2{F0: 173, F1: 42}
                        v7 -= (v15 << ((2) & 63))
                        _ = v17
                    }
                    _ = v16
                }
            }
            var v18 float64 = (0.1 + 1.5)
            _ = v18
        }
        _ = v15
    }
    var v19 uint = 18446744073709551615
    v20 := []int{42, 127}
    v21 := GenS2{F0: int(v2), F1: 1}
    v1 -= v1
    if true {
        if true {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 uint = 206
                switch (^(fi22)) {
                case 32768:
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if true {
                            v19 = (v23 + 375)
                            var v25 uint64 = 42
                            switch fi24 {
                            case 42:
                                v26 := GenS2{F0: (-(fi22)), F1: 27}
                                fmt.Println(v19, fi24)
                                v27 := GenS3{F0: 0}
                                _ = v26
                                _ = v27
                            case 2147483647:
                                v28 := GenS1{F0: 0}
                                _ = v28
                            case -1:
                                if true {
                                    var v29 bool = false
                                    v19 = (v23 << ((v23) & 63))
                                    _ = v29
                                }
                            default:
                                v23 = (v23 ^ v23)
                            }
                            _ = v25
                        } else {
                            v4 -= (v1 | 0)
                        }
                        v3 += v3
                        if true {
                            fmt.Println(v3, fi24, v4)
                        }
                        _ = fi24
                    }
                case -2147483648:
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 uint32 = (v3 / ((uint32(v2)) | 1))
                        switch ("a" + ("hello" + "foo bar")) {
                        case "":
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        case "foo bar":
                            if len(v20) > 0 {
                                v20[0] = fi30
                            }
                        case "a":
                            v19 -= uint(v2)
                        }
                        _ = fi30
                        _ = v31
                    }
                case -9223372036854775808:
                    if len(v20) > 0 {
                        v20[0] = (fi22 + fi22)
                    }
                }
                v20 = append(v20, fi22)
                _ = fi22
                _ = v23
            }
        } else {
            v19 -= v19
        }
        v32 := GenS2{F0: 4294967295, F1: 2}
        _ = v32
    } else {
        for fi33 := 0; fi33 < 7; fi33++ {
            if len(v20) > 0 {
                v20[0] = fi33
            }
            if false {
                _ = len(v20)
            }
            switch int(v7) {
            case 32767:
                if true {
                    var v34 uint = 4294967295
                    v19 -= (v34 % (((v19 * (v34 * (v19 + (v34 ^ (v34 & 65535)))))) | 1))
                    _ = v34
                }
            }
            _ = fi33
        }
        if true {
            var v35 float32 = (float32(1e10) / 1.0)
            if (1e10 <= v5) {
                var v36 int64 = (-(v7))
                _ = v36
            }
            _ = v35
        } else {
            if (76 == 2) {
                v4 = (v1 | v1)
            } else {
                switch -9223372036854775808 {
                case 0:
                    var v37 float64 = 1e10
                    _ = v37
                case -1:
                    v19 = (^(v19))
                }
                fmt.Println(v7, v5)
                var v38 int16 = 1000
                _ = v38
            }
        }
        var v39 uint64 = 0
        _ = v39
    }
    v40 := GenS3{F0: -819}
    for fi41 := 0; fi41 < 7; fi41++ {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        if len(v20) > 0 {
            v20[0] = fi41
        }
        var v42 float32 = (v5 / 1.5)
        _ = fi41
        _ = v42
    }
    if len(v20) > 0 {
        v20[0] = -402
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 float64 = 1e10
        fmt.Println(v5)
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        if (!(!false)) {
            v46 := GenS3{F0: 662}
            v47 := GenS4{F0: (v4 / (((v1 * v4)) | 1)), F1: -1}
            _ = v46
            _ = v47
        } else {
            for fi48 := 0; fi48 < 5; fi48++ {
                switch ("a" + "") {
                case "":
                    var v49 int64 = v7
                    _ = v49
                }
                var v50 string = "hello"
                v51 := GenS4{F0: 0, F1: -18}
                fmt.Println(v4, v7, v2)
                if (!(!(v2 < 679))) {
                    v3 -= v3
                } else {
                    v3 += 892
                    v5 = v5
                }
                _ = fi48
                _ = v50
                _ = v51
            }
            v3 = v3
        }
        v5 = float32((((float64(v5) + 0.1) * (((1.5 * (-2.25 - 1.5)) + 0.1) + (3.14159 + 0.0))) / 1.5))
        fmt.Println(v19)
        _ = fi45
    }
    v4 = (v1 % (((v1 ^ int16(v2))) | 1))
    for fi52 := 0; fi52 < 10; fi52++ {
        var v53 uint32 = (v3 >> (((v19 * 0)) & 63))
        var v54 float64 = (1e10 / 3.14159)
        _ = fi52
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        fmt.Println(v7)
        fmt.Println(v3)
        _ = len(v20)
        _ = fi55
    }
    if (!true) {
        var v56 float64 = (float64(v5) / ((-2.25 / float64(v5)) / 1.5))
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
