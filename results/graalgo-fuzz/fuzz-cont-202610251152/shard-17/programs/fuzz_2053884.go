package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint = 255
    v2 := []int{127, 42, -2147483648, -9223372036854775808}
    var v3 int32 = 544
    var v4 uint8 = 213
    v5 := []int{32768, 0, 32767, 4294967295}
    v3 += -2147483648
    switch "a" {
    case "a":
        var v6 uint32 = 65535
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                if true {
                    v9 := GenS1{F0: 100}
                    var v10 bool = true
                    _ = v10
                    _ = v9
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v1 = (^(v1))
                v1 = v1
                _ = len(v5)
                _ = fi8
            }
            _ = fi7
        }
    case "":
        var v11 bool = ((749 >= 249) && ((v4 << ((12) & 63)) > (v4 / ((128) | 1))))
        _ = v11
    }
    if (!true) {
        fmt.Println(v3, v4, v4)
        var v12 uint16 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
    }
    v13 := GenS2{F0: -664, F1: 429, F2: 345}
    if (false || false) {
        for fi14 := 0; fi14 < 7; fi14++ {
            if (true && (0 != 1)) {
                switch int(v1) {
                case 9223372036854775807:
                    v3 = (v3 - (v3 >> ((v1) & 63)))
                }
                switch (fi14 - fi14) {
                case 32767:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v2 = append(v2, (fi14 >> (((v1 - (v1 ^ v1))) & 63)))
                        _ = fi15
                    }
                }
                for fi16 := 0; fi16 < 11; fi16++ {
                    v4 -= (v4 / (((v4 & v4)) | 1))
                    _ = fi16
                }
            } else {
                _ = len(v5)
            }
            v17 := GenS3{F0: 128, F1: 4294967295, F2: -9223372036854775808}
            if true {
                var v18 int = (^(fi14))
                _ = v18
            } else {
                if true {
                    var v19 uint8 = 42
                    if true {
                        var v20 uint = v1
                        _ = v20
                    }
                    fmt.Println(v3, fi14)
                    _ = v19
                } else {
                    v21 := GenS4{F0: 1}
                    _ = v21
                }
                v3 = v3
            }
            _ = fi14
            _ = v17
        }
    }
    if (((^(v4)) <= (v4 - 42)) || (false || ((((!true) && false) && false) && (v4 >= (v4 >> ((0) & 63)))))) {
        _ = len(v2)
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    switch ("hello" + "a") {
    case "foo bar":
        var v22 int16 = 1000
        _ = v22
    case "hello":
        if ((false && false) && false) {
            v23 := GenS1{F0: 100}
            if len(v5) > 0 {
                v5[0] = 42
            }
            var v24 float64 = ((3.14159 / 1.0) + (0.1 * 0.1))
            _ = v23
            _ = v24
        }
    default:
        v25 := GenS4{F0: 42}
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v3 = (-(v3))
        _ = fi26
    }
    v27 := GenS3{F0: 127, F1: 2147483647, F2: -1}
    switch "a" {
    case "\n\t":
        var v28 int8 = 70
        _ = v28
    }
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    var v29 int16 = 32767
    v4 -= v4
    var v30 int8 = 75
    if false {
        v31 := GenS5{F0: 255}
        fmt.Println(v4, v29)
        _ = v31
    } else {
        if (123 >= 65535) {
            v32 := GenS6{F0: 32767, F1: v30}
            var v33 int8 = 0
            _ = v32
            _ = v33
        }
    }
    var v34 float32 = (((1.5 * 1.5) + 1.5) / float32(((1e10 - (1.5 / 1.0)) - (((1e10 - 0.0) + 1e10) * (-2.25 * (-2.25 + 1.5))))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
