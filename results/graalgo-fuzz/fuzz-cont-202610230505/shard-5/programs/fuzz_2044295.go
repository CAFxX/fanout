package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0(p0 int32) int16 {
    if true {
        if (-2.25 <= (float32((-2.25 - 3.14159)) / ((1.5 / (1e10 / 1.0)) + 3.14159))) {
            for fi7 := 0; fi7 < 19; fi7++ {
                if (false || (false || (0.1 == (3.14159 / 3.14159)))) {
                    var v8 uint16 = 255
                    switch 32768 {
                    case 32767:
                        switch (("\n\t" + "a") + "foo bar") {
                        case "":
                            v8 -= v8
                        case "hello":
                            v8 -= v8
                        case "\n\t":
                            var v9 string = ("foo bar" + ((("" + ("a" + ("\n\t" + "hello"))) + (("a" + "\n\t") + "foo bar")) + "a"))
                            _ = v9
                        }
                    default:
                        fmt.Println(v8, v8)
                    }
                    _ = v8
                }
                if (((2147483648 == 2147483648) && false) || false) {
                    v10 := GenS1{F0: 42, F1: 1, F2: -9223372036854775808}
                    _ = v10
                }
                _ = fi7
            }
        }
        switch "" {
        case "hello":
            p0 = p0
        default:
            v11 := GenS2{F0: 0, F1: -727}
            _ = v11
        }
    } else {
        p0 = (p0 % (((p0 - (p0 / ((-2147483648) | 1)))) | 1))
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        p0 = p0
        _ = fi12
    }
    p0 = (p0 / (((p0 % ((-227) | 1))) | 1))
    _ = p0
    return 976
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]int)
    var v3 uint8 = 2
    var v4 uint8 = (v3 / ((128) | 1))
    v5 := make(map[int]uint32)
    v6 := []int{-1, 2147483647, 4294967295}
    switch 128 {
    case 2147483647:
        if (580 > 1) {
            var v13 int = 2147483647
            for fi14 := 0; fi14 < 18; fi14++ {
                v3 = v3
                var v15 int32 = 1
                var v16 int32 = 0
                _ = fi14
                _ = v15
                _ = v16
            }
            switch (v13 << ((18446744073709551615) & 63)) {
            case 2147483647:
                for fi17 := 0; fi17 < 9; fi17++ {
                    if false {
                        var v18 float32 = (float32(0.1) + 1.5)
                        var v19 bool = false
                        switch "" {
                        case "foo bar":
                            var v20 uint32 = 0
                            _ = v20
                        case "hello":
                            v13 = (v13 >> ((880) & 63))
                        }
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
            default:
                fmt.Println(v1, v4, v13)
            }
            _ = v13
        }
    }
    v21 := GenS3{F0: (2147483647 > 1), F1: 0, F2: 0}
    var v22 uint = 0
    var v23 uint = (v22 % ((v22) | 1))
    v4 -= v4
    v1 = v1
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v4 += v3
    var v24 int16 = 32767
    fmt.Println(v22)
    fmt.Println(v1, v24)
    if ((((((v1 + v1) + v1) + v1) + v1) + "\n\t") > v1) {
        var v25 uint8 = v4
        v26 := GenS2{F0: v24, F1: -9223372036854775808}
        _ = v25
        _ = v26
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
