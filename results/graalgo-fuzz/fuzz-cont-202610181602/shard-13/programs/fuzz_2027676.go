package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
func genFn0(p0 uint16, p1 int, p2 uint) int64 {
    var v7 bool = true
    switch p1 {
    case 357:
        switch "hello" {
        case "hello":
            if false {
                p0 += (p0 % ((p0) | 1))
                if (!(float32(0.1) <= 3.14159)) {
                    var v8 int32 = 42
                    if v7 {
                        var v9 float64 = (1.5 / 1.5)
                        _ = v9
                    } else {
                        v10 := []int{p1, 1}
                        _ = v10
                    }
                    _ = v8
                }
            } else {
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 bool = false
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
                v14 := []int{p1, (p1 / ((1) | 1)), p1}
                v14 = append(v14, (p1 % ((p1) | 1)))
                _ = v14
            }
        case "a":
            switch p1 {
            case 4294967295:
                var v15 int32 = -325
                _ = v15
            case -261:
                fmt.Println(p0, p0)
            }
        default:
            var v16 int64 = -890
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -627
}
func genFn1(p0 uint16, p1 int) uint {
    for fi17 := 0; fi17 < 12; fi17++ {
        if (true && true) {
            var v18 string = "\n\t"
            _ = v18
        }
        var v19 uint = 18446744073709551615
        _ = fi17
        _ = v19
    }
    v20 := GenS1{F0: 363, F1: 255, F2: p1}
    var v21 float32 = (-2.25 * 1e10)
    p0 = p0
    for fi22 := 0; fi22 < 10; fi22++ {
        v21 = v21
        var v23 int8 = 95
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 65535
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = -2.25
    var v3 int16 = -160
    v4 := []int{322, 2147483647, 621}
    var v5 int8 = -92
    var v6 int = 0
    fmt.Println(v2, v5, v6)
    _ = len(v4)
    switch "" {
    case "foo bar":
        fmt.Println(v6)
    case "hello":
        var v24 uint16 = 0
        _ = v24
    default:
        switch (v6 | v6) {
        case -1:
            var v25 uint8 = (v1 & v1)
            _ = v25
        case 472:
            var v26 string = ("a" + "foo bar")
            _ = v26
        default:
            if true {
                v27 := GenS1{F0: 4294967295, F1: v1, F2: int(v1)}
                v4 = append(v4, (v6 ^ (v6 & (v6 >> ((634) & 63)))))
                switch (v6 << ((42) & 63)) {
                case 951:
                    v28 := GenS1{F0: 42, F1: (^(v1)), F2: v6}
                    _ = v28
                case 42:
                    fmt.Println(v6, v5)
                default:
                    var v29 int16 = (v3 | (v3 | v3))
                    _ = v29
                }
                _ = v27
            }
        }
    }
    fmt.Println(v5, v6)
    v30 := GenS1{F0: 2, F1: (v1 & (v1 << ((genFn1(2, 26)) & 63))), F2: (^(v6))}
    var v31 float64 = -2.25
    v4 = append(v4, (v6 >> ((893) & 63)))
    var v32 uint32 = 2147483648
    switch -1 {
    case -211:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 944:
        v3 += v3
    default:
        var v33 uint16 = 42
        _ = v33
    }
    var v34 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
