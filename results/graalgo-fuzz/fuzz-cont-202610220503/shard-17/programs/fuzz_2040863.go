package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0() int {
    v4 := GenS1{F0: false}
    var v5 int8 = 5
    if false {
        var v6 string = "\n\t"
        v7 := GenS2{F0: 133, F1: 2147483647, F2: (131 == 65535)}
        switch -9223372036854775808 {
        case 4294967295:
            v8 := []int{32768, 4294967295, 9223372036854775807, 55}
            _ = v8
        case -2147483648:
            var v9 uint8 = 30
            _ = v9
        }
        _ = v6
        _ = v7
    }
    v5 += v5
    if (!("foo bar" != "")) {
        var v10 uint = 696
        v11 := GenS2{F0: 255, F1: 59, F2: false}
        fmt.Println(v5, v5)
        _ = v10
        _ = v11
    } else {
        fmt.Println(v5)
    }
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    var v3 int8 = 100
    if (false || (!(0 != 1))) {
        fmt.Println(v3, v3)
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 = v1
            var v13 int32 = 1
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            v3 += v3
            v15 := GenS1{F0: (0 <= 0)}
            v16 := []int{(fi14 - fi14), 9223372036854775807, fi14, (^(fi14))}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    switch ((("a" + "a") + "a") + "\n\t") {
    case "hello":
        v3 = (-(v1))
    }
    v17 := GenS2{F0: 42, F1: 42, F2: ((-412 <= 0) || true)}
    _ = len(v2)
    v18 := GenS1{F0: (true && true)}
    var v19 float64 = -2.25
    v3 = (v1 << ((297) & 63))
    fmt.Println(v19, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
