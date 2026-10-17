package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 0
    var v3 uint16 = 1
    v4 := make(map[uint32]bool)
    var v5 bool = (true && (true && (v2 != v2)))
    v3 += 1
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: -656, F1: (-(v2)), F2: -431}
    var v8 uint8 = 2
    var v9 uint32 = uint32(v6)
    v9 = (^(v9))
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (v5 && v5) {
                if v5 {
                    v2 -= 974
                }
                v2 -= v2
                v12 := []int{fi10, int(v2)}
                switch "\n\t" {
                case "":
                    _ = len(v12)
                }
                if ((v5 && (255 > v6)) && ((true || true) && v5)) {
                    v1[(v9 * v9)] = ((1000 <= 445) || (float64(0.1) != 3.14159))
                }
                _ = v12
            }
            v13 := GenS2{F0: v6}
            _ = fi11
            _ = v13
        }
        switch 42 {
        case 472:
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS1{F0: 1000, F1: -819, F2: fi14}
                v3 -= (^(v3))
                _ = len(v1)
                var v16 uint64 = 4294967295
                v17 := v4[330]
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        case 42:
            var v18 uint = v6
            _ = v18
        case -1:
            v19 := []int{(fi10 - (fi10 / (((fi10 >> ((255) & 63))) | 1))), (fi10 >> ((v6) & 63)), (fi10 & int(v9))}
            _ = v19
        }
        if ((-61 > 127) || v5) {
            v8 = (v8 + v8)
            var v20 int64 = v2
            _ = v20
        }
        _ = fi10
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        switch (fi21 - (fi21 & 9223372036854775807)) {
        case 127:
            if (v5 && ("a" >= ((("" + "hello") + "hello") + ("a" + ("" + ""))))) {
                switch fi21 {
                case 9223372036854775807:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        if (-9223372036854775808 >= (v2 * 42)) {
                            v8 += (v8 + (v8 % ((v8) | 1)))
                        } else {
                            fmt.Println(v3, v2)
                            var v23 int64 = int64(fi22)
                            _ = v23
                        }
                        _ = fi22
                    }
                case -1:
                    var v24 uint8 = (v8 & 58)
                    _ = v24
                case 32767:
                    fmt.Println(v2)
                default:
                    v25 := GenS1{F0: -1, F1: (v2 >> (((v6 & v6)) & 63)), F2: 32767}
                    _ = v25
                }
                v9 -= 4294967295
            }
        case 0:
            v26 := GenS3{F0: 109, F1: 220}
            _ = v26
        }
        _ = fi21
    }
    if (-32768 >= 1) {
        if (127 != -1) {
            var v27 uint8 = (^(v8))
            var v28 string = "foo bar"
            _ = v27
            _ = v28
        } else {
            switch "foo bar" {
            case "foo bar":
                v8 += (v8 - uint8(v6))
            default:
                if (v3 == v3) {
                    fmt.Println(v8, v8, v8)
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v9 = v9
                        fmt.Println(v3, v9)
                        _ = fi29
                    }
                    fmt.Println(v2)
                } else {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        var v31 int64 = (v2 % (((v2 & v2)) | 1))
                        v4[v9] = ((100 != 25) || v5)
                        v32 := GenS4{F0: 65535, F1: "foo bar"}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            }
            var v33 uint64 = 42
            for fi34 := 0; fi34 < 11; fi34++ {
                v9 += 2
                delete(v1, uint32(v6))
                if (128 <= (^(v8))) {
                    var v35 int64 = 9223372036854775807
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 14; fi36++ {
                        var v37 string = "foo bar"
                        fmt.Println(v5, v5, v5)
                        fmt.Println(v5, v2)
                        _ = fi36
                        _ = v37
                    }
                }
                _ = fi34
            }
            _ = v33
        }
        for fi38 := 0; fi38 < 11; fi38++ {
            delete(v4, 65535)
            v5 = (!(v9 <= (v9 * 2147483648)))
            fmt.Println(v2)
            _ = fi38
        }
    }
    v8 = (^(v8))
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 string = ""
            var v42 int = 4294967295
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS1{F0: -1, F1: v2, F2: (fi39 + (-(fi39)))}
        v44 := []int{(fi39 - fi39)}
        _ = fi39
        _ = v43
        _ = v44
    }
    var v45 uint64 = uint64(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
