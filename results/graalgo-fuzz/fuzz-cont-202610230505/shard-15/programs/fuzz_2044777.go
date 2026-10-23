package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 int16 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = 0
    var v2 float64 = -2.25
    var v3 string = "foo bar"
    var v4 uint8 = 0
    var v5 int = -2147483648
    var v6 bool = true
    var v7 int = v5
    v8 := []int{(v5 % ((4294967295) | 1)), 1, v5}
    var v9 float64 = v2
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            v8 = append(v8, 1)
            if (!v6) {
                if v6 {
                    v11 := GenS1{F0: true, F1: 733, F2: (v7 & 904)}
                    var v12 bool = v6
                    switch v3 {
                    case "":
                        fmt.Println(v4)
                    case "a":
                        if len(v8) > 0 {
                            v8[0] = 2147483647
                        }
                    default:
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 uint = 4294967295
                            if true {
                                var v15 uint16 = 0
                                v14 = (v14 % ((v14) | 1))
                                _ = v15
                            } else {
                                if len(v8) > 0 {
                                    v8[0] = (v7 << (((v14 * (v14 ^ (v14 * v14)))) & 63))
                                }
                            }
                            var v16 float64 = 0.0
                            _ = fi13
                            _ = v14
                            _ = v16
                        }
                    }
                    _ = v11
                    _ = v12
                }
            } else {
                var v17 string = v3
                for fi18 := 0; fi18 < 18; fi18++ {
                    fmt.Println(v5, fi18, fi10)
                    v19 := GenS1{F0: true, F1: 25, F2: v7}
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    fmt.Println(v5)
                    v21 := GenS2{F0: "foo bar", F1: 812}
                    _ = fi20
                    _ = v21
                }
                _ = v17
            }
            fmt.Println(fi10, v9)
        }
        switch (v7 / ((v5) | 1)) {
        case -154:
            if v6 {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
        case 32768:
            if (255 <= 3) {
                if v6 {
                    v5 = (v7 / ((666) | 1))
                    switch fi10 {
                    case 4294967295:
                        for fi22 := 0; fi22 < 2; fi22++ {
                            if len(v8) > 0 {
                                v8[0] = fi10
                            }
                            var v23 uint64 = 18446744073709551615
                            _ = fi22
                            _ = v23
                        }
                    }
                } else {
                    var v24 uint = 255
                    if v6 {
                        v25 := GenS2{F0: ("\n\t" + (v3 + "a")), F1: 2147483647}
                        fmt.Println(fi10)
                        _ = v25
                    } else {
                        var v26 uint8 = (v4 ^ v4)
                        _ = v26
                    }
                    v27 := GenS2{F0: v3, F1: 613}
                    _ = v24
                    _ = v27
                }
                v2 -= v9
            } else {
                v28 := GenS3{F0: 1, F1: -433, F2: 1000}
                if len(v8) > 0 {
                    v8[0] = (v5 % ((int(v4)) | 1))
                }
                v8 = append(v8, (-(fi10)))
                _ = v28
            }
        case -532:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = fi10
    }
    v29 := GenS4{F0: 1}
    var v30 int32 = (v1 % ((-1) | 1))
    v31 := GenS3{F0: 100, F1: 894, F2: 32767}
    var v32 float64 = -2.25
    var v33 float64 = (v2 / v9)
    v33 = v32
    var v34 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
