package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int16 = -32768
    v2 := []uint32{1, 65535, 850, 42}
    var v3 uint8 = 139
    var v4 uint = 18446744073709551615
    var v5 uint64 = 535
    var v6 float32 = 0.0
    fmt.Println(v3, v5, v5)
    if (65535 > 0) {
        switch "a" {
        case "foo bar":
            if true {
                if true {
                    v7 := GenS1{F0: true, F1: 2}
                    _ = v7
                }
            }
        case "hello":
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint32 = 65535
                v6 = (0.1 * 1e10)
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(fi8)
                    switch (fi8 - (fi8 * (fi10 - 1))) {
                    case -9223372036854775808:
                        if (int(v1) != (fi8 / ((17) | 1))) {
                            _ = len(v2)
                        } else {
                            v5 -= 250
                            v11 := GenS1{F0: ((v1 * int16(v4)) >= -1), F1: 0}
                            if ((!true) && true) {
                                fmt.Println(v9, fi8, fi8)
                                v3 += (v3 * uint8(v5))
                            }
                            _ = v11
                        }
                    case 0:
                        if ("\n\t" > "\n\t") {
                            var v12 int32 = 0
                            _ = v12
                        }
                    }
                    switch ("\n\t" + "hello") {
                    case "a":
                        _ = len(v2)
                    default:
                        v13 := GenS1{F0: false, F1: 255}
                        _ = v13
                    }
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        default:
            v4 -= 2
        }
    }
    v5 = (v5 - 4294967295)
    var v14 uint64 = v5
    v2 = append(v2, 42)
    fmt.Println(v1, v3)
    v14 -= (v5 * 113)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
