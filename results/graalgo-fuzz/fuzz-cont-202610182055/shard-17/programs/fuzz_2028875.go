package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }

func main() {
    var v1 int16 = -475
    var v2 int16 = v1
    v3 := []string{""}
    var v4 uint8 = 127
    var v5 int16 = v2
    if false {
        if (2147483648 < 669) {
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "a")
            }
            v6 := GenS1{F0: false, F1: 620, F2: 18446744073709551615}
            _ = v6
        }
        var v7 uint64 = 425
        fmt.Println(v5)
        _ = v7
    }
    var v8 int64 = -1
    var v9 uint32 = 4294967295
    var v10 int16 = 391
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 -= 42
    switch 9223372036854775807 {
    case -73:
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                var v12 int32 = 42
                _ = v12
            }
            var v13 int32 = int32(v2)
            _ = fi11
            _ = v13
        }
    case -2147483648:
        var v14 int8 = 100
        _ = v14
    default:
        if true {
            var v15 uint16 = 65535
            var v16 string = (("hello" + ("hello" + ("foo bar" + ""))) + "hello")
            _ = v15
            _ = v16
        } else {
            fmt.Println(v10)
        }
    }
    var v17 int64 = (v8 << ((260) & 63))
    fmt.Println(v10, v10)
    _ = len(v3)
    switch 1 {
    case 42:
        var v18 float32 = 3.14159
        _ = v18
    case 453:
        v19 := GenS1{F0: (!((v17 | 9223372036854775807) < (v8 & (-(v17))))), F1: v9, F2: 42}
        _ = v19
    case 2147483647:
        switch -934 {
        case 9223372036854775807:
            fmt.Println(v2)
        case 32767:
            var v20 int32 = 1
            _ = v20
        case 0:
            if len(v3) > 0 {
                v3[0] = ("" + "foo bar")
            }
        }
    }
    fmt.Println(v5, v8)
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 int32 = -1
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 string = ((("foo bar" + ("a" + ("" + ("hello" + ("foo bar" + "a"))))) + "foo bar") + ("" + ("foo bar" + "\n\t")))
            if (1 > 32768) {
                var v25 uint32 = (v9 + 561)
                _ = v25
            }
            if (v8 >= (^(v8))) {
                v2 += (-(v1))
                v26 := GenS2{F0: v4, F1: (v4 - v4), F2: (v22 / ((v22) | 1))}
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
