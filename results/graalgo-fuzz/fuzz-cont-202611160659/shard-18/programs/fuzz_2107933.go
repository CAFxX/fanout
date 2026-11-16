package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 0
    v3 := make(map[int]string)
    switch 4294967295 {
    case -1:
        v4 := []int{4294967295, 127, 128, 128}
        _ = v4
    }
    v2 = (v2 >> ((4294967295) & 63))
    v5 := GenS1{F0: -128}
    v2 = (v2 | (v2 >> ((18446744073709551615) & 63)))
    var v6 uint32 = 562
    v7 := []int{32768}
    var v8 uint8 = v2
    if true {
        var v9 int16 = -1
        var v10 uint32 = v6
        switch "hello" {
        case "\n\t":
            v6 = v10
        case "hello":
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    var v12 int = (fi11 * -2147483648)
                    v6 = uint32(v2)
                    v10 -= v6
                    _ = v12
                }
                v9 += (v9 & v9)
                _ = fi11
            }
        case "foo bar":
            v2 += v8
        default:
            v13 := GenS1{F0: 1}
            _ = v13
        }
        _ = v10
        _ = v9
    } else {
        switch "hello" {
        case "hello":
            if false {
                var v14 bool = (!false)
                _ = v14
            }
        case "a":
            if ((352 != 32767) && ((1.5 + 1.5) != -2.25)) {
                fmt.Println(v6, v6)
                var v15 int32 = 42
                var v16 uint64 = 717
                _ = v15
                _ = v16
            }
        case "foo bar":
            switch "" {
            case "\n\t":
                v17 := GenS2{F0: 2147483647, F1: 42, F2: "a"}
                _ = v17
            default:
                v18 := GenS3{F0: -1, F1: 872}
                _ = v18
            }
        default:
            _ = len(v1)
        }
        v19 := GenS4{F0: -32768, F1: -840}
        v2 += (v8 | v2)
        _ = v19
    }
    if len(v7) > 0 {
        v7[0] = 42
    }
    if (!((!true) && (65535 != 65535))) {
        var v20 uint32 = 1
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int64 = int64(v20)
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        if false {
            v23 := v3[22]
            _ = v23
        }
        v24 := GenS2{F0: -751, F1: 645, F2: "a"}
        _ = v24
    }
    v2 += (v2 & v2)
    var v25 uint64 = 84
    fmt.Println(v6)
    if true {
        v26 := GenS5{F0: -175}
        var v27 int16 = 0
        _ = v26
        _ = v27
    }
    v28 := GenS6{F0: 4294967295}
    if (100 != -63) {
        v29 := GenS1{F0: 0}
        if ((v8 / ((v8) | 1)) != v8) {
            fmt.Println(v2, v25, v2)
            switch ("a" + "hello") {
            case "\n\t":
                switch "foo bar" {
                case "":
                    var v30 uint = 18446744073709551615
                    _ = v30
                case "\n\t":
                    if (v25 > v25) {
                        if true {
                            if true {
                                v25 += (v25 >> ((0) & 63))
                            } else {
                                var v31 uint32 = v6
                                _ = v31
                            }
                            var v32 string = "a"
                            _ = v32
                        }
                    } else {
                        v6 += v6
                    }
                default:
                    var v33 float32 = 1e10
                    _ = v33
                }
            case "":
                if (225 == 1) {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v25 += v25
                        v35 := GenS4{F0: -583, F1: 171}
                        _ = fi34
                        _ = v35
                    }
                    v25 = v25
                    v1["hello"] = "\n\t"
                } else {
                    delete(v1, "foo bar")
                }
            default:
                for fi36 := 0; fi36 < 11; fi36++ {
                    var v37 int16 = 1000
                    fmt.Println(v8, v2, v6)
                    var v38 uint = 941
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
        }
        _ = v29
    }
    v8 = (v2 % (((v2 | v8)) | 1))
    v25 = (v25 % ((uint64(v2)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
