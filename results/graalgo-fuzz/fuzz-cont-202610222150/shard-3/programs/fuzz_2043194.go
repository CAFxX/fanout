package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }

func main() {
    var v1 float32 = (0.0 / 1.0)
    v2 := make(map[int]bool)
    var v3 uint64 = 1
    switch "foo bar" {
    case "hello":
        v3 = (v3 / (((v3 - (v3 % ((2) | 1)))) | 1))
    case "\n\t":
        switch "" {
        case "\n\t":
            v3 += 1
        }
    case "foo bar":
        switch 42 {
        case 128:
            v4 := GenS1{F0: 105, F1: 1, F2: 42}
            _ = v4
        case -94:
            var v5 uint64 = (v3 << ((1) & 63))
            _ = v5
        case 42:
            var v6 uint8 = 164
            _ = v6
        }
    default:
        var v7 uint16 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: (fi8 & fi8), F1: 685, F2: 0}
        if (1 > -154) {
            var v10 string = "\n\t"
            if false {
                var v11 uint64 = v3
                var v12 int64 = -361
                _ = v11
                _ = v12
            }
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v13 := GenS2{F0: -9223372036854775808, F1: 42, F2: 2}
    if (514 != 42) {
        v14 := GenS3{F0: 2147483647, F1: 463, F2: -2147483648}
        if (170 != -352) {
            if false {
                v15 := GenS4{F0: 125}
                if true {
                    v16 := GenS1{F0: 0, F1: 2, F2: 95}
                    _ = v16
                }
                _ = v15
            } else {
                switch ("hello" + "hello") {
                case "\n\t":
                    if (128 <= 1) {
                        switch -475 {
                        case 9223372036854775807:
                            var v17 bool = false
                            _ = v17
                        }
                    } else {
                        fmt.Println(v1)
                        v18 := GenS3{F0: 610, F1: -1, F2: 9223372036854775807}
                        _ = v18
                    }
                case "":
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v20 := v2[fi19]
                        _ = fi19
                        _ = v20
                    }
                }
                v21 := GenS5{F0: 712}
                _ = v21
            }
            var v22 uint64 = v3
            _ = v22
        }
        var v23 int16 = -187
        _ = v14
        _ = v23
    }
    var v24 float64 = -2.25
    if false {
        if (!((true && false) || ("a" != ("a" + "foo bar")))) {
            switch "\n\t" {
            case "foo bar":
                v25 := GenS2{F0: -1, F1: 10, F2: 255}
                _ = v25
            case "\n\t":
                fmt.Println(v1, v1, v24)
            default:
                if ((!true) || (!false)) {
                    switch 9223372036854775807 {
                    case 32768:
                        var v26 int = 32767
                        _ = v26
                    case 412:
                        v1 -= v1
                    default:
                        if true {
                            v3 += v3
                        } else {
                            var v27 int8 = -1
                            var v28 string = "\n\t"
                            v29 := GenS6{F0: 128}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                }
            }
            if (!(false || true)) {
                _ = len(v2)
                if ("" > ("foo bar" + ("" + ("a" + "")))) {
                    if (4294967295 <= 4294967295) {
                        var v30 uint64 = v3
                        switch "hello" {
                        case "a":
                            v31 := GenS2{F0: -190, F1: -128, F2: 0}
                            _ = v31
                        case "hello":
                            v3 += 2
                        default:
                            var v32 int16 = 32767
                            _ = v32
                        }
                        _ = v30
                    } else {
                        switch (("foo bar" + "\n\t") + "\n\t") {
                        case "hello":
                            for fi33 := 0; fi33 < 14; fi33++ {
                                v3 += 4294967295
                                _ = fi33
                            }
                        case "\n\t":
                            v3 = 894
                        case "foo bar":
                            v2[128] = false
                        }
                        v3 -= (^(v3))
                    }
                    var v34 string = "a"
                    _ = v34
                }
            } else {
                if false {
                    _ = len(v2)
                }
                fmt.Println(v3)
                fmt.Println(v3)
            }
        } else {
            fmt.Println(v1, v1, v1)
            for fi35 := 0; fi35 < 13; fi35++ {
                if false {
                    delete(v2, fi35)
                }
                switch ("foo bar" + "") {
                case "a":
                    v3 = v3
                }
                _ = fi35
            }
            v36 := []int{32768}
            _ = v36
        }
    } else {
        v3 -= v3
        var v37 uint32 = 4294967295
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v39 := GenS7{F0: ("" + "foo bar")}
        v40 := GenS2{F0: fi38, F1: -1, F2: 158}
        if ((-(fi38)) <= fi38) {
            v41 := []int{}
            var v42 int64 = 42
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    var v43 int16 = 32767
    switch 1 {
    case 0:
        v1 += v1
    case -2147483648:
        var v44 uint16 = 834
        _ = v44
    case 291:
        fmt.Println(v1, v43, v1)
    default:
        delete(v2, -1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v43
}
