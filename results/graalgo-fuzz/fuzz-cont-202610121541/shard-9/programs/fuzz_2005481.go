package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int8) int8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 += (^(p0))
        _ = fi6
    }
    switch "hello" {
    case "\n\t":
        v7 := []int{}
        _ = v7
    case "hello":
        v8 := []int{}
        _ = v8
    case "":
        p1 = p1
    default:
        v9 := []int{1, 1}
        _ = v9
    }
    v10 := GenS1{F0: 13, F1: -9223372036854775808}
    if true {
        fmt.Println(p1, p0)
        fmt.Println(p1, p1)
        switch 4294967295 {
        case 703:
            v11 := []int{-618, 2147483647, -2147483648}
            _ = v11
        case -9223372036854775808:
            p1 = p1
        case 42:
            p1 = p1
        }
    } else {
        switch "a" {
        case "hello":
            p1 = p1
        case "":
            var v12 int64 = 42
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    var v1 bool = false
    var v2 int = 677
    var v3 uint8 = 255
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    if (955 >= 0) {
        var v13 uint32 = 0
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 float64 = 0.0
            var v16 uint8 = (v3 + v3)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        v1 = (0 != 0)
        var v17 float64 = 1e10
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 float32 = (-2.25 - -2.25)
        for fi20 := 0; fi20 < 7; fi20++ {
            v2 -= fi20
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v1, v1)
    if v1 {
        v21 := []int{-9223372036854775808, (v2 % ((v2) | 1)), (v2 + (v2 - v2))}
        if false {
            v1 = (42 >= 539)
            switch v2 {
            case 1:
                v2 += v2
            case 2147483647:
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            case -473:
                for fi22 := 0; fi22 < 4; fi22++ {
                    if v1 {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            fmt.Println(v3)
                            var v24 bool = (!v1)
                            v3 -= v3
                            _ = fi23
                            _ = v24
                        }
                    }
                    _ = fi22
                }
            }
            var v25 string = ("a" + "hello")
            _ = v25
        } else {
            for fi26 := 0; fi26 < 13; fi26++ {
                v3 -= 1
                var v27 int64 = 42
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 18; fi28++ {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                fmt.Println(v1, v1, v3)
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v3, v1, fi28)
                }
                _ = fi28
            }
        }
        _ = v21
    } else {
        v29 := []int{-1}
        fmt.Println(v1)
        _ = v29
    }
    var v30 float64 = 0.0
    for fi31 := 0; fi31 < 4; fi31++ {
        switch 2147483647 {
        case 127:
            if v1 {
                fmt.Println(v1, v2, v1)
                var v32 string = "hello"
                _ = v32
            }
        }
        _ = fi31
    }
    var v33 int32 = -902
    for fi34 := 0; fi34 < 20; fi34++ {
        v35 := GenS2{F0: 212}
        v36 := GenS1{F0: (^(v3)), F1: (v2 / ((-2147483648) | 1))}
        var v37 int64 = 9223372036854775807
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    fmt.Println(v1, v2, v30)
    v38 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v38
    _ = v4
    _ = v5
}
