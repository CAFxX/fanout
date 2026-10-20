package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0() int64 {
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 13; fi7++ {
        v6 += (v6 ^ 366)
        var v8 uint = 599
        _ = fi7
        _ = v8
    }
    _ = v6
    return 9223372036854775807
}

func main() {
    v1 := []uint32{}
    var v2 string = ("a" + "foo bar")
    v3 := make(map[string]bool)
    var v4 int8 = -76
    var v5 float32 = 0.1
    var v9 string = (v2 + "foo bar")
    v2 = v2
    v10 := GenS1{F0: -2147483648}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 65535)
    v11 := GenS2{F0: true, F1: 2}
    if (!false) {
        v12 := GenS1{F0: 1}
        v13 := GenS1{F0: 0}
        _ = v12
        _ = v13
    }
    var v14 int64 = -686
    if (v4 == (v4 ^ v4)) {
        if false {
            switch 32768 {
            case -1:
                fmt.Println(v9, v9, v2)
            case 762:
                v15 := GenS1{F0: -693}
                _ = v15
            default:
                var v16 uint64 = 864
                _ = v16
            }
            fmt.Println(v14, v5, v9)
        } else {
            v2 = "hello"
            v5 = ((-2.25 + float32((3.14159 - 1.5))) + (v5 * 1e10))
            _ = len(v1)
        }
        v17 := GenS2{F0: true, F1: 1}
        switch 4294967295 {
        case -343:
            switch -1 {
            case -2147483648:
                fmt.Println(v2)
            case -463:
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 uint16 = 339
                    _ = fi18
                    _ = v19
                }
            case 4294967295:
                fmt.Println(v14)
            default:
                var v20 float32 = v5
                _ = v20
            }
        default:
            var v21 uint32 = 0
            _ = v21
        }
        _ = v17
    } else {
        v9 = ("" + v2)
        if true {
            var v22 uint = 356
            switch 1 {
            case 336:
                v4 = v4
            }
            v23 := GenS1{F0: -1}
            _ = v22
            _ = v23
        } else {
            v4 += (v4 >> ((2) & 63))
        }
        var v24 string = v2
        _ = v24
    }
    if true {
        v2 = (v9 + "\n\t")
        var v25 int = 2147483647
        var v26 uint = 0
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 int64 = (v14 & v14)
        if false {
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(v14, v4)
                _ = len(v3)
                _ = fi29
            }
            var v30 uint = 954
            if (42 != 2) {
                v4 -= 127
                v2 = v2
                var v31 uint16 = 42
                _ = v31
            } else {
                var v32 uint32 = 2147483648
                if false {
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if false {
                            v34 := GenS2{F0: (!((!false) && (1 >= -316))), F1: 65535}
                            v5 = (0.1 + 3.14159)
                            _ = v34
                        }
                        if true {
                            v35 := GenS2{F0: (false || false), F1: (v30 * 42)}
                            var v36 uint = 545
                            v1 = append(v1, v32)
                            _ = v35
                            _ = v36
                        } else {
                            v28 -= genFn0()
                        }
                        var v37 string = v2
                        _ = fi33
                        _ = v37
                    }
                    switch "hello" {
                    case "foo bar":
                        var v38 string = v2
                        _ = v38
                    default:
                        var v39 int8 = v4
                        _ = v39
                    }
                } else {
                    fmt.Println(fi27, v9, v5)
                    v9 = v2
                }
                _ = v32
            }
            _ = v30
        } else {
            var v40 string = v2
            switch v9 {
            case "a":
                _ = len(v1)
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                v9 = v9
            }
            v28 = (v28 + 42)
            _ = v40
        }
        switch fi27 {
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -2147483648:
            for fi41 := 0; fi41 < 16; fi41++ {
                if (v14 >= (v28 + 0)) {
                    delete(v3, v9)
                } else {
                    fmt.Println(v4)
                    switch ((v2 + (v9 + v2)) + (v2 + ("foo bar" + ("hello" + (v9 + (v2 + "hello")))))) {
                    case "a":
                        var v42 int32 = -10
                        switch "a" {
                        case "hello":
                            var v43 float32 = v5
                            _ = v43
                        default:
                            v44 := GenS3{F0: v9}
                            _ = v44
                        }
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                        _ = v42
                    case "":
                        v4 = v4
                    default:
                        v2 = v2
                    }
                    v4 -= (v4 | (v4 - v4))
                    fmt.Println(v28, fi27, v9)
                    v45 := GenS4{F0: 128, F1: false}
                    var v46 int = (fi27 | (fi27 + (^(fi27))))
                    _ = v45
                    _ = v46
                }
                _ = fi41
            }
        default:
            fmt.Println(fi27)
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
