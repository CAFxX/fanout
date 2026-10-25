package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint; F2 string }
func genFn0(p0 int8, p1 uint32) string {
    var v7 uint8 = 255
    v8 := []int{2, 32767, 42}
    p1 += (p1 >> ((0) & 63))
    var v9 bool = false
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ("a" + "hello")
}

func main() {
    var v1 uint8 = 2
    var v2 uint16 = 65535
    var v3 uint32 = 999
    v4 := []string{"hello", "foo bar", ("foo bar" + (("foo bar" + (((("hello" + "") + ("hello" + "a")) + "foo bar") + "a")) + "hello"))}
    v5 := make(map[string]int)
    var v6 int64 = -1
    v1 -= 1
    var v10 bool = false
    if (false && (int64(v2) == 707)) {
        v11 := GenS1{F0: (0 >= 140), F1: v2}
        _ = v11
    }
    var v12 uint8 = v1
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 2; fi14++ {
        if (-1 < -1) {
            v15 := GenS1{F0: ((226 != -547) && (0 != 18)), F1: (v2 * v2)}
            var v16 uint32 = 66
            if false {
                v17 := GenS2{F0: 2, F1: uint(v2), F2: v13}
                _ = v17
            }
            _ = v15
            _ = v16
        } else {
            var v18 int = fi14
            _ = v18
        }
        if v10 {
            if v10 {
                if ((v3 & v3) <= (v3 >> ((148) & 63))) {
                    switch "" {
                    case "a":
                        var v19 int64 = -414
                        _ = v19
                    case "\n\t":
                        for fi20 := 0; fi20 < 6; fi20++ {
                            _ = len(v5)
                            switch fi20 {
                            case 26:
                                var v21 uint64 = 65
                                _ = v21
                            default:
                                if (3.14159 == 1.5) {
                                    v22 := GenS1{F0: (1 < 42), F1: v2}
                                    var v23 uint16 = v2
                                    fmt.Println(v6, fi20, v23)
                                    v24 := GenS1{F0: true, F1: v2}
                                    if v10 {
                                        v3 = uint32(v23)
                                        var v25 string = v13
                                        _ = v25
                                    }
                                    var v26 int64 = 462
                                    if ((v2 + (v2 & uint16(v3))) < (^(v2))) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            delete(v5, v13)
                                        default:
                                            switch (v13 + ("foo bar" + v13)) {
                                            case "\n\t":
                                                v12 += (v1 - 177)
                                            case "foo bar":
                                                v13 = v13
                                            }
                                        }
                                        v6 = 9223372036854775807
                                        var v27 float32 = -2.25
                                        _ = v27
                                    }
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v26
                                }
                            }
                            if ((v2 > v2) && v10) {
                                v3 = 973
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if true {
                                    if (!(v10 && v10)) {
                                        fmt.Println(v6)
                                    } else {
                                        _ = len(v4)
                                    }
                                }
                                v13 = v13
                            }
                            v6 += v6
                            _ = len(v4)
                            _ = fi20
                        }
                    default:
                        if v10 {
                            v10 = (true && v10)
                        }
                    }
                    v13 = v13
                }
                if v10 {
                    if v10 {
                        var v28 string = "foo bar"
                        _ = v28
                    } else {
                        v1 = 2
                        fmt.Println(v3, v10)
                        v13 = ""
                    }
                }
            }
        }
        _ = fi14
    }
    if len(v4) > 0 {
        v4[0] = v13
    }
    v29 := v5[v13]
    v30 := GenS1{F0: ("a" <= (v13 + "\n\t")), F1: (v2 & (v2 - (^(v2))))}
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v1 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
