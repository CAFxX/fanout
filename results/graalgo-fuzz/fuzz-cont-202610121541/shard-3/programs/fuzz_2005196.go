package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint64, p2 bool) string {
    fmt.Println(p0, p2, p0)
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(p2, fi4)
        var v5 int8 = -128
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 104
    if p2 {
        if ((^(v6)) < v6) {
            switch "" {
            case "\n\t":
                for fi7 := 0; fi7 < 16; fi7++ {
                    v6 = uint8(p1)
                    _ = fi7
                }
            case "a":
                v8 := GenS1{F0: v6}
                _ = v8
            case "":
                p2 = (0.1 == 1.5)
            default:
                v9 := GenS1{F0: (v6 - (^(v6)))}
                _ = v9
            }
            v6 = v6
            var v10 string = "\n\t"
            _ = v10
        }
        var v11 int = -9223372036854775808
        fmt.Println(p1, p0, p0)
        _ = v11
    } else {
        p0 = p0
        if false {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint = 2
                _ = fi12
                _ = v13
            }
            var v14 int = 127
            _ = v14
        } else {
            if p2 {
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 float64 = 0.0
                    p2 = ((v6 ^ v6) < (v6 >> ((65535) & 63)))
                    _ = fi15
                    _ = v16
                }
            }
        }
        v6 += 155
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    switch "foo bar" {
    case "hello":
        v2 -= -112
    case "foo bar":
        switch "hello" {
        case "foo bar":
            switch 259 {
            case 32768:
                v2 = v2
            default:
                if (117 <= 127) {
                    var v17 float64 = ((3.14159 + 1e10) + (1.5 * (3.14159 - 3.14159)))
                    var v18 bool = false
                    _ = v17
                    _ = v18
                }
            }
        case "hello":
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(fi19, v2)
                if (false && ((v2 - v2) == v2)) {
                    if true {
                        v1[""] = "hello"
                        fmt.Println(fi19, v2)
                    }
                    if (((3.14159 + -2.25) * ((0.1 / 1.0) / ((1e10 * 3.14159) - float32(0.0)))) != (1.5 - 0.0)) {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            var v21 uint = 1
                            _ = fi20
                            _ = v21
                        }
                        if ((0.0 <= 1e10) || ("hello" < "foo bar")) {
                            switch ("foo bar" + ("foo bar" + "foo bar")) {
                            case "\n\t":
                                v2 = (v2 % ((9223372036854775807) | 1))
                            case "foo bar":
                                delete(v1, "\n\t")
                            }
                            if false {
                                v22 := GenS1{F0: 0}
                                _ = v22
                            } else {
                                if (!true) {
                                    switch "hello" {
                                    case "\n\t":
                                        v2 = 9223372036854775807
                                    default:
                                        v23 := GenS1{F0: 88}
                                        _ = v23
                                    }
                                }
                                var v24 int32 = -250
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    v2 = (v2 | (v2 * v2))
                                    delete(v1, (("foo bar" + "\n\t") + "hello"))
                                    _ = fi25
                                }
                                _ = v24
                            }
                            for fi26 := 0; fi26 < 1; fi26++ {
                                var v27 int = fi26
                                _ = fi26
                                _ = v27
                            }
                        } else {
                            var v28 bool = false
                            _ = v28
                        }
                    }
                }
                _ = fi19
            }
        }
    case "\n\t":
        v2 = (v2 ^ (-(v2)))
    }
    v29 := v3[127]
    fmt.Println(v2, v29)
    var v30 int32 = -1
    for fi31 := 0; fi31 < 12; fi31++ {
        v3[fi31] = (v29 ^ (v29 ^ (^(v29))))
        _ = fi31
    }
    v2 = v2
    var v32 int8 = 127
    v33 := GenS1{F0: 128}
    v2 = (v2 * v2)
    fmt.Println(v2, v2)
    v34 := GenS1{F0: 255}
    switch 541 {
    case 9223372036854775807:
        var v35 uint64 = 1
        _ = v35
    }
    v2 -= (v2 - -9223372036854775808)
    if false {
        var v36 float64 = 1e10
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
}
