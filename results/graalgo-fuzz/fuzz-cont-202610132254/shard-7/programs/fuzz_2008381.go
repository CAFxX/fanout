package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16, p1 int8) uint32 {
    p0 += p0
    p0 += (p0 << ((1) & 63))
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6}
        switch (fi6 % ((fi6) | 1)) {
        case -2147483648:
            var v8 int64 = -1
            _ = v8
        case 32768:
            var v9 int = fi6
            _ = v9
        case 411:
            switch "foo bar" {
            case "\n\t":
                if len(v7) > 0 {
                    v7[0] = 1
                }
            case "hello":
                for fi10 := 0; fi10 < 9; fi10++ {
                    _ = len(v7)
                    _ = fi10
                }
            case "":
                var v11 uint8 = 127
                _ = v11
            default:
                if len(v7) > 0 {
                    v7[0] = fi6
                }
            }
        default:
            switch fi6 {
            case 32768:
                v12 := GenS1{F0: 355, F1: false}
                _ = v12
            case 127:
                v13 := GenS1{F0: 1, F1: false}
                _ = v13
            default:
                p0 -= (p0 * p0)
            }
        }
        _ = fi6
        _ = v7
    }
    var v14 uint = 305
    _ = p0
    _ = p1
    _ = v14
    return 194
}
func genFn1() uint32 {
    var v15 uint64 = 1
    v16 := GenS1{F0: (v15 - (v15 | (v15 >> ((18446744073709551615) & 63)))), F1: true}
    _ = v15
    _ = v16
    return 2
}

func main() {
    var v1 float32 = (0.1 / 3.14159)
    var v2 bool = (!("\n\t" > (("\n\t" + "\n\t") + ("a" + "a"))))
    var v3 uint8 = 0
    var v4 uint32 = 0
    v5 := make(map[uint32]bool)
    v17 := []int{2147483647, 717, -783}
    fmt.Println(v4)
    v4 -= v4
    var v18 int32 = 1
    if len(v17) > 0 {
        v17[0] = 32767
    }
    var v19 int = 42
    switch "hello" {
    case "\n\t":
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int16 = 348
            switch int(v21) {
            case -320:
                v2 = (("hello" + (("hello" + "hello") + "\n\t")) != "\n\t")
            case 42:
                switch ("hello" + (("hello" + "foo bar") + "foo bar")) {
                case "\n\t":
                    switch 32768 {
                    case 0:
                        if true {
                            switch (("foo bar" + "") + (("foo bar" + "hello") + ("" + ("hello" + (("foo bar" + "foo bar") + "a"))))) {
                            case "":
                                var v22 float32 = (v1 * 0.1)
                                _ = v22
                            case "a":
                                var v23 float64 = 1e10
                                _ = v23
                            }
                            var v24 int16 = (v21 ^ v21)
                            var v25 bool = v2
                            _ = v24
                            _ = v25
                        } else {
                            v5[(v4 << ((65535) & 63))] = (!v2)
                            var v26 uint16 = uint16(v4)
                            v17 = append(v17, (fi20 - -721))
                            _ = v26
                        }
                    }
                case "a":
                    if v2 {
                        for fi27 := 0; fi27 < 2; fi27++ {
                            var v28 uint = 461
                            v19 = (-(fi20))
                            _ = fi27
                            _ = v28
                        }
                    }
                }
            default:
                for fi29 := 0; fi29 < 8; fi29++ {
                    v3 += 1
                    if len(v17) > 0 {
                        v17[0] = int(v21)
                    }
                    _ = fi29
                }
            }
            if (v2 || v2) {
                if v2 {
                    switch 1 {
                    case -1:
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                } else {
                    v1 -= 0.0
                }
                v30 := GenS1{F0: 1, F1: ((32767 > 814) && (3.14159 >= ((-2.25 * ((0.1 + 1.5) * ((3.14159 + (0.1 * float64((v1 * (v1 * 1.5))))) + (1e10 / float64(v1))))) + -2.25)))}
                _ = v30
            } else {
                delete(v5, 4294967295)
                v2 = (1 < -88)
                var v31 uint16 = 2
                _ = v31
            }
            _ = fi20
            _ = v21
        }
    }
    var v32 int32 = (-(v18))
    if v2 {
        v33 := GenS1{F0: 18446744073709551615, F1: v2}
        var v34 uint16 = 0
        _ = v33
        _ = v34
    }
    switch (v19 & v19) {
    case 4294967295:
        var v35 float32 = v1
        _ = v35
    case 128:
        var v36 int = 32767
        _ = v36
    case -1:
        for fi37 := 0; fi37 < 11; fi37++ {
            for fi38 := 0; fi38 < 13; fi38++ {
                if (!(v2 && (38 <= 255))) {
                    if (100 >= 63) {
                        _ = len(v17)
                        v32 -= (v32 & -2147483648)
                        var v39 uint = 42
                        if len(v17) > 0 {
                            v17[0] = (fi38 | (fi38 & int(v39)))
                        }
                        var v40 uint8 = v3
                        _ = v39
                        _ = v40
                    } else {
                        var v41 uint32 = 463
                        _ = v41
                    }
                }
                var v42 uint32 = (v4 * v4)
                v43 := GenS2{F0: 65535}
                _ = fi38
                _ = v42
                _ = v43
            }
            fmt.Println(v1, v1)
            v44 := GenS2{F0: 706}
            _ = fi37
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        if ((fi45 % ((-803) | 1)) >= (fi45 / (((-(fi45))) | 1))) {
            switch (fi45 + -825) {
            case 2147483647:
                if v2 {
                    v46 := v5[(v4 * (v4 | (v4 << ((255) & 63))))]
                    _ = v46
                } else {
                    v3 -= 109
                }
            case 4294967295:
                var v47 uint64 = 18446744073709551615
                _ = v47
            }
            var v48 uint32 = 42
            var v49 uint = 381
            _ = v48
            _ = v49
        }
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 int16 = 817
            var v52 int32 = v18
            fmt.Println(v32, v3, v32)
            _ = len(v5)
            _ = fi50
            _ = v51
            _ = v52
        }
        var v53 int64 = -1
        _ = fi45
        _ = v53
    }
    v32 = (v18 | v18)
    var v54 int8 = -128
    _ = len(v5)
    fmt.Println(v1)
    v54 = (v54 + (^(v54)))
    v55 := GenS1{F0: 2, F1: false}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
