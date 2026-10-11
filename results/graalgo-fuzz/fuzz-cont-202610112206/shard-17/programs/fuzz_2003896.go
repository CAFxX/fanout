package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 int16 = 1
    var v2 int32 = -452
    var v3 bool = (100 >= 42)
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint8 = 42
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float64 = 0.0
        _ = fi7
        _ = v8
    }
    if false {
        var v9 int16 = 0
        fmt.Println(v3, v3, v1)
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{fi10, 128, -9223372036854775808, int(v1)}
        _ = fi10
        _ = v11
    }
    switch "" {
    case "\n\t":
        v3 = false
    case "hello":
        var v12 bool = (520 != 32767)
        _ = v12
    default:
        v1 = 526
    }
    var v13 uint64 = 1
    var v14 float32 = 0.0
    var v15 int = -1
    var v16 uint = 731
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint16 = 781
        switch "hello" {
        case "hello":
            switch "foo bar" {
            case "hello":
                v19 := []int{v15}
                _ = v19
            case "":
                if v3 {
                    v18 = 2
                } else {
                    v20 := []int{}
                    var v21 uint8 = 255
                    _ = v20
                    _ = v21
                }
            }
        case "":
            v22 := GenS1{F0: -1, F1: -128, F2: ""}
            _ = v22
        }
        if (!v3) {
            fmt.Println(v15)
        }
        _ = fi17
        _ = v18
    }
    if v3 {
        if v3 {
            var v23 uint64 = uint64(v16)
            _ = v23
        }
        v24 := GenS1{F0: 127, F1: 42, F2: ("a" + "")}
        switch "" {
        case "foo bar":
            fmt.Println(v2)
        }
        _ = v24
    } else {
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS2{F0: "a", F1: (v1 ^ int16(v15)), F2: v16}
            v27 := []int{v15, (-(fi25)), 2147483647}
            if (v13 < (v13 + 18446744073709551615)) {
                switch "foo bar" {
                case "hello":
                    var v28 string = (("\n\t" + "foo bar") + ("\n\t" + "\n\t"))
                    _ = v28
                case "\n\t":
                    if (!(537 >= 2)) {
                        v13 += (v13 | (v13 * (v13 - (v13 - (^(v13))))))
                    }
                default:
                    if (v3 && true) {
                        switch fi25 {
                        case 681:
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        case 4294967295:
                            if true {
                                v29 := GenS2{F0: "", F1: v1, F2: (v16 << ((v16) & 63))}
                                var v30 uint64 = v13
                                _ = v29
                                _ = v30
                            }
                        default:
                            v31 := GenS3{F0: (-(v1)), F1: 668}
                            _ = v31
                        }
                        v32 := GenS4{F0: v16, F1: 349}
                        _ = v32
                    }
                }
                var v33 uint = (v16 / ((v16) | 1))
                var v34 float32 = 0.1
                _ = v33
                _ = v34
            } else {
                v35 := GenS2{F0: "\n\t", F1: v1, F2: v16}
                v36 := GenS3{F0: 42, F1: 0}
                _ = v35
                _ = v36
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v1 -= (v1 % (((v1 | (v1 - v1))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
