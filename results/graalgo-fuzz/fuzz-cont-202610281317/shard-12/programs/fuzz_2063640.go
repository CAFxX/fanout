package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint, p2 int16) int64 {
    p2 -= -32768
    var v5 float64 = (-2.25 / (3.14159 * 0.1))
    if (true && true) {
        fmt.Println(p0, v5)
        var v6 int16 = (^(p2))
        _ = v6
    } else {
        switch 0 {
        case 4294967295:
            var v7 int64 = 9223372036854775807
            _ = v7
        case 42:
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 string = "\n\t"
                p1 = 2
                for fi10 := 0; fi10 < 3; fi10++ {
                    p0 += p0
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        default:
            var v11 uint = 2
            _ = v11
        }
    }
    v12 := []int{-631, 0, 0, 591}
    var v13 int8 = -101
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v5
    return -1
}

func main() {
    var v1 int16 = -698
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 int8 = 66
    v14 := []int{-787, 32767}
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: (v4 % ((int8(fi15)) | 1))}
        fmt.Println(fi15, v4, fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v4)
    fmt.Println(v1, v1)
    var v17 float32 = (0.0 - float32(float64(((1.5 - (0.0 + (1e10 / float32((3.14159 - ((1e10 * (-2.25 + 1e10)) - 0.0)))))) - 1e10))))
    switch "\n\t" {
    case "hello":
        fmt.Println(v1, v4)
    case "\n\t":
        var v18 uint16 = 32767
        _ = v18
    case "":
        fmt.Println(v17, v4)
    default:
        fmt.Println(v1, v4, v17)
    }
    fmt.Println(v1, v17)
    v4 -= 28
    v19 := GenS1{F0: (v4 / ((42) | 1))}
    fmt.Println(v17, v4, v17)
    fmt.Println(v1, v4)
    switch ("hello" + "\n\t") {
    case "foo bar":
        if false {
            switch int(v4) {
            case 2147483647:
                v17 = 0.1
            case 32767:
                var v20 int32 = -883
                _ = v20
            }
            switch 128 {
            case -764:
                fmt.Println(v17, v17)
            case 42:
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint32 = 511
                    fmt.Println(v17, v4)
                    _ = fi21
                    _ = v22
                }
            }
            if (-2147483648 <= 314) {
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 uint16 = 0
                    var v25 uint64 = 896
                    v4 -= 1
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 13; fi26++ {
                    v4 -= (-(v4))
                    if false {
                        v1 += v1
                    }
                    _ = fi26
                }
            }
        } else {
            for fi27 := 0; fi27 < 18; fi27++ {
                for fi28 := 0; fi28 < 11; fi28++ {
                    if true {
                        var v29 float32 = (v17 + (1.5 + (v17 + float32(1e10))))
                        v17 += v17
                        fmt.Println(v4, fi28, v29)
                        _ = v29
                    } else {
                        _ = len(v14)
                    }
                    _ = fi28
                }
                fmt.Println(v17, v1, v4)
                _ = fi27
            }
            v30 := GenS1{F0: -128}
            _ = v30
        }
    case "":
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    v31 := GenS1{F0: (v4 >> ((255) & 63))}
    v17 = v17
    _ = len(v14)
    switch ("" + ("\n\t" + "")) {
    case "":
        if true {
            switch ("a" + "a") {
            case "hello":
                v4 = 0
            case "":
                if (255 == 0) {
                    v1 += v1
                    var v32 int16 = (-(v1))
                    _ = v32
                } else {
                    v4 = v4
                }
            case "a":
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
        }
    case "foo bar":
        v33 := GenS2{F0: 507}
        _ = v33
    }
    var v34 float32 = ((v17 * (1e10 * -2.25)) * ((v17 * float32(1.5)) / float32((0.1 + (float64(float32((0.0 * float64((v17 - (0.0 * v17)))))) + 0.1)))))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
