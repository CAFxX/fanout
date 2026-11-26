package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int32, p1 int) int8 {
    p1 = (^(p1))
    v6 := []int{(p1 << ((1) & 63)), p1, (^(p1))}
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 int64) int32 {
    fmt.Println(p0)
    switch 42 {
    case -1:
        if false {
            p0 -= (^(p0))
            fmt.Println(p0)
            v7 := []int{}
            _ = v7
        }
    }
    _ = p0
    return -656
}

func main() {
    var v1 uint = 42
    var v2 int8 = 42
    var v3 float64 = 0.0
    v4 := make(map[int]uint32)
    var v5 int8 = -73
    var v8 uint = 1
    v9 := GenS1{F0: -879, F1: -180}
    var v10 uint64 = 2
    if (0.0 != 1e10) {
        v4[32768] = 31
        for fi11 := 0; fi11 < 7; fi11++ {
            switch ((("a" + "\n\t") + "\n\t") + "hello") {
            case "a":
                for fi12 := 0; fi12 < 11; fi12++ {
                    v3 = (v3 / 3.14159)
                    _ = fi12
                }
            case "":
                v13 := GenS2{F0: ((!(1.5 < (0.1 / -2.25))) && true), F1: (v10 + 42), F2: -32768}
                _ = v13
            case "\n\t":
                var v14 string = "hello"
                _ = v14
            }
            if (float32(0.1) > 3.14159) {
                v15 := GenS3{F0: 681, F1: fi11}
                fmt.Println(fi11, v3, v5)
                _ = v15
            }
            if true {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v8, fi11)
                }
            }
            _ = fi11
        }
    }
    if true {
        v16 := []int{}
        switch -746 {
        case -402:
            if true {
                if false {
                    var v17 int64 = -277
                    var v18 int = 0
                    _ = v17
                    _ = v18
                }
            } else {
                fmt.Println(v3, v5, v8)
                v5 -= (v2 & -15)
            }
        case 944:
            v1 -= 0
        }
        _ = v16
    }
    v1 = uint(v10)
    fmt.Println(v3, v3, v2)
    if ((v3 / 1.0) <= (0.0 * v3)) {
        v19 := []int{-747, 32767, -999}
        fmt.Println(v2, v1, v5)
        _ = len(v19)
        _ = v19
    }
    v3 -= (float64(3.14159) - v3)
    var v20 float32 = float32(3.14159)
    fmt.Println(v8, v3, v8)
    var v21 float32 = float32(0.0)
    v22 := v4[2147483647]
    v21 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
