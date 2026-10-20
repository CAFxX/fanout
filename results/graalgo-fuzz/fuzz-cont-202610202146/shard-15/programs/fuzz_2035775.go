package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 string }
func genFn0(p0 uint64, p1 uint8, p2 uint32) uint32 {
    v7 := []int{128, 4294967295}
    p2 -= (^(p2))
    var v8 int16 = -32768
    if true {
        switch -2147483648 {
        case 0:
            v9 := GenS1{F0: (368 != 1), F1: 32767}
            _ = v9
        case -2147483648:
            p0 -= p0
        case 822:
            var v10 uint8 = (^(p1))
            _ = v10
        default:
            if true {
                v11 := GenS1{F0: ((!false) || (255 > uint16(p2))), F1: 434}
                _ = v11
            }
        }
        var v12 uint64 = 1
        var v13 string = ("" + "a")
        _ = v12
        _ = v13
    } else {
        p1 = (p1 | p1)
        var v14 bool = ((true || (((((!false) || (!false)) && (false && false)) || false) || (("foo bar" + "") == "hello"))) && true)
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint = 156
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int) string {
    p0 = p0
    p0 = (p0 >> ((2) & 63))
    p0 = p0
    _ = p0
    return ("\n\t" + "\n\t")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    v6 := []int{2147483647, 42}
    for fi17 := 0; fi17 < 13; fi17++ {
        v3 -= (v3 - (v3 * 444))
        _ = fi17
    }
    v3 -= (v3 / ((941) | 1))
    _ = len(v1)
    v3 = (^(v3))
    var v18 uint = (v3 | v3)
    var v19 int8 = 42
    var v20 string = "hello"
    for fi21 := 0; fi21 < 10; fi21++ {
        if (true || true) {
            var v22 uint32 = 4294967295
            var v23 float32 = 1e10
            _ = v22
            _ = v23
        } else {
            switch -1 {
            case 1:
                var v24 bool = (true && false)
                _ = v24
            default:
                v25 := v4[339]
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v18 = 587
            switch "foo bar" {
            case "\n\t":
                v18 = (v3 / ((247) | 1))
                var v27 float64 = (0.1 + (float64(float32((-2.25 * -2.25))) + -2.25))
                switch (-(fi26)) {
                case 2147483647:
                    var v28 int64 = -1
                    _ = v28
                case 349:
                    var v29 uint64 = 209
                    _ = v29
                case 128:
                    fmt.Println(v20, v3, fi26)
                default:
                    var v30 uint8 = 46
                    _ = v30
                }
                _ = v27
            default:
                if false {
                    var v31 uint16 = 42
                    if (!(float32(0.0) == 3.14159)) {
                        if ((!false) || false) {
                            var v32 int = fi26
                            v18 = v18
                            _ = v32
                        }
                    }
                    v33 := GenS1{F0: ((!false) || (293 > 854)), F1: 0}
                    _ = v31
                    _ = v33
                } else {
                    v34 := GenS1{F0: false, F1: 65535}
                    fmt.Println(fi26)
                    v20 = (("a" + v20) + v20)
                    var v35 int32 = 0
                    _ = v34
                    _ = v35
                }
            }
            _ = fi26
        }
        v36 := GenS2{F0: v20, F1: v18, F2: v20}
        _ = fi21
        _ = v36
    }
    var v37 bool = (-261 > 42)
    var v38 int32 = 42
    _ = len(v2)
    v18 = v3
    if len(v6) > 0 {
        v6[0] = -723
    }
    v39 := v4[882]
    v38 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
