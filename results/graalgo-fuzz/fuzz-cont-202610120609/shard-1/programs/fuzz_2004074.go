package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 int64, p1 uint32, p2 uint16) int64 {
    if ((1 < -2147483648) || (-1 > 42)) {
        if true {
            var v6 bool = ("\n\t" != "foo bar")
            _ = v6
        } else {
            var v7 int16 = 566
            var v8 int32 = 42
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(p0, p0, p0)
                p2 += p2
                _ = fi9
            }
            _ = v7
            _ = v8
        }
    }
    if true {
        var v10 bool = (!(p0 > (-(p0))))
        var v11 int8 = 54
        for fi12 := 0; fi12 < 7; fi12++ {
            if ((false || (int64(fi12) == 1)) && false) {
                if v10 {
                    v13 := []int{-2147483648, fi12, (fi12 * (fi12 * fi12))}
                    _ = v13
                }
                if (!(v11 == (v11 << ((411) & 63)))) {
                    p0 += (p0 | p0)
                    v14 := []int{fi12, fi12, fi12}
                    _ = v14
                }
                v15 := []int{fi12, -9223372036854775808, fi12}
                _ = v15
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        p0 += (p0 % ((p0) | 1))
        _ = fi16
    }
    var v17 float64 = float64(0.1)
    var v18 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return (p0 >> ((2) & 63))
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    v3 := make(map[uint32]string)
    var v4 int = 0
    v5 := make(map[int]string)
    if false {
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS1{F0: v2, F1: -69, F2: ("foo bar" != "foo bar")}
            var v21 uint32 = 2147483648
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint32 = 4294967295
        var v23 float64 = 1e10
        _ = v22
        _ = v23
    }
    v24 := v5[(v4 ^ (v4 << ((42) & 63)))]
    v2 = genFn0(-9223372036854775808, 42, 255)
    switch (^(v4)) {
    case 1:
        for fi25 := 0; fi25 < 4; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                v4 = 782
                v27 := GenS2{F0: 65535, F1: 2, F2: true}
                _ = fi26
                _ = v27
            }
            v2 += (v2 * v2)
            var v28 uint16 = 0
            _ = fi25
            _ = v28
        }
    default:
        var v29 int64 = v2
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        if false {
            v31 := GenS1{F0: v2, F1: 100, F2: ((v2 / ((v2) | 1)) < (v2 * (-(v2))))}
            _ = v31
        }
        _ = fi30
    }
    var v32 int16 = -32768
    v33 := v1[v24]
    fmt.Println(v33, v2, v33)
    var v34 uint = 4294967295
    for fi35 := 0; fi35 < 17; fi35++ {
        if (!(int(v32) <= fi35)) {
            v33 = v24
            if (v32 >= 953) {
                var v36 int = (-(fi35))
                v24 = ((v33 + v33) + (((v24 + ((v24 + (v24 + v24)) + v33)) + "\n\t") + ""))
                delete(v5, -269)
                _ = v36
            } else {
                if (!true) {
                    fmt.Println(v34)
                    v33 = ""
                } else {
                    var v37 float64 = 0.1
                    _ = v37
                }
                v38 := []int{int(v2), fi35}
                _ = v38
            }
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 uint8 = 0
            v34 = v34
            fmt.Println(v40)
            _ = fi39
            _ = v40
        }
        v41 := GenS2{F0: 2, F1: (v34 & v34), F2: ((-2.25 * (1e10 - 1.5)) > (3.14159 + ((0.0 + -2.25) + ((float32(float64((3.14159 / -2.25))) - (((0.0 / 3.14159) - 3.14159) * 0.1)) / ((float32((3.14159 + (float64(1.5) - 0.1))) - 0.1) * (float32(3.14159) - ((-2.25 * -2.25) - 0.1)))))))}
        _ = fi35
        _ = v41
    }
    v24 = v33
    var v42 int64 = v2
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 int16 = (v32 % ((v32) | 1))
        _ = fi43
        _ = v44
    }
    var v45 float64 = ((3.14159 / 1.0) - -2.25)
    if false {
        _ = len(v3)
        var v46 float64 = v45
        fmt.Println(v32, v33)
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v45
    _ = v5
}
