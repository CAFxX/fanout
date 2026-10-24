package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 string }
func genFn0() uint16 {
    switch 0 {
    case 2147483647:
        if true {
            if true {
                var v7 int64 = 42
                _ = v7
            }
        } else {
            var v8 int64 = -412
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int32 = 42
                for fi11 := 0; fi11 < 20; fi11++ {
                    v10 -= 87
                    var v12 float32 = 0.1
                    _ = fi11
                    _ = v12
                }
                v10 = (v10 / (((v10 % (((-(v10))) | 1))) | 1))
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    case 128:
        var v13 uint16 = 65535
        _ = v13
    case 42:
        for fi14 := 0; fi14 < 18; fi14++ {
            if ((-2147483648 > (fi14 - (fi14 % (((fi14 ^ 498)) | 1)))) || false) {
                var v15 float32 = -2.25
                var v16 int64 = 994
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int32 = 42
                    v19 := GenS1{F0: (fi17 * fi14), F1: 100, F2: 2147483648}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v15
                _ = v16
            }
            if false {
                fmt.Println(fi14)
            }
            _ = fi14
        }
    }
    var v20 float64 = 1e10
    if false {
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 uint8 = 59
                var v24 uint32 = 0
                var v25 float64 = 1.5
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v20 += v20
            _ = fi21
        }
        v26 := []int{-750, 128, 513, -246}
        for fi27 := 0; fi27 < 19; fi27++ {
            v20 -= (float64((0.0 * (1e10 + (1.5 + (-2.25 * ((float32(-2.25) * (0.1 / (float32(((float64((0.1 - (0.1 * 0.0))) / float64(1.5)) / float64(1e10))) * 1.5))) - (float32(1.5) * (1.5 - (float32((float64(float32(v20)) + float64(0.1))) + 1e10))))))))) + (1e10 + v20))
            v20 = ((v20 / v20) + 0.0)
            _ = fi27
        }
        _ = v26
    } else {
        var v28 int64 = -9223372036854775808
        v29 := GenS2{F0: v28, F1: (v28 ^ (v28 >> ((4294967295) & 63)))}
        v20 += v20
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 4294967295, F1: 100, F2: 1}
    _ = v20
    _ = v30
    return 65535
}

func main() {
    var v1 string = ""
    var v2 int16 = 32767
    v3 := []int{}
    var v4 string = (("foo bar" + ("\n\t" + (((v1 + v1) + v1) + v1))) + v1)
    v5 := make(map[int]bool)
    var v6 uint16 = 800
    var v31 float64 = 1.5
    delete(v5, -507)
    v32 := GenS2{F0: -1, F1: -135}
    fmt.Println(v6, v31, v6)
    v33 := GenS2{F0: -9223372036854775808, F1: -1}
    v34 := GenS3{F0: 42}
    v35 := GenS2{F0: 717, F1: 42}
    v36 := GenS4{F0: false, F1: v4}
    v1 = ("a" + ("a" + (("a" + v4) + (v4 + ""))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
