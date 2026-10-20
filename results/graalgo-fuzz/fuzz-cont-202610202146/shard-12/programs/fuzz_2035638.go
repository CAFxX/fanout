package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 127
    var v3 int16 = 32767
    v4 := make(map[string]uint32)
    var v5 uint64 = 4294967295
    var v6 int8 = (v2 - int8(v1))
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v5)
        var v8 uint32 = 974
        _ = fi7
        _ = v8
    }
    var v9 string = (("" + "hello") + "foo bar")
    v10 := []int{9223372036854775807}
    switch -90 {
    case 42:
        v1 -= (v1 - (v1 + (v1 << ((255) & 63))))
    case 2147483647:
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: 9223372036854775807, F1: 128}
            if false {
                if ((469 > 32767) && true) {
                    delete(v4, v9)
                    switch ((v9 + v9) + (v9 + "")) {
                    case "foo bar":
                        var v13 uint32 = 2147483648
                        _ = v13
                    case "":
                        delete(v4, "hello")
                    case "hello":
                        var v14 uint16 = 42
                        _ = v14
                    default:
                        var v15 uint = 0
                        _ = v15
                    }
                } else {
                    v1 -= v1
                }
                v6 += 0
                fmt.Println(v2, v9, v1)
            } else {
                if true {
                    v16 := GenS2{F0: (v5 + v5), F1: (v1 | int32(v3)), F2: (v3 + (-(v3)))}
                    _ = v16
                }
            }
            v3 = -65
            _ = fi11
            _ = v12
        }
    case 9223372036854775807:
        fmt.Println(v3, v2, v5)
    }
    fmt.Println(v1)
    var v17 string = v9
    fmt.Println(v2, v9)
    if (false && false) {
        v2 = (v6 >> ((255) & 63))
    }
    var v18 int = 128
    fmt.Println(v2)
    switch v18 {
    case 0:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case -2147483648:
        switch (v9 + ("hello" + (v9 + "foo bar"))) {
        case "hello":
            v9 = v9
        }
    }
    v9 = (v9 + v17)
    if false {
        fmt.Println(v2, v1, v6)
        var v19 float32 = 0.1
        v17 = v17
        _ = v19
    }
    v17 = v17
    v9 = ((v17 + ("foo bar" + "")) + v17)
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
