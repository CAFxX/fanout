package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 uint, p2 int) int8 {
    p2 = 9223372036854775807
    v5 := GenS1{F0: -9223372036854775808}
    if (!(1 >= 0)) {
        var v6 float32 = 0.1
        _ = v6
    } else {
        fmt.Println(p2, p2, p2)
        fmt.Println(p2, p0)
    }
    var v7 uint32 = 885
    switch -424 {
    case -100:
        v8 := []int{}
        _ = v8
    case -2147483648:
        if (1 < 2147483647) {
            switch int(p0) {
            case -1:
                v9 := []int{}
                _ = v9
            }
        }
    case 58:
        p0 = (p0 + (p0 >> (((^(p1))) & 63)))
    default:
        var v10 uint8 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 29
}
func genFn1(p0 string, p1 int16, p2 int) uint32 {
    var v11 float32 = 0.1
    v12 := []int{4294967295, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    v1 := []string{(("foo bar" + "\n\t") + (("\n\t" + ("" + ("hello" + "\n\t"))) + "")), "a", "hello", "hello"}
    var v2 float32 = (1.5 * 0.1)
    var v3 string = ("" + ("hello" + ""))
    var v4 uint16 = 137
    switch ("foo bar" + "hello") {
    case "a":
        var v13 float32 = (v2 / v2)
        _ = v13
    case "hello":
        v14 := GenS1{F0: -108}
        _ = v14
    }
    if (((-937 < 128) || true) && (!(!(101 > 2)))) {
        v4 -= 65535
        v15 := GenS2{F0: uint(v4)}
        _ = v15
    }
    v2 = v2
    var v16 float32 = float32(1.5)
    v2 = v16
    if len(v1) > 0 {
        v1[0] = v3
    }
    v17 := GenS1{F0: 41}
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
