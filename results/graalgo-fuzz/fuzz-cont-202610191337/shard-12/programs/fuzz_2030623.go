package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8, p1 string) uint8 {
    v6 := GenS1{F0: 2, F1: (^(p0))}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi7, 1}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS2{F0: 18446744073709551615, F1: ((0.1 * 0.0) >= float64(-2.25))}
    fmt.Println(p0)
    switch (p1 + p1) {
    case "hello":
        if (0.0 <= -2.25) {
            switch 42 {
            case 2147483647:
                v11 := []int{305, 127}
                _ = v11
            case 9223372036854775807:
                v12 := []int{-2147483648, -1, 2147483647}
                _ = v12
            }
            p0 = 45
        }
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            p0 = (^(p0))
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return (p0 >> ((0) & 63))
}
func genFn1(p0 uint8, p1 int, p2 int64) uint64 {
    p2 += p2
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = genFn0((p0 / ((p0) | 1)), ((("hello" + (("foo bar" + "hello") + ("foo bar" + ("foo bar" + ("a" + (((("foo bar" + "hello") + "foo bar") + ("\n\t" + "hello")) + "")))))) + "foo bar") + ("a" + "foo bar")))
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    if (p1 > p1) {
        v17 := []int{(p1 % ((p1) | 1)), (p1 >> ((323) & 63))}
        p1 = p1
        var v18 float32 = 0.0
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := []int{9223372036854775807}
            fmt.Println(p1, p0)
            p0 = p0
            _ = fi19
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 float32 = (((float32(3.14159) * (1e10 + 1.5)) / 3.14159) / 1.0)
    var v2 float64 = 0.0
    var v3 int = -1
    v4 := make(map[string]int)
    var v5 int64 = -1
    var v21 uint8 = 1
    var v22 int64 = v5
    v23 := []int{v3}
    switch (v3 * (v3 / ((-9223372036854775808) | 1))) {
    case 32768:
        if (((-53 >= -38) && true) && (true || true)) {
            v24 := GenS1{F0: 1, F1: 0}
            if false {
                v2 = v2
                v25 := v4["hello"]
                _ = v25
            }
            _ = v24
        } else {
            fmt.Println(v3)
        }
    default:
        if true {
            var v26 uint = 1
            _ = v26
        } else {
            var v27 uint32 = 0
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v22 = 0
        _ = fi28
    }
    fmt.Println(v2, v21, v2)
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 int = fi29
        _ = fi29
        _ = v30
    }
    fmt.Println(v1, v22, v22)
    v3 += (^(v3))
    v31 := GenS2{F0: 0, F1: (float64(v1) <= (((v2 / 1.5) + (v2 * v2)) * v2))}
    for fi32 := 0; fi32 < 10; fi32++ {
        v3 -= v3
        v3 += v3
        v33 := GenS2{F0: 18446744073709551615, F1: true}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 uint32 = 2
        _ = fi34
        _ = v35
    }
    var v36 uint16 = 32767
    fmt.Println(v21, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
