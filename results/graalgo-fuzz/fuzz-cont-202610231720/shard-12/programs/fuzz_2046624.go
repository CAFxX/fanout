package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int16, p1 uint8) int8 {
    p0 = p0
    v6 := GenS1{F0: -472}
    if (0 > 42) {
        var v7 uint = 1
        _ = v7
    } else {
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                p1 -= 195
                var v9 int16 = -1
                _ = fi8
                _ = v9
            }
            var v10 uint16 = 42
            _ = v10
        } else {
            for fi11 := 0; fi11 < 18; fi11++ {
                switch fi11 {
                case -764:
                    v12 := []int{}
                    _ = v12
                default:
                    var v13 int16 = (p0 * (p0 >> ((255) & 63)))
                    _ = v13
                }
                p0 = (p0 - (p0 % (((p0 ^ p0)) | 1)))
                var v14 uint8 = p1
                _ = fi11
                _ = v14
            }
            p0 = p0
        }
        v15 := GenS1{F0: 2147483647}
        _ = v15
    }
    p1 += (p1 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 int32, p1 uint16) uint16 {
    fmt.Println(p1, p1)
    v16 := []int{42, 32767, 32768, -9223372036854775808}
    var v17 int16 = 1
    for fi18 := 0; fi18 < 17; fi18++ {
        v16 = append(v16, (^(fi18)))
        _ = fi18
    }
    var v19 float64 = 0.1
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v19
    return 42
}

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 255
    var v3 uint = 18446744073709551615
    v4 := make(map[string]bool)
    var v5 bool = ((1.5 < 1e10) && (0.1 > -2.25))
    if (0 == 1) {
        switch "\n\t" {
        case "\n\t":
            var v20 uint = v3
            _ = v20
        }
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        v5 = (-1 <= -1)
    case "hello":
        var v21 int64 = 1
        _ = v21
    case "":
        fmt.Println(v3, v1, v2)
    default:
        v22 := []int{}
        _ = v22
    }
    v23 := GenS1{F0: -742}
    switch "\n\t" {
    case "foo bar":
        v24 := GenS1{F0: -72}
        _ = v24
    case "":
        if v5 {
            v1 = (v1 * v1)
            var v25 int8 = -46
            v26 := v4["a"]
            _ = v25
            _ = v26
        } else {
            var v27 uint32 = 2
            _ = v27
        }
    default:
        fmt.Println(v2, v3, v5)
    }
    v2 = v2
    if (!false) {
        var v28 uint16 = v2
        _ = v28
    } else {
        v29 := v4["foo bar"]
        v3 += v3
        _ = v29
    }
    v30 := GenS1{F0: 127}
    v3 = 679
    var v31 string = "a"
    switch v31 {
    case "hello":
        if v5 {
            v2 = 449
            v32 := GenS1{F0: 2147483647}
            v33 := []int{-2147483648, 2147483647, 4294967295, 32767}
            _ = v32
            _ = v33
        }
    case "foo bar":
        v2 = genFn1(42, genFn1(42, genFn1(0, 360)))
    default:
        for fi34 := 0; fi34 < 7; fi34++ {
            if ((-1 != -1) || true) {
                v2 -= (^(v2))
                var v35 int64 = 9223372036854775807
                _ = v35
            }
            _ = fi34
        }
    }
    if false {
        if (v5 || ((v31 + v31) == (v31 + "foo bar"))) {
            switch 4294967295 {
            case 4294967295:
                var v36 uint8 = (^(v1))
                _ = v36
            case 127:
                var v37 int64 = 444
                _ = v37
            default:
                for fi38 := 0; fi38 < 3; fi38++ {
                    v2 = 255
                    _ = fi38
                }
            }
        } else {
            var v39 int16 = 0
            fmt.Println(v31, v2, v3)
            _ = v39
        }
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 int16 = 42
            v3 += (v3 % (((v3 % (((v3 & v3)) | 1))) | 1))
            _ = fi40
            _ = v41
        }
    }
    v2 -= v2
    if v5 {
        var v42 uint32 = 597
        v5 = v5
        _ = v42
    }
    v4["a"] = (584 != 944)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
