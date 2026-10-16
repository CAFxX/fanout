package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 int32; F2 int }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64, p1 bool) uint64 {
    fmt.Println(p0, p0)
    var v6 int16 = 865
    for fi7 := 0; fi7 < 3; fi7++ {
        v6 = v6
        v8 := GenS1{F0: fi7, F1: (p0 * (p0 + p0))}
        var v9 uint = 4294967295
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = 2147483647
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{0, 919, -1, 128}
    v3 := []int{1}
    var v4 string = "a"
    var v5 uint8 = 128
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch v4 {
        case "":
            v5 -= v5
        }
    case 528:
        var v11 int = 473
        _ = v11
    case -112:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 float64 = 3.14159
            switch (v4 + "hello") {
            case "":
                v14 := GenS2{F0: -1, F1: (-(fi12))}
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 int = -509
                    v1 = 9223372036854775807
                    v17 := GenS3{F0: (-(fi12)), F1: 914, F2: 0}
                    v16 = v16
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case "hello":
                if (true || true) {
                    v18 := GenS1{F0: (fi12 | -2147483648), F1: 0}
                    _ = v18
                } else {
                    _ = len(v3)
                    var v19 int64 = (v1 + v1)
                    v20 := GenS4{F0: v1}
                    _ = v19
                    _ = v20
                }
            default:
                v2 = append(v2, fi12)
            }
            _ = len(v2)
            _ = fi12
            _ = v13
        }
    }
    v21 := GenS1{F0: 128, F1: 18446744073709551615}
    v4 = v4
    if false {
        var v22 string = v4
        fmt.Println(v5, v22, v5)
        v4 = ("" + (v22 + v22))
        _ = v22
    } else {
        var v23 string = (v4 + ("" + "a"))
        v1 += (^(v1))
        _ = v23
    }
    switch 912 {
    case -966:
        var v24 bool = (893 > 2147483647)
        _ = v24
    case -1:
        fmt.Println(v5, v4)
    }
    v3 = append(v3, 724)
    for fi25 := 0; fi25 < 9; fi25++ {
        v4 = v4
        for fi26 := 0; fi26 < 18; fi26++ {
            if true {
                if true {
                    if (float32((3.14159 / -2.25)) < 0.1) {
                        v2 = append(v2, fi25)
                    }
                    fmt.Println(fi26)
                    switch -81 {
                    case 4294967295:
                        if false {
                            fmt.Println(v5, fi25, fi26)
                        }
                        if false {
                            v1 -= (v1 ^ -898)
                            v27 := GenS4{F0: (-(v1))}
                            _ = v27
                        }
                    default:
                        switch v4 {
                        case "foo bar":
                            v5 = v5
                        case "":
                            v28 := GenS2{F0: -80, F1: -9223372036854775808}
                            _ = v28
                        case "hello":
                            var v29 uint64 = 18446744073709551615
                            _ = v29
                        }
                    }
                }
                var v30 uint32 = 65535
                var v31 int16 = 815
                _ = v30
                _ = v31
            }
            v1 = 42
            var v32 int64 = (-(v1))
            _ = fi26
            _ = v32
        }
        fmt.Println(v1)
        _ = fi25
    }
    if len(v2) > 0 {
        v2[0] = -78
    }
    switch v4 {
    case "a":
        switch 555 {
        case -594:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 127:
            var v33 string = v4
            _ = v33
        case 32768:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    if (!true) {
        switch 663 {
        case -9223372036854775808:
            v5 = 128
        }
    }
    var v34 uint = 42
    v5 = v5
    v5 = v5
    var v35 int16 = int16(v5)
    var v36 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
