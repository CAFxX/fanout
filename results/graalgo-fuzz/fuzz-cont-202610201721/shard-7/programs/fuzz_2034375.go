package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint16 }

func main() {
    v1 := []string{}
    var v2 uint = 503
    var v3 uint32 = 42
    if true {
        v4 := GenS1{F0: -1}
        if (4294967295 <= 559) {
            for fi5 := 0; fi5 < 13; fi5++ {
                v2 += v2
                fmt.Println(fi5, v2, v2)
                var v6 int = (fi5 + fi5)
                _ = fi5
                _ = v6
            }
        }
        _ = v4
    } else {
        if (v3 > (v3 + (v3 + 4294967295))) {
            v7 := GenS2{F0: uint8(v3), F1: 115, F2: 127}
            if false {
                if ((!(1e10 != 1e10)) || (!(-611 > 1))) {
                    fmt.Println(v3)
                } else {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v2 = (v2 | v2)
                        fmt.Println(fi8, v2)
                        _ = fi8
                    }
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                }
            }
            v9 := GenS1{F0: 0}
            _ = v7
            _ = v9
        }
        var v10 bool = true
        _ = v10
    }
    switch 714 {
    case 42:
        var v11 bool = (1 <= 100)
        _ = v11
    case 128:
        v2 += (^(v2))
    default:
        v3 = 1
    }
    _ = len(v1)
    switch 127 {
    case 128:
        if false {
            switch "" {
            case "":
                v1 = append(v1, "foo bar")
            case "a":
                var v12 float32 = 1.5
                _ = v12
            case "foo bar":
                switch 78 {
                case 127:
                    if true {
                        if (42 != 940) {
                            var v13 uint64 = 771
                            if (!(0.1 > 0.1)) {
                                v3 -= v3
                                fmt.Println(v2)
                                v1 = append(v1, "\n\t")
                            }
                            var v14 float32 = 0.1
                            _ = v13
                            _ = v14
                        }
                        var v15 float64 = ((0.1 * 3.14159) * 1.5)
                        for fi16 := 0; fi16 < 10; fi16++ {
                            v2 += (v2 >> ((1) & 63))
                            v1 = append(v1, ("foo bar" + "hello"))
                            var v17 uint32 = (^(v3))
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                default:
                    v2 -= 1
                }
            }
        }
    case 0:
        var v18 int64 = 437
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 string = "a"
            v3 = 2147483648
            v20 = (v20 + v20)
            _ = fi19
            _ = v20
        }
        var v21 uint32 = (v3 ^ v3)
        _ = v21
    }
    if (0.0 != (0.1 - 1.5)) {
        v2 = uint(v3)
        for fi22 := 0; fi22 < 13; fi22++ {
            v3 += (v3 - (^(v3)))
            var v23 int = (fi22 + -9223372036854775808)
            var v24 int64 = 1
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 int8 = -1
            _ = fi25
            _ = v26
        }
    } else {
        v2 -= v2
        for fi27 := 0; fi27 < 15; fi27++ {
            switch fi27 {
            case 421:
                fmt.Println(v2, v3, v3)
            case 0:
                for fi28 := 0; fi28 < 14; fi28++ {
                    v2 = 1
                    switch "a" {
                    case "a":
                        if (("a" + "hello") >= "a") {
                            var v29 int32 = 2147483647
                            v30 := GenS3{F0: 42, F1: true}
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = fi28
                }
            case 32767:
                if (true && (243 < 18446744073709551615)) {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        v32 := GenS1{F0: -2147483648}
                        v33 := GenS4{F0: -176}
                        v1 = append(v1, ((("a" + "hello") + "hello") + ("" + ("" + ("hello" + ((("" + "foo bar") + "foo bar") + ((((("a" + (("\n\t" + "\n\t") + "hello")) + ("a" + "foo bar")) + "foo bar") + "a") + (("foo bar" + ("hello" + "hello")) + ("" + "hello")))))))))
                        var v34 uint8 = 255
                        var v35 uint64 = 0
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v3 = (v3 - 383)
                    v2 -= v2
                }
            }
            fmt.Println(v3, v3)
            _ = fi27
        }
    }
    v36 := GenS5{F0: -128, F1: 1}
    v37 := GenS3{F0: 1000, F1: (4294967295 == 2)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
