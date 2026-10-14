package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", ("\n\t" + ""), "foo bar"}
    var v3 uint32 = 42
    var v4 float64 = 0.1
    var v5 int32 = -1
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(fi7, fi7)
            _ = len(v2)
            switch 42 {
            case 32767:
                fmt.Println(v3, v4)
            case 42:
                fmt.Println(v3, fi7)
            case -1:
                delete(v1, "hello")
            }
            _ = fi8
        }
        _ = fi7
    }
    v5 = 470
    if (!true) {
        v9 := GenS2{F0: 75}
        _ = v9
    } else {
        v4 = (float64(0.1) / (3.14159 / (float64(((((((1e10 * 0.1) + (1e10 - float32(float64(1e10)))) * 1e10) / ((3.14159 / 1.5) / 1.0)) * (float32(v4) + -2.25)) / 1.0)) - 1.5)))
    }
    v3 += (v3 | 263)
    fmt.Println(v3, v4)
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v4, v3, fi10)
        _ = fi10
    }
    switch -1 {
    case -623:
        switch -1 {
        case 1:
            _ = len(v2)
        case 0:
            v4 = v4
        case 864:
            _ = len(v1)
        }
    case 4294967295:
        fmt.Println(v3, v5, v5)
    case 856:
        if true {
            var v11 uint64 = 18446744073709551615
            _ = v11
        } else {
            if (((!(!(true && (!(2 != 0))))) && (false && false)) && false) {
                var v12 int32 = -2147483648
                fmt.Println(v4, v5, v3)
                var v13 int8 = 0
                _ = v12
                _ = v13
            }
        }
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v3 += 2147483647
    case "":
        _ = len(v1)
    case "hello":
        v14 := GenS1{F0: 32767}
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v4 = 0.0
        fmt.Println(v4, fi15)
        _ = fi15
    }
    var v16 bool = true
    v4 += 1.5
    var v17 int8 = 99
    var v18 int64 = 0
    v5 += 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v18, v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
