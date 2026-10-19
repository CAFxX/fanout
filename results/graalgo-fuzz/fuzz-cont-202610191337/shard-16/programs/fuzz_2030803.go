package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint64 }

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]int)
    var v3 int8 = -73
    var v4 float32 = (0.0 - -2.25)
    var v5 int64 = -9223372036854775808
    var v6 int64 = (v5 ^ (v5 - (v5 + (v5 + 0))))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = int8(fi7)
        if (!true) {
            var v9 int32 = 402
            var v10 int64 = (v5 / ((-79) | 1))
            _ = v10
            _ = v9
        } else {
            delete(v2, 1)
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: -768, F1: 1000, F2: 521}
    v2[480] = 42
    fmt.Println(v4, v5, v3)
    var v12 uint16 = 255
    v5 = 9223372036854775807
    for fi13 := 0; fi13 < 15; fi13++ {
        if ((10 != 297) || false) {
            if true {
                fmt.Println(v6, v4)
                fmt.Println(fi13, v4, v4)
            }
        } else {
            switch "a" {
            case "a":
                v14 := v2[42]
                _ = v14
            case "\n\t":
                v1 = append(v1, "a")
            case "":
                switch ("hello" + "hello") {
                case "\n\t":
                    var v15 uint16 = (v12 - (v12 | v12))
                    _ = v15
                case "":
                    v12 += v12
                }
            }
            var v16 int = 1
            _ = v16
        }
        _ = fi13
    }
    var v17 uint = uint(v12)
    v18 := GenS1{F0: 32, F1: -878, F2: 42}
    for fi19 := 0; fi19 < 16; fi19++ {
        if (!false) {
            if false {
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint = (v17 * (v17 + v17))
                    _ = fi20
                    _ = v21
                }
                if (183 < -32768) {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v17 += uint(v12)
                        _ = fi22
                    }
                }
                v17 += (v17 / ((v17) | 1))
            }
            v4 = (v4 + 1e10)
        }
        v17 -= (^(v17))
        if len(v1) > 0 {
            v1[0] = ("hello" + ("hello" + "foo bar"))
        }
        _ = fi19
    }
    var v23 uint8 = 127
    var v24 int8 = v3
    if true {
        v12 += 65535
    } else {
        switch "\n\t" {
        case "a":
            v12 = (^(v12))
        case "foo bar":
            if (v17 < v17) {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint = (v17 / ((v17) | 1))
                    fmt.Println(v24)
                    _ = fi25
                    _ = v26
                }
                var v27 int32 = 133
                if true {
                    if true {
                        var v28 int16 = 1
                        _ = v28
                    }
                } else {
                    switch "a" {
                    case "":
                        if (139 != 0) {
                            if (v3 != v24) {
                                for fi29 := 0; fi29 < 19; fi29++ {
                                    if true {
                                        delete(v2, 4294967295)
                                        v30 := GenS2{F0: (v27 & -1)}
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 15; fi31++ {
                                        var v32 int8 = (v24 << ((0) & 63))
                                        _ = fi31
                                        _ = v32
                                    }
                                    v4 += -2.25
                                    _ = fi29
                                }
                                v4 -= v4
                            }
                        } else {
                            if true {
                                if false {
                                    v33 := GenS3{F0: (v23 & 42)}
                                    _ = v33
                                } else {
                                    if (!(!false)) {
                                        var v34 float64 = 3.14159
                                        _ = v34
                                    } else {
                                        var v35 uint = v17
                                        if false {
                                            var v36 string = "a"
                                            _ = len(v2)
                                            var v37 float32 = (0.0 - (0.1 / float32(3.14159)))
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if (true && (false || ((v4 + 1e10) > (v4 + 0.1)))) {
                                                if true {
                                                    if true {
                                                        for fi38 := 0; fi38 < 17; fi38++ {
                                                            for fi39 := 0; fi39 < 19; fi39++ {
                                                                var v40 float64 = -2.25
                                                                var v41 uint = (v17 % ((uint(v23)) | 1))
                                                                _ = fi39
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                            var v42 float32 = -2.25
                                                            _ = fi38
                                                            _ = v42
                                                        }
                                                        v1 = append(v1, "")
                                                        v43 := GenS4{F0: 65535}
                                                        _ = v43
                                                    }
                                                    fmt.Println(v3, v6, v24)
                                                } else {
                                                    var v44 int64 = (v6 % ((int64(v3)) | 1))
                                                    _ = v44
                                                }
                                                if (((v4 - ((v4 - v4) + float32((3.14159 + 3.14159)))) + float32(0.0)) == 0.0) {
                                                    for fi45 := 0; fi45 < 12; fi45++ {
                                                        v4 = v4
                                                        _ = fi45
                                                    }
                                                    fmt.Println(v17)
                                                } else {
                                                    var v46 bool = false
                                                    _ = v46
                                                }
                                            }
                                        }
                                        var v47 int16 = 1
                                        _ = v35
                                        _ = v47
                                    }
                                }
                            }
                        }
                    case "a":
                        if false {
                            switch "foo bar" {
                            case "":
                                switch -1 {
                                case -9223372036854775808:
                                    var v48 int32 = (-(v27))
                                    _ = v48
                                }
                            default:
                                v49 := GenS5{F0: -32768, F1: 2, F2: 42}
                                _ = v49
                            }
                            for fi50 := 0; fi50 < 11; fi50++ {
                                v17 += (v17 << ((v17) & 63))
                                for fi51 := 0; fi51 < 5; fi51++ {
                                    v24 -= int8(v23)
                                    _ = fi51
                                }
                                fmt.Println(v6, v27)
                                _ = fi50
                            }
                        } else {
                            if false {
                                fmt.Println(v23, v23)
                            }
                        }
                    default:
                        if true {
                            fmt.Println(v6, v6, v4)
                        }
                    }
                    var v52 int = 2147483647
                    _ = v52
                }
                _ = v27
            }
        }
        delete(v2, 42)
    }
    var v53 uint32 = 389
    var v54 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
