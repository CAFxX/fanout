package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint, p1 string) uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{(^(fi5))}
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint64 = 42
            var v9 int = fi7
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    p1 = (("a" + (p1 + "foo bar")) + "foo bar")
    v10 := []int{32767, 0, 554}
    _ = p0
    _ = p1
    _ = v10
    return 4294967295
}
func genFn1(p0 uint, p1 string, p2 int16) int32 {
    v11 := []int{534, 127, 4294967295}
    var v12 int32 = 888
    var v13 bool = (false || true)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return -619
}

func main() {
    v1 := []int{-1, -747}
    var v2 int32 = 1
    var v3 uint32 = 0
    var v4 float32 = (1e10 + (1.5 + 1e10))
    for fi14 := 0; fi14 < 20; fi14++ {
        v4 -= v4
        fmt.Println(v2)
        var v15 uint16 = 0
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v3)
    var v16 float64 = 0.1
    fmt.Println(v16, v3, v16)
    v17 := GenS1{F0: 100}
    for fi18 := 0; fi18 < 4; fi18++ {
        switch (fi18 & fi18) {
        case 364:
            if len(v1) > 0 {
                v1[0] = 42
            }
        case 674:
            var v19 uint64 = 18446744073709551615
            _ = v19
        default:
            v1 = append(v1, (fi18 - (fi18 * 33)))
        }
        var v20 int32 = (v2 | -554)
        _ = fi18
        _ = v20
    }
    v3 = (v3 - 4294967295)
    v4 -= ((v4 - v4) - (((((3.14159 * (v4 - 3.14159)) + ((3.14159 / v4) * 1.5)) * v4) * v4) * (3.14159 / v4)))
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v1) > 0 {
            v1[0] = fi21
        }
        _ = fi21
    }
    _ = len(v1)
    switch ("" + "") {
    case "":
        var v22 int16 = -506
        _ = v22
    case "hello":
        v16 -= 0.0
    case "a":
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v4 += (v4 / v4)
        default:
            v4 += v4
        }
    }
    v23 := GenS1{F0: 0}
    v1 = append(v1, -2147483648)
    var v24 bool = (!(!true))
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
