package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]string)
    var v3 uint8 = 42
    v1 -= -2147483648
    switch -1 {
    case 193:
        var v4 int8 = 42
        _ = v4
    }
    v5 := []int{-1, 127, -872}
    switch 4294967295 {
    case -9223372036854775808:
        var v6 uint32 = 1
        _ = v6
    case 609:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 0:
        v5 = append(v5, -2147483648)
    default:
        var v7 uint = 2
        _ = v7
    }
    fmt.Println(v3)
    v1 += -1
    v8 := v2[0]
    v8 = (("foo bar" + (v8 + ((v8 + "foo bar") + ((((("" + ((v8 + (("" + (v8 + ("foo bar" + "\n\t"))) + v8)) + (v8 + v8))) + v8) + "hello") + v8) + (v8 + "a"))))) + (v8 + ("foo bar" + (v8 + ((v8 + ("\n\t" + v8)) + (v8 + (v8 + ((v8 + "a") + v8))))))))
    var v9 uint = 255
    v3 += v3
    v8 = "\n\t"
    v10 := v2[-9223372036854775808]
    v1 -= v1
    v8 = ("" + (("foo bar" + ("a" + v10)) + (((((((v8 + v8) + v10) + v8) + v8) + v10) + v8) + v8)))
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 bool = (42 > (v9 ^ v9))
        for fi13 := 0; fi13 < 9; fi13++ {
            if v12 {
                v5 = append(v5, fi11)
                _ = len(v5)
                _ = len(v5)
                var v14 uint = (v9 | (v9 / ((371) | 1)))
                if true {
                    v15 := GenS1{F0: 365, F1: 18446744073709551615, F2: -11}
                    _ = v15
                } else {
                    switch (fi11 >> ((v14) & 63)) {
                    case 32768:
                        var v16 string = "\n\t"
                        v9 = v14
                        _ = v16
                    case 32767:
                        switch (v10 + (v10 + v8)) {
                        case "":
                            v17 := GenS2{F0: v14, F1: -1}
                            _ = v17
                        case "\n\t":
                            v2[(fi11 ^ (fi11 % ((fi13) | 1)))] = "a"
                            v1 -= (v1 % ((42) | 1))
                        case "a":
                            var v18 uint64 = uint64(v3)
                            _ = v18
                        default:
                            v19 := GenS1{F0: 1, F1: 833, F2: -15}
                            _ = v19
                        }
                    case 1:
                        v20 := GenS2{F0: v9, F1: int64(v3)}
                        _ = v20
                    }
                    delete(v2, (fi11 >> ((7) & 63)))
                    if v12 {
                        fmt.Println(v14, v9)
                    }
                }
                fmt.Println(v10)
                _ = v14
            }
            var v21 int32 = v1
            var v22 int16 = -32768
            _ = fi13
            _ = v21
            _ = v22
        }
        fmt.Println(v12)
        _ = fi11
        _ = v12
    }
    fmt.Println(v1)
    switch 397 {
    case -829:
        v10 = v8
    default:
        if true {
            if (true && (!(-2.25 == 0.0))) {
                v23 := GenS2{F0: (v9 << (((v9 << (((v9 / ((65535) | 1))) & 63))) & 63)), F1: 1}
                if len(v5) > 0 {
                    v5[0] = 127
                }
                var v24 bool = (v9 != 425)
                _ = v23
                _ = v24
            } else {
                if true {
                    var v25 int32 = v1
                    _ = v25
                } else {
                    v8 = v10
                    if true {
                        var v26 float32 = 3.14159
                        var v27 int16 = 831
                        _ = v26
                        _ = v27
                    }
                }
                v1 -= v1
            }
            var v28 uint16 = 32767
            _ = v28
        } else {
            var v29 int16 = 0
            switch ((v10 + (("" + "hello") + "foo bar")) + v8) {
            case "a":
                v30 := v2[32767]
                _ = v30
            case "hello":
                v8 = "foo bar"
            case "foo bar":
                v8 = (v10 + "hello")
            default:
                for fi31 := 0; fi31 < 11; fi31++ {
                    _ = len(v5)
                    v3 = 127
                    _ = fi31
                }
            }
            v1 = (v1 & v1)
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
