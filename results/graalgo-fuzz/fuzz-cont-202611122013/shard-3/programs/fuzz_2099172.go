package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool; F1 int64; F2 uint16 }
func genFn0(p0 uint32, p1 uint) uint8 {
    switch "" {
    case "foo bar":
        var v6 uint16 = 42
        _ = v6
    case "":
        v7 := []int{}
        _ = v7
    }
    v8 := GenS1{F0: 1}
    if false {
        var v9 uint8 = 89
        var v10 bool = true
        p0 = (p0 << ((240) & 63))
        _ = v10
        _ = v9
    } else {
        if ((p0 <= 1) || true) {
            for fi11 := 0; fi11 < 18; fi11++ {
                p0 += p0
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 int8 = 100
                    v14 := GenS2{F0: (-32 >= (v13 << ((p1) & 63)))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    if (p0 == p0) {
                        v16 := []int{0}
                        _ = v16
                    } else {
                        var v17 uint32 = (^(p0))
                        _ = v17
                    }
                    fmt.Println(fi15, fi15)
                    _ = fi15
                }
                _ = fi11
            }
            switch "a" {
            case "hello":
                p0 = 1
            case "foo bar":
                v18 := GenS1{F0: -32768}
                _ = v18
            case "\n\t":
                switch ("a" + "a") {
                case "\n\t":
                    var v19 int8 = -128
                    _ = v19
                case "a":
                    var v20 int64 = 9223372036854775807
                    _ = v20
                default:
                    v21 := []int{}
                    _ = v21
                }
            default:
                switch -2147483648 {
                case 919:
                    v22 := GenS2{F0: true}
                    _ = v22
                case 42:
                    var v23 int16 = 1
                    _ = v23
                case 439:
                    var v24 int = 543
                    _ = v24
                default:
                    var v25 uint64 = 312
                    _ = v25
                }
            }
            switch 9223372036854775807 {
            case -2147483648:
                p0 -= (p0 >> ((51) & 63))
            case 128:
                p0 -= (p0 & (p0 & p0))
            }
        }
        var v26 int = 259
        switch "" {
        case "hello":
            v27 := []int{v26, (v26 + v26)}
            _ = v27
        case "":
            v28 := GenS3{F0: p0, F1: 42}
            _ = v28
        }
        _ = v26
    }
    var v29 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = v29
    _ = v8
    return (v29 / ((176) | 1))
}
func genFn1() bool {
    var v30 uint32 = 601
    var v31 int32 = 0
    _ = v30
    _ = v31
    return true
}

func main() {
    var v1 int32 = -821
    var v2 uint8 = 127
    var v3 uint64 = uint64(v2)
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    if (v1 > v1) {
        var v32 uint8 = v2
        var v33 uint8 = genFn0(697, 65535)
        _ = v32
        _ = v33
    } else {
        var v34 int8 = 100
        _ = v34
    }
    if true {
        var v35 int8 = 42
        if false {
            v1 = (-(v1))
            v36 := GenS1{F0: -1}
            _ = v36
        }
        v37 := GenS3{F0: 306, F1: 32767}
        _ = v35
        _ = v37
    }
    var v38 uint32 = 0
    var v39 bool = ((!false) || false)
    v40 := v4[2]
    var v41 float32 = 0.1
    var v42 int16 = 1000
    v41 = (v41 / v41)
    var v43 int16 = v42
    v44 := GenS4{F0: v39, F1: -9223372036854775808, F2: 573}
    v45 := v4[v38]
    switch 1 {
    case -485:
        v46 := v5[127]
        _ = v46
    case -931:
        if true {
            var v47 float32 = (v41 * (((0.1 / float32(3.14159)) - (v41 / float32(float64((v41 / 0.1))))) - (v41 - (0.1 * v41))))
            v3 += 598
            v45 = v45
            _ = v47
        }
    default:
        switch v40 {
        case "hello":
            v39 = (v39 || (!((genFn0(42, 72) >= genFn0((v38 & (v38 / ((v38) | 1))), 255)) || v39)))
        case "foo bar":
            if (!(0 == uint16(v2))) {
                v43 = 1
                v45 = v45
            }
        default:
            v2 = v2
        }
    }
    v41 -= (float32(1e10) * v41)
    var v48 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v5
}
