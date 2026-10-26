package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint16 = 1
    v5 := make(map[int]uint32)
    var v6 float64 = 0.1
    switch ((("hello" + ("a" + (((("\n\t" + "") + ((("foo bar" + "") + ("foo bar" + "\n\t")) + "")) + ("hello" + ("a" + ""))) + "\n\t"))) + "hello") + ((("" + ("foo bar" + "")) + "\n\t") + "\n\t")) {
    case "":
        v4 += 42
    default:
        fmt.Println(v4, v6)
    }
    if ("foo bar" <= "hello") {
        v7 := v2["foo bar"]
        var v8 float64 = v6
        var v9 int = -48
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{9223372036854775807}
    v10 = append(v10, 537)
    var v11 int64 = 42
    switch 1 {
    case 295:
        if true {
            delete(v2, "a")
            fmt.Println(v11, v4, v11)
            if true {
                v6 -= ((0.0 - ((v6 * 0.0) - (((float64(3.14159) * ((v6 - (0.0 / 0.1)) / 1e10)) * v6) - v6))) + 1.5)
                var v12 uint8 = 206
                _ = v12
            }
        }
    }
    var v13 bool = (161 >= 255)
    v14 := GenS1{F0: (v11 >> ((4294967295) & 63)), F1: int16(v4), F2: 42}
    v15 := GenS1{F0: 1, F1: 1000, F2: 34}
    v4 = v4
    v4 -= (v4 << ((0) & 63))
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 6; fi17++ {
        v16 = v16
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
