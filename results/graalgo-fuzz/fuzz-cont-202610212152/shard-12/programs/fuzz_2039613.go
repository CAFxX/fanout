package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 int, p1 uint16, p2 uint64) string {
    fmt.Println(p1)
    switch -9223372036854775808 {
    case 9223372036854775807:
        p0 = -2147483648
    case -452:
        var v7 uint16 = p1
        _ = v7
    case 385:
        if true {
            fmt.Println(p1, p2, p1)
            p2 += p2
        } else {
            v8 := []int{}
            var v9 uint16 = p1
            for fi10 := 0; fi10 < 5; fi10++ {
                _ = len(v8)
                _ = fi10
            }
            _ = v8
            _ = v9
        }
    }
    var v11 bool = false
    fmt.Println(p0, v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return ("\n\t" + "hello")
}
func genFn1(p0 int32, p1 uint) int16 {
    var v12 bool = true
    fmt.Println(v12, p0)
    for fi13 := 0; fi13 < 14; fi13++ {
        if true {
            if (42 < 0) {
                v14 := []int{}
                p0 = -448
                _ = v14
            }
        } else {
            var v15 int64 = 771
            _ = v15
        }
        v16 := GenS1{F0: 42}
        _ = fi13
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v12
    return -32768
}

func main() {
    var v1 uint = 438
    var v2 uint32 = 0
    var v3 int = -2147483648
    var v4 string = "foo bar"
    var v5 int16 = 676
    var v6 string = ("a" + ("\n\t" + ("a" + v4)))
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    var v19 float64 = ((float64(((0.1 / -2.25) - 3.14159)) - 1e10) * ((float64(-2.25) * 0.0) + 0.1))
    if true {
        var v20 int16 = (v5 % ((6) | 1))
        v4 = "hello"
        _ = v20
    }
    v1 += v1
    fmt.Println(v3, v1, v3)
    v21 := GenS2{F0: 959, F1: (^(v3))}
    v22 := GenS1{F0: 0}
    v1 = v1
    v23 := GenS1{F0: 144}
    var v24 uint64 = 0
    var v25 int16 = (-(v5))
    fmt.Println(v5, v3)
    var v26 uint32 = v2
    switch (("\n\t" + "hello") + "hello") {
    case "\n\t":
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                v25 += genFn1(1, v1)
                v29 := []int{(fi28 | (fi28 << ((uint(v26)) & 63))), 0, fi28}
                var v30 int32 = 97
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi27
        }
    }
    var v31 uint16 = uint16(v1)
    v19 += v19
    v32 := []int{(^(v3)), v3, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
