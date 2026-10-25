package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint16; F2 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 string, p1 int16) uint {
    p0 = ((p0 + (p0 + p0)) + p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint32 = 2147483648
        var v8 string = "foo bar"
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{-663, 42, 892}
    v9 = append(v9, 4294967295)
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(p0, p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := []string{("a" + "")}
    v2 := make(map[int]string)
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]int)
    var v5 int8 = -1
    fmt.Println(v5)
    v5 += -128
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int = 127
            v13 := GenS1{F0: 65535, F1: 32767, F2: 433}
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(fi14, fi14, fi11)
                var v15 uint32 = 810
                _ = fi14
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        v16 := GenS1{F0: 2147483647, F1: 42, F2: 42}
        fmt.Println(v3)
        _ = v16
    } else {
        v5 = 127
        var v17 bool = (0 >= 42)
        switch "hello" {
        case "a":
            v18 := GenS2{F0: 42}
            _ = v18
        default:
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS3{F0: false, F1: 255, F2: true}
                var v21 int16 = -32768
                var v22 string = ("foo bar" + ("\n\t" + ""))
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = v17
    }
    var v23 float64 = 1.5
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int8 = 100
        v5 -= 0
        var v26 int64 = 488
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS4{F0: 42}
    v5 = v5
    v3 += v3
    var v28 uint64 = 340
    fmt.Println(v5)
    var v29 uint32 = 1
    v2[-472] = "foo bar"
    if false {
        v28 = 620
        if (1 != 1) {
            v29 -= v29
            for fi30 := 0; fi30 < 2; fi30++ {
                switch "hello" {
                case "a":
                    v23 = v23
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        if (true && false) {
                            v31 := GenS1{F0: (v29 / ((2147483647) | 1)), F1: 307, F2: 42}
                            fmt.Println(v23, v28)
                            _ = v31
                        }
                    case "":
                        v23 -= v23
                    case "a":
                        if false {
                            v23 -= (v23 + v23)
                        }
                    }
                case "hello":
                    _ = len(v1)
                default:
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                }
                v4[(v29 * 2147483648)] = (-(fi30))
                if (4294967295 >= 1) {
                    var v32 uint32 = (^(v29))
                    _ = v32
                } else {
                    fmt.Println(v29, v29)
                    var v33 int32 = 1
                    for fi34 := 0; fi34 < 3; fi34++ {
                        v35 := GenS2{F0: v29}
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
                _ = fi30
            }
        }
    }
    if (v5 >= v5) {
        var v36 uint = 2
        _ = len(v4)
        _ = v36
    } else {
        v4[221] = 128
    }
    if (v28 >= (v28 << ((1) & 63))) {
        v28 = 401
        var v37 uint8 = uint8(v28)
        switch (-(v3)) {
        case 128:
            v38 := GenS4{F0: 32767}
            _ = v38
        case 2147483647:
            v39 := GenS3{F0: false, F1: 0, F2: (-1 != 0)}
            _ = v39
        default:
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
        _ = v37
    } else {
        v5 = v5
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v5)
        }
        v29 -= (v29 + (v29 ^ v29))
    }
    var v40 uint32 = 0
    for fi41 := 0; fi41 < 1; fi41++ {
        fmt.Println(v23)
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
