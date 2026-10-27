package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
func genFn0() uint {
    if (1 < 549) {
        v4 := GenS1{F0: true}
        switch 4294967295 {
        case -418:
            if ((255 == 1) || ("a" == "hello")) {
                if false {
                    var v5 bool = (-240 == 2147483647)
                    for fi6 := 0; fi6 < 17; fi6++ {
                        v7 := []int{32768, (fi6 % ((9223372036854775807) | 1))}
                        _ = fi6
                        _ = v7
                    }
                    _ = v5
                }
                v8 := []int{-139, 9223372036854775807, 127}
                _ = v8
            }
        case 127:
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint16 = 1
                v10 += 65535
                _ = fi9
                _ = v10
            }
        default:
            var v11 float32 = float32(1e10)
            _ = v11
        }
        _ = v4
    } else {
        var v12 uint16 = 1
        _ = v12
    }
    v13 := []int{32768, 852}
    var v14 bool = ((!((3.14159 - 1.5) == 0.1)) && (true && false))
    _ = v13
    _ = v14
    return 253
}
func genFn1(p0 int64, p1 uint8, p2 bool) int32 {
    var v15 uint32 = 65535
    v16 := GenS2{F0: 0, F1: 0}
    v17 := []int{127, 32767, 42}
    var v18 float64 = (1e10 + 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return 74
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 1
    var v3 uint64 = 159
    v19 := GenS2{F0: genFn1(-9223372036854775808, 6, false), F1: -1}
    v3 = (v3 >> (((v2 - v2)) & 63))
    var v20 string = "foo bar"
    v20 = ("foo bar" + v20)
    fmt.Println(v20)
    var v21 float32 = float32(float64(1.5))
    var v22 uint64 = 823
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS3{F0: 9223372036854775807, F1: 127}
        v25 := GenS3{F0: -9223372036854775808, F1: 221}
        if (false || false) {
            v20 = v20
            switch v20 {
            case "a":
                var v26 uint8 = 127
                _ = v26
            case "foo bar":
                v2 += (v2 | (v2 >> ((v2) & 63)))
            case "hello":
                v22 = (v22 * v22)
            }
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS3{F0: 892, F1: 0}
                if false {
                    fmt.Println(v22)
                    if (42 < 1) {
                        v29 := []int{-9223372036854775808}
                        v20 = ("" + v20)
                        _ = v29
                    } else {
                        v30 := GenS4{F0: (("a" + v20) + ("" + "a"))}
                        v31 := GenS5{F0: 4294967295}
                        _ = v30
                        _ = v31
                    }
                }
                v21 = v21
                _ = fi27
                _ = v28
            }
        } else {
            var v32 uint32 = uint32(v2)
            _ = v32
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    if ((151 < 191) || (!false)) {
        v33 := v1[0]
        fmt.Println(v22)
        var v34 bool = ((v33 >> (((v2 + 42)) & 63)) == v33)
        _ = v33
        _ = v34
    }
    var v35 uint8 = 127
    if (v3 <= (v3 & (^(v3)))) {
        if false {
            v35 = (v35 & uint8(v3))
            var v36 uint32 = 1
            switch "\n\t" {
            case "foo bar":
                var v37 int16 = 1
                _ = v37
            case "a":
                var v38 int64 = 42
                _ = v38
            case "hello":
                var v39 int8 = 31
                _ = v39
            }
            _ = v36
        }
        v40 := GenS1{F0: false}
        v20 = "\n\t"
        _ = v40
    } else {
        for fi41 := 0; fi41 < 19; fi41++ {
            v20 = (((v20 + "a") + (("" + ((("\n\t" + "foo bar") + ("" + v20)) + "a")) + "foo bar")) + (v20 + v20))
            _ = fi41
        }
        var v42 int8 = 100
        var v43 uint64 = 884
        _ = v42
        _ = v43
    }
    v3 = 252
    var v44 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v44
}
