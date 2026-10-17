package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 uint32, p1 bool) int8 {
    p1 = false
    if (!(65535 >= 32767)) {
        if (p1 && false) {
            for fi4 := 0; fi4 < 7; fi4++ {
                for fi5 := 0; fi5 < 5; fi5++ {
                    v6 := GenS1{F0: 18446744073709551615, F1: (!((((((float64((float32(((1e10 - -2.25) - (-2.25 / ((((1e10 - -2.25) * 1.5) - -2.25) + (-2.25 - 1.5))))) / 1.0)) * (float64((1.5 - 3.14159)) / 1.5)) * 1.5) - -2.25) + ((float64((((0.1 * 1.5) / (-2.25 / 1.5)) + float32(3.14159))) * 1.5) / 3.14159)) * 1.5) == float64(1e10))), F2: 353}
                    _ = fi5
                    _ = v6
                }
                for fi7 := 0; fi7 < 8; fi7++ {
                    v8 := GenS1{F0: 1, F1: (!p1), F2: 0}
                    fmt.Println(fi4, p1, fi4)
                    _ = fi7
                    _ = v8
                }
                _ = fi4
            }
        }
        fmt.Println(p0)
        switch 32767 {
        case 4294967295:
            v9 := GenS1{F0: 2, F1: p1, F2: 882}
            _ = v9
        case -1:
            for fi10 := 0; fi10 < 7; fi10++ {
                p1 = true
                _ = fi10
            }
        }
    }
    switch "hello" {
    case "foo bar":
        var v11 uint = 42
        _ = v11
    }
    _ = p0
    _ = p1
    return 100
}

func main() {
    var v1 int = 1
    v2 := []string{"foo bar", "a"}
    var v3 uint = 511
    switch 127 {
    case 127:
        var v12 uint64 = 42
        _ = v12
    case -68:
        if false {
            v1 = v1
            v1 += v1
            if true {
                if (!((1.5 + 3.14159) < float64(0.0))) {
                    v1 = v1
                }
                var v13 uint64 = 2
                var v14 uint64 = v13
                _ = v13
                _ = v14
            }
        }
    case 128:
        v1 += v1
    }
    _ = len(v2)
    var v15 string = "foo bar"
    if true {
        for fi16 := 0; fi16 < 6; fi16++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int16 = int16(v3)
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 float32 = 3.14159
        v20 = v20
        fmt.Println(v3, v15, v15)
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v1, v1)
    v3 -= v3
    switch "hello" {
    case "a":
        var v21 bool = (!true)
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint = (v3 - (v3 - v3))
        _ = len(v2)
        v24 := GenS1{F0: 2, F1: true, F2: 336}
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        var v25 int16 = 42
        _ = v25
    }
    v15 = v15
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v1)
        if (!(42 != 42)) {
            switch "\n\t" {
            case "a":
                v27 := GenS1{F0: 26, F1: (-2.25 >= (0.1 + 1e10)), F2: 0}
                _ = v27
            default:
                if (0 <= 304) {
                    fmt.Println(v1, fi26)
                    v28 := GenS1{F0: 18446744073709551615, F1: (1.5 > 0.0), F2: 255}
                    var v29 uint16 = 2
                    _ = v28
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 14; fi30++ {
                        fmt.Println(v1, v15, fi26)
                        if true {
                            switch (fi30 % ((820) | 1)) {
                            case 525:
                                v31 := GenS1{F0: 2, F1: (true || (false || (!true))), F2: 65535}
                                _ = v31
                            case 1:
                                var v32 string = (v15 + (v15 + ("a" + "foo bar")))
                                _ = v32
                            default:
                                v1 += (fi26 + fi26)
                            }
                            var v33 float32 = 1e10
                            fmt.Println(v3, v33, fi30)
                            _ = v33
                        } else {
                            var v34 int8 = genFn0(1, (-9223372036854775808 == 0))
                            if true {
                                switch fi30 {
                                case -529:
                                    var v35 uint64 = 1
                                    _ = v35
                                case 42:
                                    v1 = int(v3)
                                default:
                                    if (18446744073709551615 != 231) {
                                        switch "hello" {
                                        case "\n\t":
                                            v1 -= fi30
                                        default:
                                            v15 = v15
                                        }
                                        var v36 int8 = (-(v34))
                                        _ = v36
                                    }
                                }
                            }
                            _ = v34
                        }
                        v37 := GenS1{F0: 42, F1: (!true), F2: 2}
                        _ = fi30
                        _ = v37
                    }
                }
            }
            v1 = fi26
            v3 = v3
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
