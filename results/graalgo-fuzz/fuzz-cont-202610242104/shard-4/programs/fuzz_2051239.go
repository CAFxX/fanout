package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16) bool {
    fmt.Println(p0)
    v5 := []int{32767, 4294967295, 9223372036854775807}
    _ = p0
    _ = v5
    return (3.14159 < 3.14159)
}
func genFn1() uint {
    var v6 int32 = -957
    v7 := GenS1{F0: false, F1: 1000}
    v6 = v6
    var v8 float32 = 3.14159
    _ = v6
    _ = v7
    _ = v8
    return 562
}

func main() {
    var v1 int8 = 72
    var v2 uint8 = 2
    v3 := []int{}
    var v4 int8 = (v1 << ((488) & 63))
    if len(v3) > 0 {
        v3[0] = -503
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += (v1 / ((0) | 1))
    _ = len(v3)
    v3 = append(v3, -9223372036854775808)
    if false {
        switch (("hello" + "") + "") {
        case "hello":
            var v9 uint16 = 2
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int = (fi10 + 1)
                _ = fi10
                _ = v11
            }
        case "a":
            if (1 != 65535) {
                v4 -= (v1 >> ((18446744073709551615) & 63))
                v2 += 189
            } else {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: true, F1: -32768}
                    _ = fi12
                    _ = v13
                }
            }
        default:
            v4 -= (^(v4))
        }
        v2 += v2
        v1 = v1
    } else {
        v4 += -128
        v1 -= v4
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            if len(v3) > 0 {
                v3[0] = -9223372036854775808
            }
            _ = fi15
        }
        _ = fi14
    }
    if true {
        if (18446744073709551615 <= 2) {
            v16 := GenS1{F0: true, F1: 1000}
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 int16 = 42
                switch fi17 {
                case 602:
                    switch "" {
                    case "":
                        if (!(-505 > 2147483647)) {
                            v19 := GenS2{F0: v1, F1: 4294967295, F2: 9223372036854775807}
                            var v20 uint = 18446744073709551615
                            _ = v19
                            _ = v20
                        }
                    default:
                        for fi21 := 0; fi21 < 9; fi21++ {
                            switch "\n\t" {
                            case "":
                                var v22 string = "foo bar"
                                _ = v22
                            default:
                                var v23 uint8 = (v2 - 128)
                                _ = v23
                            }
                            var v24 int32 = -2147483648
                            switch (("a" + ("hello" + "hello")) + ("a" + ("foo bar" + ("foo bar" + "a")))) {
                            case "foo bar":
                                _ = len(v3)
                            case "a":
                                v25 := GenS3{F0: fi17}
                                _ = v25
                            case "\n\t":
                                var v26 uint8 = v2
                                _ = v26
                            }
                            _ = fi21
                            _ = v24
                        }
                    }
                case 241:
                    v27 := GenS2{F0: (v1 * (v1 >> ((65535) & 63))), F1: 2, F2: -1}
                    _ = v27
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v3 = append(v3, -2147483648)
            fmt.Println(v1)
            switch -20 {
            case 42:
                v3 = append(v3, (fi28 << ((18446744073709551615) & 63)))
            case 128:
                for fi29 := 0; fi29 < 15; fi29++ {
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    var v30 string = "\n\t"
                    fmt.Println(fi28, fi28, v4)
                    _ = fi29
                    _ = v30
                }
            default:
                v2 = (v2 ^ (^(v2)))
            }
            _ = fi28
        }
    }
    v31 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
