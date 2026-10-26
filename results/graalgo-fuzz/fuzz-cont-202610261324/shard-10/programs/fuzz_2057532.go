package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int64 = 485
    var v2 uint16 = 897
    v3 := make(map[string]int)
    var v4 float32 = 0.0
    if (!(-2147483648 > 0)) {
        if (true && false) {
            v3["hello"] = -915
        } else {
            v5 := GenS1{F0: 42, F1: 42}
            if true {
                v4 = ((3.14159 + (-2.25 + ((v4 - ((v4 + 0.0) * (3.14159 + v4))) + 3.14159))) - (0.1 * v4))
            } else {
                v1 = v1
            }
            var v6 uint = 65535
            _ = v5
            _ = v6
        }
        switch ("\n\t" + "foo bar") {
        case "":
            v7 := []int{}
            _ = v7
        case "\n\t":
            v1 = v1
        default:
            for fi8 := 0; fi8 < 16; fi8++ {
                v1 -= (v1 / (((v1 * v1)) | 1))
                _ = fi8
            }
        }
    }
    fmt.Println(v4)
    if (!false) {
        v9 := []int{857, 2147483647, 2147483647}
        v10 := GenS1{F0: 2147483648, F1: 12}
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                v4 = v4
                _ = fi12
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    if (97 > 0) {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{(fi13 - fi13), (fi14 / (((-(fi14))) | 1)), 32767}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        if (4294967295 >= -1) {
            fmt.Println(v1, v4)
            if ((2 <= 18446744073709551615) || (255 >= 255)) {
                v16 := []int{}
                if true {
                    v1 = int64(v2)
                }
                fmt.Println(v4, v1)
                _ = v16
            } else {
                switch "hello" {
                case "hello":
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v1 = (v1 * v1)
                        _ = fi17
                    }
                }
                v3["\n\t"] = -659
            }
        } else {
            v3[(("a" + ("foo bar" + "a")) + (("hello" + ("foo bar" + "foo bar")) + "\n\t"))] = 32767
            v18 := []int{631, -370}
            v1 -= (v1 * (v1 + v1))
            _ = v18
        }
        var v19 bool = (0.1 > v4)
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v2, fi20, v4)
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    v2 -= (v2 & v2)
    fmt.Println(v2, v4, v1)
    v22 := v3[("\n\t" + "a")]
    fmt.Println(v4, v1)
    v23 := GenS1{F0: 2147483647, F1: 457}
    for fi24 := 0; fi24 < 6; fi24++ {
        v22 = v22
        v25 := GenS1{F0: 2147483648, F1: uint(v2)}
        _ = fi24
        _ = v25
    }
    if false {
        var v26 int16 = 228
        for fi27 := 0; fi27 < 17; fi27++ {
            if true {
                for fi28 := 0; fi28 < 7; fi28++ {
                    var v29 uint8 = 56
                    v22 -= 127
                    _ = fi28
                    _ = v29
                }
                fmt.Println(v4, v4)
                for fi30 := 0; fi30 < 16; fi30++ {
                    v22 += -675
                    v31 := []int{(fi27 << ((0) & 63))}
                    _ = fi30
                    _ = v31
                }
            } else {
                v32 := []int{v22, fi27}
                if false {
                    _ = len(v3)
                }
                _ = v32
            }
            v33 := GenS2{F0: 900}
            v22 = 32767
            _ = fi27
            _ = v33
        }
        v3["foo bar"] = -9223372036854775808
        _ = v26
    } else {
        if true {
            switch "" {
            case "a":
                if (true || false) {
                    v34 := GenS3{F0: 127, F1: 42}
                    _ = v34
                } else {
                    fmt.Println(v1, v1, v1)
                    v35 := GenS3{F0: 42, F1: v1}
                    _ = v35
                }
            }
            v36 := []int{42, (v22 >> ((2) & 63)), 679, 128}
            _ = v36
        } else {
            v37 := GenS1{F0: 0, F1: 484}
            _ = v37
        }
    }
    var v38 float32 = v4
    delete(v3, ("hello" + "\n\t"))
    var v39 uint16 = v2
    switch v22 {
    case 0:
        v4 += (((v38 / v4) - v4) / v38)
    case 127:
        fmt.Println(v38, v2)
    }
    switch ("\n\t" + ("" + "foo bar")) {
    case "\n\t":
        for fi40 := 0; fi40 < 2; fi40++ {
            v22 += -9223372036854775808
            _ = fi40
        }
    case "hello":
        v22 = (v22 | (v22 + (v22 ^ -2147483648)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
