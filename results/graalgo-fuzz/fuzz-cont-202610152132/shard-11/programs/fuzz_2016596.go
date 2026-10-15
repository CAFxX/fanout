package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 int8) uint16 {
    var v6 uint16 = 32767
    p0 = ("a" <= "foo bar")
    _ = p0
    _ = p1
    _ = v6
    return v6
}
func genFn1(p0 uint32, p1 string) int32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p1 = p1
        v8 := []int{fi7, (fi7 >> ((4294967295) & 63))}
        _ = fi7
        _ = v8
    }
    var v9 bool = true
    v10 := GenS1{F0: 128, F1: 926, F2: -181}
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (-2.25 / -2.25)
    var v3 uint64 = 265
    var v4 uint64 = (v3 ^ v3)
    var v5 float64 = v2
    for fi11 := 0; fi11 < 9; fi11++ {
        switch (fi11 << ((0) & 63)) {
        case 415:
            v3 = 0
        case 275:
            v3 += v4
        }
        v12 := GenS1{F0: 128, F1: 1, F2: 32767}
        if (((v4 ^ v3) <= (v4 + (v4 % ((v3) | 1)))) && (!((!true) || false))) {
            v13 := GenS2{F0: 534, F1: 9223372036854775807, F2: 0}
            v14 := v1["foo bar"]
            v2 -= 0.0
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := []int{(fi11 + 32768)}
                v5 -= 3.14159
                _ = fi15
                _ = v16
            }
            if (776 < 1) {
                v17 := []int{fi11, (fi11 | fi11), fi11, 42}
                _ = v17
            }
        }
        _ = fi11
        _ = v12
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS3{F0: 1, F1: 0}
                v4 = (v4 << ((451) & 63))
                if true {
                    v5 = ((0.0 * 3.14159) + (((((float64(0.0) - v5) - (v5 * (0.0 + 0.0))) - 1.5) + float64((0.0 / 0.1))) * (v5 - v5)))
                    v2 += ((v5 / v5) - (float64((0.1 - 1e10)) + float64(0.0)))
                    v21 := []int{32767, fi18, (fi18 << ((255) & 63)), 42}
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
        }
        v2 -= (v5 * ((v5 / 1.5) / (v2 / v5)))
        _ = fi18
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v3 += (v3 - 958)
        _ = fi22
    }
    v23 := []int{-110}
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(v3)
        v5 += v2
        _ = fi24
    }
    var v25 uint8 = 127
    if false {
        v2 += (0.0 - v5)
        if (((-2.25 + float32(float64(3.14159))) * 0.0) > (0.1 + float32(v5))) {
            if ((!false) && false) {
                fmt.Println(v5, v25)
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v28 := GenS4{F0: 1, F1: 2, F2: -736}
                        _ = len(v1)
                        v4 = v3
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(fi26)
                    fmt.Println(v5, v5, v25)
                    _ = fi26
                }
                var v29 uint16 = 131
                _ = v29
            }
            v25 = (v25 + v25)
        } else {
            var v30 int32 = 242
            _ = v30
        }
    } else {
        var v31 uint16 = 847
        fmt.Println(v2, v3, v5)
        _ = v31
    }
    switch 128 {
    case 32767:
        var v32 int8 = 89
        _ = v32
    case 759:
        var v33 bool = ((!false) && (!(!(!false))))
        _ = v33
    case -9223372036854775808:
        v25 = (v25 - v25)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
