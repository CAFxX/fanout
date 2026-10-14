package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string }

func main() {
    var v1 float64 = ((float64(1e10) - 1.5) - 3.14159)
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    switch (v2 + v2) {
    case "\n\t":
        v1 = (((v1 * 0.0) / 1.0) * (float64(0.0) - 3.14159))
    }
    if true {
        v4 := []int{-539, 0, 42, -2147483648}
        switch "foo bar" {
        case "\n\t":
            var v5 uint8 = 255
            _ = v5
        case "a":
            var v6 bool = (false || false)
            _ = v6
        default:
            fmt.Println(v2, v2)
        }
        _ = v4
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    var v9 float64 = 1e10
    v10 := []int{799, 1}
    if (128 <= 42) {
        switch 128 {
        case 0:
            if false {
                v1 = (-2.25 - v9)
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := GenS1{F0: 32767}
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v2 = ("\n\t" + (v2 + v2))
                    fmt.Println(v9, v2, v2)
                    _ = fi13
                }
            }
        case -1:
            switch 128 {
            case 977:
                var v14 string = (v2 + "\n\t")
                _ = v14
            case 42:
                v9 += 1e10
            }
        case 128:
            var v15 uint16 = 32767
            _ = v15
        default:
            if false {
                v9 = (0.1 - (v1 - v1))
                if false {
                    if len(v10) > 0 {
                        v10[0] = -9223372036854775808
                    }
                    v9 = v1
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch v2 {
                        case "a":
                            v2 = "foo bar"
                        case "":
                            v17 := GenS1{F0: 1000}
                            _ = v17
                        case "foo bar":
                            for fi18 := 0; fi18 < 11; fi18++ {
                                if (114 == -128) {
                                    var v19 uint32 = 2147483648
                                    v10 = append(v10, int(v19))
                                    var v20 uint8 = 45
                                    _ = v19
                                    _ = v20
                                }
                                var v21 int32 = 42
                                _ = fi18
                                _ = v21
                            }
                        default:
                            var v22 int32 = -2147483648
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 9; fi23++ {
                            _ = len(v3)
                            if true {
                                switch "\n\t" {
                                case "a":
                                    v24 := GenS2{F0: -1, F1: 255, F2: 42}
                                    _ = v24
                                case "\n\t":
                                    if ((91 < 175) || (fi16 >= (fi16 / (((fi16 & fi16)) | 1)))) {
                                        var v25 float64 = (0.1 * 0.0)
                                        var v26 int8 = -128
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v2 = (v2 + v2)
                                        v2 = v2
                                        v27 := v3[2147483647]
                                        _ = v27
                                    }
                                }
                                v10 = append(v10, fi16)
                            } else {
                                if (v2 < "hello") {
                                    v3[2] = false
                                } else {
                                    v9 += v1
                                    if true {
                                        var v28 float64 = ((v1 - v9) * 1.5)
                                        v9 += float64(1e10)
                                        _ = v28
                                    }
                                }
                                v29 := GenS2{F0: 127, F1: 0, F2: 118}
                                _ = v29
                            }
                            _ = fi23
                        }
                        v30 := GenS3{F0: 100}
                        _ = fi16
                        _ = v30
                    }
                }
            } else {
                var v31 int16 = 1000
                if true {
                    var v32 int8 = 42
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v31 += 1000
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v34 := GenS3{F0: 12}
                        var v35 int8 = -93
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    var v36 float32 = 0.1
                    _ = v36
                }
                v31 -= (v31 / (((v31 % ((v31) | 1))) | 1))
                _ = v31
            }
        }
        v37 := GenS3{F0: 47}
        v9 += 1e10
        _ = v37
    }
    v38 := GenS4{F0: (v2 + v2), F1: 4294967295, F2: 337}
    v2 = (v2 + "\n\t")
    v39 := GenS5{F0: "a"}
    if (!((-2.25 + (1e10 / -2.25)) == 3.14159)) {
        delete(v3, 42)
        if false {
            var v40 uint16 = 444
            _ = v40
        } else {
            for fi41 := 0; fi41 < 16; fi41++ {
                if (!true) {
                    var v42 uint32 = 0
                    var v43 uint32 = (^(v42))
                    _ = v42
                    _ = v43
                } else {
                    v2 = (v2 + (v2 + (("\n\t" + "hello") + "hello")))
                    v1 += (v1 * v1)
                }
                var v44 uint16 = 255
                _ = fi41
                _ = v44
            }
            switch 42 {
            case -940:
                fmt.Println(v1)
            case 32768:
                var v45 float64 = v9
                _ = v45
            default:
                switch 32768 {
                case 2147483647:
                    fmt.Println(v1)
                case 4294967295:
                    v2 = (v2 + "")
                default:
                    var v46 int64 = -203
                    _ = v46
                }
            }
        }
        fmt.Println(v2, v2)
    }
    var v47 float32 = (1.5 - (1.5 - float32(float64((-2.25 * 0.1)))))
    var v48 int = -1
    v49 := v3[771]
    for fi50 := 0; fi50 < 13; fi50++ {
        for fi51 := 0; fi51 < 12; fi51++ {
            v49 = true
            v2 = "foo bar"
            fmt.Println(v47)
            _ = fi51
        }
        v47 += 1e10
        switch v2 {
        case "\n\t":
            var v52 int8 = 1
            _ = v52
        case "hello":
            for fi53 := 0; fi53 < 19; fi53++ {
                var v54 uint = 332
                fmt.Println(v54, fi53)
                v55 := GenS5{F0: v2}
                var v56 int32 = -2147483648
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
        case "a":
            fmt.Println(v2, v47)
        default:
            var v57 bool = v49
            _ = v57
        }
        _ = fi50
    }
    var v58 int = (v48 << ((255) & 63))
    v2 = v2
    fmt.Println(v58)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v58:", v58)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v47
    _ = v48
    _ = v49
    _ = v58
    _ = v9
}
