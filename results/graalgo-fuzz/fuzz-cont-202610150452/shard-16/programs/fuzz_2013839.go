package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 int32) uint32 {
    if true {
        p0 += (-(p0))
    }
    v5 := GenS1{F0: 0, F1: 78}
    fmt.Println(p0)
    if false {
        p0 += -279
    }
    var v6 float64 = 3.14159
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() int16 {
    switch 4294967295 {
    case -883:
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 float64 = 1.5
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: 18446744073709551615, F1: fi7, F2: 4294967295}
            if false {
                var v11 int32 = 1
                v12 := []int{128, fi7, (fi7 ^ fi7)}
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v10
        }
    case 1:
        if (1e10 == 1e10) {
            switch "" {
            case "\n\t":
                v13 := GenS3{F0: "\n\t"}
                _ = v13
            case "a":
                v14 := []int{}
                _ = v14
            }
            var v15 int8 = 28
            _ = v15
        }
    default:
        var v16 float32 = 0.1
        _ = v16
    }
    var v17 uint32 = 2147483647
    var v18 uint = 182
    var v19 float32 = 3.14159
    fmt.Println(v17)
    _ = v17
    _ = v18
    _ = v19
    return -1
}

func main() {
    var v1 float64 = 1e10
    var v2 int16 = 42
    var v3 uint32 = 1
    var v4 int = -797
    switch (v4 / ((v4) | 1)) {
    case 127:
        fmt.Println(v2, v1, v1)
    case 42:
        v1 += (float64(-2.25) - -2.25)
    }
    switch "" {
    case "\n\t":
        var v20 bool = ((!true) || true)
        _ = v20
    case "a":
        fmt.Println(v3, v4, v2)
    case "hello":
        v3 += v3
    }
    if (0 > 42) {
        if true {
            v21 := []int{}
            var v22 int16 = (v2 / ((-46) | 1))
            var v23 bool = true
            _ = v21
            _ = v22
            _ = v23
        }
        switch 127 {
        case 269:
            if (false || (1 <= 0)) {
                v1 = v1
            }
        case 32767:
            v2 = -32768
        default:
            var v24 uint8 = 255
            _ = v24
        }
    } else {
        v25 := GenS1{F0: v2, F1: 971}
        var v26 string = "foo bar"
        _ = v25
        _ = v26
    }
    var v27 uint16 = 42
    switch v4 {
    case 42:
        v4 -= (^(v4))
    case -949:
        v28 := []int{v4}
        _ = v28
    }
    v1 = v1
    for fi29 := 0; fi29 < 12; fi29++ {
        switch ((("a" + "a") + "hello") + "a") {
        case "hello":
            fmt.Println(fi29, v3)
        case "a":
            var v30 uint64 = 1
            _ = v30
        case "foo bar":
            if false {
                v31 := []int{}
                var v32 int64 = -1
                _ = v31
                _ = v32
            } else {
                v27 -= 255
            }
        default:
            if false {
                fmt.Println(v2)
                switch (fi29 >> ((967) & 63)) {
                case 2147483647:
                    if (0 <= -9223372036854775808) {
                        fmt.Println(v1)
                    }
                default:
                    v33 := GenS2{F0: 0, F1: v4, F2: 816}
                    _ = v33
                }
                var v34 uint8 = 127
                _ = v34
            }
        }
        if (true || (int8(v3) == 1)) {
            v3 = v3
            if (9223372036854775807 == int64(v27)) {
                v35 := []int{}
                var v36 string = "\n\t"
                fmt.Println(v27, v3)
                _ = v35
                _ = v36
            } else {
                fmt.Println(v2, v4)
            }
            var v37 int8 = 1
            _ = v37
        }
        _ = fi29
    }
    v3 = 748
    var v38 uint64 = 0
    if false {
        if ("" <= ("hello" + ("a" + "\n\t"))) {
            v39 := []int{v4}
            var v40 int = 0
            v3 = v3
            _ = v39
            _ = v40
        } else {
            var v41 int8 = 0
            _ = v41
        }
    }
    v1 -= (1e10 - 1e10)
    v42 := []int{int(v3), (v4 - -1)}
    v38 = (v38 << ((597) & 63))
    v43 := GenS2{F0: 65535, F1: (v4 - (v4 & v4)), F2: v38}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
}
