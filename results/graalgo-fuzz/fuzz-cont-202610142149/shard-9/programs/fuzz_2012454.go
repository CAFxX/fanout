package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8) uint64 {
    if (859 < 1) {
        if (!(true && true)) {
            switch ("foo bar" + "") {
            case "a":
                p0 = p0
            case "foo bar":
                fmt.Println(p0)
            default:
                p0 = 42
            }
        }
    } else {
        v7 := GenS1{F0: 156}
        _ = v7
    }
    var v8 uint = 1
    v8 = 4294967295
    var v9 string = ("a" + "a")
    _ = p0
    _ = v8
    _ = v9
    return 847
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 0
    var v3 uint64 = 992
    var v4 uint = 895
    v5 := []int{812}
    var v6 uint = v4
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := v1[("hello" + "")]
        for fi12 := 0; fi12 < 16; fi12++ {
            v4 = (v6 >> (((v4 % ((v6) | 1))) & 63))
            fmt.Println(v3)
            if len(v5) > 0 {
                v5[0] = fi10
            }
            var v13 uint64 = (v3 << ((v6) & 63))
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v6, fi14)
        var v15 int = (fi14 >> (((v6 << ((358) & 63))) & 63))
        _ = fi14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS1{F0: 0}
            _ = fi16
            _ = v17
        }
        var v18 int16 = -32768
        _ = v18
    } else {
        var v19 int8 = 1
        v5 = append(v5, -2147483648)
        switch 42 {
        case -2147483648:
            v3 -= (^(v3))
        case 0:
            var v20 uint = v6
            _ = v20
        case 1:
            fmt.Println(v19, v3, v6)
        }
        _ = v19
    }
    v2 = v2
    fmt.Println(v4)
    v2 += 4294967295
    var v21 string = ""
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
