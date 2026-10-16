package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    var v1 bool = true
    v2 := []int{42, 4294967295, 0, -258}
    var v3 int64 = -663
    v4 := make(map[string]uint32)
    var v5 float64 = ((((0.0 - 1e10) * 1e10) + ((-2.25 + ((0.0 + -2.25) + (1.5 + -2.25))) - (float64(0.1) * (1.5 + 0.0)))) * (3.14159 * (0.0 * (-2.25 + 0.0))))
    if ("\n\t" > "a") {
        v6 := GenS1{F0: 1}
        var v7 string = ((("" + "a") + "") + ("a" + "a"))
        _ = v6
        _ = v7
    } else {
        var v8 float32 = 0.1
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 string = (("" + "") + ("a" + "a"))
            _ = fi9
            _ = v10
        }
        fmt.Println(v8, v5, v3)
        _ = v8
    }
    v11 := v4["a"]
    var v12 int = 995
    fmt.Println(v3, v11)
    var v13 int16 = 1000
    switch "" {
    case "":
        if v1 {
            v4["a"] = 426
            var v14 uint16 = 32767
            _ = v14
        }
    }
    fmt.Println(v1)
    v15 := GenS1{F0: v13}
    if v1 {
        if v1 {
            if (((("\n\t" + (((("hello" + "foo bar") + "foo bar") + "hello") + "")) == "") && (!(-45 >= 1))) && false) {
                switch "a" {
                case "\n\t":
                    if ((!((!(303 == 2)) || (false || (true && true)))) || (v1 || v1)) {
                        var v16 int64 = 1
                        v11 += (v11 - v11)
                        v13 += 0
                        _ = v16
                    }
                case "":
                    fmt.Println(v13, v5, v12)
                case "hello":
                    if (0 == 42) {
                        var v17 uint = 18446744073709551615
                        _ = v17
                    } else {
                        v13 += (v13 % (((v13 % ((v13) | 1))) | 1))
                    }
                }
                if v1 {
                    fmt.Println(v12)
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch ("" + (("a" + "") + "")) {
                    case "foo bar":
                        v19 := GenS2{F0: 1, F1: (v13 & v13)}
                        _ = v19
                    case "a":
                        fmt.Println(v13)
                    }
                    _ = fi18
                }
            } else {
                v20 := GenS1{F0: v13}
                switch ("hello" + "hello") {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = v20
            }
        }
    }
    if ((!false) || true) {
        v21 := GenS1{F0: (v13 + -1)}
        if (v5 == (v5 + (v5 - ((float64((-2.25 + ((-2.25 / -2.25) - float32(((v5 + v5) + (((0.0 * v5) * v5) * v5)))))) + (1.5 - v5)) * v5)))) {
            var v22 int32 = -2147483648
            var v23 uint = 1
            delete(v4, "a")
            _ = v22
            _ = v23
        } else {
            var v24 int16 = v13
            v12 += v12
            var v25 int16 = (v24 + v13)
            _ = v24
            _ = v25
        }
        _ = v21
    } else {
        var v26 uint16 = 574
        v4["hello"] = (v11 | 2147483647)
        _ = v26
    }
    var v27 int8 = 45
    if (!false) {
        var v28 int64 = 395
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 float32 = -2.25
            var v31 uint32 = 65535
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    var v32 bool = true
    fmt.Println(v1, v3, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
