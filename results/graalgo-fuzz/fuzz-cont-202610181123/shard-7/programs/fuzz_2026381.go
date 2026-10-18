package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int16 = 42
    var v5 uint8 = 143
    v6 := make(map[int]int)
    v5 = (v5 - v5)
    var v7 uint16 = 2
    v8 := GenS1{F0: -1}
    var v9 uint16 = (v7 * (v7 | (v7 << ((75) & 63))))
    for fi10 := 0; fi10 < 9; fi10++ {
        if (0.1 > 3.14159) {
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(v7, v7)
                _ = fi11
            }
            var v12 float32 = 0.0
            _ = v12
        } else {
            v9 += 42
            fmt.Println(v7, v4)
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v7 = 1
            fmt.Println(fi10, v9, v4)
            var v14 int16 = (v4 | (v4 * (v4 - (v4 << ((4294967295) & 63)))))
            v15 := GenS1{F0: 2147483647}
            _ = fi13
            _ = v14
            _ = v15
        }
        v9 = (v9 + v9)
        _ = fi10
    }
    v5 = v5
    if (-1 <= 100) {
        _ = len(v2)
    } else {
        v16 := v3["a"]
        fmt.Println(v16, v4)
        var v17 uint = 672
        _ = v16
        _ = v17
    }
    delete(v1, 155)
    var v18 int16 = v4
    if false {
        v5 += v5
        var v19 uint = 255
        v7 = v7
        _ = v19
    }
    var v20 int16 = (v4 / (((v4 - v18)) | 1))
    v21 := GenS2{F0: uint(v5), F1: v18}
    if (!(true || true)) {
        for fi22 := 0; fi22 < 13; fi22++ {
            switch ("\n\t" + "a") {
            case "a":
                v23 := GenS3{F0: 918, F1: 919}
                _ = v23
            default:
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(v7, v5)
                    switch "\n\t" {
                    case "foo bar":
                        if (-1 < -1) {
                            var v25 int64 = -26
                            fmt.Println(v7, v7)
                            if (!(0.0 <= 0.1)) {
                                switch (fi24 & int(v5)) {
                                case 1:
                                    if (42 > 184) {
                                        v26 := v2[uint32(v7)]
                                        _ = v26
                                    }
                                }
                            }
                            _ = v25
                        }
                    case "":
                        v9 = v9
                    case "\n\t":
                        var v27 uint16 = (v9 & v7)
                        _ = v27
                    default:
                        v28 := GenS1{F0: -935}
                        _ = v28
                    }
                    _ = fi24
                }
            }
            _ = fi22
        }
        fmt.Println(v18, v20, v20)
        if ((((32767 <= (v18 | (v18 * -735))) || true) || (v7 <= uint16(v5))) && (!((false && ((-128 == 100) || true)) && (!false)))) {
            v2[0] = 4294967295
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                v18 += -32768
                _ = fi29
            }
            if (1 >= 913) {
                var v30 string = "foo bar"
                if (false || ((v5 & (v5 - (v5 / ((v5) | 1)))) != v5)) {
                    var v31 int = 2147483647
                    var v32 int32 = -1
                    var v33 uint32 = 42
                    _ = v31
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 uint64 = 42
                    switch (fi34 - (fi34 | fi34)) {
                    case 662:
                        v36 := GenS4{F0: 2147483647}
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
                _ = v30
            }
        }
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := []int{}
            _ = fi37
            _ = v38
        }
        v39 := []int{-2147483648}
        var v40 bool = true
        _ = v39
        _ = v40
    }
    var v41 float32 = (-2.25 - float32((float64((-2.25 + 1.5)) * -2.25)))
    var v42 uint16 = 77
    v7 += 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
