package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool, p2 int32) uint64 {
    var v7 int16 = 42
    switch 32768 {
    case -856:
        v8 := []int{9223372036854775807, 9223372036854775807, 1}
        _ = v8
    }
    v9 := []int{}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 1
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 193
    v3 := []string{"a", "foo bar", ("a" + "hello"), ("a" + ("a" + ("hello" + ((("hello" + (("" + ((("foo bar" + "a") + "foo bar") + "foo bar")) + "foo bar")) + ("hello" + "foo bar")) + "foo bar"))))}
    var v4 uint = (v1 << (((v1 + v1)) & 63))
    v5 := []int{32767, -606, 127, -1}
    v6 := []string{"", "a"}
    v3 = append(v3, "foo bar")
    var v10 uint64 = 1
    var v11 uint64 = genFn0((v2 - uint16(v10)), (!(!false)), 1)
    fmt.Println(v1)
    v12 := GenS1{F0: 127, F1: 65535}
    if true {
        if (1.5 == float32(1.5)) {
            if true {
                var v13 uint = v1
                for fi14 := 0; fi14 < 3; fi14++ {
                    if (fi14 != fi14) {
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                    } else {
                        v15 := GenS1{F0: -82, F1: 455}
                        _ = v15
                    }
                    v16 := GenS1{F0: -1, F1: 42}
                    _ = fi14
                    _ = v16
                }
                var v17 uint = 255
                _ = v13
                _ = v17
            }
            v1 = (v1 ^ (^(v1)))
        }
        fmt.Println(v10)
    }
    v18 := GenS2{F0: -603}
    if (true && true) {
        v2 -= 1
    }
    if len(v6) > 0 {
        v6[0] = ""
    }
    v19 := GenS1{F0: -61, F1: 2147483648}
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v20 int32 = 2147483647
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint64 = (v10 % ((4294967295) | 1))
        _ = fi21
        _ = v22
    }
    _ = len(v3)
    var v23 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
