package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    v1 := []uint32{587}
    v2 := []int{-1, -395, 32768, 42}
    var v3 float32 = 0.0
    var v4 uint16 = 0
    var v5 int = 464
    var v6 string = ""
    v3 = (1e10 - (v3 + 0.1))
    if (!false) {
        v3 = 1.5
        fmt.Println(v3)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = -2.25
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int = v5
        fmt.Println(v7, v7, v5)
        _ = fi8
        _ = v9
    }
    if false {
        var v10 int32 = -14
        v10 = v10
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v7, v4, v7)
        v7 += v7
        if true {
            v7 = -1
            v5 = -479
            fmt.Println(v5, v7, v7)
        } else {
            v12 := GenS1{F0: 223}
            for fi13 := 0; fi13 < 11; fi13++ {
                if true {
                    var v14 uint64 = 2
                    _ = v14
                } else {
                    var v15 int8 = 98
                    v16 := GenS1{F0: -1}
                    var v17 uint8 = 63
                    switch fi11 {
                    case 4294967295:
                        fmt.Println(fi11, v5)
                    }
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 uint16 = (v4 + 32767)
                var v19 string = v6
                var v20 uint16 = 255
                v6 = v6
                v7 = v7
                _ = len(v2)
                if ("\n\t" > "") {
                    v21 := GenS1{F0: 584}
                    var v22 uint8 = uint8(v4)
                    var v23 bool = false
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS2{F0: "foo bar", F1: v20, F2: false}
                    _ = v24
                }
                var v25 uint16 = (^(v18))
                _ = fi13
                _ = v18
                _ = v19
                _ = v20
                _ = v25
            }
            var v26 uint8 = 0
            _ = v12
            _ = v26
        }
        _ = fi11
    }
    if ((42 != 328) || true) {
        v6 = "hello"
        var v27 float64 = ((0.1 * ((0.0 / 3.14159) * (3.14159 - 1e10))) - float64((float32(0.0) * ((0.1 - float32((3.14159 + 0.1))) / (v3 / float32(-2.25))))))
        _ = v27
    } else {
        var v28 uint32 = 4294967295
        switch "a" {
        case "\n\t":
            if false {
                var v29 bool = (false && false)
                if v29 {
                    switch "a" {
                    case "":
                        for fi30 := 0; fi30 < 10; fi30++ {
                            var v31 uint = 1
                            _ = fi30
                            _ = v31
                        }
                    default:
                        fmt.Println(v4)
                    }
                    var v32 int = v5
                    v2 = append(v2, (-(v5)))
                    _ = v32
                }
                _ = v29
            } else {
                v7 = v7
                if ((true && (!(((((v7 >> ((65535) & 63)) <= 32767) && true) && (0 < 2)) && true))) || (!false)) {
                    var v33 int8 = 1
                    _ = v33
                } else {
                    var v34 uint64 = 849
                    var v35 uint = 0
                    v28 = (v28 - v28)
                    _ = v34
                    _ = v35
                }
            }
        case "foo bar":
            v7 = 42
        case "hello":
            if true {
                if false {
                    v6 = "a"
                    var v36 int = 4294967295
                    var v37 float64 = float64(3.14159)
                    _ = v36
                    _ = v37
                }
                v7 -= (v7 + v7)
                switch ((("hello" + "a") + v6) + (v6 + v6)) {
                case "foo bar":
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 bool = false
                        var v40 uint64 = 1
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                default:
                    fmt.Println(v7, v28, v5)
                }
            } else {
                fmt.Println(v4)
                _ = len(v1)
            }
        }
        _ = v28
    }
    switch (v6 + "hello") {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v2 = append(v2, v5)
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        switch "" {
        case "a":
            v42 := GenS3{F0: 0, F1: (v7 + (v7 % ((v7) | 1))), F2: 0}
            _ = v42
        default:
            if (!false) {
                var v43 int8 = 65
                for fi44 := 0; fi44 < 6; fi44++ {
                    var v45 int32 = -119
                    var v46 int8 = v43
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                if true {
                    v4 += v4
                    v5 += v5
                }
                _ = v43
            } else {
                v4 = v4
                fmt.Println(v5, v5, fi41)
            }
        }
        v6 = v6
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
