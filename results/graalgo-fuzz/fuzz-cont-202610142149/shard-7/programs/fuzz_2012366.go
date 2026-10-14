package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{42, 32767}
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    if true {
        v5 := GenS1{F0: -2147483648}
        switch ("hello" + "hello") {
        case "\n\t":
            for fi6 := 0; fi6 < 5; fi6++ {
                _ = len(v3)
                for fi7 := 0; fi7 < 6; fi7++ {
                    switch -190 {
                    case -12:
                        switch 0 {
                        case -1:
                            fmt.Println(fi7)
                        case 128:
                            var v8 uint32 = 972
                            _ = v8
                        default:
                            var v9 int32 = 0
                            _ = v9
                        }
                    case 128:
                        var v10 int32 = -163
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v11 := GenS1{F0: v10}
                        _ = v10
                        _ = v11
                    case -47:
                        v12 := v3[65535]
                        _ = v12
                    }
                    v13 := GenS1{F0: -443}
                    if ((1 <= 1) || ((fi7 * (fi6 - (fi6 * (fi7 % ((fi6) | 1))))) >= fi7)) {
                        var v14 float32 = 0.0
                        if (128 < (fi6 & (fi7 - 127))) {
                            switch (fi7 >> ((65535) & 63)) {
                            case 1:
                                switch "\n\t" {
                                case "\n\t":
                                    if (!false) {
                                        v15 := GenS2{F0: 128}
                                        v14 = 0.1
                                        _ = v15
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(fi6, v14, fi6)
                                    }
                                case "a":
                                    v14 = -2.25
                                case "foo bar":
                                    fmt.Println(fi6, v14, fi7)
                                default:
                                    _ = len(v2)
                                }
                            case 32767:
                                v2 = append(v2, -1)
                            }
                            delete(v4, "a")
                            v14 = (0.0 / ((v14 + v14) / v14))
                        }
                        v16 := GenS2{F0: 127}
                        _ = v14
                        _ = v16
                    }
                    _ = fi7
                    _ = v13
                }
                _ = fi6
            }
        case "foo bar":
            var v17 string = "a"
            _ = v17
        case "":
            v4["foo bar"] = 1
        }
        _ = v5
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    v19 := GenS1{F0: -1}
    var v20 int64 = 1
    for fi21 := 0; fi21 < 11; fi21++ {
        if false {
            if false {
                var v22 float32 = 3.14159
                v22 = v22
                _ = v22
            } else {
                _ = len(v1)
            }
        }
        switch fi21 {
        case 2147483647:
            if false {
                var v23 uint16 = 42
                if true {
                    var v24 float32 = 3.14159
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v23 -= v23
                        var v26 uint64 = 1
                        var v27 int32 = 314
                        v1[65535] = 1
                        fmt.Println(v24, fi21)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 int32 = 0
                        var v30 uint32 = 0
                        v20 = 1
                        v30 -= (v30 | (^(v30)))
                        var v31 int32 = (v29 - v29)
                        var v32 int64 = v20
                        fmt.Println(v24)
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = v24
                } else {
                    var v33 string = "hello"
                    v34 := GenS3{F0: v23, F1: fi21, F2: -2147483648}
                    var v35 uint64 = 42
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = v23
            }
        case 127:
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 int64 = (v20 - (v20 << ((2) & 63)))
                v20 += -1
                _ = fi36
                _ = v37
            }
        }
        var v38 int8 = 127
        _ = fi21
        _ = v38
    }
    var v39 uint64 = 2
    v3[1] = ("a" + "\n\t")
    var v40 uint16 = 65535
    delete(v1, 0)
    for fi41 := 0; fi41 < 5; fi41++ {
        v40 -= 1
        var v42 uint64 = (v39 + 486)
        _ = fi41
        _ = v42
    }
    switch "\n\t" {
    case "foo bar":
        var v43 float32 = float32(1e10)
        _ = v43
    case "":
        v44 := GenS4{F0: (!(!(1 <= 1))), F1: 693}
        _ = v44
    case "\n\t":
        fmt.Println(v20, v40, v20)
    default:
        var v45 bool = false
        _ = v45
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 float64 = 3.14159
        _ = fi46
        _ = v47
    }
    var v48 int = -756
    if (v39 >= (v39 + 178)) {
        v20 -= v20
    }
    v20 = v20
    if false {
        fmt.Println(v39)
        v40 -= 876
    } else {
        v49 := GenS4{F0: (!(true || false)), F1: 211}
        v50 := v3[1]
        _ = v49
        _ = v50
    }
    v20 -= v20
    v20 = (v20 << ((271) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v48
}
