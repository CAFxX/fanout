package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 string, p1 int64, p2 int) int16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        p0 = (p0 + "")
        _ = fi6
    }
    var v7 int64 = (p1 ^ -943)
    var v8 int64 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := []int{-9223372036854775808, 128, 1, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}
func genFn1() uint8 {
    var v11 float32 = -2.25
    if (2 < 65535) {
        var v12 float32 = -2.25
        _ = v12
    }
    v11 -= (3.14159 * ((((v11 / (v11 * v11)) / v11) - v11) / (v11 - 3.14159)))
    var v13 string = ""
    _ = v11
    _ = v13
    return 1
}

func main() {
    var v1 int32 = 617
    var v2 float32 = 3.14159
    var v3 float32 = (v2 / v2)
    var v4 string = "a"
    var v5 uint16 = 48
    var v14 bool = false
    var v15 uint8 = genFn1()
    fmt.Println(v15, v15)
    v16 := GenS1{F0: genFn0("foo bar", -492, 32767)}
    for fi17 := 0; fi17 < 2; fi17++ {
        switch ("" + ("a" + v4)) {
        case "hello":
            v18 := GenS1{F0: -435}
            _ = v18
        case "":
            v4 = (("a" + ("a" + "hello")) + "")
        }
        _ = fi17
    }
    switch "\n\t" {
    case "a":
        var v19 int = 32767
        _ = v19
    }
    var v20 uint = 647
    var v21 float32 = v2
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
