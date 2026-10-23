package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "a"
    v3 := []string{v2}
    v4 := make(map[string]string)
    var v5 uint32 = 42
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v2)
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    v5 += v5
    switch "\n\t" {
    case "a":
        if (2 != 0) {
            var v9 int8 = -128
            fmt.Println(v5, v2, v2)
            v9 += (v9 + (v9 | (v9 % (((^(v9))) | 1))))
            _ = v9
        }
    }
    if true {
        v6[(v2 + v2)] = v2
    }
    v6[(v2 + "a")] = (("foo bar" + v2) + v2)
    var v10 int = 2147483647
    v5 = 214
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = append(v3, "")
        _ = fi11
    }
    switch "\n\t" {
    case "a":
        if (-2147483648 < 161) {
            v5 = (v5 | 65535)
            v3 = append(v3, v2)
            var v12 bool = ((0.1 / (-2.25 + (0.1 - 1e10))) >= -2.25)
            _ = v12
        }
    case "hello":
        v5 = 1
    }
    var v13 uint8 = 42
    var v14 int64 = 42
    fmt.Println(v14, v5, v14)
    if false {
        switch "" {
        case "":
            switch (v10 - v10) {
            case -1:
                v14 -= v14
            case 0:
                v4[("" + "\n\t")] = "foo bar"
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v15 int16 = -1
    v16 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
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
