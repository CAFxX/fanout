package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8 }
func genFn0() bool {
    v5 := GenS1{F0: ("\n\t" + "a"), F1: -798}
    var v6 int64 = 42
    if true {
        if true {
            fmt.Println(v6)
            var v7 string = "a"
            _ = v7
        }
        var v8 uint32 = 42
        _ = v8
    }
    v9 := GenS1{F0: "hello", F1: (v6 >> ((42) & 63))}
    if (42 >= 607) {
        v6 = 42
        var v10 int8 = -128
        var v11 int32 = 0
        _ = v10
        _ = v11
    } else {
        fmt.Println(v6)
    }
    _ = v5
    _ = v6
    _ = v9
    return false
}
func genFn1() int64 {
    var v12 string = "hello"
    if (1.5 < (float64(((1e10 - (-2.25 + 0.0)) - 1.5)) - (3.14159 - (1e10 - (float64((1e10 / 1.5)) / 1.0))))) {
        var v13 float32 = (((3.14159 / (0.0 - 1.5)) * 1e10) + 1e10)
        var v14 float64 = (3.14159 - 1.5)
        _ = v13
        _ = v14
    }
    _ = v12
    return -366
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 3.14159
    v3 := []string{""}
    var v4 uint8 = 128
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint32 = 199
        switch (("\n\t" + "") + "") {
        case "\n\t":
            v17 := GenS2{F0: -115}
            _ = v17
        case "a":
            var v18 int8 = 0
            _ = v18
        }
        if (!false) {
            if len(v3) > 0 {
                v3[0] = ("hello" + "a")
            }
            switch (fi15 * (fi15 << ((4294967295) & 63))) {
            case 127:
                fmt.Println(v2, fi15, v16)
            }
        }
        _ = fi15
        _ = v16
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: ("foo bar" + ("hello" + "foo bar")), F1: genFn1()}
        _ = fi19
        _ = v20
    }
    v4 += (v4 * 42)
    var v21 int32 = 1
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v21, v4)
        v23 := GenS1{F0: ("\n\t" + ("foo bar" + ("a" + "a"))), F1: -9223372036854775808}
        _ = fi22
        _ = v23
    }
    var v24 float64 = 3.14159
    v2 += (0.1 * (0.0 * v24))
    v2 += 1e10
    v25 := v1["a"]
    if (0 < 516) {
        v26 := GenS1{F0: "\n\t", F1: -352}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
