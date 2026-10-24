package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 string }
func genFn0(p0 uint32) uint64 {
    p0 = (p0 << ((65535) & 63))
    var v4 uint8 = 244
    if false {
        switch "foo bar" {
        case "\n\t":
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 float64 = 0.1
                p0 = (p0 ^ (p0 + (p0 / (((p0 >> ((871) & 63))) | 1))))
                _ = fi5
                _ = v6
            }
        case "":
            v7 := GenS1{F0: true, F1: 32767, F2: 201}
            _ = v7
        default:
            v8 := GenS1{F0: (602 >= 4294967295), F1: 127, F2: 18446744073709551615}
            _ = v8
        }
    } else {
        var v9 uint32 = p0
        if ((v4 % ((188) | 1)) <= (v4 << ((65535) & 63))) {
            var v10 string = (("a" + "\n\t") + (("a" + "a") + (("" + "") + "")))
            var v11 float64 = ((((float64(((1e10 / float32(float64(((3.14159 * 0.0) + float32((0.0 + 1e10)))))) - 0.0)) + ((((0.0 + -2.25) * (1e10 - 0.1)) / 0.1) * (3.14159 / (1.5 / -2.25)))) / 1.0) * (1.5 + (1.5 - 0.0))) / float64(3.14159))
            _ = v10
            _ = v11
        }
        _ = v9
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 float64 = (0.0 * float64(1.5))
    var v2 uint64 = 0
    var v3 float64 = (v1 / 0.1)
    var v12 string = "hello"
    v13 := []int{127, 155, 0, -629}
    fmt.Println(v1, v2)
    v12 = (("" + ((("a" + ("hello" + "")) + v12) + v12)) + ("a" + v12))
    for fi14 := 0; fi14 < 10; fi14++ {
        v3 -= 1.5
        _ = fi14
    }
    switch 472 {
    case 32768:
        fmt.Println(v2)
    case -973:
        for fi15 := 0; fi15 < 5; fi15++ {
            switch fi15 {
            case 32767:
                var v16 float32 = ((((-2.25 - 1e10) - 0.0) + -2.25) * float32(v1))
                _ = v16
            case 2147483647:
                if false {
                    fmt.Println(fi15, v1)
                    fmt.Println(v3, v1, fi15)
                    var v17 int64 = -1
                    _ = v17
                } else {
                    v3 -= -2.25
                }
            case -671:
                if false {
                    v12 = (v12 + (((v12 + "foo bar") + "a") + ((v12 + "foo bar") + v12)))
                }
            default:
                var v18 int16 = -1
                _ = v18
            }
            _ = fi15
        }
    }
    switch 140 {
    case -1:
        if false {
            v3 -= 0.1
        } else {
            var v19 string = v12
            var v20 int16 = 1000
            fmt.Println(v19, v12, v20)
            _ = v19
            _ = v20
        }
    case 0:
        for fi21 := 0; fi21 < 14; fi21++ {
            if true {
                if (544 > 65535) {
                    var v22 float32 = 1e10
                    v23 := GenS2{F0: 0, F1: 210}
                    v24 := GenS1{F0: (!true), F1: (fi21 >> ((255) & 63)), F2: (v2 % ((v2) | 1))}
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    _ = len(v13)
                    var v25 uint64 = v2
                    _ = v25
                }
                var v26 int = fi21
                _ = v26
            } else {
                v27 := GenS3{F0: ("\n\t" + ("a" + v12)), F1: 0, F2: v12}
                switch (v12 + "hello") {
                case "":
                    var v28 uint16 = 255
                    _ = v28
                case "hello":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 uint8 = 1
                        v12 = (v12 + "hello")
                        _ = fi29
                        _ = v30
                    }
                case "a":
                    var v31 uint16 = 15
                    _ = v31
                default:
                    switch fi21 {
                    case -324:
                        switch "" {
                        case "foo bar":
                            for fi32 := 0; fi32 < 4; fi32++ {
                                if ((184 == 1) && true) {
                                    v3 = v1
                                } else {
                                    v12 = "a"
                                }
                                v33 := GenS1{F0: (true && false), F1: (fi32 + (fi21 >> ((65535) & 63))), F2: genFn0(2147483647)}
                                _ = fi32
                                _ = v33
                            }
                        }
                    default:
                        v34 := GenS2{F0: -793, F1: 2147483647}
                        _ = v34
                    }
                }
                _ = v27
            }
            _ = fi21
        }
    default:
        v35 := GenS1{F0: true, F1: 32768, F2: genFn0(2)}
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        if (true || (-459 == 1)) {
            var v37 uint8 = 0
            _ = v37
        } else {
            switch v12 {
            case "a":
                v12 = "\n\t"
            }
            if len(v13) > 0 {
                v13[0] = (-(fi36))
            }
        }
        _ = fi36
    }
    var v38 float64 = v3
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 uint32 = 2147483648
        _ = fi39
        _ = v40
    }
    _ = len(v13)
    var v41 uint32 = 4294967295
    var v42 bool = (-423 == 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v41
    _ = v42
}
