package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint64 }
func genFn0(p0 uint32, p1 uint16) uint64 {
    p1 = p1
    var v7 int32 = -749
    var v8 int64 = 469
    v9 := GenS1{F0: 155, F1: (59 <= 246)}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 545
}
func genFn1(p0 int64, p1 int8) int32 {
    if (792 < 842) {
        p0 += (p0 ^ (p0 | p0))
    } else {
        p1 = 1
        fmt.Println(p1, p1)
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint32 = 998
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 0, F1: true}
        switch fi10 {
        case -691:
            fmt.Println(p0)
        case -9223372036854775808:
            if true {
                var v14 string = ("\n\t" + "foo bar")
                switch v14 {
                case "\n\t":
                    fmt.Println(fi10, p0, v14)
                case "":
                    v15 := GenS1{F0: 128, F1: (!true)}
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        switch (fi16 ^ (-(fi10))) {
                        case -320:
                            v14 = ((v14 + ((v14 + "hello") + "\n\t")) + "a")
                        }
                        if false {
                            p0 = 812
                            var v17 int = (fi10 - -9223372036854775808)
                            _ = v17
                        }
                        p0 = 723
                        var v18 int16 = 484
                        _ = fi16
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int16 = 42
                        v21 := []int{32768}
                        v22 := GenS1{F0: 127, F1: false}
                        p1 -= p1
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
                _ = v14
            } else {
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 uint64 = 420
                    _ = fi23
                    _ = v24
                }
            }
        }
        _ = fi10
        _ = v13
    }
    var v25 uint8 = 255
    p1 = (p1 << ((18446744073709551615) & 63))
    var v26 int32 = -364
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    return v26
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 uint8 = 226
    var v4 int16 = 49
    var v5 int = 109
    v6 := make(map[string]uint32)
    switch v5 {
    case 127:
        fmt.Println(v4, v4)
    case 465:
        v27 := GenS2{F0: -1, F1: -1, F2: 42}
        _ = v27
    }
    if false {
        v28 := GenS3{F0: (v5 != (v5 + v5)), F1: 0, F2: 42}
        fmt.Println(v3)
        var v29 float32 = 1.5
        _ = v28
        _ = v29
    }
    var v30 float64 = -2.25
    for fi31 := 0; fi31 < 1; fi31++ {
        v4 = v4
        v4 = (v4 ^ (v4 >> ((4294967295) & 63)))
        var v32 string = "foo bar"
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := GenS4{F0: v4}
            _ = fi34
            _ = v35
        }
        if false {
            fmt.Println(v5)
        }
        fmt.Println(v3, fi33, v5)
        _ = fi33
    }
    v36 := []int{}
    fmt.Println(v30, v3, v3)
    v37 := GenS5{F0: "\n\t", F1: 2}
    for fi38 := 0; fi38 < 15; fi38++ {
        var v39 int16 = v4
        v40 := GenS5{F0: "a", F1: 913}
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := v1[1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
