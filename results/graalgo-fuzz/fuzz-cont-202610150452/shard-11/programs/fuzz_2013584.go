package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 5; fi6++ {
        if (42 <= int8(fi6)) {
            v1 += (v1 >> ((65535) & 63))
            if (!false) {
                fmt.Println(fi6, v1)
                var v7 float32 = v4
                _ = v7
            } else {
                if ((false || false) && (1.5 < float64((v4 / 1e10)))) {
                    if false {
                        fmt.Println(v4, v1, fi6)
                        if (!false) {
                            for fi8 := 0; fi8 < 16; fi8++ {
                                fmt.Println(v1, v4)
                                _ = fi8
                            }
                            var v9 uint8 = 255
                            _ = v9
                        } else {
                            if true {
                                v10 := GenS1{F0: 1}
                                _ = v10
                            }
                        }
                        v4 -= 1.5
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v5)
                    case "":
                        var v11 uint64 = 18446744073709551615
                        _ = v11
                    }
                }
                var v12 bool = (!false)
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 += 42
                fmt.Println(fi6)
                v14 := GenS1{F0: 2147483647}
                var v15 string = "a"
                _ = fi13
                _ = v14
                _ = v15
            }
            v4 += v4
            switch fi6 {
            case 32768:
                v16 := GenS1{F0: 2147483647}
                _ = v16
            case 0:
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(fi17, v4, fi6)
                    _ = fi17
                }
            case -9223372036854775808:
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS1{F0: 42}
                    v4 += v4
                    var v20 int = (fi6 - fi6)
                    if ((v1 << ((4294967295) & 63)) == (v1 & v1)) {
                        fmt.Println(v1, fi18)
                    }
                    switch "hello" {
                    case "a":
                        var v21 uint64 = 792
                        v4 -= float32(3.14159)
                        if true {
                            v3[v1] = "\n\t"
                            v22 := []int{fi6}
                            switch fi18 {
                            case 706:
                                var v23 uint8 = 128
                                _ = v23
                            case 378:
                                v24 := v2["hello"]
                                _ = v24
                            }
                            _ = v22
                        } else {
                            v25 := []int{(v20 << ((255) & 63))}
                            v21 += 4294967295
                            v26 := GenS1{F0: -1}
                            v21 += (^(v21))
                            _ = v25
                            _ = v26
                        }
                        _ = v21
                    default:
                        var v27 uint32 = (v1 | 242)
                        _ = v27
                    }
                    var v28 int32 = 2147483647
                    fmt.Println(v1, fi18, v4)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v28
                }
            }
        }
        var v29 int32 = 2147483647
        _ = fi6
        _ = v29
    }
    var v30 uint64 = 499
    switch 42 {
    case 42:
        v31 := []int{803, 917}
        _ = v31
    case 792:
        if true {
            v30 -= v30
            v32 := GenS1{F0: 216}
            _ = v32
        }
    case 2147483647:
        v33 := GenS2{F0: 1}
        _ = v33
    default:
        switch "a" {
        case "a":
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int16 = -32768
                _ = fi34
                _ = v35
            }
        }
    }
    if true {
        switch 1 {
        case 127:
            fmt.Println(v30, v1, v1)
        case -9223372036854775808:
            v36 := []int{-907, -959}
            _ = v36
        }
        switch "" {
        case "\n\t":
            v37 := v5[483]
            _ = v37
        case "a":
            switch 9223372036854775807 {
            case -2147483648:
                fmt.Println(v1, v30, v1)
            case 862:
                if true {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v39 := []int{(fi38 >> ((255) & 63)), (fi38 - fi38)}
                        var v40 float32 = v4
                        if true {
                            var v41 uint16 = 575
                            v1 = v1
                            _ = v41
                        } else {
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                } else {
                    if ((!(!true)) || (("" + "hello") < "")) {
                        fmt.Println(v30, v1, v1)
                    } else {
                        var v42 float32 = v4
                        var v43 int32 = -993
                        v4 -= (v4 + 0.1)
                        _ = v42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 8; fi44++ {
                        var v45 uint64 = v30
                        _ = fi44
                        _ = v45
                    }
                    var v46 bool = (!(("foo bar" + ("" + "foo bar")) == "hello"))
                    _ = v46
                }
            }
        case "hello":
            fmt.Println(v30, v1)
        }
        v4 += (1.5 - (v4 - (v4 - 1.5)))
    }
    if (-1 >= -52) {
        v47 := GenS3{F0: -1, F1: 42, F2: 453}
        if true {
            v1 = 65535
            v48 := GenS1{F0: -244}
            v49 := GenS4{F0: 9223372036854775807, F1: (true && ((false || false) || true))}
            _ = v48
            _ = v49
        }
        var v50 float64 = 1.5
        _ = v47
        _ = v50
    } else {
        if ((!true) && true) {
            v51 := []int{4294967295}
            v4 += float32(0.1)
            v52 := GenS5{F0: 1000}
            _ = v51
            _ = v52
        } else {
            if (1e10 != (3.14159 / 1.0)) {
                var v53 float32 = -2.25
                var v54 int16 = -1
                if (-2.25 > (0.1 / 1.0)) {
                    v30 -= v30
                }
                _ = v53
                _ = v54
            } else {
                var v55 bool = true
                _ = v55
            }
        }
        var v56 uint8 = 1
        v57 := GenS1{F0: 42}
        _ = v56
        _ = v57
    }
    var v58 float32 = -2.25
    for fi59 := 0; fi59 < 12; fi59++ {
        v58 = 3.14159
        var v60 uint64 = 42
        _ = fi59
        _ = v60
    }
    _ = len(v2)
    v30 = (v30 << ((0) & 63))
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v58
}
