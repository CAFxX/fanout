package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 string; F2 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 bool = (-18 >= 100)
    var v4 int8 = -1
    v5 := []uint32{65535, 0, 884, 1}
    v6 := make(map[int]string)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 string = "foo bar"
    delete(v1, v7)
    fmt.Println(v7, v4, v4)
    v4 = (v4 ^ 42)
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: -77, F1: 633}
        if (((1.5 * 0.0) - 0.0) != 1.5) {
            v4 = v4
            if (3.14159 >= 1.5) {
                _ = len(v5)
                if true {
                    var v10 uint64 = 1
                    _ = v10
                } else {
                    switch (fi8 >> ((657) & 63)) {
                    case -9223372036854775808:
                        switch fi8 {
                        case -9223372036854775808:
                            for fi11 := 0; fi11 < 7; fi11++ {
                                var v12 int32 = 1
                                _ = fi11
                                _ = v12
                            }
                        case 32767:
                            for fi13 := 0; fi13 < 1; fi13++ {
                                if v3 {
                                    v3 = (float64(-2.25) == 1e10)
                                    v14 := GenS1{F0: -1, F1: 1}
                                    v3 = v3
                                    _ = v14
                                }
                                _ = fi13
                            }
                        case 32768:
                            var v15 float32 = 1e10
                            _ = v15
                        default:
                            switch (-(fi8)) {
                            case 890:
                                fmt.Println(v7, v7, v3)
                            case 841:
                                v7 = v7
                            case 2147483647:
                                v16 := GenS1{F0: 32767, F1: 255}
                                _ = v16
                            }
                        }
                    default:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            _ = len(v6)
                            v5 = append(v5, 2147483647)
                            _ = fi17
                        }
                    }
                }
                if (true && v3) {
                    v7 = ("a" + "hello")
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := GenS2{F0: 106, F1: fi8}
                        var v20 int8 = 100
                        v6[(fi8 / ((-2147483648) | 1))] = (v7 + v7)
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint = 915
                    _ = v21
                } else {
                    if true {
                        v3 = v3
                        if (true && (32767 != -1)) {
                            for fi22 := 0; fi22 < 9; fi22++ {
                                v3 = v3
                                var v23 uint = 921
                                _ = fi22
                                _ = v23
                            }
                            var v24 int = (fi8 >> ((665) & 63))
                            _ = v24
                        }
                        v25 := GenS1{F0: 442, F1: 1}
                        _ = v25
                    }
                }
            } else {
                var v26 int = fi8
                _ = v26
            }
        } else {
            var v27 uint8 = 206
            if v3 {
                var v28 int8 = v4
                v29 := GenS3{F0: (42 != 594), F1: (v7 + "")}
                _ = v28
                _ = v29
            } else {
                v4 = 100
                v27 = v27
                fmt.Println(fi8, v27)
            }
            v27 = 255
            _ = v27
        }
        _ = fi8
        _ = v9
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        switch "" {
        case "hello":
            var v31 uint32 = 0
            _ = v31
        case "foo bar":
            v7 = "\n\t"
        case "a":
            switch (v7 + v7) {
            case "foo bar":
                var v32 int64 = 1
                _ = v32
            case "":
                v7 = v7
            }
        default:
            var v33 uint8 = 128
            _ = v33
        }
        v4 = (v4 % (((v4 % (((v4 / ((-70) | 1))) | 1))) | 1))
        _ = fi30
    }
    switch 778 {
    case -9223372036854775808:
        for fi34 := 0; fi34 < 11; fi34++ {
            for fi35 := 0; fi35 < 5; fi35++ {
                v36 := GenS4{F0: 65535}
                _ = fi35
                _ = v36
            }
            v4 = 0
            _ = fi34
        }
    case -711:
        switch 4294967295 {
        case -9223372036854775808:
            switch (((v7 + v7) + "hello") + v7) {
            case "hello":
                var v37 bool = false
                _ = v37
            default:
                var v38 bool = v3
                _ = v38
            }
        default:
            for fi39 := 0; fi39 < 19; fi39++ {
                if (!(v7 < "\n\t")) {
                    fmt.Println(v4, v3)
                }
                _ = fi39
            }
        }
    case -888:
        fmt.Println(v4, v4)
    default:
        v40 := GenS1{F0: -32768, F1: 32767}
        _ = v40
    }
    if false {
        v41 := GenS5{F0: v3, F1: v7, F2: 1}
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS2{F0: 2, F1: fi42}
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    v5 = append(v5, 1)
    v44 := GenS4{F0: 1}
    var v45 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
