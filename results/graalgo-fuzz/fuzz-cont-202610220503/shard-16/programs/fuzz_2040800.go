package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 uint16) int {
    var v7 int64 = -576
    var v8 uint16 = (p1 | (p1 + 42))
    v8 = 326
    v8 = (v8 | v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []int{42}
    var v2 string = ("" + ("a" + ("\n\t" + "hello")))
    v3 := make(map[uint32]string)
    var v4 int32 = 0
    v5 := []int{128, 127, -488, 0}
    var v6 uint64 = 718
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := v3[2147483648]
        for fi11 := 0; fi11 < 2; fi11++ {
            switch v2 {
            case "":
                fmt.Println(v10, fi9)
            case "foo bar":
                switch "" {
                case "a":
                    var v12 uint16 = 1
                    _ = v12
                case "":
                    switch (v2 + ("" + "a")) {
                    case "\n\t":
                        fmt.Println(fi11, fi11, v6)
                    default:
                        v3[2147483647] = (v2 + (v10 + "hello"))
                    }
                }
            case "\n\t":
                _ = len(v5)
            default:
                fmt.Println(v2)
                switch (v10 + v2) {
                case "\n\t":
                    v13 := GenS1{F0: true}
                    _ = v13
                case "":
                    if (107 < 127) {
                        fmt.Println(fi11, v4, fi11)
                        v14 := GenS1{F0: false}
                        _ = v14
                    }
                case "a":
                    v15 := GenS1{F0: false}
                    _ = v15
                default:
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                }
                v16 := GenS1{F0: false}
                if len(v5) > 0 {
                    v5[0] = (fi11 * fi11)
                }
                _ = v16
            }
            if true {
                var v17 bool = false
                v18 := GenS1{F0: false}
                v4 += (v4 / (((v4 % (((v4 - v4)) | 1))) | 1))
                _ = v17
                _ = v18
            }
            v19 := GenS1{F0: true}
            _ = fi11
            _ = v19
        }
        _ = fi9
        _ = v10
    }
    v4 += v4
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    var v22 int32 = -24
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            if ((v6 & 636) != (v6 + 280)) {
                v25 := GenS1{F0: (!(false || ("" > "hello")))}
                if len(v5) > 0 {
                    v5[0] = (fi24 << ((42) & 63))
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v25
            }
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println(v22)
    for fi26 := 0; fi26 < 9; fi26++ {
        if true {
            var v27 float32 = 3.14159
            var v28 float64 = 1.5
            var v29 uint = 651
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 uint = 65535
            _ = v30
        }
        _ = fi26
    }
    var v31 string = ((v2 + "foo bar") + v2)
    var v32 int16 = 32767
    if (false && (0 < 100)) {
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 uint = 0
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            if (v2 < "a") {
                switch (("foo bar" + v2) + v2) {
                case "a":
                    v22 = (-(v4))
                case "hello":
                    v1 = append(v1, fi35)
                case "\n\t":
                    var v36 float32 = 1e10
                    _ = v36
                }
            } else {
                v4 = 42
                var v37 float32 = 3.14159
                _ = v37
            }
            fmt.Println(v22, v6)
            _ = fi35
        }
    }
    v31 = v2
    switch v31 {
    case "\n\t":
        var v38 int8 = -103
        _ = v38
    case "a":
        for fi39 := 0; fi39 < 2; fi39++ {
            var v40 int64 = -1
            v1 = append(v1, -1)
            v41 := v3[456]
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    var v42 int = 0
    if len(v1) > 0 {
        v1[0] = genFn0(v6, 2)
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    fmt.Println(v31, v22, v22)
    var v43 uint = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
