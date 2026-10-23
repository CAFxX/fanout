package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 uint16) int64 {
    if (2 >= 127) {
        p1 -= p1
        for fi4 := 0; fi4 < 15; fi4++ {
            p0 = (p0 - p0)
            p1 = p1
            _ = fi4
        }
    }
    p1 -= (p1 & 42)
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 int64 = 22
    v2 := make(map[string]uint32)
    var v3 string = "foo bar"
    if (2 == 255) {
        if true {
            v5 := []int{-1, 975, 9223372036854775807}
            _ = v5
        } else {
            v1 = 0
            v6 := []int{-785, -1, -928, -1}
            _ = v6
        }
        switch -1 {
        case -646:
            var v7 string = v3
            _ = v7
        }
    }
    v8 := GenS1{F0: 2147483647}
    v3 = ((v3 + v3) + "a")
    fmt.Println(v1, v3)
    switch -851 {
    case -991:
        v3 = v3
    }
    v1 -= v1
    if false {
        switch 0 {
        case 32767:
            var v9 uint16 = 661
            _ = v9
        case 127:
            for fi10 := 0; fi10 < 16; fi10++ {
                v3 = (v3 + "foo bar")
                fmt.Println(v1, v3)
                _ = fi10
            }
        case -773:
            if false {
                switch v3 {
                case "foo bar":
                    var v11 uint32 = 571
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v1 = (v1 + genFn0(-1, 0))
                        v1 = (v1 >> ((632) & 63))
                        _ = fi12
                    }
                }
            } else {
                v13 := []int{42, int(v1), -442}
                for fi14 := 0; fi14 < 3; fi14++ {
                    v1 = v1
                    if (1e10 == (1.5 + ((1.5 + (3.14159 - 1e10)) + (-2.25 * (1.5 * -2.25))))) {
                        var v15 uint16 = 442
                        if (252 >= 2147483647) {
                            v16 := GenS1{F0: -2147483648}
                            _ = v16
                        }
                        _ = v15
                    }
                    _ = fi14
                }
                _ = v13
            }
        default:
            var v17 float64 = (1e10 - (3.14159 + (0.0 + 3.14159)))
            _ = v17
        }
        v3 = ("a" + v3)
        var v18 float64 = 0.1
        _ = v18
    } else {
        v19 := v2[v3]
        var v20 int = 42
        _ = v19
        _ = v20
    }
    var v21 bool = true
    for fi22 := 0; fi22 < 11; fi22++ {
        if v21 {
            v21 = false
            v3 = v3
        } else {
            fmt.Println(v1, v3, v3)
            var v23 float64 = -2.25
            _ = v23
        }
        _ = fi22
    }
    switch 1 {
    case -887:
        v2["a"] = 300
    case -275:
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := []int{fi24, fi24}
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v8
}
