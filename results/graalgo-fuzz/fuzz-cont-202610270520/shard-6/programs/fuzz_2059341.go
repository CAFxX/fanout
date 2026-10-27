package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 int; F2 int8 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 bool; F1 int; F2 uint64 }
func genFn0(p0 uint16, p1 string, p2 uint32) int16 {
    v5 := []int{-9223372036854775808, 128, 9223372036854775807, -1}
    var v6 int = 696
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = (^(p0))
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int32 = 42
        v10 := GenS1{F0: 102}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int8 = 100
    var v2 float64 = 1e10
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    if true {
        var v11 float64 = (0.1 + v2)
        var v12 bool = (-1 < 389)
        v1 = (^(v1))
        _ = v11
        _ = v12
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v13 := GenS2{F0: 306, F1: 130, F2: 4294967295}
        _ = v13
    case "\n\t":
        v14 := []int{}
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: 114, F2: 18446744073709551615}
    v2 = float64(1.5)
    if (!((!true) || true)) {
        for fi16 := 0; fi16 < 2; fi16++ {
            delete(v3, "")
            v17 := GenS3{F0: 2147483647, F1: genFn0(178, "\n\t", 544), F2: 42}
            _ = fi16
            _ = v17
        }
        if (0.0 > (((3.14159 + 3.14159) + (((0.1 / 1e10) / -2.25) * (1e10 + 1e10))) + -2.25)) {
            v18 := []int{-2147483648, 32767}
            _ = v18
        }
    }
    var v19 int32 = 1
    for fi20 := 0; fi20 < 8; fi20++ {
        if true {
            v2 = (float64((0.1 / 3.14159)) * 0.0)
            var v21 int = fi20
            var v22 int = (fi20 / ((fi20) | 1))
            _ = v21
            _ = v22
        } else {
            switch fi20 {
            case 127:
                var v23 int32 = (v19 - (v19 ^ (v19 << ((18446744073709551615) & 63))))
                _ = v23
            case 0:
                v1 = -1
            case 2147483647:
                v2 -= (v2 / v2)
            default:
                var v24 float64 = v2
                _ = v24
            }
            v25 := []int{fi20, fi20, fi20}
            switch "" {
            case "hello":
                v19 -= v19
            case "":
                v26 := GenS1{F0: 1}
                _ = v26
            case "a":
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 uint = 182
                    var v29 uint16 = 0
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            _ = v25
        }
        _ = fi20
    }
    if true {
        var v30 float32 = 0.1
        _ = v30
    } else {
        if ((v2 <= v2) || ("foo bar" < ("\n\t" + "hello"))) {
            if (255 > 42) {
                for fi31 := 0; fi31 < 13; fi31++ {
                    v32 := GenS4{F0: 28, F1: 42, F2: 976}
                    v33 := []int{}
                    if len(v33) > 0 {
                        v33[0] = fi31
                    }
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                var v34 int8 = v1
                if (231 >= 0) {
                    v3["a"] = 1
                } else {
                    if false {
                        v35 := GenS5{F0: 42}
                        v36 := GenS6{F0: 1000}
                        _ = v35
                        _ = v36
                    }
                    v34 = 44
                }
                _ = v34
            }
            var v37 string = ""
            _ = v37
        } else {
            v1 += v1
        }
    }
    switch 0 {
    case -77:
        var v38 int8 = -100
        _ = v38
    default:
        var v39 uint16 = genFn1(255)
        _ = v39
    }
    fmt.Println(v19, v2)
    if true {
        var v40 uint32 = 2147483648
        _ = v40
    } else {
        v41 := []int{-1, -1, -526}
        _ = v41
    }
    var v42 int = 9223372036854775807
    var v43 float32 = float32(float64(0.0))
    if false {
        if (true || (!(v19 <= 42))) {
            v43 = (1e10 + ((v43 / 1e10) - float32(v2)))
            if true {
                v42 = v42
                v44 := GenS7{F0: -1}
                var v45 uint16 = 2
                _ = v44
                _ = v45
            } else {
                var v46 int = (v42 + 4294967295)
                _ = v46
            }
            var v47 uint64 = 42
            _ = v47
        }
        v48 := GenS8{F0: -1, F1: (v42 | v42), F2: 100}
        _ = len(v4)
        _ = v48
    } else {
        switch 4294967295 {
        case 581:
            var v49 bool = false
            _ = v49
        case 2147483647:
            for fi50 := 0; fi50 < 5; fi50++ {
                fmt.Println(v43, v1)
                _ = fi50
            }
        default:
            var v51 float64 = v2
            _ = v51
        }
        switch (v42 / ((v42) | 1)) {
        case -255:
            v42 += v42
        case 481:
            v52 := GenS9{F0: v1, F1: (931 == int64(v42))}
            _ = v52
        case 0:
            _ = len(v4)
        default:
            v1 = -1
        }
        var v53 int32 = (v19 << ((375) & 63))
        _ = v53
    }
    v54 := GenS10{F0: (!(1 < 1000)), F1: (-(v42)), F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v54
}
