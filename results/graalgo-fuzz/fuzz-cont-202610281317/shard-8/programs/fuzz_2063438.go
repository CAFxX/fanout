package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint64, p2 string) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: (p1 >> ((65535) & 63)), F1: 42, F2: 42}
        _ = fi4
        _ = v5
    }
    var v6 uint64 = (p0 + (p0 ^ p0))
    if true {
        p0 += 1
    } else {
        v7 := []int{9223372036854775807, -2147483648, 42}
        fmt.Println(p0, p1)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (true && true)
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    switch 42 {
    case 2147483647:
        var v8 int = -395
        _ = v8
    case 526:
        if (42 == 321) {
            var v9 int16 = 32767
            v9 = (v9 | v9)
            v9 += v9
            _ = v9
        }
    default:
        switch 32768 {
        case -978:
            var v10 float64 = (-2.25 * -2.25)
            _ = v10
        case 4294967295:
            var v11 uint = 18446744073709551615
            _ = v11
        case 32768:
            switch -30 {
            case -737:
                var v12 uint = 255
                _ = v12
            case 4294967295:
                v13 := GenS1{F0: 2, F1: 87, F2: 42}
                _ = v13
            case -2147483648:
                var v14 int = 2147483647
                _ = v14
            default:
                var v15 uint16 = 255
                _ = v15
            }
        default:
            var v16 float64 = 0.1
            _ = v16
        }
    }
    var v17 string = (("\n\t" + ("hello" + "a")) + "hello")
    if true {
        var v18 float64 = 1.5
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                _ = len(v2)
                _ = fi20
            }
            v21 := GenS1{F0: 959, F1: 67, F2: 1}
            if true {
                fmt.Println(v18, v17, v17)
            }
            _ = fi19
            _ = v21
        }
        _ = v18
    } else {
        v22 := []int{-745, 201, -143, -704}
        for fi23 := 0; fi23 < 18; fi23++ {
            v17 = v17
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi23
        }
        _ = v22
    }
    fmt.Println(v17)
    switch 2147483647 {
    case 603:
        v17 = "\n\t"
    case -487:
        var v24 uint = 1
        _ = v24
    case -722:
        var v25 bool = (3.14159 < (0.1 - 1.5))
        _ = v25
    default:
        var v26 float32 = 3.14159
        _ = v26
    }
    v27 := GenS2{F0: 65535}
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint64 = 0
        var v30 uint16 = 2
        for fi31 := 0; fi31 < 6; fi31++ {
            fmt.Println(v30)
            _ = fi31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
}
