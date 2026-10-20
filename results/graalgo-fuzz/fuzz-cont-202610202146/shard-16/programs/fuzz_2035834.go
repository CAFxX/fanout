package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 838
    var v3 string = ("\n\t" + "a")
    var v4 int = 32767
    fmt.Println(v3, v4)
    v5 := GenS1{F0: (false || true), F1: v4, F2: v3}
    fmt.Println(v2, v4)
    var v6 int32 = -1
    v2 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        v4 = -747
        fmt.Println(fi7, v3, v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            fmt.Println(v3)
            v1[(v3 + v3)] = 42
        }
        v9 := GenS1{F0: true, F1: (v4 & (-(fi8))), F2: "foo bar"}
        _ = fi8
        _ = v9
    }
    fmt.Println(v4, v2)
    fmt.Println(v3)
    fmt.Println(v3, v6)
    v10 := GenS1{F0: (!(v2 != 32767)), F1: v4, F2: ""}
    var v11 uint8 = 127
    if ((int8(v6) >= 42) && (!(false || true))) {
        switch (v4 & (v4 << ((0) & 63))) {
        case 769:
            switch (v3 + (v3 + v3)) {
            case "a":
                var v12 uint8 = v11
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 13; fi13++ {
                    v11 -= 225
                    _ = fi13
                }
            }
        default:
            v4 += 42
        }
        fmt.Println(v11, v2)
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 float32 = 3.14159
            for fi16 := 0; fi16 < 3; fi16++ {
                fmt.Println(v15)
                v3 = v3
                var v17 int16 = 733
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v2, v3, v3)
    }
    var v18 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
