package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0() bool {
    var v5 int32 = 2147483647
    var v6 float64 = ((0.0 - (-2.25 - 0.0)) * (1.5 / 1e10))
    fmt.Println(v6)
    _ = v5
    _ = v6
    return (!(v5 == -1))
}

func main() {
    var v1 bool = false
    var v2 bool = v1
    v3 := make(map[int]uint32)
    var v4 int32 = 42
    switch "hello" {
    case "a":
        var v7 uint32 = 135
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: 656, F1: 42, F2: 100}
        _ = v8
    case "hello":
        v9 := []int{127, 520}
        _ = v9
    }
    if false {
        v10 := GenS1{F0: 499, F1: 766, F2: 42}
        if ((0.1 - 0.0) == 1.5) {
            v11 := GenS2{F0: 984}
            _ = v11
        } else {
            var v12 int64 = -364
            var v13 bool = ((!((74 != 1) || v1)) || (!false))
            _ = v12
            _ = v13
        }
        _ = v10
    }
    var v14 int64 = 42
    v2 = v1
    v1 = (!true)
    v2 = (!true)
    v15 := GenS2{F0: 65535}
    if v1 {
        if true {
            for fi16 := 0; fi16 < 14; fi16++ {
                fmt.Println(v14)
                _ = fi16
            }
            var v17 uint = 4294967295
            _ = v17
        }
    } else {
        v18 := GenS3{F0: 697, F1: v14}
        _ = v18
    }
    var v19 string = ("" + "\n\t")
    if false {
        if (0 > 504) {
            switch v19 {
            case "\n\t":
                var v20 int16 = 1
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                switch "hello" {
                case "":
                    var v22 bool = v1
                    _ = v22
                case "hello":
                    var v23 float32 = (((((-2.25 * 0.0) - (1e10 - 0.1)) + (((1e10 - (0.1 + (0.0 + (1.5 + float32((1.5 - float64(1.5))))))) / (0.1 - -2.25)) + 0.1)) / 1e10) + (-2.25 * 1.5))
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v19 = v19
                        if true {
                            v4 = -135
                        } else {
                            if (-32768 < -1) {
                                var v25 int = 127
                                v26 := GenS2{F0: 65535}
                                var v27 bool = v1
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = len(v3)
                            fmt.Println(v4)
                        }
                        _ = fi24
                    }
                }
                v1 = (v1 || (!v1))
                if true {
                    var v28 uint16 = 1
                    switch (v19 + v19) {
                    case "":
                        switch 250 {
                        case 32767:
                            var v29 uint16 = 836
                            _ = v29
                        default:
                            if v1 {
                                v19 = ("hello" + v19)
                                v30 := GenS3{F0: 0, F1: v14}
                                var v31 uint64 = 4294967295
                                _ = v30
                                _ = v31
                            } else {
                                v19 = (v19 + "")
                            }
                        }
                    case "foo bar":
                        v32 := GenS4{F0: 255, F1: v28}
                        _ = v32
                    default:
                        var v33 float32 = float32((0.1 * -2.25))
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 15; fi34++ {
                        _ = len(v3)
                        v4 = -2147483648
                        v1 = v2
                        _ = fi34
                    }
                    _ = v28
                } else {
                    v4 = (^(v4))
                    v4 = (v4 << ((916) & 63))
                    var v35 int16 = 42
                    _ = v35
                }
                _ = fi21
            }
            switch v19 {
            case "\n\t":
                v14 -= v14
            case "":
                switch 128 {
                case 32768:
                    v19 = "hello"
                }
            case "hello":
                var v36 uint8 = 255
                _ = v36
            default:
                var v37 uint8 = 0
                _ = v37
            }
        } else {
            var v38 int64 = v14
            var v39 int16 = 1000
            _ = v38
            _ = v39
        }
        v4 -= 378
    }
    fmt.Println(v4, v19, v2)
    v40 := []int{42, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
