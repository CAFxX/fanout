package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int16 }
func genFn0(p0 uint16, p1 uint16, p2 uint8) int16 {
    v4 := GenS1{F0: 32767, F1: 2, F2: 387}
    if ((430 != 676) || false) {
        p0 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}
func genFn1(p0 uint16, p1 uint32) uint {
    p0 += (p0 & (p0 << ((760) & 63)))
    if true {
        switch ("hello" + "\n\t") {
        case "hello":
            v5 := GenS1{F0: 32767, F1: (^(p1)), F2: 0}
            _ = v5
        case "a":
            var v6 bool = (!false)
            _ = v6
        default:
            v7 := GenS1{F0: 42, F1: (p1 * (p1 - (p1 | (^(p1))))), F2: -1}
            _ = v7
        }
        var v8 uint64 = 938
        _ = v8
    } else {
        if (42 != -206) {
            p0 = p0
        }
        switch -958 {
        case 9223372036854775807:
            var v9 int32 = 2147483647
            _ = v9
        default:
            var v10 uint = 1
            _ = v10
        }
        p0 += p0
    }
    _ = p0
    _ = p1
    return 334
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 9223372036854775807
    var v3 uint16 = 65535
    v11 := v1[(-(v2))]
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: 0, F1: 146, F2: 1000}
        _ = fi12
        _ = v13
    }
    var v14 int64 = 1
    v14 -= v14
    v14 = (v14 | 538)
    var v15 int32 = 0
    v16 := GenS1{F0: genFn0(2, (^(v3)), 31), F1: 4294967295, F2: 1}
    var v17 uint = 1
    var v18 uint64 = 460
    var v19 uint64 = 2
    fmt.Println(v19)
    var v20 uint = genFn1(v3, 4294967295)
    v2 -= 9223372036854775807
    v21 := GenS1{F0: -486, F1: 42, F2: int16(v14)}
    v22 := []int{(^(v2)), (^(v2)), v2}
    v14 += (v14 - v14)
    var v23 int = -761
    v15 = 163
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
