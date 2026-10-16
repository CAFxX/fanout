package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int, p1 bool, p2 string) int8 {
    var v4 uint16 = 65535
    var v5 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -33
}
func genFn1(p0 uint8, p1 uint) uint16 {
    switch 9223372036854775807 {
    case 212:
        var v6 uint8 = p0
        _ = v6
    case 2147483647:
        v7 := []int{-2147483648, -2147483648, -343}
        _ = v7
    case -149:
        p1 += p1
    }
    p0 = (p0 & 255)
    p0 -= (p0 >> ((p1) & 63))
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 float32 = float32((((3.14159 - 0.1) - (1.5 + (0.1 + (1e10 / 1.0)))) / (1.5 + ((3.14159 * 0.0) + float64((1e10 / 1e10))))))
    v2 := []string{}
    v3 := []uint32{}
    var v8 uint32 = 830
    v8 += v8
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(v1, fi9)
        fmt.Println(fi9, v1, fi9)
        v8 = (v8 / ((0) | 1))
        _ = fi9
    }
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "foo bar":
        v1 -= ((float32((1e10 - 3.14159)) / v1) * v1)
    case "a":
        v1 += -2.25
    case "\n\t":
        fmt.Println(v1)
    }
    var v10 uint32 = (v8 ^ v8)
    for fi11 := 0; fi11 < 15; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v12 float64 = 0.0
        _ = fi11
        _ = v12
    }
    if false {
        v13 := GenS1{F0: 0}
        v2 = append(v2, ("\n\t" + ("" + (("hello" + "") + "hello"))))
        var v14 uint32 = 2147483648
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
}
