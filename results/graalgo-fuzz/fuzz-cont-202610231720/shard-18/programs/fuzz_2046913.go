package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 uint = 809
    v2 := make(map[int]string)
    var v3 int64 = 98
    v4 := []uint32{65535}
    fmt.Println(v1, v3)
    switch "hello" {
    case "hello":
        var v5 int64 = v3
        _ = v5
    case "\n\t":
        var v6 uint16 = 2
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi7
    }
    switch 4294967295 {
    case -948:
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi8)
                v10 := GenS1{F0: ("hello" + ("hello" + "foo bar")), F1: 42}
                if ("a" > "a") {
                    v11 := GenS1{F0: "\n\t", F1: -32768}
                    if (true && true) {
                        if false {
                            if (!true) {
                                fmt.Println(v3)
                                v4 = append(v4, 0)
                            }
                        }
                        fmt.Println(v3, v1)
                        var v12 float32 = (0.0 + -2.25)
                        _ = v12
                    }
                    _ = v11
                } else {
                    var v13 int = (fi8 >> ((v1) & 63))
                    var v14 int8 = -1
                    if (!false) {
                        v1 = (v1 >> ((v1) & 63))
                    } else {
                        var v15 float64 = (3.14159 + 1e10)
                        var v16 bool = (4294967295 != 2)
                        v4 = append(v4, 65535)
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    case 4294967295:
        switch "hello" {
        case "\n\t":
            var v17 uint32 = 130
            _ = v17
        case "foo bar":
            var v18 float64 = ((0.1 * 3.14159) / ((1e10 * (float64((3.14159 / 1.0)) * (0.1 + (-2.25 - float64(1e10))))) * (-2.25 * (0.1 - 0.0))))
            _ = v18
        case "a":
            v3 -= -1
        }
    }
    switch 193 {
    case 127:
        v4 = append(v4, 581)
    case 42:
        if ((3.14159 < 3.14159) && true) {
            v4 = append(v4, 600)
            _ = len(v2)
        } else {
            v19 := GenS1{F0: "foo bar", F1: 32767}
            _ = v19
        }
    }
    fmt.Println(v3, v1)
    fmt.Println(v1, v1)
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS1{F0: "foo bar", F1: -1}
        _ = len(v2)
        _ = fi20
        _ = v21
    }
    switch ("" + ("foo bar" + ("" + "\n\t"))) {
    case "\n\t":
        var v22 int8 = -128
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            if true {
                v2[32768] = "hello"
            }
            fmt.Println(v3, fi24)
            _ = fi24
        }
        _ = len(v2)
        _ = fi23
    }
    fmt.Println(v3, v1)
    var v25 int8 = 16
    var v26 float64 = -2.25
    var v27 int16 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v25 += 42
    var v28 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
