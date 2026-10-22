package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint; F1 string; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 uint16 = 107
    var v2 int16 = 32767
    var v3 float64 = float64(float32((3.14159 + ((0.1 - 0.0) + 0.1))))
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 8; fi5++ {
        if (0.0 == 1.5) {
            var v6 float64 = 0.1
            for fi7 := 0; fi7 < 12; fi7++ {
                fmt.Println(v2)
                var v8 uint = 685
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            var v9 bool = false
            _ = v9
        }
        if true {
            v10 := GenS1{F0: 948}
            v2 += 71
            var v11 int32 = 558
            _ = v10
            _ = v11
        }
        delete(v4, 65535)
        _ = fi5
    }
    v12 := GenS1{F0: 4294967295}
    v3 = 1e10
    v13 := GenS1{F0: 1}
    v4[2147483648] = ((42 <= 1) || true)
    v1 = v1
    fmt.Println(v2)
    var v14 bool = false
    v3 = 1.5
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            var v16 string = "a"
            _ = v16
        }
        v1 -= 65535
        if ((v2 > v2) && false) {
            fmt.Println(fi15, v1)
            v17 := GenS1{F0: 2147483647}
            _ = v17
        } else {
            v18 := []int{}
            _ = v18
        }
        _ = fi15
    }
    switch "\n\t" {
    case "\n\t":
        if (0.1 >= 1.5) {
            v2 = -89
            switch 247 {
            case 127:
                v19 := GenS2{F0: (v2 + v2)}
                _ = v19
            case 4294967295:
                v20 := GenS3{F0: "\n\t", F1: 336}
                _ = v20
            case 0:
                var v21 int64 = -1
                _ = v21
            default:
                v1 = (^(v1))
            }
            var v22 uint64 = 4294967295
            _ = v22
        }
    case "":
        v23 := GenS4{F0: 18446744073709551615, F1: uint64(v1)}
        _ = v23
    case "foo bar":
        var v24 int32 = 2147483647
        _ = v24
    default:
        for fi25 := 0; fi25 < 4; fi25++ {
            if (v14 && (42 == 349)) {
                switch "a" {
                case "":
                    var v26 uint = 499
                    _ = v26
                case "hello":
                    var v27 float64 = ((v3 / 1.0) / (v3 - v3))
                    _ = v27
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    fmt.Println(v14, v14)
                    var v29 int = (fi25 + fi25)
                    switch v29 {
                    case -162:
                        v30 := GenS5{F0: 281, F1: "hello", F2: v14}
                        _ = v30
                    case 127:
                        v3 = (float64((0.0 / -2.25)) + (float64(0.0) + float64(3.14159)))
                    }
                    _ = fi28
                    _ = v29
                }
            } else {
                v31 := GenS1{F0: 2147483647}
                var v32 int64 = 0
                _ = v31
                _ = v32
            }
            if (42 > 284) {
                if false {
                    v33 := GenS6{F0: (fi25 >> ((18446744073709551615) & 63))}
                    v34 := GenS7{F0: 1}
                    _ = v33
                    _ = v34
                } else {
                    v35 := GenS8{F0: 0}
                    switch fi25 {
                    case 32768:
                        for fi36 := 0; fi36 < 1; fi36++ {
                            if (42 > -679) {
                                if ((v2 % (((v2 >> ((1) & 63))) | 1)) <= 1000) {
                                    switch "a" {
                                    case "foo bar":
                                        v3 = float64(float32((v3 / (float64(1e10) / 3.14159))))
                                    case "hello":
                                        v37 := GenS4{F0: 18446744073709551615, F1: 42}
                                        _ = v37
                                    default:
                                        var v38 uint32 = 0
                                        _ = v38
                                    }
                                    v39 := GenS4{F0: 572, F1: 332}
                                    fmt.Println(v1, v2)
                                    _ = v39
                                }
                            } else {
                                var v40 uint = 2
                                v41 := GenS8{F0: 0}
                                _ = v40
                                _ = v41
                            }
                            _ = fi36
                        }
                    case -1:
                        switch ("" + "foo bar") {
                        case "":
                            v42 := GenS8{F0: 2147483647}
                            _ = v42
                        case "foo bar":
                            var v43 uint32 = 49
                            _ = v43
                        case "\n\t":
                            var v44 int = 574
                            _ = v44
                        }
                    case -434:
                        var v45 string = "foo bar"
                        _ = v45
                    default:
                        v1 = (v1 >> ((1) & 63))
                    }
                    v14 = (0 <= 42)
                    _ = v35
                }
                var v46 int32 = 1
                for fi47 := 0; fi47 < 11; fi47++ {
                    v1 = 0
                    v3 = 1e10
                    _ = fi47
                }
                _ = v46
            }
            _ = fi25
        }
    }
    var v48 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
