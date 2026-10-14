package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16) int {
    p0 = 255
    if (!((0.0 - float64(1e10)) > 1.5)) {
        switch 0 {
        case 285:
            var v6 string = "\n\t"
            _ = v6
        case -57:
            fmt.Println(p0)
        default:
            v7 := []int{32768}
            _ = v7
        }
        var v8 uint16 = p0
        _ = v8
    }
    _ = p0
    return 32767
}
func genFn1(p0 uint16, p1 int64) uint32 {
    var v9 float32 = ((0.0 * 1.5) * -2.25)
    if false {
        if ((867 != 4294967295) || false) {
            p1 = (p1 >> ((46) & 63))
            v9 -= 1e10
        }
        var v10 uint32 = 42
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := []int{genFn0((^(p0))), 9223372036854775807, fi11}
            for fi13 := 0; fi13 < 10; fi13++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v14 := GenS1{F0: fi13}
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        fmt.Println(p0)
    }
    v15 := GenS2{F0: 127, F1: 255, F2: 412}
    if (1e10 > -2.25) {
        v16 := GenS2{F0: -1, F1: 114, F2: 618}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v9
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 int16 = -1
    var v3 string = "foo bar"
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    v17 := GenS1{F0: 127}
    v18 := GenS1{F0: 1}
    if (true && false) {
        for fi19 := 0; fi19 < 18; fi19++ {
            v1 -= (-(v1))
            _ = fi19
        }
        switch 4294967295 {
        case 0:
            v20 := GenS3{F0: 42, F1: 255, F2: 28}
            _ = v20
        default:
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := []int{(fi21 ^ (fi21 % (((fi21 - 4294967295)) | 1))), fi21, 128, (fi21 & (^(fi21)))}
                if len(v22) > 0 {
                    v22[0] = (fi21 / ((4294967295) | 1))
                }
                v2 = (v2 * v2)
                _ = fi21
                _ = v22
            }
        }
    } else {
        var v23 uint16 = 42
        _ = v23
    }
    v24 := GenS1{F0: 561}
    switch ((v3 + (("hello" + ((v3 + "") + v3)) + (((("hello" + "foo bar") + v3) + v3) + v3))) + v3) {
    case "foo bar":
        v25 := GenS1{F0: 1}
        _ = v25
    case "":
        v2 += (v2 & 1)
    case "a":
        var v26 int16 = (v2 + 1000)
        _ = v26
    }
    var v27 uint8 = 111
    v28 := GenS1{F0: -1}
    if (false && ((v3 + v3) == (("\n\t" + v3) + "foo bar"))) {
        var v29 bool = true
        _ = v29
    }
    v30 := GenS4{F0: 127, F1: 4294967295}
    var v31 int64 = 9223372036854775807
    fmt.Println(v27, v27)
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            var v34 uint32 = 4294967295
            _ = fi33
            _ = v34
        }
        v35 := []int{fi32}
        v1 -= int8(v2)
        _ = fi32
        _ = v35
    }
    if (false || false) {
        for fi36 := 0; fi36 < 3; fi36++ {
            v31 = (v31 << ((65535) & 63))
            v37 := []int{(fi36 / ((9223372036854775807) | 1)), (fi36 & genFn0(2)), -642}
            _ = fi36
            _ = v37
        }
        switch v3 {
        case "hello":
            fmt.Println(v2, v27, v2)
        case "\n\t":
            v2 = (v2 | (v2 + v2))
        case "a":
            fmt.Println(v3, v3, v3)
        }
    } else {
        var v38 string = ((v3 + v3) + ("hello" + "hello"))
        fmt.Println(v38, v31, v31)
        _ = v38
    }
    var v39 uint8 = 0
    var v40 int64 = 42
    v41 := GenS4{F0: (v39 % ((v39) | 1)), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
