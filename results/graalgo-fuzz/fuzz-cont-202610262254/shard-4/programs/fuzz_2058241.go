package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 int) int32 {
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 17; fi5++ {
        v4 = 5
        p1 -= (fi5 & 42)
        v4 = 255
        _ = fi5
    }
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint16 = 0
            if true {
                p0 -= uint32(v7)
                var v8 uint16 = uint16(v4)
                v9 := GenS1{F0: 518, F1: (v4 + v4), F2: -747}
                _ = v8
                _ = v9
            }
            v7 = v7
            _ = fi6
            _ = v7
        }
    }
    v10 := GenS1{F0: -2147483648, F1: v4, F2: 32767}
    var v11 uint16 = 654
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v4
    return -37
}
func genFn1(p0 uint16, p1 bool, p2 uint) int32 {
    var v12 int16 = -606
    switch 2147483647 {
    case 4294967295:
        if p1 {
            for fi13 := 0; fi13 < 17; fi13++ {
                p0 += 1
                v12 += (v12 % ((v12) | 1))
                v14 := []int{(fi13 ^ (-(fi13))), fi13}
                _ = fi13
                _ = v14
            }
            v12 += v12
        } else {
            v15 := []int{}
            _ = v15
        }
    default:
        var v16 string = (("a" + (("\n\t" + "a") + "\n\t")) + "\n\t")
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2147483647
}

func main() {
    var v1 float64 = (((0.1 - (0.0 / ((1.5 * (float64((-2.25 + 1e10)) + (float64((((1e10 / 1e10) - 1e10) + (3.14159 * (3.14159 - float32((3.14159 / 3.14159)))))) + -2.25))) * (1.5 / 1.0)))) + (1.5 / float64(((float32(-2.25) + 1.5) - ((3.14159 + (-2.25 - 3.14159)) / 0.1))))) / (1.5 / 0.1))
    var v2 int8 = 100
    var v3 bool = true
    v17 := []int{-9223372036854775808}
    var v18 uint32 = 2
    if (!v3) {
        v3 = (-1 < -631)
    } else {
        v19 := GenS1{F0: 2147483647, F1: 127, F2: -20}
        _ = v19
    }
    if (1 >= 574) {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch "hello" {
            case "hello":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            }
            fmt.Println(v2, v1)
            var v21 bool = false
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: -303, F1: 2, F2: 659}
        _ = v22
    } else {
        if false {
            var v23 int32 = -2147483648
            _ = v23
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                fmt.Println(v3, v2, v2)
                var v25 int8 = 127
                var v26 uint32 = v18
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        var v27 uint16 = 0
        var v28 uint64 = uint64(v18)
        _ = v27
        _ = v28
    }
    v1 = v1
    v2 -= v2
    for fi29 := 0; fi29 < 1; fi29++ {
        v2 = v2
        var v30 string = ("a" + "hello")
        var v31 uint64 = 0
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 bool = (506 > 1)
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 float32 = 3.14159
        switch (-(fi33)) {
        case 713:
            v35 := GenS1{F0: 354, F1: 128, F2: -1}
            _ = v35
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 uint16 = 301
            switch fi36 {
            case 9223372036854775807:
                fmt.Println(v2, fi33)
            case 2147483647:
                var v38 uint = 451
                _ = v38
            default:
                var v39 uint8 = 180
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    v40 := GenS2{F0: -1}
    v1 -= ((float64(-2.25) / v1) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v40
}
