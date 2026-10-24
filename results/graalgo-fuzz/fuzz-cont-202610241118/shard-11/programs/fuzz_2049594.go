package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2, 1, 1}
    v3 := []int{-2147483648, 32767, 4294967295, -446}
    v4 := make(map[int]string)
    var v5 int16 = 0
    v6 := GenS1{F0: (v5 / (((v5 | v5)) | 1)), F1: true, F2: -301}
    _ = len(v3)
    v5 = (v5 * -26)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 float64 = (-2.25 + 1.5)
        for fi9 := 0; fi9 < 9; fi9++ {
            switch "hello" {
            case "\n\t":
                if false {
                    v5 = -582
                    var v10 uint8 = 197
                    _ = v10
                }
            case "hello":
                v11 := GenS1{F0: (v5 + 42), F1: false, F2: 2147483647}
                _ = v11
            }
            v5 = v5
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v5 -= 941
    _ = len(v2)
    v5 = v5
    if false {
        fmt.Println(v5)
        if false {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint8 = 128
                if (true && true) {
                    if true {
                        v14 := GenS1{F0: v5, F1: (true && false), F2: -754}
                        var v15 uint64 = 42
                        _ = v14
                        _ = v15
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                    }
                }
                var v16 int64 = 1
                _ = fi12
                _ = v13
                _ = v16
            }
            if true {
                _ = len(v2)
                v2 = append(v2, 65535)
            }
        } else {
            var v17 int16 = (v5 >> ((65535) & 63))
            _ = v17
        }
        v18 := GenS1{F0: v5, F1: (255 > 255), F2: 1}
        _ = v18
    } else {
        v19 := GenS2{F0: 2}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        switch "a" {
        case "":
            v5 += 0
        case "a":
            if false {
                switch "" {
                case "\n\t":
                    fmt.Println(v5, v5)
                default:
                    var v21 float64 = float64(-2.25)
                    _ = v21
                }
            } else {
                v22 := GenS1{F0: (v5 - v5), F1: (474 > 2), F2: 556}
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := GenS2{F0: 255}
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        case "hello":
            delete(v1, "\n\t")
        }
        v5 = (v5 | v5)
        fmt.Println(v5, fi20)
        _ = fi20
    }
    switch ("hello" + "") {
    case "foo bar":
        _ = len(v3)
    }
    var v25 uint = 554
    fmt.Println(v25, v25)
    v3 = append(v3, -1)
    if true {
        switch "hello" {
        case "foo bar":
            for fi26 := 0; fi26 < 1; fi26++ {
                v25 += 1
                v5 = v5
                _ = fi26
            }
        case "a":
            if (128 == 1) {
                v25 += v25
            } else {
                for fi27 := 0; fi27 < 3; fi27++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v28 := GenS1{F0: (v5 * (v5 * 789)), F1: (true && false), F2: -4}
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 int32 = 745
                    v1["foo bar"] = 2147483647
                    var v31 uint8 = 128
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                if false {
                    fmt.Println(v25, v5)
                    v32 := GenS1{F0: (-(v5)), F1: (!(1.5 == 1.5)), F2: 0}
                    _ = v32
                } else {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        var v33 uint = (^(v25))
                        _ = v33
                    case "":
                        v25 += 1
                    case "hello":
                        v4[32768] = (("hello" + ("foo bar" + "a")) + "a")
                    }
                }
            }
        default:
            if (v5 >= v5) {
                var v34 uint16 = 270
                _ = v34
            } else {
                fmt.Println(v25, v25)
                v5 = v5
                var v35 bool = (!(0 != 440))
                _ = v35
            }
        }
        if true {
            v25 = 2
            var v36 int16 = v5
            _ = v36
        } else {
            v25 = 65535
            v37 := GenS1{F0: v5, F1: false, F2: 42}
            fmt.Println(v25)
            _ = v37
        }
        var v38 int8 = -112
        _ = v38
    }
    v5 -= v5
    v39 := GenS1{F0: 43, F1: (2 != 809), F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
