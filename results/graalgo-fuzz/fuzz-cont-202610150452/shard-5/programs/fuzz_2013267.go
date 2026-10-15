package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int8; F2 uint8 }
type GenS5 struct { F0 string; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int32; F1 int8; F2 uint32 }

func main() {
    var v1 uint = 2
    var v2 int16 = 32767
    var v3 int16 = v2
    var v4 int32 = -1
    var v5 uint = v1
    var v6 float64 = 1.5
    switch -933 {
    case -627:
        v7 := []int{379, -152}
        _ = v7
    }
    if true {
        v1 = 65535
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int8 = -30
                v6 = float64(3.14159)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        v11 := []int{-848}
        _ = v11
    }
    fmt.Println(v4)
    switch ((("a" + "") + ((("\n\t" + "a") + "\n\t") + ("\n\t" + ((("\n\t" + "\n\t") + (("" + "") + ("a" + ("a" + ("a" + "foo bar"))))) + "a")))) + "foo bar") {
    case "foo bar":
        fmt.Println(v5, v2, v4)
    case "a":
        for fi12 := 0; fi12 < 9; fi12++ {
            if (!(1 > 2)) {
                if (-14 < -128) {
                    var v13 float64 = (0.0 - ((v6 - ((1.5 * v6) / (v6 / 3.14159))) / 1.0))
                    _ = v13
                } else {
                    v1 += (^(v1))
                    if (!false) {
                        fmt.Println(v3, v4)
                        v14 := GenS1{F0: v4, F1: 255}
                        v2 = (v2 >> ((2) & 63))
                        _ = v14
                    } else {
                        if false {
                            v4 -= (v4 >> ((v5) & 63))
                        } else {
                            switch fi12 {
                            case 741:
                                fmt.Println(v1, v4)
                            default:
                                var v15 uint16 = 255
                                _ = v15
                            }
                            var v16 float32 = 0.1
                            if (true && true) {
                                fmt.Println(v6, v6, v5)
                            }
                            _ = v16
                        }
                        var v17 int16 = (^(v3))
                        v6 = v6
                        _ = v17
                    }
                    if false {
                        if (v3 < 104) {
                            v5 += v1
                        } else {
                            v18 := GenS2{F0: v2}
                            for fi19 := 0; fi19 < 20; fi19++ {
                                var v20 int16 = (v2 * (v3 - v3))
                                var v21 uint = v1
                                var v22 uint32 = 2
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = v18
                        }
                    } else {
                        var v23 string = ""
                        _ = v23
                    }
                }
                var v24 int16 = v3
                _ = v24
            } else {
                v25 := GenS1{F0: -260, F1: 0}
                var v26 int16 = 1000
                _ = v25
                _ = v26
            }
            if false {
                v27 := GenS3{F0: v3}
                fmt.Println(v3)
                fmt.Println(v4)
                _ = v27
            }
            _ = fi12
        }
    case "\n\t":
        var v28 uint8 = 2
        _ = v28
    }
    v2 -= v3
    v29 := []int{}
    var v30 bool = false
    var v31 bool = true
    for fi32 := 0; fi32 < 18; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            v34 := GenS4{F0: fi33, F1: 1, F2: 2}
            v35 := GenS5{F0: "", F1: 1, F2: 1}
            _ = len(v29)
            _ = fi33
            _ = v34
            _ = v35
        }
        if (v4 < v4) {
            fmt.Println(v5, fi32)
            var v36 uint = v5
            if true {
                v29 = append(v29, -1)
                v4 -= v4
                switch (fi32 / (((fi32 * (fi32 << ((v36) & 63)))) | 1)) {
                case 32767:
                    switch ("a" + "foo bar") {
                    case "a":
                        var v37 string = ("hello" + "foo bar")
                        _ = v37
                    case "\n\t":
                        v1 -= (v36 / ((v36) | 1))
                    case "hello":
                        switch ((("foo bar" + "\n\t") + "a") + "hello") {
                        case "hello":
                            switch (fi32 / (((fi32 | -1)) | 1)) {
                            case -2147483648:
                                var v38 float64 = float64(1.5)
                                _ = v38
                            default:
                                v39 := GenS6{F0: (v4 ^ (v4 | (-(v4)))), F1: 96, F2: 65535}
                                _ = v39
                            }
                        case "foo bar":
                            v4 = v4
                        default:
                            if len(v29) > 0 {
                                v29[0] = -2147483648
                            }
                        }
                    default:
                        if false {
                            for fi40 := 0; fi40 < 2; fi40++ {
                                if len(v29) > 0 {
                                    v29[0] = fi40
                                }
                                var v41 int16 = (v2 | v3)
                                v6 -= float64(-2.25)
                                v29 = append(v29, fi40)
                                _ = fi40
                                _ = v41
                            }
                            var v42 int = (fi32 * (fi32 << ((v36) & 63)))
                            var v43 uint32 = 281
                            _ = v42
                            _ = v43
                        } else {
                            var v44 uint8 = 255
                            for fi45 := 0; fi45 < 8; fi45++ {
                                v36 = (v1 << (((v1 >> ((uint(v44)) & 63))) & 63))
                                v1 = v5
                                _ = fi45
                            }
                            v3 += (v3 - (v2 + v3))
                            _ = v44
                        }
                    }
                case 2147483647:
                    fmt.Println(v5, fi32)
                case 42:
                    v4 -= -864
                default:
                    v4 -= v4
                }
            } else {
                v29 = append(v29, (fi32 * fi32))
                if len(v29) > 0 {
                    v29[0] = fi32
                }
                v36 = (v36 | (v5 + (^(v5))))
            }
            _ = v36
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
