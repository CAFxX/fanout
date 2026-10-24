package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string, p1 uint) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    var v5 float64 = 0.0
    v6 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint8, p1 uint8, p2 int) uint64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p2 = 42
        _ = fi7
    }
    fmt.Println(p0, p2)
    if true {
        fmt.Println(p1, p1)
    } else {
        var v8 float32 = -2.25
        _ = v8
    }
    p2 = (p2 / ((969) | 1))
    v9 := []int{(p2 / ((p2) | 1)), p2, (p2 | (p2 % (((p2 - 128)) | 1))), (p2 * (p2 & p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = false
    v3 := []string{"foo bar", (("foo bar" + ("a" + "\n\t")) + "a")}
    v4 := []uint32{(v1 / (((v1 << ((329) & 63))) | 1)), v1}
    if v2 {
        var v10 uint32 = (v1 + v1)
        _ = v10
    }
    v2 = ((574 < 42) || false)
    v11 := GenS1{F0: 42}
    v2 = (v2 && (((float32(-2.25) + 3.14159) >= (1.5 / -2.25)) || v2))
    var v12 float64 = 0.1
    var v13 int = -1
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int64 = 9223372036854775807
            v4 = append(v4, v1)
            var v17 uint16 = 42
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int32 = 1
        _ = fi14
        _ = v18
    }
    switch (("" + "a") + (("foo bar" + "") + "\n\t")) {
    case "a":
        var v19 float64 = v12
        _ = v19
    default:
        var v20 bool = false
        _ = v20
    }
    v21 := GenS1{F0: -1}
    v12 += (3.14159 - 1.5)
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int = 128
        v2 = (1e10 < (3.14159 / (3.14159 * (3.14159 - (1.5 / 3.14159)))))
        v1 -= v1
        _ = fi22
        _ = v23
    }
    var v24 float32 = 1e10
    _ = len(v4)
    fmt.Println(v2)
    switch ("" + ((((("foo bar" + ("hello" + ((("" + "foo bar") + ("\n\t" + "")) + ("\n\t" + "")))) + (("" + "hello") + "foo bar")) + "hello") + "\n\t") + (((("" + (("\n\t" + "a") + "foo bar")) + (("" + (("" + "a") + "a")) + "a")) + ("\n\t" + (("a" + ("hello" + "")) + ((("a" + ("hello" + "a")) + ("a" + "foo bar")) + "a")))) + ("a" + "\n\t")))) {
    case "":
        v4 = append(v4, v1)
    case "hello":
        v2 = true
    }
    v4 = append(v4, 290)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
