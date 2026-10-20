package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
func genFn0(p0 uint32, p1 uint8) uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{fi7, fi7, (fi7 | 1), (fi7 / (((fi7 << ((4294967295) & 63))) | 1))}
        v9 := GenS1{F0: (fi7 % (((fi7 + (fi7 >> ((2) & 63)))) | 1))}
        p0 += p0
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        if false {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: (^(fi10))}
                var v12 float64 = 3.14159
                var v13 uint16 = 2
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            switch 128 {
            case 42:
                v14 := []int{}
                _ = v14
            case 128:
                if true {
                    var v15 float64 = ((-2.25 + float64(float32(-2.25))) * float64(3.14159))
                    _ = v15
                }
            case 195:
                var v16 float64 = -2.25
                _ = v16
            }
            v17 := GenS1{F0: 4294967295}
            _ = v17
        } else {
            if false {
                p1 -= (p1 >> ((1) & 63))
            } else {
                switch -903 {
                case -9223372036854775808:
                    v18 := GenS1{F0: 2147483647}
                    _ = v18
                case 2147483647:
                    fmt.Println(p1, p0)
                case 9223372036854775807:
                    var v19 int16 = 598
                    _ = v19
                }
                switch 32768 {
                case 32768:
                    var v20 int = -112
                    _ = v20
                case 0:
                    var v21 uint16 = 32767
                    _ = v21
                case -9223372036854775808:
                    p1 += p1
                }
            }
            var v22 float32 = 0.1
            _ = v22
        }
        v23 := []int{}
        p0 -= 42
        _ = v23
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint16 = 0
            _ = fi24
            _ = v25
        }
        v26 := GenS1{F0: -690}
        v27 := []int{128, 42, 2147483647}
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1(p0 int32, p1 uint32) uint16 {
    if (2 != 0) {
        var v28 int32 = p0
        var v29 uint16 = 2
        _ = v28
        _ = v29
    }
    p0 = 1
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (3.14159 - 1.5)
    v3 := make(map[int]bool)
    var v4 uint8 = 2
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    var v30 int32 = 268
    v4 = v4
    var v31 uint = 4294967295
    for fi32 := 0; fi32 < 1; fi32++ {
        v33 := []int{(fi32 - (fi32 & 42)), fi32}
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 uint16 = 42
            v36 := GenS2{F0: (v4 & (v4 - 42)), F1: (!(true && (v35 == v35))), F2: "foo bar"}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v2)
    var v37 int = -22
    var v38 uint8 = (v4 - v4)
    var v39 float64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
