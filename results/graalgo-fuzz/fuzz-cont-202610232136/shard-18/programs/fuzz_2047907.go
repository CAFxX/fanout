package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 int32; F2 int }
type GenS5 struct { F0 uint16; F1 uint; F2 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 int32) uint8 {
    var v7 bool = (p0 > p0)
    fmt.Println(p0, v7)
    v8 := []int{416, -9223372036854775808, -9223372036854775808}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint8 = 255
    var v3 string = ("" + "a")
    var v4 uint16 = 32767
    var v5 float64 = ((-2.25 / 1.0) + 3.14159)
    var v6 uint = 65535
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 = (v2 - (v2 ^ 42))
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{133}
            if false {
                var v12 int16 = 32767
                var v13 uint16 = v4
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: 0}
                switch v3 {
                case "foo bar":
                    v1 = v1
                case "":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                default:
                    if (v1 && (0 >= 2147483647)) {
                        if v1 {
                            v11 = append(v11, (fi10 - int(v4)))
                        }
                    }
                }
                if len(v11) > 0 {
                    v11[0] = fi9
                }
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        v15 := GenS2{F0: "a"}
        _ = fi9
        _ = v15
    }
    var v16 uint = (v6 & (v6 / ((uint(v4)) | 1)))
    var v17 int32 = 2147483647
    var v18 uint8 = 63
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            if v1 {
                v21 := []int{32768, (fi19 & (fi19 | (fi19 << (((v6 | v6)) & 63)))), (^(fi19)), 84}
                fmt.Println(v16, fi20)
                if ((v1 && (!((!true) || ((v5 / 1.5) != float64(float32(float64((float32((v5 - v5)) + ((0.0 * 3.14159) * (float32(v5) - ((0.1 - 3.14159) / 3.14159))))))))))) || ((v18 & genFn0(-187)) < v2)) {
                    if (!v1) {
                        v22 := GenS1{F0: 0}
                        var v23 float64 = (v5 * v5)
                        fmt.Println(v4)
                        _ = v22
                        _ = v23
                    } else {
                        v5 = 0.0
                    }
                } else {
                    switch (((v3 + (((("a" + (v3 + v3)) + "foo bar") + v3) + v3)) + "\n\t") + v3) {
                    case "\n\t":
                        if (!(!v1)) {
                            var v24 float32 = 1.5
                            _ = v24
                        }
                    case "hello":
                        v6 -= 841
                    default:
                        v25 := GenS2{F0: v3}
                        _ = v25
                    }
                    switch (("a" + v3) + (v3 + (v3 + v3))) {
                    case "foo bar":
                        switch (fi20 << ((uint(v2)) & 63)) {
                        case -881:
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            var v26 uint16 = (v4 + (v4 << ((v16) & 63)))
                            var v27 uint = 1
                            _ = v26
                            _ = v27
                        default:
                            fmt.Println(v18)
                        }
                    default:
                        v6 = v16
                    }
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                _ = v21
            } else {
                if v1 {
                    v28 := GenS3{F0: (v4 * v4), F1: (v4 | 65535), F2: -1}
                    var v29 int32 = (v17 >> ((v6) & 63))
                    fmt.Println(v4, v2)
                    v30 := GenS4{F0: v3, F1: (-(v17)), F2: fi20}
                    switch fi19 {
                    case 817:
                        var v31 int = fi20
                        _ = v31
                    default:
                        var v32 int16 = -632
                        _ = v32
                    }
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            v33 := GenS3{F0: 726, F1: v4, F2: -1}
            _ = fi20
            _ = v33
        }
        if (v1 && v1) {
            v4 = v4
        } else {
            switch "foo bar" {
            case "a":
                for fi34 := 0; fi34 < 13; fi34++ {
                    v16 = 1
                    v1 = v1
                    _ = fi34
                }
            case "\n\t":
                v16 = (^(v6))
            }
        }
        _ = fi19
    }
    var v35 float32 = float32(v5)
    if true {
        for fi36 := 0; fi36 < 12; fi36++ {
            if (v1 && (v17 > -1)) {
                v37 := []int{fi36, (fi36 << (((v16 >> ((v16) & 63))) & 63))}
                _ = v37
            } else {
                v38 := []int{}
                fmt.Println(v6, v16)
                _ = v38
            }
            _ = fi36
        }
    } else {
        var v39 int = 659
        _ = v39
    }
    var v40 int8 = 42
    switch ((v3 + "\n\t") + v3) {
    case "\n\t":
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 uint16 = 255
            v43 := []int{(fi41 / ((fi41) | 1)), (fi41 << (((v16 & v16)) & 63)), fi41}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    v44 := GenS5{F0: 65535, F1: (v6 - 42), F2: v40}
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := GenS6{F0: 42}
        _ = fi45
        _ = v46
    }
    v47 := GenS7{F0: (true || false), F1: -2147483648, F2: 1}
    v48 := []int{-893, -319, 4294967295, 0}
    _ = len(v48)
    if len(v48) > 0 {
        v48[0] = 4294967295
    }
    fmt.Println(v3)
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
