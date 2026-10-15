package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint8) uint64 {
    if true {
        for fi5 := 0; fi5 < 20; fi5++ {
            fmt.Println(fi5)
            _ = fi5
        }
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 float32 = ((1.5 * 0.0) - -2.25)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42}
    var v9 uint = 255
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{891}
    v2 := []string{"hello"}
    v3 := make(map[int]int)
    v4 := []int{-996}
    if (true || (1 >= 488)) {
        switch 646 {
        case 0:
            _ = len(v2)
        case 1:
            var v10 int16 = -32768
            _ = v10
        }
        var v11 uint32 = 0
        _ = v11
    }
    if false {
        var v12 int64 = 1
        fmt.Println(v12)
        for fi13 := 0; fi13 < 9; fi13++ {
            switch fi13 {
            case 0:
                if true {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v12 = -645
                        _ = fi14
                    }
                    if true {
                        switch (fi13 | -531) {
                        case 755:
                            v15 := GenS1{F0: genFn0(42, 0)}
                            _ = v15
                        case -297:
                            v16 := GenS1{F0: 465}
                            _ = v16
                        case 32768:
                            if false {
                                if ("foo bar" == "foo bar") {
                                    var v17 uint16 = 0
                                    fmt.Println(fi13, v12)
                                    _ = v17
                                } else {
                                    var v18 int16 = -511
                                    for fi19 := 0; fi19 < 8; fi19++ {
                                        v20 := GenS2{F0: (v12 % (((-(v12))) | 1)), F1: 470}
                                        _ = fi19
                                        _ = v20
                                    }
                                    v12 = 42
                                    _ = v18
                                }
                                var v21 uint = 4294967295
                                _ = v21
                            } else {
                                v22 := GenS1{F0: 18446744073709551615}
                                var v23 int16 = -496
                                _ = v22
                                _ = v23
                            }
                        default:
                            v24 := GenS1{F0: 895}
                            _ = v24
                        }
                    }
                }
            }
            _ = len(v1)
            v25 := GenS3{F0: 485, F1: 98, F2: 42}
            _ = fi13
            _ = v25
        }
        _ = v12
    }
    switch "a" {
    case "":
        _ = len(v3)
    case "a":
        var v26 float64 = 1.5
        _ = v26
    case "foo bar":
        if true {
            var v27 uint = 0
            _ = v27
        } else {
            switch -1 {
            case 32768:
                v28 := GenS4{F0: 0, F1: 56}
                _ = v28
            }
        }
    }
    if false {
        var v29 uint64 = 447
        v3[-595] = 128
        v29 += v29
        _ = v29
    } else {
        var v30 int32 = -1
        _ = v30
    }
    var v31 float32 = 0.0
    v32 := GenS3{F0: 1, F1: 1, F2: 73}
    var v33 int = 32767
    if ((v31 - float32(3.14159)) > 1e10) {
        var v34 string = "foo bar"
        _ = v34
    }
    switch (("hello" + (((((("foo bar" + "hello") + "a") + "foo bar") + "foo bar") + "foo bar") + "foo bar")) + "foo bar") {
    case "hello":
        var v35 uint64 = 756
        _ = v35
    case "a":
        fmt.Println(v33)
    case "foo bar":
        var v36 uint8 = 128
        _ = v36
    }
    v33 += v33
    v33 += (v33 << ((670) & 63))
    v37 := GenS5{F0: 0}
    var v38 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
}
