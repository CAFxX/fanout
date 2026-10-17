package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 int64, p1 int32) int64 {
    var v4 int64 = p0
    if false {
        v5 := []int{127, -497}
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint16 = 484
        p0 -= v4
        _ = fi6
        _ = v7
    }
    p0 -= (v4 & (-(v4)))
    var v8 int8 = -128
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return p0
}
func genFn1(p0 string, p1 string) string {
    if true {
        var v9 uint64 = 42
        v9 = 42
        v10 := GenS1{F0: -582}
        _ = v10
        _ = v9
    }
    if (192 < 2147483647) {
        if (-2.25 <= 1.5) {
            var v11 uint64 = 705
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{}
                    var v15 uint = 205
                    p0 = p1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        }
        p1 = (p1 + p0)
    }
    _ = p0
    _ = p1
    return (p0 + (p1 + p1))
}

func main() {
    var v1 string = "a"
    var v2 int64 = -463
    var v3 uint = 42
    switch 127 {
    case 554:
        v3 = v3
    case 0:
        var v16 int16 = 32767
        _ = v16
    case 480:
        fmt.Println(v2, v3, v3)
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v1 = "a"
            _ = fi18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        fmt.Println(v1, v3, v1)
        switch -1 {
        case -1:
            fmt.Println(v1, v3, v2)
        case 4294967295:
            var v20 uint16 = 42
            _ = v20
        case -417:
            v21 := []int{(^(fi19)), fi19}
            _ = v21
        default:
            v1 = ""
        }
        _ = fi19
    }
    fmt.Println(v3, v2, v3)
    for fi22 := 0; fi22 < 20; fi22++ {
        if ((!(-83 < -128)) || (0 < 42)) {
            v1 = v1
            fmt.Println(v1, fi22)
            v23 := GenS2{F0: -1, F1: (-(v2))}
            _ = v23
        }
        var v24 int = (fi22 << ((255) & 63))
        v25 := GenS3{F0: 1}
        _ = fi22
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := []int{fi26}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 15; fi29++ {
            v30 := []int{}
            var v31 int32 = 0
            v32 := GenS4{F0: v1, F1: 4, F2: 4294967295}
            var v33 uint32 = 856
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi28
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 int = (fi34 / ((-1) | 1))
        _ = fi34
        _ = v35
    }
    var v36 uint16 = 255
    if (!(true || false)) {
        var v37 uint64 = 518
        v38 := []int{128}
        _ = v37
        _ = v38
    } else {
        if true {
            fmt.Println(v3)
        } else {
            fmt.Println(v36)
            v39 := []int{}
            var v40 int8 = int8(v2)
            _ = v39
            _ = v40
        }
        switch "a" {
        case "\n\t":
            var v41 int64 = (v2 / ((42) | 1))
            _ = v41
        }
    }
    switch -52 {
    case -1:
        var v42 int32 = 42
        _ = v42
    case 42:
        if (0 <= -1) {
            var v43 uint64 = 18446744073709551615
            switch ((v1 + "a") + v1) {
            case "hello":
                fmt.Println(v2)
            default:
                v1 = v1
            }
            _ = v43
        } else {
            v44 := []int{}
            if false {
                if false {
                    var v45 int16 = 45
                    var v46 uint64 = 1
                    v47 := GenS3{F0: 928}
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    _ = len(v44)
                    v48 := GenS4{F0: v1, F1: 128, F2: 1}
                    var v49 float32 = ((float32((float64(-2.25) / 3.14159)) + (3.14159 - float32(float64(0.1)))) - 1.5)
                    _ = v48
                    _ = v49
                }
            }
            _ = v44
        }
    default:
        v50 := []int{0}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
