package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 int }
func genFn0() string {
    switch 42 {
    case 4294967295:
        var v6 int8 = 100
        _ = v6
    case 127:
        var v7 float64 = (3.14159 - ((0.1 * ((1e10 - 3.14159) / float64(-2.25))) - 1e10))
        _ = v7
    case -2147483648:
        v8 := GenS1{F0: 42, F1: (!(false || false)), F2: false}
        _ = v8
    }
    v9 := GenS1{F0: 447, F1: (0 == 205), F2: (("a" + ("a" + "\n\t")) >= "a")}
    var v10 uint32 = 65535
    v11 := GenS1{F0: 2, F1: true, F2: (float32((3.14159 * 0.1)) < 1e10)}
    v10 += v10
    _ = v10
    _ = v11
    _ = v9
    return ""
}
func genFn1(p0 int64, p1 uint8) int8 {
    v12 := []int{1, 544, 127, 4294967295}
    v12 = append(v12, 0)
    var v13 bool = ((p0 << ((65535) & 63)) >= p0)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -128
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int8 = -128
    v3 := []string{"foo bar", (("hello" + ("" + "hello")) + ("a" + ("hello" + (("a" + "a") + "foo bar")))), "\n\t", ("a" + ("\n\t" + ((("\n\t" + "a") + "") + ("foo bar" + ""))))}
    var v4 int16 = -1
    v5 := make(map[uint32]uint32)
    if (!true) {
        var v14 int16 = v4
        _ = v14
    } else {
        v2 += v2
        switch 9223372036854775807 {
        case 0:
            v15 := GenS1{F0: 1, F1: (327 <= 42), F2: false}
            _ = v15
        default:
            var v16 int64 = -9223372036854775808
            _ = v16
        }
        v4 -= v4
    }
    switch (-(v1)) {
    case 0:
        fmt.Println(v2, v1, v2)
    case 32768:
        v17 := GenS1{F0: 0, F1: false, F2: false}
        _ = v17
    case 127:
        delete(v5, 42)
    default:
        fmt.Println(v1, v4)
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint = 65535
        switch (("hello" + "hello") + (("" + ("\n\t" + "a")) + ("a" + (("\n\t" + ("foo bar" + "a")) + "foo bar")))) {
        case "foo bar":
            var v20 int = fi18
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v2)
    var v21 string = "foo bar"
    if len(v3) > 0 {
        v3[0] = v21
    }
    v22 := GenS2{F0: (-(v1))}
    delete(v5, 2147483647)
    switch (v1 | 868) {
    case 32767:
        fmt.Println(v21, v1, v4)
    }
    var v23 uint16 = 582
    fmt.Println(v23, v4)
    _ = len(v3)
    _ = len(v5)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
