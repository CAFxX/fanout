package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
func genFn0(p0 int16) int32 {
    var v6 uint32 = 65535
    switch 192 {
    case 127:
        var v7 uint32 = 2147483648
        _ = v7
    case 32767:
        var v8 int8 = -43
        _ = v8
    case 2147483647:
        if true {
            v9 := GenS1{F0: (p0 / ((1000) | 1)), F1: -2147483648}
            v10 := GenS1{F0: p0, F1: 9223372036854775807}
            _ = v10
            _ = v9
        }
    }
    var v11 uint8 = uint8(v6)
    if (0.1 >= -2.25) {
        if (false && (!(42 > 65535))) {
            p0 -= (p0 >> ((0) & 63))
            if true {
                v12 := GenS1{F0: p0, F1: int(v6)}
                v13 := GenS2{F0: 309}
                _ = v12
                _ = v13
            }
        }
        var v14 uint16 = 782
        v15 := []int{4294967295, -1}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{433, 202}
    var v2 int16 = 465
    var v3 float64 = (0.0 - ((0.0 - float64(-2.25)) * 1.5))
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    if true {
        v2 += (v2 + (v2 + (v2 | v2)))
    } else {
        var v16 float64 = float64(0.1)
        _ = v16
    }
    v2 -= v2
    if (!true) {
        var v17 int64 = 440
        _ = v17
    }
    v2 = (v2 * (v2 << ((66) & 63)))
    if false {
        v2 = (v2 + v2)
        v2 += -32768
    }
    switch "" {
    case "":
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS1{F0: (v2 % ((0) | 1)), F1: fi19}
                fmt.Println(fi18)
                v3 -= (v3 - (v3 / v3))
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v3 -= v3
                v3 = (v3 / v3)
                if (fi21 > 2147483647) {
                    if ((1 == 2) || (v2 < v2)) {
                        v3 = v3
                        var v22 float64 = 3.14159
                        _ = v22
                    } else {
                        var v23 string = "foo bar"
                        var v24 uint16 = 255
                        _ = v23
                        _ = v24
                    }
                }
                _ = fi21
            }
            _ = fi18
        }
    case "foo bar":
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint16 = 65535
            switch "a" {
            case "foo bar":
                fmt.Println(v26)
            case "a":
                switch "foo bar" {
                case "hello":
                    if ("foo bar" < "a") {
                        v3 -= (v3 - -2.25)
                    } else {
                        if true {
                            v27 := GenS1{F0: v2, F1: fi25}
                            v4[(("\n\t" + "a") + "")] = (fi25 * fi25)
                            var v28 uint32 = 2147483647
                            _ = v27
                            _ = v28
                        } else {
                            v26 = v26
                            v29 := GenS1{F0: 224, F1: fi25}
                            v3 -= (v3 - 0.0)
                            _ = v29
                        }
                    }
                default:
                    var v30 int = fi25
                    _ = v30
                }
            }
            _ = fi25
            _ = v26
        }
    case "hello":
        for fi31 := 0; fi31 < 3; fi31++ {
            switch -742 {
            case 32767:
                v32 := GenS1{F0: v2, F1: fi31}
                _ = v32
            case -1:
                var v33 int8 = 42
                _ = v33
            }
            v3 -= ((1.5 * (1.5 / v3)) * v3)
            _ = fi31
        }
    }
    fmt.Println(v3, v3)
    v34 := GenS3{F0: -9223372036854775808}
    var v35 int = -703
    var v36 uint = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
