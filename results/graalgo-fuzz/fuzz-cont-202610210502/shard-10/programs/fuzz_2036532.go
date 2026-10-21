package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 int, p1 uint64) int16 {
    v5 := []int{(-(p0)), -9223372036854775808}
    v5 = append(v5, int(p1))
    _ = p0
    _ = p1
    _ = v5
    return 1000
}

func main() {
    var v1 int64 = 312
    var v2 uint32 = 2
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v6 bool = (-1 > 42)
    v1 -= (v1 << ((592) & 63))
    if (-2147483648 < -190) {
        v7 := GenS1{F0: (!(255 != 255))}
        v1 -= v1
        _ = v7
    }
    var v8 int32 = 568
    v6 = true
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int32 = -366
            fmt.Println(v6, v1)
            _ = fi10
            _ = v11
        }
        fmt.Println(v8, fi9)
        _ = fi9
    }
    v2 += 2147483647
    fmt.Println(v8, v2)
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{fi12, 0, fi12}
        fmt.Println(v6, fi12)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint64 = 0
        v1 += 0
        var v16 int = -2147483648
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: (false || v6)}
        fmt.Println(v2)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v2 = 2147483648
        for fi20 := 0; fi20 < 8; fi20++ {
            if (0 < 361) {
                var v21 float64 = -2.25
                v4[fi20] = (fi19 << ((0) & 63))
                v22 := GenS2{F0: -17}
                v21 = (((v21 + -2.25) / v21) / (1.5 + float64((0.1 + ((0.1 / 0.1) - 0.0)))))
                v23 := GenS2{F0: 127}
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        _ = fi19
    }
    v8 -= (v8 / ((v8) | 1))
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS3{F0: true, F1: "foo bar"}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint16 = 255
        _ = len(v4)
        _ = fi26
        _ = v27
    }
    if (32767 <= 948) {
        if (("foo bar" + "foo bar") <= ("hello" + "")) {
            if (65535 <= 65535) {
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 uint8 = 0
                    v1 = (v1 & v1)
                    v29 = v29
                    _ = fi28
                    _ = v29
                }
                var v30 float64 = 3.14159
                v31 := []int{1, -2147483648, 1}
                _ = v30
                _ = v31
            }
        } else {
            fmt.Println(v6)
            if v6 {
                var v32 bool = (v1 >= (v1 + v1))
                var v33 string = "a"
                _ = v32
                _ = v33
            }
        }
        v34 := []int{929, 163}
        fmt.Println(v8, v2, v8)
        _ = v34
    } else {
        var v35 float64 = 1.5
        _ = v35
    }
    v36 := v4[0]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v6
    _ = v8
}
