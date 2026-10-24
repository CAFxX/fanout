package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 int32 = -2147483648
    var v2 int = 1
    var v3 int64 = -443
    var v4 int32 = v1
    var v5 int8 = int8(v3)
    var v6 uint = 255
    v7 := GenS1{F0: 2}
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v8 int32 = v1
        _ = v8
    }
    if true {
        if false {
            var v9 int8 = v5
            _ = v9
        }
        v10 := GenS2{F0: 29, F1: (v2 >> (((v6 >> ((v6) & 63))) & 63)), F2: 318}
        for fi11 := 0; fi11 < 18; fi11++ {
            if true {
                v12 := GenS3{F0: 2, F1: -32768, F2: (fi11 >> (((^(v6))) & 63))}
                v5 = (v5 << ((18446744073709551615) & 63))
                _ = v12
            }
            _ = fi11
        }
        _ = v10
    }
    v13 := GenS4{F0: (v2 & v2)}
    v3 = 0
    v4 = v4
    v1 = v1
    var v14 int32 = v1
    var v15 uint64 = 18446744073709551615
    v1 = (v14 / ((v4) | 1))
    v2 -= v2
    fmt.Println(v5, v5)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint = (v6 % ((v6) | 1))
        if (127 >= 163) {
            if true {
                if false {
                    v15 = v15
                    v15 = 307
                    var v18 uint16 = 1
                    _ = v18
                } else {
                    fmt.Println(v14, v6)
                    var v19 uint32 = 1
                    v3 -= v3
                    _ = v19
                }
            } else {
                var v20 uint8 = 255
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int64 = (^(v3))
                    if true {
                        switch (^(fi21)) {
                        case -838:
                            switch "foo bar" {
                            case "foo bar":
                                v6 = v17
                            case "hello":
                                v23 := GenS3{F0: uint32(v20), F1: -32768, F2: 2147483647}
                                _ = v23
                            case "a":
                                v24 := []int{}
                                _ = v24
                            default:
                                v25 := GenS5{F0: ((!true) || false), F1: (v15 * (v15 & 2))}
                                _ = v25
                            }
                        case 4294967295:
                            switch "a" {
                            case "":
                                if (v14 >= v4) {
                                    fmt.Println(v15)
                                    if false {
                                        v26 := GenS4{F0: (fi16 + fi21)}
                                        v27 := GenS5{F0: false, F1: (^(v15))}
                                        _ = v26
                                        _ = v27
                                    }
                                    v15 = v15
                                }
                            default:
                                switch fi16 {
                                case 32767:
                                    v28 := GenS6{F0: 2147483647, F1: (v5 << (((v17 - v6)) & 63))}
                                    _ = v28
                                case 42:
                                    var v29 float32 = (-2.25 + 0.1)
                                    v15 = (v15 | (v15 % ((42) | 1)))
                                    var v30 uint16 = 0
                                    _ = v29
                                    _ = v30
                                }
                            }
                        case -2147483648:
                            v1 = 0
                        default:
                            fmt.Println(v14, v14)
                        }
                    } else {
                        var v31 float32 = (-2.25 * 3.14159)
                        var v32 float32 = 1.5
                        _ = v31
                        _ = v32
                    }
                    switch (fi16 % ((v2) | 1)) {
                    case 42:
                        switch fi16 {
                        case 32767:
                            var v33 int32 = (-(v14))
                            _ = v33
                        case -349:
                            v6 -= v6
                        case 32768:
                            v34 := []int{(-(fi21))}
                            _ = v34
                        default:
                            v35 := GenS7{F0: (v15 | 2)}
                            _ = v35
                        }
                    case 2147483647:
                        var v36 int8 = (v5 >> (((v6 % (((v6 * 42)) | 1))) & 63))
                        _ = v36
                    default:
                        var v37 uint8 = v20
                        _ = v37
                    }
                    var v38 uint32 = uint32(v6)
                    fmt.Println(v4, v17, v17)
                    _ = fi21
                    _ = v22
                    _ = v38
                }
                _ = v20
            }
            v5 += (v5 | v5)
        } else {
            v39 := GenS4{F0: -428}
            _ = v39
        }
        _ = fi16
        _ = v17
    }
    if (v3 > (v3 | (v3 * (-(v3))))) {
        if true {
            v14 += int32(v2)
            var v40 int64 = (-(v3))
            _ = v40
        } else {
            v3 = 42
            v41 := GenS3{F0: 42, F1: -32768, F2: v2}
            _ = v41
        }
    }
    v3 = (v3 + int64(v2))
    var v42 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
