package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    var v3 int8 = -128
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 bool = true
    switch 128 {
    case 32768:
        v3 -= v3
    case 42:
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int16 = 0
            v9 := GenS1{F0: v3, F1: 128, F2: -1}
            _ = fi7
            _ = v8
            _ = v9
        }
    case 128:
        v10 := GenS2{F0: 1, F1: (-(v3))}
        _ = v10
    default:
        v11 := []int{-702}
        _ = v11
    }
    _ = len(v4)
    var v12 string = ""
    switch int(v3) {
    case 4294967295:
        v12 = v12
    case -9223372036854775808:
        v2 = 384
    default:
        if true {
            v2 = v2
        }
    }
    v13 := GenS3{F0: 1, F1: 2147483648}
    var v14 bool = v6
    var v15 float64 = float64(float32((float64((3.14159 / ((3.14159 + 1.5) * -2.25))) / 3.14159)))
    switch v12 {
    case "\n\t":
        v14 = v14
    case "":
        var v16 uint64 = (v2 * (v2 % ((v2) | 1)))
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := []int{fi17}
        _ = fi17
        _ = v18
    }
    if v14 {
        v19 := GenS1{F0: v3, F1: 172, F2: -2147483648}
        _ = v19
    }
    if false {
        v14 = v14
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS4{F0: -1, F1: (v3 + v3), F2: v2}
            var v22 bool = v6
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 uint16 = 552
                _ = fi23
                _ = v24
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 float64 = v15
            var v27 bool = v6
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint8 = 0
                if v14 {
                    switch (((((v12 + (("" + v12) + v12)) + "hello") + "a") + "a") + "hello") {
                    case "\n\t":
                        fmt.Println(v2)
                    case "a":
                        v30 := GenS3{F0: -2147483648, F1: 65535}
                        v31 := []int{fi25, fi28, (fi25 % (((fi25 % ((fi25) | 1))) | 1))}
                        if (602 == 65535) {
                            var v32 int8 = (v3 / (((v3 / (((v3 >> ((1) & 63))) | 1))) | 1))
                            v33 := GenS2{F0: 155, F1: -89}
                            _ = v32
                            _ = v33
                        } else {
                            switch 42 {
                            case 128:
                                var v34 int8 = (v3 / ((v3) | 1))
                                _ = v34
                            }
                            v35 := GenS2{F0: 2, F1: (v3 + v3)}
                            _ = v35
                        }
                        _ = v30
                        _ = v31
                    case "foo bar":
                        v36 := GenS1{F0: v3, F1: v29, F2: -2147483648}
                        _ = v36
                    }
                    var v37 uint8 = v29
                    v37 = 128
                    _ = v37
                } else {
                    var v38 int32 = 1
                    _ = v38
                }
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        if (2147483647 >= -1) {
            v3 += (v3 / ((-23) | 1))
            if v6 {
                var v39 int64 = 1
                if v6 {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        for fi41 := 0; fi41 < 19; fi41++ {
                            v42 := []int{fi40, (fi40 % (((fi40 ^ (fi41 & (fi40 >> ((2) & 63))))) | 1)), fi40, fi41}
                            _ = fi41
                            _ = v42
                        }
                        v12 = ("\n\t" + "a")
                        _ = fi40
                    }
                    var v43 int64 = (^(v39))
                    for fi44 := 0; fi44 < 15; fi44++ {
                        var v45 uint8 = 152
                        _ = fi44
                        _ = v45
                    }
                    _ = v43
                }
                _ = v39
            }
        } else {
            for fi46 := 0; fi46 < 7; fi46++ {
                v47 := v4[(v12 + v12)]
                _ = fi46
                _ = v47
            }
        }
    }
    fmt.Println(v12, v14, v15)
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
