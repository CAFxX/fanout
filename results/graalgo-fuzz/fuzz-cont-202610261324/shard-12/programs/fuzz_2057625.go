package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0(p0 int64, p1 uint) uint {
    if ((true && true) || (1 != uint16(p1))) {
        switch "" {
        case "\n\t":
            var v4 bool = false
            _ = v4
        case "foo bar":
            v5 := []int{0, -2147483648, 1}
            _ = v5
        default:
            p0 = (p0 >> ((p1) & 63))
        }
        v6 := []int{}
        v7 := GenS1{F0: ("a" + ("foo bar" + ("" + "foo bar"))), F1: -809, F2: (1 == -128)}
        _ = v6
        _ = v7
    } else {
        var v8 uint64 = 933
        p1 += p1
        _ = v8
    }
    v9 := GenS2{F0: -1}
    for fi10 := 0; fi10 < 1; fi10++ {
        p0 = p0
        v11 := GenS2{F0: 42}
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 129
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return p1
}

func main() {
    var v1 int64 = 42
    v2 := []int{-2147483648, -1, 1}
    v3 := make(map[int]bool)
    v1 -= (v1 >> ((255) & 63))
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v1 += (-(v1))
            var v15 int = (fi14 * fi14)
            var v16 bool = true
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int = (fi13 / (((-(fi13))) | 1))
        for fi18 := 0; fi18 < 13; fi18++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi18
        }
        _ = fi13
        _ = v17
    }
    v1 = (v1 << ((65535) & 63))
    for fi19 := 0; fi19 < 19; fi19++ {
        v1 += v1
        for fi20 := 0; fi20 < 8; fi20++ {
            v1 = v1
            var v21 int = fi20
            v22 := GenS1{F0: "", F1: 42, F2: (2 != 2)}
            switch ("hello" + ("" + "\n\t")) {
            case "a":
                v23 := GenS3{F0: 575, F1: 255, F2: -766}
                _ = v23
            case "hello":
                v21 = (v21 & 1)
            case "":
                var v24 int32 = -1
                _ = v24
            default:
                _ = len(v3)
            }
            fmt.Println(fi19)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    switch "hello" {
    case "hello":
        v2 = append(v2, 675)
    case "foo bar":
        v25 := v3[254]
        _ = v25
    case "":
        v26 := GenS2{F0: 1}
        _ = v26
    }
    _ = len(v3)
    switch (("hello" + "a") + "") {
    case "a":
        v2 = append(v2, 473)
    case "hello":
        var v27 bool = true
        _ = v27
    case "foo bar":
        for fi28 := 0; fi28 < 17; fi28++ {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 float64 = float64(3.14159)
                fmt.Println(v1)
                v2 = append(v2, -2147483648)
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    v31 := v3[-9223372036854775808]
    var v32 float64 = (1.5 - 1e10)
    _ = len(v3)
    if (-2.25 == 0.0) {
        v1 = v1
    } else {
        for fi33 := 0; fi33 < 5; fi33++ {
            v34 := GenS1{F0: ("\n\t" + ("a" + "\n\t")), F1: 2147483647, F2: (!((!true) || ((319 <= 312) || true)))}
            _ = fi33
            _ = v34
        }
        var v35 int16 = 0
        _ = v35
    }
    v36 := GenS3{F0: 509, F1: 65535, F2: 593}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
}
