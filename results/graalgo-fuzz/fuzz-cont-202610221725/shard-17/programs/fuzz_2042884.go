package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
func genFn0(p0 string, p1 uint8) bool {
    v4 := GenS1{F0: 0}
    var v5 int32 = -2147483648
    if (true && false) {
        v6 := GenS1{F0: 380}
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v9
    return false
}
func genFn1(p0 int64) uint8 {
    switch 32768 {
    case 129:
        v10 := []int{32767, 32768}
        _ = v10
    case -1:
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            fmt.Println(p0)
        case "":
            v11 := []int{416, -2147483648, -1, int(p0)}
            _ = v11
        default:
            v12 := GenS2{F0: 255, F1: 131, F2: 384}
            _ = v12
        }
    case -532:
        switch 32768 {
        case 128:
            v13 := []int{}
            _ = v13
        case 1:
            v14 := GenS2{F0: 0, F1: 2, F2: 42}
            _ = v14
        case -444:
            if false {
                p0 = (p0 / (((p0 ^ (p0 + 899))) | 1))
            } else {
                p0 = p0
            }
        default:
            if false {
                p0 = p0
            }
        }
    }
    v15 := GenS1{F0: 603}
    var v16 int16 = -1
    _ = p0
    _ = v15
    _ = v16
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    v3 := make(map[int]string)
    _ = len(v1)
    fmt.Println(v2)
    var v17 uint = 65535
    fmt.Println(v2, v2)
    v17 = v17
    var v18 int32 = 42
    var v19 uint16 = 713
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
