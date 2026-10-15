package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 int, p2 uint32) bool {
    fmt.Println(p0, p1, p0)
    p2 -= p2
    var v6 float32 = -2.25
    var v7 uint32 = p2
    if true {
        if (true && false) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = float32(0.1)
                for fi10 := 0; fi10 < 4; fi10++ {
                    fmt.Println(p0)
                    v9 += v6
                    _ = fi10
                }
                var v11 int64 = -9223372036854775808
                _ = fi8
                _ = v11
                _ = v9
            }
        }
    } else {
        var v12 float64 = (3.14159 * 1.5)
        var v13 int8 = int8(p2)
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 uint64 = 422
    var v3 int8 = 0
    var v4 uint8 = 1
    var v5 int32 = 1
    var v14 bool = true
    v5 += v5
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "")
    }
    var v15 uint32 = 2
    var v16 float64 = ((float64(1e10) / ((1e10 + -2.25) / 1e10)) * -2.25)
    switch 2147483647 {
    case 127:
        var v17 int32 = v5
        _ = v17
    case 42:
        v15 += 895
    default:
        for fi18 := 0; fi18 < 17; fi18++ {
            if v14 {
                if ((v15 >= (v15 | (v15 / ((v15) | 1)))) || v14) {
                    v14 = (false && true)
                }
                var v19 uint16 = 42
                _ = v19
            }
            _ = fi18
        }
    }
    var v20 bool = v14
    if v20 {
        var v21 int64 = -9223372036854775808
        switch ("\n\t" + "") {
        case "a":
            switch 32767 {
            case -9223372036854775808:
                if true {
                    v2 = v2
                    v22 := GenS1{F0: v15}
                    fmt.Println(v14, v3, v5)
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        var v24 uint8 = 144
                        fmt.Println(v20, v16, v21)
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        _ = fi23
                        _ = v24
                    }
                    fmt.Println(v16)
                }
            }
        default:
            v5 -= v5
        }
        if false {
            v2 = 0
        }
        _ = v21
    }
    switch "foo bar" {
    case "foo bar":
        if (!v14) {
            v25 := GenS2{F0: v4, F1: -1, F2: 963}
            v3 -= 0
            v2 -= 42
            _ = v25
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                for fi27 := 0; fi27 < 12; fi27++ {
                    v5 -= (v5 * v5)
                    v28 := GenS3{F0: true}
                    v16 -= v16
                    _ = fi27
                    _ = v28
                }
                var v29 int16 = 0
                _ = fi26
                _ = v29
            }
        }
    default:
        for fi30 := 0; fi30 < 13; fi30++ {
            switch "" {
            case "foo bar":
                var v31 uint16 = 346
                _ = v31
            case "a":
                if (42 < 1) {
                    switch (fi30 * fi30) {
                    case 0:
                        if v20 {
                            v4 -= 157
                            v5 -= (v5 << ((255) & 63))
                        }
                    }
                } else {
                    var v32 float64 = 0.1
                    _ = v32
                }
            }
            _ = fi30
        }
    }
    _ = len(v1)
    v2 = v2
    v33 := GenS3{F0: (!false)}
    var v34 uint16 = 385
    v5 += (v5 >> ((126) & 63))
    if (!((2 <= 42) || ((v5 / ((-1) | 1)) >= v5))) {
        if (v14 || ((v5 & v5) < (^(v5)))) {
            fmt.Println(v15, v14, v3)
            var v35 float32 = (float32(v16) / 1.5)
            _ = v35
        } else {
            if ((v4 - 128) < v4) {
                v36 := GenS1{F0: 2}
                _ = v36
            }
            v1 = append(v1, ("\n\t" + "\n\t"))
        }
        var v37 int16 = 32767
        _ = v37
    } else {
        v3 += v3
        switch "hello" {
        case "hello":
            var v38 int = 4294967295
            _ = v38
        }
        if (!(100 < -128)) {
            var v39 float32 = 3.14159
            _ = v39
        } else {
            v4 = 2
        }
    }
    var v40 int = -9223372036854775808
    if (!true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            if (int16(v2) < -1) {
                var v41 uint32 = v15
                var v42 uint = 866
                v41 += v41
                _ = v41
                _ = v42
            }
        }
    }
    var v43 uint = uint(v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
