package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }

func main() {
    var v1 uint = 0
    var v2 uint = (v1 | v1)
    var v3 string = ("hello" + "a")
    v4 := make(map[uint32]string)
    var v5 int8 = 10
    for fi6 := 0; fi6 < 17; fi6++ {
        switch (fi6 >> ((v2) & 63)) {
        case 4294967295:
            v5 = v5
        }
        v1 = v2
        v2 -= v1
        _ = fi6
    }
    v7 := GenS1{F0: 1, F1: -88, F2: (127 >= int(v5))}
    fmt.Println(v5, v1)
    _ = len(v4)
    switch 32767 {
    case 549:
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{(fi8 * (-(fi8)))}
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 int32 = -792
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 42:
        switch "hello" {
        case "":
            if true {
                var v12 uint32 = 869
                v3 = ""
                v12 = uint32(v2)
                _ = v12
            }
        }
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            v4[2] = (((v3 + v3) + v3) + "foo bar")
            v14 := GenS1{F0: 42, F1: (^(v5)), F2: false}
            v5 = (v5 | v5)
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        switch fi15 {
        case 306:
            switch (((v3 + ("foo bar" + "a")) + v3) + v3) {
            case "hello":
                var v16 uint64 = 0
                _ = v16
            }
        case 0:
            v3 = ("foo bar" + v3)
        case 2147483647:
            v17 := GenS1{F0: 42, F1: (v5 ^ 127), F2: (!true)}
            _ = v17
        default:
            var v18 int = 750
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: 0, F2: true}
        v3 = (((v3 + v3) + (v3 + ((v3 + (v3 + v3)) + v3))) + "foo bar")
        _ = fi15
        _ = v19
    }
    if true {
        v1 += (v1 ^ 0)
        v1 += 768
        v20 := GenS2{F0: 700, F1: (v5 / ((v5) | 1))}
        _ = v20
    } else {
        var v21 uint = (^(v1))
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 13; fi23++ {
                if false {
                    v1 = 453
                    fmt.Println(fi23)
                    switch (v3 + v3) {
                    case "foo bar":
                        if true {
                            if false {
                                v24 := []int{-57, fi22}
                                v5 = v5
                                _ = v24
                            }
                        } else {
                            v25 := GenS1{F0: -189, F1: (v5 + v5), F2: true}
                            var v26 uint16 = 0
                            var v27 float64 = (float64(float32(-2.25)) / ((float64(-2.25) / 0.1) - (((1.5 * -2.25) - -2.25) * 0.0)))
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case "a":
                        var v28 int = -1
                        _ = v28
                    default:
                        var v29 uint64 = 1
                        _ = v29
                    }
                    var v30 float64 = (1.5 + -2.25)
                    _ = v30
                } else {
                    v31 := v4[962]
                    v2 = v2
                    _ = v31
                }
                var v32 float32 = 3.14159
                _ = fi23
                _ = v32
            }
            switch fi22 {
            case 2147483647:
                if true {
                    fmt.Println(v21)
                } else {
                    v33 := GenS1{F0: 0, F1: v5, F2: (!true)}
                    _ = v33
                }
            default:
                if (32767 > (fi22 & fi22)) {
                    v34 := []int{(fi22 ^ fi22), fi22}
                    v35 := GenS1{F0: -1, F1: 100, F2: (1 != int32(v5))}
                    _ = v34
                    _ = v35
                } else {
                    v36 := []int{238, fi22}
                    if true {
                        var v37 uint16 = 42
                        var v38 float32 = 1e10
                        v39 := GenS2{F0: 297, F1: v5}
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = v36
                }
            }
            var v40 uint16 = 407
            _ = fi22
            _ = v40
        }
        switch -358 {
        case 32768:
            if false {
                if false {
                    var v41 uint32 = 1
                    _ = v41
                } else {
                    delete(v4, 2147483647)
                    var v42 bool = (1e10 <= 3.14159)
                    v4[185] = (v3 + ((v3 + ("hello" + v3)) + (v3 + "\n\t")))
                    _ = v42
                }
            }
        case 4294967295:
            var v43 int8 = v5
            _ = v43
        case -450:
            var v44 float64 = float64(0.0)
            _ = v44
        default:
            v4[754] = "a"
        }
        _ = v21
    }
    v2 = v2
    v45 := GenS3{F0: "\n\t", F1: false, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v7
}
