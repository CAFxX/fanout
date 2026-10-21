package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0() bool {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 int64 = 1
    var v8 int64 = v7
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint = 65535
        v8 += (v7 + (v7 >> ((v11) & 63)))
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int8 = -38
            v14 := GenS1{F0: v11, F1: 128, F2: 225}
            v15 := []int{(-(fi12))}
            fmt.Println(v13)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return (("" + ("\n\t" + "foo bar")) != ("foo bar" + ("\n\t" + "hello")))
}
func genFn1(p0 string, p1 uint16) int8 {
    var v16 float64 = (1.5 - -2.25)
    p1 = 0
    fmt.Println(v16, p0)
    _ = p0
    _ = p1
    _ = v16
    return 61
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 uint8 = 60
    var v5 uint32 = 4294967295
    _ = len(v3)
    v17 := GenS1{F0: 639, F1: (^(v4)), F2: (v4 - (v4 * (v4 - (^(v4)))))}
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            v5 -= 4294967295
            _ = fi19
        }
        v5 = 42
        _ = fi18
    }
    v20 := GenS2{F0: -1, F1: 100, F2: 255}
    delete(v3, 2147483648)
    var v21 int16 = -1
    for fi22 := 0; fi22 < 10; fi22++ {
        v21 -= (^(v21))
        _ = fi22
    }
    var v23 float64 = 3.14159
    var v24 bool = true
    switch "hello" {
    case "hello":
        v23 -= 3.14159
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
