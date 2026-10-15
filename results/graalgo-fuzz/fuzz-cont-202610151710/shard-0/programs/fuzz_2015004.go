package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0(p0 int32, p1 uint) uint16 {
    v4 := []int{-583, 127}
    p1 = (p1 >> ((42) & 63))
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "", ""}
    var v3 uint16 = 160
    _ = len(v2)
    v5 := GenS1{F0: -9223372036854775808, F1: 42}
    if (((905 != 65535) && false) && (100 < 100)) {
        var v6 int64 = 42
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float32 = 1e10
            var v9 int16 = 0
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(v8)
                var v11 bool = (!true)
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        if (((("hello" + "") + "hello") + "hello") <= ("a" + "foo bar")) {
            v1[4294967295] = (1e10 > 1e10)
            v6 = -1
            var v12 float32 = ((float32((-2.25 - 0.1)) * 0.1) * 1e10)
            _ = v12
        }
        _ = v6
    } else {
        v3 += 648
        _ = len(v2)
    }
    if false {
        _ = len(v1)
        fmt.Println(v3)
        switch 1 {
        case -934:
            v3 = v3
        case 630:
            var v13 uint64 = 2
            _ = v13
        case 127:
            var v14 float64 = 0.0
            _ = v14
        default:
            switch "a" {
            case "\n\t":
                if (0.0 <= (1.5 / 0.1)) {
                    if (!(false && false)) {
                        for fi15 := 0; fi15 < 2; fi15++ {
                            for fi16 := 0; fi16 < 17; fi16++ {
                                if ((!(!false)) && (!(!true))) {
                                    var v17 int64 = 0
                                    v17 -= -1
                                    var v18 int32 = -1
                                    _ = v17
                                    _ = v18
                                }
                                v3 = v3
                                _ = fi16
                            }
                            v19 := GenS1{F0: 175, F1: 15}
                            _ = fi15
                            _ = v19
                        }
                        v3 = (v3 | (v3 ^ v3))
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v21 := GenS1{F0: -9223372036854775808, F1: 42}
                            var v22 uint64 = 18446744073709551615
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 uint64 = 42
                        v25 := GenS1{F0: -102, F1: 17}
                        v24 -= (^(v24))
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
            }
        }
    }
    v3 = v3
    v26 := GenS1{F0: 1, F1: 1}
    var v27 int16 = -185
    v3 = v3
    var v28 float64 = ((((1.5 / 3.14159) - 0.1) - float64(0.0)) - ((1e10 + float64((1.5 - float32(0.0)))) / 3.14159))
    switch 9223372036854775807 {
    case -314:
        v3 += 0
    }
    var v29 float64 = float64((3.14159 - -2.25))
    if true {
        if ("hello" > "\n\t") {
            var v30 string = ("foo bar" + ((("foo bar" + "") + "") + ""))
            var v31 int = -9223372036854775808
            v28 = (v29 * 1e10)
            _ = v30
            _ = v31
        }
        if (false && true) {
            v1[65535] = true
            if (true && true) {
                v27 = v27
                switch "\n\t" {
                case "foo bar":
                    var v32 uint64 = 18446744073709551615
                    _ = v32
                case "":
                    fmt.Println(v27, v27, v28)
                case "a":
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + ("hello" + "foo bar"))
                    }
                }
            }
            v3 -= 720
        }
    } else {
        v33 := GenS1{F0: -1, F1: 0}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v5
}
