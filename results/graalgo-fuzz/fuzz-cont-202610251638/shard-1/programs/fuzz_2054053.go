package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = 1e10
    var v2 float32 = (float32(0.1) / (3.14159 / (-2.25 * float32((float64((float32(v1) / (float32((1e10 - (v1 + (v1 - (v1 + 1.5))))) / -2.25))) - (v1 / (3.14159 / v1)))))))
    v3 := []uint32{}
    var v4 bool = true
    var v5 uint8 = 255
    var v6 uint64 = 42
    var v7 bool = v4
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "\n\t" {
        case "foo bar":
            var v10 int16 = 1
            _ = v10
        }
        _ = fi9
    }
    switch ("foo bar" + ("\n\t" + (("" + "foo bar") + ""))) {
    case "\n\t":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int32 = 1
            if (0 <= -952) {
                fmt.Println(v1, v5)
                v13 := GenS1{F0: 42}
                _ = v13
            }
            if (v7 || v7) {
                v5 += (^(v5))
                if false {
                    v6 -= v6
                    _ = len(v3)
                }
            }
            _ = fi11
            _ = v12
        }
    case "":
        switch "hello" {
        case "hello":
            var v14 uint16 = 2
            _ = v14
        case "foo bar":
            v15 := GenS1{F0: 17}
            _ = v15
        case "a":
            switch -882 {
            case -2147483648:
                v1 -= v1
            }
        }
    case "hello":
        for fi16 := 0; fi16 < 20; fi16++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi16
        }
    default:
        var v17 float64 = (v1 - 0.0)
        _ = v17
    }
    v4 = (v5 != 127)
    switch ("a" + ("hello" + ("\n\t" + (("\n\t" + (("\n\t" + "\n\t") + "\n\t")) + "hello")))) {
    case "foo bar":
        fmt.Println(v1)
    case "a":
        fmt.Println(v6, v2, v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 float64 = (float64(v2) / 0.1)
    v5 = v5
    if true {
        switch "a" {
        case "hello":
            v19 := GenS1{F0: 127}
            _ = v19
        case "\n\t":
            v4 = ((!v7) || v4)
        case "a":
            var v20 uint8 = v5
            _ = v20
        }
        var v21 uint64 = v6
        _ = v21
    } else {
        var v22 bool = ((v2 / v2) == (0.0 - 1.5))
        _ = v22
    }
    v4 = (-908 != -557)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
