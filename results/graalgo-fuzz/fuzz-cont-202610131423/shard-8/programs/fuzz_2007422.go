package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    v3 := []string{("" + ("foo bar" + "a")), "a", (("a" + ("\n\t" + "\n\t")) + "")}
    var v4 float32 = float32(-2.25)
    fmt.Println(v4, v2)
    var v5 int8 = 127
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: (-(fi7)), F1: 42, F2: 65535}
        _ = fi7
        _ = v8
    }
    var v9 uint = 18446744073709551615
    var v10 int16 = (v6 & v6)
    switch -363 {
    case 42:
        delete(v1, "hello")
    case -181:
        var v11 uint32 = 65535
        _ = v11
    case 425:
        if (v4 <= 0.1) {
            v2 += (^(v2))
            if false {
                var v12 string = "foo bar"
                fmt.Println(v5, v2)
                switch ("foo bar" + v12) {
                case "":
                    fmt.Println(v9)
                default:
                    v4 += v4
                }
                _ = v12
            }
            var v13 float64 = 3.14159
            _ = v13
        } else {
            v6 = v10
            var v14 bool = true
            _ = v14
        }
    }
    if true {
        if ((!false) || true) {
            v15 := GenS1{F0: 0, F1: 0, F2: 42}
            switch 128 {
            case 32767:
                var v16 int16 = v6
                _ = v16
            case 460:
                fmt.Println(v5, v6, v2)
            case 1:
                if (uint64(v9) <= 934) {
                    v17 := GenS1{F0: 1, F1: 0, F2: 42}
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v5 += (v5 ^ v5)
                        fmt.Println(fi18, v4)
                        _ = fi18
                    }
                    _ = v17
                }
            }
            _ = v15
        } else {
            if ((!false) || ((892 != 0) || true)) {
                if true {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int8 = v5
                        if false {
                            v6 = 0
                            var v21 bool = false
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 8; fi22++ {
                                v2 = 42
                                var v23 string = "foo bar"
                                _ = fi22
                                _ = v23
                            }
                            delete(v1, (((("a" + ("hello" + "hello")) + "\n\t") + ("a" + "")) + "hello"))
                            if true {
                                v24 := GenS2{F0: 42, F1: (^(v9))}
                                _ = v24
                            }
                        }
                        v25 := GenS3{F0: 4294967295, F1: 0}
                        _ = fi19
                        _ = v20
                        _ = v25
                    }
                } else {
                    v26 := GenS4{F0: -2147483648, F1: v9}
                    fmt.Println(v5, v9)
                    _ = v26
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    if true {
                        if false {
                            v28 := GenS2{F0: 0, F1: 1}
                            if false {
                                var v29 int32 = 2147483647
                                var v30 int = (fi27 ^ fi27)
                                fmt.Println(v2, fi27, v5)
                                _ = v29
                                _ = v30
                            }
                            fmt.Println(v2, v5, v9)
                            _ = v28
                        } else {
                            v31 := GenS5{F0: (163 >= 171)}
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v2 -= 42
                            var v33 uint32 = 1
                            _ = fi32
                            _ = v33
                        }
                        v2 += (v2 % ((v2) | 1))
                    }
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v1["a"] = "a"
                        v2 = -82
                        _ = fi34
                    }
                    _ = fi27
                }
                v6 = v6
            } else {
                v3 = append(v3, (("hello" + "hello") + ""))
            }
        }
        switch "\n\t" {
        case "a":
            if len(v3) > 0 {
                v3[0] = ("foo bar" + ("" + ("foo bar" + "foo bar")))
            }
        }
        v3 = append(v3, "a")
    } else {
        var v35 uint = v9
        v1["hello"] = "\n\t"
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := GenS5{F0: (!false)}
            v4 = (((v4 / ((float32(0.0) / float32(((3.14159 * 1.5) * (1.5 * (((-2.25 * 1e10) * 1.5) - 1.5))))) + v4)) + ((v4 * (-2.25 / 1.0)) * v4)) - v4)
            fmt.Println(v35)
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        v39 := GenS6{F0: 434}
        for fi40 := 0; fi40 < 8; fi40++ {
            var v41 int16 = (v10 | 32767)
            _ = fi40
            _ = v41
        }
        v42 := GenS7{F0: -851, F1: false}
        _ = fi38
        _ = v39
        _ = v42
    }
    fmt.Println(v5)
    var v43 float64 = 0.1
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
