package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int, p1 int32) uint {
    v4 := []int{p0, (p0 >> ((933) & 63)), 32767, 9223372036854775807}
    var v5 float32 = (float32(1e10) - 1e10)
    if false {
        switch (("a" + (("a" + "foo bar") + (("\n\t" + "hello") + "\n\t"))) + "") {
        case "":
            switch (("a" + "") + "a") {
            case "a":
                v4 = append(v4, 276)
            }
        case "foo bar":
            p1 = (-(p1))
        case "hello":
            v5 -= v5
        }
        var v6 uint32 = 584
        _ = v6
    }
    p1 = -87
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 bool, p1 bool) uint16 {
    if (0 <= 127) {
        p1 = (3.14159 > 1e10)
    } else {
        v7 := []int{146, -9223372036854775808, 2147483647, 127}
        _ = v7
    }
    var v8 uint32 = 305
    _ = p0
    _ = p1
    _ = v8
    return uint16(v8)
}

func main() {
    v1 := []int{}
    var v2 string = "hello"
    var v3 int32 = -2147483648
    v9 := GenS1{F0: 75, F1: 100, F2: 2147483648}
    var v10 string = (v2 + ("hello" + "a"))
    var v11 int8 = -1
    if (!true) {
        switch v10 {
        case "a":
            v2 = (v10 + "")
        case "foo bar":
            v2 = "hello"
        case "":
            var v12 uint16 = 0
            _ = v12
        default:
            v2 = ((("\n\t" + v2) + v10) + "\n\t")
        }
    }
    v1 = append(v1, 128)
    v2 = (v10 + ("foo bar" + v10))
    v3 += v3
    v1 = append(v1, -257)
    v10 = "foo bar"
    var v13 int8 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= -2147483648
    v14 := GenS1{F0: (^(v13)), F1: (v13 & v11), F2: 0}
    for fi15 := 0; fi15 < 15; fi15++ {
        if false {
            v16 := GenS2{F0: 736}
            v1 = append(v1, 0)
            _ = v16
        }
        var v17 string = ((v2 + v10) + v2)
        _ = fi15
        _ = v17
    }
    v18 := GenS2{F0: 144}
    for fi19 := 0; fi19 < 1; fi19++ {
        v3 = v3
        v13 = 61
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
