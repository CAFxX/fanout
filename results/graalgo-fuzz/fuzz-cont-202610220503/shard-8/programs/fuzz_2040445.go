package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 128
    v4 := []int{}
    v5 := v1[631]
    for fi6 := 0; fi6 < 4; fi6++ {
        switch -152 {
        case -1:
            v3 -= (v3 + (v3 | v3))
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v3 = (v3 - v3)
        var v8 float64 = (3.14159 - (1.5 / (float64((0.0 / float32(((float64((-2.25 - 1.5)) + -2.25) / 1.0)))) + float64((float32((-2.25 * (0.0 + 3.14159))) + (0.1 + 3.14159))))))
        var v9 uint8 = v3
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v4)
    switch (-(v5)) {
    case 440:
        v10 := GenS1{F0: 844}
        _ = v10
    case -945:
        switch v5 {
        case 9223372036854775807:
            v3 = (v3 ^ (v3 + 128))
        case 1:
            var v11 int32 = 2147483647
            _ = v11
        default:
            if false {
                v3 += 255
                switch (v5 + (v5 >> ((942) & 63))) {
                case -2147483648:
                    if (true && (0 <= -9223372036854775808)) {
                        fmt.Println(v5)
                        _ = len(v2)
                        v5 = 0
                    }
                default:
                    switch "hello" {
                    case "foo bar":
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v13 := GenS1{F0: 2147483647}
                            _ = fi12
                            _ = v13
                        }
                    case "\n\t":
                        fmt.Println(v3, v3)
                    default:
                        v14 := GenS1{F0: 42}
                        _ = v14
                    }
                }
            }
        }
    case 32768:
        v15 := GenS1{F0: 2147483647}
        _ = v15
    default:
        v3 = v3
    }
    v16 := GenS2{F0: 255, F1: 42}
    var v17 uint8 = v3
    v5 = (v5 + (v5 % ((-9223372036854775808) | 1)))
    v4 = append(v4, (^(v5)))
    if false {
        var v18 string = "\n\t"
        _ = len(v4)
        _ = v18
    } else {
        if true {
            for fi19 := 0; fi19 < 16; fi19++ {
                v1[0] = fi19
                fmt.Println(v5)
                _ = fi19
            }
        } else {
            switch "a" {
            case "":
                for fi20 := 0; fi20 < 7; fi20++ {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int = v5
                        v3 = (v17 << ((0) & 63))
                        _ = fi21
                        _ = v22
                    }
                    if false {
                        fmt.Println(v3)
                    } else {
                        switch -2147483648 {
                        case -2147483648:
                            v17 = v17
                        case 32768:
                            delete(v2, 2147483647)
                        default:
                            var v23 uint8 = (v3 % (((v3 ^ (v17 >> ((65535) & 63)))) | 1))
                            _ = v23
                        }
                    }
                    fmt.Println(fi20)
                    _ = fi20
                }
            }
            var v24 uint = 830
            v17 = 128
            _ = v24
        }
    }
    v2[65535] = 335
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v17, v3)
    if (65535 > 18446744073709551615) {
        var v25 uint16 = 42
        _ = v25
    } else {
        var v26 int16 = -32768
        if false {
            switch (v5 << ((0) & 63)) {
            case 42:
                if ((!true) && (18446744073709551615 < 65535)) {
                    v3 -= v3
                    for fi27 := 0; fi27 < 20; fi27++ {
                        _ = len(v1)
                        _ = fi27
                    }
                }
            case 0:
                var v28 uint = 17
                _ = v28
            case -881:
                var v29 float64 = 1.5
                _ = v29
            }
            v3 = (v17 | (v17 * (^(v3))))
            if false {
                v30 := GenS3{F0: 103, F1: (v3 / (((v3 ^ v3)) | 1))}
                var v31 float32 = 3.14159
                switch "foo bar" {
                case "a":
                    v4 = append(v4, (v5 + (v5 + (v5 >> ((1) & 63)))))
                case "hello":
                    if false {
                        if (("hello" + "\n\t") >= "a") {
                            v31 += v31
                        }
                    }
                case "foo bar":
                    var v32 uint = 42
                    _ = v32
                }
                _ = v30
                _ = v31
            }
        } else {
            if (((v5 / (((v5 ^ v5)) | 1)) >= 975) && false) {
                v33 := GenS2{F0: 65535, F1: 0}
                _ = v33
            }
        }
        _ = v26
    }
    delete(v1, 20)
    if true {
        v34 := GenS4{F0: 0}
        v5 = 129
        _ = v34
    } else {
        if true {
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 float32 = -2.25
                _ = fi35
                _ = v36
            }
            delete(v1, 132)
        }
        v17 -= v3
        var v37 int = (v5 >> ((2) & 63))
        _ = v37
    }
    var v38 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
