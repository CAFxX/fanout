package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0() int32 {
    var v5 uint = 0
    v6 := GenS1{F0: "foo bar", F1: 10}
    if true {
        v7 := GenS1{F0: (("a" + (("" + ("foo bar" + "a")) + "foo bar")) + "hello"), F1: 128}
        _ = v7
    }
    switch "\n\t" {
    case "\n\t":
        for fi8 := 0; fi8 < 18; fi8++ {
            switch fi8 {
            case 32768:
                v9 := GenS1{F0: "a", F1: 255}
                _ = v9
            default:
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := GenS1{F0: "hello", F1: 42}
                    v5 = 65535
                    fmt.Println(fi8)
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi8
        }
    }
    _ = v5
    _ = v6
    return 201
}
func genFn1(p0 int, p1 int, p2 int64) uint64 {
    p0 = 1
    if (false || (128 > 25)) {
        switch p0 {
        case 42:
            var v12 int32 = -1
            _ = v12
        }
    } else {
        var v13 float32 = 1.5
        switch p1 {
        case 32767:
            if false {
                var v14 string = ("hello" + "hello")
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(p1, fi15)
                    p2 = (p2 ^ (p2 ^ -218))
                    var v16 float32 = (1.5 / v13)
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                var v17 uint8 = 197
                if true {
                    p0 += p1
                    v18 := []int{}
                    var v19 uint = 526
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: ((("" + "foo bar") + "a") + "hello"), F1: (v17 << ((1) & 63))}
                _ = v17
                _ = v20
            }
        case -563:
            v21 := []int{4294967295, (p0 - p0), (-(p1)), p1}
            _ = v21
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    var v4 uint64 = 0
    var v22 int8 = -1
    var v23 int = 32767
    var v24 uint32 = 0
    v24 = v24
    v25 := GenS2{F0: (v22 * (v22 + (v22 - 1)))}
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: "a", F1: 1}
            var v29 uint16 = 2
            if (uint32(v4) == (v24 << ((254) & 63))) {
                v30 := GenS2{F0: v22}
                _ = v30
            } else {
                fmt.Println(v22)
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        if (1 <= -32768) {
            v32 := []int{}
            _ = v32
        } else {
            delete(v1, "a")
            v2 -= 694
        }
        switch "hello" {
        case "foo bar":
            switch v23 {
            case 1:
                if (float64((3.14159 / 1.0)) == ((0.0 / float64(((0.1 - 1e10) * (0.1 + 3.14159)))) - ((0.0 - float64(0.0)) + (-2.25 * 0.0)))) {
                    var v33 int = (fi31 * int(v4))
                    _ = v33
                } else {
                    if false {
                        v4 += genFn1((fi31 >> ((42) & 63)), -1, -1)
                        var v34 int16 = -1
                        _ = v34
                    } else {
                        var v35 int16 = 790
                        for fi36 := 0; fi36 < 14; fi36++ {
                            var v37 uint32 = (v24 & v24)
                            v24 += v24
                            v38 := GenS3{F0: "a", F1: genFn1((fi36 * v23), v23, 42), F2: 126}
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    }
                }
            case 503:
                var v39 uint16 = 65535
                _ = v39
            default:
                v2 += int32(v4)
            }
        case "a":
            var v40 float64 = 1e10
            _ = v40
        case "":
            v2 -= v2
        default:
            v23 += fi31
        }
        var v41 uint8 = 0
        _ = fi31
        _ = v41
    }
    var v42 bool = (false && (3.14159 == float32(1e10)))
    var v43 int16 = -743
    var v44 bool = (709 > 255)
    if (0 != 953) {
        var v45 float64 = (3.14159 / 0.1)
        _ = len(v3)
        _ = v45
    }
    var v46 uint8 = 1
    for fi47 := 0; fi47 < 14; fi47++ {
        v48 := []int{v23, int(v46)}
        var v49 int16 = v43
        v23 += (^(v23))
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 20; fi50++ {
        if v44 {
            for fi51 := 0; fi51 < 15; fi51++ {
                v52 := GenS4{F0: 1}
                if (float64(3.14159) == 1.5) {
                    v23 = (-(v23))
                    var v53 uint32 = (v24 / (((v24 | v24)) | 1))
                    v24 = (v53 >> ((463) & 63))
                    _ = v53
                }
                _ = fi51
                _ = v52
            }
            if (!(!false)) {
                v42 = v44
            } else {
                if (v43 != v43) {
                    if (float64(1.5) >= (1.5 / ((0.0 + -2.25) + 3.14159))) {
                        v42 = false
                    }
                } else {
                    v24 = uint32(v46)
                    v54 := []int{(v23 - 32768), -837, (v23 >> ((514) & 63)), v23}
                    _ = v54
                }
            }
        } else {
            v3[v24] = int(v22)
        }
        v2 = (v2 / ((genFn0()) | 1))
        _ = fi50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v46
}
