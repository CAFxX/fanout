package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 uint16; F2 int32 }

func main() {
    var v1 int32 = 1
    var v2 int64 = -9223372036854775808
    var v3 bool = (true && ("hello" < ("\n\t" + ((("" + "foo bar") + "") + "\n\t"))))
    var v4 uint8 = 204
    var v5 uint64 = 42
    var v6 float64 = float64(1e10)
    v4 = (^(v4))
    fmt.Println(v4, v5)
    var v7 uint32 = 0
    if ("hello" == "") {
        if true {
            v5 = 1
        }
        switch "" {
        case "hello":
            switch "" {
            case "":
                v8 := []int{2147483647, 127, 0}
                _ = v8
            case "\n\t":
                var v9 uint8 = v4
                _ = v9
            default:
                v7 -= (^(v7))
            }
        case "":
            v10 := []int{32767}
            _ = v10
        default:
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v3, v5)
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := []int{(fi11 / ((fi11) | 1))}
                    _ = fi12
                    _ = v13
                }
                var v14 bool = (fi11 == fi11)
                _ = fi11
                _ = v14
            }
        }
        if v3 {
            if v3 {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 int16 = -848
                    var v17 float32 = (3.14159 - (float32(v6) * 3.14159))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 9; fi18++ {
                    v5 += (^(v5))
                    _ = fi18
                }
            } else {
                if true {
                    v19 := []int{}
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 float64 = (0.1 - 0.0)
                        var v22 int8 = 100
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    fmt.Println(v7)
                    _ = v19
                }
                fmt.Println(v7, v2, v5)
                if (v1 == (v1 & v1)) {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v24 := []int{(^(fi23)), int(v2), 299}
                        _ = fi23
                        _ = v24
                    }
                }
            }
            v6 = float64(float32((v6 + v6)))
        }
    }
    var v25 float32 = 3.14159
    v26 := []int{}
    if (!v3) {
        var v27 string = ("a" + "foo bar")
        for fi28 := 0; fi28 < 3; fi28++ {
            switch 645 {
            case 32768:
                for fi29 := 0; fi29 < 16; fi29++ {
                    fmt.Println(v3)
                    var v30 float64 = v6
                    _ = fi29
                    _ = v30
                }
            case 403:
                fmt.Println(v7)
            case 88:
                v6 -= v6
            }
            var v31 bool = (42 != 255)
            for fi32 := 0; fi32 < 7; fi32++ {
                v2 -= v2
                _ = fi32
            }
            _ = fi28
            _ = v31
        }
        _ = v27
    }
    var v33 float32 = float32(((0.0 - float64((v25 * (v25 + (v25 / 3.14159))))) / float64(v25)))
    switch 128 {
    case -1:
        fmt.Println(v2, v3, v33)
    case 301:
        if len(v26) > 0 {
            v26[0] = 32768
        }
    }
    if ((v1 / (((^(v1))) | 1)) > v1) {
        for fi34 := 0; fi34 < 15; fi34++ {
            v26 = append(v26, 531)
            fmt.Println(v1)
            switch ("\n\t" + "hello") {
            case "hello":
                v35 := GenS1{F0: v4, F1: v4, F2: 1}
                _ = v35
            case "\n\t":
                if v3 {
                    if v3 {
                        switch int(v7) {
                        case -9223372036854775808:
                            fmt.Println(v1, v1)
                        case -220:
                            v26 = append(v26, 32767)
                        case -35:
                            switch ("\n\t" + "foo bar") {
                            case "hello":
                                fmt.Println(v6)
                            default:
                                v26 = append(v26, 127)
                            }
                        }
                        switch "a" {
                        case "":
                            var v36 int32 = v1
                            _ = v36
                        }
                        v37 := GenS2{F0: 1, F1: (v5 >> ((18446744073709551615) & 63))}
                        _ = v37
                    } else {
                        var v38 float32 = (v33 + v25)
                        _ = v38
                    }
                    v39 := GenS3{F0: 827, F1: v2}
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v2 = v2
                        _ = fi40
                    }
                }
            }
            _ = fi34
        }
        switch "foo bar" {
        case "foo bar":
            v7 = (^(v7))
        case "a":
            for fi41 := 0; fi41 < 4; fi41++ {
                if v3 {
                    v26 = append(v26, fi41)
                    _ = len(v26)
                } else {
                    v7 += (v7 ^ uint32(v4))
                    var v42 uint64 = uint64(v4)
                    var v43 bool = false
                    _ = v42
                    _ = v43
                }
                _ = fi41
            }
        default:
            var v44 int64 = (v2 - 1)
            _ = v44
        }
    } else {
        if v3 {
            switch 2147483647 {
            case 545:
                v45 := GenS4{F0: (32768 > 32768), F1: uint16(v5), F2: v1}
                _ = v45
            case 2147483647:
                var v46 int = 974
                _ = v46
            }
            if (-32768 <= 488) {
                v26 = append(v26, -201)
                var v47 float32 = float32(v6)
                _ = v47
            }
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                v33 -= (v33 - v33)
            }
        } else {
            switch ("hello" + "hello") {
            case "\n\t":
                var v48 int16 = int16(v1)
                _ = v48
            case "":
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v26 = append(v26, 127)
                case "a":
                    fmt.Println(v3)
                }
            }
        }
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    v2 += v2
    for fi49 := 0; fi49 < 6; fi49++ {
        v5 = 2
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
