package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
func genFn0(p0 int16, p1 int, p2 uint8) int8 {
    p0 = p0
    if false {
        var v7 float32 = float32(1.5)
        v8 := []int{p1, (p1 ^ p1), (p1 >> ((0) & 63)), 32768}
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0)
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    v2 = v3
    if true {
        var v9 float64 = 3.14159
        fmt.Println(v2, v2, v2)
        _ = v9
    }
    var v10 string = ("a" + ("foo bar" + (v3 + v2)))
    v11 := v5[564]
    for fi12 := 0; fi12 < 9; fi12++ {
        if false {
            var v13 int16 = 0
            switch v10 {
            case "foo bar":
                switch -2147483648 {
                case -9223372036854775808:
                    v2 = (v2 + v3)
                case 128:
                    v3 = v2
                case 32767:
                    for fi14 := 0; fi14 < 2; fi14++ {
                        delete(v6, ("\n\t" + v3))
                        v10 = v10
                        _ = fi14
                    }
                default:
                    if v11 {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v2 = "hello"
                            var v16 int = 127
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        var v17 int16 = (v13 - (v13 * (-(v13))))
                        _ = len(v6)
                        switch 283 {
                        case 9223372036854775807:
                            v3 = ("hello" + (v10 + (v10 + "")))
                        case -378:
                            var v18 int = fi12
                            _ = v18
                        }
                        _ = v17
                    }
                }
            case "":
                v2 = "a"
            }
            if (v11 && v11) {
                v19 := GenS1{F0: 0, F1: 1, F2: v13}
                var v20 uint16 = 0
                var v21 uint16 = 2
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v13
        }
        if (("foo bar" + v3) == ("foo bar" + (v2 + v10))) {
            var v22 uint8 = 128
            _ = v22
        }
        _ = fi12
    }
    fmt.Println(v2)
    switch ((v2 + "foo bar") + ((v10 + ("foo bar" + v3)) + v10)) {
    case "foo bar":
        var v23 int8 = 127
        _ = v23
    default:
        for fi24 := 0; fi24 < 18; fi24++ {
            v11 = (!(!v11))
            fmt.Println(v10)
            v25 := []int{1, (fi24 & (fi24 | fi24)), fi24}
            _ = fi24
            _ = v25
        }
    }
    v11 = v11
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint = 2
        _ = fi26
        _ = v27
    }
    switch 735 {
    case -723:
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 string = "\n\t"
            _ = fi28
            _ = v29
        }
    case 136:
        var v30 int16 = -540
        _ = v30
    default:
        v2 = v3
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint8 = 1
        fmt.Println(v2, v3, fi31)
        var v33 uint = 255
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
