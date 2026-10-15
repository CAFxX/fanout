package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    v1 := []uint32{932, 2147483648}
    var v2 string = "foo bar"
    v3 := []int{}
    _ = len(v1)
    v2 = ("" + v2)
    v4 := GenS1{F0: 2}
    v2 = ("" + ("" + v2))
    if (1 != -1) {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint8 = 157
            _ = fi5
            _ = v6
        }
        v7 := GenS2{F0: -9223372036854775808}
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS3{F0: 127, F1: fi8, F2: 0}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                v2 = "\n\t"
                var v11 int = (fi8 + -493)
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v2 = "a"
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (1 <= 0) {
        if ((0.0 - ((float32(1.5) / 1.5) - 1.5)) > float32((1e10 - (((3.14159 - float64((1e10 / 3.14159))) + float64(1.5)) / 3.14159)))) {
            if true {
                if (0 <= 819) {
                    if false {
                        var v12 bool = (true || (!false))
                        v1 = append(v1, 634)
                        _ = v12
                    } else {
                        fmt.Println(v2)
                    }
                } else {
                    v2 = v2
                    var v13 uint32 = 1
                    _ = v13
                }
                switch ("" + (((("" + v2) + v2) + v2) + v2)) {
                case "foo bar":
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        var v14 float32 = (((0.0 - (1e10 + 0.1)) / float32(3.14159)) * (-2.25 * 0.1))
                        v2 = "\n\t"
                        _ = v14
                    } else {
                        v15 := GenS2{F0: -937}
                        v16 := GenS2{F0: 1}
                        switch 0 {
                        case 32767:
                            var v17 bool = (42 != 100)
                            _ = v17
                        case 128:
                            for fi18 := 0; fi18 < 19; fi18++ {
                                var v19 float64 = 3.14159
                                _ = fi18
                                _ = v19
                            }
                        }
                        _ = v15
                        _ = v16
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = 69
                    }
                }
                v2 = ("foo bar" + "foo bar")
            }
            v2 = v2
        } else {
            switch "\n\t" {
            case "\n\t":
                switch 32768 {
                case -866:
                    _ = len(v1)
                }
            case "hello":
                v20 := GenS1{F0: 952}
                _ = v20
            }
        }
    } else {
        if ((276 != 1) && (1 < 2)) {
            for fi21 := 0; fi21 < 14; fi21++ {
                v2 = ("\n\t" + v2)
                v1 = append(v1, 4294967295)
                _ = fi21
            }
            switch "hello" {
            case "\n\t":
                if true {
                    if ((!(true && (!(false && (108 >= 135))))) || false) {
                        v2 = ("\n\t" + (v2 + v2))
                        fmt.Println(v2)
                        if false {
                            var v22 int32 = -1
                            var v23 bool = true
                            _ = v22
                            _ = v23
                        }
                    }
                    var v24 uint8 = 0
                    _ = v24
                } else {
                    _ = len(v3)
                }
            default:
                var v25 uint32 = 65535
                _ = v25
            }
        } else {
            fmt.Println(v2)
            v26 := GenS1{F0: 461}
            var v27 uint64 = 18446744073709551615
            _ = v26
            _ = v27
        }
        v1 = append(v1, 994)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v28 uint32 = 2147483648
    _ = len(v1)
    for fi29 := 0; fi29 < 16; fi29++ {
        if true {
            v1 = append(v1, v28)
            v1 = append(v1, v28)
        }
        _ = fi29
    }
    v2 = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
