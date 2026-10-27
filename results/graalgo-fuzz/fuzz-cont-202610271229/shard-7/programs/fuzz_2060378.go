package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool, p1 int16) uint8 {
    p1 += (p1 ^ -1)
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "hello" {
        case "foo bar":
            var v8 int = fi7
            _ = v8
        case "\n\t":
            if p0 {
                p0 = true
                var v9 uint64 = 893
                _ = v9
            }
        case "a":
            p1 -= (^(p1))
        }
        p0 = ((1e10 < 3.14159) && ((-2.25 == 3.14159) || (-1 <= -422)))
        _ = fi7
    }
    if (false && (-1 >= 0)) {
        fmt.Println(p0)
        v10 := []int{2147483647}
        if p0 {
            var v11 int16 = p1
            _ = v11
        } else {
            var v12 uint64 = 1
            v12 = (v12 - v12)
            fmt.Println(p1, p1, p0)
            _ = v12
        }
        _ = v10
    }
    if (3.14159 < (3.14159 / 0.1)) {
        v13 := []int{-465, -9223372036854775808, 9223372036854775807}
        _ = len(v13)
        if (p0 || (-128 <= 1)) {
            if len(v13) > 0 {
                v13[0] = -2147483648
            }
        } else {
            var v14 uint16 = 0
            var v15 int8 = 0
            _ = v14
            _ = v15
        }
        _ = v13
    }
    p0 = (0.1 > 0.0)
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 uint8, p1 int16, p2 string) uint8 {
    var v16 int32 = 546
    p0 += (p0 << ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (p0 ^ genFn0(false, 42))
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{581, 350}
    v3 := []string{"\n\t", "foo bar"}
    var v4 string = "a"
    var v5 string = ((v4 + "") + "hello")
    var v6 int16 = v1
    if false {
        var v17 uint64 = 1
        var v18 float32 = 1e10
        var v19 uint16 = 2
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint32 = 488
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint = 255
        _ = fi21
        _ = v22
    }
    var v23 int = 32768
    if (812 == 44) {
        fmt.Println(v4)
    }
    var v24 int = v23
    v25 := GenS1{F0: (v6 + -32768)}
    var v26 int8 = 1
    switch v5 {
    case "foo bar":
        var v27 uint8 = genFn1(127, (v6 ^ -32768), v5)
        _ = v27
    case "a":
        v28 := GenS1{F0: (^(v1))}
        _ = v28
    case "\n\t":
        var v29 float64 = float64(0.0)
        _ = v29
    }
    v23 += v23
    var v30 float64 = (-2.25 + (0.1 + -2.25))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v31 float64 = v30
    if false {
        var v32 uint8 = 255
        v20 -= (v20 ^ (v20 + (v20 << ((0) & 63))))
        v33 := GenS1{F0: v6}
        _ = v32
        _ = v33
    }
    var v34 float32 = 0.1
    v4 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
