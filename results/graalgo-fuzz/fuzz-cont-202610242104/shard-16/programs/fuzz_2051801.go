package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    var v1 bool = ((65535 < 2) || true)
    var v2 int8 = 0
    var v3 uint32 = 1
    var v4 uint = 2
    var v5 string = "foo bar"
    var v6 int32 = -935
    v7 := GenS1{F0: "foo bar"}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float64 = 0.0
        _ = fi8
        _ = v9
    }
    switch (v5 + v5) {
    case "\n\t":
        switch "\n\t" {
        case "a":
            if (!v1) {
                v10 := []int{-1, 128, 4294967295, 1}
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS2{F0: 565, F1: 128, F2: (32767 != (fi11 >> ((v4) & 63)))}
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v10) > 0 {
                        v10[0] = 1
                    }
                    _ = fi13
                }
                _ = v10
            } else {
                v14 := GenS3{F0: 61, F1: 255}
                v15 := GenS2{F0: 313, F1: 72, F2: (v1 && (v1 || v1))}
                switch -2147483648 {
                case -815:
                    v5 = v5
                }
                _ = v14
                _ = v15
            }
        }
    case "a":
        if (2 > 0) {
            fmt.Println(v5, v2)
        } else {
            v16 := []int{2147483647, -855, 42}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            var v17 string = (("hello" + v5) + v5)
            _ = v16
            _ = v17
        }
    }
    v5 = v5
    var v18 uint32 = v3
    var v19 uint16 = 42
    fmt.Println(v1)
    v6 = (v6 / ((2147483647) | 1))
    v20 := GenS4{F0: (v6 << ((v4) & 63)), F1: 127}
    switch 647 {
    case 957:
        var v21 int16 = -406
        _ = v21
    }
    v22 := GenS1{F0: ((v5 + v5) + v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
