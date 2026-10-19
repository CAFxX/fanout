package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) string {
    var v4 float32 = 3.14159
    var v5 int32 = 0
    var v6 int64 = 1
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ((("foo bar" + "a") + "hello") + ((("foo bar" + ("foo bar" + "foo bar")) + ((("hello" + "hello") + ("\n\t" + "a")) + "foo bar")) + ("\n\t" + "foo bar")))
}
func genFn1(p0 string, p1 int, p2 uint32) uint8 {
    var v7 int8 = int8(p1)
    v8 := GenS1{F0: "foo bar"}
    var v9 float64 = -2.25
    var v10 int64 = 820
    v7 = int8(v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    v1 -= genFn1(("foo bar" + "hello"), -527, 1)
    var v11 int64 = 0
    var v12 int16 = v2
    fmt.Println(v11, v12)
    delete(v3, 408)
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := []int{179}
            _ = len(v15)
            fmt.Println(fi13, v12)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v16 int8 = 75
    if (true || false) {
        v17 := []int{4294967295, 2147483647, 0}
        _ = v17
    }
    var v18 int64 = (v11 % ((v11) | 1))
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v16 -= 105
            v21 := []int{fi20, fi19}
            var v22 uint64 = 18446744073709551615
            fmt.Println(v11)
            var v23 int16 = v2
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi19
    }
    if (-574 < -252) {
        if false {
            v24 := GenS2{F0: ((true || (!false)) && true)}
            var v25 float32 = -2.25
            _ = len(v3)
            _ = v24
            _ = v25
        } else {
            v1 = 169
            var v26 int8 = -128
            _ = v26
        }
        var v27 string = (("hello" + "\n\t") + "hello")
        v28 := GenS1{F0: v27}
        _ = v27
        _ = v28
    }
    if false {
        var v29 bool = (838 < 1)
        _ = v29
    } else {
        v30 := v3[74]
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
