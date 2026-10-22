package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []int{2147483647, 0, -1}
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 int = -636
    var v5 float64 = 0.0
    var v6 uint32 = 65535
    var v7 uint16 = 170
    var v8 int32 = 54
    v2[("hello" + ("foo bar" + "foo bar"))] = (v6 << ((0) & 63))
    delete(v3, 127)
    for fi9 := 0; fi9 < 7; fi9++ {
        if (!false) {
            v10 := GenS1{F0: true, F1: -9223372036854775808, F2: 127}
            _ = v10
        }
        v11 := GenS1{F0: true, F1: 10, F2: -1}
        var v12 int = (fi9 + fi9)
        _ = fi9
        _ = v11
        _ = v12
    }
    var v13 uint32 = uint32(v7)
    for fi14 := 0; fi14 < 18; fi14++ {
        if true {
            v6 = v6
            if false {
                if (-2.25 <= 1e10) {
                    v15 := GenS2{F0: ""}
                    _ = v15
                } else {
                    var v16 string = "\n\t"
                    _ = v16
                }
                var v17 int8 = 100
                var v18 float32 = 1e10
                _ = v17
                _ = v18
            } else {
                var v19 uint = 0
                fmt.Println(v13, v19)
                _ = v19
            }
        } else {
            var v20 int = v4
            v21 := GenS2{F0: "\n\t"}
            v13 += v6
            _ = v20
            _ = v21
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v5)
        case "a":
            fmt.Println(v13)
        }
        var v22 uint = 1
        _ = fi14
        _ = v22
    }
    v5 += float64((-2.25 + (-2.25 + (-2.25 / 3.14159))))
    for fi23 := 0; fi23 < 1; fi23++ {
        if true {
            if (18446744073709551615 >= 18446744073709551615) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (916 == 11) {
                        v8 = 995
                        fmt.Println(fi24, v5, fi24)
                        v6 += v13
                    } else {
                        v5 = v5
                        v25 := v3[2]
                        _ = v25
                    }
                    _ = fi24
                }
                v13 = (v13 * (v6 << ((597) & 63)))
            }
        }
        _ = fi23
    }
    v5 = v5
    var v26 string = "\n\t"
    fmt.Println(v7)
    switch "" {
    case "":
        switch ((((v26 + v26) + (v26 + v26)) + v26) + ("foo bar" + v26)) {
        case "foo bar":
            if true {
                fmt.Println(v8, v4, v6)
                var v27 int64 = 1
                var v28 float64 = v5
                _ = v27
                _ = v28
            } else {
                fmt.Println(v5, v6, v5)
                if ((v13 - (v13 - 0)) != v13) {
                    switch (v4 % (((v4 & v4)) | 1)) {
                    case 4294967295:
                        v29 := GenS3{F0: 29}
                        _ = v29
                    case 2147483647:
                        switch (v4 | (v4 + 127)) {
                        case 32767:
                            v30 := GenS1{F0: (42 < 117), F1: 127, F2: -1}
                            _ = v30
                        }
                    default:
                        var v31 int64 = -71
                        _ = v31
                    }
                    fmt.Println(v5, v7)
                    if true {
                        fmt.Println(v6)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                switch (v26 + ("a" + v26)) {
                case "a":
                    v32 := GenS3{F0: 0}
                    _ = v32
                }
            }
        default:
            v8 += int32(v7)
        }
    case "\n\t":
        for fi33 := 0; fi33 < 15; fi33++ {
            for fi34 := 0; fi34 < 12; fi34++ {
                if true {
                    v4 = (v4 / (((v4 >> ((2) & 63))) | 1))
                    v4 = (fi33 ^ v4)
                    _ = len(v1)
                    v13 += 719
                    if ((!(v7 == (^(v7)))) || (!(int64(v8) > -1))) {
                        v7 -= 77
                        var v35 int8 = 95
                        _ = v35
                    } else {
                        v3[-1] = (v6 * (v13 % ((v13) | 1)))
                        var v36 float64 = float64(0.0)
                        if len(v1) > 0 {
                            v1[0] = (fi33 >> ((2) & 63))
                        }
                        _ = v36
                    }
                }
                _ = len(v1)
                _ = fi34
            }
            _ = fi33
        }
    case "a":
        for fi37 := 0; fi37 < 8; fi37++ {
            v38 := v3[503]
            _ = fi37
            _ = v38
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
