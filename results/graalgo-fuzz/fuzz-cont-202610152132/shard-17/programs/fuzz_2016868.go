package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint8 = 1
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    var v4 uint8 = v1
    var v5 uint8 = v4
    var v6 float64 = (1e10 / 1.5)
    v1 = (v1 & 128)
    var v7 int32 = -1
    var v8 int16 = 852
    if (uint32(v2) <= 2147483647) {
        switch 4294967295 {
        case 32768:
            v4 = v5
        case 861:
            v6 = v6
        case 32767:
            fmt.Println(v8, v5)
        default:
            var v9 float32 = 3.14159
            _ = v9
        }
        var v10 bool = (v2 < (v2 << ((v2) & 63)))
        var v11 uint8 = v5
        _ = v10
        _ = v11
    } else {
        v4 = 65
        for fi12 := 0; fi12 < 16; fi12++ {
            if false {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := []int{(fi13 >> ((v2) & 63)), (fi13 - (fi12 - fi12))}
                    var v15 string = ""
                    switch "a" {
                    case "":
                        if (!(0 >= 9223372036854775807)) {
                            v5 = v4
                            v16 := GenS1{F0: 127}
                            var v17 float64 = -2.25
                            _ = v16
                            _ = v17
                        }
                    default:
                        fmt.Println(fi13, v15, v15)
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v18 := GenS1{F0: 49}
                if true {
                    v7 = (v7 & (^(v7)))
                }
                _ = v18
            }
            v8 -= v8
            _ = fi12
        }
        var v19 int64 = int64(v7)
        _ = v19
    }
    var v20 uint8 = (v5 - 127)
    if (439 < 42) {
        var v21 uint16 = 42
        _ = v21
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            v8 = v8
            v23 := GenS1{F0: 93}
            v20 = 2
            _ = fi22
            _ = v23
        }
        v5 += (v1 << ((v2) & 63))
    }
    v1 = v5
    v2 += 144
    v24 := GenS2{F0: 42, F1: 9223372036854775807}
    v25 := []int{-1, 127, -677, 42}
    for fi26 := 0; fi26 < 4; fi26++ {
        if false {
            fmt.Println(v2, v20, fi26)
        }
        _ = fi26
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    if true {
        v27 := v3["hello"]
        v25 = append(v25, 737)
        v28 := GenS3{F0: (-(v8)), F1: (^(v8)), F2: (v1 | (v5 | (v1 / (((^(v1))) | 1))))}
        _ = v27
        _ = v28
    } else {
        v29 := GenS1{F0: 1}
        if (v7 != 698) {
            if true {
                v30 := GenS4{F0: (("" + (((("\n\t" + "\n\t") + "") + ("hello" + (((("" + "") + "a") + "\n\t") + ""))) + ("hello" + "\n\t"))) + "foo bar"), F1: (^(v7)), F2: (v5 | (v4 * (v5 + v4)))}
                v31 := GenS5{F0: 1, F1: v4}
                v32 := GenS6{F0: 65535}
                _ = v30
                _ = v31
                _ = v32
            }
            var v33 int32 = 2147483647
            if true {
                for fi34 := 0; fi34 < 1; fi34++ {
                    switch "a" {
                    case "hello":
                        var v35 int8 = 1
                        _ = v35
                    }
                    var v36 int = 128
                    _ = fi34
                    _ = v36
                }
                v4 = v5
            } else {
                v37 := v3[("a" + "\n\t")]
                if (((210 <= v2) || (v7 <= (v7 & v33))) || (639 < 130)) {
                    var v38 float32 = -2.25
                    v20 += v5
                    _ = v38
                }
                _ = v37
            }
            _ = v33
        }
        _ = v29
    }
    var v39 float64 = 0.0
    if (-1 >= 9223372036854775807) {
        fmt.Println(v2, v5)
        if true {
            var v40 float64 = v39
            v20 -= 42
            _ = v40
        } else {
            var v41 uint8 = v20
            var v42 bool = ((v39 - (-2.25 + ((v6 * v39) / (1.5 / (v39 - v6))))) >= v6)
            _ = v41
            _ = v42
        }
        v5 = (v20 / ((v5) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
