package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0() uint8 {
    var v7 uint16 = 1
    v8 := []int{32767}
    if (v7 == (v7 % ((0) | 1))) {
        var v9 int32 = 42
        _ = v9
    } else {
        var v10 int = 751
        v10 -= v10
        v10 += (v10 ^ (-(v10)))
        _ = v10
    }
    v7 += v7
    v7 -= 831
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int) int32 {
    v11 := GenS1{F0: 42, F1: 455}
    p0 = -2147483648
    _ = p0
    _ = v11
    return -2147483648
}

func main() {
    var v1 int64 = 42
    v2 := []string{"hello", "", "a", "a"}
    var v3 uint = 146
    var v4 uint32 = 65535
    var v5 string = "a"
    var v6 uint = (v3 | v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint32 = v4
        for fi14 := 0; fi14 < 13; fi14++ {
            if false {
                var v15 int32 = -2147483648
                v16 := GenS2{F0: 32768}
                var v17 string = "a"
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: fi12}
            v1 = (v1 | v1)
            _ = fi14
            _ = v18
        }
        var v19 bool = true
        _ = fi12
        _ = v13
        _ = v19
    }
    if (true && ((v4 << (((v3 % ((0) | 1))) & 63)) <= v4)) {
        _ = len(v2)
        v3 = (v6 & (v3 & (v6 << ((v3) & 63))))
        v2 = append(v2, v5)
    }
    v4 = (v4 / ((v4) | 1))
    v3 += (v3 | (v6 * (v6 * 1)))
    fmt.Println(v3, v3, v5)
    v3 = (v6 * v3)
    var v20 bool = false
    var v21 int = 0
    v2 = append(v2, v5)
    if (true || v20) {
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS3{F0: 1000, F1: 42}
            v20 = (!((-(v1)) > v1))
            _ = fi22
            _ = v23
        }
    }
    switch (-(v21)) {
    case 288:
        v24 := GenS1{F0: (v1 << ((v3) & 63)), F1: genFn1(1)}
        _ = v24
    case 1:
        v4 = 2147483648
    case 968:
        if len(v2) > 0 {
            v2[0] = v5
        }
    default:
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint64 = 2
                fmt.Println(v5, v4, v20)
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    v6 = (v6 - v6)
    for fi28 := 0; fi28 < 12; fi28++ {
        switch fi28 {
        case 0:
            v29 := GenS2{F0: (v21 >> (((v3 % ((0) | 1))) & 63))}
            _ = v29
        }
        fmt.Println(v5, v1, v21)
        if ((!v20) || (!v20)) {
            for fi30 := 0; fi30 < 11; fi30++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v2 = append(v2, v5)
                if (!(((1e10 < -2.25) || (fi28 <= (-(fi28)))) && (fi30 == 4294967295))) {
                    v5 = v5
                    var v31 uint = v6
                    _ = v31
                } else {
                    fmt.Println(v20, fi30, v20)
                }
                _ = fi30
            }
            var v32 string = ((v5 + "foo bar") + v5)
            _ = v32
        } else {
            v4 = (^(v4))
        }
        _ = fi28
    }
    v2 = append(v2, "\n\t")
    v20 = v20
    for fi33 := 0; fi33 < 17; fi33++ {
        v3 -= v3
        fmt.Println(v1, v4, v20)
        var v34 bool = (v20 && v20)
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
