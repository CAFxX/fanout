package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32; F1 int64; F2 uint8 }

func main() {
    var v1 uint8 = 1
    var v2 float64 = -2.25
    var v3 float64 = ((3.14159 + 1.5) * v2)
    var v4 int32 = -1
    var v5 int = 9223372036854775807
    var v6 int32 = v4
    for fi7 := 0; fi7 < 6; fi7++ {
        if true {
            var v8 int8 = 112
            if false {
                var v9 uint = 287
                _ = v9
            } else {
                v10 := []int{(fi7 | (fi7 + 2147483647)), v5, v5, (v5 % (((fi7 ^ fi7)) | 1))}
                var v11 uint64 = 565
                _ = v10
                _ = v11
            }
            _ = v8
        } else {
            v4 += v6
            switch (v5 << ((4294967295) & 63)) {
            case 128:
                fmt.Println(v3)
            case -2147483648:
                if false {
                    if (4294967295 > 18446744073709551615) {
                        v12 := GenS1{F0: 265, F1: 1000, F2: (^(v1))}
                        _ = v12
                    }
                    if ("\n\t" <= "hello") {
                        v6 = v4
                        fmt.Println(v6, v6)
                        v5 = (v5 << ((4294967295) & 63))
                    } else {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            if ((^(v1)) == (v1 >> ((65535) & 63))) {
                                var v14 bool = (!((true || (0.0 <= (1.5 + 0.0))) || (0 > (v6 << ((683) & 63)))))
                                _ = v14
                            }
                            v15 := GenS2{F0: 275, F1: 42}
                            _ = fi13
                            _ = v15
                        }
                    }
                    var v16 uint32 = 2147483648
                    _ = v16
                }
            case 9223372036854775807:
                var v17 int8 = -128
                _ = v17
            default:
                v1 = v1
            }
            var v18 int16 = 1000
            _ = v18
        }
        var v19 bool = true
        var v20 float32 = 3.14159
        _ = fi7
        _ = v19
        _ = v20
    }
    var v21 uint8 = (v1 + (v1 - 222))
    switch ("hello" + "a") {
    case "":
        v22 := GenS1{F0: 1, F1: 94, F2: 63}
        _ = v22
    case "foo bar":
        v23 := []int{int(v1), 628, v5, (v5 << ((18446744073709551615) & 63))}
        _ = v23
    case "a":
        fmt.Println(v2, v2, v2)
    default:
        fmt.Println(v21, v1, v1)
    }
    v24 := GenS3{F0: 65535, F1: true}
    fmt.Println(v3, v5, v4)
    var v25 uint64 = 0
    fmt.Println(v2, v5)
    var v26 string = (("a" + "hello") + "foo bar")
    v25 = v25
    v6 = v4
    if true {
        fmt.Println(v4, v1, v6)
    }
    if (true || true) {
        v26 = "a"
        var v27 uint64 = v25
        fmt.Println(v26, v4, v2)
        _ = v27
    } else {
        v28 := []int{(v5 - (v5 << ((65535) & 63))), -2147483648, (v5 + (v5 & 32768)), 1}
        var v29 uint8 = 22
        v30 := GenS4{F0: (-(v4)), F1: -1, F2: v21}
        _ = v28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        switch v26 {
        case "\n\t":
            var v32 float64 = 0.1
            _ = v32
        case "hello":
            v33 := GenS2{F0: 1, F1: 703}
            _ = v33
        }
        v34 := GenS1{F0: 0, F1: 372, F2: v1}
        var v35 uint16 = 255
        _ = fi31
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
