package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool; F2 string }

func main() {
    var v1 int = 32767
    var v2 int8 = -67
    var v3 uint16 = 746
    if (255 > v3) {
        v4 := []int{(v1 + v1), (v1 + 32768), -1, (v1 - 4294967295)}
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 19; fi6++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v7 int = (fi5 & v1)
                _ = fi6
                _ = v7
            }
            var v8 bool = false
            _ = fi5
            _ = v8
        }
        _ = v4
    } else {
        v9 := []int{(v1 | (^(v1)))}
        v2 -= 127
        _ = v9
    }
    if (!(false && true)) {
        v2 = (-(v2))
        var v10 float32 = 1e10
        v11 := GenS1{F0: (v1 >> ((0) & 63))}
        _ = v10
        _ = v11
    } else {
        if false {
            var v12 float32 = 1.5
            _ = v12
        }
        if ((v3 & (^(v3))) >= v3) {
            if (v3 == (v3 + 32767)) {
                v3 = 205
                var v13 uint64 = 18446744073709551615
                _ = v13
            } else {
                var v14 uint8 = 2
                fmt.Println(v1, v2, v3)
                var v15 uint64 = 4294967295
                _ = v14
                _ = v15
            }
        } else {
            var v16 int16 = 42
            var v17 uint16 = (v3 ^ v3)
            v2 -= v2
            _ = v16
            _ = v17
        }
        var v18 int32 = -1
        _ = v18
    }
    switch (-(v1)) {
    case -826:
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := []int{int(v3)}
            v21 := GenS2{F0: "foo bar", F1: false, F2: ""}
            v2 -= 127
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    switch ("hello" + "hello") {
    case "a":
        var v22 string = ("hello" + "a")
        _ = v22
    default:
        v1 = v1
    }
    var v23 int16 = 1000
    v24 := []int{(v1 & v1), (v1 / ((v1) | 1))}
    v3 = 32767
    fmt.Println(v1, v1)
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS1{F0: (-(v1))}
        var v27 uint8 = 2
        if true {
            var v28 float32 = -2.25
            _ = v28
        } else {
            switch ("a" + "") {
            case "a":
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 uint16 = (v3 << ((1) & 63))
                    v31 := GenS1{F0: fi29}
                    var v32 uint16 = v3
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            case "hello":
                v23 -= (v23 + v23)
            default:
                if (1e10 == -2.25) {
                    switch "a" {
                    case "a":
                        _ = len(v24)
                    case "\n\t":
                        v23 -= (v23 & (v23 + 42))
                    default:
                        v23 = v23
                    }
                    for fi33 := 0; fi33 < 16; fi33++ {
                        var v34 float32 = ((1e10 - (3.14159 / 1.0)) - 0.1)
                        var v35 bool = (false || false)
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    if true {
                        fmt.Println(v2, fi25)
                        v2 = -1
                    }
                    var v36 int16 = 0
                    for fi37 := 0; fi37 < 19; fi37++ {
                        var v38 uint32 = 302
                        if false {
                            switch "a" {
                            case "":
                                switch "\n\t" {
                                case "a":
                                    var v39 uint16 = (v3 & v3)
                                    _ = v39
                                }
                            case "a":
                                var v40 int8 = -123
                                _ = v40
                            default:
                                v1 = (fi25 << ((255) & 63))
                            }
                            if ((18446744073709551615 >= 1) && true) {
                                if ((0.0 / 1.5) <= float32(1e10)) {
                                    if len(v24) > 0 {
                                        v24[0] = (fi37 & fi25)
                                    }
                                }
                                v1 += int(v38)
                                _ = len(v24)
                            }
                            switch ("\n\t" + "a") {
                            case "hello":
                                fmt.Println(v27)
                            case "foo bar":
                                var v41 int16 = v36
                                _ = v41
                            case "a":
                                v24 = append(v24, 2147483647)
                            default:
                                fmt.Println(fi25, v23, v38)
                            }
                        } else {
                            var v42 uint32 = v38
                            if (4294967295 > 2) {
                                fmt.Println(fi37)
                            } else {
                                var v43 uint16 = 255
                                fmt.Println(v3, v38, v38)
                                v3 = v3
                                _ = v43
                            }
                            v23 += v36
                            _ = v42
                        }
                        if true {
                            v36 = v23
                        } else {
                            v27 -= v27
                            var v44 uint8 = 218
                            _ = len(v24)
                            var v45 float64 = (1e10 + -2.25)
                            _ = v44
                            _ = v45
                        }
                        _ = len(v24)
                        if (0 >= 255) {
                            v23 = -679
                            fmt.Println(fi25, fi25, v3)
                        } else {
                            fmt.Println(v27)
                        }
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                }
            }
            v3 += 65535
            v27 += v27
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
