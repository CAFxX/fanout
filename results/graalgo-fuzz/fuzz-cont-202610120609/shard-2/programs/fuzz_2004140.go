package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 int8) uint16 {
    var v7 int16 = int16(p1)
    v8 := []int{32768, -506}
    for fi9 := 0; fi9 < 2; fi9++ {
        if (true && true) {
            var v10 uint64 = 998
            v10 = v10
            _ = v10
        } else {
            fmt.Println(p0)
            fmt.Println(p1, p1)
        }
        _ = fi9
    }
    var v11 int8 = 14
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 489
}
func genFn1(p0 bool, p1 int32, p2 uint) int32 {
    for fi12 := 0; fi12 < 3; fi12++ {
        switch ("a" + "\n\t") {
        case "":
            for fi13 := 0; fi13 < 10; fi13++ {
                p2 = p2
                switch fi13 {
                case -737:
                    var v14 int32 = p1
                    _ = v14
                case 602:
                    var v15 int = (fi13 | fi12)
                    _ = v15
                default:
                    var v16 uint16 = 42
                    _ = v16
                }
                var v17 uint8 = 127
                p0 = ((-9223372036854775808 >= -1) && p0)
                var v18 int = 1
                _ = fi13
                _ = v17
                _ = v18
            }
        case "foo bar":
            if ((!(!true)) || (false || ((fi12 | (fi12 ^ (fi12 - -249))) > 128))) {
                p1 = -2147483648
                p2 -= 972
            }
        default:
            p0 = ("hello" <= "")
        }
        fmt.Println(p1)
        if p0 {
            var v19 float32 = 1.5
            for fi20 := 0; fi20 < 3; fi20++ {
                switch "" {
                case "foo bar":
                    switch -595 {
                    case 4294967295:
                        p0 = (0 >= -80)
                    case 32768:
                        var v21 int8 = 71
                        _ = v21
                    }
                case "\n\t":
                    fmt.Println(v19, fi20, p0)
                case "":
                    v19 -= ((-2.25 * v19) + float32(float64(3.14159)))
                }
                var v22 int64 = 9223372036854775807
                _ = fi20
                _ = v22
            }
            _ = v19
        }
        _ = fi12
    }
    if true {
        var v23 float64 = -2.25
        _ = v23
    }
    switch 2147483647 {
    case -1:
        fmt.Println(p1, p0)
    case 32767:
        v24 := []int{2147483647, 9223372036854775807, 2147483647, 984}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int8 = 42
    var v2 int16 = 0
    v3 := []uint32{608, 2, 140}
    var v4 int64 = 0
    v5 := []int{}
    var v6 float64 = 0.1
    v25 := GenS1{F0: 935, F1: 2147483647, F2: -1}
    var v26 string = "a"
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 int8 = v1
        v6 += v6
        v6 = -2.25
        _ = fi27
        _ = v28
    }
    v4 = v4
    v1 += (v1 % (((v1 * v1)) | 1))
    v6 = float64(0.1)
    var v29 uint16 = 32767
    _ = len(v5)
    var v30 int32 = 42
    var v31 float32 = 1.5
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if false {
        v32 := GenS2{F0: v26, F1: (^(v1)), F2: 65535}
        v6 -= (v6 + (v6 / (1.5 - v6)))
        _ = v32
    } else {
        fmt.Println(v26, v6, v29)
        if true {
            var v33 float64 = 1.5
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 int16 = (^(v2))
                v3 = append(v3, 1)
                _ = fi34
                _ = v35
            }
            v5 = append(v5, 0)
            _ = v33
        }
        v3 = append(v3, 2147483647)
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        if (!true) {
            fmt.Println(v30, v6, v2)
        }
        if (1 < 255) {
            if false {
                if (true && false) {
                    var v37 string = v26
                    v2 = -818
                    _ = v37
                } else {
                    var v38 float32 = 1e10
                    _ = v38
                }
            }
        }
        var v39 int8 = v1
        _ = fi36
        _ = v39
    }
    switch -9223372036854775808 {
    case 32767:
        var v40 int16 = 32767
        _ = v40
    case -316:
        var v41 int = 24
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
