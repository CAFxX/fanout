package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 uint8) int32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        switch (fi5 ^ 128) {
        case -9223372036854775808:
            p2 += (p2 << ((255) & 63))
        }
        for fi6 := 0; fi6 < 17; fi6++ {
            if (-2.25 == 0.0) {
                switch "" {
                case "":
                    var v7 int64 = 42
                    _ = v7
                case "hello":
                    v8 := GenS1{F0: true, F1: 32767}
                    _ = v8
                default:
                    var v9 string = "foo bar"
                    _ = v9
                }
            }
            var v10 int32 = -1
            var v11 float64 = 1.5
            _ = fi6
            _ = v10
            _ = v11
        }
        _ = fi5
    }
    v12 := []int{242, 32767, 127}
    for fi13 := 0; fi13 < 15; fi13++ {
        if true {
            v14 := GenS1{F0: true, F1: 0}
            p0 -= (p0 & p0)
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: false, F1: 1000}
            var v17 int64 = 0
            p2 += (p2 % ((p2) | 1))
            v18 := GenS2{F0: uint16(p2), F1: true}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    var v19 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v19
    return v19
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 uint64 = 2
    var v4 string = ""
    switch -9223372036854775808 {
    case -249:
        var v20 int32 = -857
        _ = v20
    case -9223372036854775808:
        if true {
            fmt.Println(v4, v3, v4)
            for fi21 := 0; fi21 < 2; fi21++ {
                for fi22 := 0; fi22 < 7; fi22++ {
                    v3 = (v3 >> ((1) & 63))
                    v23 := GenS2{F0: v2, F1: (-32 < -1)}
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
            v2 -= 2
        }
    }
    fmt.Println(v4)
    var v24 string = (v4 + v4)
    v25 := GenS1{F0: (-55 < 42), F1: -1}
    if false {
        v26 := []int{0, 128}
        switch v24 {
        case "foo bar":
            switch v24 {
            case "":
                v24 = v4
            }
        }
        _ = v26
    }
    fmt.Println(v24, v24)
    v27 := GenS3{F0: 19, F1: 0, F2: (((1 != 6) || false) || true)}
    if (255 <= 2) {
        var v28 uint = 255
        v3 += v3
        _ = v28
    } else {
        v24 = ""
        var v29 uint8 = 2
        switch -9223372036854775808 {
        case 42:
            v30 := v1[0]
            _ = v30
        case 190:
            v31 := GenS4{F0: (v3 << ((45) & 63))}
            _ = v31
        case 9223372036854775807:
            v1[uint32(v29)] = v24
        default:
            var v32 string = ((v24 + v4) + (v4 + v24))
            _ = v32
        }
        _ = v29
    }
    v2 = 32767
    v4 = v24
    if true {
        _ = len(v1)
        fmt.Println(v4, v24)
        if (-766 >= 1) {
            v33 := GenS3{F0: -128, F1: v2, F2: false}
            _ = v33
        } else {
            v34 := v1[1]
            v35 := v1[0]
            _ = v34
            _ = v35
        }
    } else {
        switch -195 {
        case 42:
            v36 := v1[4294967295]
            _ = v36
        default:
            fmt.Println(v24, v3)
        }
        delete(v1, 2147483647)
    }
    v2 += v2
    for fi37 := 0; fi37 < 20; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            v24 = v4
            if false {
                v24 = "hello"
            }
            var v39 int32 = 2147483647
            v39 -= 42
            _ = fi38
            _ = v39
        }
        v4 = "foo bar"
        _ = fi37
    }
    v40 := []int{}
    v40 = append(v40, 32767)
    var v41 int16 = 42
    v3 = uint64(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
