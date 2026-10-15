package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 uint }
type GenS6 struct { F0 int64; F1 int; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{787, 485}
    var v3 uint8 = 128
    var v4 int16 = 308
    v5 := make(map[int]string)
    var v6 float32 = (-2.25 * (0.0 * (float32((3.14159 / 0.1)) / 1.0)))
    var v7 bool = true
    fmt.Println(v3)
    v4 -= (v4 + (v4 & v4))
    v4 += (v4 | v4)
    fmt.Println(v4)
    var v8 uint8 = 255
    switch 0 {
    case 9223372036854775807:
        v6 += v6
    case 127:
        v9 := GenS1{F0: 2, F1: 42}
        _ = v9
    default:
        var v10 int32 = -1
        _ = v10
    }
    var v11 bool = v7
    for fi12 := 0; fi12 < 5; fi12++ {
        if v11 {
            var v13 uint32 = 42
            var v14 uint8 = 0
            if true {
                _ = len(v2)
                var v15 int = 675
                switch "" {
                case "\n\t":
                    v16 := GenS2{F0: (^(v15)), F1: v8}
                    _ = v16
                case "hello":
                    if v7 {
                        if false {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                var v18 uint = 255
                                _ = fi17
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v3 += 42
                                var v20 float32 = (v6 + 1e10)
                                v4 -= v4
                                _ = fi19
                                _ = v20
                            }
                        }
                        if false {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v11 = v11
                                v7 = ("hello" == ("" + ""))
                                _ = fi21
                            }
                        } else {
                            v2 = append(v2, v15)
                        }
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v14 = (v14 + 0)
                            _ = fi22
                        }
                    }
                }
                _ = v15
            }
            _ = v13
            _ = v14
        } else {
            v4 -= 42
        }
        _ = len(v2)
        v3 += 255
        _ = fi12
    }
    var v23 float32 = 0.0
    v11 = v11
    v2 = append(v2, 617)
    if (!(v3 != v3)) {
        v24 := v5[-812]
        if true {
            fmt.Println(v6, v24, v24)
        } else {
            if v11 {
                v24 = ""
                v6 -= v6
                v7 = true
            } else {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 uint8 = v8
                    switch 128 {
                    case 32768:
                        switch (v24 + v24) {
                        case "foo bar":
                            v27 := GenS3{F0: 1, F1: "foo bar"}
                            _ = v27
                        }
                    case 1:
                        fmt.Println(v11, v26, fi25)
                    case -2147483648:
                        var v28 float32 = 0.0
                        _ = v28
                    }
                    switch ((((("\n\t" + ("" + "")) + "") + "hello") + v24) + v24) {
                    case "a":
                        for fi29 := 0; fi29 < 9; fi29++ {
                            if false {
                                if ((!(false && (v24 != ""))) || (0 == 1)) {
                                    var v30 int32 = 0
                                    _ = v30
                                }
                                v4 = (v4 ^ -32768)
                                if false {
                                    fmt.Println(v7, v7)
                                    fmt.Println(fi25, v7, v11)
                                    switch fi29 {
                                    case -826:
                                        fmt.Println(v23)
                                    case 2147483647:
                                        delete(v5, fi29)
                                    default:
                                        switch v24 {
                                        case "hello":
                                            v8 += 31
                                        case "a":
                                            switch 4294967295 {
                                            case 4294967295:
                                                if ((float32(-2.25) * float32(3.14159)) <= float32(1e10)) {
                                                    v6 += v6
                                                    v11 = false
                                                } else {
                                                    var v31 int = (fi25 >> ((1) & 63))
                                                    if v7 {
                                                        switch (fi29 | -1) {
                                                        case 128:
                                                            v32 := GenS4{F0: 362}
                                                            _ = v32
                                                        case 479:
                                                            v33 := GenS1{F0: 0, F1: 42}
                                                            _ = v33
                                                        default:
                                                            v8 += 255
                                                        }
                                                        var v34 bool = (2 != 0)
                                                        _ = v34
                                                    }
                                                    _ = v31
                                                }
                                            case 1:
                                                var v35 uint64 = 525
                                                _ = v35
                                            default:
                                                v36 := GenS5{F0: 785, F1: v4, F2: 716}
                                                _ = v36
                                            }
                                        default:
                                            var v37 bool = v7
                                            _ = v37
                                        }
                                        switch (fi29 ^ fi25) {
                                        case -779:
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v3)
                                            case "foo bar":
                                                var v38 int = fi29
                                                _ = v38
                                            }
                                        }
                                        var v39 int64 = int64(v3)
                                        _ = v39
                                    }
                                }
                                v24 = (v24 + (v24 + (v24 + v24)))
                            }
                            v40 := GenS4{F0: (v4 >> ((0) & 63))}
                            v23 += 1.5
                            _ = fi29
                            _ = v40
                        }
                    case "foo bar":
                        fmt.Println(v11, v26)
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v41 bool = v11
                        _ = v41
                    }
                    _ = fi25
                    _ = v26
                }
            }
        }
        v42 := GenS6{F0: 42, F1: 128, F2: -450}
        _ = v24
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
