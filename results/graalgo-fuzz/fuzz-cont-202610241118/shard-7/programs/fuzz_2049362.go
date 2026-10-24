package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 uint64) uint64 {
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint16 = (^(p0))
            fmt.Println(v6, fi8, fi7)
            var v10 float32 = 1e10
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = fi7
        _ = v11
    }
    p0 += (p0 & 2)
    p1 = 128
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 | p2)
}

func main() {
    var v1 int16 = 0
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    var v4 uint16 = 255
    v5 := make(map[uint32]int)
    for fi12 := 0; fi12 < 20; fi12++ {
        if ("" == "a") {
            for fi13 := 0; fi13 < 9; fi13++ {
                v2 = v2
                fmt.Println(v1)
                var v14 uint32 = 74
                _ = fi13
                _ = v14
            }
            var v15 string = "a"
            _ = v15
        } else {
            switch fi12 {
            case 4294967295:
                var v16 string = "foo bar"
                _ = v16
            case 128:
                switch (fi12 * (fi12 ^ fi12)) {
                case 0:
                    v2 = v2
                case 127:
                    v17 := GenS1{F0: -2147483648, F1: (-(v1))}
                    _ = v17
                case 1:
                    var v18 bool = (-1 < 880)
                    _ = v18
                }
            default:
                var v19 int64 = 9223372036854775807
                _ = v19
            }
        }
        _ = fi12
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        switch fi20 {
        case 9223372036854775807:
            v5[815] = fi20
        default:
            v21 := []int{fi20, (fi20 + 32768), fi20, fi20}
            _ = v21
        }
        var v22 uint8 = 111
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 uint = 255
            var v25 float64 = -2.25
            var v26 int16 = (v1 >> ((v24) & 63))
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi20
        _ = v22
    }
    switch "\n\t" {
    case "":
        v2 -= 0
    case "\n\t":
        switch 32768 {
        case 32768:
            if (-860 < 9223372036854775807) {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v2 = v2
                    switch fi27 {
                    case 483:
                        var v28 int64 = 0
                        _ = v28
                    case 135:
                        _ = len(v5)
                    case 2147483647:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 int8 = 42
                            v1 = (v1 * 564)
                            var v31 uint16 = (v4 - 42)
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                    _ = fi27
                }
                v32 := []int{2147483647, 127, 1, 4294967295}
                v1 = v1
                _ = v32
            }
        }
    case "hello":
        v4 = 0
    default:
        _ = len(v3)
    }
    switch "\n\t" {
    case "a":
        v33 := GenS1{F0: 1, F1: v1}
        _ = v33
    case "\n\t":
        v4 = (v4 % ((558) | 1))
    case "hello":
        v1 += -793
    default:
        var v34 int64 = int64(v4)
        _ = v34
    }
    if true {
        v35 := []int{9223372036854775807, -1, -9223372036854775808, 2147483647}
        _ = v35
    } else {
        fmt.Println(v4, v2, v1)
    }
    fmt.Println(v1, v2)
    switch (("hello" + "\n\t") + "a") {
    case "a":
        v5[871] = 631
    }
    if true {
        switch -210 {
        case 32768:
            v4 = v4
        case 269:
            var v36 string = ("a" + "a")
            _ = v36
        }
    } else {
        var v37 int32 = 2147483647
        var v38 int8 = 100
        v39 := []int{9223372036854775807, -2147483648}
        _ = v37
        _ = v38
        _ = v39
    }
    v40 := v5[4294967295]
    v41 := GenS1{F0: 1, F1: v1}
    v1 -= (v1 >> ((4294967295) & 63))
    delete(v3, (v40 % ((v40) | 1)))
    for fi42 := 0; fi42 < 5; fi42++ {
        for fi43 := 0; fi43 < 15; fi43++ {
            if (false && (-1 <= 42)) {
                v4 -= v4
                v44 := GenS2{F0: 18446744073709551615}
                v45 := GenS2{F0: 765}
                v46 := GenS3{F0: 2147483648}
                v40 = v40
                v47 := []int{(fi43 >> ((4294967295) & 63)), (fi43 ^ (fi42 * fi43)), v40}
                _ = v44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi43
        }
        if (("hello" + "\n\t") != "foo bar") {
            v48 := v3[(fi42 % (((fi42 + v40)) | 1))]
            if ((fi42 | fi42) != fi42) {
                v49 := GenS1{F0: 0, F1: v1}
                v48 += (v48 - (v48 & (^(v48))))
                _ = v49
            }
            _ = v48
        }
        var v50 int = (fi42 / ((fi42) | 1))
        _ = fi42
        _ = v50
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        var v52 float32 = 0.0
        fmt.Println(v52, v40)
        _ = fi51
        _ = v52
    }
    v53 := GenS1{F0: 42, F1: int16(v4)}
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v53
}
