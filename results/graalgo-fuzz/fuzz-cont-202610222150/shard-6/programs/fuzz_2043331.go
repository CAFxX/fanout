package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 int16) uint32 {
    if false {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        var v7 string = "hello"
        var v8 uint8 = 255
        _ = v7
        _ = v8
    }
    var v9 uint64 = 423
    for fi10 := 0; fi10 < 4; fi10++ {
        if (!(3.14159 >= -2.25)) {
            fmt.Println(fi10, fi10, v9)
            for fi11 := 0; fi11 < 5; fi11++ {
                if false {
                    v12 := []int{fi11, 128, -772, fi11}
                    _ = len(v12)
                    _ = v12
                }
                _ = fi11
            }
            fmt.Println(p0, p0, fi10)
        }
        var v13 string = ("hello" + ("\n\t" + "hello"))
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    var v3 uint32 = 0
    v4 := make(map[uint32]int)
    var v5 int64 = 1
    var v6 string = ""
    var v14 uint32 = v3
    v3 = (v3 * 4294967295)
    switch -434 {
    case 1:
        v14 += 65535
    case 0:
        if false {
            if ((1e10 * (3.14159 + 0.0)) > (float64(3.14159) * (float64(0.1) - -2.25))) {
                _ = len(v4)
                if (9223372036854775807 <= -9223372036854775808) {
                    v4[(v3 ^ v14)] = 128
                }
            } else {
                var v15 uint32 = genFn0(v1)
                v16 := GenS1{F0: -1, F1: 937}
                _ = v15
                _ = v16
            }
            var v17 uint32 = 4294967295
            _ = v17
        }
    case 207:
        delete(v4, (v3 & genFn0(v1)))
    }
    if (v5 > v5) {
        var v18 int = 32767
        fmt.Println(v3, v18)
        v1 = 1
        _ = v18
    }
    _ = len(v2)
    fmt.Println(v6, v1, v14)
    var v19 int16 = (v1 << ((696) & 63))
    v14 = 735
    for fi20 := 0; fi20 < 14; fi20++ {
        v6 = "\n\t"
        _ = fi20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v2[(v14 << ((415) & 63))] = (3.14159 < (1.5 * float64(3.14159)))
        var v22 uint32 = (v14 + (v14 | v14))
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
