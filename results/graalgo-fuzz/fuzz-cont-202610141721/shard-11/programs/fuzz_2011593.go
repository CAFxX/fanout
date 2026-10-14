package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 float32 = ((0.1 / (3.14159 + 0.0)) * (0.1 + -2.25))
    var v2 int64 = 0
    v3 := []uint32{65535}
    v4 := []int{-1}
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 1}
    if false {
        switch 86 {
        case 128:
            for fi7 := 0; fi7 < 11; fi7++ {
                if false {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        v2 -= 42
                        _ = fi8
                    }
                }
                _ = fi7
            }
        default:
            switch -9223372036854775808 {
            case 0:
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 uint = 0
                    var v11 string = "\n\t"
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                var v12 uint = 65535
                _ = v12
            }
        }
        var v13 uint64 = 4294967295
        var v14 string = "foo bar"
        _ = v13
        _ = v14
    }
    v3 = append(v3, 2)
    switch 128 {
    case 548:
        v2 = v2
    default:
        v1 = v1
    }
    v1 += ((1e10 - ((float32(((0.0 + float64(v1)) * -2.25)) * float32((3.14159 - 0.0))) + ((1.5 / 1.0) / 3.14159))) * -2.25)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 int64 = v2
        fmt.Println(v16)
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    var v17 int16 = -622
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(fi18, v1)
        switch ("a" + (("hello" + "a") + ((((("" + (("" + ("a" + "foo bar")) + ("foo bar" + "hello"))) + (("a" + ("a" + "\n\t")) + "hello")) + ("\n\t" + (("" + "foo bar") + (((("a" + "a") + "hello") + "hello") + ("" + "\n\t"))))) + (("a" + "hello") + ("hello" + "\n\t"))) + ""))) {
        case "foo bar":
            if (v2 == (v2 + 240)) {
                fmt.Println(fi18, fi18, fi18)
                var v19 uint64 = 616
                _ = v19
            }
        case "hello":
            v1 = 1e10
        case "a":
            for fi20 := 0; fi20 < 3; fi20++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v21 int16 = (v17 - (v17 ^ (v17 / (((v17 - v17)) | 1))))
                var v22 uint = 946
                if len(v4) > 0 {
                    v4[0] = (fi20 ^ (fi20 % ((fi18) | 1)))
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi18
    }
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            v2 -= int64(v17)
            if false {
                v2 = 852
                v24 := GenS1{F0: v2}
                _ = v24
            }
            if false {
                var v25 uint16 = 42
                _ = v25
            }
            _ = fi23
        }
        var v26 int8 = -122
        _ = v26
    } else {
        var v27 int = 42
        _ = v27
    }
    v2 = (v2 * (v2 + (v2 - v2)))
    v1 += (v1 * (float32(-2.25) / -2.25))
    if (!(!(!(true || true)))) {
        v17 = (v17 * (v17 + v17))
        fmt.Println(v1, v2, v17)
    } else {
        v1 = v1
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := GenS2{F0: 50, F1: (^(fi28))}
                if true {
                    v31 := GenS2{F0: 1, F1: fi29}
                    switch "hello" {
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                        v2 -= -351
                        var v32 uint32 = 0
                        fmt.Println(v17, v1)
                        switch (-(fi29)) {
                        case 0:
                            if false {
                                v33 := v5[929]
                                var v34 int64 = v2
                                _ = v33
                                _ = v34
                            } else {
                                v1 = 0.0
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    fmt.Println(fi28, v17, v2)
                                case "hello":
                                    var v35 bool = (false && true)
                                    _ = v35
                                }
                                v36 := GenS3{F0: true}
                                _ = v36
                            }
                        case 32768:
                            var v37 uint64 = 18446744073709551615
                            _ = v37
                        }
                        _ = v32
                    case "\n\t":
                        var v38 uint32 = 1
                        _ = v38
                    }
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
        v39 := v5[2147483648]
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
