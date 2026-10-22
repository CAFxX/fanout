package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8, p1 uint32, p2 int) string {
    var v7 int8 = 26
    v8 := GenS1{F0: (p0 & p0)}
    switch p2 {
    case 502:
        var v9 int16 = -637
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("foo bar" + "hello")
}

func main() {
    var v1 uint32 = 2147483647
    var v2 string = ("a" + "foo bar")
    var v3 float64 = (-2.25 / (1.5 * (1e10 + 0.1)))
    var v4 uint32 = (^(v1))
    v5 := []int{1}
    v6 := []int{2147483647, -444, 518}
    var v10 int64 = 0
    _ = len(v6)
    if len(v5) > 0 {
        v5[0] = 128
    }
    if (!true) {
        switch 32767 {
        case 42:
            switch (("hello" + v2) + (v2 + "\n\t")) {
            case "a":
                for fi11 := 0; fi11 < 9; fi11++ {
                    if false {
                        switch fi11 {
                        case 32768:
                            if (((-2.25 - (0.0 - (3.14159 * 3.14159))) * 1.5) != -2.25) {
                                v1 = v4
                                v3 -= ((1e10 * (1.5 - ((3.14159 / 1.5) - 0.1))) * v3)
                                v1 = 2147483647
                            } else {
                                var v12 string = v2
                                v3 -= float64(1.5)
                                _ = v12
                            }
                        case -2147483648:
                            v6 = append(v6, fi11)
                        default:
                            _ = len(v5)
                        }
                        var v13 int16 = 17
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 18446744073709551615}
                        switch "foo bar" {
                        case "a":
                            if (42 >= 128) {
                                var v15 uint16 = 42
                                _ = len(v6)
                                var v16 float32 = float32(v3)
                                _ = v15
                                _ = v16
                            } else {
                                fmt.Println(v10)
                                if false {
                                    if false {
                                        var v17 string = ("a" + (v2 + (v2 + "foo bar")))
                                        for fi18 := 0; fi18 < 1; fi18++ {
                                            var v19 uint8 = 171
                                            var v20 uint64 = 42
                                            v21 := GenS3{F0: 772, F1: 42, F2: 32767}
                                            _ = fi18
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        v3 -= -2.25
                                        _ = v17
                                    } else {
                                        v22 := GenS4{F0: true}
                                        _ = v22
                                    }
                                    switch 149 {
                                    case -843:
                                        switch "\n\t" {
                                        case "hello":
                                            if true {
                                                for fi23 := 0; fi23 < 4; fi23++ {
                                                    var v24 uint8 = 128
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi23 >> ((1) & 63))
                                                    }
                                                    v25 := GenS3{F0: 32767, F1: 42, F2: 65535}
                                                    _ = fi23
                                                    _ = v24
                                                    _ = v25
                                                }
                                                var v26 uint = 742
                                                v27 := GenS1{F0: 221}
                                                _ = v26
                                                _ = v27
                                            }
                                        case "":
                                            v6 = append(v6, (-(fi11)))
                                        }
                                    case 127:
                                        switch "foo bar" {
                                        case "":
                                            if (0 >= 0) {
                                                v28 := GenS4{F0: false}
                                                fmt.Println(v3)
                                                _ = len(v6)
                                                _ = v28
                                            } else {
                                                v1 += v4
                                                for fi29 := 0; fi29 < 3; fi29++ {
                                                    fmt.Println(v1, v10)
                                                    var v30 int8 = 100
                                                    _ = fi29
                                                    _ = v30
                                                }
                                            }
                                        case "a":
                                            var v31 int16 = 32767
                                            _ = v31
                                        }
                                    }
                                    v32 := GenS2{F0: 99}
                                    _ = v32
                                }
                            }
                        case "\n\t":
                            fmt.Println(v3, v10)
                        case "hello":
                            if len(v5) > 0 {
                                v5[0] = (fi11 << ((255) & 63))
                            }
                        }
                        for fi33 := 0; fi33 < 5; fi33++ {
                            switch (fi11 ^ fi11) {
                            case 0:
                                fmt.Println(v10)
                            case 1:
                                switch fi11 {
                                case 985:
                                    v4 -= 4294967295
                                case -218:
                                    var v34 int = -1
                                    _ = v34
                                default:
                                    switch "a" {
                                    case "hello":
                                        v35 := GenS4{F0: true}
                                        _ = v35
                                    case "\n\t":
                                        switch ((v2 + "foo bar") + (v2 + v2)) {
                                        case "\n\t":
                                            var v36 int16 = 0
                                            _ = v36
                                        case "":
                                            v1 = 2147483648
                                        case "foo bar":
                                            v37 := GenS4{F0: true}
                                            _ = v37
                                        default:
                                            switch fi11 {
                                            case 1:
                                                if true {
                                                    var v38 int16 = -52
                                                    if ((float32(((v3 / v3) / -2.25)) + 3.14159) < float32(v3)) {
                                                        v4 = (v4 - (v1 % (((v1 + v1)) | 1)))
                                                        var v39 int64 = (v10 + v10)
                                                        if (!true) {
                                                            fmt.Println(v4, v10, fi11)
                                                        }
                                                        _ = v39
                                                    } else {
                                                        fmt.Println(fi11, v10, fi33)
                                                        v40 := GenS4{F0: true}
                                                        _ = v40
                                                    }
                                                    if false {
                                                        var v41 float32 = (float32(0.1) + -2.25)
                                                        switch (fi33 % ((fi11) | 1)) {
                                                        case -9223372036854775808:
                                                            var v42 bool = false
                                                            _ = v42
                                                        }
                                                        _ = v41
                                                    } else {
                                                        var v43 uint32 = v4
                                                        var v44 int = -819
                                                        if true {
                                                            v3 = float64(1e10)
                                                            var v45 string = "\n\t"
                                                            _ = v45
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    _ = v38
                                                }
                                            }
                                        }
                                    }
                                }
                            case 4294967295:
                                v3 += (0.0 + (v3 * ((((((0.0 / v3) + float64(3.14159)) * (((v3 / v3) + v3) - v3)) - ((v3 * 0.1) / (v3 * 0.1))) * v3) * ((3.14159 + (v3 + v3)) / v3))))
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            fmt.Println(fi11, fi11, v3)
                            _ = fi33
                        }
                        _ = v14
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi11
                }
            }
        case 372:
            v46 := GenS5{F0: 1, F1: 1}
            _ = v46
        }
    } else {
        switch "hello" {
        case "\n\t":
            switch int(v1) {
            case 9223372036854775807:
                v47 := GenS5{F0: 42, F1: 128}
                _ = v47
            default:
                v1 = v1
            }
        case "hello":
            _ = len(v6)
        default:
            if (!(128 == -948)) {
                fmt.Println(v4, v10, v4)
                switch int(v1) {
                case 0:
                    v48 := GenS5{F0: 1, F1: 176}
                    _ = v48
                case 128:
                    fmt.Println(v2)
                case 42:
                    var v49 int64 = (v10 % ((v10) | 1))
                    _ = v49
                }
                var v50 int64 = (v10 << ((42) & 63))
                _ = v50
            }
        }
        fmt.Println(v1, v3, v10)
        var v51 int = -9223372036854775808
        _ = v51
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        v53 := GenS2{F0: 511}
        _ = fi52
        _ = v53
    }
    switch "\n\t" {
    case "a":
        v4 = v1
    }
    if false {
        if true {
            v1 += v1
        }
    } else {
        var v54 uint8 = 98
        _ = v54
    }
    for fi55 := 0; fi55 < 12; fi55++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v56 int8 = -108
        v6 = append(v6, fi55)
        _ = fi55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
