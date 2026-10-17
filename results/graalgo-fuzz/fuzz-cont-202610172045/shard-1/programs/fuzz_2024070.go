package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]string)
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    var v5 int = -1
    v3 = (v3 & v3)
    v3 += 2
    var v6 int16 = 953
    for fi7 := 0; fi7 < 1; fi7++ {
        if false {
            if false {
                var v8 int8 = 42
                v9 := []int{}
                _ = v8
                _ = v9
            } else {
                if (!(false || (!(fi7 < (v5 << ((1) & 63)))))) {
                    v5 = fi7
                    var v10 float64 = (float64(v1) + (3.14159 + (0.0 / 3.14159)))
                    var v11 uint = 613
                    _ = v10
                    _ = v11
                } else {
                    fmt.Println(fi7, v6)
                    var v12 int16 = v6
                    fmt.Println(v5, v5)
                    _ = v12
                }
            }
        } else {
            if false {
                v13 := GenS1{F0: 18446744073709551615, F1: v3}
                var v14 string = ("hello" + "hello")
                v15 := GenS2{F0: 127, F1: 977, F2: 2147483648}
                _ = v13
                _ = v14
                _ = v15
            }
            fmt.Println(v6)
        }
        fmt.Println(v3, fi7, v5)
        v6 = (v6 | v6)
        _ = fi7
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int64 = 1
        var v18 int64 = (-(v17))
        for fi19 := 0; fi19 < 2; fi19++ {
            fmt.Println(fi19, fi19, v3)
            var v20 float32 = (v1 - v1)
            switch (fi16 - -1) {
            case 493:
                if false {
                    v21 := GenS2{F0: 100, F1: 317, F2: 4294967295}
                    _ = v21
                } else {
                    if true {
                        v20 -= -2.25
                    }
                }
            case -2147483648:
                var v22 int16 = (^(v6))
                _ = v22
            }
            if (!((((false || ((0.1 - -2.25) >= float64((v20 - -2.25)))) || true) || false) && (true && true))) {
                fmt.Println(v3, v17)
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    fmt.Println(v18, v1)
                default:
                    v23 := []int{v5, (v5 + (fi16 | (v5 * fi16))), 9223372036854775807, fi16}
                    _ = v23
                }
                var v24 int16 = v6
                var v25 int64 = v17
                _ = v24
                _ = v25
            } else {
                if true {
                    v20 = v1
                    switch (v5 >> ((924) & 63)) {
                    case 4294967295:
                        var v26 uint8 = (v3 >> ((225) & 63))
                        _ = v26
                    case -273:
                        v18 = -9223372036854775808
                    case -236:
                        var v27 uint = 2
                        _ = v27
                    }
                    var v28 int = 32768
                    fmt.Println(v5)
                    _ = v28
                }
                var v29 uint = 1
                v30 := GenS3{F0: 898, F1: -1}
                _ = v29
                _ = v30
            }
            _ = fi19
            _ = v20
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v31 float32 = (v1 + 1.5)
    fmt.Println(v3)
    var v32 int64 = 939
    v4[v5] = 91
    v1 = v1
    var v33 uint8 = (^(v3))
    v34 := GenS2{F0: 1, F1: 32767, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
