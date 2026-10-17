package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint16; F1 int8; F2 int32 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint32 }
func genFn0() uint32 {
    switch 144 {
    case 127:
        if false {
            v7 := GenS1{F0: 255, F1: 0}
            var v8 bool = (false || true)
            _ = v7
            _ = v8
        }
    case -2147483648:
        switch 1 {
        case 1:
            var v9 bool = (false || (!(783 < -132)))
            _ = v9
        }
    }
    if (4294967295 != -2147483648) {
        v10 := []int{-372, 2147483647, 1, 127}
        var v11 int32 = 2147483647
        _ = v10
        _ = v11
    } else {
        var v12 int16 = -225
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(v12, fi13)
            var v14 uint8 = 255
            _ = fi13
            _ = v14
        }
        if true {
            if (18446744073709551615 >= 904) {
                if true {
                    v15 := GenS1{F0: 4294967295, F1: 1}
                    v12 += (v12 + v12)
                    switch ("foo bar" + "hello") {
                    case "hello":
                        v16 := GenS2{F0: 1000, F1: 127, F2: 1}
                        _ = v16
                    case "a":
                        for fi17 := 0; fi17 < 17; fi17++ {
                            var v18 uint32 = 4294967295
                            _ = fi17
                            _ = v18
                        }
                    default:
                        if ((805 != 42) && (((!false) && (0 <= 0)) || (!false))) {
                            var v19 float64 = 0.1
                            v20 := GenS1{F0: 0, F1: -1}
                            _ = v19
                            _ = v20
                        }
                    }
                    _ = v15
                } else {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v22 := []int{128, 32768}
                        var v23 string = ("" + ("a" + "\n\t"))
                        var v24 uint8 = 2
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    var v25 uint = 18446744073709551615
                    _ = v25
                }
                if true {
                    v12 = 0
                }
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(fi26)
                v27 := GenS3{F0: true}
                _ = fi26
                _ = v27
            }
        } else {
            if false {
                v28 := []int{32768, -234}
                _ = v28
            } else {
                v29 := []int{1, 1, 42, 602}
                v30 := GenS4{F0: 2147483647}
                v12 = v12
                _ = v29
                _ = v30
            }
        }
        _ = v12
    }
    return 0
}

func main() {
    var v1 int = 127
    v2 := []uint32{197, 42, 0}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 32767
    v5 := []string{"foo bar", ("foo bar" + (((("a" + ("hello" + "foo bar")) + "foo bar") + "") + "hello"))}
    var v6 int32 = 2147483647
    if ((v4 ^ 32767) < v4) {
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 int = v1
            _ = fi31
            _ = v32
        }
        var v33 int64 = -1
        v6 -= (v6 * v6)
        _ = v33
    }
    var v34 uint16 = v4
    v35 := GenS5{F0: 9223372036854775807, F1: 36}
    var v36 uint64 = 4294967295
    v4 = v34
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := GenS6{F0: (v4 - v4), F1: 114, F2: (v6 - -1)}
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: 1000, F1: v1, F2: 0}
    var v40 float32 = float32((1e10 - 0.0))
    v41 := GenS7{F0: 42, F1: (1 == 54)}
    v1 += (v1 & (v1 - v1))
    var v42 int16 = 1
    v36 += v36
    v1 = v1
    switch int(v6) {
    case 127:
        v43 := GenS8{F0: 65535}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
