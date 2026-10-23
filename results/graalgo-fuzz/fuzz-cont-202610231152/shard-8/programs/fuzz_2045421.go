package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() bool {
    var v7 int32 = -573
    v8 := []int{32768, -781, 237, 9223372036854775807}
    _ = v7
    _ = v8
    return true
}
func genFn1() uint16 {
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float64 = 0.0
        var v11 uint32 = 0
        v12 := GenS1{F0: (fi9 ^ fi9)}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if false {
        switch "\n\t" {
        case "hello":
            var v13 int16 = 511
            _ = v13
        }
    } else {
        var v14 int64 = 1
        if ((true || false) && true) {
            if false {
                v14 = v14
                var v15 uint16 = 2
                v15 -= (v15 + (v15 / ((v15) | 1)))
                _ = v15
            } else {
                v16 := []int{127, 42}
                _ = v16
            }
            fmt.Println(v14)
        } else {
            var v17 uint = 65535
            v18 := []int{}
            _ = v17
            _ = v18
        }
        v14 += v14
        _ = v14
    }
    switch 856 {
    case 32768:
        var v19 float32 = 3.14159
        _ = v19
    case 4294967295:
        v20 := []int{}
        _ = v20
    default:
        if false {
            var v21 string = "a"
            _ = v21
        }
    }
    var v22 bool = false
    _ = v22
    return 65535
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 42
    var v3 int32 = -2147483648
    v4 := make(map[int]uint32)
    v5 := []string{"foo bar", ("a" + ""), ((("hello" + "foo bar") + ("a" + "")) + "\n\t")}
    v6 := make(map[int]bool)
    if true {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    var v23 uint8 = 2
    v24 := GenS1{F0: 348}
    for fi25 := 0; fi25 < 2; fi25++ {
        if (((fi25 ^ (fi25 << ((912) & 63))) == fi25) && (4294967295 <= 724)) {
            if ("" >= (("a" + "") + "\n\t")) {
                var v26 float64 = 1e10
                v2 = (v2 * (-(v2)))
                v27 := GenS1{F0: 9223372036854775807}
                _ = v26
                _ = v27
            }
        } else {
            for fi28 := 0; fi28 < 19; fi28++ {
                if true {
                    fmt.Println(v3, v3)
                } else {
                    if ((0.1 * 0.0) == float64(1e10)) {
                        delete(v6, fi28)
                    }
                    var v29 uint8 = v23
                    v1 = v1
                    _ = v29
                }
                var v30 int32 = (^(v3))
                _ = fi28
                _ = v30
            }
            for fi31 := 0; fi31 < 8; fi31++ {
                v5 = append(v5, "a")
                var v32 bool = ((fi25 >> ((296) & 63)) >= (^(fi31)))
                v3 += (v3 >> ((4294967295) & 63))
                _ = fi31
                _ = v32
            }
        }
        v33 := GenS1{F0: fi25}
        v3 = int32(v2)
        _ = fi25
        _ = v33
    }
    fmt.Println(v1, v1, v1)
    var v34 float64 = -2.25
    switch 30 {
    case 42:
        v1 = (v1 - 3.14159)
    case 127:
        var v35 uint8 = (v23 * (v23 + v23))
        _ = v35
    case 67:
        v2 = (^(v2))
    default:
        switch 454 {
        case 276:
            switch "\n\t" {
            case "\n\t":
                v34 += ((((1e10 + -2.25) - v34) - (0.1 - (v34 / (v34 / -2.25)))) / v34)
            case "":
                v34 = 1.5
            case "hello":
                var v36 float32 = v1
                _ = v36
            }
        case -1:
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 uint16 = 65535
                for fi39 := 0; fi39 < 12; fi39++ {
                    var v40 int = fi39
                    v6[(-(fi37))] = (1 < 82)
                    _ = fi39
                    _ = v40
                }
                if false {
                    if (-721 == 1) {
                        var v41 int64 = -9223372036854775808
                        _ = v41
                    }
                }
                _ = fi37
                _ = v38
            }
        }
    }
    fmt.Println(v23, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
