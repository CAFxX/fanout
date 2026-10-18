package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int32; F2 int }

func main() {
    var v1 int32 = -2147483648
    v2 := []int{1, 32767, -9223372036854775808}
    var v3 string = "a"
    v4 := []int{42, 127, 48}
    var v5 uint8 = 2
    var v6 int8 = 0
    v5 -= 40
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: (!false)}
        v9 := GenS1{F0: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 828 {
    case -1:
        var v10 int64 = -9223372036854775808
        _ = v10
    case 4294967295:
        v3 = (v3 + v3)
    }
    v2 = append(v2, 0)
    var v11 int32 = v1
    v12 := GenS1{F0: true}
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS2{F0: 0}
        v1 = v1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch v3 {
        case "foo bar":
            v5 = 145
        case "hello":
            if ((!true) && (-169 < 835)) {
                v16 := GenS3{F0: (v11 / (((v1 & v11)) | 1)), F1: (("a" + v3) + v3), F2: 606}
                fmt.Println(v1, fi15)
                _ = v16
            } else {
                var v17 int64 = -1
                switch (fi15 * (fi15 ^ fi15)) {
                case 127:
                    var v18 uint = 4294967295
                    _ = v18
                case -628:
                    v19 := GenS4{F0: 0}
                    _ = v19
                default:
                    var v20 float32 = 0.1
                    _ = v20
                }
                v11 = (v11 * (v11 + 42))
                _ = v17
            }
        case "":
            var v21 float32 = 1.5
            _ = v21
        default:
            if (true && (!(0.1 >= 1e10))) {
                v22 := GenS1{F0: false}
                v6 = v6
                for fi23 := 0; fi23 < 5; fi23++ {
                    switch (fi15 & fi23) {
                    case 4294967295:
                        var v24 float64 = 1e10
                        _ = v24
                    }
                    v5 += v5
                    fmt.Println(fi15, v3, v5)
                    _ = fi23
                }
                _ = v22
            }
        }
        switch (^(fi15)) {
        case 32767:
            v25 := GenS5{F0: 4294967295, F1: v5}
            _ = v25
        case 9223372036854775807:
            if (832 <= 0) {
                v11 = v1
                for fi26 := 0; fi26 < 18; fi26++ {
                    if false {
                        v6 += (v6 / ((v6) | 1))
                    }
                    _ = fi26
                }
                var v27 float64 = 0.1
                _ = v27
            } else {
                var v28 bool = false
                v11 += 481
                _ = v28
            }
        default:
            v3 = (v3 + ((v3 + (("" + (v3 + "hello")) + (v3 + v3))) + v3))
        }
        v2 = append(v2, fi15)
        _ = fi15
    }
    v3 = "foo bar"
    if false {
        if (v5 == v5) {
            v3 = (v3 + v3)
            v1 -= (-(v1))
        }
        v2 = append(v2, 1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v6 = 0
        if false {
            v5 -= 255
            v5 += (v5 | (v5 & v5))
            v3 = (v3 + "\n\t")
        }
    }
    v5 = (^(v5))
    var v29 uint64 = 697
    var v30 int64 = 221
    v31 := GenS6{F0: (v29 >> ((497) & 63)), F1: v1, F2: 32767}
    fmt.Println(v3)
    var v32 int64 = v30
    fmt.Println(v5, v5, v32)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
