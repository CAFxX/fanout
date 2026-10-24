package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -128
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := []string{("hello" + ((("hello" + "\n\t") + ("hello" + "a")) + "foo bar")), ("\n\t" + "")}
    v5 := []string{(("hello" + "\n\t") + "")}
    if false {
        v2 += (0.1 + (v2 * 3.14159))
    } else {
        if (true && true) {
            fmt.Println(v1, v2)
            fmt.Println(v2)
            if len(v4) > 0 {
                v4[0] = ""
            }
        }
    }
    var v6 int32 = 42
    v6 -= (v6 | v6)
    var v7 bool = false
    v1 = (v1 ^ (^(v1)))
    v1 = -1
    v4 = append(v4, "")
    var v8 int64 = -9223372036854775808
    if false {
        if (!(3.14159 <= ((1e10 - 1.5) / (0.1 * -2.25)))) {
            fmt.Println(v1)
        }
    } else {
        switch (("foo bar" + ("\n\t" + ((("" + "hello") + ("hello" + ("hello" + ""))) + "\n\t"))) + ("" + "hello")) {
        case "":
            var v9 uint16 = 42
            _ = v9
        case "hello":
            v8 = (v8 & v8)
        case "foo bar":
            var v10 uint = 0
            _ = v10
        default:
            for fi11 := 0; fi11 < 3; fi11++ {
                v8 = -233
                _ = fi11
            }
        }
        var v12 int = -2147483648
        v6 = (v6 & v6)
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
