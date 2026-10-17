package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint16, p1 bool, p2 uint64) uint16 {
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int64 = 9223372036854775807
        if (2147483648 != 437) {
            var v10 uint32 = 65535
            fmt.Println(fi8)
            var v11 float64 = float64(((float32(1e10) * (0.0 - -2.25)) + 0.0))
            _ = v10
            _ = v11
        } else {
            switch "a" {
            case "hello":
                var v12 uint64 = 18446744073709551615
                _ = v12
            case "a":
                if (!p1) {
                    var v13 float32 = 1.5
                    _ = v13
                }
            case "\n\t":
                fmt.Println(p0, p0, fi8)
            }
            var v14 int64 = v9
            _ = v14
        }
        v15 := GenS1{F0: 0}
        _ = fi8
        _ = v15
        _ = v9
    }
    v16 := []int{127, 1, -1}
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: 42, F1: 100, F2: (fi17 % ((128) | 1))}
        for fi19 := 0; fi19 < 18; fi19++ {
            fmt.Println(p2, v7)
            _ = fi19
        }
        fmt.Println(p2)
        _ = fi17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    return (p0 * p0)
}
func genFn1(p0 int16, p1 int16, p2 uint16) int32 {
    v20 := GenS2{F0: -72, F1: 127, F2: -9223372036854775808}
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: 42}
        fmt.Println(p1)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := []int{fi23, (^(fi23))}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            switch "foo bar" {
            case "foo bar":
                var v27 int16 = (p0 / (((-(p1))) | 1))
                _ = v27
            default:
                v28 := []int{(fi26 >> ((570) & 63)), (fi26 << ((4294967295) & 63))}
                _ = v28
            }
            v29 := GenS3{F0: 128, F1: "\n\t"}
            _ = fi26
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            p1 -= p1
            p2 += genFn0(42, (18446744073709551615 >= 42), 0)
            _ = fi30
        }
        _ = fi25
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v32 := []int{(^(fi31)), fi31, fi31, 32768}
        switch (fi31 >> ((1) & 63)) {
        case 42:
            var v33 uint8 = 255
            _ = v33
        case 1:
            var v34 uint8 = 128
            _ = v34
        default:
            var v35 int32 = 301
            _ = v35
        }
        v36 := GenS2{F0: -1, F1: 1, F2: 4294967295}
        _ = fi31
        _ = v32
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 42
}

func main() {
    v1 := []uint32{2, 926}
    v2 := []int{}
    v3 := []string{""}
    var v4 uint = 0
    var v5 int32 = -935
    v6 := make(map[uint32]string)
    v5 += v5
    var v37 int32 = (v5 / ((0) | 1))
    v5 += (^(v5))
    fmt.Println(v4, v4)
    var v38 string = ((("" + "\n\t") + "") + "")
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 int = (fi40 << ((0) & 63))
            fmt.Println(v5, v37, v4)
            _ = fi40
            _ = v41
        }
        v38 = "a"
        _ = fi39
    }
    if true {
        for fi42 := 0; fi42 < 18; fi42++ {
            v37 = (v5 >> ((v4) & 63))
            if true {
                var v43 int32 = -522
                _ = v43
            }
            v44 := GenS4{F0: fi42, F1: v37}
            _ = fi42
            _ = v44
        }
    } else {
        var v45 int8 = 42
        var v46 uint16 = 0
        if false {
            if false {
                for fi47 := 0; fi47 < 14; fi47++ {
                    delete(v6, uint32(v4))
                    var v48 int32 = 0
                    _ = fi47
                    _ = v48
                }
                for fi49 := 0; fi49 < 3; fi49++ {
                    var v50 uint16 = v46
                    _ = fi49
                    _ = v50
                }
                v37 += (v37 * v5)
            }
        } else {
            v51 := GenS5{F0: genFn0(genFn0((v46 * 32767), false, 42), (false || true), 2)}
            _ = v51
        }
        _ = v45
        _ = v46
    }
    var v52 int16 = 526
    v53 := GenS6{F0: 894}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
