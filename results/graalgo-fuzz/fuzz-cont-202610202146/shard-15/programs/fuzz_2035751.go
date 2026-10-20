package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 int16) uint32 {
    if true {
        p0 = (p0 + "\n\t")
        p1 += (p1 + (p1 << ((42) & 63)))
        if true {
            v5 := GenS1{F0: 42, F1: -488}
            _ = v5
        } else {
            switch -552 {
            case 312:
                var v6 int16 = (-(p1))
                _ = v6
            case -438:
                var v7 uint = 255
                _ = v7
            case 127:
                for fi8 := 0; fi8 < 5; fi8++ {
                    p0 = "a"
                    switch "\n\t" {
                    case "a":
                        v9 := GenS2{F0: 84}
                        _ = v9
                    case "hello":
                        switch fi8 {
                        case 32768:
                            p1 += p1
                        case -74:
                            var v10 int = (fi8 / (((fi8 | -574)) | 1))
                            _ = v10
                        case 42:
                            v11 := []int{}
                            _ = v11
                        }
                    case "":
                        var v12 uint16 = 460
                        _ = v12
                    default:
                        var v13 uint8 = 11
                        _ = v13
                    }
                    fmt.Println(p1, p1, p0)
                    _ = fi8
                }
            default:
                v14 := []int{1}
                _ = v14
            }
            var v15 float64 = 1e10
            _ = v15
        }
    }
    if false {
        var v16 uint16 = 2
        v17 := GenS2{F0: 127}
        _ = v16
        _ = v17
    }
    p0 = ("hello" + p0)
    var v18 bool = true
    for fi19 := 0; fi19 < 14; fi19++ {
        if v18 {
            p1 = p1
            switch "" {
            case "":
                var v20 int32 = -2147483648
                _ = v20
            case "\n\t":
                var v21 int32 = -939
                _ = v21
            }
        } else {
            v22 := GenS2{F0: -53}
            var v23 int64 = -1
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "a":
            v24 := GenS1{F0: (p1 * 1000), F1: (fi19 * (fi19 - fi19))}
            _ = v24
        default:
            var v25 uint32 = 450
            _ = v25
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = v18
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    var v3 int32 = 0
    var v4 int16 = -418
    switch 984 {
    case 32767:
        v3 += v3
    default:
        var v26 bool = true
        _ = v26
    }
    v4 = v4
    fmt.Println(v4, v2)
    switch 42 {
    case -957:
        v27 := GenS2{F0: 1}
        _ = v27
    case 648:
        var v28 int8 = 42
        _ = v28
    default:
        if (-2.25 > 1e10) {
            var v29 uint32 = 42
            var v30 string = "foo bar"
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        fmt.Println(v4)
        if (true || false) {
            var v32 uint = 4294967295
            v3 = (-(v3))
            _ = v32
        }
        var v33 bool = ((964 != 2) && true)
        _ = fi31
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 uint8 = v2
        v2 = 0
        _ = fi34
        _ = v35
    }
    fmt.Println(v2, v2, v4)
    v4 += -1
    var v36 uint = 18446744073709551615
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := GenS1{F0: 1, F1: fi37}
        _ = fi37
        _ = v38
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
