package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64; F2 int }
func genFn0(p0 int16, p1 uint, p2 bool) uint {
    var v6 int64 = -9223372036854775808
    switch 2147483647 {
    case 2147483647:
        var v7 int = -581
        _ = v7
    case -935:
        var v8 string = ("" + ((("\n\t" + "a") + "foo bar") + "a"))
        _ = v8
    }
    p0 = (-(p0))
    var v9 int64 = v6
    var v10 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v9
    return 760
}
func genFn1(p0 int32) int64 {
    var v11 int64 = 9223372036854775807
    var v12 int8 = 0
    v13 := []int{}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int = 127
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    v4 := make(map[string]string)
    v5 := []uint32{}
    fmt.Println(v1, v3)
    v1 = -2147483648
    v14 := GenS1{F0: "", F1: 0, F2: 876}
    if true {
        switch "\n\t" {
        case "foo bar":
            var v15 float64 = 3.14159
            _ = v15
        }
    }
    _ = len(v5)
    if ((genFn1(-942) == 1) || true) {
        if false {
            if false {
                fmt.Println(v1, v3)
                if (false || (290 < v3)) {
                    var v16 int8 = 109
                    _ = v16
                } else {
                    v1 = -2147483648
                }
            }
            if (1e10 > 1.5) {
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 uint32 = 1
                    _ = fi17
                    _ = v18
                }
            }
            var v19 int8 = 42
            _ = v19
        }
        var v20 int16 = -32768
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 bool = true
            v23 := GenS2{F0: ("" + (("\n\t" + "") + "\n\t")), F1: 431, F2: fi21}
            v20 += (v20 << ((947) & 63))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint64 = 4294967295
        var v26 int64 = 1
        _ = fi24
        _ = v25
        _ = v26
    }
    if ((42 < -779) && (42 <= 42)) {
        v3 = v3
        var v27 string = "a"
        v4[v27] = ""
        _ = v27
    } else {
        var v28 uint16 = 2
        v1 = v1
        if ((v1 % ((4294967295) | 1)) < int(v28)) {
            v3 += v3
            v3 -= v3
        }
        _ = v28
    }
    var v29 uint8 = 2
    if (false && true) {
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint64 = 426
                if (uint64(v29) <= v32) {
                    v3 += v3
                    v1 = 4294967295
                    var v33 uint8 = (v29 / ((v29) | 1))
                    var v34 string = ((("\n\t" + "") + "") + "\n\t")
                    _ = v33
                    _ = v34
                }
                _ = fi31
                _ = v32
            }
            var v35 float32 = 0.0
            fmt.Println(v1)
            _ = fi30
            _ = v35
        }
        var v36 int32 = -408
        _ = v36
    } else {
        v1 = 109
        v29 -= v29
    }
    var v37 uint32 = 70
    fmt.Println(v37, v37, v29)
    v38 := GenS2{F0: "foo bar", F1: 336, F2: int(v29)}
    fmt.Println(v37)
    switch (v1 * 0) {
    case 9223372036854775807:
        v39 := GenS2{F0: (("a" + ("hello" + "")) + ("a" + (("hello" + "a") + ("" + ("foo bar" + "\n\t"))))), F1: 1, F2: (v1 % ((-2147483648) | 1))}
        _ = v39
    case 32768:
        v29 += 127
    default:
        v29 -= 255
    }
    fmt.Println(v29)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
