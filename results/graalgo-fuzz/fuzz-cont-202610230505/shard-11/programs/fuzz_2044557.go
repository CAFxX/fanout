package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 int; F1 int; F2 int }

func main() {
    var v1 int8 = 115
    var v2 int16 = -1
    var v3 int32 = 892
    v4 := make(map[int]string)
    var v5 int8 = 127
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint16 = 260
        v6 = (v6 << ((0) & 63))
        _ = fi7
        _ = v8
    }
    if ((v2 / (((v2 & v2)) | 1)) <= (v2 * -32768)) {
        switch (("" + "hello") + "foo bar") {
        case "foo bar":
            var v9 int16 = 32767
            _ = v9
        case "a":
            if false {
                v10 := []int{-59, 568, 32767, 1}
                v2 = v2
                _ = v10
            } else {
                var v11 uint64 = 763
                _ = v11
            }
        default:
            if false {
                var v12 float32 = float32((-2.25 + 3.14159))
                _ = v12
            } else {
                v3 = -493
            }
        }
        v1 -= (v1 + (v1 - 127))
        v6 -= (v6 >> ((4294967295) & 63))
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{(fi13 / ((fi13) | 1)), (fi13 - 42), 4294967295, 598}
        v3 += v3
        _ = fi13
        _ = v14
    }
    switch ("a" + ("foo bar" + "foo bar")) {
    case "hello":
        if true {
            v15 := []int{537, 67, 32767, int(v2)}
            for fi16 := 0; fi16 < 15; fi16++ {
                if false {
                    var v17 int = -1
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 uint16 = 42
                        _ = fi18
                        _ = v19
                    }
                    _ = v17
                } else {
                    var v20 float64 = (3.14159 / (float64(1.5) - 1e10))
                    _ = v20
                }
                _ = fi16
            }
            _ = v15
        }
    }
    var v21 float32 = float32((1e10 / 0.1))
    var v22 uint16 = 943
    if false {
        v23 := GenS1{F0: v6}
        _ = v23
    } else {
        var v24 int8 = (v5 & (v5 ^ (v5 >> ((4294967295) & 63))))
        v25 := GenS2{F0: 32768}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 int8 = v1
        v28 := GenS2{F0: 128}
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            var v31 float64 = 1e10
            switch "foo bar" {
            case "a":
                _ = len(v4)
            default:
                if false {
                    switch ("hello" + "") {
                    case "\n\t":
                        v2 += v2
                    case "":
                        switch 42 {
                        case -244:
                            v32 := []int{}
                            _ = v32
                        case -2147483648:
                            var v33 bool = false
                            _ = v33
                        }
                    }
                    var v34 float64 = 0.1
                    v35 := GenS3{F0: (fi30 | fi30), F1: 78, F2: v1}
                    _ = v34
                    _ = v35
                } else {
                    v1 = int8(v2)
                }
            }
            var v36 int16 = 1000
            v37 := v4[fi30]
            _ = fi30
            _ = v31
            _ = v36
            _ = v37
        }
        if (3.14159 <= (1e10 - 0.1)) {
            v38 := GenS4{F0: (fi29 - (fi29 & fi29)), F1: fi29, F2: fi29}
            fmt.Println(v5, v21)
            for fi39 := 0; fi39 < 4; fi39++ {
                v22 -= 921
                v5 += v5
                if (false || false) {
                    v3 += v3
                    v22 = (^(v22))
                }
                _ = fi39
            }
            _ = v38
        }
        if false {
            for fi40 := 0; fi40 < 7; fi40++ {
                v41 := []int{fi29}
                _ = fi40
                _ = v41
            }
        }
        _ = fi29
    }
    v22 += (v22 % (((v22 & v22)) | 1))
    var v42 int8 = 127
    _ = len(v4)
    fmt.Println(v5, v42)
    v43 := GenS3{F0: 4294967295, F1: (v3 - 42), F2: int8(v22)}
    var v44 int32 = 42
    if ((false || ((v3 == 1) && (("\n\t" + ("foo bar" + "foo bar")) <= "hello"))) && (v21 > (((1e10 * (v21 + (float32(3.14159) / (0.1 * v21)))) - 1e10) / ((v21 - v21) + (v21 / 0.1))))) {
        if false {
            v5 += (v5 - v1)
            var v45 int16 = v2
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
