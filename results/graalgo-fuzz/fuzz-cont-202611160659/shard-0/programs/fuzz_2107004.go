package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool, p1 int8, p2 int) uint8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        p2 = -2147483648
        switch "hello" {
        case "foo bar":
            if (true && ((((((("hello" + "foo bar") + "a") + "foo bar") + "") + "a") + ("\n\t" + ("\n\t" + ("a" + "foo bar")))) > "\n\t")) {
                v6 := []int{(p2 << ((18446744073709551615) & 63)), 128}
                _ = v6
            }
        case "":
            fmt.Println(p0, fi5, p0)
        case "\n\t":
            var v7 uint = 4294967295
            _ = v7
        default:
            if true {
                switch "" {
                case "a":
                    var v8 uint32 = 925
                    _ = v8
                default:
                    p0 = p0
                }
                var v9 float64 = 0.0
                _ = v9
            } else {
                v10 := GenS1{F0: 9223372036854775807, F1: 1, F2: 255}
                v11 := []int{(^(fi5)), (-(fi5))}
                switch 32768 {
                case 297:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        fmt.Println(p0, fi12)
                        p1 = p1
                        fmt.Println(p0, fi5)
                        _ = fi12
                    }
                case -1:
                    v13 := GenS1{F0: 42, F1: 831, F2: 127}
                    _ = v13
                case 897:
                    p0 = ((p1 + p1) < (-(p1)))
                }
                _ = v10
                _ = v11
            }
        }
        _ = fi5
    }
    var v14 float64 = (1.5 * (1e10 - (0.0 * float64((float32(1e10) / 3.14159)))))
    for fi15 := 0; fi15 < 5; fi15++ {
        p1 += 43
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 211
    v18 := []int{42, (p2 | (^(p2)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    _ = v18
    return 0
}
func genFn1(p0 uint) int64 {
    v19 := GenS1{F0: -777, F1: 675, F2: 255}
    var v20 uint64 = 2
    p0 -= (p0 | (p0 / ((p0) | 1)))
    _ = p0
    _ = v19
    _ = v20
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 4294967295
    var v3 int32 = -917
    v4 := []uint32{(v2 >> ((896) & 63)), (^(v2)), (v2 * (^(v2))), (^(v2))}
    var v21 int64 = 1
    var v22 int = 729
    v21 += v21
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v22, fi23)
            v25 := GenS1{F0: 0, F1: 909, F2: 1}
            v21 += int64(v2)
            if (false && (int8(v2) != -46)) {
                v2 = v2
                var v26 int = 128
                v22 = (fi24 >> ((309) & 63))
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        fmt.Println(v2, v3, v22)
        _ = fi23
    }
    v27 := GenS2{F0: v3, F1: 0}
    var v28 uint = 314
    for fi29 := 0; fi29 < 2; fi29++ {
        switch "\n\t" {
        case "":
            v4 = append(v4, 813)
        case "foo bar":
            fmt.Println(v2, v28)
        }
        _ = fi29
    }
    switch ((((("\n\t" + "foo bar") + ("hello" + ("" + ""))) + "\n\t") + (("" + (("a" + ("hello" + ((("a" + ("hello" + "foo bar")) + "hello") + "\n\t"))) + "foo bar")) + "\n\t")) + (((("hello" + ("hello" + "a")) + (("" + ("a" + "a")) + "a")) + ("hello" + (("hello" + "foo bar") + ("\n\t" + "\n\t")))) + "a")) {
    case "foo bar":
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 int = (^(v22))
            _ = fi30
            _ = v31
        }
    default:
        v32 := GenS3{F0: 42, F1: v3, F2: (!((127 < -128) || ((float64(-2.25) / 1.0) >= (((1.5 + 3.14159) * 3.14159) / 1.0))))}
        _ = v32
    }
    v3 = -387
    if ((-(v22)) < v22) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 float64 = float64(0.0)
            _ = fi33
            _ = v34
        }
    }
    switch (v22 & 4294967295) {
    case 32767:
        switch v22 {
        case -652:
            v35 := GenS4{F0: v2}
            _ = v35
        case 343:
            var v36 int = (v22 & (-(v22)))
            _ = v36
        case 42:
            var v37 string = "a"
            _ = v37
        }
    case 127:
        v38 := GenS3{F0: 1, F1: -1, F2: (!false)}
        _ = v38
    case 877:
        v39 := GenS4{F0: 1}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
