package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 312
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    v6 := []int{607, -527}
    v7 := GenS1{F0: -544, F1: (v1 & 42)}
    _ = len(v6)
    var v8 int8 = 100
    switch ("\n\t" + "a") {
    case "foo bar":
        var v9 uint = 255
        _ = v9
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                v1 = (^(v8))
                _ = fi11
            }
            _ = fi10
        }
    }
    var v12 int32 = 1
    delete(v2, (v3 ^ v3))
    var v13 int16 = -1
    if false {
        fmt.Println(v12, v8, v12)
        fmt.Println(v8)
    }
    v1 += (v8 << ((656) & 63))
    var v14 int32 = -153
    v3 = v3
    for fi15 := 0; fi15 < 5; fi15++ {
        v8 = 1
        if true {
            v13 = (v13 * (v13 << ((550) & 63)))
            switch "\n\t" {
            case "":
                v16 := GenS2{F0: 0, F1: ("hello" + ("a" + ("a" + ("a" + (("\n\t" + ("\n\t" + "\n\t")) + ("hello" + (("a" + "") + ("hello" + ("foo bar" + (("\n\t" + "a") + ""))))))))))}
                _ = v16
            }
            fmt.Println(v13, v1)
        } else {
            if ((42 > 255) && (fi15 == (^(fi15)))) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
