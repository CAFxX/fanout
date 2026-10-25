package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 bool, p2 int) uint8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: ((((((("\n\t" + ("hello" + "foo bar")) + "") + (("" + "") + ("hello" + ""))) + "foo bar") + "") + ("a" + "foo bar")) + ("foo bar" + "hello"))}
        _ = fi5
        _ = v6
    }
    var v7 float64 = 1e10
    v8 := []int{(p2 - p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int64, p1 string) int8 {
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            if (p1 < p1) {
                switch (fi10 ^ fi9) {
                case 366:
                    var v11 float32 = ((0.0 * 0.1) + 0.0)
                    _ = v11
                default:
                    v12 := GenS2{F0: 42, F1: 845, F2: int64(fi9)}
                    fmt.Println(fi9, p0, fi9)
                    var v13 bool = false
                    _ = v12
                    _ = v13
                }
                var v14 int8 = -128
                _ = v14
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        p1 = "foo bar"
        _ = fi15
    }
    v16 := GenS2{F0: -2147483648, F1: 2, F2: (p0 % ((p0) | 1))}
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := []int{9223372036854775807, (fi17 + (-(fi17))), (fi17 + (fi17 + 4294967295)), fi17}
        switch fi17 {
        case 4294967295:
            v18 = append(v18, (fi17 / ((fi17) | 1)))
        case 128:
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float64 = float64(-2.25)
                var v21 uint16 = 125
                var v22 string = ""
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS1{F0: p1}
                switch fi17 {
                case -2147483648:
                    var v25 int32 = 241
                    _ = v25
                case -1:
                    if (-754 >= 1) {
                        if false {
                            p1 = "a"
                            var v26 float64 = 1.5
                            _ = v26
                        }
                    }
                case -206:
                    var v27 int32 = 1
                    _ = v27
                }
                var v28 int8 = 100
                _ = fi23
                _ = v24
                _ = v28
            }
        }
        fmt.Println(fi17, p1)
        _ = fi17
        _ = v18
    }
    var v29 uint = 1
    _ = p0
    _ = p1
    _ = v16
    _ = v29
    return 1
}

func main() {
    var v1 uint8 = 101
    var v2 uint64 = 290
    v3 := []uint32{2147483647, 65535}
    v4 := make(map[int]int)
    fmt.Println(v2, v2)
    switch (("" + "foo bar") + "hello") {
    case "a":
        v30 := GenS1{F0: ""}
        _ = v30
    case "hello":
        fmt.Println(v1, v2)
    case "\n\t":
        var v31 int8 = 127
        _ = v31
    default:
        if (945 <= 559) {
            v32 := GenS2{F0: 1, F1: 255, F2: -772}
            if true {
                var v33 uint32 = 211
                if false {
                    var v34 float64 = (0.0 - float64(0.1))
                    _ = v34
                }
                _ = v33
            }
            var v35 int16 = 42
            _ = v32
            _ = v35
        }
    }
    var v36 int16 = 42
    v1 = v1
    v36 -= (v36 & 0)
    var v37 int32 = 0
    for fi38 := 0; fi38 < 18; fi38++ {
        if (true || (!false)) {
            if true {
                v37 = 42
            }
        }
        _ = fi38
    }
    v1 -= (^(v1))
    var v39 uint8 = 1
    v36 = (^(v36))
    v39 = v1
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v40 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
}
