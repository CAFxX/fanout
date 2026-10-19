package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
func genFn0(p0 int8, p1 int) uint16 {
    v6 := GenS1{F0: 420, F1: 2147483647, F2: 255}
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7, p1, p1)
        switch fi7 {
        case 4294967295:
            p1 = (fi7 & (p1 | (p1 & p1)))
        }
        var v8 uint16 = 579
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 65535
}
func genFn1() int64 {
    v9 := []int{42, 127, 42, -1}
    v9 = append(v9, -212)
    var v10 uint64 = 1
    var v11 uint64 = v10
    if (42 <= -660) {
        v11 += (v11 >> ((255) & 63))
        _ = len(v9)
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            fmt.Println(v11)
            _ = fi12
        }
        v11 += 42
    }
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[string]string)
    v3 := []string{}
    v4 := []string{}
    var v5 bool = true
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = -9223372036854775808
        v14 = (v14 + 1)
        var v15 float64 = 0.1
        _ = fi13
        _ = v14
        _ = v15
    }
    delete(v2, "a")
    v16 := GenS2{F0: "foo bar"}
    v5 = v5
    if false {
        for fi17 := 0; fi17 < 12; fi17++ {
            if (0 <= -955) {
                v18 := GenS1{F0: -1, F1: 1, F2: 2}
                _ = v18
            }
            _ = fi17
        }
        var v19 float32 = ((((float32((3.14159 / 3.14159)) / 1.5) + -2.25) + 3.14159) - 3.14159)
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v5 = ((fi20 >> ((1) & 63)) == -9223372036854775808)
        if true {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        } else {
            v21 := GenS3{F0: 42, F1: genFn1()}
            for fi22 := 0; fi22 < 6; fi22++ {
                _ = len(v2)
                _ = fi22
            }
            _ = v21
        }
        var v23 uint32 = 2147483648
        _ = fi20
        _ = v23
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        _ = len(v2)
    case "a":
        fmt.Println(v5)
    default:
        if false {
            v5 = v5
            fmt.Println(v5)
            v24 := GenS4{F0: 18446744073709551615}
            _ = v24
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            if (false && (-669 == -2147483648)) {
                v5 = (-2.25 <= -2.25)
                v27 := GenS5{F0: fi26}
                var v28 uint = 1
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        _ = fi25
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := GenS4{F0: 765}
            var v32 float64 = 1.5
            v32 = v32
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v1 = append(v1, 4294967295)
        v34 := GenS4{F0: 18446744073709551615}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 bool = false
        _ = fi35
        _ = v36
    }
    v37 := GenS1{F0: 0, F1: 304, F2: 0}
    var v38 float64 = (((((0.1 - ((3.14159 + 1e10) + float64(1.5))) - 1.5) + float64(0.0)) / 0.1) * -2.25)
    fmt.Println(v38)
    for fi39 := 0; fi39 < 14; fi39++ {
        v5 = (1 != -1)
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
