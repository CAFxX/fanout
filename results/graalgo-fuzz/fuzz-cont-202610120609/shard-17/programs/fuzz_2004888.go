package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 int64) uint64 {
    switch ((("" + ("" + ("" + ("a" + "\n\t")))) + "\n\t") + "hello") {
    case "a":
        v4 := GenS1{F0: (false || true)}
        _ = v4
    case "\n\t":
        for fi5 := 0; fi5 < 5; fi5++ {
            p1 -= 42
            v6 := GenS2{F0: 100, F1: 4294967295, F2: 18446744073709551615}
            _ = fi5
            _ = v6
        }
    case "hello":
        p1 += p1
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            p1 += p0
            _ = fi8
        }
        if true {
            var v9 uint64 = 18446744073709551615
            fmt.Println(v9, v9)
            switch 42 {
            case -2147483648:
                var v10 bool = (!false)
                _ = v10
            case 128:
                v11 := []int{(-(fi7))}
                _ = v11
            case 117:
                fmt.Println(fi7, p0)
            }
            _ = v9
        }
        var v12 float64 = (3.14159 / (1.5 + (-2.25 * (1e10 - 3.14159))))
        _ = fi7
        _ = v12
    }
    var v13 int64 = (p0 % ((p0) | 1))
    if (42 >= 32767) {
        var v14 float32 = float32(1.5)
        _ = v14
    } else {
        v15 := GenS3{F0: 19}
        p1 = 42
        if ((!true) && false) {
            p0 = (v13 | v13)
            var v16 uint16 = 65535
            fmt.Println(p0, v16)
            _ = v16
        } else {
            fmt.Println(p0, v13, p0)
            switch -215 {
            case 128:
                v17 := []int{128, 127}
                _ = v17
            case 32767:
                v13 += (-(v13))
            case 622:
                var v18 int = 391
                _ = v18
            }
            var v19 int64 = (p1 - v13)
            _ = v19
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = ((0.1 * 0.0) - (0.1 + (1e10 - 1e10)))
    var v3 int8 = 2
    v20 := GenS1{F0: false}
    if (-9223372036854775808 != -9223372036854775808) {
        var v21 uint64 = 1
        var v22 uint16 = 828
        v3 = v3
        _ = v21
        _ = v22
    } else {
        v23 := GenS1{F0: true}
        for fi24 := 0; fi24 < 9; fi24++ {
            if true {
                for fi25 := 0; fi25 < 6; fi25++ {
                    v3 = v3
                    var v26 string = "a"
                    _ = fi25
                    _ = v26
                }
            } else {
                var v27 uint8 = 127
                v28 := GenS1{F0: true}
                delete(v1, 65535)
                _ = v27
                _ = v28
            }
            _ = fi24
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 20; fi30++ {
                var v31 uint8 = 127
                fmt.Println(fi30, v2)
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        _ = v23
    }
    fmt.Println(v3, v3)
    var v32 uint64 = 1
    switch 127 {
    case 9223372036854775807:
        delete(v1, 1)
    case 637:
        var v33 int = 0
        _ = v33
    default:
        if false {
            v1[1] = 65535
            v34 := v1[2]
            _ = v34
        } else {
            var v35 int64 = -138
            var v36 int64 = (v35 + (v35 & -732))
            _ = v35
            _ = v36
        }
    }
    var v37 int16 = 1
    v2 += (((1e10 / v2) + v2) * 3.14159)
    for fi38 := 0; fi38 < 3; fi38++ {
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 uint = 65535
            fmt.Println(fi39)
            v37 = v37
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    var v41 int16 = v37
    v42 := GenS3{F0: 32767}
    v43 := []int{42, 32767, 194}
    fmt.Println(v32, v3)
    switch "\n\t" {
    case "foo bar":
        if true {
            var v44 uint64 = v32
            v3 += 1
            _ = v44
        }
    }
    if false {
        v3 = v3
    } else {
        fmt.Println(v41, v32)
    }
    var v45 int = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v37
    _ = v41
    _ = v42
    _ = v43
    _ = v45
}
