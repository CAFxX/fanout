package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64, p1 uint8) uint16 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{-186, (fi5 + fi5)}
        var v7 float32 = (3.14159 - -2.25)
        if false {
            switch "" {
            case "hello":
                var v8 float32 = v7
                _ = v8
            case "a":
                var v9 uint = 4294967295
                _ = v9
            case "foo bar":
                v7 += float32(0.1)
            }
            var v10 uint = 4294967295
            _ = v10
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 -= (-(p0))
    v11 := GenS1{F0: -2147483648, F1: 0, F2: 127}
    v12 := GenS1{F0: 1, F1: (^(p1)), F2: 23}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 32767
}
func genFn1(p0 int64, p1 int) int {
    v13 := GenS2{F0: 65535, F1: (!false), F2: -841}
    var v14 uint = 0
    v15 := []int{p1, (p1 & p1), 972, p1}
    if len(v15) > 0 {
        v15[0] = p1
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return p1
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 1
    v3 := make(map[string]bool)
    var v4 int16 = 1000
    if (0.1 < ((-2.25 + (1.5 - (1.5 / (-2.25 / -2.25)))) + 1e10)) {
        var v16 float64 = 0.0
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(fi17)
            fmt.Println(fi17)
            fmt.Println(v16, v16)
            _ = fi17
        }
        _ = v16
    } else {
        v4 = v4
    }
    v2 -= (v2 - v2)
    v2 += (v2 * 1)
    var v18 uint64 = 4294967295
    if true {
        if ((!false) || false) {
            var v19 int8 = 42
            _ = v19
        }
        fmt.Println(v2, v18)
    } else {
        v20 := GenS1{F0: 187, F1: 220, F2: 91}
        switch (v2 << ((250) & 63)) {
        case 2147483647:
            if (!false) {
                v2 = -703
                var v21 int64 = -9223372036854775808
                fmt.Println(v18)
                _ = v21
            } else {
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v18)
                    _ = fi22
                }
            }
        }
        var v23 uint16 = 42
        _ = v20
        _ = v23
    }
    v24 := GenS3{F0: v4}
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 bool = (9223372036854775807 >= -1)
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := []int{}
            _ = fi27
            _ = v28
        }
        v18 = (v18 - v18)
        _ = fi25
        _ = v26
    }
    var v29 int64 = 9223372036854775807
    var v30 float32 = 0.1
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint = 0
        var v33 uint32 = 27
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
}
