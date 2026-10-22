package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 string }

func main() {
    var v1 uint16 = 0
    v2 := []int{244, 9223372036854775807, 128}
    var v3 int16 = 0
    var v4 uint8 = 122
    v5 := GenS1{F0: v3, F1: (false || (!false)), F2: 42}
    if (v3 < (v3 << ((821) & 63))) {
        v1 += v1
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS2{F0: fi6, F1: 1, F2: "foo bar"}
            _ = fi6
            _ = v7
        }
    }
    if false {
        v1 -= (v1 + (^(v1)))
        if ((1 != 253) || false) {
            fmt.Println(v4, v3, v1)
            if ((1 == v1) || false) {
                var v8 string = ""
                _ = v8
            }
            v9 := GenS3{F0: "a"}
            _ = v9
        }
    }
    v1 = (v1 + 1)
    var v10 bool = ((!true) && (-1 > -128))
    v11 := GenS1{F0: v3, F1: (351 == 0), F2: -1}
    v10 = ((float64(0.1) / 1.0) != ((1.5 + 1.5) * 0.0))
    v10 = ((!((v4 & uint8(v1)) < (v4 * 1))) || v10)
    if true {
        for fi12 := 0; fi12 < 9; fi12++ {
            v1 = v1
            _ = fi12
        }
    } else {
        if ((^(v1)) > v1) {
            v1 += (v1 + (v1 + 42))
        }
        var v13 string = (("foo bar" + "") + "foo bar")
        v1 = (v1 >> ((0) & 63))
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        switch "hello" {
        case "hello":
            switch (("hello" + "a") + "hello") {
            case "hello":
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS2{F0: fi14, F1: -2147483648, F2: "foo bar"}
                    v3 += v3
                    v10 = (1.5 >= 1e10)
                    fmt.Println(fi14)
                    var v17 int32 = 815
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case "foo bar":
                switch 0 {
                case -242:
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                case 42:
                    fmt.Println(v4, v4, v10)
                default:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if len(v2) > 0 {
                            v2[0] = -260
                        }
                        _ = fi18
                    }
                }
            default:
                if true {
                    if len(v2) > 0 {
                        v2[0] = (fi14 | fi14)
                    }
                    switch "\n\t" {
                    case "":
                        if (-32768 != (v3 % ((v3) | 1))) {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                _ = len(v2)
                                _ = fi19
                            }
                            v10 = (true || (-32768 > v3))
                            fmt.Println(fi14, fi14)
                        }
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = fi14
                        }
                    case "foo bar":
                        fmt.Println(v4, v3, v1)
                    }
                } else {
                    var v20 uint16 = v1
                    _ = v20
                }
            }
        case "foo bar":
            var v21 uint8 = (v4 % (((v4 | v4)) | 1))
            _ = v21
        }
        var v22 string = ""
        _ = fi14
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v1 = v1
        _ = fi23
    }
    var v24 uint8 = uint8(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
