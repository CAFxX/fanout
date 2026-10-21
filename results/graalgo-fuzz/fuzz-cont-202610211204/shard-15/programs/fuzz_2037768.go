package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 string) int64 {
    fmt.Println(p0, p1)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int64 = 42
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 -= 42
    p0 += (-(p0))
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint32, p1 uint, p2 uint8) uint64 {
    var v8 uint32 = p0
    v9 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    var v2 uint = 2
    v3 := make(map[string]string)
    var v4 float64 = 1.5
    for fi10 := 0; fi10 < 13; fi10++ {
        if v1 {
            fmt.Println(v4, v1)
            for fi11 := 0; fi11 < 6; fi11++ {
                v2 = (v2 & v2)
                v4 += ((float64(((1.5 / (3.14159 / 1e10)) * 1e10)) * -2.25) / v4)
                v4 = v4
                _ = fi11
            }
            var v12 int32 = -367
            _ = v12
        } else {
            if false {
                var v13 uint = (v2 ^ v2)
                fmt.Println(v13)
                _ = v13
            }
            var v14 float64 = 0.1
            v15 := v3["hello"]
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    if true {
        var v16 string = ("hello" + ("foo bar" + ("a" + "\n\t")))
        _ = v16
    } else {
        var v17 uint64 = 18446744073709551615
        switch 4294967295 {
        case 0:
            switch "\n\t" {
            case "\n\t":
                if v1 {
                    v1 = v1
                    var v18 int = 127
                    _ = v18
                }
            case "":
                v4 = 1e10
            default:
                var v19 int32 = -1
                _ = v19
            }
        case 32768:
            v20 := GenS1{F0: 42}
            _ = v20
        case 518:
            var v21 int32 = -127
            _ = v21
        }
        _ = v17
    }
    v2 = (v2 << (((v2 * 65535)) & 63))
    v2 = v2
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS1{F0: 255}
        for fi24 := 0; fi24 < 2; fi24++ {
            v4 -= 0.1
            v1 = v1
            _ = fi24
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v2 -= 2
            _ = fi25
        }
        _ = fi22
        _ = v23
    }
    v2 = 757
    _ = len(v3)
    v2 = (^(v2))
    if v1 {
        for fi26 := 0; fi26 < 14; fi26++ {
            v2 = v2
            for fi27 := 0; fi27 < 5; fi27++ {
                v4 = v4
                v4 = v4
                _ = fi27
            }
            var v28 uint8 = 83
            _ = fi26
            _ = v28
        }
        v3["foo bar"] = "foo bar"
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint8 = 127
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
