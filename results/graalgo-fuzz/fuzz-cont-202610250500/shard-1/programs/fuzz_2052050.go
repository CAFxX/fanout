package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8) int64 {
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            p0 = 13
            v6 := []int{(fi5 | int(p0)), (fi5 % ((int(p0)) | 1))}
            v7 := GenS1{F0: 127, F1: 62, F2: 9223372036854775807}
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{2147483647, 127, 75, -2147483648}
        _ = v8
    }
    p0 = p0
    v9 := []int{2147483647, -1, -9223372036854775808}
    _ = p0
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 42
    var v2 bool = false
    var v3 bool = false
    var v4 uint = 255
    var v10 string = "a"
    if (((3.14159 / 1e10) != 1e10) && v3) {
        v3 = false
    }
    v11 := GenS1{F0: -38, F1: 29, F2: genFn0(2)}
    v12 := GenS2{F0: 1, F1: (v1 | (v1 >> ((v4) & 63)))}
    v13 := GenS1{F0: int8(v1), F1: 8, F2: -1}
    v2 = true
    v14 := []int{658, 42}
    v2 = (0.0 != -2.25)
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int16 = 51
        _ = fi15
        _ = v16
    }
    var v17 int16 = -32768
    v10 = (((v10 + (v10 + v10)) + (v10 + (v10 + v10))) + v10)
    v18 := GenS3{F0: -1, F1: v17}
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS4{F0: 63, F1: 2}
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v3, v4, v1)
        case "hello":
            switch v10 {
            case "foo bar":
                for fi21 := 0; fi21 < 1; fi21++ {
                    v3 = v3
                    var v22 uint64 = 284
                    var v23 int8 = 127
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case "":
                v1 -= (v1 << ((uint(v1)) & 63))
            case "a":
                fmt.Println(fi19, fi19, v2)
            }
        case "\n\t":
            if v3 {
                switch "hello" {
                case "":
                    v2 = true
                case "\n\t":
                    var v24 int16 = (v17 << (((v4 % (((^(v4))) | 1))) & 63))
                    _ = v24
                }
            }
        }
        var v25 int = fi19
        _ = fi19
        _ = v20
        _ = v25
    }
    var v26 int64 = -1
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS4{F0: 495, F1: 563}
        if (32767 <= v17) {
            if false {
                if v2 {
                    if (!true) {
                        var v29 bool = v2
                        v30 := GenS5{F0: "\n\t", F1: v26}
                        _ = v29
                        _ = v30
                    } else {
                        var v31 uint16 = v1
                        _ = v31
                    }
                    var v32 int8 = 42
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        v34 := GenS1{F0: 0, F1: 255, F2: v26}
                        var v35 int32 = int32(v26)
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    fmt.Println(fi27)
                    v14 = append(v14, fi27)
                }
                if v2 {
                    v36 := GenS6{F0: false, F1: 165}
                    fmt.Println(v10, v17)
                    _ = v36
                }
                v2 = ((42 < 210) || v3)
            }
            var v37 float64 = (float64((((-2.25 * -2.25) + -2.25) + 1e10)) + (1e10 - float64((0.0 * (3.14159 * 1.5)))))
            for fi38 := 0; fi38 < 7; fi38++ {
                v39 := GenS7{F0: v1}
                var v40 float64 = (1e10 / v37)
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = v37
        }
        var v41 bool = (v3 || false)
        _ = fi27
        _ = v28
        _ = v41
    }
    v1 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
