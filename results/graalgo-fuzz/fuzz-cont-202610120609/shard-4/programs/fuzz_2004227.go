package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 162
    var v3 uint = 2
    var v4 string = ("\n\t" + "")
    v5 := GenS1{F0: 2}
    v3 = v3
    _ = len(v1)
    var v6 float64 = ((-2.25 + 1e10) + -2.25)
    var v7 int16 = -909
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: true}
        var v10 int64 = int64(v7)
        v3 -= 2
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        v2 += v2
    } else {
        fmt.Println(v3)
        var v11 uint = (^(v3))
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int32 = 149
        v14 := GenS3{F0: 127, F1: 109, F2: v13}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := []int{int(v2), -1}
    v3 = (v3 << ((571) & 63))
    var v16 int64 = 0
    if ((v6 * float64(3.14159)) > float64((0.1 * (-2.25 - (1.5 - -2.25))))) {
        fmt.Println(v4, v16, v3)
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi18
            }
            if true {
                if (195 < (^(v2))) {
                    v19 := GenS4{F0: (fi17 | (fi17 | fi17)), F1: -1, F2: 65535}
                    _ = v19
                } else {
                    v20 := GenS4{F0: -1, F1: (v7 / ((v7) | 1)), F2: 2147483647}
                    if true {
                        switch v4 {
                        case "\n\t":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        fmt.Println(v2)
                        var v21 uint16 = 956
                        _ = v21
                    }
                    v2 += (v2 / ((v2) | 1))
                    _ = v20
                }
            }
            _ = fi17
        }
        var v22 int16 = (-(v7))
        _ = v22
    } else {
        fmt.Println(v3, v16, v16)
    }
    if (255 == v3) {
        fmt.Println(v3)
        fmt.Println(v6)
        var v23 uint64 = 331
        _ = v23
    } else {
        _ = len(v1)
        switch 129 {
        case 251:
            if ((-2.25 != ((1e10 * 0.1) + -2.25)) && true) {
                var v24 uint = v3
                v25 := v1[2]
                for fi26 := 0; fi26 < 4; fi26++ {
                    v7 = v7
                    _ = fi26
                }
                _ = v24
                _ = v25
            } else {
                v15 = append(v15, 1)
                switch ((v4 + v4) + "") {
                case "foo bar":
                    var v27 int = 128
                    _ = v27
                case "\n\t":
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if true {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                var v30 string = "hello"
                                switch v30 {
                                case "a":
                                    v6 = (-2.25 * v6)
                                    _ = len(v1)
                                    v7 -= v7
                                case "\n\t":
                                    var v31 uint64 = 42
                                    _ = v31
                                case "hello":
                                    if false {
                                        var v32 uint8 = 128
                                        var v33 uint64 = 4294967295
                                        var v34 float32 = float32(v6)
                                        switch (^(fi29)) {
                                        case -2147483648:
                                            v3 = (v3 & 4294967295)
                                        case 32768:
                                            v4 = ""
                                        case 4294967295:
                                            v35 := GenS4{F0: 128, F1: 1, F2: 4294967295}
                                            _ = v35
                                        default:
                                            if false {
                                                var v36 uint64 = (v33 | 1)
                                                _ = len(v1)
                                                _ = v36
                                            } else {
                                                var v37 uint = 4294967295
                                                _ = v37
                                            }
                                        }
                                        fmt.Println(v2, v34, fi29)
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    } else {
                                        var v38 int64 = -290
                                        _ = v38
                                    }
                                    fmt.Println(fi29, v30)
                                default:
                                    v39 := GenS5{F0: 42}
                                    _ = v39
                                }
                                if (0 <= 0) {
                                    v30 = (v4 + v4)
                                }
                                _ = fi29
                                _ = v30
                            }
                        }
                        _ = fi28
                    }
                case "":
                    var v40 uint64 = 734
                    _ = v40
                }
                v41 := GenS5{F0: 1}
                _ = v41
            }
        }
        var v42 bool = true
        _ = v42
    }
    v6 -= 0.1
    v2 = (v2 >> (((v3 >> ((v3) & 63))) & 63))
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint32 = 68
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
