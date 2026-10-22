package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int16 = -738
    v2 := []string{}
    var v3 float32 = (0.1 + (-2.25 + 1e10))
    var v4 uint8 = 0
    var v5 bool = true
    for fi6 := 0; fi6 < 5; fi6++ {
        switch fi6 {
        case 179:
            switch "" {
            case "\n\t":
                var v7 bool = true
                _ = v7
            case "a":
                v2 = append(v2, ("a" + ""))
            case "":
                var v8 uint = 847
                _ = v8
            }
        case 42:
            var v9 int8 = 0
            _ = v9
        }
        v2 = append(v2, "hello")
        _ = fi6
    }
    var v10 uint32 = 836
    v10 = v10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            v10 += (v10 >> ((18446744073709551615) & 63))
            _ = fi12
        }
        _ = fi11
    }
    v1 += v1
    if (!false) {
        v13 := GenS1{F0: (v10 / ((65535) | 1))}
        v14 := GenS2{F0: 1, F1: 1, F2: v4}
        switch 2147483647 {
        case 0:
            v15 := GenS2{F0: 2, F1: -299, F2: v4}
            _ = v15
        }
        _ = v13
        _ = v14
    }
    var v16 uint = 0
    var v17 float64 = -2.25
    if v5 {
        v18 := GenS1{F0: v10}
        v1 = 1000
        v4 += 176
        _ = v18
    } else {
        var v19 bool = v5
        fmt.Println(v16, v3)
        var v20 int = int(v4)
        _ = v19
        _ = v20
    }
    fmt.Println(v16, v4)
    for fi21 := 0; fi21 < 20; fi21++ {
        if v5 {
            fmt.Println(v3, v4)
            v2 = append(v2, ("a" + ("" + (("hello" + ("foo bar" + "hello")) + "hello"))))
        } else {
            if true {
                v4 -= v4
                fmt.Println(v10)
                for fi22 := 0; fi22 < 7; fi22++ {
                    v17 += float64(0.1)
                    if v5 {
                        v10 = v10
                        switch fi22 {
                        case 42:
                            if v5 {
                                v2 = append(v2, "a")
                            }
                            if (false || v5) {
                                _ = len(v2)
                                v17 -= float64(v3)
                            } else {
                                var v23 uint64 = 33
                                var v24 bool = true
                                if v5 {
                                    v17 = 0.0
                                    if (!((^(v23)) > v23)) {
                                        v25 := GenS1{F0: v10}
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(fi21, v4, fi21)
                                        default:
                                            v26 := GenS3{F0: 258, F1: v4}
                                            _ = v26
                                        }
                                        _ = v25
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v24 = false
                                    case "a":
                                        var v27 uint16 = 0
                                        _ = v27
                                    case "hello":
                                        v4 = 128
                                    }
                                    fmt.Println(v3, v5)
                                    v23 += 18446744073709551615
                                } else {
                                    var v28 uint64 = (v23 ^ (v23 ^ v23))
                                    v24 = v5
                                    if v5 {
                                        if v24 {
                                            if v24 {
                                                fmt.Println(v23, v5)
                                                fmt.Println(v4, v23)
                                                var v29 float64 = v17
                                                var v30 uint8 = v4
                                                _ = v29
                                                _ = v30
                                            } else {
                                                _ = len(v2)
                                                v2 = append(v2, ("hello" + ("\n\t" + ("\n\t" + (("hello" + "hello") + "\n\t")))))
                                                v10 += v10
                                            }
                                            _ = len(v2)
                                        }
                                    }
                                    var v31 int64 = 42
                                    _ = v28
                                    _ = v31
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        switch "\n\t" {
                        case "hello":
                            v3 = (v3 / (0.0 + 0.1))
                        case "\n\t":
                            v2 = append(v2, "\n\t")
                        case "a":
                            v2 = append(v2, "\n\t")
                        default:
                            fmt.Println(v3, v16)
                        }
                    }
                    var v32 int = fi21
                    _ = fi22
                    _ = v32
                }
            }
        }
        v33 := GenS3{F0: -544, F1: 127}
        v4 += 0
        _ = fi21
        _ = v33
    }
    v34 := GenS2{F0: 18446744073709551615, F1: 399, F2: (v4 ^ (v4 >> ((v16) & 63)))}
    var v35 uint8 = (v4 << ((921) & 63))
    switch 844 {
    case 4294967295:
        fmt.Println(v4, v17, v17)
    case 344:
        switch 127 {
        case 944:
            var v36 int32 = -449
            _ = v36
        }
    case 128:
        v2 = append(v2, "")
    default:
        if ((v10 % ((979) | 1)) >= v10) {
            v3 = 1.5
        }
    }
    var v37 bool = v5
    v16 += (v16 / (((v16 << (((v16 + uint(v35))) & 63))) | 1))
    v35 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
