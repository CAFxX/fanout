package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    var v3 bool = (!(0 != -2147483648))
    v4 := make(map[string]string)
    v5 := []string{}
    switch -627 {
    case -583:
        switch "" {
        case "a":
            if true {
                if v3 {
                    var v6 uint64 = 42
                    v3 = (42 < 1)
                    v1 += (v1 / ((v1) | 1))
                    _ = v6
                } else {
                    var v7 float64 = (((1.5 / ((((3.14159 / 1.0) + 1e10) - -2.25) * 1e10)) - float64(1e10)) / 1.0)
                    _ = v7
                }
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 int16 = -1
                    _ = fi8
                    _ = v9
                }
            }
        }
    case 853:
        if len(v5) > 0 {
            v5[0] = ""
        }
    }
    _ = len(v5)
    var v10 int16 = 1
    var v11 uint8 = 128
    switch 4294967295 {
    case 127:
        if v3 {
            v10 = (v10 | (v10 % (((v10 / ((v10) | 1))) | 1)))
            var v12 float32 = 0.0
            _ = v12
        }
    case -1:
        if (669 > 1) {
            v13 := GenS1{F0: ((("hello" + ("hello" + "a")) + "") + ""), F1: -9223372036854775808}
            _ = v13
        }
    case -220:
        v10 -= (v10 + (v10 / (((v10 | v10)) | 1)))
    }
    v11 = v11
    v14 := GenS1{F0: ("\n\t" + "foo bar"), F1: 127}
    v11 = v11
    var v15 uint32 = 69
    if true {
        delete(v4, "")
        var v16 float32 = -2.25
        _ = v16
    }
    var v17 int = 9223372036854775807
    if (-1 < 9223372036854775807) {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, v3, v1)
            _ = len(v2)
            _ = fi18
        }
        v15 -= v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
