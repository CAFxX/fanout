package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint64, p1 int32) bool {
    switch "" {
    case "hello":
        var v6 uint = 2
        _ = v6
    }
    var v7 uint8 = 1
    _ = p0
    _ = p1
    _ = v7
    return false
}
func genFn1(p0 uint32) uint64 {
    switch "" {
    case "\n\t":
        switch 128 {
        case 4294967295:
            var v8 uint = 657
            _ = v8
        case 84:
            v9 := []int{792, -1, 32768, 796}
            _ = v9
        }
    case "a":
        var v10 int32 = -46
        _ = v10
    }
    var v11 string = ((("foo bar" + "\n\t") + "a") + ("\n\t" + "\n\t"))
    _ = p0
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int32 = 0
    v4 := make(map[uint32]string)
    var v5 int = 32768
    v12 := v4[2147483647]
    fmt.Println(v12, v12)
    var v13 int16 = -32768
    var v14 float32 = 1.5
    var v15 float64 = 1.5
    v5 += v5
    if true {
        if (165 <= 0) {
            for fi16 := 0; fi16 < 7; fi16++ {
                for fi17 := 0; fi17 < 5; fi17++ {
                    fmt.Println(fi17, v13, v5)
                    _ = fi17
                }
                _ = fi16
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                if false {
                    var v19 uint16 = 42
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(v14, v5)
                        _ = fi20
                    }
                    _ = v19
                }
                _ = fi18
            }
        }
    }
    if (!true) {
        switch v5 {
        case 32768:
            var v21 int = (v5 % (((v5 + (v5 | v5))) | 1))
            _ = v21
        case 127:
            v22 := GenS1{F0: 128, F1: 0}
            _ = v22
        case 9223372036854775807:
            v23 := GenS1{F0: 2, F1: v3}
            _ = v23
        }
    } else {
        v24 := GenS1{F0: 792, F1: -338}
        fmt.Println(v12, v5)
        _ = v24
    }
    v25 := GenS1{F0: 0, F1: 313}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
