package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int8, p1 int) uint8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: (("hello" + "a") + "a"), F1: 1, F2: 1}
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 200
    _ = p0
    _ = p1
    _ = v8
    return v8
}
func genFn1() uint32 {
    v9 := GenS1{F0: ("hello" + "foo bar"), F1: 2, F2: 42}
    var v10 int8 = -128
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    var v4 int16 = 678
    var v5 uint8 = (v1 | 220)
    if ((-308 == 0) && ((712 >= 4294967295) && false)) {
        v11 := GenS1{F0: "a", F1: 18446744073709551615, F2: 0}
        if true {
            fmt.Println(v5, v1)
            var v12 string = ("hello" + "foo bar")
            v13 := v2[65535]
            _ = v12
            _ = v13
        }
        _ = v11
    } else {
        var v14 uint64 = 904
        v3 = v3
        v3 = v3
        _ = v14
    }
    delete(v2, 101)
    var v15 uint64 = 4294967295
    fmt.Println(v1)
    fmt.Println(v3, v3, v3)
    fmt.Println(v4)
    var v16 int64 = -420
    if false {
        if false {
            v16 -= (-(v16))
            var v17 uint32 = genFn1()
            _ = v17
        } else {
            fmt.Println(v4)
            fmt.Println(v3, v1, v4)
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := []int{-2147483648}
        fmt.Println(fi18)
        v4 += (^(v4))
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: "foo bar", F1: 0, F2: 32767}
    if (v5 >= (^(v1))) {
        if (false && ((false || (3.14159 == 0.0)) && false)) {
            v21 := GenS2{F0: 1}
            if (!true) {
                var v22 int16 = 0
                _ = v22
            } else {
                v2[genFn1()] = (!(1 <= 127))
            }
            _ = v21
        } else {
            var v23 uint16 = 0
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
