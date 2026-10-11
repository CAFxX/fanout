package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }

func main() {
    var v1 uint32 = 879
    v2 := []int{0, 4294967295}
    v3 := make(map[int]uint32)
    var v4 uint16 = uint16(v1)
    var v5 uint8 = 171
    v1 += (v1 | 2)
    if (0.1 == ((float64((0.0 / (3.14159 + (1.5 / 1.0)))) * float64(0.0)) / 1.5)) {
        if ((((3.14159 / 1.0) + 0.0) != -2.25) && false) {
            fmt.Println(v5, v5)
            var v6 uint32 = (v1 << ((1) & 63))
            if true {
                v7 := GenS1{F0: v4}
                _ = v7
            } else {
                var v8 uint16 = 255
                _ = v8
            }
            _ = v6
        }
    }
    var v9 bool = false
    switch 2147483647 {
    case 128:
        for fi10 := 0; fi10 < 4; fi10++ {
            switch int(v4) {
            case 127:
                if (v4 < v4) {
                    var v11 uint16 = v4
                    _ = v11
                }
            case 2147483647:
                v5 += (v5 / (((v5 ^ (v5 << ((42) & 63)))) | 1))
            case -545:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v4 -= (v4 + 42)
                    var v13 int32 = -1
                    v2 = append(v2, (fi12 >> ((1) & 63)))
                    fmt.Println(v5, fi12)
                    if (1.5 != ((1e10 / 0.1) / (0.1 + ((-2.25 * (0.0 - (-2.25 / 1.0))) + 1.5)))) {
                        var v14 float32 = 0.0
                        _ = v14
                    } else {
                        delete(v3, (fi10 ^ fi10))
                    }
                    _ = fi12
                    _ = v13
                }
            default:
                var v15 int16 = int16(v5)
                _ = v15
            }
            var v16 float32 = 1e10
            v4 += v4
            _ = fi10
            _ = v16
        }
    case 470:
        for fi17 := 0; fi17 < 3; fi17++ {
            if v9 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v18 := GenS1{F0: 581}
                var v19 int16 = 1
                _ = v18
                _ = v19
            }
            var v20 uint32 = uint32(v4)
            fmt.Println(v5, v1)
            _ = fi17
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            if (42 == (v1 | (^(v1)))) {
                switch ((("hello" + "\n\t") + "\n\t") + "foo bar") {
                case "":
                    var v23 uint32 = (v1 / ((61) | 1))
                    _ = v23
                case "hello":
                    if false {
                        if true {
                            if (v9 || (!(!(v9 || (1.5 == (((0.0 + (-2.25 + 0.0)) + 3.14159) / 0.1)))))) {
                                if v9 {
                                    var v24 bool = true
                                    _ = v24
                                }
                            }
                            _ = len(v3)
                        }
                        var v25 int = fi21
                        var v26 int32 = 0
                        _ = v25
                        _ = v26
                    }
                case "\n\t":
                    v9 = v9
                }
                if false {
                    var v27 int8 = -1
                    _ = v27
                }
                var v28 int32 = 2147483647
                _ = v28
            }
            _ = fi22
        }
        _ = fi21
    }
    v29 := GenS2{F0: 100, F1: 670, F2: v4}
    var v30 float32 = (1e10 * 0.0)
    var v31 string = "hello"
    v9 = (-2147483648 == 2147483647)
    var v32 int64 = 0
    v5 = 0
    fmt.Println(v31)
    if (-2147483648 < 1) {
        v5 = (v5 * v5)
        var v33 uint8 = uint8(v4)
        var v34 float32 = v30
        _ = v33
        _ = v34
    }
    if true {
        v35 := GenS1{F0: v4}
        if v9 {
            var v36 uint64 = 1
            fmt.Println(v4)
            v31 = v31
            _ = v36
        } else {
            var v37 uint16 = v4
            _ = v37
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
