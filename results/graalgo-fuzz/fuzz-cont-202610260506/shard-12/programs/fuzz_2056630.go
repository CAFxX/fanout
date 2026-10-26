package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 string = "hello"
    var v2 int8 = 0
    var v3 int64 = 496
    if true {
        v3 = 42
    } else {
        switch int(v2) {
        case 4294967295:
            switch v1 {
            case "foo bar":
                v4 := GenS1{F0: (^(v2))}
                _ = v4
            default:
                fmt.Println(v3, v2)
            }
        case 128:
            v5 := GenS2{F0: (v2 / (((v2 % ((127) | 1))) | 1)), F1: "foo bar"}
            _ = v5
        default:
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := []int{fi6, (fi6 ^ fi6), fi6, (fi6 / ((-283) | 1))}
                for fi8 := 0; fi8 < 10; fi8++ {
                    v1 = v1
                    v1 = (v1 + v1)
                    _ = fi8
                }
                v3 = v3
                _ = fi6
                _ = v7
            }
        }
    }
    v3 -= (-(v3))
    if (true || false) {
        if (4294967295 != 287) {
            v3 = (v3 - (^(v3)))
        } else {
            v9 := []int{32767}
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{}
            switch -2147483648 {
            case -2147483648:
                v1 = v1
            case -976:
                var v12 int16 = -664
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        var v13 string = v1
        _ = v13
    } else {
        v1 = v1
        v14 := GenS1{F0: v2}
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: 120, F1: (("a" + (v1 + v1)) + v1)}
        switch ("" + "\n\t") {
        case "\n\t":
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                v17 := GenS2{F0: 1, F1: "\n\t"}
                _ = v17
            case "\n\t":
                var v18 int64 = 9223372036854775807
                _ = v18
            case "":
                if (713 >= 42) {
                    switch v1 {
                    case "hello":
                        v2 = (v2 * (v2 + (v2 % ((v2) | 1))))
                    case "\n\t":
                        var v19 uint8 = 2
                        _ = v19
                    case "a":
                        if false {
                            var v20 uint32 = 4294967295
                            _ = v20
                        }
                    }
                    v21 := []int{(-(fi15)), 9223372036854775807, fi15}
                    _ = v21
                } else {
                    var v22 int = fi15
                    v23 := GenS3{F0: 940}
                    _ = v22
                    _ = v23
                }
            default:
                var v24 uint16 = 681
                _ = v24
            }
        case "":
            switch fi15 {
            case 4294967295:
                for fi25 := 0; fi25 < 2; fi25++ {
                    fmt.Println(fi25)
                    _ = fi25
                }
            }
        case "foo bar":
            switch ("hello" + ("foo bar" + "a")) {
            case "hello":
                v2 = v2
            case "a":
                fmt.Println(v3, v3)
            case "\n\t":
                var v26 int8 = 127
                _ = v26
            default:
                var v27 float32 = 1e10
                _ = v27
            }
        default:
            var v28 string = v1
            _ = v28
        }
        _ = fi15
        _ = v16
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 int64 = (v3 * v3)
        var v31 string = v1
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int32 = -953
    var v33 int64 = v3
    for fi34 := 0; fi34 < 14; fi34++ {
        v2 = v2
        v3 = (v3 + (v33 >> ((18446744073709551615) & 63)))
        _ = fi34
    }
    var v35 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
}
