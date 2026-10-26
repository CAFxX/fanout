package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint8 = 128
    var v2 int = -1
    v3 := make(map[string]bool)
    var v4 int8 = 1
    v3[""] = true
    var v5 int = v2
    var v6 int16 = 1000
    fmt.Println(v2)
    v7 := []int{v2, v5}
    v8 := GenS1{F0: 2147483647}
    if ((0.1 - (((((0.0 + 1e10) * 0.1) / (0.0 + (1.5 / (0.0 + -2.25)))) / 1e10) * 3.14159)) <= 0.0) {
        switch (v2 << ((65535) & 63)) {
        case -1:
            if (!(true && (65535 <= 848))) {
                switch v2 {
                case 9223372036854775807:
                    var v9 int64 = 42
                    _ = v9
                case -663:
                    if false {
                        var v10 int64 = 9223372036854775807
                        if (true || (!(false && false))) {
                            _ = len(v3)
                            if len(v7) > 0 {
                                v7[0] = 32767
                            }
                            var v11 uint8 = (v1 * 2)
                            _ = v11
                        } else {
                            for fi12 := 0; fi12 < 13; fi12++ {
                                fmt.Println(v2, v10)
                                v13 := GenS1{F0: 220}
                                v14 := GenS1{F0: 65535}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        }
                        fmt.Println(v1, v2)
                        _ = v10
                    }
                case 32767:
                    var v15 int = (v2 % ((v5) | 1))
                    _ = v15
                }
                var v16 int8 = (v4 & v4)
                var v17 uint16 = 32767
                _ = v16
                _ = v17
            }
        }
        _ = len(v7)
    } else {
        v4 = -128
        v18 := GenS1{F0: 42}
        _ = v18
    }
    v3[((("\n\t" + "\n\t") + "\n\t") + "\n\t")] = (!((true && (255 < 256)) && false))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
