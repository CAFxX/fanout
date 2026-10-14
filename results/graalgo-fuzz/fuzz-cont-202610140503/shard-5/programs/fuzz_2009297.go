package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    var v5 bool = (((true && false) || true) && (42 <= 255))
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{"hello", "\n\t", (("foo bar" + (("\n\t" + ("hello" + ("a" + ("hello" + "a")))) + "a")) + "hello"), "\n\t"}
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    v3 += v3
    var v7 int = genFn0(0)
    v4["hello"] = ""
    var v8 int = (v7 / ((v7) | 1))
    var v9 float64 = (v1 - float64(0.1))
    v10 := v4["foo bar"]
    _ = len(v4)
    var v11 uint = 0
    v12 := GenS1{F0: -87, F1: (!true), F2: 1}
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint16 = 65535
        switch genFn0(v3) {
        case 127:
            switch (("\n\t" + v10) + (("" + (("\n\t" + v10) + (v10 + ""))) + v10)) {
            case "a":
                fmt.Println(v7, v10)
            case "\n\t":
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS1{F0: -9, F1: true, F2: -32768}
                    var v17 string = v10
                    v18 := GenS1{F0: 127, F1: true, F2: 214}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            default:
                switch (((("" + "\n\t") + ("a" + v10)) + (v10 + v10)) + "") {
                case "foo bar":
                    v1 -= (-2.25 / float64((3.14159 - 1e10)))
                case "a":
                    v14 -= (^(v14))
                }
            }
        case 952:
            for fi19 := 0; fi19 < 13; fi19++ {
                v3 -= v3
                _ = fi19
            }
        case 4294967295:
            delete(v4, "")
        }
        _ = fi13
        _ = v14
    }
    var v20 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
