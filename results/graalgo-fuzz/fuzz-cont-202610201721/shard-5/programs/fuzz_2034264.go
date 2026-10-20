package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int64 }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 1
    var v3 bool = true
    var v4 uint8 = (^(v2))
    v2 += 255
    switch -154 {
    case -1:
        v5 := []int{-1, 2147483647, -237}
        _ = v5
    case 4294967295:
        var v6 float64 = 3.14159
        _ = v6
    case -295:
        v7 := []int{311, 128}
        _ = v7
    }
    fmt.Println(v2, v2, v2)
    if v3 {
        var v8 float32 = ((1.5 - float32((1e10 + (((float64(1e10) * v1) / (v1 / v1)) - -2.25)))) + float32(v1))
        _ = v8
    } else {
        v9 := GenS1{F0: 1, F1: 0}
        v1 = 1e10
        if (!((v1 / -2.25) <= v1)) {
            var v10 int16 = 42
            _ = v10
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := []int{42}
        switch "\n\t" {
        case "":
            v1 += 1e10
        case "foo bar":
            for fi13 := 0; fi13 < 7; fi13++ {
                if (false && ((fi13 * fi11) > (fi13 >> ((65535) & 63)))) {
                    if (1000 != -856) {
                        v14 := GenS2{F0: 1, F1: 493}
                        if true {
                            switch "" {
                            case "\n\t":
                                var v15 uint64 = 1
                                _ = v15
                            case "hello":
                                fmt.Println(fi11)
                            case "":
                                if v3 {
                                    if false {
                                        v16 := GenS3{F0: ("hello" + "foo bar"), F1: 996, F2: (fi11 | (fi11 >> ((18446744073709551615) & 63)))}
                                        _ = v16
                                    } else {
                                        switch -319 {
                                        case 4294967295:
                                            switch ("\n\t" + "\n\t") {
                                            case "hello":
                                                fmt.Println(fi13, fi13, fi11)
                                            case "\n\t":
                                                if (v2 >= 255) {
                                                    v3 = (!v3)
                                                }
                                            case "a":
                                                if v3 {
                                                    var v17 int8 = 42
                                                    var v18 float64 = v1
                                                    var v19 int = -1
                                                    _ = v17
                                                    _ = v18
                                                    _ = v19
                                                }
                                            default:
                                                switch fi11 {
                                                case 9223372036854775807:
                                                    var v20 int8 = 127
                                                    _ = v20
                                                case 156:
                                                    v21 := GenS2{F0: 42, F1: 42}
                                                    _ = v21
                                                case 932:
                                                    var v22 uint32 = 436
                                                    _ = v22
                                                default:
                                                    var v23 float64 = 3.14159
                                                    _ = v23
                                                }
                                            }
                                        case 32767:
                                            if (2147483648 <= 2147483648) {
                                                var v24 uint8 = (v2 | v2)
                                                _ = v24
                                            } else {
                                                var v25 float32 = -2.25
                                                switch fi11 {
                                                case 0:
                                                    v26 := GenS4{F0: v3, F1: ("a" + ("a" + (("\n\t" + ("a" + "a")) + "foo bar")))}
                                                    v27 := GenS4{F0: (!((!v3) && (-413 <= 42))), F1: (("" + (("a" + ("" + "\n\t")) + "foo bar")) + "foo bar")}
                                                    _ = v26
                                                    _ = v27
                                                case -2147483648:
                                                    v28 := GenS1{F0: 42, F1: 18446744073709551615}
                                                    _ = v28
                                                default:
                                                    _ = len(v12)
                                                }
                                                _ = v25
                                            }
                                        }
                                    }
                                } else {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                            }
                            var v29 string = "hello"
                            _ = v29
                        }
                        var v30 uint = 2
                        _ = v14
                        _ = v30
                    }
                    if len(v12) > 0 {
                        v12[0] = -774
                    }
                }
                _ = fi13
            }
        }
        _ = fi11
        _ = v12
    }
    switch ("a" + "") {
    case "":
        var v31 uint32 = 98
        _ = v31
    case "a":
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS5{F0: fi32}
            var v34 float32 = (1e10 * 3.14159)
            _ = fi32
            _ = v33
            _ = v34
        }
    case "\n\t":
        v1 = v1
    default:
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 float64 = 0.1
            _ = fi35
            _ = v36
        }
    }
    var v37 float32 = 3.14159
    v38 := GenS6{F0: 42, F1: 42}
    var v39 int16 = 0
    v40 := GenS6{F0: 657, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
