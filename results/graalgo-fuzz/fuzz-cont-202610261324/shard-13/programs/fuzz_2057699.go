package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int8, p1 bool) uint64 {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 = -51
        _ = fi4
    }
    var v5 int32 = -1
    _ = p0
    _ = p1
    _ = v5
    return 2
}
func genFn1() uint {
    var v6 uint32 = 42
    switch -1 {
    case 4294967295:
        v6 -= v6
    case 1:
        v7 := []int{}
        _ = v7
    default:
        v8 := []int{695, 43, 355, 4294967295}
        _ = v8
    }
    fmt.Println(v6)
    switch 32768 {
    case 9223372036854775807:
        v9 := []int{9223372036854775807, 9223372036854775807, 32767}
        _ = v9
    case -2147483648:
        v6 -= v6
    case -1:
        for fi10 := 0; fi10 < 20; fi10++ {
            v6 = (^(v6))
            var v11 float64 = -2.25
            for fi12 := 0; fi12 < 14; fi12++ {
                if true {
                    v13 := GenS1{F0: fi12}
                    _ = v13
                }
                var v14 uint16 = 2
                if (false || true) {
                    v11 -= (v11 / (v11 + v11))
                    if (-138 > -677) {
                        if false {
                            var v15 int = fi12
                            _ = v15
                        } else {
                            v16 := []int{}
                            var v17 int16 = -90
                            _ = v16
                            _ = v17
                        }
                    }
                }
                _ = fi12
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 int8 = 39
            _ = fi18
            _ = v19
        }
    }
    if ("foo bar" <= (("hello" + ("foo bar" + ("\n\t" + "hello"))) + "a")) {
        v20 := []int{-299}
        var v21 float32 = ((float32(1e10) * float32(3.14159)) / 1e10)
        _ = v20
        _ = v21
    }
    _ = v6
    return 0
}

func main() {
    v1 := []int{127}
    v2 := make(map[string]string)
    var v3 uint8 = 127
    v3 = (v3 & (v3 * (^(v3))))
    var v22 uint32 = 103
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v1) > 0 {
            v1[0] = (fi23 % ((fi23) | 1))
        }
        if true {
            v24 := GenS1{F0: fi23}
            _ = v24
        } else {
            var v25 float32 = (0.0 * (1.5 * 1e10))
            if false {
                var v26 int32 = -2147483648
                _ = v26
            }
            _ = v25
        }
        _ = len(v2)
        _ = fi23
    }
    v27 := GenS1{F0: -1}
    switch -2147483648 {
    case 1:
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := GenS1{F0: int(v22)}
            _ = fi28
            _ = v29
        }
    }
    if (-28 == 57) {
        var v30 float64 = 0.0
        if true {
            v3 -= 255
        }
        fmt.Println(v22)
        _ = v30
    } else {
        v3 = v3
    }
    fmt.Println(v3, v3)
    v31 := GenS1{F0: 0}
    if (!false) {
        var v32 string = "hello"
        v33 := v2[v32]
        _ = v32
        _ = v33
    }
    switch 4294967295 {
    case 402:
        switch ((("" + "a") + (("a" + ("hello" + "a")) + "foo bar")) + ("hello" + "a")) {
        case "\n\t":
            v22 -= (^(v22))
        case "":
            v34 := GenS2{F0: 4294967295}
            _ = v34
        case "a":
            var v35 float32 = -2.25
            _ = v35
        }
    case 4294967295:
        var v36 uint8 = 255
        _ = v36
    case 1:
        switch "hello" {
        case "a":
            if ((false || (1 < -1)) && true) {
                var v37 float32 = 3.14159
                _ = v37
            }
        case "foo bar":
            v3 -= 1
        default:
            for fi38 := 0; fi38 < 5; fi38++ {
                switch fi38 {
                case 2147483647:
                    var v39 uint16 = uint16(v3)
                    _ = v39
                case -2147483648:
                    v3 += v3
                default:
                    fmt.Println(v3)
                }
                v3 += v3
                _ = fi38
            }
        }
    default:
        v3 = (v3 & (v3 + v3))
    }
    var v40 int64 = 542
    var v41 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v31
    _ = v40
    _ = v41
}
