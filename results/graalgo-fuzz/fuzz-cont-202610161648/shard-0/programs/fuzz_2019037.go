package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }

func main() {
    var v1 bool = (!(!(18446744073709551615 == 728)))
    var v2 int64 = 489
    v3 := make(map[string]uint32)
    var v4 int8 = -128
    var v5 bool = true
    fmt.Println(v4)
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(v4, v2, v1)
        var v7 bool = v5
        _ = fi6
        _ = v7
    }
    if v5 {
        v8 := GenS1{F0: 0, F1: 42, F2: 42}
        v1 = v1
        var v9 float32 = (0.0 + float32((1.5 * 1.5)))
        _ = v8
        _ = v9
    } else {
        var v10 bool = (18446744073709551615 >= 42)
        switch "a" {
        case "\n\t":
            v11 := v3["hello"]
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 1, F1: 32767, F2: 18446744073709551615}
            _ = v12
        case "hello":
            fmt.Println(v5, v2, v4)
        }
        _ = v10
    }
    v1 = true
    _ = len(v3)
    if (4294967295 <= 4294967295) {
        if v5 {
            var v13 uint32 = 0
            var v14 int = 1
            switch 42 {
            case 0:
                v15 := GenS2{F0: (v13 << ((2) & 63)), F1: 42}
                _ = v15
            }
            _ = v13
            _ = v14
        } else {
            v16 := []int{2}
            var v17 bool = (0.0 == (((0.1 + 1e10) * float64(1.5)) / 1e10))
            if (!false) {
                var v18 uint8 = 1
                v4 = 127
                _ = v18
            } else {
                fmt.Println(v4, v2, v5)
                var v19 int8 = (v4 << ((65535) & 63))
                _ = v19
            }
            _ = v16
            _ = v17
        }
        v2 = v2
        v20 := GenS1{F0: 0, F1: -31, F2: 409}
        _ = v20
    } else {
        var v21 bool = v5
        v4 = v4
        v22 := []int{127}
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 2, F1: 65535}
    if ("" != "hello") {
        var v24 float32 = 3.14159
        for fi25 := 0; fi25 < 6; fi25++ {
            if v1 {
                switch "" {
                case "":
                    v26 := []int{fi25, (fi25 / ((735) | 1))}
                    _ = v26
                case "a":
                    if v5 {
                        v27 := []int{fi25, (^(fi25)), (fi25 % ((fi25) | 1)), 9223372036854775807}
                        _ = v27
                    } else {
                        var v28 uint16 = 32767
                        v2 = (v2 & (v2 >> ((18446744073709551615) & 63)))
                        _ = v28
                    }
                case "hello":
                    var v29 float64 = -2.25
                    _ = v29
                }
                v30 := GenS1{F0: 0, F1: (fi25 << ((741) & 63)), F2: 2}
                v1 = (false || false)
                _ = v30
            }
            if (!(float32(0.1) >= v24)) {
                v1 = false
                var v31 uint64 = 398
                _ = v31
            } else {
                switch "\n\t" {
                case "foo bar":
                    v32 := []int{fi25}
                    _ = v32
                case "":
                    v2 += 1
                default:
                    var v33 string = ("foo bar" + "\n\t")
                    _ = v33
                }
                var v34 uint = 18446744073709551615
                fmt.Println(v5, fi25, fi25)
                _ = v34
            }
            for fi35 := 0; fi35 < 1; fi35++ {
                var v36 int32 = 808
                v37 := []int{fi35, 127, (^(fi35))}
                if v1 {
                    v2 += int64(fi35)
                    v24 = v24
                } else {
                    fmt.Println(v36)
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi25
        }
        v4 -= -103
        _ = v24
    }
    v38 := []int{-1, 378, 62, 32768}
    for fi39 := 0; fi39 < 19; fi39++ {
        switch (fi39 << ((544) & 63)) {
        case -2147483648:
            v5 = true
        }
        v40 := GenS3{F0: (v4 ^ 87), F1: fi39, F2: 0}
        v41 := GenS2{F0: 42, F1: 0}
        _ = fi39
        _ = v40
        _ = v41
    }
    if (!v5) {
        if v1 {
            v3["foo bar"] = 2147483648
        }
        var v42 int = 32767
        _ = v42
    }
    fmt.Println(v5)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
