package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0(p0 bool, p1 uint64, p2 uint8) int16 {
    var v4 int = -635
    v5 := []int{v4}
    if ((!(!true)) && p0) {
        p2 = p2
        switch "\n\t" {
        case "\n\t":
            var v6 float32 = 0.1
            _ = v6
        }
    }
    v7 := GenS1{F0: p2, F1: 2, F2: 32767}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    return 581
}

func main() {
    var v1 int = 32768
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v8 := GenS1{F0: 128, F1: 65535, F2: genFn0((!true), 447, 127)}
    var v9 bool = (true && true)
    v1 = v1
    v1 = 42
    var v10 int16 = 42
    for fi11 := 0; fi11 < 9; fi11++ {
        switch fi11 {
        case 1:
            var v12 uint16 = 65535
            _ = v12
        case 9223372036854775807:
            switch fi11 {
            case 4294967295:
                var v13 int = 9223372036854775807
                _ = v13
            case -146:
                fmt.Println(v1)
            case -9223372036854775808:
                if ((709 >= -886) && ((!(-1 >= 1)) && false)) {
                    v9 = v9
                    v9 = v9
                } else {
                    if v9 {
                        v9 = (457 > 386)
                        v14 := GenS1{F0: 2, F1: 65535, F2: 32767}
                        _ = v14
                    }
                    var v15 float32 = 1e10
                    _ = v15
                }
            default:
                v9 = (float32(1.5) >= 1.5)
            }
        case 128:
            fmt.Println(v10)
        default:
            _ = len(v2)
        }
        _ = fi11
    }
    if ((128 != 194) || (-128 <= 100)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint32 = 311
            var v18 int8 = 1
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int8 = 0
            var v21 bool = v9
            _ = fi19
            _ = v20
            _ = v21
        }
        delete(v3, "")
        fmt.Println(v1)
    }
    var v22 string = (("a" + "hello") + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
