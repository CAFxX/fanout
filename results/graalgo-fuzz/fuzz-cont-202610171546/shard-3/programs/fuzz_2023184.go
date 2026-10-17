package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
func genFn0(p0 uint64, p1 bool, p2 uint8) string {
    v6 := []int{-2147483648, -9223372036854775808, 42}
    switch (("\n\t" + ("hello" + "a")) + "") {
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 float64 = 1e10
            if ((!(!(false || (1 < (p0 | p0))))) && p1) {
                var v9 int8 = 102
                p0 -= 2
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    if false {
        v6 = append(v6, 148)
        fmt.Println(p2, p2, p1)
    } else {
        p1 = (!(!((p0 ^ (p0 - 1)) < p0)))
        p2 -= (p2 % ((p2) | 1))
        p2 = uint8(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}
func genFn1(p0 string) string {
    fmt.Println(p0)
    v10 := GenS1{F0: 65535, F1: 0, F2: -842}
    fmt.Println(p0)
    var v11 uint = 4294967295
    switch 81 {
    case 32767:
        v12 := []int{737}
        _ = v12
    case 282:
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v11, p0)
            var v14 int64 = -9223372036854775808
            switch p0 {
            case "a":
                v15 := []int{fi13, (^(fi13)), (fi13 / (((fi13 - fi13)) | 1))}
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    case -1:
        p0 = p0
    default:
        v16 := []int{-29, -9223372036854775808}
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v11
    return ("" + p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 float64 = 3.14159
    v4 := []int{2147483647, 0}
    var v5 string = "a"
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v3, v5, v2)
        v5 = v5
        _ = fi17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 -= float64(1e10)
        v5 = v5
        _ = fi18
    }
    v19 := v1[-2147483648]
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 string = v5
        fmt.Println(fi20, v19)
        _ = fi20
        _ = v21
    }
    v2 = -170
    fmt.Println(v2, v19, v3)
    var v22 uint8 = 2
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(v5)
        switch (fi23 >> ((465) & 63)) {
        case 4294967295:
            for fi24 := 0; fi24 < 9; fi24++ {
                v1[127] = (^(v19))
                var v25 int32 = (v2 | (v2 * 0))
                _ = fi24
                _ = v25
            }
        case 928:
            if (!false) {
                switch (v5 + "\n\t") {
                case "foo bar":
                    fmt.Println(v5, v22, v5)
                default:
                    switch v5 {
                    case "a":
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 string = ((("\n\t" + ((((v5 + v5) + (("" + v5) + "a")) + v5) + v5)) + v5) + v5)
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            }
        }
        _ = fi23
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 float32 = -2.25
        var v30 float64 = (v3 * v3)
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v2 -= (v2 + (v2 * 2147483647))
        var v32 bool = (false || (100 > -111))
        switch ((v5 + v5) + v5) {
        case "a":
            var v33 int64 = 1
            _ = v33
        case "hello":
            v19 = (v19 | (v19 + v19))
        case "":
            var v34 int8 = -58
            _ = v34
        default:
            v35 := v1[fi31]
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(v2, v2)
    var v36 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
