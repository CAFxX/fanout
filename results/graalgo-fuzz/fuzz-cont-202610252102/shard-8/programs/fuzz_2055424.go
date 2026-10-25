package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    v3 := make(map[string]string)
    var v4 float32 = 1.5
    v5 := make(map[int]bool)
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            v4 -= (float32((0.1 / ((1.5 / (1.5 + 3.14159)) * float64(-2.25)))) / v4)
            _ = fi6
        }
    }
    fmt.Println(v2, v4)
    var v7 int8 = 127
    var v8 int32 = (v2 >> ((65535) & 63))
    v8 -= v2
    fmt.Println(v8, v2, v8)
    v2 = (v2 | v8)
    switch 4294967295 {
    case 4294967295:
        fmt.Println(v4)
    case -9223372036854775808:
        var v9 uint8 = 160
        _ = v9
    case 0:
        if true {
            var v10 uint = 0
            v11 := GenS1{F0: 47, F1: (-(v7))}
            switch 128 {
            case 741:
                for fi12 := 0; fi12 < 2; fi12++ {
                    v2 += v2
                    var v13 string = "foo bar"
                    _ = fi12
                    _ = v13
                }
            }
            _ = v10
            _ = v11
        } else {
            v14 := GenS1{F0: 147, F1: v7}
            v4 = v4
            for fi15 := 0; fi15 < 20; fi15++ {
                if true {
                    if false {
                        if true {
                            if false {
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    v7 = 1
                                default:
                                    fmt.Println(v7, v8, v7)
                                }
                            } else {
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    if ((true || false) || false) {
                                        var v17 bool = false
                                        var v18 uint32 = 2
                                        var v19 uint16 = 32767
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := []int{-9223372036854775808, -917, fi16, (fi15 & 127)}
                                        var v21 uint = 1
                                        var v22 uint32 = 0
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    fmt.Println(v4, fi15)
                                    if (-179 > fi15) {
                                        v8 -= v2
                                        if (0 >= 0) {
                                            fmt.Println(v4, v4)
                                            switch "" {
                                            case "foo bar":
                                                v23 := GenS1{F0: 2, F1: -78}
                                                _ = v23
                                            case "a":
                                                if (true || true) {
                                                    var v24 float32 = (1.5 - float32((1e10 - (-2.25 / -2.25))))
                                                    v25 := GenS2{F0: 246}
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v4 = v4
                                            default:
                                                var v26 int16 = 0
                                                _ = v26
                                            }
                                            fmt.Println(fi16, v4, v7)
                                        } else {
                                            if (true && (!true)) {
                                                v27 := GenS3{F0: false}
                                                var v28 uint16 = 42
                                                fmt.Println(fi16)
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 int64 = -1
                                            var v30 int16 = 1000
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 uint = 255
                                        _ = v31
                                    }
                                    v32 := GenS2{F0: 1}
                                    _ = fi16
                                    _ = v32
                                }
                                if true {
                                    v33 := GenS4{F0: -487, F1: 1}
                                    var v34 uint = 4294967295
                                    v35 := []int{(fi15 >> ((42) & 63)), fi15, fi15, fi15}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            fmt.Println(v4, v4)
                        } else {
                            switch "\n\t" {
                            case "":
                                if true {
                                    var v36 int32 = v8
                                    _ = v36
                                }
                            }
                            v4 += 3.14159
                            fmt.Println(fi15)
                        }
                        v37 := []int{-9223372036854775808, fi15}
                        _ = v37
                    }
                    var v38 int = -971
                    switch v38 {
                    case 1:
                        fmt.Println(v38)
                    default:
                        v7 += v7
                    }
                    _ = v38
                }
                _ = fi15
            }
            _ = v14
        }
    default:
        fmt.Println(v7, v8)
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := []int{}
        v8 = 527
        delete(v5, (fi39 * (-(fi39))))
        _ = fi39
        _ = v40
    }
    var v41 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v7
    _ = v8
}
