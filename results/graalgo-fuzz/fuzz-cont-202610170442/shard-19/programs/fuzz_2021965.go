package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0() int64 {
    var v7 string = "a"
    if true {
        v8 := []int{128, 32767, -9223372036854775808, -659}
        _ = v8
    }
    var v9 string = "hello"
    v9 = v7
    v10 := GenS1{F0: 42, F1: 42, F2: 2}
    _ = v10
    _ = v7
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := []int{538, -9223372036854775808, 127}
    var v4 int = -9223372036854775808
    var v5 int64 = -1
    var v6 int16 = -925
    var v11 float64 = 3.14159
    switch v4 {
    case 852:
        var v12 int8 = 42
        _ = v12
    case 279:
        switch (("hello" + "hello") + "\n\t") {
        case "a":
            switch (v4 / (((v4 % ((v4) | 1))) | 1)) {
            case 4294967295:
                var v13 int64 = (-(v5))
                _ = v13
            default:
                var v14 string = "foo bar"
                _ = v14
            }
        default:
            fmt.Println(v5, v6)
        }
    default:
        switch (v4 & v4) {
        case -714:
            v11 = (0.0 * v11)
        case 233:
            switch "a" {
            case "\n\t":
                var v15 int16 = v6
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v6 = 1000
        _ = fi16
    }
    v6 -= v6
    switch v4 {
    case -484:
        var v17 int8 = 42
        _ = v17
    case -2147483648:
        v4 = v4
    case -579:
        for fi18 := 0; fi18 < 1; fi18++ {
            delete(v1, (fi18 << ((0) & 63)))
            _ = len(v1)
            _ = len(v1)
            _ = fi18
        }
    }
    v19 := GenS2{F0: -2147483648, F1: 65535}
    var v20 uint8 = 58
    v3 = append(v3, -384)
    var v21 uint32 = 65535
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS2{F0: 1, F1: 42}
        _ = fi22
        _ = v23
    }
    v21 += v21
    v20 = v20
    for fi24 := 0; fi24 < 10; fi24++ {
        if (0.1 == ((0.0 + -2.25) + (((-2.25 + (0.0 / 0.1)) + 1e10) / ((float32(1e10) / 1e10) + ((3.14159 - float32(v11)) - 0.0))))) {
            var v25 int8 = 127
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint8 = 255
            if true {
                _ = len(v3)
                var v28 bool = (!true)
                if v28 {
                    var v29 uint16 = 0
                    if true {
                        if (!(v28 || (v28 && (!(!(v4 != v4)))))) {
                            var v30 int32 = 607
                            var v31 int8 = 100
                            var v32 int = v4
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            var v33 uint32 = v21
                            v34 := GenS2{F0: 1, F1: 4294967295}
                            fmt.Println(fi24, v29)
                            _ = v33
                            _ = v34
                        }
                        var v35 int8 = 100
                        var v36 int64 = genFn0()
                        fmt.Println(v21, v11)
                        _ = len(v1)
                        var v37 int64 = v5
                        fmt.Println(v29, v37, v4)
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    var v38 uint32 = (v21 | (v21 & (v21 >> ((uint(v20)) & 63))))
                    _ = v29
                    _ = v38
                }
                _ = v28
            }
            _ = fi26
            _ = v27
        }
        _ = fi24
    }
    var v39 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
