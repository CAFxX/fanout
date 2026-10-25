package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint8 }

func main() {
    var v1 int16 = -602
    var v2 int32 = -2147483648
    v3 := []string{}
    v4 := make(map[uint32]bool)
    if false {
        var v5 uint8 = 94
        switch 4294967295 {
        case 48:
            var v6 string = "foo bar"
            _ = v6
        case -157:
            var v7 uint64 = 0
            _ = v7
        default:
            var v8 uint8 = 42
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            v1 += v1
            _ = fi9
        }
        _ = v5
    } else {
        if (42 > -89) {
            v2 += v2
            v1 += (-(v1))
            v2 -= (v2 * (v2 | 1))
        }
        var v10 bool = (!true)
        _ = v10
    }
    var v11 int8 = 127
    if false {
        switch 42 {
        case 32768:
            v2 += (v2 >> ((42) & 63))
        case 42:
            v12 := GenS1{F0: "", F1: ((false || true) || (0.0 >= 0.0)), F2: 128}
            _ = v12
        }
        fmt.Println(v2, v11, v11)
    }
    var v13 int8 = v11
    if (((!(0 <= (v2 | v2))) && true) || true) {
        var v14 int32 = v2
        var v15 float32 = -2.25
        var v16 int32 = v2
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint64 = 0
    if (2 < 2) {
        if (0 > -411) {
            v18 := GenS1{F0: "\n\t", F1: true, F2: 0}
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v11, v1, v1)
            v17 -= v17
            _ = fi19
        }
        var v20 int = int(v11)
        _ = v20
    }
    v1 = (v1 % ((553) | 1))
    if ((false || true) && (42 != 185)) {
        switch -509 {
        case 2147483647:
            if (!(("hello" + ("\n\t" + "")) > (("a" + "hello") + "a"))) {
                delete(v4, 673)
            } else {
                var v21 int64 = 1
                var v22 string = "a"
                _ = v21
                _ = v22
            }
        default:
            if ((!false) && ((!true) && true)) {
                switch "foo bar" {
                case "hello":
                    switch (((("a" + "hello") + ("\n\t" + (("a" + ("hello" + "a")) + ""))) + "") + ("\n\t" + ("a" + ("" + "")))) {
                    case "a":
                        var v23 uint = 4294967295
                        _ = v23
                    }
                default:
                    fmt.Println(v13, v1, v11)
                }
                var v24 float64 = 0.0
                _ = v24
            }
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        switch -278 {
        case 249:
            if (128 == 255) {
                v13 = (v11 & 42)
                var v25 uint64 = v17
                _ = v25
            }
        }
    }
    if false {
        v17 -= 18446744073709551615
        if false {
            var v26 int32 = v2
            fmt.Println(v2, v11)
            if true {
                v3 = append(v3, "foo bar")
            }
            _ = v26
        }
        v2 -= -1
    } else {
        if true {
            switch "foo bar" {
            case "hello":
                v11 = v11
            case "":
                var v27 float64 = (0.1 - 0.1)
                _ = v27
            case "\n\t":
                fmt.Println(v13, v13, v17)
            }
            v11 -= v13
            v13 = 100
        } else {
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int = fi28
                var v30 int16 = v1
                v30 = v30
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 float32 = 0.1
            _ = v31
        }
        var v32 float64 = 3.14159
        fmt.Println(v2, v17, v2)
        _ = v32
    }
    var v33 float32 = (float32((float64(3.14159) + (3.14159 + 0.0))) + 0.1)
    var v34 uint = 255
    var v35 uint32 = 2147483647
    v34 -= v34
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
