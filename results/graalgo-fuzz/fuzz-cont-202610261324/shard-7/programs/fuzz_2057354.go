package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 int32 = 969
    var v2 int64 = -583
    v3 := make(map[int]int)
    v1 += (v1 * v1)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: (v2 >> ((1) & 63))}
        _ = fi4
        _ = v5
    }
    v2 = (^(v2))
    var v6 bool = ((v1 + (v1 + 42)) <= (v1 ^ v1))
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(v6, v6, fi7)
            v6 = (v2 != (v2 + (v2 + -9223372036854775808)))
            _ = fi8
        }
        var v9 int64 = (v2 / ((-1) | 1))
        _ = fi7
        _ = v9
    }
    if (9223372036854775807 <= -9223372036854775808) {
        v6 = true
        fmt.Println(v6)
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            _ = len(v3)
            _ = fi10
        }
        v11 := []int{127}
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(fi12, v1, v2)
            if true {
                var v13 int32 = (^(v1))
                switch "hello" {
                case "foo bar":
                    if (!v6) {
                        v14 := GenS2{F0: -108, F1: v6}
                        fmt.Println(v13, v1, v2)
                        var v15 int32 = v1
                        _ = v14
                        _ = v15
                    } else {
                        _ = len(v11)
                        if len(v11) > 0 {
                            v11[0] = -207
                        }
                    }
                }
                v6 = v6
                _ = v13
            }
            v2 = (-(v2))
            _ = fi12
        }
        _ = v11
    }
    var v16 int = -1
    v17 := GenS2{F0: 1, F1: v6}
    var v18 string = ("\n\t" + ("\n\t" + (("a" + "foo bar") + "a")))
    v6 = v6
    if v6 {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 bool = v6
            _ = fi19
            _ = v20
        }
        v1 += (v1 + -748)
        for fi21 := 0; fi21 < 12; fi21++ {
            fmt.Println(v2)
            v22 := []int{(v16 - v16), (v16 >> ((255) & 63)), v16}
            if v6 {
                var v23 uint = 59
                v22 = append(v22, v16)
                _ = v23
            }
            _ = fi21
            _ = v22
        }
    } else {
        switch v16 {
        case -624:
            if true {
                var v24 uint32 = 65535
                _ = v24
            }
        case 4:
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 string = v18
                var v27 uint32 = 92
                _ = fi25
                _ = v26
                _ = v27
            }
        case 32768:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int8 = 127
                fmt.Println(v16)
                var v30 uint8 = 127
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            v2 += (-(v2))
        }
        fmt.Println(v16)
        var v31 uint32 = 2147483648
        _ = v31
    }
    var v32 uint32 = 693
    if ((v16 << ((4294967295) & 63)) < v16) {
        fmt.Println(v18)
    } else {
        switch ("foo bar" + (v18 + v18)) {
        case "hello":
            for fi33 := 0; fi33 < 1; fi33++ {
                v2 -= (v2 >> ((598) & 63))
                _ = fi33
            }
        }
        var v34 int = v16
        _ = v34
    }
    var v35 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v6
}
