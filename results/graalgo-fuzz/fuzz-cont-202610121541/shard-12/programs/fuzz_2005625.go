package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
func genFn0(p0 uint8) int16 {
    v6 := GenS1{F0: 407, F1: 986}
    if false {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 string = (("foo bar" + ("foo bar" + (("foo bar" + (("a" + "") + "\n\t")) + "foo bar"))) + ((("\n\t" + ("a" + "\n\t")) + "\n\t") + "a"))
            v9 := []int{(fi7 >> ((255) & 63)), fi7, 0}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        switch "" {
        case "":
            p0 -= (p0 - (p0 % ((p0) | 1)))
        case "hello":
            fmt.Println(p0)
        default:
            p0 += p0
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                p0 = (p0 << ((812) & 63))
                var v12 uint32 = 0
                fmt.Println(fi11, v12)
                _ = fi11
                _ = v12
            }
            if (42 < 32767) {
                p0 -= p0
            } else {
                if true {
                    fmt.Println(p0, fi10)
                    var v13 int32 = 2147483647
                    _ = v13
                }
                p0 = 127
                var v14 int64 = -1
                _ = v14
            }
            _ = fi10
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -973
}
func genFn1(p0 int8) int16 {
    v15 := []int{127, 2147483647, 32768, 4294967295}
    v15 = append(v15, 4294967295)
    if (!false) {
        var v16 int8 = (p0 / ((p0) | 1))
        var v17 int = 104
        if ((42 == -1) || false) {
            if len(v15) > 0 {
                v15[0] = -1
            }
            if false {
                p0 -= p0
                var v18 int16 = 42
                v19 := GenS2{F0: 4294967295, F1: 128, F2: 461}
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                p0 -= (p0 + (p0 & v16))
                var v21 uint16 = 2
                _ = fi20
                _ = v21
            }
        } else {
            fmt.Println(v17, p0)
            p0 -= (v16 | p0)
        }
        _ = v16
        _ = v17
    }
    p0 = p0
    var v22 int16 = -708
    _ = p0
    _ = v15
    _ = v22
    return v22
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]int)
    var v3 float64 = (3.14159 + 1e10)
    var v4 float32 = 1.5
    var v5 uint16 = 0
    var v23 float64 = (v3 / ((v3 + v3) / 1e10))
    _ = len(v2)
    fmt.Println(v3, v23)
    var v24 float32 = float32((-2.25 / -2.25))
    v5 += v5
    v25 := GenS3{F0: 2, F1: "foo bar", F2: v1}
    if false {
        var v26 uint16 = 32767
        v26 -= v5
        _ = v26
    }
    _ = len(v2)
    if (1e10 > float64((v4 - v24))) {
        fmt.Println(v24, v1)
        v24 -= (1e10 * v4)
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                v4 -= v4
                fmt.Println(v4)
                fmt.Println(v3)
                _ = fi28
            }
            _ = fi27
        }
    }
    v29 := v2[42]
    v1 = 1
    fmt.Println(v23)
    if ("foo bar" > ("" + "")) {
        v3 -= v23
        for fi30 := 0; fi30 < 19; fi30++ {
            v1 = 67
            v5 -= 255
            _ = fi30
        }
        if ((0.0 * (((3.14159 - v23) * (0.1 + (float64(v24) / v23))) + 0.1)) == (((((-2.25 * float64((v24 / 1.0))) + v3) + v3) * v3) - 1e10)) {
            var v31 bool = (genFn1(v1) <= 1000)
            var v32 int16 = -108
            fmt.Println(v3, v1)
            _ = v31
            _ = v32
        }
    }
    v1 = (v1 + (v1 & int8(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
