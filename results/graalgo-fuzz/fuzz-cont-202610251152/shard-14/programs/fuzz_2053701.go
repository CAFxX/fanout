package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int64 }
func genFn0() uint {
    switch "" {
    case "\n\t":
        for fi4 := 0; fi4 < 13; fi4++ {
            switch "\n\t" {
            case "\n\t":
                for fi5 := 0; fi5 < 20; fi5++ {
                    var v6 uint32 = 4294967295
                    v6 += 2147483648
                    switch (fi4 + fi5) {
                    case 1:
                        if (!(255 == 128)) {
                            var v7 uint32 = (v6 << ((65535) & 63))
                            _ = v7
                        }
                    case -604:
                        v6 -= v6
                    case 32768:
                        v8 := GenS1{F0: false}
                        _ = v8
                    default:
                        v9 := []int{(-(fi4)), fi5, (fi4 / ((fi5) | 1))}
                        _ = v9
                    }
                    _ = fi5
                    _ = v6
                }
            case "hello":
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 int = (fi10 - (-(fi10)))
                    if false {
                        v12 := GenS1{F0: false}
                        v11 = 627
                        _ = v12
                    }
                    v11 += (fi4 << ((0) & 63))
                    _ = fi10
                    _ = v11
                }
            default:
                v13 := GenS2{F0: 2}
                _ = v13
            }
            _ = fi4
        }
    case "foo bar":
        var v14 uint16 = 812
        _ = v14
    case "":
        v15 := []int{-475, 42, 127, 9223372036854775807}
        _ = v15
    }
    var v16 uint32 = 198
    v17 := GenS1{F0: false}
    _ = v16
    _ = v17
    return 255
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{0, 1, 1, 750}
    var v3 int32 = -1
    var v18 int = -1
    fmt.Println(v3, v1, v3)
    switch (v18 ^ 32767) {
    case 42:
        for fi19 := 0; fi19 < 1; fi19++ {
            v1 = 1
            _ = fi19
        }
    case 245:
        fmt.Println(v1, v18, v3)
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(v18, fi21)
                v22 := GenS3{F0: 42, F1: 42}
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v2 = append(v2, 968)
                    v2 = append(v2, 65535)
                    v3 += (v3 / ((v3) | 1))
                    v23 := GenS3{F0: 2, F1: 113}
                    switch 990 {
                    case -678:
                        v3 += v3
                    case 2147483647:
                        var v24 int32 = (v3 * (v3 - v3))
                        _ = v24
                    }
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            if false {
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 uint64 = 305
                    _ = fi25
                    _ = v26
                }
                fmt.Println(v1, v1)
                v3 = v3
            } else {
                if false {
                    var v27 int32 = v3
                    fmt.Println(v18)
                    v27 += v3
                    _ = v27
                } else {
                    v2 = append(v2, 1)
                    var v28 float64 = 1.5
                    _ = v28
                }
            }
            var v29 uint = 718
            _ = fi20
            _ = v29
        }
    }
    var v30 uint64 = uint64(v1)
    if (1e10 > (((-2.25 / 3.14159) / 0.1) / ((3.14159 - 1e10) * 3.14159))) {
        v31 := GenS3{F0: 633, F1: 1}
        var v32 int16 = 1000
        v32 = (v32 / ((-1) | 1))
        _ = v31
        _ = v32
    } else {
        v30 = 2
        var v33 uint32 = uint32(v1)
        _ = v33
    }
    v34 := GenS3{F0: 290, F1: -128}
    if ((v30 % (((v30 >> ((genFn0()) & 63))) | 1)) <= 0) {
        var v35 float32 = -2.25
        if true {
            v36 := GenS4{F0: v3, F1: -92}
            switch (v18 & v18) {
            case 909:
                fmt.Println(v18, v30)
            case 42:
                var v37 int32 = (v3 << ((0) & 63))
                _ = v37
            default:
                var v38 uint32 = 2147483648
                _ = v38
            }
            _ = v36
        }
        v39 := GenS5{F0: -1}
        _ = v35
        _ = v39
    }
    if false {
        v18 = (v18 | (v18 / ((v18) | 1)))
        v2 = append(v2, 2147483648)
        v18 -= 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v34
}
