package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 int8 = 1
    var v2 bool = (((42 < 963) || false) && (!true))
    var v3 uint16 = 2
    var v4 uint32 = 346
    v2 = (true || v2)
    v5 := []int{620, 4, -9223372036854775808, -755}
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 13; fi7++ {
        if (v2 || v2) {
            var v8 int32 = -1
            if false {
                _ = len(v5)
                fmt.Println(v2, v3, v8)
            } else {
                switch ((((v6 + v6) + v6) + (v6 + v6)) + ((v6 + (v6 + "")) + v6)) {
                case "":
                    if (18446744073709551615 <= 42) {
                        fmt.Println(v1)
                    }
                }
            }
            if len(v5) > 0 {
                v5[0] = (fi7 - int(v4))
            }
            _ = v8
        } else {
            v9 := GenS1{F0: 359, F1: 42}
            _ = v9
        }
        _ = fi7
    }
    v4 = 2147483648
    v10 := GenS1{F0: 594, F1: (v1 % ((127) | 1))}
    var v11 uint = 1
    fmt.Println(v1, v11, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
