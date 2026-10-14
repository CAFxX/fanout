package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 int; F1 int32; F2 string }
type GenS9 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint64, p1 bool) int32 {
    fmt.Println(p1)
    var v6 uint = 0
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 uint8) string {
    if true {
        var v7 int64 = 593
        switch 128 {
        case 83:
            p0 = p0
        case -678:
            switch -868 {
            case -839:
                v8 := []int{-2147483648, 127, -9223372036854775808}
                _ = v8
            case 84:
                v9 := GenS1{F0: 14}
                _ = v9
            case -679:
                v10 := []int{}
                _ = v10
            default:
                v11 := []int{655, 128, 4294967295, -865}
                _ = v11
            }
        }
        v12 := GenS2{F0: 1}
        _ = v12
        _ = v7
    } else {
        var v13 int16 = -27
        for fi14 := 0; fi14 < 16; fi14++ {
            p0 -= 127
            var v15 int = (fi14 >> ((640) & 63))
            switch ("foo bar" + "foo bar") {
            case "a":
                p0 -= (^(p0))
            default:
                if (1 < 114) {
                    var v16 uint8 = (p0 / ((128) | 1))
                    switch "foo bar" {
                    case "hello":
                        v17 := GenS3{F0: "", F1: 238}
                        _ = v17
                    case "foo bar":
                        var v18 bool = ("foo bar" <= "hello")
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v20 := []int{}
                            _ = fi19
                            _ = v20
                        }
                    }
                    v21 := GenS4{F0: 756, F1: 1}
                    _ = v16
                    _ = v21
                } else {
                    v15 = fi14
                    var v22 uint32 = 42
                    switch "\n\t" {
                    case "hello":
                        if (true || (v13 >= (v13 * v13))) {
                            fmt.Println(v13, p0, v22)
                            v23 := GenS4{F0: 451, F1: (v22 - (^(v22)))}
                            fmt.Println(p0)
                            _ = v23
                        }
                    case "":
                        v15 = (v15 - v15)
                    case "a":
                        v24 := GenS3{F0: "hello", F1: v22}
                        _ = v24
                    }
                    _ = v22
                }
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v25 := GenS3{F0: ("\n\t" + "foo bar"), F1: 729}
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS5{F0: 1}
        p0 += 128
        p0 = p0
        _ = fi26
        _ = v27
    }
    v28 := []int{32767, 32768, 32767, 32768}
    if (3.14159 <= float32(float64((float32(((float64((1e10 - 3.14159)) / (1e10 - 0.1)) + 3.14159)) / (1e10 / 0.1))))) {
        var v29 int = -2147483648
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := GenS6{F0: p0}
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        v32 := GenS7{F0: (p0 & p0), F1: 1}
        fmt.Println(p0)
        _ = v32
    }
    _ = p0
    _ = v25
    _ = v28
    return "hello"
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + "a"), (("a" + ("\n\t" + "foo bar")) + ("\n\t" + ("" + ("foo bar" + "a"))))}
    var v2 string = "foo bar"
    var v3 uint64 = 42
    var v4 int64 = -136
    v5 := make(map[uint32]uint32)
    for fi33 := 0; fi33 < 8; fi33++ {
        switch (fi33 * 1) {
        case 1:
            for fi34 := 0; fi34 < 2; fi34++ {
                switch (fi33 + fi34) {
                case 1:
                    var v35 float64 = (1e10 + float64(-2.25))
                    _ = v35
                case 127:
                    var v36 uint8 = 42
                    _ = v36
                case -2147483648:
                    var v37 string = (v2 + (v2 + ("foo bar" + (v2 + ("" + v2)))))
                    _ = v37
                }
                v1 = append(v1, ((v2 + ("hello" + ((((v2 + (v2 + ((("\n\t" + "a") + v2) + ""))) + v2) + (((v2 + ((v2 + ("hello" + "hello")) + "")) + v2) + v2)) + v2))) + (v2 + v2)))
                v38 := GenS8{F0: (fi34 / ((fi33) | 1)), F1: 828, F2: (v2 + v2)}
                _ = fi34
                _ = v38
            }
        case 128:
            var v39 float64 = float64((1e10 * -2.25))
            _ = v39
        case 9223372036854775807:
            var v40 uint8 = 0
            _ = v40
        }
        _ = fi33
    }
    v2 = (v2 + "")
    if false {
        v3 = v3
        v41 := GenS4{F0: 2, F1: 2}
        _ = v41
    } else {
        v2 = "\n\t"
        v42 := GenS9{F0: 255, F1: 881}
        for fi43 := 0; fi43 < 7; fi43++ {
            v2 = v2
            var v44 int16 = 469
            if false {
                var v45 int8 = 127
                var v46 uint8 = 0
                _ = v45
                _ = v46
            } else {
                var v47 int16 = -32768
                _ = v47
            }
            _ = fi43
            _ = v44
        }
        _ = v42
    }
    var v48 int64 = (-(v4))
    v1 = append(v1, v2)
    if (true || false) {
        v2 = "foo bar"
        var v49 float32 = (((1.5 / (0.1 * float32((1.5 + 3.14159)))) / 1.0) - 0.1)
        _ = v49
    }
    _ = len(v5)
    fmt.Println(v3, v2)
    var v50 bool = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v50
}
