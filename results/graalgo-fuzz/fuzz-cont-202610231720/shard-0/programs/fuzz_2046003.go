package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint16 }
func genFn0(p0 bool) uint32 {
    p0 = (!(-561 < 32768))
    var v4 uint32 = 610
    v4 = (v4 * v4)
    p0 = p0
    _ = p0
    _ = v4
    return v4
}
func genFn1(p0 uint32, p1 uint32) int64 {
    p0 = genFn0((!(false && true)))
    v5 := []int{}
    v6 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float64 = (1.5 - 0.0)
    var v2 int = 128
    v3 := make(map[int]int)
    if ((!(false || (1.5 <= ((1e10 / (((1e10 + 0.0) * 1.5) - float32((0.0 * v1)))) * 3.14159)))) && (-469 > genFn1(0, 2))) {
        v7 := GenS1{F0: 2}
        if false {
            v1 = v1
            v8 := GenS1{F0: 2}
            v2 = (v2 ^ (v2 ^ v2))
            _ = v8
        }
        var v9 uint16 = 255
        _ = v7
        _ = v9
    } else {
        v10 := GenS1{F0: 255}
        switch (v2 / ((v2) | 1)) {
        case -750:
            var v11 bool = (!false)
            _ = v11
        case -1:
            switch (v2 | (v2 >> ((4294967295) & 63))) {
            case -220:
                var v12 int32 = 0
                _ = v12
            case 128:
                v1 = (0.1 + v1)
            default:
                var v13 int16 = 1000
                _ = v13
            }
        case -418:
            switch "a" {
            case "hello":
                v1 = 0.0
            case "":
                var v14 uint32 = 957
                _ = v14
            default:
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := v3[fi15]
                    v17 := GenS1{F0: 39}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        default:
            for fi18 := 0; fi18 < 5; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    fmt.Println(fi19)
                    _ = fi19
                }
                v20 := GenS1{F0: 42}
                var v21 uint8 = 100
                _ = fi18
                _ = v20
                _ = v21
            }
        }
        fmt.Println(v2, v2)
        _ = v10
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(v2, fi22, v2)
        v1 -= 1.5
        if ((true || true) || false) {
            fmt.Println(fi22, v1)
        }
        _ = fi22
    }
    if (!false) {
        delete(v3, -2147483648)
        var v23 bool = true
        _ = v23
    } else {
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS1{F0: 127}
            var v26 int32 = -448
            v1 -= 1.5
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 float64 = (v1 * v1)
        if true {
            var v28 int16 = 32767
            switch "a" {
            case "":
                var v29 string = "foo bar"
                _ = v29
            case "a":
                switch "a" {
                case "\n\t":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 uint8 = 1
                        switch "" {
                        case "":
                            v27 -= v27
                        case "foo bar":
                            if ("" > "a") {
                                if true {
                                    v27 = 0.1
                                }
                                v31 = (v31 | 0)
                                var v32 int = fi30
                                _ = v32
                            } else {
                                fmt.Println(v27)
                            }
                        case "\n\t":
                            v33 := GenS1{F0: v31}
                            _ = v33
                        }
                        v2 += v2
                        _ = fi30
                        _ = v31
                    }
                }
            }
            var v34 int = -728
            _ = v28
            _ = v34
        } else {
            fmt.Println(v27)
        }
        _ = v27
    }
    if (true || (v2 == v2)) {
        for fi35 := 0; fi35 < 10; fi35++ {
            v36 := []int{(fi35 & 2147483647)}
            _ = fi35
            _ = v36
        }
        v37 := GenS2{F0: (v2 + v2)}
        _ = v37
    } else {
        switch (-(v2)) {
        case 2147483647:
            v38 := []int{(^(v2)), (v2 + (^(v2)))}
            _ = v38
        default:
            if false {
                v2 += v2
            } else {
                v39 := GenS1{F0: 0}
                _ = v39
            }
        }
    }
    if true {
        v40 := []int{(v2 - (v2 * (^(v2)))), -9223372036854775808, 240, (v2 % ((v2) | 1))}
        v41 := GenS3{F0: 42, F1: 32767}
        _ = v40
        _ = v41
    } else {
        var v42 int8 = 1
        v43 := GenS1{F0: 0}
        _ = v42
        _ = v43
    }
    v2 -= 42
    var v44 string = "a"
    v1 = (v1 + ((0.1 / v1) / ((0.0 - v1) + v1)))
    var v45 float64 = (1.5 * ((float64(0.1) / v1) / (3.14159 * float64(1e10))))
    var v46 uint64 = 136
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
}
