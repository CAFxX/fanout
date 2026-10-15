package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 bool) int64 {
    fmt.Println(p1, p0)
    v6 := []int{0}
    _ = p0
    _ = p1
    _ = v6
    return 513
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 uint16 = 0
    var v5 int64 = -9223372036854775808
    v5 = v5
    v7 := GenS1{F0: 755, F1: (!(566 > (^(v4)))), F2: (v4 & v4)}
    if false {
        fmt.Println(v4, v5)
    } else {
        if (false || true) {
            var v8 uint32 = 65535
            v9 := []int{-2147483648, -1}
            var v10 uint = 255
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            switch "foo bar" {
            case "\n\t":
                var v12 uint8 = 127
                _ = v12
            case "foo bar":
                v5 = (v5 / ((-1) | 1))
            }
            switch (("foo bar" + "\n\t") + "foo bar") {
            case "a":
                var v13 float32 = -2.25
                _ = v13
            case "\n\t":
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS1{F0: -541, F1: ("a" < "hello"), F2: 32767}
                    _ = fi14
                    _ = v15
                }
            default:
                if true {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        if true {
                            fmt.Println(fi11)
                            v5 = v5
                        } else {
                            switch (fi16 * 0) {
                            case 586:
                                fmt.Println(fi16, fi16)
                            case -713:
                                if (!(!(false && (0.1 > float64(((1.5 - 0.1) * 1e10)))))) {
                                    if (583 < 2) {
                                        var v17 string = "a"
                                        fmt.Println(v17, v5)
                                        if (((true || false) && false) || true) {
                                            v4 -= 255
                                            var v18 int16 = -32768
                                            v19 := GenS2{F0: 32, F1: ((909 < v4) || true)}
                                            var v20 int32 = 0
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        }
                                        _ = v17
                                    }
                                    var v21 int = fi11
                                    _ = v21
                                }
                            case 756:
                                v4 -= v4
                            default:
                                v5 = v5
                            }
                            v4 += 65535
                            var v22 int16 = 0
                            _ = v22
                        }
                        var v23 int = fi16
                        _ = fi16
                        _ = v23
                    }
                    v4 += 1
                } else {
                    switch "" {
                    case "":
                        if true {
                            if true {
                                fmt.Println(v5, fi11, v5)
                                v4 = (v4 & 643)
                            }
                            var v24 string = (((("" + ("hello" + "\n\t")) + "\n\t") + "foo bar") + (("hello" + ("foo bar" + (((("a" + "foo bar") + "hello") + "hello") + ("hello" + "foo bar")))) + (("foo bar" + ("a" + ((("foo bar" + "") + "foo bar") + (("" + "a") + "\n\t")))) + "a")))
                            _ = v24
                        }
                    }
                }
            }
            _ = fi11
        }
    }
    v5 += (v5 % ((genFn0(0, (206 != 248))) | 1))
    var v25 int = 125
    v4 += 217
    var v26 float32 = -2.25
    v25 += v25
    fmt.Println(v4)
    for fi27 := 0; fi27 < 2; fi27++ {
        if (v5 >= -249) {
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(v25)
                v29 := GenS3{F0: v25}
                v30 := GenS2{F0: 127, F1: ((!true) || (false || (((!(!true)) || false) && (3.14159 > (1e10 * 0.0)))))}
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            v25 = fi27
        }
        _ = fi27
    }
    var v31 int8 = -1
    var v32 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
