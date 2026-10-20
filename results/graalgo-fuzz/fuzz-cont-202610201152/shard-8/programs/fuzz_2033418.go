package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string) int {
    v6 := []int{}
    p0 = (("\n\t" + p0) + p0)
    p0 = ""
    _ = p0
    _ = v6
    return -15
}
func genFn1(p0 uint) uint64 {
    v7 := []int{128, -9223372036854775808, 128}
    var v8 int = 4294967295
    fmt.Println(v8, v8)
    v8 = genFn0(("\n\t" + "foo bar"))
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []string{"foo bar", ("a" + "\n\t"), ("" + "a")}
    var v4 bool = true
    v5 := make(map[int]int)
    if (!v4) {
        v9 := GenS1{F0: v4}
        _ = v9
    } else {
        var v10 float64 = (1.5 / ((1e10 - 1.5) - float64(float32(0.1))))
        v10 -= 1.5
        _ = v10
    }
    if v4 {
        v11 := GenS1{F0: v4}
        _ = v11
    }
    switch ("\n\t" + ("\n\t" + ("" + "a"))) {
    case "a":
        v4 = (255 >= 2)
    }
    if ((1 == 2) || v4) {
        var v12 uint16 = 1
        _ = v12
    } else {
        v4 = v4
    }
    var v13 int64 = -269
    if v4 {
        if (!(1.5 == 1e10)) {
            var v14 uint = 680
            v4 = (0.1 != (-2.25 * 0.1))
            _ = v14
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                v13 = 0
                _ = fi15
            }
        }
        v16 := GenS2{F0: 193}
        v13 -= (v13 | v13)
        _ = v16
    }
    if (!v4) {
        v13 = 42
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        if (((("\n\t" + "\n\t") + "") + "\n\t") <= (("hello" + "hello") + (("foo bar" + (("foo bar" + "") + "")) + "hello"))) {
            var v18 string = "\n\t"
            _ = v18
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            v4 = (527 <= 2)
            _ = fi19
        }
        _ = fi17
    }
    v3 = append(v3, "foo bar")
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int8 = -26
        v22 := GenS2{F0: 127}
        v23 := GenS2{F0: 0}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v4 = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
