package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64, p1 bool) uint32 {
    switch "" {
    case "":
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 255}
            switch 2147483647 {
            case 0:
                v9 := []int{601, (fi7 * -676), (fi7 + -1), (fi7 & 32767)}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    case "hello":
        if true {
            v10 := []int{32767, 4294967295, 143, 127}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                p1 = p1
                fmt.Println(p1, p0)
                var v12 int64 = 42
                _ = fi11
                _ = v12
            }
        }
    }
    switch 128 {
    case 127:
        v13 := []int{}
        _ = v13
    default:
        v14 := []int{519, 128, 128}
        _ = v14
    }
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 42
    var v3 bool = (!(1 <= 18446744073709551615))
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]bool)
    var v6 uint = 2
    v15 := GenS2{F0: (^(v6))}
    fmt.Println(v6, v3, v3)
    v6 = v6
    for fi16 := 0; fi16 < 5; fi16++ {
        v6 = (v6 / (((v6 << (((v6 << (((v6 - (^(v6)))) & 63))) & 63))) | 1))
        _ = fi16
    }
    fmt.Println(v6, v3)
    var v17 float64 = 1e10
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int64 = (v2 | v2)
        fmt.Println(v3, v2)
        var v20 int32 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 bool = true
    var v22 int64 = v2
    var v23 int32 = -1
    v22 += 42
    v23 -= -360
    for fi24 := 0; fi24 < 6; fi24++ {
        v17 = 1e10
        _ = fi24
    }
    v5[1] = (0.1 > v17)
    v25 := []int{720, 452, -297, 32768}
    if v21 {
        switch ("foo bar" + ((("a" + ((((("" + ("hello" + ((("\n\t" + "a") + (("" + ("hello" + ((("hello" + "\n\t") + "a") + "hello"))) + "")) + "hello"))) + "hello") + ("" + "\n\t")) + "") + "\n\t")) + (("a" + ((("a" + "foo bar") + (("\n\t" + "foo bar") + "hello")) + "foo bar")) + "\n\t")) + "hello")) {
        case "hello":
            v22 = v22
        case "\n\t":
            v26 := GenS3{F0: -2147483648}
            _ = v26
        }
        v27 := GenS1{F0: 2}
        v28 := GenS4{F0: -1}
        _ = v27
        _ = v28
    }
    v17 = float64(((-2.25 / -2.25) + -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
