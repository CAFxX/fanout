package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }

func main() {
    v1 := []int{127, -1, 127, 32767}
    var v2 int16 = 1
    var v3 float32 = 0.1
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 623}
            v2 += v2
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            if false {
                fmt.Println(fi7, v3, v3)
                if (((1e10 - 0.1) == float64(((((((3.14159 + 0.1) + v3) / v3) / 1e10) / v3) - (-2.25 * 1e10)))) || (!true)) {
                    v8 := GenS1{F0: 572}
                    if len(v1) > 0 {
                        v1[0] = fi4
                    }
                    _ = len(v1)
                    v2 = v2
                    var v9 string = (("a" + "hello") + ("" + ("foo bar" + ((("hello" + "\n\t") + (("\n\t" + "a") + ("" + ("foo bar" + (("" + "foo bar") + (("hello" + "\n\t") + "foo bar")))))) + ""))))
                    _ = v8
                    _ = v9
                }
            }
            switch -2147483648 {
            case -2147483648:
                v10 := GenS1{F0: 2}
                _ = v10
            }
            _ = fi7
        }
        v11 := GenS2{F0: 0, F1: "a", F2: 47}
        _ = fi4
        _ = v11
    }
    var v12 int = 1
    var v13 string = "foo bar"
    if false {
        var v14 string = "foo bar"
        switch v14 {
        case "":
            var v15 int64 = 34
            _ = v15
        default:
            v16 := GenS1{F0: 321}
            _ = v16
        }
        if (!false) {
            if true {
                switch v12 {
                case -146:
                    v2 = (v2 | v2)
                case -1:
                    fmt.Println(v14, v12, v12)
                default:
                    v17 := GenS3{F0: (!false), F1: 42, F2: 622}
                    _ = v17
                }
            } else {
                v3 = v3
                fmt.Println(v13, v2, v14)
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 uint64 = 36
                    _ = fi18
                    _ = v19
                }
            }
            fmt.Println(v14)
        }
        _ = v14
    }
    v2 += v2
    v12 = (v12 * 0)
    for fi20 := 0; fi20 < 16; fi20++ {
        v12 += 0
        _ = fi20
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        fmt.Println(v12, v3, v13)
    } else {
        var v21 string = "foo bar"
        _ = v21
    }
    if (true && (!(false && true))) {
        var v22 uint64 = 498
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
    } else {
        var v23 float64 = 1.5
        v13 = (v13 + v13)
        v12 += 2147483647
        _ = v23
    }
    fmt.Println(v3, v12, v3)
    v12 = (v12 & v12)
    v12 -= 42
    var v24 uint = 54
    _ = len(v1)
    fmt.Println(v13, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
}
