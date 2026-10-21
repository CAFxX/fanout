package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int16 = -1
    var v2 string = "a"
    var v3 float32 = (0.0 - -2.25)
    v4 := make(map[string]bool)
    v5 := []uint32{}
    var v6 int64 = 9223372036854775807
    var v7 int64 = v6
    switch v2 {
    case "hello":
        v2 = ""
    case "":
        v8 := GenS1{F0: 11, F1: -1, F2: 605}
        _ = v8
    default:
        var v9 uint8 = 192
        _ = v9
    }
    v10 := GenS2{F0: false, F1: 0}
    v6 -= (v6 ^ v6)
    v1 = v1
    fmt.Println(v7, v7)
    if false {
        v2 = "\n\t"
    } else {
        var v11 float64 = 0.1
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                v2 = ((v2 + "hello") + (("hello" + v2) + (v2 + "foo bar")))
                if (v7 >= (v6 | 1)) {
                    switch "foo bar" {
                    case "\n\t":
                        v5 = append(v5, 2147483647)
                    default:
                        for fi13 := 0; fi13 < 7; fi13++ {
                            fmt.Println(v11)
                            if len(v5) > 0 {
                                v5[0] = 4294967295
                            }
                            _ = fi13
                        }
                    }
                    var v14 string = v2
                    _ = v14
                }
            } else {
                if (65535 != 42) {
                    var v15 string = v2
                    _ = v15
                } else {
                    var v16 uint16 = 68
                    _ = v16
                }
                var v17 int = fi12
                _ = v17
            }
            v7 = (v7 >> ((4294967295) & 63))
            if (!false) {
                switch "a" {
                case "a":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        if false {
                            var v19 uint64 = 2
                            _ = len(v5)
                            _ = v19
                        }
                        v20 := GenS3{F0: ("hello" + (((("a" + v2) + v2) + (v2 + v2)) + v2)), F1: 0}
                        _ = fi18
                        _ = v20
                    }
                case "":
                    if true {
                        var v21 int8 = int8(v6)
                        _ = v21
                    }
                case "\n\t":
                    v7 += v6
                }
                for fi22 := 0; fi22 < 18; fi22++ {
                    fmt.Println(fi12, fi22)
                    _ = fi22
                }
            }
            _ = fi12
        }
        v1 = v1
        _ = v11
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if false {
            v24 := GenS1{F0: 100, F1: -19, F2: 255}
            _ = v24
        } else {
            switch (fi23 & (fi23 | fi23)) {
            case 127:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            fmt.Println(v2, v1, v7)
            fmt.Println(v3, v6, fi23)
        }
        if true {
            switch fi23 {
            case 128:
                switch (fi23 >> ((276) & 63)) {
                case 127:
                    v25 := GenS3{F0: v2, F1: 405}
                    _ = v25
                case -1:
                    v26 := GenS3{F0: v2, F1: 142}
                    _ = v26
                }
            case 4294967295:
                var v27 uint8 = 81
                _ = v27
            case 32767:
                var v28 uint8 = 206
                _ = v28
            default:
                var v29 uint = 65535
                _ = v29
            }
            var v30 string = v2
            _ = len(v5)
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 uint32 = 2147483647
            var v33 int32 = 1
            v34 := GenS4{F0: int8(v32), F1: (v33 ^ 276)}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi23
    }
    var v35 int64 = v6
    var v36 string = (("foo bar" + v2) + ((v2 + v2) + v2))
    v6 = 9223372036854775807
    _ = len(v4)
    v37 := GenS5{F0: 1, F1: 18446744073709551615}
    switch int(v35) {
    case -513:
        if (!true) {
            for fi38 := 0; fi38 < 6; fi38++ {
                v35 += int64(v1)
                _ = fi38
            }
        } else {
            if (!((false || (v3 <= 0.0)) || (!(v3 < (v3 - 1e10))))) {
                for fi39 := 0; fi39 < 5; fi39++ {
                    var v40 string = v36
                    _ = fi39
                    _ = v40
                }
                if (!true) {
                    v3 -= v3
                    if (-425 > v1) {
                        v3 += -2.25
                        if false {
                            var v41 float64 = 0.0
                            var v42 int = -2147483648
                            var v43 float64 = ((1.5 * (v41 * (v41 - v41))) + 1e10)
                            _ = v41
                            _ = v42
                            _ = v43
                        } else {
                            _ = len(v5)
                            _ = len(v5)
                            if false {
                                v1 = v1
                                var v44 float64 = 0.0
                                var v45 uint8 = 152
                                _ = v44
                                _ = v45
                            }
                        }
                        _ = len(v4)
                    }
                } else {
                    var v46 int = -2147483648
                    v1 = 645
                    _ = v46
                }
            }
        }
    default:
        _ = len(v5)
    }
    var v47 uint32 = 496
    var v48 uint8 = 0
    if ((1 <= 255) || ((!(!(!false))) && (1 > 831))) {
        v6 = int64(v48)
    } else {
        v49 := GenS3{F0: (v36 + "foo bar"), F1: uint64(v47)}
        delete(v4, v36)
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
