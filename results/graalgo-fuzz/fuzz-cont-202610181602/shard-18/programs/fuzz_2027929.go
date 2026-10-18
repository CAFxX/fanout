package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int8 = 90
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]bool)
    v4 := []int{32767, 923}
    var v5 uint32 = v2
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint8 = 42
        _ = fi6
        _ = v7
    }
    if (42 > 42) {
        v8 := GenS1{F0: v1, F1: 0}
        if true {
            v9 := GenS2{F0: v5}
            var v10 string = "\n\t"
            fmt.Println(v1, v2)
            _ = v10
            _ = v9
        } else {
            var v11 string = ((("a" + "\n\t") + "foo bar") + (("" + ("foo bar" + "")) + ("hello" + ("foo bar" + "hello"))))
            _ = v11
        }
        if (true || false) {
            switch "foo bar" {
            case "\n\t":
                v1 = v1
            case "a":
                for fi12 := 0; fi12 < 2; fi12++ {
                    if (1 >= 128) {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            if (!(!(3.14159 != (3.14159 * float32((0.0 + (3.14159 * 1e10))))))) {
                                if (90 > 0) {
                                    var v14 float64 = 0.1
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        v15 := GenS1{F0: (v1 << ((2) & 63)), F1: 201}
                                        _ = v15
                                    default:
                                        var v16 bool = true
                                        _ = v16
                                    }
                                    fmt.Println(fi12, v1)
                                    var v17 bool = (!false)
                                    _ = v14
                                    _ = v17
                                }
                                v18 := GenS3{F0: 2, F1: -266, F2: fi13}
                                var v19 int8 = 29
                                v19 -= 1
                                _ = v18
                                _ = v19
                            } else {
                                var v20 int64 = -343
                                var v21 float64 = (float64(float32(1e10)) * (((float64(1e10) + 1.5) + 0.1) - 0.1))
                                v20 = v20
                                _ = v20
                                _ = v21
                            }
                            v5 += 2147483647
                            var v22 int = -218
                            _ = fi13
                            _ = v22
                        }
                    } else {
                        if (2 < 181) {
                            v23 := GenS4{F0: (v2 - 838)}
                            _ = v23
                        } else {
                            switch (fi12 << ((255) & 63)) {
                            case -2147483648:
                                if (false || (!true)) {
                                    var v24 uint16 = 545
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 20; fi25++ {
                                        v26 := GenS5{F0: 42}
                                        v27 := GenS5{F0: 156}
                                        if false {
                                            v2 = v2
                                            _ = len(v4)
                                        } else {
                                            if true {
                                                if true {
                                                    var v28 uint16 = 2
                                                    v28 -= v28
                                                    _ = v28
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v1 += 88
                                                    case "a":
                                                        v29 := GenS4{F0: v2}
                                                        _ = v29
                                                    }
                                                    var v30 float64 = 0.1
                                                    _ = v30
                                                }
                                                v31 := GenS6{F0: 0}
                                                _ = v31
                                            }
                                            var v32 int = -2147483648
                                            v1 -= v1
                                            _ = v32
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (4294967295 <= 1) {
                                                if false {
                                                    _ = len(v3)
                                                    var v33 int16 = -517
                                                    var v34 float64 = -2.25
                                                    var v35 float64 = ((((v34 - v34) * v34) * -2.25) / float64(3.14159))
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v1 = v1
                                                    var v36 int = (fi25 & 0)
                                                    v37 := GenS2{F0: v2}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            } else {
                                                switch fi12 {
                                                case 790:
                                                    v38 := GenS1{F0: v1, F1: 1}
                                                    _ = v38
                                                case 128:
                                                    var v39 float64 = 3.14159
                                                    _ = v39
                                                }
                                            }
                                        }
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                    if false {
                                        if (42 > v2) {
                                            var v40 int32 = 0
                                            _ = v40
                                        }
                                    } else {
                                        fmt.Println(fi12, v2, v1)
                                        v5 = v5
                                    }
                                    v2 = v5
                                }
                            case -1:
                                var v41 int = (fi12 + (fi12 << ((0) & 63)))
                                _ = v41
                            default:
                                var v42 uint = 1
                                _ = v42
                            }
                        }
                        for fi43 := 0; fi43 < 4; fi43++ {
                            fmt.Println(fi43, v5, v5)
                            var v44 uint64 = 18446744073709551615
                            v4 = append(v4, (-(fi12)))
                            _ = fi43
                            _ = v44
                        }
                    }
                    _ = fi12
                }
            }
        }
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v45 uint = 866
        _ = v45
    case "foo bar":
        for fi46 := 0; fi46 < 2; fi46++ {
            if (-2.25 > 3.14159) {
                if (177 <= 230) {
                    var v47 uint8 = 127
                    _ = v47
                }
                fmt.Println(v1, fi46)
                var v48 int = 2147483647
                _ = v48
            }
            if (v1 > v1) {
                var v49 float64 = 1e10
                switch fi46 {
                case 244:
                    fmt.Println(fi46, v2)
                case 0:
                    var v50 uint16 = 2
                    _ = v50
                case 42:
                    var v51 uint32 = (v5 & (v2 << ((uint(v2)) & 63)))
                    _ = v51
                }
                _ = v49
            }
            _ = fi46
        }
    case "hello":
        for fi52 := 0; fi52 < 10; fi52++ {
            var v53 uint64 = 2
            _ = fi52
            _ = v53
        }
    default:
        for fi54 := 0; fi54 < 6; fi54++ {
            v1 -= v1
            v55 := GenS5{F0: 15}
            v2 -= v2
            _ = fi54
            _ = v55
        }
    }
    var v56 uint16 = 2
    delete(v3, 2147483648)
    switch "foo bar" {
    case "a":
        v56 = 255
    default:
        if false {
            v2 += 0
            _ = len(v4)
        }
    }
    v4 = append(v4, -676)
    switch 4294967295 {
    case 1:
        var v57 int32 = 0
        _ = v57
    }
    fmt.Println(v2)
    var v58 int = 958
    switch 4294967295 {
    case 2147483647:
        delete(v3, (^(v5)))
    }
    var v59 int64 = -469
    v60 := v3[1]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v59
    _ = v60
}
