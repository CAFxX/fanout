package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = (1.5 * (1e10 + float64(0.1)))
    var v2 float64 = ((v1 - 0.0) - -2.25)
    v3 := []string{("foo bar" + ("hello" + "hello"))}
    var v4 int8 = -100
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint64 = 4294967295
        _ = len(v3)
        _ = fi5
        _ = v6
    }
    v4 += -125
    switch 795 {
    case 129:
        var v7 bool = (!(false || (2 != 210)))
        _ = v7
    case 32768:
        v8 := GenS1{F0: 0, F1: 42, F2: 0}
        _ = v8
    default:
        v4 = -45
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        switch ("a" + "\n\t") {
        case "\n\t":
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(v4, v2)
                _ = fi10
            }
        case "":
            v4 = v4
        case "hello":
            fmt.Println(v4, v1)
        }
        v4 = v4
        v11 := GenS1{F0: 4294967295, F1: v4, F2: 9223372036854775807}
        _ = fi9
        _ = v11
    }
    v3 = append(v3, "\n\t")
    v2 = (v2 * (((1e10 - 3.14159) + v2) - v1))
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v4 -= v4
            v3 = append(v3, ("\n\t" + ("hello" + ((("\n\t" + "\n\t") + (((("\n\t" + "") + (("hello" + "foo bar") + (("hello" + "\n\t") + ("\n\t" + ("" + "a"))))) + "\n\t") + "\n\t")) + "a"))))
            _ = fi13
        }
        var v14 bool = false
        _ = fi12
        _ = v14
    }
    switch "" {
    case "\n\t":
        v2 += v1
    case "foo bar":
        v15 := GenS1{F0: 1, F1: -16, F2: -616}
        _ = v15
    }
    v16 := GenS2{F0: true}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v2, v4)
    if true {
        _ = len(v3)
        for fi17 := 0; fi17 < 4; fi17++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            v18 := GenS1{F0: 4294967295, F1: v4, F2: 9223372036854775807}
            _ = fi17
            _ = v18
        }
    } else {
        var v19 int8 = v4
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
