package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 int64 }

func main() {
    var v1 bool = (0.1 == ((float64(1.5) + (3.14159 + float64((float32(3.14159) + 1.5)))) / -2.25))
    v2 := []uint32{}
    v3 := []string{"hello"}
    v4 := make(map[string]string)
    var v5 int16 = 1
    v1 = v1
    var v6 int8 = 127
    delete(v4, "a")
    for fi7 := 0; fi7 < 1; fi7++ {
        if v1 {
            delete(v4, "hello")
        } else {
            if true {
                fmt.Println(v1, fi7)
                v8 := GenS1{F0: (fi7 | fi7), F1: 2}
                _ = v8
            }
        }
        fmt.Println(fi7, v6)
        _ = fi7
    }
    var v9 uint64 = 42
    for fi10 := 0; fi10 < 2; fi10++ {
        if (-271 != 42) {
            if v1 {
                switch -204 {
                case 0:
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v9 += v9
                        _ = fi11
                    }
                case 128:
                    v6 = (^(v6))
                default:
                    if true {
                        fmt.Println(v5, v1)
                        v5 = v5
                    } else {
                        v9 = v9
                        if (true || (727 >= 0)) {
                            var v12 uint16 = uint16(v9)
                            delete(v4, ("foo bar" + "a"))
                            _ = v12
                        }
                    }
                }
                v5 += (v5 / ((v5) | 1))
            } else {
                var v13 int64 = 42
                v14 := GenS2{F0: (!(710 > (^(v9))))}
                v1 = v1
                _ = v13
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 bool = true
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int16 = int16(v6)
            var v19 int8 = (v6 | v6)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi10
    }
    v5 = 334
    for fi20 := 0; fi20 < 19; fi20++ {
        v9 += v9
        v21 := GenS2{F0: v1}
        var v22 int8 = v6
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = 145
    v1 = (750 <= 65535)
    v5 -= (v5 | (-(v5)))
    if true {
        for fi24 := 0; fi24 < 14; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                switch (fi25 >> ((131) & 63)) {
                case 127:
                    if (!v1) {
                        var v26 int = fi24
                        v26 = 9223372036854775807
                        _ = v26
                    }
                    _ = len(v3)
                case 609:
                    v27 := GenS2{F0: (0 < (v9 | v9))}
                    _ = v27
                case 2147483647:
                    if ((^(fi25)) >= (-(fi25))) {
                        switch (("hello" + ("" + "")) + "foo bar") {
                        case "a":
                            var v28 uint32 = 2
                            _ = v28
                        case "foo bar":
                            switch (fi25 % ((127) | 1)) {
                            case 42:
                                var v29 string = (("" + (("foo bar" + ((("foo bar" + "foo bar") + "foo bar") + (("\n\t" + ((("hello" + "hello") + ("a" + ("a" + (("" + "a") + "\n\t")))) + "hello")) + "hello"))) + "a")) + "hello")
                                _ = v29
                            case -978:
                                v30 := GenS2{F0: v1}
                                _ = v30
                            }
                        default:
                            var v31 uint16 = 1
                            _ = v31
                        }
                        v6 -= (v6 | (v6 * (-(v6))))
                    } else {
                        v6 += (v6 - v6)
                    }
                default:
                    if v1 {
                        switch (fi24 * (-(fi25))) {
                        case 285:
                            _ = len(v4)
                        }
                    }
                }
                v32 := GenS3{F0: v9, F1: false, F2: -1}
                var v33 uint16 = 2
                _ = fi25
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 3; fi34++ {
                delete(v4, "")
                _ = fi34
            }
            _ = fi24
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(v23, v5, v5)
            _ = fi35
        }
        v1 = false
    }
    var v36 bool = (9223372036854775807 <= 128)
    v37 := GenS2{F0: ((true || (-2.25 >= ((float64(-2.25) * 1e10) * ((1.5 / 1e10) + float64(((1.5 / 1.0) - ((((1.5 / 1e10) - 3.14159) - 0.1) + 1.5))))))) || (v23 <= v23))}
    _ = len(v4)
    var v38 float64 = 3.14159
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
