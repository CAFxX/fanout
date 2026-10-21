package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint; F1 int32; F2 int8 }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 string = ("a" + "foo bar")
    var v5 int = 42
    var v6 float32 = ((((1e10 - (-2.25 - ((1e10 - float32(float64(float32(3.14159)))) * 1e10))) / (0.1 - (0.0 * ((0.0 - -2.25) * (0.1 + 1.5))))) - 0.0) - 1.5)
    v7 := []int{v5, v5, 0}
    v5 = v5
    fmt.Println(v4, v6)
    switch (("" + "\n\t") + v4) {
    case "":
        switch v5 {
        case 42:
            var v8 uint64 = 1
            _ = v8
        case 397:
            v9 := GenS1{F0: (v5 ^ v5), F1: 4294967295}
            _ = v9
        case 32768:
            _ = len(v7)
        default:
            _ = len(v2)
        }
    }
    var v10 bool = true
    if (3.14159 != 3.14159) {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS2{F0: 128, F1: 438}
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint64 = 1
                fmt.Println(fi13)
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        switch (v5 >> ((0) & 63)) {
        case 0:
            if v1 {
                v15 := GenS1{F0: v5, F1: 3}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 19; fi16++ {
                    if (!(1 == 18446744073709551615)) {
                        v17 := GenS2{F0: fi16, F1: 0}
                        switch (((v4 + v4) + ((((v4 + v4) + v4) + (("a" + "hello") + v4)) + ((v4 + "a") + v4))) + "\n\t") {
                        case "a":
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v1 = v10
                                switch "foo bar" {
                                case "\n\t":
                                    if v10 {
                                        fmt.Println(v1, v1)
                                        if (0 != 190) {
                                            if len(v7) > 0 {
                                                v7[0] = -33
                                            }
                                            if (65535 > 747) {
                                                v5 += (fi16 & 32767)
                                                if (1 == 41) {
                                                    v1 = v10
                                                } else {
                                                    var v19 int8 = -2
                                                    v10 = true
                                                    var v20 uint8 = 176
                                                    _ = v19
                                                    _ = v20
                                                }
                                                fmt.Println(v4, fi16, v1)
                                            }
                                        }
                                    }
                                case "":
                                    v21 := GenS2{F0: (fi18 + 635), F1: 255}
                                    _ = v21
                                case "a":
                                    if v10 {
                                        v22 := GenS3{F0: 394, F1: 18446744073709551615, F2: (fi16 << ((255) & 63))}
                                        v23 := GenS1{F0: fi16, F1: 42}
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 bool = v10
                                    _ = v24
                                default:
                                    var v25 uint = 0
                                    _ = v25
                                }
                                v26 := GenS2{F0: (fi18 - (fi16 * fi16)), F1: 0}
                                _ = fi18
                                _ = v26
                            }
                        case "\n\t":
                            var v27 float64 = (-2.25 + -2.25)
                            _ = v27
                        case "hello":
                            fmt.Println(v4)
                        }
                        _ = v17
                    }
                    _ = fi16
                }
                fmt.Println(v10, v5)
            }
        }
        if (false && true) {
            v5 = v5
            var v28 uint64 = 4294967295
            _ = v28
        } else {
            v29 := GenS4{F0: 1, F1: 588, F2: ("hello" + (((v4 + ("foo bar" + v4)) + "foo bar") + "foo bar"))}
            _ = v29
        }
    }
    fmt.Println(v6, v5)
    switch v5 {
    case 1:
        v5 += (v5 - v5)
    }
    v30 := GenS1{F0: 67, F1: 65535}
    if true {
        if false {
            var v31 uint8 = 1
            v32 := GenS5{F0: 42, F1: 0}
            _ = v31
            _ = v32
        }
        if len(v7) > 0 {
            v7[0] = v5
        }
        v33 := GenS6{F0: 4294967295, F1: 1, F2: 100}
        _ = v33
    } else {
        var v34 int16 = 0
        _ = v34
    }
    if v10 {
        var v35 int16 = -1
        v36 := GenS7{F0: 1}
        var v37 int16 = v35
        _ = v35
        _ = v36
        _ = v37
    }
    var v38 uint32 = 42
    v39 := GenS3{F0: 4294967295, F1: 18446744073709551615, F2: v5}
    for fi40 := 0; fi40 < 2; fi40++ {
        switch v5 {
        case 32768:
            var v41 float32 = -2.25
            _ = v41
        case 296:
            if v1 {
                if (!v10) {
                    var v42 uint16 = 32767
                    for fi43 := 0; fi43 < 7; fi43++ {
                        var v44 int64 = 9223372036854775807
                        _ = fi43
                        _ = v44
                    }
                    _ = v42
                } else {
                    v5 = fi40
                }
                for fi45 := 0; fi45 < 19; fi45++ {
                    fmt.Println(v5)
                    fmt.Println(v6, v10)
                    _ = len(v7)
                    _ = fi45
                }
            }
        case 87:
            switch fi40 {
            case 128:
                switch v4 {
                case "foo bar":
                    var v46 uint32 = 2147483647
                    _ = v46
                default:
                    var v47 int8 = 91
                    _ = v47
                }
            case 0:
                v1 = ((("a" + v4) <= ("" + v4)) && (1.5 >= 0.0))
            case 213:
                v10 = true
            default:
                var v48 int32 = 42
                _ = v48
            }
        }
        if false {
            v10 = (2 <= 2)
            for fi49 := 0; fi49 < 6; fi49++ {
                v50 := GenS6{F0: 42, F1: -2147483648, F2: 1}
                var v51 float32 = v6
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
