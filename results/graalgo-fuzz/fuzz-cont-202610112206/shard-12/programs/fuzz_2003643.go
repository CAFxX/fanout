package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int32 = 2147483647
    var v2 string = "hello"
    var v3 int64 = 42
    v4 := []string{"\n\t", v2}
    var v5 string = v2
    v2 = v5
    v5 = v5
    var v6 uint32 = 0
    switch (((v2 + ("hello" + (v2 + ((v5 + v5) + ((v5 + "") + v5))))) + v2) + "a") {
    case "":
        var v7 uint16 = 42
        _ = v7
    case "\n\t":
        switch -1 {
        case 4294967295:
            v8 := GenS1{F0: (false || (-162 <= 336)), F1: 9223372036854775807}
            _ = v8
        }
    case "foo bar":
        fmt.Println(v2, v5, v3)
    default:
        fmt.Println(v5, v2, v1)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    _ = len(v4)
    for fi9 := 0; fi9 < 13; fi9++ {
        v4 = append(v4, "hello")
        var v10 float64 = (1e10 + ((((-2.25 * 1.5) * float64((1e10 + (0.1 * 3.14159)))) - 3.14159) + 1e10))
        for fi11 := 0; fi11 < 5; fi11++ {
            v1 += int32(v6)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if (false && ((32767 < -32768) && (false && (false || (65535 != 65535))))) {
        var v12 int8 = -33
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint = 568
    fmt.Println(v3, v13)
    v1 = -2147483648
    fmt.Println(v1, v2)
    var v14 int64 = v3
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: false, F1: (fi15 * -2147483648)}
            var v17 float32 = float32(((1e10 / float64(0.1)) / 1.5))
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 bool = true
        if ((18446744073709551615 <= 42) || (!v18)) {
            var v19 uint64 = 1
            _ = v19
        } else {
            if len(v4) > 0 {
                v4[0] = ("" + ("foo bar" + ((v2 + "a") + "foo bar")))
            }
        }
        _ = v18
    } else {
        v5 = (v2 + (v2 + v2))
        v4 = append(v4, "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
