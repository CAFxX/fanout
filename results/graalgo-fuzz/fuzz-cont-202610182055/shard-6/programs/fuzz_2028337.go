package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 bool; F1 string }
func genFn0(p0 int64) uint16 {
    fmt.Println(p0)
    p0 = p0
    for fi7 := 0; fi7 < 16; fi7++ {
        switch ("foo bar" + "") {
        case "a":
            if true {
                var v8 string = "\n\t"
                _ = v8
            }
        default:
            p0 = -842
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            p0 = p0
            _ = fi9
        }
        v10 := GenS1{F0: 754, F1: -1}
        _ = fi7
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = v11
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 string = (((("hello" + "hello") + "hello") + "\n\t") + "hello")
    var v3 int = 4294967295
    v4 := []string{v2, (v2 + v2), "hello", v2}
    var v5 int = v3
    var v6 bool = false
    v4 = append(v4, (("hello" + (("hello" + v2) + v2)) + ((v2 + "") + "\n\t")))
    if len(v4) > 0 {
        v4[0] = v2
    }
    _ = len(v4)
    v12 := GenS2{F0: 42, F1: v6}
    v2 = (v2 + (v2 + (("foo bar" + v2) + v2)))
    switch (v5 % ((255) | 1)) {
    case -759:
        v5 += (^(v3))
    case 4294967295:
        v2 = "\n\t"
    case -107:
        switch v5 {
        case 932:
            var v13 string = ((("foo bar" + "hello") + (v2 + v2)) + v2)
            _ = v13
        default:
            v2 = "\n\t"
        }
    default:
        var v14 uint = 4294967295
        _ = v14
    }
    v15 := GenS3{F0: -2147483648, F1: 2, F2: 1}
    if ((0 > -12) || (!v6)) {
        fmt.Println(v5, v6)
        switch (v2 + ("\n\t" + ((v2 + "") + (v2 + v2)))) {
        case "":
            v16 := GenS3{F0: 2147483647, F1: 495, F2: 127}
            _ = v16
        default:
            var v17 int64 = -208
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            if ((!v6) || (!(!(v6 || (v2 != v2))))) {
                v4 = append(v4, (v2 + (v2 + ((v2 + v2) + (((v2 + v2) + v2) + v2)))))
                switch ((v2 + ("" + v2)) + (("" + v2) + ("\n\t" + "hello"))) {
                case "":
                    fmt.Println(v1)
                case "hello":
                    fmt.Println(v3, v2, v1)
                default:
                    v2 = v2
                }
                if v6 {
                    var v19 uint8 = (v1 - (v1 / ((v1) | 1)))
                    v20 := GenS4{F0: (v2 != "a"), F1: v5}
                    var v21 int8 = 100
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v6 = v6
                    if (("\n\t" + "") <= (v2 + v2)) {
                        fmt.Println(v3)
                        switch v2 {
                        case "hello":
                            switch (v5 * -170) {
                            case -1:
                                v22 := GenS5{F0: v6}
                                _ = v22
                            }
                        case "\n\t":
                            var v23 uint16 = genFn0(-430)
                            _ = v23
                        }
                        switch v2 {
                        case "":
                            v24 := GenS1{F0: 191, F1: -32768}
                            _ = v24
                        case "a":
                            var v25 bool = ((v1 >> ((0) & 63)) != (v1 & (v1 - 0)))
                            _ = v25
                        case "hello":
                            v2 = (v2 + "a")
                        }
                    }
                }
            }
            if v6 {
                v1 = 1
                var v26 uint = 674
                fmt.Println(v2, v2, v3)
                _ = v26
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                if (v2 <= v2) {
                    fmt.Println(v3, v1)
                } else {
                    var v28 uint64 = 42
                    var v29 float64 = 3.14159
                    if false {
                        v1 = 2
                        fmt.Println(fi27)
                        if (!true) {
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                        } else {
                            if (1 >= 65535) {
                                if true {
                                    v29 = v29
                                }
                                var v30 uint32 = 2147483648
                                var v31 uint = 1
                                _ = v30
                                _ = v31
                            }
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            _ = fi18
        }
    }
    v32 := GenS6{F0: 255, F1: 173}
    if true {
        fmt.Println(v3, v3, v6)
        fmt.Println(v6, v2, v3)
        switch (v3 % ((v3) | 1)) {
        case 32767:
            v1 += v1
        case 32768:
            v2 = ("" + ("" + (v2 + v2)))
        }
    } else {
        v3 = v5
        if (v3 < 9223372036854775807) {
            switch (v2 + "foo bar") {
            case "\n\t":
                var v33 float64 = 0.0
                _ = v33
            default:
                v34 := GenS4{F0: v6, F1: 833}
                _ = v34
            }
            v4 = append(v4, "a")
        }
    }
    v35 := GenS7{F0: (v6 || (!v6)), F1: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
