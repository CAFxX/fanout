package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 bool, p1 int, p2 int32) int8 {
    v6 := GenS1{F0: 0, F1: 127}
    v7 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 int64) int16 {
    fmt.Println(p1)
    p1 += p1
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []string{"a", "\n\t", (("" + "a") + ("\n\t" + "a")), "\n\t"}
    var v2 bool = true
    var v3 uint32 = 1
    v4 := []int{}
    v5 := []string{((("hello" + ("hello" + "foo bar")) + ("foo bar" + "a")) + "\n\t")}
    var v8 int = 4294967295
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: -32768, F1: -116}
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    v2 = true
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 1
        fmt.Println(v8, v8, v3)
        _ = fi11
        _ = v12
    }
    _ = len(v1)
    v8 = 997
    if v2 {
        v1 = append(v1, (("a" + "\n\t") + (("a" + "hello") + ("hello" + "a"))))
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint8 = 1
                _ = fi14
                _ = v15
            }
            v2 = (4294967295 != 42)
            if true {
                var v16 uint8 = 1
                var v17 uint64 = 2
                _ = v16
                _ = v17
            } else {
                v2 = false
                var v18 uint64 = 728
                v4 = append(v4, -687)
                _ = v18
            }
            _ = fi13
        }
        if (437 != (v3 + (v3 | (v3 >> ((473) & 63))))) {
            v3 -= (^(v3))
            if v2 {
                var v19 float64 = 0.1
                if (!(!false)) {
                    var v20 uint16 = 2
                    v21 := GenS2{F0: false, F1: v8, F2: 42}
                    _ = v20
                    _ = v21
                } else {
                    v3 += (v3 >> ((791) & 63))
                }
                _ = v19
            }
            fmt.Println(v8, v2)
        }
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v1 = append(v1, ((("a" + ("\n\t" + "\n\t")) + "") + "hello"))
        _ = fi22
    }
    var v23 int32 = -2147483648
    var v24 bool = (!true)
    var v25 uint8 = 0
    if (((!v24) && true) || (2 <= 42)) {
        v26 := GenS3{F0: false, F1: v3, F2: v3}
        var v27 int32 = (v23 >> ((255) & 63))
        _ = v26
        _ = v27
    } else {
        var v28 uint16 = 2
        var v29 float64 = 3.14159
        _ = v28
        _ = v29
    }
    var v30 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
