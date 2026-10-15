package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 uint32, p2 uint64) string {
    v5 := []int{p0, (p0 >> ((886) & 63))}
    p1 += p1
    p0 = p0
    switch (p0 & 4294967295) {
    case -9223372036854775808:
        v6 := GenS1{F0: 14, F1: -1}
        _ = v6
    case -1:
        v7 := GenS1{F0: (p0 * p0), F1: 0}
        _ = v7
    case 9223372036854775807:
        v8 := GenS1{F0: 4294967295, F1: -9223372036854775808}
        _ = v8
    }
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "hello"
}
func genFn1(p0 int32, p1 string) int8 {
    p0 = (p0 << ((18446744073709551615) & 63))
    var v9 int64 = 1
    _ = p0
    _ = p1
    _ = v9
    return 100
}

func main() {
    var v1 int32 = -54
    var v2 int16 = 619
    var v3 string = "\n\t"
    v4 := make(map[int]uint32)
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint32 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v1, v3, v2)
        if true {
            v3 = "hello"
            _ = len(v4)
            v13 := []int{(fi12 | -2147483648), (fi12 << ((255) & 63))}
            _ = v13
        }
        _ = fi12
    }
    v14 := GenS2{F0: 2147483647}
    var v15 string = (("foo bar" + ("foo bar" + "foo bar")) + v3)
    switch (("" + (v3 + v15)) + "foo bar") {
    case "hello":
        var v16 int = 42
        _ = v16
    }
    var v17 uint64 = 0
    if true {
        v18 := GenS2{F0: 0}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := []int{fi19, fi19}
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v15 = (((v3 + v3) + v15) + v3)
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int64 = -582
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            switch "hello" {
            case "":
                v2 = (v2 % ((v2) | 1))
            case "a":
                var v25 uint32 = 2147483647
                _ = v25
            }
            v3 = ((v15 + (v15 + v3)) + "hello")
            _ = fi24
        }
        _ = fi21
    }
    var v26 uint8 = 53
    var v27 bool = true
    if v27 {
        v26 = v26
        v2 -= v2
    } else {
        if (!(v17 == (v17 | v17))) {
            var v28 int16 = (v2 % ((32767) | 1))
            _ = len(v4)
            _ = v28
        }
        v29 := []int{0, 798, 228}
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        if (v27 && false) {
            for fi31 := 0; fi31 < 9; fi31++ {
                if false {
                    var v32 uint8 = v26
                    v27 = false
                    v33 := v4[fi30]
                    _ = v32
                    _ = v33
                } else {
                    var v34 float32 = float32((((1e10 + (-2.25 - ((3.14159 + float64(float32(1e10))) + 0.0))) / 1.5) * 0.0))
                    v35 := []int{-1, (fi30 % ((0) | 1)), fi31}
                    _ = v34
                    _ = v35
                }
                _ = fi31
            }
            v27 = v27
            for fi36 := 0; fi36 < 5; fi36++ {
                v3 = v3
                _ = fi36
            }
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            fmt.Println(v26)
            v17 = 4294967295
            if (fi37 >= fi30) {
                var v38 float64 = 0.0
                var v39 int64 = -9223372036854775808
                _ = v38
                _ = v39
            } else {
                v40 := GenS2{F0: 994}
                var v41 float64 = -2.25
                _ = v40
                _ = v41
            }
            _ = fi37
        }
        for fi42 := 0; fi42 < 9; fi42++ {
            v26 = v26
            v43 := GenS1{F0: (fi30 / ((2147483647) | 1)), F1: 450}
            _ = fi42
            _ = v43
        }
        _ = fi30
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 int16 = -32768
        delete(v4, fi44)
        _ = fi44
        _ = v45
    }
    var v46 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v46
}
