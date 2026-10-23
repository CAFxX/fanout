package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }
type GenS6 struct { F0 int32; F1 int; F2 uint32 }
type GenS7 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 uint64) int8 {
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            p0 += p0
            fmt.Println(fi6, p0)
            p0 = (p0 * (p0 >> ((18446744073709551615) & 63)))
            _ = fi6
        }
        if true {
            var v7 bool = (!false)
            if (2 != 1) {
                v8 := GenS1{F0: 42, F1: 677}
                _ = v8
            } else {
                if true {
                    v9 := GenS2{F0: p0, F1: 420, F2: 1}
                    _ = v9
                }
            }
            _ = v7
        } else {
            switch "\n\t" {
            case "hello":
                var v10 uint16 = 255
                _ = v10
            case "\n\t":
                p0 -= p0
            default:
                var v11 float32 = ((((3.14159 / 3.14159) * 3.14159) - 1e10) - 0.1)
                _ = v11
            }
            p0 += p0
        }
        v12 := GenS1{F0: 0, F1: 116}
        _ = v12
    } else {
        p0 -= p0
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float64 = 0.0
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: -9223372036854775808, F1: -2147483648}
    v16 := GenS3{F0: 65535, F1: 994, F2: p0}
    var v17 int32 = int32(p0)
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return -54
}
func genFn1(p0 int, p1 uint64) bool {
    v18 := GenS4{F0: 42, F1: 331}
    fmt.Println(p1)
    v19 := GenS1{F0: -1, F1: 0}
    switch (p0 << ((996) & 63)) {
    case 4294967295:
        p1 = p1
    case 32768:
        var v20 uint64 = 4294967295
        _ = v20
    case 0:
        var v21 int16 = -182
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("" + "a")
    v3 := []string{"a"}
    var v4 float64 = (3.14159 * -2.25)
    var v5 int = 42
    v2 = "a"
    var v22 uint = 255
    var v23 int32 = -672
    for fi24 := 0; fi24 < 10; fi24++ {
        v22 += v22
        _ = fi24
    }
    if false {
        var v25 bool = ((v4 - float64(1.5)) != 0.1)
        v3 = append(v3, v2)
        if len(v3) > 0 {
            v3[0] = v2
        }
        _ = v25
    } else {
        switch v5 {
        case 4294967295:
            switch (v2 + ("a" + v2)) {
            case "hello":
                var v26 int32 = 797
                _ = v26
            default:
                v4 = float64(1.5)
            }
        }
        var v27 uint = v22
        var v28 bool = (42 == 0)
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 42, F1: (v23 / ((-1) | 1))}
    v30 := GenS5{F0: 64, F1: 49, F2: "a"}
    v31 := GenS6{F0: v23, F1: (v5 % ((-507) | 1)), F2: 2147483648}
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            delete(v1, ((v2 + v2) + v2))
            v22 += v22
            _ = fi33
        }
        _ = fi32
    }
    v34 := GenS6{F0: (v23 * (^(v23))), F1: (v5 << (((v22 % ((214) | 1))) & 63)), F2: 361}
    if true {
        if true {
            v35 := GenS1{F0: -9223372036854775808, F1: (v23 / ((v23) | 1))}
            _ = v35
        }
        v36 := v1[v2]
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 string = (v2 + (("a" + v2) + ("foo bar" + v36)))
            var v39 uint16 = 720
            switch ((v2 + (v2 + (v38 + v36))) + "hello") {
            case "":
                v36 = ("hello" + "")
            case "\n\t":
                _ = len(v3)
            case "a":
                var v40 float64 = float64(0.1)
                _ = v40
            }
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v36
    } else {
        _ = len(v3)
    }
    var v41 uint8 = 1
    if len(v3) > 0 {
        v3[0] = v2
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        v43 := GenS1{F0: -9223372036854775808, F1: 27}
        v44 := GenS7{F0: v22, F1: 113, F2: 367}
        v22 = (v22 & v22)
        _ = fi42
        _ = v43
        _ = v44
    }
    v41 += v41
    var v45 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v41
    _ = v45
    _ = v5
}
