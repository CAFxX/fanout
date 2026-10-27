package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint64 = 0
    var v2 int8 = -66
    var v3 int8 = 0
    if true {
        if (127 != 128) {
            if false {
                switch 4294967295 {
                case 2147483647:
                    fmt.Println(v1, v1)
                case -896:
                    v4 := GenS1{F0: false, F1: 356}
                    _ = v4
                }
                var v5 uint32 = 0
                _ = v5
            }
            v3 += (^(v3))
        } else {
            v6 := GenS1{F0: true, F1: (v1 + (v1 ^ v1))}
            switch 127 {
            case -2147483648:
                var v7 int16 = 32767
                _ = v7
            default:
                if true {
                    v8 := GenS2{F0: 394}
                    _ = v8
                }
            }
            var v9 bool = (57 < 0)
            _ = v6
            _ = v9
        }
        if (!(4294967295 <= 42)) {
            fmt.Println(v2)
            var v10 string = ("\n\t" + ((("a" + ("foo bar" + "foo bar")) + "") + ("a" + ((("a" + "hello") + "hello") + ("\n\t" + ((("foo bar" + "a") + "a") + "hello"))))))
            v11 := []int{93, 42}
            _ = v10
            _ = v11
        } else {
            var v12 float64 = 1.5
            v13 := []int{-559, 343, 2147483647, -2147483648}
            _ = v12
            _ = v13
        }
        var v14 uint = 880
        _ = v14
    }
    v15 := GenS2{F0: 307}
    v2 = (v2 << ((65535) & 63))
    switch "foo bar" {
    case "":
        var v16 int32 = 501
        _ = v16
    }
    var v17 uint16 = 1
    var v18 int = -896
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(v18, v2, fi19)
        if (1.5 != float32(1e10)) {
            v3 = v3
            fmt.Println(v2)
            fmt.Println(v17, v3)
        }
        _ = fi19
    }
    v20 := GenS2{F0: 42}
    if (!false) {
        v3 = 100
        v21 := []int{v18, v18, v18}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
