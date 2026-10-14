package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int32 = 716
    v5 := []string{v1, v1, (v1 + "\n\t"), "hello"}
    var v6 uint8 = 205
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float64 = 1e10
        v4 = v4
        switch "a" {
        case "\n\t":
            v8 -= (v8 + v8)
        default:
            v5 = append(v5, (v1 + (("" + "\n\t") + v1)))
        }
        _ = fi7
        _ = v8
    }
    if ((981 != 2) && true) {
        var v9 int8 = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = 42
        v11 = (v11 << ((2) & 63))
        var v12 int = (fi10 << ((255) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 float32 = -2.25
    var v14 float32 = float32(-2.25)
    switch "foo bar" {
    case "":
        var v15 uint16 = 0
        _ = v15
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v13, v13, v14)
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint = 764
        switch "a" {
        case "foo bar":
            fmt.Println(v6)
        case "hello":
            var v18 int64 = -208
            _ = v18
        case "":
            if (((3.14159 * 0.1) + -2.25) == float64(v13)) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    switch ((v1 + (v1 + ("hello" + ""))) + v1) {
                    case "foo bar":
                        if (0 <= 83) {
                            fmt.Println(v4, v14)
                        } else {
                            var v20 int = fi16
                            _ = v20
                        }
                    }
                    v14 = (v14 + v14)
                    v21 := GenS1{F0: 0}
                    _ = fi19
                    _ = v21
                }
            }
        default:
            for fi22 := 0; fi22 < 9; fi22++ {
                fmt.Println(v13)
                var v23 uint8 = (v6 - 255)
                if false {
                    switch (("foo bar" + (v1 + "")) + v1) {
                    case "a":
                        var v24 float64 = 1e10
                        _ = v24
                    case "\n\t":
                        v17 = (v17 | (v17 * v17))
                    default:
                        var v25 string = (v1 + v1)
                        _ = v25
                    }
                    var v26 string = "foo bar"
                    var v27 int16 = 247
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS2{F0: 878, F1: 127, F2: 0}
                    v13 = ((v14 / ((float32(1e10) / v14) - float32(1.5))) * v13)
                    var v29 int = int(v4)
                    _ = v28
                    _ = v29
                }
                _ = fi22
                _ = v23
            }
        }
        _ = fi16
        _ = v17
    }
    v4 += (v4 | (v4 >> ((861) & 63)))
    if (v6 <= (^(v6))) {
        for fi30 := 0; fi30 < 3; fi30++ {
            if true {
                v31 := GenS2{F0: 0, F1: 1, F2: 110}
                var v32 uint64 = 18446744073709551615
                _ = v31
                _ = v32
            }
            v1 = (v1 + (v1 + "hello"))
            v33 := GenS1{F0: 42}
            _ = fi30
            _ = v33
        }
        v34 := GenS2{F0: 1, F1: 1, F2: 311}
        _ = v34
    } else {
        v1 = v1
        _ = len(v3)
        if (v4 <= v4) {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 uint = 2
                v13 -= 3.14159
                v37 := GenS3{F0: (!(fi35 <= (fi35 / (((fi35 & 995)) | 1)))), F1: uint32(v36), F2: 18446744073709551615}
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    }
    fmt.Println(v6, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
