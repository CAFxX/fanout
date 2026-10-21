package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
func genFn0(p0 uint32, p1 int16, p2 int8) uint16 {
    var v6 uint16 = 760
    var v7 int8 = (p2 ^ (p2 << ((1) & 63)))
    p0 += p0
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: p2, F1: (-(v7))}
            var v12 uint = 65535
            if true {
                var v13 uint64 = 18446744073709551615
                var v14 uint8 = (v8 >> ((v12) & 63))
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v15 := []int{127}
        var v16 int8 = v7
        _ = fi9
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v6
}

func main() {
    var v1 uint = 1
    var v2 int = 4294967295
    var v3 int8 = 0
    var v4 float32 = ((float32(0.0) - 3.14159) / (0.0 + 0.1))
    var v5 bool = ((42 > 1000) || false)
    v5 = v5
    v17 := []int{(v2 | (v2 << ((1) & 63))), 4294967295, v2, 127}
    var v18 uint64 = 539
    v19 := GenS2{F0: (("a" + "") + ("a" + ""))}
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v4, v4)
        var v21 bool = ((2147483648 < 912) || (0 <= 42))
        _ = fi20
        _ = v21
    }
    v2 -= (v2 | v2)
    for fi22 := 0; fi22 < 13; fi22++ {
        v2 = -1
        _ = fi22
    }
    v23 := GenS2{F0: ("hello" + "a")}
    _ = len(v17)
    switch 32767 {
    case 1:
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 bool = true
            v4 += v4
            _ = fi24
            _ = v25
        }
    default:
        if (!v5) {
            for fi26 := 0; fi26 < 4; fi26++ {
                fmt.Println(v4)
                _ = fi26
            }
            if (int16(v3) == 0) {
                var v27 int32 = int32(v1)
                _ = v27
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        } else {
            v18 = (v18 / ((0) | 1))
            for fi28 := 0; fi28 < 1; fi28++ {
                if ((v4 * float32((0.0 / -2.25))) != (v4 * v4)) {
                    if len(v17) > 0 {
                        v17[0] = v2
                    }
                }
                _ = fi28
            }
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS1{F0: v3, F1: v3}
                fmt.Println(v2)
                _ = fi29
                _ = v30
            }
        }
    }
    if false {
        v1 = v1
        var v31 bool = (!(1 == 0))
        var v32 uint8 = 2
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int16 = 1000
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 float32 = v4
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 string = "foo bar"
            var v39 uint8 = 255
            fmt.Println(fi35, v1)
            _ = fi37
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 4; fi40++ {
            switch ("hello" + ("\n\t" + "a")) {
            case "":
                v2 -= 32768
            case "\n\t":
                fmt.Println(v2, fi35, fi35)
            case "a":
                v2 -= (fi40 << (((v1 % ((v1) | 1))) & 63))
            }
            v36 += 0.1
            _ = fi40
        }
        _ = fi35
        _ = v36
    }
    v41 := GenS2{F0: ("hello" + ("" + ""))}
    switch (v2 * (v2 + (v2 / ((v2) | 1)))) {
    case 828:
        var v42 float64 = ((-2.25 + -2.25) + 3.14159)
        _ = v42
    case 2147483647:
        switch "hello" {
        case "":
            v17 = append(v17, int(v3))
        }
    case 0:
        v18 += v18
    }
    switch v2 {
    case 127:
        if true {
            v3 -= v3
        }
    case 0:
        v43 := GenS3{F0: ((("" + ("hello" + "")) + "\n\t") + "a")}
        _ = v43
    case 4294967295:
        v44 := GenS1{F0: v3, F1: (v3 >> ((v1) & 63))}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
