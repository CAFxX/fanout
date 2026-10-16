package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 uint, p2 uint) string {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint64 = 2
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int = -9223372036854775808
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    v12 := []int{2147483647, 182, 448}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return "hello"
}

func main() {
    var v1 uint32 = 65535
    v2 := []int{9223372036854775807, 289, 127}
    var v3 uint32 = (v1 / (((v1 / ((272) | 1))) | 1))
    var v4 int64 = -9223372036854775808
    var v5 uint32 = v3
    v6 := []string{"", "", ("\n\t" + "foo bar")}
    if ((!false) || (919 <= 951)) {
        v3 = v1
        v13 := GenS1{F0: 32, F1: -2147483648}
        if true {
            for fi14 := 0; fi14 < 5; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    v2 = append(v2, (-(fi15)))
                    var v16 uint8 = 127
                    _ = fi15
                    _ = v16
                }
                v5 += v5
                fmt.Println(fi14)
                _ = fi14
            }
            var v17 bool = true
            _ = v17
        } else {
            var v18 float64 = float64(0.1)
            _ = v18
        }
        _ = v13
    } else {
        v5 += (v5 >> ((484) & 63))
        v1 = v1
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        if (0 < 1) {
            v6 = append(v6, "hello")
        } else {
            v4 = v4
        }
        var v20 uint16 = 42
        v21 := GenS1{F0: 0, F1: 921}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if len(v2) > 0 {
            v2[0] = (fi22 >> ((4294967295) & 63))
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: 128, F1: 317}
            var v25 int8 = 64
            var v26 int64 = (^(v4))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v26 -= 111
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
    }
    var v27 bool = (false && (v4 >= 1))
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 int32 = int32(v5)
        _ = fi28
        _ = v29
    }
    if (255 < 1) {
        v30 := GenS1{F0: 1, F1: 0}
        switch 128 {
        case 577:
            v2 = append(v2, 42)
        case 128:
            if len(v2) > 0 {
                v2[0] = 127
            }
        case 127:
            var v31 bool = (0 <= 974)
            _ = v31
        }
        _ = v30
    } else {
        var v32 uint32 = 593
        switch ("\n\t" + "foo bar") {
        case "":
            switch "\n\t" {
            case "a":
                v6 = append(v6, "hello")
            case "\n\t":
                _ = len(v6)
            }
        case "foo bar":
            fmt.Println(v27, v27)
        }
        _ = v32
    }
    switch "\n\t" {
    case "":
        var v33 uint = 65535
        _ = v33
    case "\n\t":
        for fi34 := 0; fi34 < 3; fi34++ {
            _ = len(v2)
            var v35 bool = (!(!((((1.5 * 0.0) / 3.14159) / 1.5) != 1.5)))
            var v36 uint16 = uint16(v3)
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v37 := GenS1{F0: 128, F1: -438}
    v38 := GenS2{F0: v27}
    v6 = append(v6, ("a" + ("a" + "")))
    v39 := GenS3{F0: -469, F1: 368}
    if true {
        v5 = v3
        v4 = (-(v4))
        var v40 float32 = (((1e10 / (0.1 * 1e10)) + 0.0) * 1.5)
        _ = v40
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5, v5, v1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v27 = false
    if ((v5 | v3) != 0) {
        if (!false) {
            v1 += (v5 ^ (^(v5)))
            v2 = append(v2, -605)
        } else {
            v27 = ((0.1 - 0.1) == ((0.0 / ((-2.25 / float32((0.0 + 1e10))) / 1e10)) / (-2.25 / 0.1)))
        }
    }
    if (42 <= 127) {
        if v27 {
            var v41 float32 = float32(1e10)
            var v42 string = "a"
            if len(v2) > 0 {
                v2[0] = 699
            }
            _ = v41
            _ = v42
        }
    } else {
        fmt.Println(v3)
        v43 := GenS4{F0: v27}
        v5 -= (v5 & (^(v3)))
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
