package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 string, p1 int8) uint {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int64 = -9223372036854775808
        p1 += -128
        var v9 float64 = 1e10
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!(((p0 + p0) < p0) && (0.0 != float32((float64((0.1 + ((1e10 * (((0.1 + (3.14159 / (float32(0.1) * 1.5))) / -2.25) + 0.1)) - 1.5))) - 1.5))))) {
        v10 := []int{0}
        _ = v10
    } else {
        if (true && true) {
            var v11 float32 = 1e10
            v12 := GenS1{F0: 128, F1: 582}
            _ = v11
            _ = v12
        } else {
            fmt.Println(p0)
        }
    }
    switch -9223372036854775808 {
    case -314:
        var v13 int64 = 1
        _ = v13
    case -968:
        v14 := GenS1{F0: 2, F1: 1}
        _ = v14
    }
    v15 := []int{-2147483648, 2147483647, -533}
    switch -2147483648 {
    case 2147483647:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    case 127:
        v15 = append(v15, -2147483648)
    }
    _ = p0
    _ = p1
    _ = v15
    return 1
}
func genFn1(p0 int64, p1 int) string {
    var v16 int = p1
    for fi17 := 0; fi17 < 20; fi17++ {
        switch (-(p1)) {
        case 4294967295:
            fmt.Println(p0, fi17)
        case -2147483648:
            if true {
                fmt.Println(v16)
            }
        case 0:
            v18 := GenS2{F0: 0, F1: (("\n\t" + ("foo bar" + "a")) + "hello")}
            _ = v18
        }
        p1 += fi17
        switch 32768 {
        case 128:
            v19 := []int{-1, -2147483648, fi17}
            _ = v19
        case 32767:
            fmt.Println(v16, p1)
        }
        _ = fi17
    }
    v20 := GenS2{F0: 677, F1: ("foo bar" + "foo bar")}
    _ = p0
    _ = p1
    _ = v16
    _ = v20
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 0.0
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 string = "a"
    v6 := make(map[string]uint32)
    fmt.Println(v2, v5)
    for fi21 := 0; fi21 < 19; fi21++ {
        if false {
            v22 := v4[275]
            var v23 int16 = 1000
            var v24 int16 = -1
            _ = v22
            _ = v23
            _ = v24
        }
        v25 := []int{}
        _ = fi21
        _ = v25
    }
    fmt.Println(v2, v5)
    v26 := []int{589, 32768}
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS2{F0: genFn0("", 42), F1: (v5 + v5)}
        v2 -= (1e10 / (v2 / (v2 * v2)))
        var v29 uint = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS3{F0: 1}
    fmt.Println(v2, v5)
    _ = len(v26)
    v31 := GenS4{F0: (0 != 2), F1: 171}
    v32 := v3[-9223372036854775808]
    v2 = v2
    if len(v26) > 0 {
        v26[0] = 9223372036854775807
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
