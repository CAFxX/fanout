package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int, p1 int16, p2 uint8) int8 {
    var v7 string = "a"
    v8 := GenS1{F0: -2147483648}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint = 2
        var v11 uint64 = 0
        v10 += (^(v10))
        _ = fi9
        _ = v10
        _ = v11
    }
    switch (("hello" + v7) + v7) {
    case "a":
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{fi12, fi12, (p0 >> ((0) & 63)), p0}
            _ = fi12
            _ = v13
        }
    case "\n\t":
        for fi14 := 0; fi14 < 10; fi14++ {
            if false {
                v15 := GenS1{F0: 1}
                for fi16 := 0; fi16 < 4; fi16++ {
                    fmt.Println(p1, fi14)
                    _ = fi16
                }
                var v17 uint32 = 65535
                _ = v15
                _ = v17
            } else {
                p2 = (p2 + 181)
                var v18 uint64 = 18446744073709551615
                _ = v18
            }
            _ = fi14
        }
    }
    var v19 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 0
    var v3 int8 = 100
    var v4 uint64 = v1
    v5 := []int{}
    v6 := make(map[uint32]uint32)
    switch 127 {
    case 301:
        switch 810 {
        case 2147483647:
            switch "\n\t" {
            case "hello":
                switch -9223372036854775808 {
                case 127:
                    if ((true && false) || (1e10 <= (0.1 * 3.14159))) {
                        v1 -= v1
                        for fi20 := 0; fi20 < 10; fi20++ {
                            var v21 bool = true
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        var v22 uint8 = 128
                        var v23 int = 127
                        v23 = int(v1)
                        _ = v22
                        _ = v23
                    }
                }
            case "":
                v24 := GenS1{F0: 708}
                _ = v24
            case "foo bar":
                v1 += (v1 % ((v4) | 1))
            default:
                for fi25 := 0; fi25 < 8; fi25++ {
                    if true {
                        v1 += v1
                        for fi26 := 0; fi26 < 20; fi26++ {
                            fmt.Println(v1, v4)
                            v6[2] = 0
                            var v27 bool = true
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        v28 := v6[1]
                        v5 = append(v5, int(v3))
                        _ = v28
                    }
                    _ = fi25
                }
            }
        case 32767:
            var v29 float64 = (float64(0.0) + float64(3.14159))
            _ = v29
        default:
            for fi30 := 0; fi30 < 12; fi30++ {
                if (fi30 < fi30) {
                    v31 := GenS2{F0: uint16(v1), F1: "\n\t"}
                    _ = v31
                }
                var v32 uint64 = v1
                _ = fi30
                _ = v32
            }
        }
    }
    v33 := GenS3{F0: -1}
    if (!true) {
        for fi34 := 0; fi34 < 9; fi34++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v3)
            _ = fi34
        }
        var v35 int = int(v4)
        v36 := GenS4{F0: 32767, F1: 2147483647}
        _ = v35
        _ = v36
    }
    v4 = (v4 % (((v4 + 42)) | 1))
    v37 := GenS2{F0: 1, F1: (("" + "\n\t") + (("foo bar" + ("foo bar" + ("" + "\n\t"))) + ""))}
    delete(v6, 0)
    var v38 bool = true
    _ = len(v5)
    var v39 uint32 = 65535
    v40 := GenS1{F0: 0}
    if true {
        v41 := GenS2{F0: 1, F1: ("\n\t" + "foo bar")}
        _ = v41
    } else {
        v42 := GenS3{F0: -9223372036854775808}
        var v43 int32 = 0
        _ = v42
        _ = v43
    }
    v44 := GenS5{F0: v2}
    v4 = (v4 | (v1 + (v1 % ((v4) | 1))))
    v45 := GenS3{F0: 565}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
