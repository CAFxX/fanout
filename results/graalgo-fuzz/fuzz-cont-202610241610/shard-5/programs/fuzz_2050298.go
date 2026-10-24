package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := []string{(("hello" + ("" + ("hello" + ""))) + "\n\t"), ("hello" + "foo bar"), "\n\t"}
    v4 := make(map[string]uint32)
    var v5 float64 = 0.0
    var v6 float64 = 1.5
    switch -66 {
    case 983:
        v7 := GenS1{F0: "\n\t", F1: -2147483648}
        _ = v7
    case 0:
        if true {
            v8 := GenS1{F0: ("hello" + ("" + "a")), F1: -2147483648}
            var v9 uint16 = 0
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                if ((3.14159 * 1.5) <= ((1.5 - (-2.25 * -2.25)) / (-2.25 / 3.14159))) {
                    delete(v1, 2)
                    if true {
                        v11 := GenS1{F0: ("" + "hello"), F1: fi10}
                        _ = v11
                    } else {
                        var v12 int64 = 183
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(v5, v6, fi13)
                        _ = len(v3)
                        var v14 int8 = -64
                        var v15 uint8 = 0
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v5 = ((0.1 / v5) * v5)
                        var v17 float64 = (float64(1.5) * v5)
                        _ = fi16
                        _ = v17
                    }
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
                var v18 uint64 = 535
                _ = fi10
                _ = v18
            }
        }
    }
    if false {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int = (fi19 | -9223372036854775808)
            _ = fi19
            _ = v20
        }
        _ = len(v3)
        for fi21 := 0; fi21 < 1; fi21++ {
            if (!(((("\n\t" + ((("foo bar" + "") + "hello") + ("foo bar" + ("a" + "a")))) + ("\n\t" + "")) + "hello") <= "\n\t")) {
                switch 42 {
                case 128:
                    fmt.Println(fi21, v5, v5)
                default:
                    var v22 uint32 = 822
                    _ = v22
                }
                fmt.Println(fi21, v6, v5)
            }
            var v23 float32 = 0.1
            for fi24 := 0; fi24 < 19; fi24++ {
                if (false || false) {
                    v5 -= v6
                } else {
                    v5 += 1.5
                    var v25 uint64 = 2
                    v25 += 224
                    _ = v25
                }
                _ = fi24
            }
            _ = fi21
            _ = v23
        }
    }
    _ = len(v4)
    switch 9223372036854775807 {
    case -2147483648:
        var v26 float64 = v6
        _ = v26
    case 529:
        var v27 int64 = 349
        _ = v27
    default:
        switch -2147483648 {
        case -9223372036854775808:
            v5 += (-2.25 * (float64(0.0) / 0.1))
        case 9223372036854775807:
            v3 = append(v3, "foo bar")
        }
    }
    var v28 uint32 = 4294967295
    if (!false) {
        if false {
            v1[978] = true
        } else {
            var v29 int = 32767
            _ = v29
        }
        v28 -= 0
    } else {
        fmt.Println(v5)
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int32 = -552
            _ = fi30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        if (!((!true) || ((!(true && (1 == 0))) && true))) {
            fmt.Println(v28, v28)
            delete(v4, "a")
        } else {
            var v33 string = "\n\t"
            _ = v33
        }
        _ = fi32
    }
    var v34 float64 = ((float64(0.1) - ((v5 / (v6 * (((0.0 * v5) * 1e10) * 3.14159))) - v5)) * 1e10)
    v35 := GenS1{F0: "", F1: 754}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
