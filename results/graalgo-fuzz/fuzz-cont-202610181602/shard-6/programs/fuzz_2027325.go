package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 uint8) uint16 {
    var v5 uint64 = uint64(p0)
    var v6 int32 = -1
    switch "a" {
    case "foo bar":
        if (false || (true && false)) {
            var v7 float64 = (-2.25 + (1e10 / -2.25))
            var v8 int16 = 1
            var v9 uint64 = v5
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := []int{42, 4294967295}
            _ = v10
        }
    case "hello":
        v6 += v6
    case "":
        v11 := []int{998, 2147483647, -9223372036854775808, -611}
        _ = v11
    default:
        var v12 string = ""
        _ = v12
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint32, p1 int16) string {
    fmt.Println(p0, p0)
    v13 := GenS1{F0: p1, F1: 18446744073709551615, F2: false}
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = v13
    return "foo bar"
}

func main() {
    var v1 float64 = float64((float32(float64(1e10)) + (1.5 - -2.25)))
    v2 := []uint32{2, 736, 2, 805}
    var v3 int8 = -1
    v4 := make(map[int]int)
    v3 += 79
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 += v3
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 465, F1: 65535, F2: false}
            var v17 uint32 = 1
            var v18 string = "hello"
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    v1 -= v1
    v3 -= v3
    if (!(!(!(42 > genFn0(0))))) {
        v1 = (v1 * 0.1)
    }
    v3 += -103
    v19 := GenS1{F0: 1, F1: 1, F2: (false && true)}
    switch 1 {
    case 1:
        if true {
            var v20 int32 = -1
            fmt.Println(v3, v3)
            for fi21 := 0; fi21 < 3; fi21++ {
                if true {
                    var v22 uint8 = 127
                    var v23 int = (fi21 + (fi21 | (fi21 * -711)))
                    _ = v22
                    _ = v23
                } else {
                    v3 = (-(v3))
                    v24 := GenS1{F0: -1, F1: 65535, F2: true}
                    _ = v24
                }
                switch (fi21 / ((fi21) | 1)) {
                case 127:
                    if true {
                        fmt.Println(v3, v3, v3)
                    }
                case 952:
                    var v25 int8 = v3
                    _ = v25
                case 32768:
                    fmt.Println(v3)
                }
                _ = fi21
            }
            _ = v20
        } else {
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 float64 = v1
                v2 = append(v2, 2147483648)
                v27 -= ((v27 / 1.5) / 1e10)
                _ = fi26
                _ = v27
            }
            switch "\n\t" {
            case "a":
                v28 := v4[884]
                _ = v28
            case "\n\t":
                var v29 int64 = 42
                _ = v29
            case "hello":
                v1 = v1
            }
            v2 = append(v2, 0)
        }
    case 2147483647:
        v30 := GenS1{F0: 701, F1: 255, F2: false}
        _ = v30
    default:
        v1 += ((v1 + v1) + v1)
    }
    if (v1 != 0.0) {
        _ = len(v4)
    }
    var v31 float64 = float64((-2.25 - (0.1 / 1e10)))
    if true {
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            fmt.Println(v31, v3, v31)
        case "hello":
            var v32 int32 = 331
            _ = v32
        case "a":
            for fi33 := 0; fi33 < 16; fi33++ {
                fmt.Println(v3, v3, v3)
                v34 := GenS2{F0: 42, F1: -1, F2: false}
                _ = fi33
                _ = v34
            }
        }
        switch 9223372036854775807 {
        case 735:
            var v35 string = ("hello" + "foo bar")
            _ = v35
        }
        if (((v3 % (((-(v3))) | 1)) != (v3 / ((v3) | 1))) && (-1 <= 0)) {
            var v36 string = ""
            v31 = float64(float32((0.1 + float64(0.1))))
            if true {
                fmt.Println(v3)
            }
            _ = v36
        } else {
            for fi37 := 0; fi37 < 2; fi37++ {
                switch (fi37 % ((fi37) | 1)) {
                case 1:
                    switch (fi37 * (fi37 >> ((0) & 63))) {
                    case 642:
                        v38 := GenS2{F0: 485, F1: v3, F2: (1 != 792)}
                        _ = v38
                    case 1:
                        _ = len(v4)
                    case -2147483648:
                        v31 = (v1 + ((v1 + ((float64(3.14159) + (v31 * ((v31 + v31) * (1.5 * 0.1)))) - (float64(float32(v1)) - v1))) - float64(0.0)))
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                default:
                    v3 = (v3 % (((v3 | (^(v3)))) | 1))
                }
                _ = fi37
            }
            var v39 float32 = -2.25
            fmt.Println(v3)
            _ = v39
        }
    }
    switch "hello" {
    case "a":
        var v40 int32 = 1
        _ = v40
    case "foo bar":
        if true {
            fmt.Println(v31, v31)
        } else {
            var v41 uint8 = 107
            switch int(v41) {
            case -197:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if true {
                v31 = 0.1
                var v42 float32 = 1e10
                _ = v42
            } else {
                var v43 float32 = 1.5
                _ = v43
            }
            _ = v41
        }
    case "\n\t":
        fmt.Println(v1, v3)
    default:
        v44 := GenS2{F0: 32767, F1: (v3 >> ((1) & 63)), F2: (!(1 != -794))}
        _ = v44
    }
    var v45 float32 = 3.14159
    v3 = (v3 % ((41) | 1))
    v1 += (3.14159 * (1e10 * v1))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v45
}
