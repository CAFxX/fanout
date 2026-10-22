package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }

func main() {
    var v1 int16 = -896
    var v2 uint64 = 18446744073709551615
    var v3 float64 = (((((((float64(((1.5 + 3.14159) / 1e10)) + 1.5) + 0.0) - -2.25) / 1.5) + -2.25) - ((1.5 * ((((3.14159 - 0.0) * (-2.25 - float64(0.1))) - 0.1) / 1.5)) * 1.5)) / 1e10)
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v6 uint16 = 255
    v5[42] = "a"
    var v7 int16 = (^(v1))
    fmt.Println(v7, v3, v6)
    v8 := []int{2147483647, -9223372036854775808, -9223372036854775808, 32767}
    v9 := GenS1{F0: false, F1: 9223372036854775807, F2: ((("" + ("foo bar" + "\n\t")) + "foo bar") + "\n\t")}
    if false {
        var v10 int16 = v1
        var v11 uint32 = 135
        v5[32768] = "a"
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int64 = 9223372036854775807
            if false {
                if true {
                    switch fi12 {
                    case 128:
                        switch int(v2) {
                        case -842:
                            switch "foo bar" {
                            case "\n\t":
                                v3 = ((v3 + (-2.25 * v3)) * float64(0.1))
                            case "a":
                                if false {
                                    switch (fi12 / ((fi12) | 1)) {
                                    case 9223372036854775807:
                                        var v14 string = ("a" + ("a" + "\n\t"))
                                        _ = v14
                                    case 2147483647:
                                        var v15 int8 = 1
                                        _ = v15
                                    }
                                    v16 := v5[fi12]
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v16
                                } else {
                                    var v17 int16 = 1000
                                    v17 += (-(v17))
                                    switch 0 {
                                    case -2147483648:
                                        v18 := GenS1{F0: ("hello" != "a"), F1: v13, F2: "a"}
                                        _ = v18
                                    case 127:
                                        v19 := GenS1{F0: ((((!(42 < 124)) && true) || ((0.1 * 1.5) <= float64(1e10))) || (42 < 0)), F1: v13, F2: "hello"}
                                        _ = v19
                                    case 2147483647:
                                        var v20 int32 = 2147483647
                                        _ = v20
                                    }
                                    _ = v17
                                }
                            }
                        case 712:
                            v21 := GenS1{F0: true, F1: (v13 - (v13 ^ -689)), F2: ("" + "foo bar")}
                            _ = v21
                        }
                    case -2147483648:
                        v6 += uint16(v2)
                    default:
                        if true {
                            _ = len(v8)
                            v1 += v7
                        } else {
                            var v22 uint = 65535
                            fmt.Println(v13, v22, v22)
                            var v23 string = ""
                            _ = v22
                            _ = v23
                        }
                    }
                }
                fmt.Println(fi12)
            }
            var v24 int16 = v7
            _ = fi12
            _ = v13
            _ = v24
        }
    }
    v3 = v3
    v2 -= 151
    fmt.Println(v7, v2)
    for fi25 := 0; fi25 < 9; fi25++ {
        switch "a" {
        case "":
            v1 -= (v1 >> ((2) & 63))
        }
        _ = fi25
    }
    switch -545 {
    case 788:
        v4[9223372036854775807] = 0
    case 127:
        fmt.Println(v6, v7)
    case -211:
        if ((("hello" + "\n\t") + ("\n\t" + "\n\t")) != "a") {
            var v26 uint64 = (v2 - (v2 - 863))
            _ = v26
        } else {
            _ = len(v4)
            delete(v5, 42)
            var v27 bool = ((-605 <= 42) || true)
            _ = v27
        }
    }
    v6 += v6
    var v28 uint16 = v6
    var v29 string = ("hello" + ("hello" + (("" + "") + ("\n\t" + "hello"))))
    v3 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
