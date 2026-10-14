package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint16, p2 int32) int16 {
    p1 = 2
    v6 := GenS1{F0: 255, F1: (p0 % ((p0) | 1))}
    p2 -= -1
    var v7 int8 = 42
    switch ("a" + "foo bar") {
    case "a":
        p1 -= p1
    case "":
        v8 := GenS2{F0: (v7 - (v7 - 0))}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -388
}
func genFn1(p0 uint8, p1 bool) int {
    var v9 int64 = -9223372036854775808
    switch 4294967295 {
    case 42:
        v10 := GenS1{F0: 0, F1: 4294967295}
        _ = v10
    case 32767:
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                v9 = -23
                _ = fi12
            }
            var v13 string = "\n\t"
            _ = fi11
            _ = v13
        }
    case 32768:
        var v14 string = "foo bar"
        _ = v14
    }
    var v15 float32 = 3.14159
    fmt.Println(p1)
    for fi16 := 0; fi16 < 3; fi16++ {
        if (42 <= 42) {
            switch 32767 {
            case 403:
                var v17 uint64 = 2
                _ = v17
            default:
                v15 -= 0.0
            }
            p0 += p0
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint32 = 299
                v20 := []int{int(v19), 127, (fi18 + -1)}
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 int8 = 0
                _ = fi21
                _ = v22
            }
        }
        var v23 uint = 1
        v9 -= (v9 ^ v9)
        _ = fi16
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v9
    return -713
}

func main() {
    v1 := []uint32{2147483647}
    var v2 float32 = (-2.25 + ((3.14159 / (0.1 + 1e10)) + 1e10))
    v3 := []int{0, 1, 2147483647, 402}
    v4 := []int{4294967295}
    var v5 uint32 = 783
    v5 -= 360
    var v24 uint = 37
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 float64 = -2.25
        _ = fi25
        _ = v26
    }
    var v27 uint32 = 1
    fmt.Println(v5, v2, v24)
    var v28 int32 = -124
    var v29 string = ("" + (("\n\t" + ("" + "foo bar")) + "\n\t"))
    fmt.Println(v5)
    var v30 uint64 = 831
    v29 = (("" + v29) + (("hello" + v29) + ((v29 + ((((v29 + "a") + v29) + v29) + ((v29 + "") + "\n\t"))) + (v29 + v29))))
    var v31 int = 32767
    switch genFn1(127, (-9223372036854775808 <= 825)) {
    case -2147483648:
        var v32 int8 = -128
        _ = v32
    case -488:
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS2{F0: 118}
            _ = len(v3)
            var v35 int = genFn1(35, true)
            _ = fi33
            _ = v34
            _ = v35
        }
    case 4294967295:
        v36 := GenS2{F0: 42}
        _ = v36
    }
    var v37 uint = uint(v27)
    v5 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
