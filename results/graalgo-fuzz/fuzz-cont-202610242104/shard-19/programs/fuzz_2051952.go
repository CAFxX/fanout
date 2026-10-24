package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int8 }
func genFn0() string {
    v6 := GenS1{F0: -1, F1: 550}
    var v7 uint = 816
    v8 := []int{}
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 uint64, p1 int, p2 int8) uint64 {
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint32 = 42
            v12 := GenS2{F0: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(p2, p0, p0)
        switch fi9 {
        case 9223372036854775807:
            v13 := []int{(fi9 % ((p1) | 1)), -248, int(p0), 128}
            _ = v13
        case 128:
            v14 := []int{}
            _ = v14
        case 0:
            v15 := []int{p1}
            _ = v15
        default:
            v16 := []int{42, (p1 / ((9223372036854775807) | 1)), (p1 - (p1 / ((fi9) | 1))), (fi9 << ((18446744073709551615) & 63))}
            _ = v16
        }
        _ = fi9
    }
    p0 = (p0 << ((1) & 63))
    _ = p0
    _ = p1
    _ = p2
    return (p0 % (((^(p0))) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 65535
    v4 := make(map[uint32]int)
    var v5 int = -1
    v3 -= v3
    var v17 float64 = ((1e10 / 1.0) + float64(0.1))
    v5 = v5
    fmt.Println(v5)
    if false {
        v3 = 0
        v18 := []int{v5}
        if ((v5 / (((-(v5))) | 1)) != (v5 + (v5 | v5))) {
            if false {
                if (!false) {
                    v19 := GenS2{F0: 4294967295}
                    var v20 int = 370
                    _ = v19
                    _ = v20
                }
                var v21 uint8 = 127
                _ = v21
            }
        }
        _ = v18
    }
    var v22 bool = (-243 <= 331)
    var v23 uint16 = 2
    v22 = (v22 && (0 == -835))
    v24 := v2[2147483647]
    var v25 int32 = -805
    if true {
        v23 = (^(v23))
        fmt.Println(v5)
    } else {
        for fi26 := 0; fi26 < 4; fi26++ {
            v22 = (((float32(v17) + ((1.5 - (float32(v17) + (1e10 - -2.25))) / (1.5 * 1.5))) * 1e10) != 0.0)
            fmt.Println(v23)
            _ = fi26
        }
        v27 := []int{}
        _ = v27
    }
    v17 = float64((3.14159 + 0.1))
    var v28 float64 = v17
    v3 = 65535
    v29 := GenS1{F0: 2147483647, F1: 65535}
    switch ("hello" + ("hello" + (("foo bar" + "a") + (((((("hello" + "a") + "a") + ("hello" + "foo bar")) + "a") + "") + ("" + ("\n\t" + "hello")))))) {
    case "":
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 int = v5
            switch (v5 - v31) {
            case 772:
                var v32 string = "foo bar"
                _ = v32
            case 389:
                if (42 < 255) {
                    if v22 {
                        v33 := GenS3{F0: (1 < 18446744073709551615), F1: (v24 | (v24 << ((1) & 63)))}
                        v34 := GenS3{F0: true, F1: v24}
                        v35 := v2[(^(v24))]
                        _ = v33
                        _ = v34
                        _ = v35
                    } else {
                        fmt.Println(v25, v17)
                        fmt.Println(v3, v22)
                    }
                    var v36 uint32 = v24
                    v5 += (fi30 | 909)
                    _ = v36
                }
            default:
                var v37 string = ((("" + ("\n\t" + "\n\t")) + ("hello" + ("" + ("foo bar" + "")))) + "\n\t")
                _ = v37
            }
            v17 = v17
            _ = fi30
            _ = v31
        }
    case "\n\t":
        v38 := GenS4{F0: 127}
        _ = v38
    default:
        v28 += -2.25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
