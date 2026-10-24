package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    var v4 float64 = 1e10
    var v5 string = "a"
    if false {
        v2 = 1
        if len(v1) > 0 {
            v1[0] = 0
        }
        var v6 uint32 = 2147483648
        _ = v6
    } else {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: (!((1 <= -9223372036854775808) && true)), F1: 127}
            delete(v3, (((((((v5 + v5) + v5) + v5) + v5) + (((("foo bar" + "\n\t") + ("" + v5)) + v5) + v5)) + v5) + (v5 + "")))
            _ = fi7
            _ = v8
        }
        v2 = v2
        v2 = (v2 % ((-129) | 1))
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println(v2, v2)
    v2 = (^(v2))
    v9 := GenS1{F0: (((!((-8 < 127) && (42 == 255))) || false) && (!(true && ("" != "hello")))), F1: 128}
    fmt.Println(v2, v5)
    switch 9223372036854775807 {
    case 42:
        for fi10 := 0; fi10 < 13; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                v2 += v2
                v11 := GenS1{F0: (0 < 906), F1: 207}
                switch (-(fi10)) {
                case 128:
                    v4 = (-2.25 - 0.1)
                case 237:
                    switch fi10 {
                    case 32767:
                        for fi12 := 0; fi12 < 10; fi12++ {
                            v4 -= 0.0
                            _ = fi12
                        }
                    case -9223372036854775808:
                        switch fi10 {
                        case -462:
                            switch (fi10 + (^(fi10))) {
                            case 1:
                                delete(v3, "foo bar")
                            default:
                                delete(v3, v5)
                            }
                        case 2147483647:
                            v4 = 3.14159
                        case 332:
                            v13 := GenS2{F0: 42, F1: 988, F2: -1}
                            _ = v13
                        }
                    default:
                        if (0 >= 0) {
                            var v14 int32 = 2147483647
                            delete(v3, v5)
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 14; fi15++ {
                                switch (fi10 >> ((0) & 63)) {
                                case 128:
                                    v5 = v5
                                    if true {
                                        v5 = ((v5 + v5) + (v5 + ("\n\t" + v5)))
                                    }
                                    fmt.Println(v2, v5, v4)
                                case 2147483647:
                                    delete(v3, ("" + ((v5 + ("\n\t" + "")) + "")))
                                case 58:
                                    var v16 string = "foo bar"
                                    _ = v16
                                }
                                var v17 float32 = 1e10
                                _ = fi15
                                _ = v17
                            }
                        }
                    }
                }
                _ = v11
            } else {
                fmt.Println(v2, v2)
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch (fi18 << ((2) & 63)) {
                    case 0:
                        v4 -= -2.25
                        switch "foo bar" {
                        case "":
                            var v19 float64 = -2.25
                            _ = v19
                        }
                        v2 = v2
                    case 9223372036854775807:
                        var v20 uint = 255
                        _ = v20
                    case 128:
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    default:
                        if true {
                            v21 := GenS3{F0: fi18, F1: 65535, F2: 42}
                            var v22 string = (v5 + v5)
                            fmt.Println(fi10)
                            _ = v21
                            _ = v22
                        }
                    }
                    fmt.Println(v4, v2, fi18)
                    _ = fi18
                }
                var v23 float64 = 0.1
                _ = v23
            }
            _ = fi10
        }
    default:
        v3["\n\t"] = 2147483647
    }
    fmt.Println(v5)
    v24 := GenS1{F0: true, F1: 1}
    v5 = (("a" + v5) + v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
