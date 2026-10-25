package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0() uint {
    v7 := GenS1{F0: "", F1: 299, F2: 18446744073709551615}
    var v8 uint = 0
    v9 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: (v8 % (((v8 & v8)) | 1))}
    if ((false && true) || false) {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch (fi10 >> ((v8) & 63)) {
            case 128:
                switch (fi10 | (fi10 + (^(fi10)))) {
                case 127:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v8 = (v8 & v8)
                        if (714 != fi10) {
                            v8 -= v8
                            var v12 int64 = int64(v8)
                            _ = v12
                        }
                        _ = fi11
                    }
                case 635:
                    v13 := GenS1{F0: "\n\t", F1: 42, F2: 18446744073709551615}
                    _ = v13
                case 558:
                    if true {
                        v14 := GenS1{F0: "\n\t", F1: 1, F2: (v8 * (v8 & (^(v8))))}
                        v15 := []int{9223372036854775807, (fi10 >> ((255) & 63)), fi10}
                        _ = v14
                        _ = v15
                    } else {
                        v16 := []int{(fi10 + (fi10 | fi10)), 8}
                        _ = v16
                    }
                default:
                    v8 = (v8 ^ 42)
                }
            case 276:
                v8 -= v8
            case 2147483647:
                if true {
                    fmt.Println(fi10)
                    var v17 bool = (true || true)
                    switch fi10 {
                    case -679:
                        v18 := []int{(fi10 | fi10)}
                        _ = v18
                    case 456:
                        fmt.Println(v17)
                    case -9223372036854775808:
                        if v17 {
                            fmt.Println(fi10, v17, fi10)
                        } else {
                            v8 = (v8 | (v8 & v8))
                        }
                    }
                    _ = v17
                } else {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 uint64 = 18446744073709551615
                        _ = fi19
                        _ = v20
                    }
                    var v21 uint16 = 32767
                    _ = v21
                }
            default:
                v8 -= 65535
            }
            var v22 float64 = 1e10
            _ = fi10
            _ = v22
        }
    } else {
        v8 = 0
        v8 = (v8 * 2)
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 uint = 255
            var v25 uint8 = 108
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 2147483647
    v3 := []int{-9223372036854775808, 620}
    v4 := []int{128, 42, 2147483647}
    v5 := make(map[string]string)
    var v6 int16 = 414
    for fi26 := 0; fi26 < 16; fi26++ {
        switch ("foo bar" + (("hello" + "a") + "foo bar")) {
        case "\n\t":
            v27 := GenS1{F0: "", F1: -740, F2: 65535}
            _ = v27
        default:
            for fi28 := 0; fi28 < 4; fi28++ {
                v5[("\n\t" + "hello")] = "\n\t"
                delete(v5, ("\n\t" + "foo bar"))
                v29 := v5["hello"]
                v30 := GenS1{F0: ("foo bar" + "a"), F1: 1, F2: 384}
                v6 = (v6 >> ((4294967295) & 63))
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi26
    }
    v31 := GenS1{F0: ("foo bar" + "a"), F1: 0, F2: genFn0()}
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v1 = ((0.0 + 1.5) * -2.25)
    case "a":
        fmt.Println(v2, v1)
    case "\n\t":
        _ = len(v3)
    }
    if len(v3) > 0 {
        v3[0] = 32767
    }
    v6 = (-(v6))
    var v32 uint8 = 246
    if false {
        var v33 uint32 = 4294967295
        _ = v33
    } else {
        v2 += (-(v2))
    }
    v34 := GenS1{F0: "\n\t", F1: 149, F2: 18446744073709551615}
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS2{F0: 2147483648, F1: false, F2: 65535}
        v32 = v32
        if (234 == 42) {
            var v37 uint64 = 42
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    var v38 int = 32767
    v39 := GenS3{F0: v38, F1: 2, F2: 65535}
    if true {
        var v40 bool = false
        v40 = true
        if v40 {
            if (((v2 * (v2 / (((v2 ^ v2)) | 1))) > v2) && v40) {
                var v41 bool = ((!(-2.25 == (float64(0.0) + float64(float32(v1))))) || ("foo bar" > "\n\t"))
                if (v40 && (!(!false))) {
                    if (255 < 330) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v40 = v40
                        v40 = v40
                    }
                } else {
                    v32 += v32
                }
                if len(v3) > 0 {
                    v3[0] = v38
                }
                _ = v41
            } else {
                v42 := GenS2{F0: 2147483648, F1: v40, F2: 181}
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 string = "\n\t"
                    v45 := GenS3{F0: fi43, F1: 0, F2: 1}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = v42
            }
        } else {
            var v46 int64 = -28
            v40 = ((v2 / ((2147483647) | 1)) < v2)
            _ = v46
        }
        _ = v40
    }
    var v47 uint8 = (v32 >> ((2) & 63))
    var v48 uint8 = (^(v47))
    v1 = v1
    var v49 float32 = (1.5 + float32(((v1 * (v1 * (0.0 - v1))) + 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
