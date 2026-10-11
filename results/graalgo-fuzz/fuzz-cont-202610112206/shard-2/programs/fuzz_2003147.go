package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 uint8) uint8 {
    fmt.Println(p1)
    if (-9223372036854775808 < -732) {
        p1 -= (p1 | (p1 * (p1 | 0)))
        v7 := GenS1{F0: (true || ((32768 > 4294967295) || true)), F1: (p0 ^ p0), F2: 1}
        _ = v7
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1() int16 {
    var v8 uint32 = 2147483648
    v8 += (v8 % ((1) | 1))
    v8 = v8
    _ = v8
    return -561
}

func main() {
    var v1 int = -839
    v2 := make(map[string]int)
    var v3 int16 = 32767
    var v4 uint32 = 2147483648
    v5 := make(map[int]int)
    var v6 int16 = v3
    if (-2.25 != ((1e10 * -2.25) * 0.1)) {
        v5[(^(v1))] = 9223372036854775807
        if false {
            v9 := v5[(v1 & 127)]
            var v10 float64 = 3.14159
            v11 := []int{2147483647, -1}
            _ = v10
            _ = v11
            _ = v9
        }
        switch "" {
        case "":
            v4 = v4
        case "a":
            var v12 float32 = 0.1
            _ = v12
        default:
            v13 := []int{(v1 % ((v1) | 1)), (v1 - v1), v1, 978}
            _ = v13
        }
    } else {
        v6 = 1000
        v1 = (v1 ^ v1)
    }
    v14 := []int{9223372036854775807}
    v14 = append(v14, v1)
    var v15 int16 = 0
    if true {
        var v16 uint8 = genFn0(-1, 127)
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS2{F0: 4294967295}
            var v19 int8 = 42
            var v20 float32 = 1e10
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v16
    }
    fmt.Println(v3)
    v21 := v5[v1]
    v22 := GenS1{F0: false, F1: -325, F2: 42}
    if (17 <= 255) {
        v23 := GenS2{F0: 65535}
        var v24 int64 = 934
        _ = len(v2)
        _ = v23
        _ = v24
    } else {
        fmt.Println(v21)
        fmt.Println(v3, v21, v1)
    }
    var v25 string = ("" + (((("a" + ("\n\t" + "\n\t")) + "") + "\n\t") + (("hello" + "a") + "\n\t")))
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = len(v14)
    if true {
        if (((("" + v25) + "hello") + v25) != v25) {
            for fi26 := 0; fi26 < 18; fi26++ {
                v25 = ""
                v27 := GenS2{F0: 510}
                for fi28 := 0; fi28 < 2; fi28++ {
                    v6 += -32768
                    _ = fi28
                }
                _ = fi26
                _ = v27
            }
            switch v1 {
            case 127:
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 bool = false
                    v31 := v2[v25]
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case -9223372036854775808:
                v32 := GenS2{F0: 255}
                _ = v32
            }
            var v33 float32 = 0.1
            _ = v33
        } else {
            var v34 uint32 = v4
            var v35 uint32 = 426
            _ = v34
            _ = v35
        }
        var v36 uint8 = 2
        _ = len(v14)
        _ = v36
    } else {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            var v38 uint16 = 2
            _ = fi37
            _ = v38
        }
        if (-50 >= 1) {
            var v39 int32 = -1
            v6 -= genFn1()
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
