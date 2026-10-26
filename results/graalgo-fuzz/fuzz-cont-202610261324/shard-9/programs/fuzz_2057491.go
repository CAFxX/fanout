package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    v3 := make(map[uint32]int)
    fmt.Println(v2)
    delete(v3, 2)
    v1[1] = v2
    for fi4 := 0; fi4 < 8; fi4++ {
        v2 = v2
        _ = fi4
    }
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v2 = (v2 + v2)
        if true {
            switch fi6 {
            case 127:
                switch (v2 + v2) {
                case "\n\t":
                    switch fi6 {
                    case 1:
                        switch v2 {
                        case "a":
                            v2 = v2
                        case "hello":
                            for fi7 := 0; fi7 < 18; fi7++ {
                                if true {
                                    v5 = v5
                                }
                                _ = len(v3)
                                _ = fi7
                            }
                        case "foo bar":
                            var v8 uint8 = v5
                            _ = v8
                        default:
                            v2 = "\n\t"
                        }
                    case 32768:
                        if false {
                            switch v2 {
                            case "a":
                                switch (fi6 - fi6) {
                                case 0:
                                    var v9 int = fi6
                                    _ = v9
                                case -9223372036854775808:
                                    v10 := GenS1{F0: 1, F1: uint64(v5), F2: (fi6 + (^(fi6)))}
                                    _ = v10
                                case 127:
                                    var v11 int8 = 42
                                    _ = v11
                                default:
                                    for fi12 := 0; fi12 < 12; fi12++ {
                                        v2 = v2
                                        _ = fi12
                                    }
                                }
                            case "\n\t":
                                v13 := []int{fi6, (-(fi6))}
                                _ = v13
                            }
                            v14 := []int{}
                            for fi15 := 0; fi15 < 15; fi15++ {
                                var v16 uint8 = 77
                                _ = fi15
                                _ = v16
                            }
                            _ = v14
                        }
                    default:
                        v5 = v5
                    }
                case "hello":
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v18 := []int{}
                        _ = fi17
                        _ = v18
                    }
                }
            default:
                if false {
                    var v19 int = -1
                    v20 := GenS2{F0: v19, F1: 100}
                    _ = v19
                    _ = v20
                } else {
                    var v21 int64 = 42
                    var v22 uint = 2
                    var v23 string = ("a" + (v2 + ("" + "a")))
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        } else {
            var v24 int = (fi6 | (fi6 + fi6))
            v24 = v24
            v24 = 127
            _ = v24
        }
        _ = fi6
    }
    v25 := GenS2{F0: 128, F1: -75}
    v2 = ""
    v26 := GenS1{F0: 623, F1: 347, F2: 464}
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 int8 = 4
        var v29 int64 = -1
        _ = fi27
        _ = v28
        _ = v29
    }
    switch -1 {
    case 32767:
        switch v2 {
        case "":
            if (0 <= 245) {
                v30 := v3[65535]
                for fi31 := 0; fi31 < 14; fi31++ {
                    switch ("hello" + (v2 + (v2 + v2))) {
                    case "":
                        var v32 int16 = 1
                        _ = v32
                    case "\n\t":
                        switch ("foo bar" + v2) {
                        case "":
                            if (false || (fi31 < v30)) {
                                v2 = "\n\t"
                                for fi33 := 0; fi33 < 13; fi33++ {
                                    fmt.Println(v2)
                                    _ = fi33
                                }
                                var v34 int64 = -9223372036854775808
                                _ = v34
                            }
                        }
                    default:
                        v2 = ("" + v2)
                    }
                    fmt.Println(v5, v2)
                    fmt.Println(v2, v5, v30)
                    _ = fi31
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 uint32 = 916
                    for fi37 := 0; fi37 < 14; fi37++ {
                        var v38 bool = false
                        _ = fi37
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                }
                _ = v30
            } else {
                if (255 != 42) {
                    if false {
                        var v39 uint16 = 149
                        fmt.Println(v2)
                        v40 := GenS3{F0: 255}
                        _ = v39
                        _ = v40
                    } else {
                        fmt.Println(v5)
                        switch "a" {
                        case "a":
                            fmt.Println(v2, v2)
                        }
                        v2 = "a"
                    }
                    for fi41 := 0; fi41 < 13; fi41++ {
                        v2 = v2
                        _ = fi41
                    }
                }
                fmt.Println(v5, v2)
            }
        case "hello":
            fmt.Println(v2)
        }
    }
    var v42 int8 = -128
    v43 := []int{}
    var v44 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
