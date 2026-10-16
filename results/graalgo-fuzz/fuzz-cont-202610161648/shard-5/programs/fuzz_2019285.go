package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (0.1 + 0.0)
    var v3 uint16 = 0
    var v4 uint16 = 32767
    if (-79 != (v1 * -2147483648)) {
        fmt.Println(v2, v1)
    }
    var v5 bool = true
    v5 = (v5 || ((799 > 2147483647) && v5))
    var v6 float32 = v2
    v7 := GenS1{F0: -1, F1: 227}
    switch 2147483647 {
    case 32767:
        v8 := GenS2{F0: v4, F1: 0, F2: 28}
        _ = v8
    case 2147483647:
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch ("a" + (("" + "\n\t") + "foo bar")) {
                case "\n\t":
                    v11 := GenS1{F0: (v1 / ((1) | 1)), F1: (fi9 - (^(fi10)))}
                    _ = v11
                default:
                    if (int(v4) == fi10) {
                        v3 -= v4
                        v1 = (v1 ^ -954)
                        v6 = v2
                    }
                }
                _ = fi10
            }
            switch (("" + (("\n\t" + "a") + ("hello" + "a"))) + "\n\t") {
            case "":
                v12 := GenS2{F0: v4, F1: -128, F2: 42}
                _ = v12
            case "hello":
                fmt.Println(v5)
            case "\n\t":
                var v13 uint64 = 580
                _ = v13
            }
            _ = fi9
        }
    case -415:
        v5 = v5
    }
    var v14 int8 = 0
    var v15 float32 = (v2 / 0.1)
    var v16 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
