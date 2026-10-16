package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int16 = 1000
    var v2 string = (("foo bar" + "") + "\n\t")
    var v3 int32 = 42
    if false {
        for fi4 := 0; fi4 < 6; fi4++ {
            v3 = v3
            v2 = "hello"
            _ = fi4
        }
        v5 := []int{0, 127}
        _ = v5
    } else {
        if false {
            v3 = v3
        }
        v6 := []int{-9223372036854775808, 128, 32767}
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        switch -361 {
        case 0:
            var v8 int16 = -724
            _ = v8
        case -1:
            v3 = v3
        case 161:
            var v9 int16 = v1
            _ = v9
        default:
            v1 += (v1 * v1)
        }
        if false {
            v10 := GenS1{F0: v3}
            v2 = v2
            fmt.Println(fi7)
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v2, v3)
    v2 = v2
    var v11 int16 = (v1 | -495)
    if (32767 == 1) {
        var v12 uint64 = 0
        v2 = (v2 + v2)
        v13 := []int{0, 42}
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: (^(v3))}
        fmt.Println(v11, v3, v1)
        _ = v14
    }
    var v15 uint32 = 881
    var v16 float32 = (float32(1.5) / 1.0)
    switch int(v11) {
    case 829:
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := []int{32767, 4294967295}
            v19 := GenS2{F0: 604}
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        var v20 int8 = -32
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
