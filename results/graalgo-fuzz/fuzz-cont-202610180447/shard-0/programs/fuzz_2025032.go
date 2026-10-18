package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint32 }
func genFn0() uint32 {
    if (32767 > -1) {
        var v7 float32 = 0.1
        fmt.Println(v7)
        var v8 int64 = 257
        _ = v7
        _ = v8
    }
    v9 := []int{}
    var v10 int32 = -591
    v10 = 0
    _ = v10
    _ = v9
    return 681
}
func genFn1(p0 uint, p1 uint8, p2 uint) int32 {
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(fi11, fi11)
        p1 -= (p1 - 255)
        var v12 uint8 = p1
        _ = fi11
        _ = v12
    }
    if (true && false) {
        for fi13 := 0; fi13 < 6; fi13++ {
            if false {
                v14 := GenS1{F0: 4294967295, F1: false}
                _ = v14
            }
            var v15 uint64 = 323
            fmt.Println(p2)
            _ = fi13
            _ = v15
        }
    }
    p2 = 2
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 int16 = -664
    var v2 float64 = 0.0
    var v3 string = (((("\n\t" + "foo bar") + ("\n\t" + ((("\n\t" + "hello") + ("\n\t" + "foo bar")) + ""))) + "\n\t") + "hello")
    var v4 uint16 = 1
    v5 := []int{4294967295, 127, 1, 951}
    v6 := make(map[uint32]bool)
    v16 := GenS1{F0: 342, F1: (v3 <= ((("\n\t" + "a") + ("hello" + (v3 + v3))) + "hello"))}
    switch (("a" + (v3 + "a")) + "a") {
    case "a":
        switch ((v3 + "a") + "a") {
        case "":
            var v17 uint = 42
            _ = v17
        case "hello":
            var v18 int8 = -1
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 uint64 = 927
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        switch 32767 {
        case 0:
            if len(v5) > 0 {
                v5[0] = 647
            }
        case -9223372036854775808:
            switch fi21 {
            case -216:
                v6[1] = (v4 < (v4 >> ((255) & 63)))
            case 32768:
                switch ("foo bar" + (v3 + "")) {
                case "foo bar":
                    var v22 int = fi21
                    _ = v22
                case "\n\t":
                    switch 127 {
                    case -38:
                        v1 = v1
                    }
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            case -1:
                v1 = v1
            default:
                var v23 int8 = 0
                _ = v23
            }
        case -2147483648:
            var v24 uint = 1
            _ = v24
        default:
            for fi25 := 0; fi25 < 20; fi25++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch (fi25 / ((fi21) | 1)) {
                case -934:
                    v3 = ("" + ((v3 + v3) + ""))
                case 594:
                    _ = len(v5)
                case -1:
                    v4 = v4
                }
                v26 := GenS1{F0: 2, F1: (((0.1 / -2.25) + -2.25) <= float32(v2))}
                _ = fi25
                _ = v26
            }
        }
        v3 = ("" + ("" + (v3 + v3)))
        _ = fi21
    }
    var v27 string = v3
    v28 := GenS1{F0: 18446744073709551615, F1: (false || (659 > 2147483648))}
    for fi29 := 0; fi29 < 6; fi29++ {
        fmt.Println(v3, v27)
        switch 128 {
        case 42:
            v1 = v1
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            fmt.Println(v4, v1)
            v31 := GenS1{F0: 4294967295, F1: (42 <= 42)}
            var v32 float64 = v2
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    var v33 string = v27
    v3 = "a"
    var v34 int16 = (-(v1))
    _ = len(v6)
    v2 = float64(-2.25)
    v35 := GenS2{F0: 0, F1: 42, F2: 0}
    for fi36 := 0; fi36 < 9; fi36++ {
        if (0 != 4294967295) {
            for fi37 := 0; fi37 < 16; fi37++ {
                fmt.Println(v4)
                _ = fi37
            }
        } else {
            if (!true) {
                v27 = (v3 + "")
            } else {
                for fi38 := 0; fi38 < 6; fi38++ {
                    v1 += (v1 + 42)
                    _ = fi38
                }
            }
            v39 := GenS1{F0: 37, F1: (!((2147483647 >= 2) && false))}
            var v40 uint = 255
            _ = v39
            _ = v40
        }
        _ = fi36
    }
    var v41 uint64 = 2
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 int8 = -1
        fmt.Println(v27)
        _ = fi42
        _ = v43
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
