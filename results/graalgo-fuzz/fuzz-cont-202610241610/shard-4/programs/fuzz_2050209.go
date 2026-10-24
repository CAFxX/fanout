package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 uint, p2 int16) bool {
    p1 = 255
    fmt.Println(p1, p2, p2)
    p1 = (p1 / ((p1) | 1))
    var v7 uint32 = 724
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]bool)
    var v3 int32 = 2147483647
    var v4 string = ("\n\t" + ("a" + ""))
    v5 := make(map[uint32]string)
    v6 := []string{v4}
    if len(v6) > 0 {
        v6[0] = v4
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        if (-2.25 == 1e10) {
            var v9 string = v4
            var v10 uint8 = 127
            v11 := v2[(("\n\t" + ("\n\t" + v4)) + "hello")]
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 int8 = 85
            fmt.Println(v4, v3, v12)
            _ = v12
        }
        v13 := GenS1{F0: 0, F1: fi8, F2: 42}
        v1 += v1
        _ = fi8
        _ = v13
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v1)
    v4 = ("hello" + "foo bar")
    if ((v4 < ((("a" + "a") + "a") + v4)) && false) {
        for fi14 := 0; fi14 < 2; fi14++ {
            if (false && true) {
                v1 -= 0
                switch v4 {
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                v4 = v4
            } else {
                v15 := GenS1{F0: 0, F1: fi14, F2: 493}
                switch -223 {
                case 719:
                    var v16 bool = (1 < 42)
                    _ = v16
                case -1:
                    var v17 float32 = 1.5
                    _ = v17
                }
                _ = v15
            }
            _ = fi14
        }
    }
    switch 4294967295 {
    case -1:
        if (-1 != 9223372036854775807) {
            v18 := v2[v4]
            _ = v18
        }
    case -341:
        v1 = 255
    case -880:
        switch 32767 {
        case 0:
            var v19 float32 = 1e10
            _ = v19
        case -28:
            switch v4 {
            case "a":
                fmt.Println(v3, v4)
            case "hello":
                for fi20 := 0; fi20 < 12; fi20++ {
                    if false {
                        if false {
                            var v21 int8 = 42
                            fmt.Println(v21, fi20, v1)
                            _ = v21
                        }
                        v6 = append(v6, v4)
                        v22 := GenS1{F0: 255, F1: (fi20 | (fi20 | fi20)), F2: 42}
                        _ = v22
                    }
                    var v23 uint = v1
                    if false {
                        var v24 uint16 = 2
                        switch 1 {
                        case 555:
                            v23 -= 0
                        case -861:
                            v25 := GenS1{F0: 128, F1: (fi20 & 0), F2: int16(fi20)}
                            _ = v25
                        case 229:
                            v26 := GenS1{F0: 255, F1: fi20, F2: -32768}
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 10; fi27++ {
                                v1 = 1
                                var v28 uint8 = 2
                                _ = fi27
                                _ = v28
                            }
                        }
                        var v29 uint16 = v24
                        _ = v24
                        _ = v29
                    }
                    _ = fi20
                    _ = v23
                }
            case "":
                if (3.14159 <= float32(3.14159)) {
                    var v30 uint32 = 2147483647
                    _ = v30
                } else {
                    _ = len(v5)
                }
            }
        }
    default:
        v31 := GenS1{F0: 1, F1: 1, F2: -32768}
        _ = v31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v6 = append(v6, v4)
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
