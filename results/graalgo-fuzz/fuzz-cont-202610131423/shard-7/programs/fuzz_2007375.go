package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0() int64 {
    if (!(!false)) {
        if false {
            var v7 int16 = -540
            v8 := GenS1{F0: (65535 == 189), F1: (((("a" + ("" + "foo bar")) + "") + ("hello" + "hello")) > ("hello" + "hello"))}
            _ = v7
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
    }
    if (!((true && (true && (true || true))) && false)) {
        v11 := []int{-1, -9223372036854775808, 32768}
        _ = v11
    }
    return -1
}

func main() {
    var v1 int64 = -1
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 123
    var v6 float64 = 1e10
    delete(v3, 195)
    fmt.Println(v6, v2, v5)
    v2 = (v1 != 1)
    switch "foo bar" {
    case "\n\t":
        v1 = v1
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        switch (fi12 | 4294967295) {
        case 32768:
            fmt.Println(v6, v1)
        case 9223372036854775807:
            delete(v4, 2147483647)
        case 398:
            v2 = true
        default:
            switch (("\n\t" + "foo bar") + "hello") {
            case "":
                switch "" {
                case "":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 uint = 18446744073709551615
                        var v15 bool = ((v6 * float64(-2.25)) >= 3.14159)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case "\n\t":
                    var v16 bool = (!v2)
                    _ = v16
                default:
                    fmt.Println(v6, v2, v6)
                }
            case "hello":
                if v2 {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        fmt.Println(v2)
                        _ = fi17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 int64 = v1
                        v4[0] = false
                        fmt.Println(v1)
                        _ = fi18
                        _ = v19
                    }
                }
            }
        }
        _ = fi12
    }
    if (296 > 2147483647) {
        v20 := GenS2{F0: 1}
        _ = v20
    } else {
        switch "a" {
        case "foo bar":
            if false {
                if ((!v2) && v2) {
                    v21 := GenS3{F0: 1, F1: 32767, F2: 2147483647}
                    _ = v21
                }
            }
        default:
            v1 += genFn0()
        }
        v6 -= v6
        var v22 int = 32767
        _ = v22
    }
    v23 := GenS1{F0: true, F1: ((0 < -2147483648) && v2)}
    if ((1 != 32767) && v2) {
        v1 += 132
        v6 -= v6
    } else {
        if v2 {
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := []int{fi24, (fi24 | fi24), fi24, (fi24 * (fi24 * fi24))}
                _ = fi24
                _ = v25
            }
        }
        delete(v3, 2147483647)
        var v26 uint8 = (v5 << ((849) & 63))
        _ = v26
    }
    v6 = float64(1.5)
    var v27 string = "a"
    switch 209 {
    case 1:
        var v28 string = v27
        _ = v28
    case -2147483648:
        v29 := v4[65535]
        _ = v29
    }
    if (((v6 / 1e10) >= (v6 - 3.14159)) || (false && ((42 >= 32767) && (false && true)))) {
        if (!v2) {
            v27 = v27
            switch 42 {
            case 1:
                v6 = 0.0
            default:
                v30 := GenS2{F0: -2147483648}
                _ = v30
            }
        } else {
            v31 := GenS4{F0: 2147483648, F1: 1, F2: 1}
            _ = v31
        }
        fmt.Println(v6, v1)
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        v33 := GenS3{F0: 4294967295, F1: 1, F2: 42}
        switch v27 {
        case "":
            if v2 {
                v2 = v2
                var v34 float32 = 3.14159
                v3[0] = 2
                _ = v34
            } else {
                fmt.Println(v1)
                var v35 bool = ((!v2) || (false || ((!(!v2)) || true)))
                _ = v35
            }
        default:
            switch (fi32 / ((fi32) | 1)) {
            case 2147483647:
                v36 := GenS4{F0: 855, F1: 912, F2: 2}
                _ = v36
            default:
                if (v2 && false) {
                    if false {
                        v5 -= (v5 >> ((42) & 63))
                    }
                    var v37 float32 = 0.1
                    _ = v37
                }
            }
        }
        _ = fi32
        _ = v33
    }
    var v38 int16 = -782
    var v39 float32 = 3.14159
    if ((v6 + -2.25) != v6) {
        var v40 uint64 = 1
        _ = v40
    } else {
        var v41 int8 = 42
        _ = v41
    }
    var v42 int16 = (v38 ^ (^(v38)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
