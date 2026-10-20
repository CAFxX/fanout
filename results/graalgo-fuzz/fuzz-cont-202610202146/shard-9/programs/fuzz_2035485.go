package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32, p1 int) int16 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        if false {
            if true {
                v7 := []int{fi6, p1}
                _ = v7
            } else {
                p0 = (-(p0))
            }
        }
        _ = fi6
    }
    fmt.Println(p1)
    p1 = 2147483647
    _ = p0
    _ = p1
    return 263
}
func genFn1(p0 int8, p1 uint16) int16 {
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = 2147483647
        v10 := []int{(-(fi8))}
        v9 = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = 3.14159
    for fi12 := 0; fi12 < 13; fi12++ {
        if (v11 > v11) {
            switch "hello" {
            case "a":
                var v13 int16 = 32767
                _ = v13
            }
            switch ("a" + "a") {
            case "a":
                var v14 uint8 = 0
                _ = v14
            default:
                for fi15 := 0; fi15 < 17; fi15++ {
                    if true {
                        v11 -= 0.0
                        p1 = (p1 ^ p1)
                    } else {
                        switch -9223372036854775808 {
                        case 349:
                            v16 := []int{fi12}
                            _ = v16
                        }
                        v17 := GenS1{F0: 255}
                        fmt.Println(p0, fi15, v11)
                        _ = v17
                    }
                    p1 = p1
                    _ = fi15
                }
            }
        } else {
            p1 += p1
            var v18 int64 = 42
            _ = v18
        }
        var v19 uint64 = 2
        _ = fi12
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    var v3 bool = (0.1 > 0.0)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    switch 0 {
    case 32768:
        if (30 > -85) {
            fmt.Println(v3)
            if (59 <= 42) {
                fmt.Println(v3, v2)
                if (!true) {
                    v2 -= -157
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    for fi21 := 0; fi21 < 1; fi21++ {
                        v22 := GenS1{F0: 809}
                        var v23 float64 = (0.0 * float64(-2.25))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 19; fi24++ {
                        switch "" {
                        case "hello":
                            v25 := GenS1{F0: 0}
                            _ = v25
                        case "foo bar":
                            v3 = (65535 > 0)
                        case "":
                            v3 = (127 != -56)
                        }
                        _ = fi24
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 int16 = -1
                        var v28 int32 = 2147483647
                        v29 := []int{fi26}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi20
                }
            } else {
                for fi30 := 0; fi30 < 11; fi30++ {
                    for fi31 := 0; fi31 < 9; fi31++ {
                        switch (("hello" + ((("hello" + (("a" + (((("foo bar" + "a") + ("hello" + "hello")) + "a") + ("foo bar" + ""))) + "a")) + "foo bar") + "\n\t")) + "\n\t") {
                        case "":
                            v32 := []int{(fi31 * (fi30 << ((255) & 63))), -9223372036854775808}
                            _ = v32
                        }
                        fmt.Println(v3, v2)
                        _ = fi31
                    }
                    _ = fi30
                }
                _ = len(v4)
            }
        }
    }
    switch -413 {
    case 4294967295:
        var v33 uint8 = 127
        _ = v33
    default:
        var v34 uint32 = 2
        _ = v34
    }
    v2 = v2
    var v35 float64 = (0.0 * float64(float32(1e10)))
    v36 := []int{42, 42}
    if v3 {
        var v37 bool = (0 >= 42)
        _ = v37
    } else {
        v3 = v3
        var v38 int64 = 0
        fmt.Println(v35, v38, v35)
        _ = v38
    }
    switch "\n\t" {
    case "hello":
        var v39 bool = false
        _ = v39
    case "a":
        v40 := GenS2{F0: 127}
        _ = v40
    }
    v41 := GenS2{F0: 226}
    _ = len(v1)
    var v42 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
