package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint16, p1 uint8, p2 uint) string {
    p1 = p1
    var v4 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ""
}
func genFn1(p0 uint64, p1 uint64) uint32 {
    if false {
        if (!false) {
            var v5 string = "foo bar"
            v6 := GenS1{F0: 0, F1: (1 < 128)}
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 bool = (!false)
            p1 = (p0 / (((p0 & (p1 & 42))) | 1))
            _ = fi7
            _ = v8
        }
    }
    var v9 int8 = 1
    _ = p0
    _ = p1
    _ = v9
    return 844
}

func main() {
    var v1 uint64 = 2
    var v2 int8 = 42
    var v3 string = ("\n\t" + (("\n\t" + ("a" + (("foo bar" + (("\n\t" + ("hello" + "foo bar")) + "hello")) + "\n\t"))) + ((("a" + "hello") + "foo bar") + (("hello" + "") + ("\n\t" + "\n\t")))))
    var v10 int16 = -32768
    for fi11 := 0; fi11 < 1; fi11++ {
        v2 += (v2 / ((-65) | 1))
        v2 += v2
        v3 = v3
        _ = fi11
    }
    v12 := GenS1{F0: 0, F1: true}
    v10 -= (v10 & 1)
    v10 -= v10
    v1 += 0
    var v13 uint64 = v1
    v14 := GenS2{F0: 1, F1: 83}
    fmt.Println(v10, v13)
    v15 := GenS1{F0: 2147483647, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
