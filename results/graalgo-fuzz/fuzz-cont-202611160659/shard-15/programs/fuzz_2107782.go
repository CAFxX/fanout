package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }

func main() {
    var v1 bool = (-9223372036854775808 > -1)
    v2 := []string{(("\n\t" + "foo bar") + "a"), ("foo bar" + (("hello" + "a") + ("foo bar" + "")))}
    var v3 bool = (103 < 2147483648)
    var v4 int32 = 1
    var v5 int16 = 42
    var v6 float64 = (0.0 + 0.0)
    if true {
        fmt.Println(v3)
        fmt.Println(v1)
    }
    if true {
        v3 = (!(!(1 == 127)))
        fmt.Println(v3, v3, v4)
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: (v4 * 1)}
        v9 := GenS1{F0: (v4 & (v4 & 529))}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        if (((1.5 - (1.5 / 1.5)) >= 1e10) || (v3 && v1)) {
            if (73 != 42) {
                fmt.Println(fi10, v5, v1)
            } else {
                v11 := GenS2{F0: 0}
                _ = len(v2)
                _ = v11
            }
            v12 := GenS1{F0: v4}
            var v13 uint = 852
            _ = v12
            _ = v13
        }
        var v14 int32 = 208
        switch "foo bar" {
        case "":
            _ = len(v2)
        case "hello":
            v15 := GenS1{F0: (v14 / (((v4 / ((2147483647) | 1))) | 1))}
            _ = v15
        case "a":
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
        _ = fi10
        _ = v14
    }
    if len(v2) > 0 {
        v2[0] = ((((("\n\t" + ("" + "foo bar")) + "foo bar") + "hello") + ("hello" + ("a" + ("a" + "a")))) + "\n\t")
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS2{F0: 1}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v3 = v3
        fmt.Println(v4)
        v2 = append(v2, ("a" + "a"))
        _ = fi19
    }
    v20 := GenS3{F0: 128, F1: 991, F2: -2147483648}
    fmt.Println(v4)
    switch "a" {
    case "hello":
        v3 = (39 > 22)
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
