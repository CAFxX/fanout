package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint16 = 42
    var v2 uint = 243
    v3 := []string{"hello"}
    var v4 string = ("foo bar" + ("foo bar" + "foo bar"))
    var v5 uint8 = 1
    fmt.Println(v4, v5, v4)
    var v6 int = 9223372036854775807
    v7 := GenS1{F0: (("\n\t" + ((v4 + v4) + v4)) + "\n\t"), F1: 34, F2: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v4 + v4) {
    case "\n\t":
        var v8 string = ("\n\t" + v4)
        _ = v8
    case "hello":
        var v9 int16 = int16(v1)
        _ = v9
    case "a":
        var v10 int64 = -221
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    if len(v3) > 0 {
        v3[0] = v4
    }
    var v12 string = v4
    v6 -= (v6 ^ v6)
    v13 := GenS2{F0: 2}
    switch (-(v6)) {
    case -368:
        fmt.Println(v12)
    case -9223372036854775808:
        v14 := GenS2{F0: 18446744073709551615}
        _ = v14
    case 4294967295:
        v4 = (v4 + v4)
    }
    v3 = append(v3, (v12 + ""))
    var v15 uint8 = 102
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int64 = 42
        fmt.Println(v15, v2, v15)
        if len(v3) > 0 {
            v3[0] = v12
        }
        _ = fi16
        _ = v17
    }
    if (-706 < 1) {
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(fi18, v6)
            var v19 bool = true
            _ = fi18
            _ = v19
        }
        if true {
            _ = len(v3)
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 bool = (true && true)
                var v22 int32 = 218
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
