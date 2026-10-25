package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16, p1 int16, p2 uint64) uint {
    if false {
        fmt.Println(p2, p2, p2)
        fmt.Println(p1, p2, p0)
        for fi5 := 0; fi5 < 2; fi5++ {
            p0 = -1
            switch fi5 {
            case 267:
                fmt.Println(p1)
            case -1:
                p1 = (p1 | p1)
            case 1:
                for fi6 := 0; fi6 < 18; fi6++ {
                    if ((("a" + "\n\t") + "foo bar") >= "a") {
                        if true {
                            var v7 uint64 = 240
                            var v8 uint = 18446744073709551615
                            _ = v7
                            _ = v8
                        } else {
                            if false {
                                if false {
                                    v9 := GenS1{F0: ("" + "foo bar")}
                                    if false {
                                        p1 = p1
                                    }
                                    if false {
                                        fmt.Println(p1, fi6)
                                        if true {
                                            p2 = (p2 | 18446744073709551615)
                                        }
                                        var v10 float32 = 3.14159
                                        _ = v10
                                    } else {
                                        v11 := GenS1{F0: "a"}
                                        switch "hello" {
                                        case "":
                                            var v12 uint32 = 0
                                            _ = v12
                                        case "a":
                                            var v13 float64 = -2.25
                                            _ = v13
                                        }
                                        _ = v11
                                    }
                                    v14 := []int{}
                                    _ = v14
                                    _ = v9
                                }
                            } else {
                                p2 = 1
                            }
                            var v15 string = ("foo bar" + "hello")
                            var v16 int64 = 1
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: "foo bar"}
                        _ = v17
                    }
                    v18 := GenS1{F0: "\n\t"}
                    var v19 float32 = 3.14159
                    _ = fi6
                    _ = v18
                    _ = v19
                }
            }
            v20 := GenS1{F0: "\n\t"}
            _ = fi5
            _ = v20
        }
    }
    v21 := []int{-9223372036854775808, -2147483648, 448, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 42
}
func genFn1(p0 uint16, p1 int8, p2 int) uint8 {
    fmt.Println(p2)
    var v22 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    v3 := []int{127, 4294967295}
    v4 := make(map[uint32]uint32)
    if false {
        switch 0 {
        case -904:
            v1 = v2
        case 32767:
            var v23 float32 = 3.14159
            _ = v23
        }
        var v24 int32 = 0
        fmt.Println(v24, v2, v1)
        _ = v24
    }
    v1 = ((v1 + v1) + v2)
    var v25 uint16 = 0
    if (255 <= 588) {
        if true {
            for fi26 := 0; fi26 < 11; fi26++ {
                switch v1 {
                case "foo bar":
                    v1 = ((v2 + "hello") + v2)
                case "hello":
                    v27 := GenS1{F0: v2}
                    _ = v27
                case "\n\t":
                    v2 = v1
                default:
                    for fi28 := 0; fi28 < 16; fi28++ {
                        v3 = append(v3, (fi26 % ((fi26) | 1)))
                        var v29 int8 = 42
                        fmt.Println(v2, v2, fi28)
                        _ = fi28
                        _ = v29
                    }
                }
                var v30 uint = 42
                _ = fi26
                _ = v30
            }
            var v31 uint16 = (^(v25))
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 15; fi32++ {
            v1 = v2
            if true {
                v25 = v25
            }
            _ = fi32
        }
    }
    v4[42] = 4294967295
    if (261 >= 42) {
        if (1000 <= 1000) {
            v1 = "foo bar"
            if (false || (-9223372036854775808 <= -35)) {
                for fi33 := 0; fi33 < 1; fi33++ {
                    v34 := GenS1{F0: v2}
                    v1 = (("a" + "\n\t") + (v2 + ("hello" + "\n\t")))
                    _ = fi33
                    _ = v34
                }
            } else {
                var v35 int8 = 127
                if true {
                    v2 = "\n\t"
                    v25 = (v25 >> ((1) & 63))
                    var v36 int8 = v35
                    _ = v36
                }
                _ = v35
            }
            for fi37 := 0; fi37 < 5; fi37++ {
                v38 := GenS2{F0: 35, F1: 61, F2: -1}
                _ = fi37
                _ = v38
            }
        }
        v39 := GenS3{F0: 42}
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 uint8 = 128
            v2 = v1
            for fi42 := 0; fi42 < 5; fi42++ {
                v1 = "foo bar"
                _ = fi42
            }
            _ = fi40
            _ = v41
        }
        _ = v39
    } else {
        v2 = v1
        var v43 int8 = -27
        if false {
            v44 := v4[2147483648]
            v1 = ("hello" + v2)
            _ = v44
        }
        _ = v43
    }
    v25 = (^(v25))
    v45 := GenS4{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v45
}
