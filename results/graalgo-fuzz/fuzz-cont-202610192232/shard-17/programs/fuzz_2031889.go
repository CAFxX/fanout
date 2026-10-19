package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0() int {
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 4; fi8++ {
        if (42 == 2) {
            v7 = v7
            var v9 int8 = -1
            for fi10 := 0; fi10 < 8; fi10++ {
                if true {
                    v7 = (^(v7))
                    v9 -= -128
                }
                v11 := GenS1{F0: 250}
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        var v12 bool = (-662 != 42)
        _ = fi8
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 uint16 = v7
        v15 := []int{(fi13 / ((fi13) | 1)), fi13, fi13}
        var v16 uint64 = 42
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if (!true) {
            v18 := []int{}
            _ = len(v18)
            switch fi17 {
            case -674:
                v19 := GenS1{F0: -32768}
                _ = v19
            case 4294967295:
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            }
            _ = v18
        }
        _ = fi17
    }
    _ = v7
    return 490
}
func genFn1(p0 int, p1 int16) uint32 {
    v20 := []int{p0, p0, p0, (^(p0))}
    switch "foo bar" {
    case "a":
        p0 += p0
    case "hello":
        if false {
            v21 := GenS1{F0: p1}
            if ((p1 >= p1) || true) {
                var v22 int16 = 0
                var v23 string = "a"
                p0 -= 42
                _ = v22
                _ = v23
            } else {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                fmt.Println(p1)
            }
            _ = v21
        }
    }
    p1 -= (p1 * 1)
    for fi24 := 0; fi24 < 11; fi24++ {
        if true {
            switch fi24 {
            case 1:
                v25 := GenS2{F0: 255, F1: 4294967295, F2: (true && (746 == 744))}
                _ = v25
            case 889:
                var v26 int64 = 9223372036854775807
                _ = v26
            case -57:
                var v27 uint16 = 0
                _ = v27
            default:
                if len(v20) > 0 {
                    v20[0] = fi24
                }
            }
            var v28 float64 = 0.0
            _ = v28
        }
        if false {
            var v29 float64 = 0.1
            for fi30 := 0; fi30 < 2; fi30++ {
                v29 += -2.25
                _ = fi30
            }
            p1 -= p1
            _ = v29
        }
        _ = fi24
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        if (!false) {
            p1 = 0
            if (42 > 4294967295) {
                for fi32 := 0; fi32 < 11; fi32++ {
                    if true {
                        var v33 int16 = p1
                        _ = v33
                    }
                    _ = fi32
                }
                if true {
                    p0 = 275
                    for fi34 := 0; fi34 < 18; fi34++ {
                        p1 = (-(p1))
                        _ = fi34
                    }
                    for fi35 := 0; fi35 < 6; fi35++ {
                        var v36 float32 = float32(1.5)
                        v37 := GenS1{F0: p1}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
                p1 = p1
            } else {
                if (42 >= -1) {
                    v38 := GenS3{F0: 677, F1: 32767, F2: 51}
                    for fi39 := 0; fi39 < 19; fi39++ {
                        var v40 uint8 = 1
                        p0 += int(p1)
                        _ = fi39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 1; fi41++ {
                        p1 = p1
                        p1 += p1
                        _ = fi41
                    }
                    _ = v38
                }
                v42 := GenS2{F0: 1, F1: 2147483648, F2: false}
                if (926 <= 65535) {
                    v43 := GenS1{F0: p1}
                    fmt.Println(p0, fi31)
                    _ = v43
                } else {
                    var v44 int32 = -495
                    _ = v44
                }
                _ = v42
            }
            var v45 uint16 = 701
            _ = v45
        }
        v46 := GenS1{F0: p1}
        p1 = (^(p1))
        _ = fi31
        _ = v46
    }
    _ = p0
    _ = p1
    _ = v20
    return 660
}

func main() {
    var v1 uint = 388
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    if false {
        var v47 uint8 = 255
        v48 := []int{4294967295, -443, 1, 127}
        fmt.Println(v47, v1)
        _ = v47
        _ = v48
    } else {
        v1 = v1
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 uint8 = 242
            _ = fi49
            _ = v50
        }
    }
    v1 = (v1 & (v1 | (v1 - (^(v1)))))
    fmt.Println(v1)
    var v51 uint = v1
    if (!true) {
        v1 += (v51 >> (((^(v51))) & 63))
        v1 = (v51 | 395)
    }
    v1 += 148
    if false {
        var v52 uint64 = 42
        v53 := GenS4{F0: 1, F1: true}
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 15; fi54++ {
        v51 += 758
        if (!(uint8(v51) <= 127)) {
            for fi55 := 0; fi55 < 10; fi55++ {
                v51 += v1
                v56 := []int{(fi54 << ((253) & 63)), (fi55 + (fi54 - fi54)), genFn0()}
                fmt.Println(fi55, v51)
                _ = fi55
                _ = v56
            }
            switch (fi54 * (fi54 >> ((2) & 63))) {
            case -324:
                if (!(0.1 > 0.0)) {
                    v51 = v1
                } else {
                    if false {
                        v51 += 315
                    }
                    for fi57 := 0; fi57 < 1; fi57++ {
                        v51 = v1
                        var v58 int32 = -2147483648
                        _ = fi57
                        _ = v58
                    }
                }
            default:
                v59 := GenS4{F0: 600, F1: (1 == 2)}
                _ = v59
            }
        } else {
            v51 += v1
            for fi60 := 0; fi60 < 8; fi60++ {
                var v61 string = ("foo bar" + ("foo bar" + ("a" + ("a" + ("\n\t" + ((("\n\t" + "foo bar") + "a") + "hello"))))))
                _ = fi60
                _ = v61
            }
            for fi62 := 0; fi62 < 3; fi62++ {
                delete(v6, 595)
                _ = fi62
            }
        }
        for fi63 := 0; fi63 < 4; fi63++ {
            v64 := v4[("" + ("foo bar" + "hello"))]
            v1 += (v1 - uint(v64))
            _ = len(v2)
            _ = fi63
            _ = v64
        }
        _ = fi54
    }
    v65 := GenS4{F0: -677, F1: true}
    for fi66 := 0; fi66 < 8; fi66++ {
        v1 = v1
        for fi67 := 0; fi67 < 2; fi67++ {
            switch ("a" + ("hello" + "a")) {
            case "a":
                v1 += (v1 & (v51 - 4294967295))
            case "\n\t":
                fmt.Println(fi66)
            }
            v1 = 2
            var v68 int = (^(fi66))
            _ = fi67
            _ = v68
        }
        fmt.Println(v51)
        _ = fi66
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v65
}
