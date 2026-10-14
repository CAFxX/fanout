package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }

func main() {
    var v1 float32 = 1e10
    v2 := []uint32{2147483648}
    var v3 uint16 = 259
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 string = "foo bar"
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: v3, F1: 536, F2: 4294967295}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    v7 := GenS1{F0: (^(v3)), F1: 4294967295, F2: 0}
    var v8 int64 = 0
    var v9 uint = 682
    v1 = float32(((float64(v1) + (((3.14159 * (1e10 * -2.25)) / (float64(v1) - ((1.5 - (float64(v1) / 0.1)) + 1.5))) / float64(v1))) + -2.25))
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: v3, F1: 18446744073709551615, F2: uint64(v9)}
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint8 = 2
            fmt.Println(fi10, v8, fi12)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint = v9
    if (0 > 173) {
        v1 -= v1
        if false {
            v8 -= 9223372036854775807
            v2 = append(v2, 2147483648)
        } else {
            var v15 uint16 = v3
            _ = v15
        }
    } else {
        v9 = v14
        v2 = append(v2, 219)
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint32 = 65535
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "a" {
        case "hello":
            var v19 int16 = 983
            _ = v19
        case "":
            switch ("\n\t" + (("\n\t" + ("hello" + ("hello" + "foo bar"))) + ("foo bar" + "hello"))) {
            case "\n\t":
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS1{F0: (^(v3)), F1: 2, F2: 411}
                    _ = fi20
                    _ = v21
                }
            case "":
                switch "foo bar" {
                case "":
                    switch (("a" + "hello") + (("foo bar" + ("foo bar" + "foo bar")) + (("hello" + ((("" + "") + "a") + "hello")) + ("\n\t" + "foo bar")))) {
                    case "hello":
                        var v22 int8 = 42
                        _ = v22
                    }
                case "foo bar":
                    v23 := GenS1{F0: v3, F1: 18446744073709551615, F2: 4294967295}
                    _ = v23
                case "\n\t":
                    fmt.Println(v3, fi18, v8)
                }
            default:
                if ((!true) && true) {
                    fmt.Println(v9)
                }
            }
        default:
            if true {
                v14 += 701
            } else {
                switch "" {
                case "\n\t":
                    if (((0.0 / -2.25) + -2.25) != float64((v1 / v1))) {
                        for fi24 := 0; fi24 < 2; fi24++ {
                            fmt.Println(fi24)
                            _ = fi24
                        }
                    }
                case "":
                    if (v8 != -475) {
                        v9 = (v14 >> ((v14) & 63))
                    } else {
                        v14 -= v14
                    }
                case "a":
                    var v25 uint64 = 42
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 int64 = v8
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
            _ = fi26
            _ = v27
        }
        _ = fi18
    }
    var v28 uint = v14
    switch "foo bar" {
    case "a":
        for fi29 := 0; fi29 < 6; fi29++ {
            v14 -= v28
            _ = fi29
        }
    case "\n\t":
        fmt.Println(v8, v3, v3)
    case "hello":
        _ = len(v2)
    }
    var v30 int64 = v8
    if (uint64(v3) < 18446744073709551615) {
        if len(v2) > 0 {
            v2[0] = 2
        }
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
