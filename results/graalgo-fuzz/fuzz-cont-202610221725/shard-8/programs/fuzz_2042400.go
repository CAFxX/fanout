package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 uint32, p1 uint32, p2 int64) int8 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 827
    }
    var v8 uint = 346
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int8) bool {
    v9 := []int{2147483647, -9223372036854775808, -644, 32768}
    if (!false) {
        fmt.Println(p0)
        switch 2147483647 {
        case 32768:
            switch -2147483648 {
            case -939:
                switch "foo bar" {
                case "foo bar":
                    if (!true) {
                        fmt.Println(p0)
                        v10 := GenS1{F0: "foo bar", F1: 255, F2: 401}
                        v9 = append(v9, 4294967295)
                        _ = v10
                    } else {
                        var v11 bool = (65535 > 123)
                        _ = v11
                    }
                case "a":
                    p0 = genFn0(2, 4, 1)
                case "":
                    for fi12 := 0; fi12 < 14; fi12++ {
                        if true {
                            switch 810 {
                            case -270:
                                if true {
                                    p0 += genFn0(520, 65535, 1)
                                    p0 -= 0
                                    p0 = (p0 * p0)
                                }
                            }
                            p0 = p0
                        } else {
                            fmt.Println(fi12, p0)
                        }
                        var v13 int16 = -709
                        _ = fi12
                        _ = v13
                    }
                default:
                    if ((!false) && true) {
                        var v14 int16 = -865
                        _ = v14
                    } else {
                        p0 += p0
                    }
                }
            }
        case 128:
            p0 -= (p0 - genFn0(2, 947, 42))
        }
        v15 := GenS1{F0: "", F1: 65535, F2: 32767}
        _ = v15
    }
    switch ("a" + "") {
    case "hello":
        for fi16 := 0; fi16 < 8; fi16++ {
            switch (^(fi16)) {
            case 227:
                fmt.Println(p0, fi16)
            }
            _ = fi16
        }
    }
    if len(v9) > 0 {
        v9[0] = -839
    }
    switch -1 {
    case 1:
        var v17 uint8 = 1
        _ = v17
    case 127:
        var v18 uint64 = 2
        _ = v18
    }
    _ = p0
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 161
    var v2 int = 32767
    var v3 bool = ((!(65535 == (v1 / ((v1) | 1)))) || true)
    var v4 uint = 1
    var v5 int = int(v1)
    var v6 int16 = -32768
    for fi19 := 0; fi19 < 13; fi19++ {
        if (v3 && false) {
            v4 = (^(v4))
            v20 := []int{(v5 + fi19), int(v1), (v2 + v2)}
            _ = v20
        } else {
            fmt.Println(v4)
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            switch "" {
            case "":
                var v22 int16 = -1
                _ = v22
            case "hello":
                var v23 uint = v4
                _ = v23
            default:
                if ((4294967295 < v4) && v3) {
                    v5 += (v2 + (^(fi21)))
                    switch -618 {
                    case -801:
                        var v24 uint32 = 2147483647
                        _ = v24
                    case 32768:
                        v6 = v6
                    }
                    var v25 float64 = 1.5
                    _ = v25
                } else {
                    v26 := GenS1{F0: "foo bar", F1: v4, F2: 42}
                    if ((!v3) && ((int8(v6) > 17) || v3)) {
                        v3 = v3
                    }
                    _ = v26
                }
            }
            _ = fi21
        }
        _ = fi19
    }
    var v27 uint8 = 42
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS1{F0: ("a" + "a"), F1: (v4 % ((255) | 1)), F2: v1}
        _ = fi28
        _ = v29
    }
    fmt.Println(v5)
    var v30 uint16 = 255
    if false {
        v31 := []int{v2, (v5 >> ((v4) & 63)), v5}
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 int16 = 0
            _ = fi32
            _ = v33
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            v35 := GenS1{F0: ("\n\t" + ("hello" + ("hello" + ""))), F1: (v4 - (^(v4))), F2: (v30 >> ((325) & 63))}
            _ = fi34
            _ = v35
        }
        _ = v31
    }
    v6 -= v6
    v5 = (v5 & (v5 % (((v2 << (((^(v4))) & 63))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
