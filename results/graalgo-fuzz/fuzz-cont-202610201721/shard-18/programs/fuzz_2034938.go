package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint32 }
func genFn0(p0 uint8, p1 bool) uint8 {
    p1 = (1000 >= 1)
    if ((float32(1e10) * (-2.25 * 3.14159)) == float32(((-2.25 / 1.5) - (1e10 / -2.25)))) {
        var v7 string = ""
        switch v7 {
        case "":
            var v8 string = v7
            _ = v8
        }
        if ((!(9223372036854775807 <= 424)) && false) {
            var v9 string = v7
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := []int{(^(fi10)), (fi10 / ((fi10) | 1)), -335}
                v11 = append(v11, fi10)
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        _ = v7
    } else {
        v12 := GenS1{F0: 128, F1: 617, F2: 121}
        p1 = p1
        _ = v12
    }
    if (((1.5 + 1e10) - ((0.1 + ((0.1 / ((3.14159 - 1e10) - 3.14159)) * ((1e10 - 1.5) - (0.1 * (1e10 * 1e10))))) / -2.25)) < 0.0) {
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                p0 = 255
                p1 = (42 == -54)
                var v15 int32 = int32(p0)
                v15 -= v15
                _ = fi14
                _ = v15
            }
            switch ("hello" + (("" + ("foo bar" + ("a" + "a"))) + "foo bar")) {
            case "a":
                switch (fi13 + fi13) {
                case 9223372036854775807:
                    v16 := []int{}
                    _ = v16
                }
            case "":
                for fi17 := 0; fi17 < 13; fi17++ {
                    v18 := []int{fi13, fi13, fi13, 128}
                    v19 := GenS1{F0: (fi17 + (fi17 + (fi13 * fi13))), F1: 32767, F2: 100}
                    if false {
                        p1 = false
                        v20 := GenS1{F0: fi17, F1: 0, F2: 100}
                        if p1 {
                            if len(v18) > 0 {
                                v18[0] = (fi17 << ((1) & 63))
                            }
                            v21 := GenS1{F0: fi17, F1: 32767, F2: -101}
                            if p1 {
                                v22 := GenS1{F0: fi17, F1: 358, F2: 1}
                                _ = v22
                            } else {
                                var v23 uint64 = 42
                                _ = len(v18)
                                _ = v23
                            }
                            _ = v21
                        } else {
                            p1 = true
                            var v24 int8 = 100
                            if (4294967295 > 255) {
                                v25 := GenS2{F0: p1, F1: 65535}
                                switch "hello" {
                                case "foo bar":
                                    v26 := GenS3{F0: 660, F1: 2147483647, F2: 2147483647}
                                    _ = v26
                                }
                                _ = v25
                            }
                            _ = v24
                        }
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                for fi27 := 0; fi27 < 1; fi27++ {
                    if (true || ((p0 << ((42) & 63)) > (p0 / ((114) | 1)))) {
                        var v28 uint16 = 566
                        _ = v28
                    }
                    p1 = p1
                    _ = fi27
                }
            }
            _ = fi13
        }
        var v29 uint32 = 226
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int64 = 42
            v32 := GenS2{F0: true, F1: 42}
            v31 -= v31
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = v29
    }
    _ = p0
    _ = p1
    return p0
}
func genFn1(p0 int64, p1 uint64, p2 uint64) int8 {
    var v33 float64 = ((1e10 - float64(-2.25)) / -2.25)
    v34 := []int{-959}
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    return 127
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 511
    var v4 uint16 = 765
    var v5 float32 = (3.14159 + 0.0)
    var v6 uint16 = v4
    v35 := []int{}
    for fi36 := 0; fi36 < 19; fi36++ {
        v4 = v6
        _ = fi36
    }
    var v37 int16 = -533
    if true {
        fmt.Println(v37, v37)
        for fi38 := 0; fi38 < 20; fi38++ {
            v5 -= float32(1e10)
            v39 := v2[v3]
            _ = fi38
            _ = v39
        }
        v40 := GenS1{F0: -9223372036854775808, F1: (v6 | v4), F2: -38}
        _ = v40
    }
    v1 += (v1 + (1e10 + v1))
    v3 = (v3 % ((v3) | 1))
    v6 = 236
    v41 := GenS1{F0: 142, F1: v6, F2: genFn1(733, 1, 4294967295)}
    if true {
        fmt.Println(v6, v5, v5)
    }
    v1 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
