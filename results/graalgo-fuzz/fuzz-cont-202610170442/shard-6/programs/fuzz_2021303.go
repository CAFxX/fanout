package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint16 }
func genFn0() bool {
    if (!true) {
        if true {
            var v5 int32 = 42
            switch "hello" {
            case "a":
                fmt.Println(v5)
            default:
                for fi6 := 0; fi6 < 2; fi6++ {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        var v8 int64 = 429
                        _ = fi7
                        _ = v8
                    }
                    var v9 bool = ((0 < 1) || ((false || ((562 <= 0) || (1000 != 695))) || true))
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 bool = v9
                        _ = fi10
                        _ = v11
                    }
                    _ = fi6
                    _ = v9
                }
            }
            v12 := GenS1{F0: 42, F1: v5, F2: ""}
            _ = v12
            _ = v5
        }
        var v13 bool = false
        var v14 float32 = (((1e10 - 3.14159) - -2.25) + float32(((1e10 / (1.5 * ((-2.25 * (0.0 / 0.1)) + 1.5))) / 1.0)))
        _ = v13
        _ = v14
    } else {
        var v15 uint16 = 255
        _ = v15
    }
    var v16 float32 = -2.25
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := []int{32767, fi17, fi17, fi17}
        _ = fi17
        _ = v18
    }
    if true {
        var v19 uint64 = 623
        v16 -= v16
        _ = v19
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := GenS2{F0: true, F1: 1}
                if false {
                    var v23 uint32 = 451
                    v23 = v23
                    _ = v23
                } else {
                    if (32767 > 42) {
                        v24 := []int{32768, (fi20 << ((4294967295) & 63)), (fi20 / ((2147483647) | 1)), (^(fi21))}
                        _ = v24
                    } else {
                        v16 -= v16
                        v25 := GenS3{F0: 0, F1: 0}
                        switch fi20 {
                        case 2147483647:
                            var v26 uint8 = 127
                            _ = v26
                        }
                        _ = v25
                    }
                    v27 := []int{fi20, -534}
                    _ = v27
                }
                v28 := []int{fi21, fi21, 42, 2147483647}
                if (true || (127 != 122)) {
                    fmt.Println(fi20)
                } else {
                    v29 := GenS4{F0: 751}
                    var v30 int32 = 0
                    _ = v29
                    _ = v30
                }
                _ = fi21
                _ = v22
                _ = v28
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := []int{-192}
                _ = fi31
                _ = v32
            }
            var v33 uint = 768
            _ = fi20
            _ = v33
        }
        if (false || false) {
            v34 := []int{32767, 555, 9223372036854775807, 712}
            if len(v34) > 0 {
                v34[0] = 127
            }
            _ = v34
        } else {
            v35 := []int{32767, 9223372036854775807}
            _ = v35
        }
        var v36 bool = ("hello" > "")
        _ = v36
    }
    v37 := GenS1{F0: 714, F1: 901, F2: ("hello" + ("\n\t" + "hello"))}
    _ = v16
    _ = v37
    return ((!true) || false)
}

func main() {
    var v1 bool = false
    var v2 bool = (1.5 > (0.0 * (1.5 + (float64(0.0) / (0.0 + (0.1 / 1.0))))))
    v3 := []uint32{}
    var v4 uint64 = 4294967295
    v38 := GenS2{F0: true, F1: -753}
    if true {
        switch "foo bar" {
        case "hello":
            switch 1 {
            case 1:
                fmt.Println(v4)
            case 775:
                var v39 int64 = 0
                _ = v39
            }
        case "a":
            var v40 uint16 = 0
            _ = v40
        case "":
            v4 -= 1
        }
    } else {
        var v41 int64 = 9223372036854775807
        _ = v41
    }
    fmt.Println(v2, v4)
    var v42 int = 661
    var v43 uint32 = 42
    v44 := GenS5{F0: v43, F1: 2, F2: (^(v43))}
    var v45 int8 = -27
    _ = len(v3)
    var v46 bool = (65535 >= 1)
    for fi47 := 0; fi47 < 3; fi47++ {
        if (533 >= -720) {
            v48 := GenS4{F0: 786}
            _ = v48
        } else {
            v1 = v1
        }
        var v49 uint16 = uint16(v43)
        _ = fi47
        _ = v49
    }
    v50 := GenS6{F0: 962}
    for fi51 := 0; fi51 < 3; fi51++ {
        if len(v3) > 0 {
            v3[0] = (^(v43))
        }
        _ = fi51
    }
    fmt.Println(v43)
    v2 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v50
}
