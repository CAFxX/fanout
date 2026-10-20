package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    var v2 int16 = 0
    v3 := make(map[string]string)
    v4 := []int{32768}
    fmt.Println(v2)
    delete(v3, ("foo bar" + (("hello" + "") + "foo bar")))
    if false {
        if false {
            var v5 uint32 = 2147483647
            for fi6 := 0; fi6 < 19; fi6++ {
                v1 = v1
                v7 := GenS1{F0: 0, F1: 4294967295}
                fmt.Println(v2)
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        var v8 string = (("hello" + "\n\t") + ("foo bar" + "\n\t"))
        _ = v8
    } else {
        v2 = v2
    }
    var v9 uint8 = 42
    v10 := GenS1{F0: 255, F1: 0}
    fmt.Println(v2, v1, v9)
    for fi11 := 0; fi11 < 5; fi11++ {
        v1 = v1
        var v12 uint64 = 2
        _ = fi11
        _ = v12
    }
    if (true && ((!(!(0.0 < (0.0 + 3.14159)))) || false)) {
        var v13 float64 = (0.1 / float64((float32(v1) + 3.14159)))
        switch "hello" {
        case "":
            v1 = 3.14159
        }
        _ = v13
    }
    v9 = 164
    switch 9223372036854775807 {
    case 42:
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int64 = -1
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: 42, F1: 953}
                v18 := GenS2{F0: (-(v2))}
                fmt.Println(fi16, v2)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    case 0:
        fmt.Println(v9, v9)
    case -646:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
