package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 uint8 = 127
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int32 = 42
            _ = len(v1)
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 1, F1: 42}
        _ = fi4
        _ = v7
    }
    if (302 >= 868) {
        if (true || (1 < 0)) {
            switch ("foo bar" + "\n\t") {
            case "a":
                v3 -= (v3 << ((266) & 63))
            case "\n\t":
                if true {
                    delete(v1, 326)
                } else {
                    if (!(182 <= (v3 / ((v3) | 1)))) {
                        v8 := GenS1{F0: 1, F1: 0}
                        _ = v8
                    } else {
                        var v9 string = ""
                        for fi10 := 0; fi10 < 9; fi10++ {
                            if len(v2) > 0 {
                                v2[0] = v9
                            }
                            v11 := v1[0]
                            _ = fi10
                            _ = v11
                        }
                        _ = v9
                    }
                }
            }
        } else {
            if true {
                v3 = v3
            } else {
                fmt.Println(v3)
                fmt.Println(v3)
                var v12 uint16 = 376
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int32 = 40
            _ = fi13
            _ = v14
        }
        var v15 uint8 = (v3 << ((2) & 63))
        _ = v15
    }
    var v16 float64 = (((1e10 + (((1e10 / -2.25) - 0.0) + 3.14159)) * 3.14159) - 1.5)
    for fi17 := 0; fi17 < 19; fi17++ {
        v16 = (float64(1.5) + float64(1e10))
        v3 = 255
        _ = fi17
    }
    var v18 uint = 587
    var v19 float32 = 0.0
    switch "\n\t" {
    case "a":
        for fi20 := 0; fi20 < 12; fi20++ {
            if ((!(true || false)) || ((!((v18 << ((801) & 63)) != v18)) && (864 > 18446744073709551615))) {
                v3 = v3
                if false {
                    var v21 float32 = 0.0
                    switch (("a" + "\n\t") + ("" + "")) {
                    case "a":
                        if false {
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v21 = v19
                                v23 := GenS1{F0: -848, F1: 255}
                                _ = fi22
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 3; fi24++ {
                                v16 += float64(v19)
                                _ = fi24
                            }
                            var v25 bool = false
                            _ = v25
                        } else {
                            var v26 float32 = v21
                            var v27 string = ("hello" + (("hello" + "") + "foo bar"))
                            _ = len(v1)
                            _ = v26
                            _ = v27
                        }
                    case "foo bar":
                        v28 := GenS1{F0: 1, F1: 1}
                        _ = v28
                    case "hello":
                        v19 = 1.5
                    default:
                        fmt.Println(v16, v3, v3)
                    }
                    v29 := GenS2{F0: 32767}
                    _ = v21
                    _ = v29
                }
            } else {
                var v30 float64 = float64(v19)
                if true {
                    var v31 int16 = -1
                    if true {
                        fmt.Println(fi20)
                        var v32 int64 = int64(v31)
                        _ = v32
                    } else {
                        v33 := GenS1{F0: 1, F1: 255}
                        v30 = (((0.0 + v30) * v30) * v16)
                        _ = v33
                    }
                    _ = v31
                }
                _ = v30
            }
            v34 := GenS3{F0: (^(v18)), F1: 42}
            for fi35 := 0; fi35 < 6; fi35++ {
                switch (-(fi20)) {
                case 32768:
                    if (((!true) || (v3 <= v3)) || ((v18 >> ((v18) & 63)) < (^(v18)))) {
                        var v36 int16 = 1000
                        switch "\n\t" {
                        case "a":
                            v37 := GenS4{F0: (-1 != 0)}
                            _ = v37
                        default:
                            v18 -= v18
                        }
                        _ = v36
                    } else {
                        v38 := GenS3{F0: 452, F1: (v3 | (^(v3)))}
                        var v39 int64 = 9223372036854775807
                        _ = v38
                        _ = v39
                    }
                case 128:
                    var v40 int16 = 447
                    _ = v40
                case 0:
                    v41 := GenS2{F0: 0}
                    _ = v41
                }
                var v42 int = fi20
                var v43 uint32 = 4294967295
                _ = fi35
                _ = v42
                _ = v43
            }
            _ = fi20
            _ = v34
        }
    case "hello":
        for fi44 := 0; fi44 < 5; fi44++ {
            delete(v1, 189)
            _ = fi44
        }
    }
    v16 -= 1.5
    var v45 float64 = (v16 + 0.0)
    var v46 int8 = -128
    var v47 uint = v18
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 uint32 = 0
        v49 -= v49
        var v50 float32 = float32(v45)
        _ = fi48
        _ = v49
        _ = v50
    }
    var v51 int16 = 1000
    v18 += (v47 * (^(v18)))
    v1[771] = "hello"
    var v52 int = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v51
    _ = v52
}
