package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    var v3 uint16 = 255
    var v4 float64 = 0.0
    var v5 int = 437
    var v6 uint8 = uint8(v3)
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            if true {
                var v8 uint16 = (v3 & 255)
                _ = v8
            } else {
                var v9 string = ("a" + "\n\t")
                fmt.Println(v1)
                _ = v9
            }
            _ = fi7
        }
    } else {
        v10 := GenS1{F0: 1, F1: (v6 ^ (v6 & v6))}
        var v11 uint = 838
        _ = v10
        _ = v11
    }
    v6 -= (v6 % ((128) | 1))
    v3 += v3
    var v12 uint8 = v6
    var v13 uint8 = (v6 >> ((300) & 63))
    var v14 int16 = -1
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 float64 = (v4 / 1.0)
        if false {
            fmt.Println(v6, v4)
        }
        _ = fi15
        _ = v16
    }
    v13 = (v13 | 127)
    switch v5 {
    case 416:
        v17 := GenS2{F0: v1}
        _ = v17
    case -192:
        var v18 int = (v5 * v5)
        _ = v18
    }
    if (1 <= 1) {
        for fi19 := 0; fi19 < 9; fi19++ {
            switch "" {
            case "hello":
                v5 += 42
            }
            var v20 bool = (!((^(v1)) >= v1))
            _ = fi19
            _ = v20
        }
        fmt.Println(v13, v5)
    } else {
        var v21 int32 = 1
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint = 2
        _ = fi22
        _ = v23
    }
    v12 = 127
    v6 -= (v12 >> ((517) & 63))
    v6 = (v13 * (v12 % ((v6) | 1)))
    v24 := GenS1{F0: 2, F1: v12}
    v25 := GenS2{F0: v1}
    for fi26 := 0; fi26 < 11; fi26++ {
        if true {
            fmt.Println(v5)
        } else {
            v27 := []int{(-(v5)), (^(fi26)), (v5 / (((v5 % ((v5) | 1))) | 1))}
            for fi28 := 0; fi28 < 6; fi28++ {
                v14 = int16(v5)
                var v29 string = ("" + "a")
                v4 = 1.5
                v4 = float64(-2.25)
                fmt.Println(fi26, v12)
                v30 := GenS3{F0: (v14 << ((966) & 63)), F1: v14}
                fmt.Println(v12, v12)
                v31 := GenS2{F0: 1}
                var v32 float32 = float32((float64((1.5 + (((1e10 * float32(v4)) / (-2.25 - (-2.25 - 3.14159))) / float32(v4)))) + v4))
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            if len(v27) > 0 {
                v27[0] = (-(fi26))
            }
            _ = v27
        }
        switch "\n\t" {
        case "\n\t":
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS3{F0: (v14 << ((0) & 63)), F1: (v14 | v14)}
                if true {
                    v5 = (-(v5))
                    v3 -= uint16(v12)
                    var v35 int32 = (v1 | v1)
                    switch (v5 * (fi33 / ((v5) | 1))) {
                    case 625:
                        var v36 uint16 = (^(v3))
                        var v37 int64 = -1
                        var v38 float64 = float64(float32(v4))
                        var v39 int64 = (v37 | v37)
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    case 371:
                        var v40 float64 = v4
                        _ = v40
                    }
                    v41 := []int{}
                    fmt.Println(v3)
                    _ = v35
                    _ = v41
                }
                _ = fi33
                _ = v34
            }
        case "":
            v14 -= (-(v14))
        case "a":
            v42 := v2[(v5 / ((v5) | 1))]
            _ = v42
        }
        var v43 uint16 = 982
        _ = fi26
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
