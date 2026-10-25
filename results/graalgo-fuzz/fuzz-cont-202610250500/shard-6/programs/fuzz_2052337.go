package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }

func main() {
    var v1 int64 = -706
    var v2 uint8 = 128
    var v3 uint = 65535
    var v4 int32 = 0
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: v4, F1: (!false), F2: "foo bar"}
        v3 = v3
        for fi9 := 0; fi9 < 6; fi9++ {
            v4 -= -842
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if false {
        if false {
            fmt.Println(v2, v1, v4)
            switch 42 {
            case 2147483647:
                v2 = v2
            }
        }
        v5["\n\t"] = 65535
        v10 := GenS1{F0: v4, F1: false, F2: ("hello" + "\n\t")}
        _ = v10
    }
    _ = len(v5)
    fmt.Println(v4)
    var v11 uint64 = 0
    v12 := v5[""]
    v4 = -2147483648
    switch 42 {
    case 1:
        _ = len(v6)
    default:
        for fi13 := 0; fi13 < 6; fi13++ {
            v2 += v2
            var v14 int8 = 0
            for fi15 := 0; fi15 < 18; fi15++ {
                fmt.Println(v3)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    }
    switch "" {
    case "\n\t":
        if false {
            var v16 string = ""
            fmt.Println(v2)
            _ = v16
        } else {
            v17 := GenS1{F0: v4, F1: ("a" > (("a" + "foo bar") + ("\n\t" + ("foo bar" + (("foo bar" + "") + "hello"))))), F2: "a"}
            delete(v5, (("" + ("hello" + "\n\t")) + ("a" + "foo bar")))
            _ = v17
        }
    case "hello":
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint8 = (^(v2))
            _ = fi18
            _ = v19
        }
    }
    var v20 float32 = (0.0 / 1.0)
    v21 := GenS1{F0: (v4 << (((v3 & uint(v11))) & 63)), F1: (((-2.25 <= (1.5 / float64(v20))) && ((v2 + 1) > 127)) || false), F2: "foo bar"}
    v3 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
