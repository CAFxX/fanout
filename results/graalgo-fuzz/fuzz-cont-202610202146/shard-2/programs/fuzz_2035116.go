package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 uint32, p1 uint32) int16 {
    if (true || (true && (((false && (1 == 32767)) && (0 != 928)) && (false || false)))) {
        p0 = (^(p1))
    }
    fmt.Println(p1, p0)
    switch "foo bar" {
    case "hello":
        for fi5 := 0; fi5 < 14; fi5++ {
            p1 = 1
            fmt.Println(p0)
            _ = fi5
        }
    case "a":
        var v6 string = ("a" + "")
        _ = v6
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint8) uint16 {
    p0 -= p0
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 string = (("foo bar" + "foo bar") + ("hello" + ""))
            v9 = (v9 + v9)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int32 = 1
    _ = p0
    _ = v10
    return 32767
}

func main() {
    var v1 int32 = -328
    v2 := make(map[string]string)
    v3 := []string{"hello"}
    var v4 int16 = -1
    if ((v4 >> ((996) & 63)) <= (v4 << ((150) & 63))) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 -= v4
    } else {
        var v11 int64 = -595
        _ = v11
    }
    _ = len(v2)
    v1 -= v1
    if true {
        fmt.Println(v4, v4)
    }
    if (!("foo bar" != "hello")) {
        v12 := GenS1{F0: 0, F1: 9223372036854775807}
        var v13 int32 = (^(v1))
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := v2[(("hello" + ((("foo bar" + "foo bar") + (("" + "foo bar") + ("\n\t" + ("hello" + "a")))) + "")) + "\n\t")]
            if len(v3) > 0 {
                v3[0] = v15
            }
            if ((!(!false)) || (fi14 == (fi14 + fi14))) {
                fmt.Println(v1, v4)
                var v16 int = fi14
                var v17 uint8 = 1
                _ = v16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v1)
        v18 := GenS1{F0: 119, F1: -725}
        _ = v18
    }
    v4 = -764
    _ = len(v3)
    switch "foo bar" {
    case "a":
        var v19 float32 = (float32(((1.5 + -2.25) - float64(1e10))) * float32(1.5))
        _ = v19
    case "\n\t":
        v20 := GenS1{F0: 42, F1: 42}
        _ = v20
    default:
        v4 -= (v4 | 1000)
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
