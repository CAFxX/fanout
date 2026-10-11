package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint32, p2 string) uint64 {
    v7 := GenS1{F0: p0, F1: 4294967295}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{fi8, (^(fi8)), fi8}
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int32 = 42
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(fi12)
        var v13 uint32 = p1
        _ = fi12
        _ = v13
    }
    fmt.Println(p0, p0)
    p2 = ("hello" + (p2 + "foo bar"))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return uint64(p1)
}
func genFn1(p0 uint32, p1 uint64, p2 uint) int64 {
    if true {
        v14 := GenS2{F0: ("hello" + "")}
        var v15 int = 2147483647
        _ = v14
        _ = v15
    }
    v16 := []int{-9223372036854775808}
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 0
}

func main() {
    var v1 float32 = (float32(0.0) - ((0.0 / 1.5) * float32(-2.25)))
    var v2 string = (((("hello" + "\n\t") + "hello") + "\n\t") + "hello")
    var v3 uint16 = 951
    var v4 uint16 = v3
    var v5 int16 = 653
    var v6 int64 = -820
    v17 := GenS3{F0: -2147483648, F1: (255 < 127)}
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            switch (fi18 & (-(fi19))) {
            case -973:
                v6 = (v6 | (v6 & v6))
            }
            var v20 int = (^(fi18))
            _ = fi19
            _ = v20
        }
        v21 := GenS2{F0: "\n\t"}
        _ = fi18
        _ = v21
    }
    if false {
        if true {
            v22 := GenS1{F0: 4294967295, F1: 2147483647}
            if false {
                for fi23 := 0; fi23 < 16; fi23++ {
                    switch "hello" {
                    case "hello":
                        v24 := []int{(fi23 + 0), (^(fi23)), 9223372036854775807}
                        _ = v24
                    case "a":
                        v4 -= 0
                    case "":
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v4 = v3
                            _ = fi25
                        }
                    default:
                        v3 = v3
                    }
                    v26 := []int{465, 32767, 32767}
                    var v27 uint8 = 0
                    _ = fi23
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 8; fi28++ {
                    v29 := GenS3{F0: 0, F1: false}
                    _ = fi28
                    _ = v29
                }
            }
            if false {
                for fi30 := 0; fi30 < 8; fi30++ {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        if true {
                            v1 = float32(3.14159)
                            var v32 float64 = 0.0
                            _ = v32
                        }
                        switch -950 {
                        case -963:
                            fmt.Println(fi30, v3)
                        case 230:
                            v33 := GenS1{F0: 42, F1: -795}
                            _ = v33
                        case 4294967295:
                            v34 := []int{280}
                            v4 = v3
                            _ = v34
                        default:
                            var v35 bool = false
                            _ = v35
                        }
                        _ = fi31
                    }
                    v36 := GenS4{F0: 0}
                    _ = fi30
                    _ = v36
                }
                var v37 float32 = v1
                _ = v37
            }
            _ = v22
        } else {
            v6 = (v6 | -9223372036854775808)
            v4 += v4
        }
    }
    v6 -= genFn1(2147483647, 0, 446)
    v38 := GenS1{F0: genFn0(690, 4294967295, (v2 + ("foo bar" + (v2 + v2)))), F1: 128}
    var v39 int16 = (v5 ^ 42)
    v39 = v5
    if (((1e10 - v1) >= 1e10) || (!(42 < 1))) {
        v3 = 0
    } else {
        fmt.Println(v3, v5, v3)
    }
    var v40 int16 = v39
    fmt.Println(v5)
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 uint16 = (v4 - (^(v4)))
        if (1 <= 22) {
            var v43 uint32 = 2
            var v44 bool = false
            if (v44 || (219 != 127)) {
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 int8 = -1
                    _ = fi45
                    _ = v46
                }
                v3 = 2
                v47 := GenS2{F0: ""}
                _ = v47
            }
            _ = v43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
