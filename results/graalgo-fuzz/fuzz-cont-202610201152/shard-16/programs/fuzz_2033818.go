package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int64) uint {
    p0 -= p0
    for fi5 := 0; fi5 < 13; fi5++ {
        p0 = p0
        for fi6 := 0; fi6 < 15; fi6++ {
            if false {
                v7 := []int{fi5, 889}
                var v8 float64 = -2.25
                p0 -= (p0 % ((p0) | 1))
                _ = v7
                _ = v8
            } else {
                p0 = -60
            }
            v9 := []int{}
            _ = fi6
            _ = v9
        }
        v10 := GenS1{F0: 42}
        _ = fi5
        _ = v10
    }
    _ = p0
    return 0
}
func genFn1(p0 uint16, p1 string) uint {
    p0 += (p0 | (p0 >> ((1) & 63)))
    v11 := []int{127, -9223372036854775808}
    v12 := GenS2{F0: 112}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return genFn0(1)
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    var v4 uint64 = 2
    switch -204 {
    case -225:
        v13 := GenS1{F0: 1}
        _ = v13
    case 980:
        v4 = (^(v4))
    default:
        v3 = append(v3, 709)
    }
    fmt.Println(v4)
    switch ("\n\t" + "") {
    case "hello":
        v3 = append(v3, 1)
    default:
        if false {
            var v14 uint32 = 2
            fmt.Println(v4)
            v14 += 2147483647
            _ = v14
        }
    }
    var v15 int32 = 42
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int = -437
        if ((v15 < v15) || (!(!((1e10 != (3.14159 * ((0.1 + 1.5) / 1.0))) || true)))) {
            v18 := GenS2{F0: 182}
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    v15 = v15
    v15 = (-(v15))
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            if (569 <= -9223372036854775808) {
                switch 4294967295 {
                case 1:
                    var v21 int32 = (v15 << ((0) & 63))
                    _ = v21
                default:
                    _ = len(v3)
                }
                var v22 uint32 = 1
                v23 := GenS1{F0: (v4 & (v4 * (v4 ^ 2)))}
                _ = v22
                _ = v23
            }
            var v24 float32 = 1.5
            var v25 uint16 = 553
            _ = fi20
            _ = v24
            _ = v25
        }
        if (!false) {
            v26 := GenS1{F0: 956}
            fmt.Println(v4, fi19)
            _ = len(v3)
            _ = v26
        } else {
            v27 := GenS2{F0: -1}
            _ = v27
        }
        _ = fi19
    }
    v28 := GenS3{F0: -1}
    v29 := GenS2{F0: -45}
    v30 := GenS4{F0: "hello"}
    fmt.Println(v4)
    fmt.Println(v15, v15)
    for fi31 := 0; fi31 < 14; fi31++ {
        switch (-(fi31)) {
        case 284:
            fmt.Println(v4, fi31, fi31)
        case 42:
            v4 = (v4 * (^(v4)))
        default:
            v32 := GenS1{F0: (^(v4))}
            _ = v32
        }
        v15 += v15
        if false {
            v15 += v15
            var v33 uint8 = 147
            _ = v33
        }
        _ = fi31
    }
    var v34 int32 = 1
    v15 = (v34 >> ((1) & 63))
    for fi35 := 0; fi35 < 1; fi35++ {
        switch "" {
        case "foo bar":
            v34 = v34
        case "\n\t":
            switch 729 {
            case -9223372036854775808:
                v36 := GenS5{F0: v4}
                _ = v36
            case 629:
                v15 = 1
            }
        case "hello":
            switch fi35 {
            case 4:
                var v37 uint = 255
                _ = v37
            }
        default:
            var v38 int = fi35
            _ = v38
        }
        v39 := v2[944]
        _ = fi35
        _ = v39
    }
    v40 := GenS2{F0: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v40
}
