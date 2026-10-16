package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    var v1 int16 = 1
    var v2 int8 = 23
    var v3 float64 = float64(1e10)
    v3 += 3.14159
    switch 32767 {
    case -557:
        var v4 uint32 = 4294967295
        _ = v4
    default:
        var v5 string = "hello"
        _ = v5
    }
    switch 32767 {
    case 0:
        if (0 != 32767) {
            if (v3 > float64((1.5 + float32(((float64((1.5 * (float32(float64(float32(((v3 / v3) - v3)))) - 3.14159))) * v3) * v3))))) {
                fmt.Println(v1, v2, v1)
            }
        }
    }
    var v6 float64 = v3
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch 128 {
            case 32768:
                var v9 float64 = float64(3.14159)
                _ = v9
            case 127:
                v10 := GenS1{F0: "hello", F1: ("hello" + "\n\t"), F2: -1}
                _ = v10
            }
            _ = fi8
        }
        v2 = (v2 + 1)
        _ = fi7
    }
    if (v3 != ((v3 * (v3 + (v6 - float64(-2.25)))) - v6)) {
        var v11 uint64 = 957
        _ = v11
    }
    v12 := GenS1{F0: "a", F1: ((("" + ((("hello" + "foo bar") + "foo bar") + "hello")) + "hello") + "foo bar"), F2: 1}
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: ((("hello" + "a") + ("a" + (("foo bar" + "a") + ((("\n\t" + "") + "\n\t") + ((("hello" + ("\n\t" + ("" + "foo bar"))) + "") + ("hello" + "")))))) + ("foo bar" + ("" + "hello"))), F1: "hello", F2: 42}
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float32 = 1.5
            v17 := GenS2{F0: 2, F1: "a"}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
}
