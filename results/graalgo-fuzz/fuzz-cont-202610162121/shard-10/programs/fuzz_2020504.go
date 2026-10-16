package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint64) int64 {
    var v4 string = "a"
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if false {
                p0 -= p0
            }
            v4 = "\n\t"
            _ = fi6
        }
        var v7 uint64 = p0
        _ = fi5
        _ = v7
    }
    v4 = ((v4 + "foo bar") + v4)
    _ = p0
    _ = v4
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64(-2.25)
    var v3 int32 = 2147483647
    v8 := v1[(("hello" + ("" + "")) + ("a" + ("\n\t" + "foo bar")))]
    v8 -= v8
    v9 := v1["\n\t"]
    v8 += v8
    v2 = v2
    switch -1 {
    case -2147483648:
        var v10 int = 2147483647
        _ = v10
    case 128:
        switch 0 {
        case 127:
            v3 += -1
        case -46:
            v9 += (v9 % (((v8 % ((953) | 1))) | 1))
        case 1:
            v8 -= (v9 << ((520) & 63))
        }
    }
    if ((-(v3)) != v3) {
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 7; fi11++ {
                switch fi11 {
                case 4294967295:
                    v12 := []int{-9223372036854775808, 32767, fi11}
                    _ = v12
                }
                _ = fi11
            }
        case "hello":
            if (-947 > -951) {
                var v13 uint8 = 127
                var v14 uint32 = (v9 + v9)
                switch -2147483648 {
                case 0:
                    v13 -= (v13 | 46)
                }
                _ = v13
                _ = v14
            }
        }
    }
    var v15 float64 = 0.1
    if true {
        v2 += 1.5
        var v16 uint = 1
        if false {
            for fi17 := 0; fi17 < 16; fi17++ {
                switch ("hello" + "") {
                case "":
                    v1[("hello" + "\n\t")] = 65535
                }
                _ = fi17
            }
        }
        _ = v16
    } else {
        var v18 int16 = 441
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int16 = 1
            delete(v1, (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + (("\n\t" + "a") + "")))
            fmt.Println(fi19)
            _ = fi19
            _ = v20
        }
        if (1 == 0) {
            v18 -= (v18 << ((65535) & 63))
            var v21 string = ("foo bar" + ("" + (("a" + ("hello" + ("foo bar" + "hello"))) + ("" + ("foo bar" + "hello")))))
            _ = v21
        } else {
            var v22 uint8 = 42
            fmt.Println(v15, v8, v3)
            var v23 int32 = (v3 * (^(v3)))
            _ = v22
            _ = v23
        }
        _ = v18
    }
    var v24 uint = 0
    if true {
        if (false || (!false)) {
            var v25 uint32 = (v9 ^ 2147483648)
            _ = v25
        }
        v8 = (v9 & (v8 / ((v8) | 1)))
        if ((!true) && true) {
            var v26 int16 = -598
            v2 += ((v2 * v15) + float64(float32(1.5)))
            _ = v26
        }
    }
    if true {
        var v27 uint8 = 128
        for fi28 := 0; fi28 < 1; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                v8 = 42
                v8 = (v8 ^ 2)
                v27 += v27
                _ = fi29
            }
            var v30 int64 = 719
            v15 = ((v2 / (v15 / 1.5)) + v2)
            _ = fi28
            _ = v30
        }
        v24 = (^(v24))
        _ = v27
    }
    v31 := GenS1{F0: 2147483647, F1: uint8(v24), F2: (^(v3))}
    switch -2147483648 {
    case -1:
        var v32 float32 = ((0.1 / -2.25) * ((3.14159 / -2.25) - ((1e10 + float32(v15)) + 1.5)))
        _ = v32
    case 4294967295:
        var v33 uint16 = 353
        _ = v33
    case 9:
        fmt.Println(v8, v24, v3)
    }
    switch 2147483647 {
    case 404:
        v34 := []int{-690, 0, 42}
        _ = v34
    case 128:
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := GenS2{F0: (-(fi35))}
            _ = fi35
            _ = v36
        }
    case -626:
        switch "hello" {
        case "\n\t":
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int64 = 1
                _ = fi37
                _ = v38
            }
        default:
            for fi39 := 0; fi39 < 14; fi39++ {
                for fi40 := 0; fi40 < 2; fi40++ {
                    var v41 float64 = v2
                    _ = fi40
                    _ = v41
                }
                switch "hello" {
                case "a":
                    v15 = ((v2 + v2) / ((float64(0.0) - 0.1) / v15))
                case "foo bar":
                    v42 := GenS1{F0: fi39, F1: 128, F2: -969}
                    _ = v42
                case "hello":
                    var v43 float64 = v15
                    _ = v43
                default:
                    var v44 uint = (v24 & 1)
                    _ = v44
                }
                v45 := []int{fi39, fi39, fi39, (fi39 | (fi39 % ((fi39) | 1)))}
                _ = fi39
                _ = v45
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v8
    _ = v9
}
