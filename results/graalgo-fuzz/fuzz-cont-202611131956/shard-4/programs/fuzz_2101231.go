package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint64 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 string = "hello"
    delete(v2, "a")
    switch v4 {
    case "a":
        switch v4 {
        case "hello":
            var v5 uint = 65535
            _ = v5
        }
    default:
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                if true {
                    if (false && (!((-128 != 0) && (65535 >= 0)))) {
                        if true {
                            switch v4 {
                            case "foo bar":
                                v1 -= v1
                            case "hello":
                                fmt.Println(fi7, fi6)
                            }
                        }
                    }
                    v8 := []int{(^(fi6)), 128, (^(fi7)), fi7}
                    v1 += (v1 + -2.25)
                    _ = v8
                }
                _ = fi7
            }
            v1 += v1
            _ = fi6
        }
    }
    switch v4 {
    case "foo bar":
        fmt.Println(v1, v4)
    case "hello":
        v9 := GenS1{F0: 127}
        _ = v9
    default:
        switch v4 {
        case "foo bar":
            if true {
                fmt.Println(v4)
                for fi10 := 0; fi10 < 11; fi10++ {
                    switch v4 {
                    case "hello":
                        v1 = v1
                    case "foo bar":
                        for fi11 := 0; fi11 < 5; fi11++ {
                            fmt.Println(v4, v4, fi11)
                            v12 := GenS2{F0: -345, F1: 1, F2: -1}
                            var v13 string = ("" + (v4 + "\n\t"))
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    case "a":
                        v14 := GenS3{F0: 1}
                        _ = v14
                    default:
                        var v15 bool = false
                        _ = v15
                    }
                    _ = fi10
                }
            } else {
                switch 32767 {
                case 42:
                    var v16 float32 = 0.1
                    _ = v16
                default:
                    v17 := GenS4{F0: 0, F1: v4, F2: v4}
                    _ = v17
                }
            }
        case "\n\t":
            delete(v3, 993)
        default:
            v18 := []int{}
            _ = v18
        }
    }
    if false {
        v1 += (v1 * v1)
        for fi19 := 0; fi19 < 17; fi19++ {
            v4 = v4
            _ = fi19
        }
        fmt.Println(v4)
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS3{F0: 1}
            _ = fi20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 16; fi22++ {
            v4 = ((v4 + v4) + v4)
            _ = fi22
        }
        v23 := GenS5{F0: 757}
        var v24 float32 = float32(v1)
        _ = v23
        _ = v24
    }
    v4 = "hello"
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v1, v4, v4)
        _ = fi25
    }
    _ = len(v3)
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := []int{fi26}
        _ = fi26
        _ = v27
    }
    v1 = v1
    var v28 float64 = (v1 - v1)
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(fi29)
        _ = fi29
    }
    if false {
        v1 += v1
        switch 128 {
        case 1:
            var v30 uint64 = 42
            _ = v30
        case 4294967295:
            delete(v2, ("\n\t" + v4))
        }
    }
    _ = len(v2)
    v31 := GenS4{F0: 0, F1: v4, F2: v4}
    if true {
        var v32 string = v4
        if (0 >= 42) {
            var v33 uint16 = 65535
            _ = v33
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            v35 := GenS6{F0: 2147483647}
            if true {
                if false {
                    var v36 float32 = -2.25
                    _ = v36
                }
                var v37 float32 = float32((-2.25 - 1.5))
                _ = v37
            } else {
                switch v32 {
                case "hello":
                    switch (fi34 * 0) {
                    case 965:
                        for fi38 := 0; fi38 < 19; fi38++ {
                            fmt.Println(fi38)
                            fmt.Println(fi34)
                            _ = fi38
                        }
                    case 358:
                        v39 := GenS7{F0: -1, F1: 0}
                        _ = v39
                    case 1:
                        for fi40 := 0; fi40 < 13; fi40++ {
                            switch (v4 + "hello") {
                            case "":
                                fmt.Println(v4, v4)
                            case "foo bar":
                                v41 := GenS4{F0: 1, F1: v4, F2: (v4 + "")}
                                _ = v41
                            }
                            if (((v32 <= v32) && true) && true) {
                                var v42 string = "hello"
                                _ = v42
                            } else {
                                v4 = "foo bar"
                                var v43 uint64 = 807
                                _ = v43
                            }
                            _ = fi40
                        }
                    }
                case "foo bar":
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v28 = 3.14159
                        var v45 int64 = -9223372036854775808
                        _ = fi44
                        _ = v45
                    }
                default:
                    fmt.Println(v1, v28, v32)
                }
            }
            _ = fi34
            _ = v35
        }
        _ = v32
    }
    if (2147483647 > 42) {
        switch 9223372036854775807 {
        case 325:
            v46 := []int{}
            _ = v46
        case 70:
            v28 = (v1 / ((v1 + v28) - (v28 + (v28 * -2.25))))
        }
        var v47 uint8 = 202
        delete(v3, 128)
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
