package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int32, p1 uint16) bool {
    var v6 int8 = 127
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = (p0 / ((p0) | 1))
        v8 += v8
        var v9 uint = 4294967295
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{fi10}
        if false {
            v12 := GenS1{F0: fi10}
            _ = v12
        } else {
            if false {
                v6 += 42
                var v13 uint32 = 2147483648
                _ = v13
            } else {
                v6 = 42
            }
            p0 += 2147483647
            for fi14 := 0; fi14 < 2; fi14++ {
                switch (fi14 / ((127) | 1)) {
                case 1:
                    var v15 float64 = (1e10 + (1.5 * (1e10 + 0.0)))
                    _ = v15
                case 0:
                    fmt.Println(fi10)
                default:
                    p0 += p0
                }
                fmt.Println(p0, p1, v6)
                _ = fi14
            }
        }
        _ = fi10
        _ = v11
    }
    var v16 string = ("hello" + "")
    p1 -= (p1 & (p1 ^ p1))
    _ = p0
    _ = p1
    _ = v16
    _ = v6
    return false
}
func genFn1() uint16 {
    var v17 uint32 = 288
    switch "a" {
    case "hello":
        var v18 int = 32768
        _ = v18
    case "a":
        v19 := []int{9223372036854775807}
        _ = v19
    case "":
        var v20 uint = 1
        _ = v20
    }
    var v21 float32 = (3.14159 * 3.14159)
    v21 = v21
    v21 = (v21 - 0.0)
    _ = v17
    _ = v21
    return 32767
}

func main() {
    var v1 float32 = (3.14159 * -2.25)
    var v2 int = 0
    var v3 uint16 = 0
    v4 := []int{v2, (v2 ^ v2), v2, v2}
    v5 := []int{(v2 + 909), v2, -1}
    var v22 int64 = 0
    var v23 float32 = (1e10 * -2.25)
    for fi24 := 0; fi24 < 6; fi24++ {
        v3 = 42
        v23 = v23
        _ = fi24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v26 uint16 = 255
        if len(v4) > 0 {
            v4[0] = fi25
        }
        _ = fi25
        _ = v26
    }
    var v27 string = "\n\t"
    fmt.Println(v1, v2)
    var v28 int = 862
    v22 = 811
    fmt.Println(v3, v22)
    if len(v4) > 0 {
        v4[0] = (v2 / (((v28 / ((2147483647) | 1))) | 1))
    }
    if (-2.25 >= 1.5) {
        var v29 int32 = 2147483647
        v5 = append(v5, (^(v28)))
        fmt.Println(v27, v2, v29)
        _ = v29
    } else {
        var v30 uint64 = 42
        v23 += (1e10 * v23)
        if (42 != 2) {
            var v31 int8 = 95
            _ = v31
        }
        _ = v30
    }
    var v32 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
