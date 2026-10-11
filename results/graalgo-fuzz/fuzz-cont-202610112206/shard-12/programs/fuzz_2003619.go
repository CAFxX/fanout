package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0() uint64 {
    var v5 float32 = (3.14159 * (float32(3.14159) * 1e10))
    if (685 == 42) {
        v5 = v5
    }
    _ = v5
    return 0
}

func main() {
    var v1 int = 1
    var v2 uint = 0
    v3 := []string{}
    v4 := make(map[uint32]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint8 = 255
        v3 = append(v3, "a")
        v8 := GenS1{F0: 0, F1: uint16(v7)}
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 = v2
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int16 = -1
        _ = fi9
        _ = v10
    }
    if true {
        fmt.Println(v1, v1)
        v1 -= (v1 * (v1 >> (((v2 * v2)) & 63)))
    }
    if ((4294967295 <= 0) && true) {
        v11 := GenS1{F0: -2147483648, F1: 828}
        _ = v11
    }
    switch ("\n\t" + "a") {
    case "a":
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint32 = 0
            var v14 uint32 = (v13 * 2147483647)
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS2{F0: 2}
                var v17 int32 = 2147483647
                delete(v4, v14)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    case "foo bar":
        var v18 float32 = 3.14159
        _ = v18
    default:
        if false {
            switch ((("" + "foo bar") + ("hello" + (("a" + "") + "\n\t"))) + "foo bar") {
            case "hello":
                v1 -= (^(v1))
            default:
                v2 = v2
            }
        } else {
            var v19 bool = ((float32(0.0) <= 1.5) || false)
            v2 -= (v2 & v2)
            if ((2147483648 != 65535) && ((float32(((3.14159 - -2.25) * 3.14159)) - float32((0.1 - 0.0))) == (1e10 * 0.1))) {
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 float32 = (-2.25 / (0.1 / float32(((1.5 + 0.1) + 3.14159))))
                    v22 := GenS1{F0: -2147483648, F1: 681}
                    var v23 int64 = int64(v1)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v2)
            } else {
                v3 = append(v3, (("a" + ((((("foo bar" + "foo bar") + "hello") + "hello") + "") + ("\n\t" + ("a" + "\n\t")))) + ("hello" + ("a" + "foo bar"))))
                switch 128 {
                case -99:
                    v19 = v19
                case 32767:
                    v24 := GenS1{F0: -939, F1: 65535}
                    _ = v24
                default:
                    switch (^(v1)) {
                    case 42:
                        var v25 int64 = 1
                        _ = v25
                    default:
                        if ((108 <= 1) && (32767 <= 756)) {
                            _ = len(v3)
                            var v26 float32 = 1e10
                            var v27 string = ("\n\t" + "\n\t")
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS3{F0: 0, F1: (2 == 42), F2: 42}
                            if (577 == 972) {
                                v1 = (v1 << ((v2) & 63))
                            }
                            v2 -= v2
                            _ = v28
                        }
                    }
                }
            }
            _ = v19
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v3 = append(v3, "foo bar")
        v1 = v1
        var v30 string = ("foo bar" + "\n\t")
        _ = fi29
        _ = v30
    }
    if (851 > v2) {
        v1 += v1
        fmt.Println(v2, v2)
        delete(v4, 2147483648)
    } else {
        var v31 float64 = ((3.14159 + 0.0) / 1.0)
        _ = v31
    }
    v32 := GenS1{F0: 0, F1: 65535}
    v1 += (v1 & 9223372036854775807)
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint16 = 776
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
