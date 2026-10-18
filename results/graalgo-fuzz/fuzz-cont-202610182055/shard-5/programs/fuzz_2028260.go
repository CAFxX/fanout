package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint; F2 int16 }
func genFn0() uint {
    var v6 string = "hello"
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint8 = 255
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: ((v6 + v6) + ("a" + (v6 + ""))), F1: 154, F2: false}
    v10 := GenS2{F0: 42, F1: ((!((true && true) || (0 <= 1000))) || (!(!(59 == 2147483647))))}
    var v11 int = 983
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 397
}

func main() {
    var v1 int64 = 736
    var v2 uint8 = 107
    v3 := make(map[string]int)
    var v4 float64 = (0.1 + (-2.25 * 0.0))
    var v5 int64 = v1
    v12 := []int{-2147483648, 0, 127}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v13 float32 = ((((0.0 * 1e10) * (-2.25 * 0.0)) + 0.1) / (3.14159 + (-2.25 * 3.14159)))
    v5 = v5
    if len(v12) > 0 {
        v12[0] = 287
    }
    v13 = v13
    if (3.14159 >= v4) {
        if false {
            if len(v12) > 0 {
                v12[0] = 259
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                if (2 <= 1) {
                    _ = len(v12)
                    var v15 float64 = float64((1e10 / 1.5))
                    _ = v15
                } else {
                    v16 := GenS3{F0: 18446744073709551615, F1: 65535, F2: 65535}
                    _ = v16
                }
                _ = fi14
            }
            v17 := GenS4{F0: "foo bar", F1: 65535}
            _ = v17
        }
    }
    var v18 int16 = 0
    if len(v12) > 0 {
        v12[0] = 843
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if ("a" >= ("foo bar" + "hello")) {
            var v20 uint = 18446744073709551615
            if false {
                v13 -= (v13 / (3.14159 * -2.25))
                fmt.Println(fi19, v2)
                if false {
                    if false {
                        fmt.Println(v4, v5, v1)
                    } else {
                        var v21 bool = false
                        _ = v21
                    }
                }
            }
            _ = v20
        }
        fmt.Println(v5)
        _ = fi19
    }
    v1 = (^(v5))
    switch -2147483648 {
    case -297:
        v2 += 61
    case 247:
        var v22 bool = false
        _ = v22
    case 4294967295:
        if true {
            var v23 float32 = v13
            _ = v23
        } else {
            var v24 bool = (v13 > v13)
            _ = v24
        }
    default:
        if false {
            if true {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            switch "hello" {
            case "hello":
                if (!true) {
                    _ = len(v12)
                }
            case "":
                var v25 string = ((("" + "\n\t") + "\n\t") + "\n\t")
                _ = v25
            }
            fmt.Println(v1, v2, v2)
        } else {
            var v26 uint64 = 2
            if true {
                v27 := GenS4{F0: ((("\n\t" + "\n\t") + (("\n\t" + ((("\n\t" + ("hello" + "")) + "foo bar") + "a")) + (("" + "hello") + "foo bar"))) + ""), F1: 913}
                _ = v27
            }
            if (v1 == (v1 >> ((65535) & 63))) {
                var v28 uint8 = (v2 - (^(v2)))
                for fi29 := 0; fi29 < 8; fi29++ {
                    fmt.Println(v2, v28)
                    _ = fi29
                }
                delete(v3, "")
                _ = v28
            } else {
                v5 = v5
                var v30 int32 = 2147483647
                _ = v30
            }
            _ = v26
        }
    }
    fmt.Println(v2)
    var v31 int = 42
    fmt.Println(v13, v13)
    var v32 uint16 = 65535
    switch ("a" + "foo bar") {
    case "a":
        v33 := GenS5{F0: v18, F1: 4294967295, F2: 1000}
        _ = v33
    case "hello":
        var v34 uint16 = v32
        _ = v34
    default:
        var v35 uint = genFn0()
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
