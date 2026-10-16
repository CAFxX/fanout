package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = ("foo bar" == "\n\t")
    var v3 float32 = 0.0
    var v4 bool = v2
    var v5 float32 = v3
    var v6 int = -730
    var v7 int8 = 57
    v8 := GenS1{F0: (765 <= 1)}
    switch ("a" + "hello") {
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            v7 = v7
            _ = len(v1)
            v6 = fi9
            _ = fi9
        }
    case "a":
        if v4 {
            fmt.Println(v5, v7)
            if false {
                var v10 string = "foo bar"
                for fi11 := 0; fi11 < 5; fi11++ {
                    v3 = v5
                    var v12 uint32 = 394
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            v4 = ((false && v2) || (true || v2))
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                switch ("\n\t" + "hello") {
                case "foo bar":
                    var v14 uint8 = 127
                    _ = v14
                case "":
                    fmt.Println(v6, v2)
                case "hello":
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v6 = 238
                        switch (v6 * fi13) {
                        case 127:
                            v6 -= fi15
                            v5 = v3
                        case 0:
                            v16 := v1[v6]
                            _ = v16
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v2)
                            }
                        }
                        _ = fi15
                    }
                default:
                    fmt.Println(v3, v4, v5)
                }
                switch fi13 {
                case 9223372036854775807:
                    v17 := GenS2{F0: 2, F1: 22, F2: 0}
                    _ = v17
                case 158:
                    switch "\n\t" {
                    case "":
                        v18 := GenS2{F0: 128, F1: 9223372036854775807, F2: 128}
                        _ = v18
                    default:
                        v4 = ("\n\t" == "hello")
                    }
                default:
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v4 = (v4 || (!false))
                        _ = fi19
                    }
                }
                switch v6 {
                case 0:
                    switch (fi13 - 32767) {
                    case 4294967295:
                        v20 := GenS2{F0: 42, F1: 9223372036854775807, F2: 128}
                        _ = v20
                    case -1:
                        var v21 int = (-(fi13))
                        _ = v21
                    default:
                        v22 := []int{(fi13 & (v6 | (fi13 ^ fi13))), v6, 243}
                        _ = v22
                    }
                }
                _ = fi13
            }
            v23 := GenS3{F0: 32767, F1: 1, F2: 0}
            var v24 int32 = 2147483647
            _ = v23
            _ = v24
        }
    case "\n\t":
        v25 := GenS2{F0: 0, F1: -384, F2: 128}
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v6, v3, v7)
        _ = fi26
    }
    v1[(v6 + (v6 % ((v6) | 1)))] = (v6 & (v6 ^ v6))
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 9; fi28++ {
            v2 = v4
            var v29 uint64 = 4294967295
            if (!v2) {
                fmt.Println(fi28, v4)
                var v30 float64 = -2.25
                v30 = (float64(v5) + (3.14159 / 1.5))
                _ = len(v1)
                _ = v30
            }
            _ = fi28
            _ = v29
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v4 = ("hello" == (((("a" + (("a" + ("hello" + ("hello" + "\n\t"))) + ("" + ""))) + "") + ("foo bar" + "\n\t")) + "foo bar"))
            _ = fi31
        }
        v2 = (42 != 725)
        _ = fi27
    }
    if (4294967295 >= 2147483647) {
        if v4 {
            var v32 float64 = (float64(v3) - 1.5)
            v33 := GenS2{F0: 42, F1: -9223372036854775808, F2: 0}
            _ = v32
            _ = v33
        } else {
            v34 := []int{v6}
            var v35 int32 = 192
            var v36 float64 = (-2.25 / (-2.25 + (-2.25 * ((1.5 * 1e10) - float64((v3 + 1e10))))))
            _ = v34
            _ = v35
            _ = v36
        }
    }
    var v37 uint = 239
    v38 := v1[v6]
    switch "" {
    case "foo bar":
        for fi39 := 0; fi39 < 15; fi39++ {
            v2 = (2147483647 == 2)
            switch "a" {
            case "a":
                v38 -= 32768
            case "hello":
                v38 -= v6
            default:
                if (v2 && v4) {
                    v6 += fi39
                    switch "" {
                    case "foo bar":
                        v40 := v1[(v38 + 0)]
                        _ = v40
                    case "\n\t":
                        if v2 {
                            v5 -= 3.14159
                        }
                    case "a":
                        v41 := GenS2{F0: 74, F1: -1, F2: 76}
                        _ = v41
                    }
                } else {
                    v42 := v1[128]
                    if ((v2 || (-1 > -9223372036854775808)) && ("foo bar" <= ((((("\n\t" + "") + "") + "") + "") + "a"))) {
                        v6 = v6
                    }
                    if ((float64(v3) * (1.5 - float64((0.0 / v3)))) < 0.1) {
                        for fi43 := 0; fi43 < 20; fi43++ {
                            v42 = v38
                            v44 := GenS1{F0: ((0 <= 607) || (883 <= 0))}
                            _ = fi43
                            _ = v44
                        }
                        fmt.Println(v6, v4, v6)
                    }
                    _ = v42
                }
            }
            _ = fi39
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
