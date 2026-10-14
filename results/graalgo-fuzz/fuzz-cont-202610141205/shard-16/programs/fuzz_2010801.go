package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint, p1 bool, p2 uint64) uint64 {
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 0, F1: int8(p2), F2: -1}
        p2 += p2
        switch (fi7 - 32768) {
        case 127:
            v9 := []int{fi7}
            _ = v9
        case 2147483647:
            if ((p0 ^ 255) != (p0 + 493)) {
                p0 = 65535
            } else {
                if p1 {
                    fmt.Println(p1, p0)
                }
                switch fi7 {
                case 42:
                    var v10 string = "a"
                    _ = v10
                case -2147483648:
                    fmt.Println(fi7)
                default:
                    v11 := []int{}
                    _ = v11
                }
            }
        case -635:
            var v12 uint8 = 168
            _ = v12
        default:
            if (42 > -9223372036854775808) {
                p0 -= p0
                var v13 int64 = -544
                _ = v13
            }
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 % (((p2 + 2)) | 1))
}
func genFn1(p0 string, p1 uint8) uint64 {
    var v14 int8 = -115
    var v15 float32 = 0.0
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := []uint32{267, 1, 848, 2}
    v2 := make(map[int]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    var v5 uint32 = 65535
    if true {
        v3 = (v3 | (-(v3)))
    }
    var v16 uint32 = 0
    _ = len(v4)
    v17 := v2[4294967295]
    if (255 >= 112) {
        if true {
            switch "\n\t" {
            case "foo bar":
                for fi18 := 0; fi18 < 2; fi18++ {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v5 = v17
                        v17 = v16
                        v16 = (v17 - v17)
                        _ = fi19
                    }
                    var v20 float64 = 1.5
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v3 -= -9223372036854775808
                        var v22 uint = 2
                        var v23 int32 = 1
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = fi18
                    _ = v20
                }
            case "":
                for fi24 := 0; fi24 < 7; fi24++ {
                    v3 += (v3 - (v3 >> ((255) & 63)))
                    if (!true) {
                        var v25 float64 = float64(0.0)
                        _ = v25
                    }
                    _ = fi24
                }
            case "hello":
                v26 := GenS1{F0: 42, F1: 39, F2: v3}
                _ = v26
            }
            if true {
                if false {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v3 -= (v3 + v3)
                        if (-32768 > 42) {
                            v5 -= 0
                            v28 := GenS2{F0: 846}
                            _ = v28
                        }
                        if (false && false) {
                            var v29 float64 = 1.5
                            var v30 int64 = -9223372036854775808
                            v16 = v16
                            _ = v29
                            _ = v30
                        }
                        _ = fi27
                    }
                } else {
                    v3 -= v3
                    var v31 int64 = v3
                    _ = v31
                }
                v16 = 2147483647
            } else {
                v3 = 528
                v17 = (^(v16))
            }
        } else {
            var v32 float64 = -2.25
            var v33 float64 = 1e10
            v17 = (^(v17))
            _ = v32
            _ = v33
        }
    } else {
        var v34 float32 = 1e10
        _ = v34
    }
    var v35 uint8 = 2
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint = 2
        for fi38 := 0; fi38 < 2; fi38++ {
            switch fi36 {
            case -916:
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi38, fi36)
                }
            case 2147483647:
                var v39 int = fi36
                _ = v39
            }
            if (!false) {
                switch "a" {
                case "":
                    v16 = v5
                case "foo bar":
                    if false {
                        v3 = 42
                        var v40 int16 = 1
                        _ = v40
                    } else {
                        v5 = (v5 + v16)
                        v3 += (-(v3))
                        v17 = (v16 >> ((v37) & 63))
                    }
                case "hello":
                    v2[567] = v16
                default:
                    var v41 bool = ((65535 > (v5 ^ (v5 | v16))) && (!(1000 <= -999)))
                    _ = v41
                }
            }
            v42 := GenS2{F0: 42}
            _ = fi38
            _ = v42
        }
        _ = fi36
        _ = v37
    }
    v5 = v17
    switch "\n\t" {
    case "a":
        var v43 uint16 = 255
        _ = v43
    case "":
        v44 := GenS3{F0: (^(v35))}
        _ = v44
    case "foo bar":
        var v45 string = (("\n\t" + "\n\t") + "a")
        _ = v45
    }
    var v46 uint = 2
    if ((4294967295 <= 18446744073709551615) && false) {
        var v47 int = 32767
        var v48 uint = 0
        _ = v47
        _ = v48
    }
    var v49 int = 0
    v35 = 0
    for fi50 := 0; fi50 < 17; fi50++ {
        for fi51 := 0; fi51 < 14; fi51++ {
            var v52 int64 = v3
            v3 += (v52 | (v3 - v3))
            _ = fi51
            _ = v52
        }
        var v53 uint64 = genFn0((v46 | v46), (0 != -360), 883)
        for fi54 := 0; fi54 < 19; fi54++ {
            v16 -= 2147483647
            _ = fi54
        }
        _ = fi50
        _ = v53
    }
    v55 := GenS4{F0: 899, F1: ""}
    var v56 bool = false
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v55
    _ = v56
}
