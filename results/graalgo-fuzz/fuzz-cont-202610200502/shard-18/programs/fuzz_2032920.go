package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0() uint {
    var v6 bool = (1 >= -887)
    v7 := GenS1{F0: 42, F1: 1, F2: -80}
    _ = v6
    _ = v7
    return 855
}

func main() {
    var v1 uint64 = 2
    var v2 uint64 = 655
    v3 := []uint32{42, 65535, 2147483647, 1}
    var v4 uint8 = 238
    v5 := make(map[uint32]uint32)
    v4 = v4
    if (2 > 2) {
        switch "hello" {
        case "\n\t":
            if true {
                fmt.Println(v4, v2)
                if true {
                    switch 9223372036854775807 {
                    case -1:
                        delete(v5, 1)
                    case 32768:
                        _ = len(v5)
                    case 32767:
                        var v8 bool = false
                        _ = v8
                    }
                    fmt.Println(v4)
                }
            }
        }
    }
    var v9 bool = (!((-9223372036854775808 != -1) && false))
    var v10 string = ""
    v11 := GenS1{F0: (v4 & 42), F1: 255, F2: -1}
    v5[2] = 886
    switch v10 {
    case "hello":
        if true {
            switch -800 {
            case 604:
                v10 = v10
            case 32767:
                _ = len(v3)
            case -351:
                if ((v4 & (v4 << ((uint(v4)) & 63))) <= (v4 >> ((4294967295) & 63))) {
                    _ = len(v3)
                    v5[2147483647] = 1
                    fmt.Println(v2, v4)
                }
            }
            v12 := v5[440]
            _ = v12
        }
    case "foo bar":
        var v13 uint8 = (^(v4))
        _ = v13
    }
    v4 += 127
    var v14 float32 = (3.14159 - float32(-2.25))
    v3 = append(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
