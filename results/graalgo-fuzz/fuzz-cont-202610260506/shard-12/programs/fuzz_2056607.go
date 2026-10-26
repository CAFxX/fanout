package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint; F2 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    v4 := make(map[int]int)
    v5 := GenS1{F0: 0}
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := v1["foo bar"]
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(fi8)
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    if true {
        var v10 float64 = 1.5
        for fi11 := 0; fi11 < 20; fi11++ {
            if (2 >= 42) {
                switch "a" {
                case "foo bar":
                    switch (fi11 << ((255) & 63)) {
                    case -864:
                        v12 := GenS2{F0: fi11}
                        _ = v12
                    case -9223372036854775808:
                        var v13 int = 2147483647
                        _ = v13
                    case 32767:
                        for fi14 := 0; fi14 < 20; fi14++ {
                            var v15 uint16 = 0
                            _ = fi14
                            _ = v15
                        }
                    default:
                        v10 += (3.14159 - 0.1)
                    }
                default:
                    var v16 float32 = ((0.1 - 1e10) / 0.1)
                    _ = v16
                }
                fmt.Println(v10, fi11)
                if true {
                    v10 = (1e10 + ((-2.25 + v10) * float64(float32(float64(1.5)))))
                    var v17 string = ("foo bar" + ("foo bar" + "hello"))
                    var v18 int = fi11
                    _ = v17
                    _ = v18
                }
            } else {
                fmt.Println(v10)
                v3 = append(v3, 42)
                var v19 uint32 = 2147483648
                _ = v19
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v20 float32 = 0.1
        v21 := v4[2147483647]
        v20 -= v20
        _ = v20
        _ = v21
    }
    if ((1.5 - float64(0.1)) != (((0.0 * -2.25) - (float64(3.14159) - float64(0.1))) / 1.5)) {
        if false {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int32 = -252
                var v24 uint32 = 4294967295
                var v25 int64 = 42
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
    } else {
        if false {
            var v26 uint64 = 1
            v2[(("" + "foo bar") + ("\n\t" + ""))] = 755
            _ = v26
        }
        var v27 int16 = -1
        var v28 string = "foo bar"
        _ = v27
        _ = v28
    }
    var v29 int16 = 728
    var v30 float64 = (1e10 * (((float64(((float32(0.0) - 0.1) * float32(float64((3.14159 / 1.0))))) / 1e10) / 1.5) + 1e10))
    v29 += (v29 << ((255) & 63))
    fmt.Println(v30, v29)
    v29 = v29
    if false {
        fmt.Println(v30)
        switch 128 {
        case -2147483648:
            v29 = v29
        case 127:
            v31 := GenS3{F0: (v29 << ((794) & 63))}
            _ = v31
        default:
            for fi32 := 0; fi32 < 12; fi32++ {
                v29 += v29
                _ = fi32
            }
        }
    } else {
        for fi33 := 0; fi33 < 10; fi33++ {
            v30 += v30
            switch 32768 {
            case -212:
                if (42 >= 298) {
                    fmt.Println(v30, v30)
                    for fi34 := 0; fi34 < 14; fi34++ {
                        var v35 int64 = 1
                        v36 := GenS4{F0: 65535}
                        v37 := GenS3{F0: v29}
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                } else {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 int = fi38
                        var v40 float32 = 1.5
                        if ((!(1 != 76)) || true) {
                            v41 := GenS1{F0: 65535}
                            delete(v1, "")
                            _ = v41
                        } else {
                            v29 += (^(v29))
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                }
            }
            _ = fi33
        }
        if (v29 <= (v29 / ((v29) | 1))) {
            if (255 < 195) {
                var v42 string = "a"
                if (!true) {
                    if false {
                        v42 = v42
                        var v43 uint = 42
                        if (!(-812 >= -388)) {
                            if (uint64(v43) > 4294967295) {
                                v30 = (v30 / v30)
                                v43 -= v43
                                var v44 float64 = v30
                                _ = v44
                            } else {
                                v45 := GenS5{F0: uint8(v43)}
                                v3 = append(v3, 485)
                                if false {
                                    v46 := GenS3{F0: (v29 ^ (v29 >> (((v43 / ((v43) | 1))) & 63)))}
                                    fmt.Println(v29, v43)
                                    fmt.Println(v43, v30)
                                    _ = v46
                                }
                                _ = v45
                            }
                        }
                        _ = v43
                    } else {
                        var v47 int8 = 85
                        v48 := GenS6{F0: 25, F1: 0, F2: 42}
                        v42 = v42
                        _ = v47
                        _ = v48
                    }
                    v49 := GenS3{F0: v29}
                    _ = len(v3)
                    _ = v49
                }
                _ = v42
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
