package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 int }
func genFn0(p0 string, p1 int16) int {
    var v7 uint64 = 0
    switch "" {
    case "foo bar":
        for fi8 := 0; fi8 < 13; fi8++ {
            if (2 >= 4294967295) {
                var v9 bool = false
                v10 := GenS1{F0: 0}
                if (-468 != -9223372036854775808) {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v9 = (!(v9 || false))
                        fmt.Println(fi11)
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 14; fi12++ {
                        v13 := GenS1{F0: v7}
                        var v14 int64 = 0
                        var v15 int8 = -1
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                } else {
                    if (v9 && v9) {
                        fmt.Println(p0, fi8, fi8)
                        v16 := GenS2{F0: 0}
                        _ = v16
                    }
                    v17 := GenS3{F0: 870, F1: (v9 || v9), F2: fi8}
                    var v18 string = ((("" + p0) + p0) + p0)
                    _ = v17
                    _ = v18
                }
                _ = v10
                _ = v9
            } else {
                var v19 int16 = p1
                if (-2.25 == -2.25) {
                    v20 := GenS3{F0: (fi8 - 4294967295), F1: false, F2: (fi8 | (fi8 << ((18446744073709551615) & 63)))}
                    _ = v20
                } else {
                    switch 128 {
                    case -432:
                        p1 = 772
                    }
                }
                var v21 float64 = 0.1
                _ = v19
                _ = v21
            }
            var v22 int16 = (p1 >> ((255) & 63))
            switch -9223372036854775808 {
            case -2147483648:
                v7 = v7
            default:
                switch 127 {
                case 759:
                    fmt.Println(v7, fi8)
                default:
                    fmt.Println(fi8, p0, p0)
                }
            }
            _ = fi8
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            p1 = (p1 | (^(p1)))
            _ = fi24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v7 += (^(v7))
        if true {
            var v26 uint8 = 0
            _ = v26
        }
        v7 = v7
        _ = fi25
    }
    _ = p0
    _ = p1
    _ = v7
    return -270
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{("" + "hello"), "", (("hello" + ("" + "")) + ("hello" + ("hello" + ("hello" + "foo bar")))), (("hello" + ("hello" + "\n\t")) + "a")}
    var v3 uint = 129
    v4 := make(map[uint32]string)
    var v5 int32 = 42
    v6 := []int{265, 32768, 818}
    fmt.Println(v3, v5)
    fmt.Println(v5, v5)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v27 uint32 = 400
    var v28 uint32 = v27
    switch -2147483648 {
    case -466:
        if ((v5 - (v5 << (((v3 & 42)) & 63))) == v5) {
            for fi29 := 0; fi29 < 7; fi29++ {
                v27 += (v28 >> (((v3 + 212)) & 63))
                var v30 uint8 = 224
                _ = fi29
                _ = v30
            }
        } else {
            v27 = v27
        }
    }
    v31 := GenS2{F0: (v28 ^ 65535)}
    fmt.Println(v27, v3)
    if (!false) {
        switch 128 {
        case 128:
            v32 := GenS3{F0: 142, F1: true, F2: -606}
            _ = v32
        case 165:
            var v33 uint8 = 1
            _ = v33
        }
        fmt.Println(v28)
        fmt.Println(v28, v5)
    } else {
        var v34 uint = v3
        v35 := v1["foo bar"]
        _ = v34
        _ = v35
    }
    v28 += v28
    fmt.Println(v28)
    if (18446744073709551615 >= 258) {
        for fi36 := 0; fi36 < 7; fi36++ {
            v3 += v3
            v3 = v3
            var v37 uint32 = (v27 + v28)
            _ = fi36
            _ = v37
        }
        if false {
            v28 -= v27
            var v38 int32 = v5
            _ = v38
        }
    } else {
        fmt.Println(v5, v5)
        fmt.Println(v28)
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
