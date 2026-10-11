package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 string, p1 int16, p2 uint32) uint {
    fmt.Println(p1)
    switch -291 {
    case -2147483648:
        switch (p0 + p0) {
        case "hello":
            p1 -= (p1 * p1)
        case "foo bar":
            v7 := GenS1{F0: 42}
            _ = v7
        case "":
            p1 += p1
        default:
            v8 := []int{42, 479, -398}
            _ = v8
        }
    case -112:
        switch 42 {
        case 42:
            v9 := []int{-1}
            _ = v9
        }
    case 719:
        p1 += 320
    default:
        p0 = ("\n\t" + "\n\t")
    }
    if (!(!(-128 < 100))) {
        p0 = "foo bar"
        v10 := []int{42, 2147483647, 128}
        _ = v10
    } else {
        v11 := []int{-2147483648, 1, 32768}
        var v12 uint32 = p2
        switch (p0 + "\n\t") {
        case "a":
            var v13 float32 = (1.5 * (0.1 * 3.14159))
            _ = v13
        case "":
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS1{F0: 32767}
                    var v17 int = (-(fi15))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        }
        _ = v11
        _ = v12
    }
    v18 := GenS1{F0: 255}
    if (!(42 > 128)) {
        var v19 int64 = 770
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 37
}

func main() {
    var v1 bool = (65535 != 435)
    v2 := make(map[string]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{2147483648, 1, 559}
    var v5 float64 = 1e10
    var v6 int = 577
    fmt.Println(v5, v3)
    _ = len(v2)
    var v20 string = ""
    var v21 int32 = 1
    v5 += 3.14159
    if (false || (((v1 || v1) || (681 < v3)) || v1)) {
        fmt.Println(v6, v6, v3)
    }
    v1 = (0 >= genFn0(v20, -769, 674))
    var v22 int32 = v21
    var v23 uint16 = 1
    delete(v2, v20)
    fmt.Println(v5, v20)
    fmt.Println(v6, v6)
    switch "\n\t" {
    case "a":
        var v24 int32 = v21
        _ = v24
    case "hello":
        switch "" {
        case "a":
            if v1 {
                var v25 float64 = (float64((0.0 - 1.5)) - (float64(0.1) - v5))
                var v26 bool = v1
                _ = v25
                _ = v26
            } else {
                switch 972 {
                case 9223372036854775807:
                    if true {
                        _ = len(v4)
                        v20 = v20
                        v27 := GenS1{F0: 577}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v1 = (1e10 != 0.1)
                            var v29 int = v6
                            _ = fi28
                            _ = v29
                        }
                    }
                }
                var v30 float32 = float32(v5)
                var v31 int32 = (v22 >> ((18446744073709551615) & 63))
                _ = v30
                _ = v31
            }
        case "hello":
            for fi32 := 0; fi32 < 20; fi32++ {
                if (v1 && true) {
                    var v33 uint = 65535
                    _ = v33
                } else {
                    v34 := GenS2{F0: 128, F1: 100, F2: 0}
                    _ = v34
                }
                _ = fi32
            }
        case "":
            v35 := GenS3{F0: (-(v21)), F1: 1}
            _ = v35
        default:
            for fi36 := 0; fi36 < 12; fi36++ {
                v20 = v20
                _ = fi36
            }
        }
    default:
        var v37 int16 = -1
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        delete(v2, v20)
        if true {
            delete(v2, v20)
        }
        v39 := GenS2{F0: v6, F1: 100, F2: 42}
        _ = fi38
        _ = v39
    }
    v1 = (true && (0 < 18446744073709551615))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
