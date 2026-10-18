package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint = 42
    var v2 float64 = 1e10
    var v3 uint32 = 2
    v4 := []string{"", "a", ("hello" + (("hello" + ("hello" + "\n\t")) + "\n\t")), "\n\t"}
    var v5 uint16 = 42
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    fmt.Println(v1)
    var v9 int = 128
    fmt.Println(v2)
    switch v9 {
    case 9223372036854775807:
        _ = len(v6)
    default:
        switch "a" {
        case "hello":
            var v10 uint16 = v5
            _ = v10
        case "\n\t":
            var v11 int16 = int16(v1)
            _ = v11
        case "a":
            if (false || false) {
                var v12 int32 = -2147483648
                _ = v12
            } else {
                switch 32768 {
                case 32768:
                    if false {
                        var v13 int64 = -718
                        v9 += v9
                        _ = v13
                    } else {
                        var v14 int8 = int8(v3)
                        if (-2.25 == 1.5) {
                            fmt.Println(v1, v3)
                            fmt.Println(v14, v1, v2)
                        }
                        _ = v14
                    }
                default:
                    if true {
                        v15 := GenS1{F0: 42, F1: 1}
                        switch v9 {
                        case -1:
                            if true {
                                var v16 int32 = -1
                                _ = v16
                            } else {
                                var v17 bool = (v5 <= v5)
                                _ = v17
                            }
                        case 4294967295:
                            switch v9 {
                            case 32768:
                                var v18 int16 = int16(v9)
                                _ = v18
                            case 135:
                                v5 += 42
                            case 0:
                                switch (("a" + ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "\n\t"))) + ("hello" + ("hello" + ("foo bar" + "a")))) {
                                case "\n\t":
                                    v19 := GenS2{F0: "foo bar"}
                                    _ = v19
                                default:
                                    if (-2147483648 != 1) {
                                        var v20 uint = 255
                                        switch (v9 & (v9 >> ((911) & 63))) {
                                        case 128:
                                            v9 = (v9 ^ (-(v9)))
                                        case -9223372036854775808:
                                            _ = len(v4)
                                        }
                                        if true {
                                            fmt.Println(v3)
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                if true {
                                                    v6[(v9 << (((v1 << (((v1 << ((65535) & 63))) & 63))) & 63))] = (v9 << ((v1) & 63))
                                                    var v21 uint16 = 42
                                                    _ = v21
                                                }
                                            case "":
                                                switch (v9 - v9) {
                                                case 782:
                                                    v9 -= 2147483647
                                                case 0:
                                                    var v22 int32 = 1
                                                    _ = v22
                                                case 319:
                                                    if (!((65535 != (v1 - uint(v5))) && true)) {
                                                        v23 := GenS3{F0: (v3 & v3)}
                                                        v24 := GenS4{F0: (^(v9))}
                                                        var v25 float64 = 0.1
                                                        _ = v23
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                }
                                            case "hello":
                                                switch (((("hello" + "\n\t") + ("hello" + "foo bar")) + "foo bar") + ("a" + "foo bar")) {
                                                case "":
                                                    var v26 string = "\n\t"
                                                    _ = v26
                                                case "a":
                                                    switch (v9 ^ 2147483647) {
                                                    case -241:
                                                        for fi27 := 0; fi27 < 16; fi27++ {
                                                            if true {
                                                                fmt.Println(v20, v3, v3)
                                                                v2 -= v2
                                                                switch (((("" + "") + "a") + "a") + "a") {
                                                                case "foo bar":
                                                                    var v28 bool = (42 <= 161)
                                                                    _ = v28
                                                                }
                                                            }
                                                            var v29 float64 = v2
                                                            _ = fi27
                                                            _ = v29
                                                        }
                                                    case 967:
                                                        v4 = append(v4, ("foo bar" + ("hello" + ("foo bar" + ("hello" + ("foo bar" + "hello"))))))
                                                    default:
                                                        v30 := v6[(v9 + (v9 % (((v9 * (v9 + (v9 ^ (v9 | (-(v9))))))) | 1)))]
                                                        _ = v30
                                                    }
                                                case "hello":
                                                    for fi31 := 0; fi31 < 19; fi31++ {
                                                        fmt.Println(fi31, v1)
                                                        if false {
                                                            v9 -= v9
                                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                                v3 += v3
                                                                _ = fi32
                                                            }
                                                            v5 = uint16(v20)
                                                        }
                                                        fmt.Println(v9, v1)
                                                        _ = fi31
                                                    }
                                                }
                                            }
                                            _ = len(v4)
                                            v5 += v5
                                        }
                                        _ = v20
                                    } else {
                                        var v33 uint64 = 781
                                        delete(v6, -527)
                                        _ = v33
                                    }
                                }
                            default:
                                var v34 int64 = 9223372036854775807
                                _ = v34
                            }
                        case 32768:
                            fmt.Println(v3, v2)
                        default:
                            var v35 uint32 = v3
                            _ = v35
                        }
                        _ = v15
                    }
                }
            }
        }
    }
    if (((v3 & v3) <= 44) && ((^(v9)) > v9)) {
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v2, v1)
        case "foo bar":
            if false {
                fmt.Println(v2, v9)
            }
        }
        for fi36 := 0; fi36 < 18; fi36++ {
            fmt.Println(v9, v1, v1)
            _ = fi36
        }
    } else {
        v9 = (-(v9))
        v37 := GenS3{F0: v3}
        var v38 uint = v1
        _ = v37
        _ = v38
    }
    v5 += (v5 - (^(v5)))
    var v39 bool = (!(true && (0 == -2147483648)))
    switch "hello" {
    case "foo bar":
        var v40 bool = v39
        _ = v40
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
