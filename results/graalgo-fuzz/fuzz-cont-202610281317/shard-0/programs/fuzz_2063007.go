package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0(p0 uint16, p1 bool, p2 bool) string {
    var v6 uint = 0
    if ((("a" + ("foo bar" + "")) < (("" + (("hello" + ((("" + ("a" + "hello")) + ((("" + (("" + ("foo bar" + "foo bar")) + "")) + (("a" + ("a" + "\n\t")) + "foo bar")) + "a")) + "a")) + "\n\t")) + "hello")) || p2) {
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                p2 = false
            }
            _ = fi7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            p0 = (p0 * (p0 / ((p0) | 1)))
            v9 := GenS1{F0: 4294967295}
            var v10 int8 = 127
            _ = fi8
            _ = v10
            _ = v9
        }
        v6 = (v6 * (v6 & 841))
    } else {
        var v11 float64 = 0.1
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = v6
            v15 := GenS2{F0: -9223372036854775808, F1: 52}
            fmt.Println(v14, fi12, fi12)
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v6, p0)
        if (2 <= (p0 >> ((v6) & 63))) {
            p2 = p1
            var v16 uint16 = (p0 << (((v6 ^ (v6 & 65535))) & 63))
            _ = v16
        } else {
            if (-32768 > -1) {
                p1 = p1
            } else {
                v17 := GenS1{F0: 460}
                switch ("\n\t" + "") {
                case "\n\t":
                    v18 := GenS2{F0: 167, F1: 2}
                    _ = v18
                case "foo bar":
                    v19 := []int{9223372036854775807, (fi12 | fi12), (fi12 % ((562) | 1))}
                    _ = v19
                case "hello":
                    fmt.Println(v6, p0)
                default:
                    var v20 int64 = -887
                    _ = v20
                }
                v6 -= 4294967295
                _ = v17
            }
            p2 = false
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 float64 = 1.5
                var v23 uint64 = 18446744073709551615
                v24 := []int{(-(fi12)), (fi21 << ((v6) & 63))}
                if (fi12 != (^(fi21))) {
                    v23 += (v23 % (((v23 - uint64(p0))) | 1))
                    p0 = (p0 / (((p0 + 42)) | 1))
                }
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        }
        _ = fi12
    }
    fmt.Println(p1, v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}
func genFn1(p0 uint8) uint32 {
    switch "" {
    case "hello":
        v25 := []int{42}
        _ = v25
    case "\n\t":
        if ((42 != 42) && true) {
            v26 := []int{-926, -447, -284, -1}
            fmt.Println(p0)
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint64 = 1
        _ = fi27
        _ = v28
    }
    _ = p0
    return 65535
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := []int{418, 127, -1}
    fmt.Println(v1)
    v1 = ((v1 + -2.25) * (v1 - 0.0))
    var v29 uint32 = 65535
    _ = len(v3)
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS1{F0: 4294967295}
        _ = fi30
        _ = v31
    }
    if true {
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 uint = 65535
            _ = fi32
            _ = v33
        }
        v29 = (^(v29))
        v1 -= (v1 - 3.14159)
    } else {
        for fi34 := 0; fi34 < 5; fi34++ {
            v1 = v1
            _ = fi34
        }
    }
    _ = len(v5)
    var v35 int32 = 0
    var v36 int = -945
    if (true && true) {
        v36 = (-(v36))
    } else {
        fmt.Println(v29)
    }
    v29 = genFn1(211)
    if (!true) {
        switch -1 {
        case 9223372036854775807:
            if false {
                v37 := GenS3{F0: 181, F1: (v29 - (v29 / ((v29) | 1))), F2: 1}
                var v38 float32 = float32(-2.25)
                _ = v37
                _ = v38
            }
        default:
            _ = len(v5)
        }
        var v39 float64 = (float64((0.0 * (float32(-2.25) + float32(v1)))) * float64((0.1 + 0.0)))
        _ = v39
    } else {
        v35 += v35
        _ = len(v5)
    }
    v5 = append(v5, v36)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
