package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{-9223372036854775808, -2147483648, 2147483647, -9223372036854775808}
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 int8 = 0
    var v5 float32 = float32(0.1)
    if len(v1) > 0 {
        v1[0] = 848
    }
    fmt.Println(v4)
    v4 = 127
    if false {
        fmt.Println(v5, v3)
        v3 = v3
        switch "hello" {
        case "hello":
            v4 += (v4 / ((v4) | 1))
        }
    } else {
        if v3 {
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 int64 = 0
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch ("" + ("a" + "")) {
                    case "a":
                        var v9 float32 = v5
                        _ = v9
                    case "":
                        v7 = v7
                    }
                    v2[2147483648] = 65535
                    _ = len(v2)
                    _ = fi8
                }
                _ = fi6
                _ = v7
            }
        } else {
            v4 -= (-(v4))
            v10 := GenS1{F0: 426, F1: (-(v4)), F2: 1}
            _ = v10
        }
        _ = len(v2)
        v11 := GenS1{F0: 701, F1: (v4 + -128), F2: 9223372036854775807}
        _ = v11
    }
    v3 = v3
    if ((v4 - v4) == (-(v4))) {
        if v3 {
            switch 1 {
            case 853:
                if v3 {
                    var v12 int64 = -9223372036854775808
                    _ = v12
                }
            case 128:
                v13 := GenS1{F0: 141, F1: (v4 << ((583) & 63)), F2: 1}
                _ = v13
            case 32767:
                var v14 string = "\n\t"
                _ = v14
            default:
                fmt.Println(v5)
            }
            switch "a" {
            case "foo bar":
                for fi15 := 0; fi15 < 20; fi15++ {
                    v1 = append(v1, (fi15 & 0))
                    _ = fi15
                }
            case "":
                var v16 uint = 42
                _ = v16
            }
            switch ("a" + "a") {
            case "foo bar":
                var v17 uint16 = 32767
                _ = v17
            case "a":
                var v18 uint = 42
                _ = v18
            case "":
                v5 += v5
            }
        } else {
            fmt.Println(v4, v4, v3)
            if (605 <= 42) {
                v19 := GenS2{F0: v3}
                _ = v19
            }
            var v20 int32 = 0
            _ = v20
        }
    } else {
        var v21 string = "hello"
        v22 := GenS2{F0: (0 >= 902)}
        _ = v21
        _ = v22
    }
    v3 = (42 == 1)
    v5 += 0.1
    if false {
        var v23 uint8 = 255
        if true {
            v24 := GenS3{F0: 255, F1: 255}
            if (v4 < v4) {
                v3 = (v3 && false)
                fmt.Println(v5, v5, v23)
                v4 += (v4 >> ((1) & 63))
            }
            _ = v24
        } else {
            v4 -= (v4 + v4)
            var v25 int16 = 32767
            var v26 uint16 = 0
            _ = v25
            _ = v26
        }
        switch -2147483648 {
        case -2147483648:
            v5 += v5
        default:
            switch -510 {
            case 350:
                for fi27 := 0; fi27 < 5; fi27++ {
                    v5 += ((v5 * (3.14159 / 3.14159)) + v5)
                    var v28 float64 = 0.0
                    _ = fi27
                    _ = v28
                }
            case -1:
                for fi29 := 0; fi29 < 8; fi29++ {
                    v23 = 42
                    v23 += v23
                    _ = fi29
                }
            }
        }
        _ = v23
    }
    switch "\n\t" {
    case "foo bar":
        if (128 == 255) {
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS2{F0: (32767 == 1)}
                _ = fi30
                _ = v31
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                if false {
                    var v32 uint = 746
                    _ = v32
                }
            case "hello":
                var v33 float64 = 0.0
                _ = v33
            default:
                if true {
                    var v34 uint8 = 0
                    fmt.Println(v34, v5, v3)
                    _ = v34
                } else {
                    v3 = (true || v3)
                    switch 42 {
                    case 9223372036854775807:
                        v4 += (v4 - 100)
                    }
                }
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                v5 = v5
                v4 = v4
                _ = fi35
            }
            var v36 float64 = 0.0
            _ = v36
        }
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if v3 {
        fmt.Println(v5, v3, v3)
        v37 := GenS1{F0: 65535, F1: (^(v4)), F2: 42}
        _ = v37
    }
    v1 = append(v1, 42)
    v38 := GenS1{F0: 42, F1: (v4 - -28), F2: 1}
    if (true || (!v3)) {
        v39 := GenS4{F0: 42}
        v4 += 100
        _ = v39
    } else {
        switch 1 {
        case 32767:
            fmt.Println(v4)
        case 1:
            if true {
                if ((44 > 0) && false) {
                    v3 = v3
                } else {
                    if false {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v5 -= v5
                            _ = fi40
                        }
                        v41 := GenS4{F0: 825}
                        v42 := GenS3{F0: 0, F1: 2}
                        _ = v41
                        _ = v42
                    } else {
                        delete(v2, 987)
                    }
                    var v43 string = "a"
                    var v44 string = ("" + "\n\t")
                    _ = v43
                    _ = v44
                }
                v5 += v5
            } else {
                var v45 uint16 = 2
                _ = v45
            }
        case -9223372036854775808:
            switch -1 {
            case 128:
                var v46 uint16 = 2
                _ = v46
            case 0:
                var v47 int = 32767
                _ = v47
            default:
                for fi48 := 0; fi48 < 8; fi48++ {
                    var v49 bool = (!true)
                    var v50 bool = (874 == 614)
                    switch (fi48 & int(v4)) {
                    case 127:
                        delete(v2, 1)
                    default:
                        _ = len(v1)
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            }
        }
        fmt.Println(v5, v4)
        if (!false) {
            if true {
                var v51 uint32 = 2147483648
                if (32767 != 32767) {
                    var v52 bool = v3
                    var v53 int16 = 420
                    _ = v52
                    _ = v53
                } else {
                    var v54 string = ("" + "\n\t")
                    v5 = float32(3.14159)
                    _ = v54
                }
                _ = v51
            } else {
                var v55 bool = (!false)
                v56 := GenS1{F0: 201, F1: -128, F2: -9223372036854775808}
                _ = v55
                _ = v56
            }
            if v3 {
                var v57 string = "hello"
                v58 := GenS4{F0: 2}
                v57 = "\n\t"
                _ = v57
                _ = v58
            }
            v3 = (!v3)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
