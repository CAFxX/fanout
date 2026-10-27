package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0() int64 {
    var v7 uint64 = 1
    var v8 uint64 = (v7 >> ((255) & 63))
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint {
    var v9 uint64 = 4294967295
    if (false && (42 > 42)) {
        switch "" {
        case "":
            v10 := GenS1{F0: 1, F1: "", F2: 243}
            _ = v10
        case "foo bar":
            v11 := []int{32767, -257, 392, -274}
            _ = v11
        }
        fmt.Println(v9)
        v12 := GenS2{F0: ("\n\t" + "\n\t"), F1: 1, F2: 255}
        _ = v12
    } else {
        v13 := GenS1{F0: 42, F1: "\n\t", F2: 65535}
        v14 := []int{-1, 4294967295}
        var v15 int = 915
        _ = v13
        _ = v14
        _ = v15
    }
    v9 += 42
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 float64 = (((1e10 * float64(1.5)) * (-2.25 - 3.14159)) + float64((-2.25 / float32(-2.25))))
        v9 = 1
        if ((float64((1e10 + 0.1)) * ((((v17 * v17) - (0.1 * v17)) - (float64(0.1) - (0.0 - (v17 + ((float64(-2.25) + 1e10) / v17))))) + (v17 / v17))) <= v17) {
            v18 := []int{(fi16 + fi16)}
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    var v19 int8 = 100
    _ = v19
    _ = v9
    return 2
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]string)
    v3 := []uint32{4294967295}
    var v4 uint32 = 694
    var v5 float64 = (((-2.25 * 1.5) - (0.0 * (0.1 / 1.0))) * 3.14159)
    var v6 int = 900
    v20 := GenS2{F0: "hello", F1: -903, F2: 0}
    var v21 float64 = v5
    v5 = (float64((0.1 / 1e10)) * (3.14159 * float64((3.14159 + 0.1))))
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int32 = -2147483648
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS2{F0: ((("a" + "\n\t") + "\n\t") + "\n\t"), F1: 0, F2: 255}
            _ = fi24
            _ = v25
        }
        fmt.Println(v1, v5)
        _ = fi22
        _ = v23
    }
    fmt.Println(v1, v1)
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int = fi27
            v21 -= (float64(float32((-2.25 + (0.1 + float64(-2.25))))) + (((v5 + v5) - v21) + v5))
            _ = fi27
            _ = v28
        }
        var v29 uint16 = 32767
        var v30 float64 = v5
        _ = fi26
        _ = v29
        _ = v30
    }
    switch ("" + ("\n\t" + ("foo bar" + "\n\t"))) {
    case "":
        fmt.Println(v6, v4, v4)
    case "hello":
        v4 += v4
    default:
        v31 := GenS3{F0: 65535}
        _ = v31
    }
    v32 := GenS2{F0: "a", F1: 889, F2: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v33 int = v6
    v33 = 1
    fmt.Println(v6, v33, v21)
    v21 = ((((v21 - (v21 - 1e10)) * 1e10) / 1.5) / v21)
    if false {
        var v34 string = "foo bar"
        var v35 float32 = 3.14159
        _ = v34
        _ = v35
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
