package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0() uint32 {
    v7 := GenS1{F0: ("hello" + ("hello" + "hello")), F1: 9223372036854775807}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            if true {
                var v10 int16 = 32767
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    v11 := []int{-971}
    v11 = append(v11, 128)
    _ = v11
    _ = v7
    return 1
}
func genFn1() uint {
    var v12 uint32 = 65535
    switch -1 {
    case 32768:
        if (-128 >= 23) {
            v13 := []int{-1, 127, 2147483647}
            fmt.Println(v12)
            _ = v13
        } else {
            v14 := GenS1{F0: (("" + "") + "foo bar"), F1: -1}
            switch ((("foo bar" + "") + ("foo bar" + "")) + "hello") {
            case "":
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 bool = false
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 int16 = 0
                        var v19 uint16 = 42
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: 42, F1: -95, F2: 94}
                    _ = fi15
                    _ = v16
                    _ = v20
                }
            default:
                if (18446744073709551615 >= 1) {
                    var v21 uint8 = 42
                    fmt.Println(v12, v21)
                    var v22 bool = true
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v24 := []int{fi23}
                        if true {
                            var v25 bool = false
                            _ = v25
                        } else {
                            fmt.Println(fi23)
                            if (true && (9223372036854775807 == 0)) {
                                if true {
                                    var v26 int64 = -429
                                    var v27 int16 = 42
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(v12, v12)
                                }
                                v12 += v12
                            }
                            var v28 bool = (255 != 4294967295)
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                    }
                    if (("" <= ("hello" + ("hello" + "\n\t"))) && (0.1 >= (0.1 * 1e10))) {
                        v29 := GenS3{F0: 100}
                        v30 := GenS2{F0: 42, F1: 45, F2: 42}
                        fmt.Println(v12)
                        _ = v29
                        _ = v30
                    }
                }
            }
            v31 := []int{32767, 157, -872}
            _ = v14
            _ = v31
        }
    case 4294967295:
        v32 := GenS3{F0: -1}
        _ = v32
    default:
        var v33 bool = (-1 < 4294967295)
        _ = v33
    }
    v34 := []int{}
    if true {
        v12 -= genFn0()
        fmt.Println(v12)
        v12 = (v12 - 0)
    } else {
        fmt.Println(v12)
        switch "hello" {
        case "":
            var v35 float32 = 1e10
            _ = v35
        case "a":
            switch "\n\t" {
            case "hello":
                fmt.Println(v12)
            }
        }
        var v36 int16 = -689
        _ = v36
    }
    _ = v12
    _ = v34
    return 2
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint32 = (v1 % (((v1 & 584)) | 1))
    v3 := []uint32{v2, (v2 >> ((0) & 63)), (^(v2))}
    v4 := make(map[uint32]int)
    v5 := []string{"foo bar"}
    v6 := make(map[int]string)
    var v37 float32 = float32(float64((float32((((3.14159 - float64(1.5)) * 3.14159) + 1e10)) - 1.5)))
    if len(v3) > 0 {
        v3[0] = (v2 >> ((65535) & 63))
    }
    var v38 int16 = 1000
    v38 = v38
    v1 = (v1 & genFn0())
    v1 = v1
    switch 737 {
    case 9223372036854775807:
        var v39 int32 = 128
        _ = v39
    case -890:
        switch ("a" + (("a" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t"))) {
        case "hello":
            v38 = -296
        case "a":
            v37 += v37
        default:
            switch -226 {
            case -290:
                if true {
                    v40 := GenS1{F0: "hello", F1: 0}
                    v1 -= (v2 % ((v2) | 1))
                    _ = v40
                }
            case 2147483647:
                v41 := GenS4{F0: 335, F1: 118, F2: v1}
                _ = v41
            case -332:
                v1 = (^(v1))
            }
        }
    case 2147483647:
        v42 := GenS3{F0: -128}
        _ = v42
    default:
        v37 -= v37
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        fmt.Println(v38, v38, v37)
        v2 = v1
        _ = fi43
    }
    var v44 uint = 2
    fmt.Println(v37, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
