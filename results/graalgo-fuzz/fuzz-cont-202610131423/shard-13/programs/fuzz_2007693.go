package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 int; F1 uint; F2 string }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int64; F1 bool }
func genFn0(p0 uint16, p1 int16, p2 uint64) int8 {
    v5 := []int{127}
    switch "hello" {
    case "a":
        _ = len(v5)
    default:
        v6 := GenS1{F0: 1}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1() int16 {
    switch "hello" {
    case "foo bar":
        v7 := GenS2{F0: ("\n\t" + "\n\t"), F1: 1}
        _ = v7
    case "hello":
        var v8 int32 = -98
        _ = v8
    }
    var v9 bool = false
    v10 := GenS1{F0: 0}
    _ = v10
    _ = v9
    return -32768
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    v1 += v1
    var v11 uint16 = 0
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                fmt.Println(v1, v11, v1)
                if (false || true) {
                    fmt.Println(fi13)
                    v14 := []int{}
                    if (42 >= 986) {
                        v11 -= v11
                        fmt.Println(v1, fi13)
                    } else {
                        v1 -= (v1 + (v1 & v1))
                        v15 := GenS1{F0: -931}
                        v2[(fi12 - -272)] = (57 <= 0)
                        v14 = append(v14, fi12)
                        var v16 int8 = genFn0(255, int16(fi13), uint64(v1))
                        _ = v15
                        _ = v16
                    }
                    _ = v14
                } else {
                    var v17 uint64 = 42
                    if (v11 != v11) {
                        switch "hello" {
                        case "hello":
                            v11 += (v11 >> ((255) & 63))
                        }
                    }
                    v18 := []int{fi12, fi12, fi13}
                    _ = v17
                    _ = v18
                }
                var v19 uint = 0
                var v20 uint16 = v11
                _ = v19
                _ = v20
            } else {
                v11 += v11
                if (!false) {
                    v21 := GenS3{F0: genFn1()}
                    v1 += (v1 - (v1 & 4294967295))
                    _ = v21
                } else {
                    var v22 string = "a"
                    _ = v22
                }
            }
            switch "hello" {
            case "foo bar":
                v4[0] = (42 < 1)
            case "a":
                var v23 int8 = 0
                _ = v23
            }
            v11 -= 255
            _ = fi13
        }
        switch "hello" {
        case "\n\t":
            if (0.1 >= 0.0) {
                for fi24 := 0; fi24 < 12; fi24++ {
                    v11 = (v11 ^ 932)
                    var v25 int32 = 421
                    _ = fi24
                    _ = v25
                }
            } else {
                var v26 float32 = (-2.25 / 1.0)
                v27 := GenS4{F0: 0, F1: 830, F2: 0}
                _ = v26
                _ = v27
            }
        case "":
            delete(v2, fi12)
        }
        v28 := GenS5{F0: (fi12 & fi12), F1: 0, F2: ("" + "")}
        _ = fi12
        _ = v28
    }
    switch "a" {
    case "foo bar":
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 string = "hello"
            _ = fi29
            _ = v30
        }
    case "":
        switch ("foo bar" + "") {
        case "\n\t":
            v31 := []int{9223372036854775807, 2147483647, -9223372036854775808, -1}
            _ = v31
        case "hello":
            v32 := GenS6{F0: 65535, F1: genFn1()}
            _ = v32
        }
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        switch "\n\t" {
        case "":
            var v33 bool = false
            _ = v33
        case "hello":
            v34 := GenS7{F0: 1, F1: v11}
            _ = v34
        }
    default:
        v35 := GenS8{F0: 4294967295, F1: 1}
        _ = v35
    }
    var v36 string = ""
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 int16 = -801
        v3[v1] = (v1 / ((v1) | 1))
        var v39 string = (("\n\t" + v36) + "\n\t")
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 uint = 896
    switch 32767 {
    case -2147483648:
        v41 := []int{4294967295, 4294967295}
        _ = v41
    case 0:
        fmt.Println(v1, v40, v36)
    }
    delete(v4, 39)
    var v42 bool = true
    v43 := GenS9{F0: 161, F1: (^(v11))}
    fmt.Println(v1)
    v40 -= v40
    if (v36 >= (("\n\t" + v36) + ("a" + (v36 + ("" + v36))))) {
        v42 = v42
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 int = (-(fi44))
            for fi46 := 0; fi46 < 8; fi46++ {
                if (fi44 < fi44) {
                    var v47 bool = ((!(0.1 > ((0.1 + ((1.5 * 3.14159) + 1e10)) + 1e10))) || v42)
                    v48 := GenS3{F0: 71}
                    _ = v47
                    _ = v48
                } else {
                    var v49 uint16 = 65535
                    v50 := GenS6{F0: v40, F1: 462}
                    _ = v49
                    _ = v50
                }
                var v51 uint8 = 1
                _ = fi46
                _ = v51
            }
            fmt.Println(v1, v45, v11)
            _ = fi44
            _ = v45
        }
    }
    v52 := GenS10{F0: 412, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v52
}
