package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int; F1 uint; F2 uint }
type GenS9 struct { F0 uint64; F1 int32 }

func main() {
    var v1 string = ("" + "a")
    var v2 uint8 = 128
    var v3 int16 = -32768
    v4 := make(map[int]string)
    v3 = v3
    v5 := GenS1{F0: v3, F1: v2}
    var v6 uint = 716
    v1 = (("" + ("hello" + v1)) + v1)
    v7 := []int{}
    v8 := GenS2{F0: 823}
    var v9 int = -9223372036854775808
    v10 := GenS3{F0: 481, F1: 50, F2: (v2 ^ v2)}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS4{F0: 18446744073709551615, F1: 42}
        var v13 uint8 = v2
        _ = fi11
        _ = v12
        _ = v13
    }
    v1 = (("foo bar" + (v1 + "foo bar")) + v1)
    var v14 uint64 = 85
    var v15 string = ""
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint16 = 398
            switch (fi16 >> ((v6) & 63)) {
            case -858:
                fmt.Println(v9)
            case 9223372036854775807:
                fmt.Println(v6, v18)
            case 2147483647:
                var v19 uint8 = uint8(v18)
                _ = v19
            }
            switch fi16 {
            case 107:
                var v20 int = (fi16 * (fi16 * (fi17 | fi16)))
                v1 = v1
                _ = v20
            case -2147483648:
                var v21 int16 = (v3 & (v3 | 1000))
                _ = v21
            case 9223372036854775807:
                _ = len(v4)
            }
            if false {
                if true {
                    _ = len(v4)
                    v22 := GenS5{F0: 100, F1: v2, F2: -13}
                    var v23 uint64 = (v14 + 125)
                    _ = v22
                    _ = v23
                } else {
                    _ = len(v4)
                    var v24 int = fi17
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v24
                }
            } else {
                switch fi17 {
                case 42:
                    v25 := GenS6{F0: v15}
                    _ = v25
                case 1:
                    if (!true) {
                        if false {
                            var v26 float64 = 1.5
                            v27 := GenS7{F0: (!((1 < 4294967295) && (-82 != 0))), F1: 65535}
                            _ = v26
                            _ = v27
                        }
                        var v28 int16 = v3
                        _ = v28
                    }
                }
                switch "\n\t" {
                case "":
                    v6 = v6
                }
            }
            v14 += v14
            if (!(true || (false && true))) {
                var v29 int = -814
                v18 = v18
                switch (v15 + v1) {
                case "\n\t":
                    v18 -= v18
                case "hello":
                    var v30 int8 = 0
                    _ = v30
                }
                var v31 uint32 = 639
                v14 -= (v14 | 0)
                _ = v29
                _ = v31
            } else {
                v32 := GenS2{F0: -789}
                _ = v32
            }
            if true {
                v33 := v4[(fi16 % (((fi16 / (((-(fi16))) | 1))) | 1))]
                _ = v33
            }
            var v34 uint64 = (v14 % (((v14 - 441)) | 1))
            if (9223372036854775807 != 434) {
                delete(v4, (fi17 + (fi16 / ((1) | 1))))
                v6 = v6
            } else {
                v4[v9] = "foo bar"
                v2 += 145
                v3 += v3
            }
            _ = fi17
            _ = v18
            _ = v34
        }
        var v35 int16 = 155
        fmt.Println(fi16, v2, v2)
        _ = fi16
        _ = v35
    }
    if (42 <= 0) {
        v7 = append(v7, 32767)
    }
    var v36 uint16 = 255
    if true {
        for fi37 := 0; fi37 < 4; fi37++ {
            if true {
                var v38 float32 = (1e10 + 1e10)
                v1 = (((v1 + v15) + (v15 + v1)) + (v1 + (v15 + "hello")))
                var v39 string = v15
                _ = v38
                _ = v39
            } else {
                v40 := GenS8{F0: -2147483648, F1: v6, F2: uint(v36)}
                _ = v40
            }
            if false {
                var v41 int64 = 1
                switch v15 {
                case "a":
                    var v42 float64 = (1.5 - 0.1)
                    _ = v42
                case "\n\t":
                    for fi43 := 0; fi43 < 13; fi43++ {
                        fmt.Println(v6, v9, v2)
                        _ = fi43
                    }
                }
                v36 = (v36 / ((v36) | 1))
                _ = v41
            } else {
                v44 := GenS1{F0: v3, F1: uint8(v14)}
                v3 = (-(v3))
                switch (v1 + ("hello" + (v15 + (v15 + (("hello" + v15) + "a"))))) {
                case "\n\t":
                    var v45 uint64 = uint64(v2)
                    _ = v45
                case "":
                    for fi46 := 0; fi46 < 16; fi46++ {
                        var v47 uint32 = 2147483647
                        switch v9 {
                        case -172:
                            v48 := GenS3{F0: (^(v6)), F1: 42, F2: 255}
                            _ = v48
                        }
                        if false {
                            v49 := GenS9{F0: uint64(v6), F1: 2147483647}
                            _ = v49
                        }
                        _ = fi46
                        _ = v47
                    }
                case "a":
                    fmt.Println(v36)
                }
                _ = v44
            }
            var v50 float64 = (0.1 + (0.0 / float64(3.14159)))
            _ = fi37
            _ = v50
        }
    }
    v9 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
