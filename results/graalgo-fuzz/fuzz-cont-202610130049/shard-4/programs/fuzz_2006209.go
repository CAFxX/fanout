package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
type GenS4 struct { F0 int8 }
func genFn0() int8 {
    switch 32767 {
    case 876:
        var v6 int16 = 0
        _ = v6
    case 32768:
        var v7 uint32 = 972
        _ = v7
    case 2147483647:
        var v8 uint = 0
        _ = v8
    default:
        var v9 int16 = 177
        _ = v9
    }
    if (!false) {
        v10 := []int{9223372036854775807, 32768}
        _ = v10
    }
    var v11 string = "hello"
    v11 = v11
    v11 = v11
    _ = v11
    return 0
}
func genFn1(p0 string, p1 bool) int16 {
    if (0 != 0) {
        p1 = (0.1 != ((0.0 + (3.14159 * (1e10 + float32(1.5)))) * -2.25))
    }
    if true {
        for fi12 := 0; fi12 < 16; fi12++ {
            if p1 {
                var v13 uint64 = 1
                var v14 uint32 = 1
                var v15 float32 = float32(1.5)
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 uint8 = 19
                var v17 int = (fi12 % ((0) | 1))
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := []int{0, (fi18 % (((fi12 - (v17 & (fi18 % (((-(fi12))) | 1))))) | 1))}
                    v19 = append(v19, fi18)
                    var v20 int64 = -450
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                _ = v16
                _ = v17
            }
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 int16 = 42
                p1 = false
                _ = fi21
                _ = v22
            }
            _ = fi12
        }
        var v23 string = "hello"
        _ = v23
    } else {
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(p1, p1)
                var v26 int16 = 142
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        v27 := []int{0, 320, -2147483648}
        switch -446 {
        case 128:
            v28 := GenS1{F0: 1, F1: 2}
            _ = v28
        case 127:
            var v29 uint = 2
            _ = v29
        case 1:
            v30 := GenS1{F0: 36, F1: 42}
            _ = v30
        default:
            var v31 int = 660
            _ = v31
        }
        _ = v27
    }
    p1 = (p1 || true)
    if (35 != 128) {
        p1 = (p1 && ((true && true) || false))
        var v32 int64 = 42
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 bool = p1
            v35 := GenS2{F0: 0, F1: 164, F2: 42}
            v36 := []int{fi33}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        _ = v32
    }
    _ = p0
    _ = p1
    return -786
}

func main() {
    var v1 int16 = -482
    var v2 float64 = 3.14159
    v3 := make(map[uint32]string)
    v4 := []int{296, 0}
    var v5 string = "\n\t"
    var v37 string = "\n\t"
    if false {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v4 = append(v4, 4294967295)
            switch 127 {
            case 42:
                if (true && (false || (v5 != v37))) {
                    switch -902 {
                    case -2147483648:
                        switch "a" {
                        case "":
                            v1 += (v1 >> ((4294967295) & 63))
                        case "\n\t":
                            v2 -= 1e10
                        default:
                            var v38 uint8 = 255
                            _ = v38
                        }
                    case 202:
                        v5 = (("foo bar" + (v37 + ((("\n\t" + "a") + v5) + "a"))) + ("" + "foo bar"))
                    default:
                        if false {
                            v39 := GenS3{F0: 4294967295, F1: 127, F2: 803}
                            v4 = append(v4, 768)
                            fmt.Println(v37)
                            _ = v39
                        } else {
                            var v40 uint8 = 42
                            _ = v40
                        }
                    }
                    _ = len(v3)
                } else {
                    if ((false || (42 <= 418)) || (2 == 255)) {
                        for fi41 := 0; fi41 < 9; fi41++ {
                            var v42 uint8 = 127
                            _ = fi41
                            _ = v42
                        }
                        fmt.Println(v5, v37, v5)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = len(v3)
                }
            case 32767:
                v4 = append(v4, 42)
            default:
                v43 := GenS2{F0: 0, F1: -2147483648, F2: 127}
                _ = v43
            }
        } else {
            if ("" == ("a" + v5)) {
                if len(v4) > 0 {
                    v4[0] = 127
                }
                switch 0 {
                case 2147483647:
                    var v44 uint8 = 255
                    _ = v44
                case -1:
                    for fi45 := 0; fi45 < 13; fi45++ {
                        delete(v3, 42)
                        fmt.Println(v1, v1)
                        _ = fi45
                    }
                }
                v37 = ""
            }
            if len(v4) > 0 {
                v4[0] = 127
            }
            v46 := GenS2{F0: -495, F1: 0, F2: 9223372036854775807}
            _ = v46
        }
    } else {
        v3[778] = v37
        v1 -= v1
        for fi47 := 0; fi47 < 9; fi47++ {
            for fi48 := 0; fi48 < 10; fi48++ {
                v5 = (v37 + v5)
                _ = fi48
            }
            var v49 int32 = 0
            _ = fi47
            _ = v49
        }
    }
    v50 := GenS3{F0: 0, F1: 127, F2: 0}
    v51 := GenS1{F0: 0, F1: 2147483647}
    var v52 uint8 = 42
    if (255 <= 0) {
        var v53 bool = true
        var v54 int16 = 606
        v55 := GenS4{F0: 42}
        _ = v53
        _ = v54
        _ = v55
    }
    var v56 float32 = (float32(((0.0 / v2) / 1.5)) * 1e10)
    var v57 uint = 2
    v58 := GenS4{F0: 0}
    fmt.Println(v1, v37, v5)
    for fi59 := 0; fi59 < 18; fi59++ {
        for fi60 := 0; fi60 < 11; fi60++ {
            v2 = 3.14159
            var v61 float64 = v2
            _ = fi60
            _ = v61
        }
        _ = fi59
    }
    v62 := GenS3{F0: 35, F1: 1, F2: 9223372036854775807}
    var v63 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v56
    _ = v57
    _ = v58
    _ = v62
    _ = v63
}
