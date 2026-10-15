package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0(p0 int64, p1 int16, p2 int32) int {
    var v5 string = "foo bar"
    var v6 uint32 = 378
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}
func genFn1() bool {
    switch genFn0(-621, 0, 42) {
    case -429:
        var v7 int8 = 127
        _ = v7
    case 2147483647:
        var v8 int64 = 0
        _ = v8
    case -169:
        var v9 int64 = -9223372036854775808
        _ = v9
    default:
        var v10 uint64 = 1
        _ = v10
    }
    switch "a" {
    case "foo bar":
        if (!(2 <= 255)) {
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: 1000, F1: (-(fi11))}
                if false {
                    v13 := []int{-221, -555, (^(fi11))}
                    if true {
                        fmt.Println(fi11)
                    }
                    _ = v13
                } else {
                    var v14 float32 = 3.14159
                    _ = v14
                }
                var v15 float32 = (((1.5 / 1.5) * 3.14159) - -2.25)
                _ = fi11
                _ = v12
                _ = v15
            }
            switch 32767 {
            case 9223372036854775807:
                var v16 int16 = 1000
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if true {
            var v18 uint32 = 522
            fmt.Println(fi17, v18)
            v19 := GenS2{F0: (-(fi17))}
            _ = v18
            _ = v19
        }
        if (9223372036854775807 < 0) {
            var v20 uint64 = 4294967295
            var v21 uint32 = 4294967295
            if (false || (-947 != 1)) {
                v22 := GenS3{F0: 42}
                if ((0.1 / 0.1) >= 0.0) {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS2{F0: fi17}
                        v25 := []int{fi23, (fi23 << ((818) & 63))}
                        v26 := GenS4{F0: 65535, F1: (fi23 == (fi23 + fi23))}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    if (!(false || (18446744073709551615 != 0))) {
                        var v27 int8 = 127
                        switch (fi17 * (fi17 << ((493) & 63))) {
                        case -2147483648:
                            var v28 uint16 = 65535
                            _ = v28
                        default:
                            v29 := []int{(^(fi17)), 1}
                            _ = v29
                        }
                        var v30 uint = 18446744073709551615
                        _ = v27
                        _ = v30
                    }
                    fmt.Println(v20, v20)
                }
                v31 := []int{fi17, fi17}
                _ = v22
                _ = v31
            } else {
                for fi32 := 0; fi32 < 12; fi32++ {
                    switch fi17 {
                    case -2147483648:
                        v33 := GenS2{F0: 221}
                        _ = v33
                    case -2:
                        v34 := []int{genFn0(-17, 0, 1), (fi32 >> ((0) & 63)), (fi17 + (fi17 ^ (fi17 * fi17)))}
                        _ = v34
                    }
                    v21 = (v21 * v21)
                    var v35 uint = 42
                    v36 := GenS4{F0: 255, F1: true}
                    v37 := []int{(fi17 - (fi32 * genFn0(-9223372036854775808, 811, 42)))}
                    _ = fi32
                    _ = v35
                    _ = v36
                    _ = v37
                }
                fmt.Println(fi17, fi17)
            }
            _ = v20
            _ = v21
        }
        _ = fi17
    }
    return true
}

func main() {
    var v1 uint = 42
    var v2 string = "hello"
    var v3 int64 = -9223372036854775808
    var v4 string = v2
    v38 := []int{-236, 0}
    v39 := GenS1{F0: -1, F1: 361}
    var v40 uint8 = 0
    var v41 uint64 = 42
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 int32 = 1
        var v44 int64 = 42
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 uint16 = uint16(v41)
    v3 += (v3 * (v3 ^ (v3 ^ 711)))
    var v46 int8 = 42
    var v47 int16 = -32768
    v48 := GenS5{F0: uint32(v45), F1: v47}
    v49 := GenS2{F0: 956}
    fmt.Println(v41, v46)
    _ = len(v38)
    v50 := GenS6{F0: v46, F1: v45, F2: v40}
    v2 = ""
    fmt.Println(v3)
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
