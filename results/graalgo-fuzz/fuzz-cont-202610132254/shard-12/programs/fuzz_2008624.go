package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 string, p2 uint) string {
    fmt.Println(p0, p1)
    switch p1 {
    case "\n\t":
        for fi5 := 0; fi5 < 10; fi5++ {
            fmt.Println(p1, fi5, p1)
            _ = fi5
        }
    case "foo bar":
        switch -10 {
        case 32768:
            var v6 uint = p2
            _ = v6
        case 127:
            fmt.Println(p2, p0)
        default:
            fmt.Println(p0, p2)
        }
    }
    p2 += (p2 ^ 1)
    v7 := GenS1{F0: (p1 + p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 + "")
}
func genFn1(p0 uint, p1 int64, p2 int16) int16 {
    v8 := GenS2{F0: -1, F1: 42}
    v9 := []int{4294967295}
    var v10 uint8 = 79
    p1 -= 9223372036854775807
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            for fi12 := 0; fi12 < 3; fi12++ {
                p0 = 42
                _ = fi12
            }
        } else {
            switch fi11 {
            case -2147483648:
                v13 := GenS1{F0: "foo bar"}
                _ = v13
            case 32767:
                _ = len(v9)
            case 9223372036854775807:
                v9 = append(v9, 1)
            default:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint32 = 42
                v10 = (^(v10))
                _ = fi14
                _ = v15
            }
        }
        var v16 bool = (2147483648 >= 42)
        _ = fi11
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (p2 * p2)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 26
    v3 := []uint32{634, v1}
    var v4 int8 = v2
    _ = len(v3)
    var v17 uint8 = 2
    var v18 uint8 = v17
    if false {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int16 = 42
            var v21 int32 = 469
            _ = fi19
            _ = v20
            _ = v21
        }
        v17 += v17
    }
    fmt.Println(v2)
    v22 := GenS2{F0: 0, F1: 1}
    v3 = append(v3, v1)
    var v23 int8 = v4
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    fmt.Println(v4, v4)
    for fi24 := 0; fi24 < 12; fi24++ {
        switch -205 {
        case 2147483647:
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 string = "foo bar"
                var v27 uint16 = 878
                v28 := GenS1{F0: v26}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        case -9223372036854775808:
            var v29 float64 = 0.1
            _ = v29
        case -1:
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 uint16 = 0
                fmt.Println(v1)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    fmt.Println(v18)
                }
                var v32 string = ("a" + "foo bar")
                fmt.Println(v17, v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v33 float32 = -2.25
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        v18 = (^(v17))
        fmt.Println(v17, v23)
        _ = fi24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    v34 := GenS3{F0: 4294967295}
    fmt.Println(v17)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
}
