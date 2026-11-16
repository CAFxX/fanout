package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        switch 183 {
        case -294:
            v8 := GenS1{F0: 42, F1: 189, F2: ((-1 > 42) && true)}
            _ = v8
        default:
            var v9 uint64 = 139
            _ = v9
        }
        var v10 string = "\n\t"
        switch 0 {
        case 2147483647:
            var v11 uint8 = 0
            _ = v11
        case 42:
            fmt.Println(v10)
        case 4294967295:
            v10 = ("foo bar" + (v10 + "foo bar"))
        default:
            v12 := GenS1{F0: 294, F1: 183, F2: false}
            _ = v12
        }
        _ = fi7
        _ = v10
    }
    if true {
        var v13 int16 = 1000
        var v14 uint8 = 127
        v15 := []int{}
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 uint8 = 255
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float32 = 3.14159
            var v19 int16 = int16(v16)
            _ = fi17
            _ = v18
            _ = v19
        }
        switch "a" {
        case "foo bar":
            switch ("\n\t" + "foo bar") {
            case "":
                fmt.Println(v16)
            case "hello":
                v20 := []int{-868}
                _ = v20
            case "foo bar":
                if true {
                    if (1 > 760) {
                        fmt.Println(v16)
                        var v21 string = "\n\t"
                        if false {
                            fmt.Println(v21)
                        } else {
                            v16 -= (v16 << ((65535) & 63))
                            var v22 int16 = -459
                            _ = v22
                        }
                        _ = v21
                    } else {
                        if true {
                            v23 := GenS1{F0: 215, F1: 18446744073709551615, F2: true}
                            _ = v23
                        } else {
                            var v24 uint8 = v16
                            var v25 float32 = -2.25
                            var v26 uint64 = 2
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: 0, F1: 0, F2: (!false)}
                        v16 -= v16
                        _ = v27
                    }
                    switch 1 {
                    case -847:
                        v16 = 103
                    default:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v29 := []int{-15, fi28, (fi28 / ((fi28) | 1)), (fi28 / (((fi28 | (fi28 / ((fi28) | 1)))) | 1))}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
            }
        }
        _ = v16
    }
    v30 := []int{1, 1, 4294967295}
    var v31 uint = 919
    _ = v30
    _ = v31
    return 379
}
func genFn1(p0 int32, p1 string) uint8 {
    p0 = (p0 % (((p0 + (p0 % ((p0) | 1)))) | 1))
    var v32 uint16 = 589
    _ = p0
    _ = p1
    _ = v32
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := []int{32767, 1}
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]uint32)
    switch 32768 {
    case 1:
        if true {
            var v33 float64 = float64(-2.25)
            _ = v33
        }
    case -308:
        var v34 uint8 = 1
        _ = v34
    }
    var v35 bool = (!false)
    v35 = v35
    v35 = v35
    fmt.Println(v35)
    var v36 bool = false
    var v37 int = 9223372036854775807
    if v36 {
        switch 127 {
        case -189:
            for fi38 := 0; fi38 < 3; fi38++ {
                _ = len(v1)
                v3 = (v3 - 2147483647)
                _ = fi38
            }
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            v40 := GenS2{F0: -1}
            _ = fi39
            _ = v40
        }
        var v41 uint64 = genFn0()
        _ = v41
    } else {
        if (!v36) {
            if true {
                var v42 int16 = 42
                v43 := GenS3{F0: (v37 >> ((2) & 63)), F1: 42}
                _ = len(v6)
                _ = v42
                _ = v43
            }
        }
    }
    var v44 int32 = v3
    var v45 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
