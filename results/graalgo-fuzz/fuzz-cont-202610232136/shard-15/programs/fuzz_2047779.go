package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int64 }

func main() {
    var v1 uint8 = 2
    var v2 uint = 65535
    var v3 float32 = 0.0
    var v4 int8 = 1
    fmt.Println(v4, v4)
    if false {
        fmt.Println(v2, v1)
        if ((!(!true)) || ((v3 + float32((-2.25 - ((((0.1 - float64(((v3 - float32(((1.5 / ((0.1 + -2.25) - (((0.0 - 3.14159) * 1.5) - -2.25))) / float64(v3)))) / float32(float64(float32(1e10)))))) - ((3.14159 - 1e10) / float64(v3))) * 1e10) - -2.25)))) <= 0.1)) {
            var v5 int32 = -1
            v5 = (v5 * 42)
            _ = v5
        }
        if false {
            switch 128 {
            case 42:
                v6 := GenS1{F0: 1000}
                _ = v6
            case 127:
                for fi7 := 0; fi7 < 2; fi7++ {
                    v2 = (^(v2))
                    v2 += 42
                    var v8 int = -331
                    _ = fi7
                    _ = v8
                }
            case -9223372036854775808:
                var v9 int8 = v4
                _ = v9
            }
            if false {
                if false {
                    v2 += v2
                    v4 = -108
                }
                var v10 uint64 = 42
                if true {
                    v3 = v3
                    for fi11 := 0; fi11 < 2; fi11++ {
                        if true {
                            var v12 uint8 = 127
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                switch "a" {
                                case "\n\t":
                                    if (!false) {
                                        v2 = v2
                                        var v14 uint8 = v1
                                        var v15 int8 = (v4 & 42)
                                        _ = v14
                                        _ = v15
                                    }
                                default:
                                    if false {
                                        switch fi11 {
                                        case -9223372036854775808:
                                            v10 -= (^(v10))
                                        default:
                                            var v16 int8 = (^(v4))
                                            _ = v16
                                        }
                                    } else {
                                        v3 -= float32(-2.25)
                                        var v17 uint16 = 345
                                        var v18 uint32 = 905
                                        v19 := []int{(^(fi11)), fi13, (^(fi11)), 867}
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                _ = fi13
                            }
                            v20 := GenS2{F0: false, F1: false, F2: (v4 ^ v4)}
                            var v21 float32 = 0.0
                            _ = v20
                            _ = v21
                        }
                        _ = fi11
                    }
                    var v22 string = ""
                    _ = v22
                }
                _ = v10
            }
            var v23 float64 = 3.14159
            _ = v23
        } else {
            switch "\n\t" {
            case "\n\t":
                v24 := GenS3{F0: v4}
                _ = v24
            case "hello":
                var v25 uint64 = 42
                _ = v25
            case "foo bar":
                v26 := []int{1}
                _ = v26
            }
        }
    } else {
        v27 := []int{658, -9223372036854775808}
        _ = v27
    }
    var v28 uint8 = 232
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(fi29, v28)
        _ = fi29
    }
    switch (("foo bar" + "foo bar") + ("foo bar" + (("" + "hello") + "foo bar"))) {
    case "":
        v30 := GenS4{F0: -9223372036854775808}
        _ = v30
    case "hello":
        v31 := GenS5{F0: v1, F1: ("" + "")}
        _ = v31
    default:
        for fi32 := 0; fi32 < 9; fi32++ {
            for fi33 := 0; fi33 < 5; fi33++ {
                fmt.Println(v4, fi32, v2)
                var v34 string = "a"
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 8; fi35++ {
                fmt.Println(v28, v4, v2)
                v36 := []int{-40}
                _ = fi35
                _ = v36
            }
            _ = fi32
        }
    }
    v37 := GenS6{F0: 1, F1: 42}
    var v38 float64 = 1.5
    v28 += 127
    switch "foo bar" {
    case "\n\t":
        switch 45 {
        case 127:
            v4 = v4
        default:
            var v39 float32 = (v3 + ((float32(v38) / -2.25) * ((v3 + ((3.14159 - v3) - v3)) - (0.1 / v3))))
            _ = v39
        }
    default:
        switch (("foo bar" + "") + "hello") {
        case "":
            for fi40 := 0; fi40 < 7; fi40++ {
                v41 := GenS3{F0: v4}
                _ = fi40
                _ = v41
            }
        case "foo bar":
            fmt.Println(v1, v3)
        case "\n\t":
            v42 := GenS5{F0: v28, F1: "a"}
            _ = v42
        }
    }
    v43 := []int{0, 2147483647, 42}
    _ = len(v43)
    fmt.Println(v1, v28)
    if (-1 <= 0) {
        for fi44 := 0; fi44 < 13; fi44++ {
            var v45 int32 = -455
            _ = fi44
            _ = v45
        }
        fmt.Println(v3, v1)
    }
    switch "hello" {
    case "a":
        v3 += float32(v38)
    case "":
        v1 = (v1 - v28)
    }
    v46 := GenS7{F0: -124, F1: 42}
    switch -28 {
    case 1:
        for fi47 := 0; fi47 < 18; fi47++ {
            v38 = -2.25
            fmt.Println(v3)
            _ = fi47
        }
    case -965:
        if (v4 > v4) {
            v48 := GenS8{F0: 550}
            for fi49 := 0; fi49 < 10; fi49++ {
                v50 := GenS7{F0: v4, F1: 42}
                fmt.Println(v2, v28, v3)
                if len(v43) >= 2 {
                    v43 = v43[1:len(v43)]
                }
                _ = fi49
                _ = v50
            }
            v4 += (v4 - 127)
            _ = v48
        }
    case -1:
        var v51 bool = (111 <= v4)
        _ = v51
    }
    var v52 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v46
    _ = v52
}
