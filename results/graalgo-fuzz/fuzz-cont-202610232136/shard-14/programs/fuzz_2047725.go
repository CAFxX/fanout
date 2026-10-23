package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32) int8 {
    p0 -= -2147483648
    v6 := GenS1{F0: "hello"}
    p0 = p0
    var v7 int64 = 395
    var v8 uint = 2
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 5
}

func main() {
    var v1 uint32 = 1
    var v2 int32 = 42
    var v3 bool = (83 < 1)
    var v4 int = -2147483648
    var v5 float32 = 1e10
    v1 = (v1 >> ((1) & 63))
    v5 = (v5 * v5)
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: 32767}
        fmt.Println(v2)
        _ = fi9
        _ = v10
    }
    if (!true) {
        fmt.Println(v4, v3, v1)
        v11 := []int{127, v4, (v4 / ((v4) | 1)), v4}
        _ = v11
    }
    v12 := GenS1{F0: "hello"}
    fmt.Println(v1, v2)
    v2 += (-(v2))
    v5 += v5
    v4 -= (v4 + v4)
    for fi13 := 0; fi13 < 9; fi13++ {
        if (!(!false)) {
            var v14 int16 = 32767
            var v15 uint64 = 1
            _ = v14
            _ = v15
        } else {
            if v3 {
                if (-9223372036854775808 >= 1) {
                    var v16 int32 = (v2 >> ((822) & 63))
                    _ = v16
                }
                v1 -= (v1 * (v1 ^ (v1 % (((v1 / (((v1 + (v1 & (v1 + 2)))) | 1))) | 1))))
            } else {
                fmt.Println(v5, fi13)
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            switch (v4 - fi13) {
            case -1:
                v3 = (!v3)
            case 0:
                v18 := GenS1{F0: "a"}
                _ = v18
            case 425:
                switch (fi17 % ((fi17) | 1)) {
                case 2147483647:
                    var v19 uint32 = 0
                    _ = v19
                case -9223372036854775808:
                    if false {
                        v20 := GenS3{F0: 146}
                        _ = v20
                    } else {
                        fmt.Println(v4, v2, v2)
                        var v21 uint8 = 1
                        _ = v21
                    }
                case 42:
                    v5 += (v5 * (v5 + (v5 + 1.5)))
                }
            }
            fmt.Println(v5, v3, v2)
            var v22 uint = 486
            if (!(v2 > (v2 ^ 42))) {
                var v23 float32 = (((v5 * (v5 + (((0.0 + (-2.25 / 1.5)) - 0.0) - (0.1 + 1.5)))) * float32(1e10)) - 0.1)
                fmt.Println(fi17, v2, v1)
                _ = v23
            }
            var v24 int64 = 0
            _ = fi17
            _ = v22
            _ = v24
        }
        _ = fi13
    }
    if (32767 == 904) {
        var v25 int64 = -9223372036854775808
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 int8 = genFn0((v2 % (((v2 % ((2147483647) | 1))) | 1)))
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        _ = v25
    } else {
        var v29 int16 = 1000
        var v30 bool = (v3 && (!true))
        _ = v29
        _ = v30
    }
    v31 := []int{(v4 - (v4 % ((v4) | 1))), (v4 | (-(v4))), (v4 - v4), v4}
    for fi32 := 0; fi32 < 7; fi32++ {
        v4 += (v4 + fi32)
        v33 := GenS3{F0: (v1 ^ v1)}
        _ = fi32
        _ = v33
    }
    v3 = (0 != 1)
    var v34 uint64 = 770
    var v35 int32 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
