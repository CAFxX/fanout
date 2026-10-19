package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    delete(v2, "hello")
    v5 := []int{4294967295, -1, -518}
    v6 := GenS1{F0: (255 != 2)}
    if (369 >= 32767) {
        var v7 uint16 = 42
        var v8 int = 127
        v9 := GenS1{F0: false}
        _ = v7
        _ = v8
        _ = v9
    } else {
        if len(v5) > 0 {
            v5[0] = 0
        }
        switch ("hello" + ("a" + "a")) {
        case "foo bar":
            var v10 int8 = -1
            _ = v10
        case "":
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    fmt.Println(fi12, fi12)
                    _ = fi12
                }
                _ = fi11
            }
        case "\n\t":
            var v13 uint64 = 4294967295
            _ = v13
        }
    }
    var v14 uint16 = 0
    v15 := GenS1{F0: (!true)}
    var v16 uint64 = 2
    var v17 int32 = 0
    switch 4294967295 {
    case 1:
        if (0 == 0) {
            if len(v5) > 0 {
                v5[0] = 1
            }
        } else {
            v16 -= v16
            var v18 int32 = (^(v17))
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            _ = v18
        }
    case 127:
        v19 := GenS1{F0: ((^(v16)) == 1)}
        _ = v19
    default:
        switch -2147483648 {
        case -569:
            v20 := GenS1{F0: true}
            _ = v20
        case 42:
            for fi21 := 0; fi21 < 9; fi21++ {
                v17 += (-(v17))
                v17 -= (v17 - (v17 >> ((2) & 63)))
                var v22 int16 = 42
                _ = fi21
                _ = v22
            }
        case 144:
            for fi23 := 0; fi23 < 17; fi23++ {
                _ = len(v5)
                switch (fi23 & fi23) {
                case 2147483647:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    if true {
                        switch ("foo bar" + "\n\t") {
                        case "":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        default:
                            v14 += v14
                        }
                        var v24 float64 = (float64(float32(((((-2.25 - -2.25) - 1.5) / (-2.25 * 0.1)) / 1e10))) / 1e10)
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 uint32 = 42
                            v26 -= v26
                            _ = fi25
                            _ = v26
                        }
                    }
                }
                if len(v5) > 0 {
                    v5[0] = fi23
                }
                _ = fi23
            }
        }
    }
    v14 = v14
    switch 127 {
    case 4294967295:
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 string = ((("" + "") + ((("a" + "foo bar") + "hello") + "\n\t")) + ("a" + ((("hello" + "a") + "hello") + (("\n\t" + "foo bar") + "foo bar"))))
            delete(v1, -1)
            _ = fi27
            _ = v28
        }
    case 0:
        for fi29 := 0; fi29 < 11; fi29++ {
            for fi30 := 0; fi30 < 15; fi30++ {
                v14 = (v14 >> ((42) & 63))
                if true {
                    fmt.Println(v14)
                }
                var v31 uint64 = 42
                var v32 uint32 = 4294967295
                fmt.Println(fi30, v17)
                v33 := GenS2{F0: ((fi30 * fi30) < fi30)}
                if false {
                    if false {
                        var v34 int32 = v17
                        v5 = append(v5, fi29)
                        _ = v34
                    }
                    var v35 uint16 = 2
                    _ = v35
                }
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            var v36 int16 = 1
            _ = fi29
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        v17 = (v17 + -2147483648)
        v38 := GenS3{F0: 783, F1: (^(v14))}
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 float32 = 1e10
            if true {
                v41 := GenS4{F0: fi39, F1: true}
                switch fi37 {
                case 127:
                    fmt.Println(v40, v40, v17)
                case -9223372036854775808:
                    fmt.Println(v17, fi37)
                default:
                    v16 += v16
                }
                _ = v41
            } else {
                v16 -= (^(v16))
            }
            _ = fi39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    var v42 uint16 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v42
    _ = v5
    _ = v6
}
