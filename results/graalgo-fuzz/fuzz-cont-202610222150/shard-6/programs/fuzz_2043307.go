package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0() int16 {
    var v5 float64 = ((float64(1e10) - -2.25) - 0.1)
    fmt.Println(v5)
    _ = v5
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    v3 := []int{-2147483648, -184, 585}
    var v4 int8 = 100
    for fi6 := 0; fi6 < 3; fi6++ {
        v4 += v4
        _ = fi6
    }
    v4 = (v4 >> ((255) & 63))
    var v7 int8 = 1
    if ((float32(((0.1 + 0.1) + (1e10 * -2.25))) + (1e10 - -2.25)) > (1e10 * (float32(0.1) + float32(-2.25)))) {
        fmt.Println(v7, v7)
        switch 32767 {
        case -196:
            var v8 int = int(v4)
            _ = v8
        case -9223372036854775808:
            v9 := GenS1{F0: 255}
            _ = v9
        default:
            v2 = append(v2, 2)
        }
    } else {
        var v10 int32 = 42
        for fi11 := 0; fi11 < 1; fi11++ {
            if (!("" > "")) {
                fmt.Println(v4, v4, fi11)
                v4 -= (^(v4))
                var v12 int32 = v10
                _ = v12
            }
            _ = fi11
        }
        var v13 uint32 = 65535
        _ = v10
        _ = v13
    }
    switch (((("a" + "hello") + (("a" + (("a" + "foo bar") + ("\n\t" + ""))) + "\n\t")) + "\n\t") + "\n\t") {
    case "a":
        v4 = -1
    case "":
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v7)
            _ = fi14
        }
    default:
        for fi15 := 0; fi15 < 12; fi15++ {
            if true {
                _ = len(v3)
                var v16 float32 = -2.25
                v4 -= v4
                _ = v16
            } else {
                var v17 int32 = 0
                for fi18 := 0; fi18 < 4; fi18++ {
                    if true {
                        var v19 string = (("" + "\n\t") + "a")
                        var v20 uint64 = 4294967295
                        if false {
                            v1["hello"] = (!(0 != 32767))
                            var v21 int64 = 468
                            v21 = 9223372036854775807
                            _ = v21
                        } else {
                            v19 = "foo bar"
                            v19 = (((v19 + "foo bar") + v19) + ((((("" + v19) + "\n\t") + (v19 + v19)) + "a") + v19))
                            v22 := GenS1{F0: 2}
                            _ = v22
                        }
                        v7 = -1
                        var v23 uint8 = 42
                        if true {
                            v7 += (v4 << ((215) & 63))
                        } else {
                            v17 = (v17 / ((v17) | 1))
                            v24 := v1[(v19 + v19)]
                            var v25 int = fi18
                            var v26 uint = 587
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = v19
                        _ = v20
                        _ = v23
                    }
                    if false {
                        v27 := GenS1{F0: 2}
                        _ = v27
                    } else {
                        switch -2147483648 {
                        case -426:
                            var v28 uint8 = 0
                            _ = v28
                        default:
                            if len(v3) > 0 {
                                v3[0] = fi18
                            }
                        }
                        v7 -= (v4 / (((v7 - v7)) | 1))
                        v29 := GenS2{F0: 133, F1: 32767}
                        _ = v29
                    }
                    _ = fi18
                }
                _ = v17
            }
            _ = fi15
        }
    }
    var v30 int16 = 788
    var v31 int8 = (v7 >> ((844) & 63))
    switch "" {
    case "foo bar":
        var v32 bool = (1 <= 1)
        _ = v32
    case "\n\t":
        var v33 bool = false
        _ = v33
    case "hello":
        v31 -= v7
    }
    v7 += (^(v31))
    fmt.Println(v31, v4)
    if true {
        for fi34 := 0; fi34 < 20; fi34++ {
            if (2147483647 <= 920) {
                if (true || false) {
                    var v35 uint16 = 434
                    v30 -= 42
                    _ = v35
                }
                _ = len(v2)
            }
            v4 = v4
            v1["a"] = (("hello" + "\n\t") > ("" + "foo bar"))
            _ = fi34
        }
        v36 := GenS2{F0: 2, F1: genFn0()}
        _ = v36
    } else {
        for fi37 := 0; fi37 < 2; fi37++ {
            v7 = -34
            v2 = append(v2, 2147483648)
            _ = fi37
        }
        v31 = v7
        v30 = (v30 % (((v30 ^ (v30 % ((463) | 1)))) | 1))
    }
    switch "" {
    case "hello":
        v38 := GenS3{F0: 128, F1: (^(v30))}
        _ = v38
    case "\n\t":
        v39 := GenS3{F0: 9223372036854775807, F1: v30}
        _ = v39
    case "":
        var v40 int32 = -928
        _ = v40
    }
    _ = len(v2)
    v41 := GenS2{F0: 246, F1: 1000}
    v4 -= 0
    fmt.Println(v4, v4, v4)
    v42 := GenS4{F0: 0, F1: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v7
}
