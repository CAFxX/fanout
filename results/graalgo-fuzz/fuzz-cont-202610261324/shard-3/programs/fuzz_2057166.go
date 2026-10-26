package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0() int32 {
    v5 := GenS1{F0: 2}
    v6 := GenS2{F0: ((0.0 > (3.14159 + 1e10)) && ((false && true) || false))}
    var v7 uint16 = 61
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 float32 = (1e10 - (float32(((0.1 + 1.5) * (3.14159 * (float64((0.1 + float32(1e10))) - 1.5)))) / float32((1.5 / (float64(0.0) - 1.5)))))
    v2 := []uint32{}
    var v3 bool = ("a" == (("hello" + "a") + "hello"))
    v4 := make(map[int]string)
    v3 = ((0 != -1) || false)
    switch 155 {
    case 42:
        v2 = append(v2, 4294967295)
    default:
        fmt.Println(v1)
    }
    v8 := GenS3{F0: false}
    v9 := GenS3{F0: (v3 && (!false))}
    var v10 float32 = v1
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 bool = true
        if v12 {
            v10 = v1
        } else {
            switch "foo bar" {
            case "\n\t":
                if (1000 >= -32768) {
                    switch fi11 {
                    case 32768:
                        var v13 uint8 = 255
                        _ = v13
                    case 2147483647:
                        v3 = (!false)
                    case 191:
                        for fi14 := 0; fi14 < 2; fi14++ {
                            var v15 int64 = -1
                            var v16 int32 = int32(v15)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    default:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v18 := GenS4{F0: int16(fi11)}
                            var v19 float64 = ((1.5 - -2.25) - 1e10)
                            v2 = append(v2, 1)
                            v12 = v12
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                }
            case "a":
                fmt.Println(v10, v10)
            }
            _ = len(v4)
        }
        if (v3 || (1 >= 775)) {
            v1 = v10
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS5{F0: 128, F1: 4294967295}
                var v22 uint16 = 65535
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS4{F0: -380}
            _ = v23
        } else {
            var v24 uint64 = 270
            _ = v24
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v3, v3)
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 int = 32767
        v27 := GenS2{F0: v3}
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v3, v10)
    var v28 uint8 = 2
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS5{F0: fi29, F1: 42}
        var v31 uint32 = 65535
        switch -1 {
        case 254:
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 int64 = 42
                _ = fi32
                _ = v33
            }
        default:
            if (42 >= 1) {
                var v34 int64 = 565
                var v35 int32 = genFn0()
                var v36 int32 = genFn0()
                _ = v34
                _ = v35
                _ = v36
            }
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    var v37 uint64 = 2
    for fi38 := 0; fi38 < 6; fi38++ {
        switch ("foo bar" + "") {
        case "a":
            if false {
                for fi39 := 0; fi39 < 9; fi39++ {
                    var v40 uint = 469
                    _ = fi39
                    _ = v40
                }
                v3 = (!(42 != -1))
            }
        case "hello":
            for fi41 := 0; fi41 < 11; fi41++ {
                fmt.Println(v10)
                _ = fi41
            }
        default:
            fmt.Println(v37)
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
