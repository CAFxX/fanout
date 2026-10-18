package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 string, p2 uint64) int32 {
    var v5 uint = 1
    var v6 uint16 = 223
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int, p1 int16) int {
    p0 += p0
    v7 := []int{(p0 + p0), -9223372036854775808, p0}
    if len(v7) > 0 {
        v7[0] = (p0 - (p0 - (p0 ^ p0)))
    }
    var v8 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -352
}

func main() {
    var v1 float64 = ((1e10 - (0.0 * -2.25)) * 0.1)
    var v2 int32 = 1
    var v3 string = ((("hello" + (("a" + ("hello" + "foo bar")) + ("hello" + ""))) + "") + "\n\t")
    var v4 int64 = 234
    v9 := GenS1{F0: 4294967295}
    for fi10 := 0; fi10 < 20; fi10++ {
        v4 += (-(v4))
        _ = fi10
    }
    v1 += float64(0.1)
    v1 = ((float64(-2.25) / ((v1 - v1) * v1)) + (-2.25 + v1))
    if (42 > 706) {
        var v11 int64 = v4
        var v12 uint32 = 275
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 += v2
            v14 := []int{genFn1((fi13 + 128), -32768)}
            _ = fi13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    v3 = (("foo bar" + (("foo bar" + (("hello" + v3) + ((("foo bar" + v3) + ("foo bar" + "a")) + "a"))) + v3)) + (v3 + v3))
    var v15 uint64 = 139
    var v16 float64 = (v1 - 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
