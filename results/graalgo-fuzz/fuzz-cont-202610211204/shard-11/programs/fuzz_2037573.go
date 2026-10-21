package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 int64; F2 int64 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint64 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]uint32)
    var v3 uint64 = (v1 | v1)
    v4 := make(map[string]uint32)
    var v5 uint16 = 255
    v6 := GenS1{F0: 457, F1: (-2.25 != ((3.14159 + 1.5) + ((-2.25 / 1e10) - (float32(float64((float32(float64((0.1 + 3.14159))) + 0.1))) / 1.0)))), F2: 32768}
    if ("foo bar" == "a") {
        var v7 int64 = 9223372036854775807
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 bool = (67 <= 0)
            switch fi8 {
            case -518:
                v1 = v1
            case 4294967295:
                fmt.Println(v5)
            case 1:
                var v10 float32 = 0.1
                _ = v10
            default:
                delete(v4, "")
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v11 := GenS2{F0: (v5 * (^(v5)))}
    var v12 int64 = int64(v3)
    var v13 bool = true
    v1 += v1
    v14 := GenS3{F0: 32768}
    v1 -= 0
    v3 = v1
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint16 = 65535
        _ = fi15
        _ = v16
    }
    if (168 >= (v3 & (v1 << ((0) & 63)))) {
        if false {
            v4[""] = 65535
            switch ("\n\t" + "") {
            case "a":
                v17 := GenS4{F0: 42}
                _ = v17
            case "":
                if (0 > 42) {
                    var v18 uint16 = v5
                    _ = v18
                }
            case "foo bar":
                var v19 uint16 = (v5 | (^(v5)))
                _ = v19
            }
        } else {
            v20 := GenS1{F0: int16(v12), F1: ("\n\t" > "a"), F2: 262}
            _ = v20
        }
        if v13 {
            fmt.Println(v12, v5)
        }
        if ("a" != "\n\t") {
            v21 := v2[1]
            var v22 int = 0
            _ = v21
            _ = v22
        }
    } else {
        var v23 bool = (((849 >= -32768) && (-128 <= 42)) || (!false))
        _ = v23
    }
    fmt.Println(v5, v12, v1)
    var v24 float32 = 0.0
    switch 127 {
    case 127:
        if (!(1000 > -1)) {
            v25 := GenS5{F0: -9223372036854775808}
            var v26 uint8 = 207
            var v27 uint32 = 104
            _ = v25
            _ = v26
            _ = v27
        } else {
            v24 += v24
            v24 -= ((1e10 + 1.5) - v24)
        }
    case 32767:
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "a" {
            case "a":
                for fi29 := 0; fi29 < 2; fi29++ {
                    switch (fi28 * (^(fi29))) {
                    case 2147483647:
                        v30 := GenS5{F0: (fi29 - (fi28 + 2147483647))}
                        _ = v30
                    default:
                        switch "hello" {
                        case "a":
                            v31 := GenS6{F0: 255, F1: -9223372036854775808, F2: (v12 ^ v12)}
                            _ = v31
                        default:
                            v13 = (606 == 65535)
                        }
                        if (0 != 0) {
                            switch "a" {
                            case "a":
                                switch ("hello" + ("hello" + "a")) {
                                case "\n\t":
                                    var v32 bool = v13
                                    _ = v32
                                }
                            case "foo bar":
                                var v33 uint8 = 0
                                _ = v33
                            default:
                                var v34 uint16 = 2
                                _ = v34
                            }
                            v24 = float32((3.14159 * 3.14159))
                            if v13 {
                                var v35 uint32 = 0
                                _ = v35
                            } else {
                                fmt.Println(v13, v12)
                                if (2147483647 != -2147483648) {
                                    v36 := GenS1{F0: -32768, F1: false, F2: fi28}
                                    _ = v36
                                } else {
                                    var v37 uint8 = 255
                                    var v38 int64 = (v12 + (v12 / ((1) | 1)))
                                    v24 = v24
                                    _ = v37
                                    _ = v38
                                }
                            }
                        } else {
                            var v39 float64 = (-2.25 * ((3.14159 - (float64(v24) * float64(0.0))) + 0.1))
                            v3 -= 4294967295
                            _ = v39
                        }
                        var v40 uint64 = v3
                        _ = v40
                    }
                    var v41 string = "hello"
                    _ = fi29
                    _ = v41
                }
            case "\n\t":
                if (2 != 1) {
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        var v42 bool = v13
                        _ = v42
                    case "hello":
                        v43 := GenS7{F0: 2147483647, F1: 473}
                        _ = v43
                    default:
                        delete(v4, "hello")
                    }
                    for fi44 := 0; fi44 < 3; fi44++ {
                        switch (^(fi28)) {
                        case 9223372036854775807:
                            v3 += (v1 << ((0) & 63))
                        case -682:
                            v45 := GenS3{F0: (fi28 ^ (fi44 % (((fi44 >> ((255) & 63))) | 1)))}
                            _ = v45
                        case -9223372036854775808:
                            v46 := GenS8{F0: (v12 / ((v12) | 1)), F1: 867}
                            _ = v46
                        }
                        if ((v12 << ((685) & 63)) < v12) {
                            fmt.Println(fi28, v12)
                            var v47 int32 = int32(fi28)
                            var v48 uint = 710
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS9{F0: 100, F1: v1}
                            v5 = (v5 & (v5 + 255))
                            _ = v49
                        }
                        _ = fi44
                    }
                } else {
                    v12 = v12
                    v1 += v1
                    if false {
                        var v50 bool = v13
                        for fi51 := 0; fi51 < 6; fi51++ {
                            v2[fi51] = 279
                            v1 -= 0
                            _ = fi51
                        }
                        _ = v50
                    }
                }
            }
            _ = fi28
        }
    case 9223372036854775807:
        v52 := []int{32767}
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
