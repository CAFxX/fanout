package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
func genFn0(p0 int64) bool {
    v6 := []int{}
    var v7 int8 = -35
    if true {
        fmt.Println(p0, v7)
        var v8 int = 4294967295
        switch (v8 - 2147483647) {
        case 1:
            for fi9 := 0; fi9 < 6; fi9++ {
                if len(v6) > 0 {
                    v6[0] = fi9
                }
                _ = fi9
            }
        default:
            fmt.Println(p0, v8)
        }
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return (18446744073709551615 != 42)
}
func genFn1(p0 uint64, p1 uint8, p2 uint64) int32 {
    p1 -= 127
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int8 = 100
        switch "foo bar" {
        case "hello":
            v13 := []int{-44, (fi11 << ((637) & 63))}
            _ = v13
        default:
            v12 += v12
        }
        v10 = (v10 % (((v10 % (((v10 % (((^(v10))) | 1))) | 1))) | 1))
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        if false {
            var v15 int32 = 1
            p2 = p2
            _ = v15
        } else {
            fmt.Println(v10, p2)
            if (p0 != (^(p0))) {
                p1 += (v10 ^ v10)
            } else {
                var v16 int16 = 1000
                _ = v16
            }
        }
        v17 := GenS1{F0: 182, F1: 42, F2: 32767}
        _ = fi14
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 float32 = ((1e10 * -2.25) - 0.0)
    v3 := []uint32{0, 2147483647, 2, 4294967295}
    v4 := []uint32{}
    var v5 uint32 = 0
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 string = ("foo bar" + "")
        _ = fi18
        _ = v19
    }
    if (2147483647 > 756) {
        var v20 int32 = 923
        _ = v20
    } else {
        for fi21 := 0; fi21 < 19; fi21++ {
            v4 = append(v4, v5)
            switch (-(fi21)) {
            case 128:
                for fi22 := 0; fi22 < 19; fi22++ {
                    v23 := GenS1{F0: 240, F1: v5, F2: 1}
                    fmt.Println(fi21, fi21)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v24 bool = false
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            if false {
                var v25 uint64 = 474
                _ = v25
            }
            _ = fi21
        }
        fmt.Println(v1)
        fmt.Println(v2, v1, v2)
    }
    if false {
        var v26 float32 = v2
        var v27 bool = false
        var v28 int = 127
        _ = v26
        _ = v27
        _ = v28
    } else {
        fmt.Println(v2)
    }
    var v29 int8 = 1
    if (v1 == v1) {
        var v30 float32 = v2
        var v31 float32 = (float32(((float64((-2.25 - v2)) + 0.1) - (1.5 + -2.25))) + 0.1)
        fmt.Println(v5, v1, v29)
        _ = v30
        _ = v31
    } else {
        v5 = (v5 / (((v5 * (v5 * (v5 ^ (v5 & v5))))) | 1))
    }
    v5 -= (^(v5))
    if ((497 > 689) || (!true)) {
        if (2 == 165) {
            v32 := GenS1{F0: 1, F1: v5, F2: 0}
            _ = v32
        }
    } else {
        var v33 int32 = 2147483647
        _ = v33
    }
    if (18446744073709551615 > 966) {
        var v34 uint = 935
        _ = v34
    } else {
        if false {
            v29 -= -1
        }
        v5 -= v5
        v35 := GenS1{F0: 0, F1: v5, F2: 35}
        _ = v35
    }
    var v36 string = ""
    switch 79 {
    case 4294967295:
        switch -165 {
        case -9223372036854775808:
            v5 -= v5
        case -258:
            fmt.Println(v5)
        }
    }
    if false {
        v37 := GenS2{F0: 364, F1: "foo bar", F2: 0}
        _ = v37
    } else {
        fmt.Println(v5)
        v1 = (v1 - 1.5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (380 == 32767) {
        var v38 int32 = 0
        var v39 bool = false
        _ = v38
        _ = v39
    }
    fmt.Println(v2)
    var v40 uint64 = 68
    for fi41 := 0; fi41 < 4; fi41++ {
        v1 = float64((0.1 - float32(float64(3.14159))))
        _ = fi41
    }
    v29 -= 92
    var v42 int = -601
    var v43 string = ("foo bar" + v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
}
