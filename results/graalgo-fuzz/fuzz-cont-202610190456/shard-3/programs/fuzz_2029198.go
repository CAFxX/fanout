package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
func genFn0() bool {
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    switch -1 {
    case 1:
        v6 := []int{9223372036854775807, 32768}
        _ = v6
    }
    if true {
        var v7 float64 = 3.14159
        switch 42 {
        case -2147483648:
            v8 := GenS1{F0: 344}
            _ = v8
        default:
            v9 := GenS1{F0: 65535}
            _ = v9
        }
        _ = v7
    } else {
        var v10 int8 = 9
        var v11 float64 = 0.0
        if (float32(((v11 - 3.14159) / (0.0 / v11))) == (float32((0.0 - (v11 + (((float64(float32(-2.25)) + -2.25) - v11) / v11)))) * (-2.25 + (0.1 / 1e10)))) {
            v11 += v11
            var v12 uint16 = 42
            _ = v12
        } else {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 uint = 255
                if true {
                    v15 := []int{-822, 0}
                    _ = v15
                } else {
                    var v16 float32 = 1.5
                    fmt.Println(fi13, v14)
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            v17 := []int{}
            v10 += -25
            _ = v17
        }
        _ = v10
        _ = v11
    }
    return (217 <= 42)
}
func genFn1(p0 bool, p1 int16, p2 string) uint8 {
    var v18 int32 = 2147483647
    var v19 int64 = -213
    for fi20 := 0; fi20 < 1; fi20++ {
        if (0 < 507) {
            switch (fi20 % (((fi20 >> ((944) & 63))) | 1)) {
            case -1:
                v21 := GenS2{F0: p0}
                _ = v21
            case -2147483648:
                v22 := GenS2{F0: (!false)}
                _ = v22
            }
            v23 := GenS3{F0: v19}
            _ = v23
        } else {
            v24 := GenS2{F0: (!(true && (true && false)))}
            _ = v24
        }
        fmt.Println(v19, v19)
        _ = fi20
    }
    var v25 int32 = (v18 % ((v18) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v25
    return 237
}

func main() {
    v1 := []string{}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v26 := v3[56]
    var v27 uint64 = 799
    if (-1 == 472) {
        fmt.Println(v27)
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 int16 = 32767
            _ = fi28
            _ = v29
        }
    }
    _ = len(v1)
    v27 = (^(v27))
    if true {
        for fi30 := 0; fi30 < 11; fi30++ {
            v26 -= 32768
            _ = fi30
        }
        switch ("\n\t" + "\n\t") {
        case "a":
            if true {
                switch (v26 >> ((0) & 63)) {
                case 32768:
                    if true {
                        var v31 uint32 = 158
                        var v32 uint16 = 2
                        var v33 int64 = -837
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case 0:
                    var v34 int32 = -1
                    _ = v34
                }
                var v35 int = v26
                v27 = v27
                _ = v35
            }
        case "hello":
            delete(v3, 26)
        case "foo bar":
            var v36 uint16 = 2
            _ = v36
        default:
            var v37 int64 = 779
            _ = v37
        }
        var v38 uint8 = 43
        _ = v38
    } else {
        v39 := GenS1{F0: 42}
        v27 = (v27 + v27)
        _ = v39
    }
    v40 := v2[(v26 * (v26 << ((219) & 63)))]
    v26 = 0
    switch "a" {
    case "\n\t":
        var v41 int32 = 590
        _ = v41
    case "hello":
        var v42 float32 = 1e10
        _ = v42
    case "":
        v40 = (v40 % ((uint32(v27)) | 1))
    default:
        for fi43 := 0; fi43 < 18; fi43++ {
            var v44 uint64 = 1
            var v45 bool = true
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
}
