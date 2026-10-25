package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint }
type GenS6 struct { F0 bool; F1 int; F2 bool }
func genFn0(p0 int32, p1 int64, p2 int64) bool {
    p2 = -9223372036854775808
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = p0
        if true {
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint16 = 2
                var v10 uint8 = 117
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: "a"}
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint32 = 0
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        v14 := GenS1{F0: ("hello" + "hello")}
        _ = fi7
        _ = v14
    }
    switch "" {
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: (((((("hello" + "hello") + "foo bar") + "a") + ("foo bar" + "foo bar")) + (("foo bar" + "hello") + "hello")) + "hello")}
            _ = fi15
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (((255 <= 1) && false) || (925 != p2))
}

func main() {
    v1 := []uint32{0, 2147483648, 0}
    var v2 string = ("hello" + "\n\t")
    var v3 int32 = 2147483647
    var v4 uint64 = 42
    v5 := make(map[uint32]bool)
    var v6 string = v2
    var v17 float64 = (-2.25 + 0.1)
    v18 := GenS1{F0: v2}
    v19 := GenS2{F0: (v3 & (v3 >> ((0) & 63))), F1: v2}
    v20 := GenS1{F0: (("foo bar" + v2) + v6)}
    for fi21 := 0; fi21 < 11; fi21++ {
        if true {
            if (true || false) {
                v2 = "hello"
            } else {
                var v22 uint32 = 679
                v2 = ((v2 + "hello") + "hello")
                _ = v22
            }
            v23 := GenS3{F0: 0, F1: 0, F2: 2}
            _ = v23
        }
        _ = fi21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v24 := GenS1{F0: ("hello" + v2)}
    if (v17 >= (v17 * float64(float32(3.14159)))) {
        if true {
            _ = len(v1)
        }
        var v25 uint16 = 255
        var v26 uint = 1
        _ = v25
        _ = v26
    }
    v2 = ("\n\t" + (v6 + "\n\t"))
    fmt.Println(v6, v3, v4)
    if (!(42 <= -1)) {
        v3 = 489
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 bool = false
            if false {
                for fi29 := 0; fi29 < 3; fi29++ {
                    if (128 <= (fi29 >> ((604) & 63))) {
                        switch (fi27 | (^(fi29))) {
                        case 9223372036854775807:
                            switch ("foo bar" + v2) {
                            case "a":
                                if v28 {
                                    v1 = append(v1, 636)
                                }
                            case "\n\t":
                                delete(v5, 42)
                            }
                        case 0:
                            var v30 int16 = 42
                            _ = v30
                        case -264:
                            v31 := GenS3{F0: -1, F1: 23, F2: 0}
                            _ = v31
                        default:
                            var v32 bool = v28
                            _ = v32
                        }
                        if true {
                            v33 := GenS4{F0: 0, F1: 42, F2: -32768}
                            v3 = v3
                            fmt.Println(v6, v4)
                            _ = v33
                        }
                        switch fi29 {
                        case 414:
                            switch int(v4) {
                            case 32768:
                                if ((!(1 <= -128)) || (!v28)) {
                                    delete(v5, 65535)
                                    var v34 bool = v28
                                    _ = v34
                                }
                            }
                        case 456:
                            v17 = float64((1.5 / 1e10))
                        case 2147483647:
                            var v35 string = ((("hello" + (v6 + v2)) + v6) + "")
                            _ = v35
                        default:
                            fmt.Println(fi27)
                        }
                    } else {
                        var v36 uint32 = 1
                        switch (fi29 | fi27) {
                        case 234:
                            v37 := GenS4{F0: v36, F1: 127, F2: 366}
                            _ = v37
                        case 127:
                            v38 := GenS5{F0: 255, F1: 255, F2: 65535}
                            v1 = append(v1, (v36 + (v36 / (((v36 | v36)) | 1))))
                            _ = v38
                        case -804:
                            v1 = append(v1, (v36 | (v36 % ((v36) | 1))))
                        }
                        if len(v1) > 0 {
                            v1[0] = 152
                        }
                        _ = v36
                    }
                    _ = fi29
                }
            }
            var v39 int64 = 9223372036854775807
            _ = fi27
            _ = v28
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            for fi41 := 0; fi41 < 4; fi41++ {
                v6 = ""
                if (!false) {
                    var v42 int32 = (v3 >> ((0) & 63))
                    v43 := v5[2147483647]
                    v44 := GenS3{F0: 0, F1: 200, F2: 705}
                    _ = v42
                    _ = v43
                    _ = v44
                }
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                _ = fi41
            }
            fmt.Println(fi40)
            v17 -= 3.14159
            _ = fi40
        }
    }
    if true {
        v4 += v4
        var v45 string = v6
        _ = v45
    }
    var v46 bool = (808 == 0)
    if (!((true && v46) && (1 >= 2147483647))) {
        var v47 bool = (!true)
        v48 := GenS6{F0: (true || (127 != uint8(v4))), F1: 32768, F2: v46}
        if (uint8(v4) == 199) {
            delete(v5, 2147483648)
            _ = len(v1)
        }
        _ = v47
        _ = v48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
