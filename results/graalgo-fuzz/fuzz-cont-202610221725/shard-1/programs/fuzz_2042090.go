package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0() uint8 {
    switch "" {
    case "hello":
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
    case "\n\t":
        if ((true && false) || false) {
            v8 := GenS1{F0: 65535, F1: (false && (0.0 < 0.0)), F2: "a"}
            for fi9 := 0; fi9 < 7; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 int8 = -1
                    var v12 uint64 = 18446744073709551615
                    v12 = v12
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 int32 = 42
                _ = fi9
                _ = v13
            }
            _ = v8
        }
    default:
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint64 = 0
            _ = fi14
            _ = v15
        }
    }
    var v16 uint = 1
    var v17 float32 = 1e10
    if (true && (!true)) {
        if ((true || true) && (-646 > -239)) {
            switch 540 {
            case 155:
                v18 := GenS2{F0: 2147483648}
                _ = v18
            case 9223372036854775807:
                fmt.Println(v17, v16)
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                v17 -= 0.0
                _ = fi19
            }
            var v20 uint32 = 655
            _ = v20
        } else {
            v21 := []int{1, 42, 4294967295, 385}
            v17 -= v17
            var v22 float32 = float32(float64((1e10 / 3.14159)))
            _ = v21
            _ = v22
        }
        if (1 != uint64(v16)) {
            var v23 int8 = -1
            _ = v23
        }
    } else {
        var v24 uint = 4294967295
        v25 := []int{32768, 32767, 128}
        v26 := GenS3{F0: (v24 - 560), F1: 0, F2: 127}
        _ = v24
        _ = v25
        _ = v26
    }
    if (!true) {
        if (!true) {
            var v27 bool = (42 < 18446744073709551615)
            _ = v27
        } else {
            if true {
                switch 1 {
                case -9223372036854775808:
                    for fi28 := 0; fi28 < 16; fi28++ {
                        if false {
                            if true {
                                var v29 string = ("a" + ("" + "a"))
                                _ = v29
                            }
                            v16 += (v16 >> ((679) & 63))
                        }
                        v17 += ((0.1 + 0.1) * v17)
                        var v30 uint32 = 885
                        _ = fi28
                        _ = v30
                    }
                case 32768:
                    var v31 int = 4294967295
                    _ = v31
                default:
                    if false {
                        if (true || false) {
                            v17 = v17
                        } else {
                            v17 -= v17
                        }
                        for fi32 := 0; fi32 < 17; fi32++ {
                            switch "hello" {
                            case "":
                                var v33 uint = v16
                                _ = v33
                            case "a":
                                fmt.Println(fi32)
                            case "foo bar":
                                v34 := []int{fi32}
                                _ = v34
                            default:
                                fmt.Println(fi32, fi32)
                            }
                            var v35 int = (fi32 | (fi32 + fi32))
                            var v36 string = (("a" + "\n\t") + "a")
                            _ = fi32
                            _ = v35
                            _ = v36
                        }
                        v17 += float32(-2.25)
                    }
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    fmt.Println(v16, v16)
                    _ = fi37
                }
            } else {
                v17 -= (1e10 / 1.0)
            }
            v38 := GenS4{F0: 0}
            _ = v38
        }
        v17 = 0.0
        var v39 uint32 = 0
        _ = v39
    } else {
        if false {
            switch ("a" + "foo bar") {
            case "foo bar":
                var v40 uint16 = 42
                _ = v40
            case "hello":
                for fi41 := 0; fi41 < 16; fi41++ {
                    v16 = (v16 % (((v16 / ((v16) | 1))) | 1))
                    _ = fi41
                }
            default:
                v42 := GenS4{F0: 940}
                _ = v42
            }
        } else {
            v43 := []int{32767, -160}
            var v44 bool = (!((0 > v16) && (!(-596 < 1))))
            _ = v43
            _ = v44
        }
        if ((true && false) && (741 == -1)) {
            var v45 bool = (false || true)
            switch ("\n\t" + "\n\t") {
            case "hello":
                var v46 uint16 = 1
                _ = v46
            case "\n\t":
                if true {
                    fmt.Println(v45, v17)
                    v45 = (("\n\t" != "a") && v45)
                    v47 := GenS1{F0: (^(v16)), F1: true, F2: "hello"}
                    _ = v47
                }
            }
            v48 := GenS3{F0: v16, F1: 128, F2: 36}
            _ = v45
            _ = v48
        } else {
            fmt.Println(v16, v17)
        }
        var v49 uint = v16
        _ = v49
    }
    _ = v16
    _ = v17
    return 127
}
func genFn1(p0 bool, p1 uint8) uint16 {
    v50 := []int{42, 42}
    var v51 int8 = -128
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v50
    _ = v51
    return 14
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]int)
    var v3 uint = 377
    var v4 int = -1
    var v5 uint64 = 60
    v6 := make(map[int]int)
    switch "a" {
    case "a":
        if ((v4 ^ (v4 + v4)) >= v4) {
            var v52 uint8 = 0
            var v53 int64 = 326
            _ = v52
            _ = v53
        }
    case "hello":
        v5 += v5
    default:
        v3 += (v3 << ((uint(v1)) & 63))
    }
    var v54 float64 = 1e10
    var v55 uint = 4294967295
    if (true && (42 == 42)) {
        v56 := v6[v4]
        fmt.Println(v5, v3, v56)
        _ = v56
    } else {
        var v57 string = "foo bar"
        _ = v57
    }
    v4 = (v4 / ((int(v55)) | 1))
    var v58 float64 = ((float64((1.5 * 3.14159)) * v54) / (-2.25 / (v54 / (0.1 / -2.25))))
    if true {
        var v59 int = (v4 + 128)
        v5 -= v5
        _ = v59
    } else {
        v4 = 32767
    }
    v60 := GenS2{F0: (v1 ^ v1)}
    v1 += v1
    v55 = (^(v55))
    if true {
        v4 -= -827
    } else {
        v61 := GenS5{F0: 1, F1: genFn0()}
        v55 += v3
        _ = v61
    }
    if (false && false) {
        switch v4 {
        case -9223372036854775808:
            if true {
                fmt.Println(v5)
            }
        case 4294967295:
            var v62 uint = v55
            _ = v62
        default:
            for fi63 := 0; fi63 < 4; fi63++ {
                v64 := []int{(fi63 - 127), fi63, -1, v4}
                _ = fi63
                _ = v64
            }
        }
    } else {
        fmt.Println(v55)
    }
    v65 := GenS1{F0: 242, F1: false, F2: "\n\t"}
    v58 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v58
    _ = v6
    _ = v60
    _ = v65
}
