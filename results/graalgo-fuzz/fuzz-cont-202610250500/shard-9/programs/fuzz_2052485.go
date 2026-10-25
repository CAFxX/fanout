package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 44
    var v3 int32 = v2
    var v4 bool = false
    v2 = 1
    var v5 uint32 = 2147483648
    var v6 int16 = 42
    v2 = v3
    for fi7 := 0; fi7 < 9; fi7++ {
        v2 -= v2
        _ = fi7
    }
    var v8 string = "foo bar"
    delete(v1, 438)
    if (!(42 != 125)) {
        v9 := GenS1{F0: -128}
        fmt.Println(v8, v3, v2)
        switch -832 {
        case 32768:
            v10 := []int{}
            _ = v10
        default:
            v4 = (!(v4 && v4))
        }
        _ = v9
    } else {
        if v4 {
            delete(v1, 1)
        } else {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS2{F0: fi11, F1: (!v4), F2: ("hello" + v8)}
                v13 := GenS2{F0: fi11, F1: v4, F2: "hello"}
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(v2)
                _ = fi14
            }
            var v15 int16 = 132
            _ = v15
        }
        v3 -= v3
    }
    if v4 {
        v5 = 2147483648
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint = 0
            fmt.Println(fi16, v4)
            _ = len(v1)
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int16 = v6
            _ = fi18
            _ = v19
        }
    }
    switch 696 {
    case 127:
        v6 = (v6 ^ (v6 >> ((506) & 63)))
    case 42:
        var v20 int64 = 732
        _ = v20
    case -1:
        v8 = (("a" + (v8 + v8)) + (v8 + ((((v8 + v8) + v8) + v8) + ((("\n\t" + "a") + (v8 + v8)) + v8))))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
