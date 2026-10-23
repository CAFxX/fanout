package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 bool, p1 int8, p2 int64) uint {
    v4 := GenS1{F0: p1, F1: 65535, F2: 538}
    switch -1 {
    case 392:
        v5 := []int{2147483647, 9223372036854775807, -150}
        _ = v5
    default:
        if p0 {
            for fi6 := 0; fi6 < 8; fi6++ {
                if ((3.14159 + ((-2.25 / 1.0) + (0.1 - 1.5))) == 3.14159) {
                    var v7 float64 = (3.14159 / 1e10)
                    _ = v7
                }
                if (float32(1.5) > float32((3.14159 * -2.25))) {
                    var v8 int8 = (-(p1))
                    _ = v8
                }
                switch "a" {
                case "hello":
                    var v9 int = -949
                    _ = v9
                case "foo bar":
                    if false {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            fmt.Println(p1, fi6, p2)
                            _ = fi10
                        }
                        for fi11 := 0; fi11 < 13; fi11++ {
                            switch "hello" {
                            case "hello":
                                p1 -= p1
                            case "\n\t":
                                var v12 uint8 = 42
                                _ = v12
                            }
                            p2 += (p2 | (p2 / ((p2) | 1)))
                            var v13 uint8 = 127
                            _ = fi11
                            _ = v13
                        }
                    }
                case "":
                    v14 := GenS1{F0: (p1 << ((2) & 63)), F1: 1, F2: 2147483647}
                    _ = v14
                default:
                    var v15 int8 = (p1 >> ((18446744073709551615) & 63))
                    _ = v15
                }
                _ = fi6
            }
            fmt.Println(p1, p0)
        }
    }
    var v16 bool = (0 > 4294967295)
    if v16 {
        for fi17 := 0; fi17 < 11; fi17++ {
            switch "" {
            case "":
                for fi18 := 0; fi18 < 15; fi18++ {
                    switch (fi17 + 2147483647) {
                    case 4294967295:
                        v19 := GenS1{F0: p1, F1: 2, F2: 42}
                        _ = v19
                    }
                    _ = fi18
                }
            }
            v20 := []int{fi17}
            _ = fi17
            _ = v20
        }
    }
    fmt.Println(v16, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v4
    return 145
}

func main() {
    var v1 int16 = 0
    var v2 float32 = (1.5 / (3.14159 / ((1e10 * (((((((3.14159 + (0.1 / 3.14159)) + 0.1) + (float32(1.5) + (-2.25 / 0.1))) + float32(((0.0 - 3.14159) + 3.14159))) - float32(-2.25)) + 3.14159) + (1.5 + float32(3.14159)))) * 3.14159)))
    var v3 int32 = 456
    if true {
        switch "" {
        case "foo bar":
            var v21 uint32 = 777
            _ = v21
        default:
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: 73, F1: 65535, F2: 530}
                for fi24 := 0; fi24 < 3; fi24++ {
                    v2 += ((v2 / ((1.5 - 1e10) - v2)) / float32(1.5))
                    _ = fi24
                }
                v25 := []int{fi22, (^(fi22)), (fi22 << ((2) & 63))}
                _ = fi22
                _ = v23
                _ = v25
            }
        }
        v26 := GenS2{F0: -1, F1: 2, F2: 1}
        _ = v26
    } else {
        switch (("" + ("a" + "hello")) + "foo bar") {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                v27 := []int{}
                _ = v27
            case "\n\t":
                if true {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        v29 := GenS2{F0: 100, F1: 2, F2: 42}
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v3, v1)
                }
            case "":
                v3 -= v3
            }
        case "hello":
            fmt.Println(v2, v2, v1)
        case "a":
            var v30 int8 = -1
            _ = v30
        default:
            var v31 int64 = -9223372036854775808
            _ = v31
        }
        if true {
            v32 := GenS2{F0: 127, F1: 678, F2: -9223372036854775808}
            v2 += -2.25
            _ = v32
        }
        v2 = (0.0 - v2)
    }
    v3 = (v3 - v3)
    v1 += v1
    var v33 float64 = ((0.1 / float64(1.5)) - 1e10)
    var v34 uint32 = 2
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := GenS3{F0: -32768, F1: 127, F2: ""}
        _ = fi35
        _ = v36
    }
    if ((true || (false && (true && true))) || false) {
        v37 := []int{}
        v38 := GenS2{F0: 63, F1: 42, F2: 1}
        _ = v37
        _ = v38
    } else {
        switch 42 {
        case 130:
            var v39 string = "hello"
            _ = v39
        case 2147483647:
            v34 -= v34
        case 128:
            v40 := GenS3{F0: (v1 & (v1 ^ v1)), F1: -361, F2: "\n\t"}
            _ = v40
        default:
            var v41 uint16 = uint16(v34)
            _ = v41
        }
    }
    switch 128 {
    case 128:
        v34 = v34
    case -505:
        v3 = v3
    case 42:
        if (44 == 255) {
            v1 = v1
        }
    }
    var v42 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v42
}
