package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 uint64, p2 uint16) string {
    v6 := []int{-417}
    _ = len(v6)
    v7 := GenS1{F0: "\n\t"}
    v8 := GenS1{F0: ("hello" + ((("foo bar" + "\n\t") + "\n\t") + "hello"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 int64, p1 uint64) uint16 {
    fmt.Println(p0, p0)
    var v9 float32 = 0.0
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: "\n\t"}
        var v12 int64 = p0
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]int)
    var v5 uint16 = 746
    v2 += v2
    v2 = v2
    delete(v4, 2147483647)
    if (!false) {
        v5 += 32767
    }
    switch 1000 {
    case -9223372036854775808:
        switch 631 {
        case -9223372036854775808:
            var v13 float64 = 1.5
            _ = v13
        case 32767:
            for fi14 := 0; fi14 < 14; fi14++ {
                v3 -= v3
                _ = fi14
            }
        }
    default:
        fmt.Println(v5, v2, v5)
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int = fi15
        _ = fi15
        _ = v16
    }
    v3 += 917
    fmt.Println(v3, v5, v3)
    var v17 uint64 = 780
    v18 := GenS1{F0: ""}
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            switch "foo bar" {
            case "a":
                var v21 uint = 767
                _ = v21
            }
            var v22 int8 = 1
            v23 := GenS1{F0: ("a" + ("hello" + (("" + "foo bar") + ("a" + ""))))}
            _ = fi20
            _ = v22
            _ = v23
        }
        var v24 int16 = 1000
        v2 = (v2 ^ 4294967295)
        _ = fi19
        _ = v24
    }
    var v25 int16 = 665
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
