package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 bool }
type GenS5 struct { F0 uint }

func main() {
    var v1 uint = 42
    var v2 uint64 = 18446744073709551615
    var v3 int8 = 52
    v4 := make(map[string]int)
    var v5 uint8 = 0
    var v6 int16 = -32768
    var v7 int16 = v6
    if false {
        v8 := GenS1{F0: v5, F1: (v7 | v7), F2: -1}
        v9 := []int{1, 299}
        v10 := v4["foo bar"]
        _ = v10
        _ = v8
        _ = v9
    }
    if (32767 == -2147483648) {
        var v11 uint8 = (v5 << ((65535) & 63))
        fmt.Println(v3)
        v12 := GenS1{F0: 127, F1: v7, F2: 159}
        _ = v11
        _ = v12
    } else {
        var v13 uint64 = (v2 * 765)
        fmt.Println(v7, v2, v6)
        _ = v13
    }
    switch 0 {
    case 286:
        var v14 uint8 = (v5 & v5)
        _ = v14
    case -784:
        v1 += v1
    case 128:
        var v15 bool = (297 <= 32767)
        _ = v15
    }
    v3 += 1
    var v16 uint16 = 42
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS2{F0: 0}
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        if true {
            if false {
                var v20 uint32 = 723
                var v21 int8 = v3
                _ = v20
                _ = v21
            }
        }
        _ = fi19
    }
    if true {
        var v22 bool = false
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := []int{(fi23 + int(v2)), int(v16), -275}
            if ((1.5 * float32(1e10)) <= float32(0.0)) {
                v25 := GenS3{F0: (^(v7)), F1: (v3 / ((-111) | 1))}
                _ = v25
            } else {
                if false {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v27 := GenS2{F0: -1}
                        switch fi23 {
                        case -1:
                            var v28 int64 = 50
                            switch ("\n\t" + ("a" + "a")) {
                            case "hello":
                                if true {
                                    var v29 uint32 = 65535
                                    fmt.Println(v2)
                                    var v30 int32 = int32(v2)
                                    _ = v29
                                    _ = v30
                                }
                            case "\n\t":
                                if false {
                                    if false {
                                        v28 += int64(v16)
                                        var v31 int = int(v2)
                                        v31 = v31
                                        _ = len(v4)
                                        var v32 string = ("\n\t" + ((((("foo bar" + (("\n\t" + (("" + ("foo bar" + ("hello" + ("" + "")))) + ("\n\t" + "\n\t"))) + ("" + ("hello" + "\n\t")))) + "hello") + "hello") + "foo bar") + ("\n\t" + ("" + "\n\t"))))
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            }
                            v33 := GenS1{F0: 255, F1: (^(v6)), F2: -703}
                            _ = v28
                            _ = v33
                        case 128:
                            v4["a"] = fi26
                        default:
                            v34 := GenS4{F0: -1, F1: 42, F2: (v1 != (v1 - v1))}
                            _ = len(v24)
                            v6 += v6
                            _ = v34
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                v35 := GenS2{F0: 458}
                _ = v35
            }
            v36 := GenS5{F0: 65535}
            _ = fi23
            _ = v24
            _ = v36
        }
        _ = v22
    }
    v16 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
