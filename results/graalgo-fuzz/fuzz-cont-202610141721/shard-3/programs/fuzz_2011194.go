package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 uint64; F2 uint64 }
func genFn0(p0 uint) int32 {
    var v5 bool = false
    var v6 string = "hello"
    for fi7 := 0; fi7 < 15; fi7++ {
        if (!true) {
            fmt.Println(p0, fi7)
            if v5 {
                v8 := []int{}
                _ = v8
            }
            switch ((v6 + v6) + ("hello" + v6)) {
            case "hello":
                v9 := GenS1{F0: 2}
                _ = v9
            case "foo bar":
                if v5 {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 uint64 = 20
                        if (1e10 < 0.0) {
                            v12 := []int{(fi10 | 127)}
                            fmt.Println(p0, fi10, fi7)
                            if v5 {
                                fmt.Println(fi10, fi10, v5)
                                switch (v6 + "") {
                                case "\n\t":
                                    if v5 {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v13 := GenS2{F0: 1, F1: (p0 & p0)}
                                            _ = v13
                                        case "a":
                                            if v5 {
                                                v14 := GenS3{F0: 1000}
                                                var v15 bool = true
                                                _ = v14
                                                _ = v15
                                            } else {
                                                var v16 float32 = ((float32(1.5) - (0.1 / 1.0)) / 1.0)
                                                var v17 uint8 = 127
                                                fmt.Println(v11, v5, v6)
                                                _ = v16
                                                _ = v17
                                            }
                                        default:
                                            fmt.Println(fi7)
                                        }
                                        var v18 float64 = 1e10
                                        v19 := GenS4{F0: 255}
                                        fmt.Println(v11)
                                        _ = v18
                                        _ = v19
                                    }
                                case "foo bar":
                                    _ = len(v12)
                                }
                                switch (v6 + "") {
                                case "":
                                    v20 := GenS2{F0: 0, F1: p0}
                                    _ = v20
                                case "a":
                                    var v21 int32 = -987
                                    _ = v21
                                case "hello":
                                    v5 = (4294967295 < 953)
                                default:
                                    if v5 {
                                        fmt.Println(fi10, v5, v6)
                                    } else {
                                        var v22 uint16 = uint16(v11)
                                        _ = v22
                                    }
                                }
                                switch fi10 {
                                case 1:
                                    switch v6 {
                                    case "a":
                                        if v5 {
                                            var v23 float64 = 0.1
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            _ = v23
                                        }
                                    case "":
                                        var v24 uint32 = 2147483648
                                        _ = v24
                                    case "foo bar":
                                        fmt.Println(v5)
                                    default:
                                        v25 := GenS5{F0: 0, F1: 2147483647}
                                        _ = v25
                                    }
                                    v26 := GenS6{F0: 127, F1: v5}
                                    _ = v26
                                case 0:
                                    var v27 string = (((v6 + v6) + (("hello" + "") + v6)) + "foo bar")
                                    _ = v27
                                case -722:
                                    p0 += (p0 % ((p0) | 1))
                                }
                            }
                            _ = v12
                        }
                        v28 := GenS7{F0: 1, F1: (v11 ^ (v11 | (v11 + v11)))}
                        _ = fi10
                        _ = v11
                        _ = v28
                    }
                    v5 = ((fi7 & (-(fi7))) > (-(fi7)))
                    var v29 float32 = 1.5
                    _ = v29
                }
            }
        }
        v6 = (v6 + ("foo bar" + v6))
        switch fi7 {
        case -217:
            v5 = true
        case 42:
            var v30 uint8 = 127
            _ = v30
        case 380:
            v31 := []int{515, (-(fi7))}
            _ = v31
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := []string{"hello", "a", ("hello" + "a"), "a"}
    var v2 float64 = 0.1
    var v3 int8 = 1
    var v4 int32 = -2147483648
    switch ("" + (("a" + "\n\t") + ("hello" + "a"))) {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        fmt.Println(v4)
    }
    fmt.Println(v4, v4, v2)
    v3 += v3
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS8{F0: "\n\t"}
        if (!(!(!(1 < 658)))) {
            switch (("a" + "\n\t") + "a") {
            case "a":
                for fi34 := 0; fi34 < 16; fi34++ {
                    v4 = 938
                    _ = fi34
                }
            case "":
                v1 = append(v1, ("a" + "\n\t"))
            case "\n\t":
                switch (fi32 | (fi32 % ((fi32) | 1))) {
                case 128:
                    if (v3 > (v3 >> ((255) & 63))) {
                        var v35 uint16 = 327
                        if true {
                            for fi36 := 0; fi36 < 5; fi36++ {
                                var v37 uint64 = 341
                                _ = fi36
                                _ = v37
                            }
                            var v38 uint = 65535
                            fmt.Println(fi32, v4, fi32)
                            _ = v38
                        } else {
                            if ((42 > 0) && true) {
                                for fi39 := 0; fi39 < 11; fi39++ {
                                    v40 := GenS9{F0: 2, F1: 1, F2: 736}
                                    var v41 int16 = 32767
                                    if (false || true) {
                                        var v42 int64 = 753
                                        v41 += v41
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v42
                                    }
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                v43 := GenS1{F0: (v35 + 42)}
                                fmt.Println(fi32, v3)
                                _ = v43
                            }
                        }
                        v44 := GenS6{F0: -1, F1: (925 == 1)}
                        _ = v35
                        _ = v44
                    } else {
                        var v45 float32 = 3.14159
                        v4 = v4
                        v46 := GenS3{F0: -1}
                        _ = v45
                        _ = v46
                    }
                }
            }
            var v47 uint = 65535
            if true {
                for fi48 := 0; fi48 < 17; fi48++ {
                    var v49 float32 = ((1.5 / float32((v2 / 3.14159))) / 3.14159)
                    v4 = (v4 & genFn0(735))
                    _ = fi48
                    _ = v49
                }
                v2 = (((((v2 * float64((0.0 / (((1e10 - (-2.25 / 1.0)) - 0.0) - (1.5 - 3.14159))))) / (v2 - ((v2 - (((-2.25 + v2) / 0.1) * v2)) + 1.5))) / float64(float32(1.5))) * 1e10) / v2)
                v4 -= int32(v3)
            } else {
                var v50 int16 = 1
                _ = v50
            }
            _ = v47
        }
        for fi51 := 0; fi51 < 19; fi51++ {
            v4 = v4
            _ = fi51
        }
        _ = fi32
        _ = v33
    }
    v3 += -128
    fmt.Println(v4)
    for fi52 := 0; fi52 < 6; fi52++ {
        if (1 > 42) {
            fmt.Println(v2)
        } else {
            v3 = -1
            v53 := GenS1{F0: 81}
            _ = v53
        }
        _ = fi52
    }
    var v54 float64 = v2
    v2 = (v54 + v54)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
