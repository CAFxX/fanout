package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    v7 := GenS1{F0: 204, F1: "", F2: 0}
    v8 := GenS1{F0: 1, F1: "", F2: 42}
    switch "foo bar" {
    case "":
        var v9 uint16 = 2
        _ = v9
    case "foo bar":
        var v10 uint = 852
        _ = v10
    default:
        if (!true) {
            v11 := []int{}
            switch 42 {
            case 345:
                var v12 int32 = -289
                _ = v12
            }
            _ = v11
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 float64 = 0.0
                _ = fi13
                _ = v14
            }
            fmt.Println(p1, p0)
        }
    }
    v15 := []int{450, 0, 460, 701}
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = 371
    v2 := make(map[string]bool)
    v3 := []string{"hello", ("\n\t" + "a"), "hello"}
    var v4 uint8 = 211
    var v5 int32 = 0
    v6 := make(map[string]string)
    fmt.Println(v4, v5, v4)
    var v16 string = "hello"
    if (0 <= -432) {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint = 2
            for fi19 := 0; fi19 < 17; fi19++ {
                _ = len(v3)
                v3 = append(v3, "\n\t")
                if (577 > 32767) {
                    if ((fi17 << ((4294967295) & 63)) >= (fi19 + 0)) {
                        var v20 uint16 = 738
                        var v21 string = v16
                        var v22 uint32 = 2147483648
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        _ = len(v6)
                        switch "" {
                        case "foo bar":
                            v23 := GenS1{F0: (v4 % (((v4 ^ 128)) | 1)), F1: (v16 + "\n\t"), F2: genFn0(319, 486)}
                            _ = v23
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "":
                            v6[(v16 + (v16 + ((v16 + ("hello" + "")) + (v16 + "a"))))] = (v16 + "hello")
                        default:
                            v24 := GenS2{F0: v5, F1: 0, F2: 26}
                            _ = v24
                        }
                    }
                    var v25 uint32 = 424
                    _ = v25
                } else {
                    v26 := GenS1{F0: 96, F1: ("hello" + "a"), F2: (v4 + v4)}
                    _ = v26
                }
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        var v27 uint32 = 77
        v28 := GenS1{F0: v4, F1: ("foo bar" + "a"), F2: (v4 & (v4 & (v4 % ((2) | 1))))}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int16 = 1
        _ = fi29
        _ = v30
    }
    v31 := v2[v16]
    v32 := GenS3{F0: 2, F1: (v16 + "foo bar"), F2: v16}
    for fi33 := 0; fi33 < 10; fi33++ {
        if ((uint32(v4) == 65535) && v31) {
            v3 = append(v3, v16)
        } else {
            v34 := GenS4{F0: 0, F1: 18446744073709551615, F2: 18446744073709551615}
            _ = v34
        }
        v16 = (v16 + ((v16 + "foo bar") + (v16 + v16)))
        _ = fi33
    }
    var v35 uint = 42
    v35 = v35
    var v36 int32 = v5
    var v37 int16 = int16(v1)
    v38 := GenS1{F0: v4, F1: "hello", F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
