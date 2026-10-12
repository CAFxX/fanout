package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 - 3.14159)
    var v3 string = "a"
    var v4 int32 = 200
    v5 := make(map[string]int)
    if false {
        if (-159 >= 42) {
            _ = len(v1)
            for fi6 := 0; fi6 < 20; fi6++ {
                v7 := v5["foo bar"]
                var v8 bool = (!false)
                var v9 uint16 = 346
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 string = v3
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if (float32((0.1 - 3.14159)) != 0.1) {
            v12 := []int{(fi11 + (^(fi11))), 2147483647, -484, (-(fi11))}
            fmt.Println(v2)
            _ = v12
        } else {
            fmt.Println(v3, v3)
        }
        v2 += -2.25
        v4 = v4
        _ = fi11
    }
    var v13 uint16 = 32767
    if true {
        v13 += v13
        v14 := GenS1{F0: -1}
        v15 := []int{1, -634, 127}
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: 1000}
        if (false && true) {
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v3, v2)
                    _ = fi18
                }
                _ = fi17
            }
        } else {
            for fi19 := 0; fi19 < 20; fi19++ {
                if false {
                    fmt.Println(fi19)
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := GenS2{F0: -438, F1: (false && false), F2: (v3 + v3)}
                        _ = fi20
                        _ = v21
                    }
                } else {
                    fmt.Println(v3, v3)
                }
                _ = fi19
            }
            v4 += v4
            v22 := GenS1{F0: 1000}
            _ = v22
        }
        _ = v16
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if true {
            v4 = (v4 + v4)
            for fi24 := 0; fi24 < 9; fi24++ {
                v4 = v4
                _ = fi24
            }
        }
        _ = fi23
    }
    fmt.Println(v4)
    fmt.Println(v3, v4, v3)
    fmt.Println(v4)
    fmt.Println(v4, v2)
    fmt.Println(v13, v4, v3)
    v3 = v3
    if (1 > v4) {
        v25 := GenS1{F0: -32768}
        _ = v25
    }
    switch 32767 {
    case -734:
        v2 = v2
    }
    if false {
        if true {
            fmt.Println(v13, v13, v2)
        } else {
            if true {
                fmt.Println(v3)
                var v26 int64 = 1
                _ = v26
            }
            for fi27 := 0; fi27 < 13; fi27++ {
                v28 := GenS3{F0: 2147483648}
                _ = fi27
                _ = v28
            }
            switch v3 {
            case "":
                v3 = v3
            }
        }
        var v29 float32 = v2
        _ = v29
    } else {
        v30 := []int{286, 641, -816, -259}
        for fi31 := 0; fi31 < 19; fi31++ {
            if (true || false) {
                if (false && (4294967295 >= 501)) {
                    _ = len(v30)
                } else {
                    var v32 int32 = (v4 ^ v4)
                    if true {
                        for fi33 := 0; fi33 < 12; fi33++ {
                            var v34 uint = 1
                            _ = fi33
                            _ = v34
                        }
                        var v35 int8 = 42
                        _ = v35
                    }
                    if false {
                        var v36 float64 = 0.0
                        _ = v36
                    }
                    _ = v32
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    if false {
                        var v38 int32 = 0
                        var v39 uint64 = 4294967295
                        _ = v38
                        _ = v39
                    }
                    if (662 <= 1) {
                        if ((1e10 != ((3.14159 * float64((v2 / v2))) - -2.25)) && (1000 < 0)) {
                            var v40 float64 = 0.0
                            _ = v40
                        } else {
                            var v41 uint32 = 65535
                            switch (fi37 + (fi37 + fi37)) {
                            case 2147483647:
                                fmt.Println(v3, v4, v41)
                            case 32767:
                                switch v3 {
                                case "\n\t":
                                    var v42 uint8 = 1
                                    _ = v42
                                case "":
                                    v43 := GenS4{F0: 74, F1: 0}
                                    _ = v43
                                }
                            default:
                                var v44 int64 = -9223372036854775808
                                _ = v44
                            }
                            if false {
                                var v45 uint32 = 1
                                var v46 int64 = -434
                                v2 -= (3.14159 + float32((0.1 + 1.5)))
                                _ = v45
                                _ = v46
                            } else {
                                switch v3 {
                                case "a":
                                    switch "a" {
                                    case "a":
                                        v2 = 1.5
                                    case "foo bar":
                                        v4 = v4
                                    }
                                default:
                                    v47 := GenS2{F0: -257, F1: false, F2: "\n\t"}
                                    _ = v47
                                }
                            }
                            switch v3 {
                            case "":
                                v2 -= v2
                            case "hello":
                                v3 = ""
                            default:
                                v13 -= v13
                                switch "\n\t" {
                                case "a":
                                    var v48 uint64 = 2
                                    _ = v48
                                case "":
                                    fmt.Println(fi31, v13)
                                }
                                if (897 < 656) {
                                    v2 = 1.5
                                }
                                var v49 int = -277
                                v13 = (v13 ^ 255)
                                _ = v49
                            }
                            v50 := GenS5{F0: 42}
                            var v51 uint8 = 0
                            _ = v41
                            _ = v50
                            _ = v51
                        }
                    } else {
                        var v52 int = 2147483647
                        _ = len(v5)
                        _ = v52
                    }
                    delete(v1, fi31)
                    v53 := GenS6{F0: 127}
                    fmt.Println(v2, v3)
                    _ = fi37
                    _ = v53
                }
            }
            v3 = v3
            var v54 uint8 = 0
            _ = fi31
            _ = v54
        }
        _ = v30
    }
    var v55 string = ""
    fmt.Println(v3, v3)
    var v56 string = v55
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
}
