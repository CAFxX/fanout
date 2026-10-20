package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int; F2 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{}
    var v3 string = ("foo bar" + "")
    var v4 bool = true
    v4 = v4
    v1["a"] = -81
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(v3)
        switch (fi5 & fi5) {
        case -1:
            if v4 {
                switch fi5 {
                case 32768:
                    for fi6 := 0; fi6 < 18; fi6++ {
                        if (-32768 >= -878) {
                            v3 = ""
                            if false {
                                var v7 int8 = 1
                                _ = v7
                            }
                            v3 = "foo bar"
                            if true {
                                if v4 {
                                    var v8 uint8 = 1
                                    _ = v8
                                }
                            }
                        }
                        v9 := GenS1{F0: 4, F1: 87, F2: 127}
                        _ = fi6
                        _ = v9
                    }
                case 2147483647:
                    v3 = "hello"
                case -1:
                    if false {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            delete(v1, v3)
                            v11 := GenS1{F0: 0, F1: 968, F2: 128}
                            fmt.Println(v4, v3)
                            _ = fi10
                            _ = v11
                        }
                        var v12 int = (fi5 ^ fi5)
                        _ = v12
                    }
                }
                var v13 int = fi5
                delete(v1, "\n\t")
                _ = v13
            } else {
                var v14 float64 = 0.1
                v4 = v4
                _ = v14
            }
        case 1:
            if (false || ((false && true) && false)) {
                var v15 string = (v3 + v3)
                var v16 int64 = -817
                var v17 float64 = 3.14159
                _ = v15
                _ = v16
                _ = v17
            }
        case 2147483647:
            if v4 {
                fmt.Println(v3)
                v18 := GenS1{F0: 1, F1: 1, F2: 88}
                _ = v18
            } else {
                switch fi5 {
                case 616:
                    switch (("\n\t" + v3) + v3) {
                    case "hello":
                        if v4 {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((fi5 >> ((414) & 63)) < fi5) {
                                var v19 uint = 65535
                                _ = v19
                            }
                            switch 678 {
                            case -283:
                                for fi20 := 0; fi20 < 3; fi20++ {
                                    var v21 int = (fi20 | fi5)
                                    if (701 < 352) {
                                        if v4 {
                                            v4 = (2147483647 > 2147483647)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = v3
                                        }
                                    } else {
                                        v4 = v4
                                        v4 = true
                                        fmt.Println(v4, v3, v4)
                                    }
                                    var v22 string = "a"
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                    }
                }
                v4 = v4
            }
        }
        fmt.Println(v3, fi5)
        _ = fi5
    }
    fmt.Println(v3)
    for fi23 := 0; fi23 < 15; fi23++ {
        switch v3 {
        case "hello":
            for fi24 := 0; fi24 < 19; fi24++ {
                switch ("hello" + (("hello" + v3) + v3)) {
                case "\n\t":
                    var v25 uint8 = 54
                    _ = v25
                case "":
                    if (!true) {
                        v4 = v4
                        switch v3 {
                        case "foo bar":
                            var v26 bool = (v4 && v4)
                            _ = v26
                        }
                        fmt.Println(fi23)
                    }
                }
                v27 := GenS1{F0: 42, F1: 42, F2: 42}
                switch "" {
                case "a":
                    switch 32768 {
                    case 4294967295:
                        v3 = v3
                    case -442:
                        switch fi23 {
                        case 32767:
                            v28 := GenS2{F0: 42, F1: true, F2: 32767}
                            _ = v28
                        case 0:
                            var v29 string = "\n\t"
                            _ = v29
                        case -1:
                            v3 = ""
                        }
                    case 0:
                        if (1 != -61) {
                            var v30 int = -849
                            v31 := GenS3{F0: (fi24 != (fi23 & (-(v30)))), F1: (v30 / (((fi23 % ((fi24) | 1))) | 1)), F2: 127}
                            var v32 int = 1
                            v2 = append(v2, (v3 + v3))
                            v32 += v30
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    default:
                        var v33 uint = 0
                        _ = v33
                    }
                }
                _ = fi24
                _ = v27
            }
        case "\n\t":
            v34 := GenS3{F0: ((1000 <= 32767) && (4294967295 <= 209)), F1: fi23, F2: 100}
            _ = v34
        default:
            var v35 string = (((((v3 + v3) + (((v3 + (v3 + v3)) + ("\n\t" + "foo bar")) + v3)) + "foo bar") + v3) + v3)
            _ = v35
        }
        for fi36 := 0; fi36 < 18; fi36++ {
            delete(v1, ("foo bar" + ("a" + v3)))
            _ = fi36
        }
        _ = fi23
    }
    if (!false) {
        var v37 int64 = -1
        _ = v37
    }
    var v38 int32 = -1
    var v39 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
