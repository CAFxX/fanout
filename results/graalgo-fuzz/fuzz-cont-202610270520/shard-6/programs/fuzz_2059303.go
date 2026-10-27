package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint64, p1 int32, p2 int) string {
    var v7 int16 = -32768
    var v8 uint16 = 227
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (("foo bar" + "hello") + "hello")
}
func genFn1(p0 int) bool {
    v9 := GenS1{F0: 128, F1: -353}
    p0 -= (p0 - -2147483648)
    _ = p0
    _ = v9
    return (false && (!(128 <= 4294967295)))
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 1}
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := []string{(("hello" + "\n\t") + "foo bar"), (("\n\t" + ("hello" + "foo bar")) + "a"), (("\n\t" + "") + ("hello" + "\n\t")), ("" + "\n\t")}
    var v10 uint8 = 128
    switch -9223372036854775808 {
    case 127:
        var v11 float64 = -2.25
        _ = v11
    case 4294967295:
        var v12 string = ""
        _ = v12
    case -705:
        v13 := v2[2147483647]
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v10)
        if true {
            v10 -= v10
        } else {
            switch (fi14 & fi14) {
            case 32768:
                v15 := GenS2{F0: 896, F1: 1}
                _ = v15
            }
            var v16 uint = 4294967295
            _ = v16
        }
        if len(v3) > 0 {
            v3[0] = 936
        }
        _ = fi14
    }
    fmt.Println(v10, v1)
    v17 := GenS1{F0: v10, F1: (-(v1))}
    var v18 string = "a"
    delete(v2, -1)
    delete(v4, v18)
    fmt.Println(v18)
    v19 := GenS2{F0: 2, F1: -490}
    delete(v4, (v18 + (v18 + (v18 + ((v18 + v18) + (v18 + v18))))))
    fmt.Println(v1, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
