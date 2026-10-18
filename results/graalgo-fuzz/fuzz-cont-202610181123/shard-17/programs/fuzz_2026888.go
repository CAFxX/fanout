package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0() int16 {
    var v5 string = (("foo bar" + "a") + "foo bar")
    var v6 uint16 = 32767
    v7 := GenS1{F0: 1, F1: 32767}
    var v8 int64 = -9223372036854775808
    switch -71 {
    case 236:
        v9 := []int{1, 2147483647}
        _ = v9
    default:
        v10 := GenS1{F0: 100, F1: 127}
        _ = v10
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -778
    var v2 uint16 = 1
    var v3 float64 = float64(((1.5 - (float32(-2.25) + float32((0.1 / -2.25)))) * 0.1))
    var v4 float32 = 1.5
    var v11 uint64 = 0
    var v12 uint64 = v11
    v2 += v2
    v13 := []int{-1, 0}
    var v14 int8 = -84
    if true {
        var v15 uint = 65535
        var v16 int64 = -1
        for fi17 := 0; fi17 < 18; fi17++ {
            if ((^(v11)) != v12) {
                fmt.Println(v3, v11, v11)
            }
            switch fi17 {
            case 32768:
                var v18 uint32 = 42
                _ = v18
            }
            if false {
                var v19 uint8 = 255
                if true {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if (genFn0() == (v1 * (-(v1)))) {
                            if false {
                                if (false || true) {
                                    var v21 int32 = 42
                                    switch (-(fi17)) {
                                    case 447:
                                        _ = len(v13)
                                    case 500:
                                        switch 128 {
                                        case -258:
                                            fmt.Println(fi17)
                                            v13 = append(v13, -1)
                                            var v22 uint8 = 255
                                            _ = v22
                                        }
                                    case 2147483647:
                                        var v23 uint64 = (v11 | v12)
                                        _ = v23
                                    }
                                    _ = v21
                                } else {
                                    if true {
                                        fmt.Println(v19)
                                        v13 = append(v13, -479)
                                    } else {
                                        if (false || (!true)) {
                                            _ = len(v13)
                                        }
                                    }
                                    v24 := GenS2{F0: v2, F1: 2, F2: 42}
                                    v14 = v14
                                    _ = v24
                                }
                                v1 += v1
                            } else {
                                var v25 uint64 = 1
                                v26 := GenS3{F0: uint64(v2)}
                                fmt.Println(v11, fi20)
                                var v27 uint64 = (v11 >> ((v15) & 63))
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                        _ = fi20
                    }
                    var v28 uint = (v15 & uint(v2))
                    switch int(v1) {
                    case -121:
                        if true {
                            var v29 uint32 = 360
                            _ = v29
                        } else {
                            v30 := GenS2{F0: (v2 | (v2 ^ 255)), F1: v19, F2: uint64(v19)}
                            _ = v30
                        }
                    case 42:
                        var v31 uint16 = (v2 ^ (v2 + v2))
                        _ = v31
                    }
                    _ = v28
                } else {
                    v32 := GenS4{F0: (v2 - 42), F1: (v14 ^ (v14 >> ((v15) & 63)))}
                    v11 = (v11 ^ (^(v11)))
                    fmt.Println(v11, fi17)
                    _ = v32
                }
                _ = v19
            } else {
                v33 := GenS3{F0: 4294967295}
                var v34 bool = false
                _ = v33
                _ = v34
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    var v35 uint8 = 0
    var v36 int16 = -32768
    var v37 int64 = 9223372036854775807
    v38 := GenS1{F0: int8(v2), F1: -772}
    _ = len(v13)
    var v39 uint8 = v35
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v40 bool = true
    var v41 bool = v40
    v1 -= int16(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
