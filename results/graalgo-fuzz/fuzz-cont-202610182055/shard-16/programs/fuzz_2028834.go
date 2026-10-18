package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint16; F2 bool }

func main() {
    var v1 uint8 = 127
    var v2 int64 = 0
    var v3 uint32 = 840
    var v4 uint64 = 1
    var v5 float64 = 3.14159
    var v6 int64 = v2
    var v7 uint16 = 280
    v8 := GenS1{F0: 127, F1: (^(v3)), F2: (v3 + (^(v3)))}
    fmt.Println(v1, v2)
    fmt.Println(v5, v1)
    var v9 float32 = (1.5 * 0.0)
    var v10 bool = (false && ((v9 + v9) >= v9))
    v1 -= 128
    var v11 int16 = 213
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint = 255
        _ = fi12
        _ = v13
    }
    fmt.Println(v7, v11)
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 string = ("\n\t" + "foo bar")
            v1 = 2
            v17 := GenS1{F0: 1, F1: uint32(v4), F2: v3}
            _ = fi15
            _ = v16
            _ = v17
        }
        if v10 {
            if (!v10) {
                var v18 int8 = 100
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                if v10 {
                    if ((v3 * 649) >= v3) {
                        v3 -= 65535
                        v20 := []int{fi19, (fi14 & 265)}
                        var v21 uint = 18446744073709551615
                        _ = v20
                        _ = v21
                    }
                    if (!((fi19 - (-(fi19))) < (fi19 >> ((255) & 63)))) {
                        var v22 uint = 65535
                        if (255 < v7) {
                            v11 = v11
                            if v10 {
                                if v10 {
                                    v23 := []int{int(v11), fi14}
                                    var v24 uint8 = (v1 * (^(v1)))
                                    _ = v23
                                    _ = v24
                                } else {
                                    fmt.Println(fi19, v6, v7)
                                }
                            } else {
                                if v10 {
                                    v9 = v9
                                    v25 := GenS1{F0: 1, F1: 2147483647, F2: (v3 + 928)}
                                    _ = v25
                                } else {
                                    v11 = (v11 + v11)
                                    var v26 uint16 = v7
                                    v27 := []int{(fi19 ^ (-(fi19))), (^(fi14)), (fi19 & (-(fi19)))}
                                    switch "hello" {
                                    case "":
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v28 uint64 = (^(v4))
                                            _ = v28
                                        }
                                    case "foo bar":
                                        v7 += v26
                                    case "a":
                                        var v29 float64 = v5
                                        _ = v29
                                    default:
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                    }
                                    _ = v26
                                    _ = v27
                                }
                                var v30 float32 = 1e10
                                var v31 int64 = v2
                                v32 := GenS1{F0: 127, F1: 1, F2: v3}
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        } else {
                            if (true || v10) {
                                switch (("hello" + "\n\t") + "hello") {
                                case "\n\t":
                                    var v33 int = int(v4)
                                    _ = v33
                                case "hello":
                                    var v34 int16 = -248
                                    _ = v34
                                }
                                fmt.Println(v2, v2)
                            } else {
                                var v35 float64 = v5
                                _ = v35
                            }
                        }
                        v11 = (v11 / ((1000) | 1))
                        v36 := []int{(fi14 * int(v3))}
                        _ = v22
                        _ = v36
                    }
                    if (v3 > v3) {
                        var v37 int16 = (v11 & 32767)
                        if (0.1 <= float64(v9)) {
                            var v38 string = "a"
                            v9 -= v9
                            _ = v38
                        }
                        v39 := GenS2{F0: (^(v6)), F1: v7, F2: v10}
                        _ = v37
                        _ = v39
                    } else {
                        v5 += v5
                        switch "\n\t" {
                        case "\n\t":
                            v6 -= v2
                        case "":
                            v10 = (true && v10)
                        }
                    }
                    v9 -= v9
                }
                _ = fi19
            }
            fmt.Println(v11, v10)
        }
        v40 := []int{(fi14 | 127), fi14}
        _ = fi14
        _ = v40
    }
    v2 -= int64(v4)
    v41 := GenS2{F0: v2, F1: (v7 % ((255) | 1)), F2: (((v1 * v1) == uint8(v7)) || (int8(v1) > 100))}
    v42 := []int{438}
    v7 -= (v7 / (((v7 * 314)) | 1))
    if false {
        v4 += (v4 << ((4294967295) & 63))
        for fi43 := 0; fi43 < 14; fi43++ {
            for fi44 := 0; fi44 < 12; fi44++ {
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
                var v45 uint = 65535
                _ = fi44
                _ = v45
            }
            var v46 float32 = 3.14159
            fmt.Println(v9, v7)
            _ = fi43
            _ = v46
        }
        v47 := GenS2{F0: (v2 << ((1) & 63)), F1: v7, F2: (v10 || true)}
        _ = v47
    } else {
        var v48 uint64 = v4
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
