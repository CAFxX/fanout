package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 uint32, p1 int) bool {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    switch (p1 % ((p1) | 1)) {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 string = ("hello" + "foo bar")
            var v9 float32 = (((1.5 - -2.25) - (1e10 * (1.5 + -2.25))) - 1e10)
            v10 := GenS1{F0: p1, F1: 127, F2: (!false)}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        v11 := []int{9223372036854775807, 2147483647, (-(p1))}
        _ = v11
    }
    _ = p0
    _ = p1
    return (true && (float32(float64(3.14159)) < ((-2.25 * 0.0) - float32(-2.25))))
}

func main() {
    var v1 int32 = 42
    var v2 bool = false
    var v3 int8 = 42
    var v4 uint = 565
    v2 = (!(!(!true)))
    switch "hello" {
    case "":
        v12 := []int{251, -614}
        _ = v12
    }
    v2 = true
    switch (("hello" + "\n\t") + ("hello" + "a")) {
    case "\n\t":
        fmt.Println(v2)
    }
    v1 = (v1 | v1)
    v1 -= (v1 >> ((v4) & 63))
    for fi13 := 0; fi13 < 10; fi13++ {
        if (!(65535 != 1)) {
            v3 += (v3 - (v3 | (v3 | v3)))
            v14 := []int{}
            var v15 uint32 = 710
            _ = v14
            _ = v15
        }
        var v16 float64 = 3.14159
        _ = fi13
        _ = v16
    }
    fmt.Println(v1)
    switch 42 {
    case 42:
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 float32 = ((0.0 * 0.0) * 3.14159)
            _ = fi17
            _ = v18
        }
    case 32768:
        fmt.Println(v3, v1, v4)
    case 127:
        var v19 string = ("a" + "")
        _ = v19
    default:
        var v20 bool = false
        _ = v20
    }
    v21 := GenS1{F0: 127, F1: 100, F2: false}
    var v22 uint32 = 1
    for fi23 := 0; fi23 < 20; fi23++ {
        v4 -= 41
        v1 -= v1
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
