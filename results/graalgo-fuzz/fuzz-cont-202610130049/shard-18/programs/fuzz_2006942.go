package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int }
func genFn0() int {
    if (!(float64(float32(1.5)) == 1.5)) {
        if (("a" + "\n\t") <= "") {
            var v6 uint8 = 128
            _ = v6
        } else {
            var v7 uint = 742
            v8 := GenS1{F0: -128, F1: 0}
            v9 := []int{32767, 4294967295, 265, -2147483648}
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: 65535, F1: 1}
        _ = v10
    } else {
        var v11 uint32 = 32
        var v12 bool = true
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                v11 = (v11 * v11)
                var v14 int16 = 990
                _ = v14
            } else {
                v11 = v11
                var v15 int16 = -11
                _ = v15
            }
            v11 += v11
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    var v16 int16 = -32768
    switch "a" {
    case "hello":
        var v17 bool = false
        _ = v17
    default:
        fmt.Println(v16)
    }
    _ = v16
    return 128
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 2147483647
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]int)
    v3 = 255
    if true {
        v18 := []int{}
        _ = v18
    }
    var v19 bool = (2147483647 != -2147483648)
    switch v1 {
    case "":
        for fi20 := 0; fi20 < 8; fi20++ {
            if ((0.1 + 0.1) == 0.0) {
                fmt.Println(fi20, v2)
                v21 := []int{fi20}
                _ = v21
            }
            var v22 uint16 = v3
            _ = fi20
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 13; fi23++ {
            if true {
                v1 = "foo bar"
                switch v1 {
                case "\n\t":
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v2 -= v2
                        _ = fi24
                    }
                case "":
                    switch fi23 {
                    case 2147483647:
                        if v19 {
                            fmt.Println(v1)
                            var v25 uint = 1
                            var v26 float32 = ((((1e10 / (-2.25 * float32(1.5))) + (-2.25 - 0.1)) + 1e10) + (((float32(1e10) / 3.14159) + (-2.25 + float32((0.1 * (((3.14159 * (3.14159 / 1.5)) * 0.1) / 1.0))))) - -2.25))
                            _ = v25
                            _ = v26
                        }
                    case 0:
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 float32 = float32(((float64(float32(0.1)) * 3.14159) + 1.5))
                            fmt.Println(v2, fi27)
                            fmt.Println(fi27)
                            v29 := GenS2{F0: 0, F1: 0}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                case "foo bar":
                    fmt.Println(v1)
                }
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                v1 = (v1 + v1)
                _ = fi30
            }
            _ = fi23
        }
    }
    v3 = (v3 | (v3 + v3))
    fmt.Println(v2, v19, v1)
    v3 -= (v3 + (^(v3)))
    var v31 uint = 1
    v32 := GenS1{F0: 1, F1: 54}
    if (!false) {
        if v19 {
            var v33 float64 = 0.0
            var v34 string = ((v1 + v1) + v1)
            v35 := []int{2147483647, 870}
            _ = v33
            _ = v34
            _ = v35
        }
        var v36 int32 = -2147483648
        _ = v36
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint8 = 0
        var v39 string = v1
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 int8 = -128
            switch genFn0() {
            case -631:
                v42 := v4[(v2 ^ 0)]
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    v43 := GenS3{F0: 1}
    v44 := GenS2{F0: (v31 << ((65535) & 63)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
