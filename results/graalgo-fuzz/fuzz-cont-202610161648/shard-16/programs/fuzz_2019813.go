package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string, p1 uint8, p2 uint16) uint64 {
    v6 := []int{9223372036854775807, 127, -826, 32768}
    if false {
        var v7 uint64 = 42
        var v8 int32 = 42
        _ = v7
        _ = v8
    }
    var v9 int64 = 42
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint8 = 0
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int = (fi12 / ((fi12) | 1))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 42
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 uint64 = 1
    v4 := make(map[int]int)
    var v5 uint8 = 128
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: -791}
        var v16 int = (-(fi14))
        _ = fi14
        _ = v15
        _ = v16
    }
    v5 -= 255
    var v17 string = ((("hello" + "\n\t") + "hello") + "\n\t")
    switch ("" + v17) {
    case "hello":
        var v18 int8 = 42
        _ = v18
    case "\n\t":
        fmt.Println(v17)
    case "":
        var v19 int32 = -551
        _ = v19
    }
    if (!(1 > 163)) {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 uint16 = 874
            v22 := GenS2{F0: (^(v5)), F1: 0}
            for fi23 := 0; fi23 < 13; fi23++ {
                v3 += 1
                var v24 uint32 = 65535
                v5 -= 137
                _ = fi23
                _ = v24
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        if (!true) {
            v17 = v17
        } else {
            var v25 float32 = ((3.14159 + -2.25) * float32(1.5))
            v17 = v17
            _ = v25
        }
        var v26 uint32 = 2147483648
        _ = v26
    }
    fmt.Println(v17, v17, v5)
    var v27 uint = 255
    v28 := GenS2{F0: 255, F1: 238}
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            v31 := v4[42]
            var v32 int = fi30
            _ = fi30
            _ = v31
            _ = v32
        }
        if ((false && false) && true) {
            if true {
                for fi33 := 0; fi33 < 1; fi33++ {
                    fmt.Println(v3, fi33, fi33)
                    if (65535 != 255) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        switch (v17 + "\n\t") {
                        case "hello":
                            _ = len(v1)
                        }
                        v17 = ("hello" + ((v17 + v17) + ("\n\t" + (v17 + v17))))
                        _ = len(v2)
                    }
                    switch v17 {
                    case "\n\t":
                        if (true && true) {
                            if (!false) {
                                v34 := GenS3{F0: (^(v5))}
                                v3 = 1
                                delete(v2, "\n\t")
                                _ = v34
                            } else {
                                v27 = (^(v27))
                                if true {
                                    _ = len(v1)
                                }
                            }
                            if false {
                                v3 -= 806
                            }
                        }
                    case "hello":
                        v3 = v3
                    default:
                        switch (fi33 - (fi33 | fi33)) {
                        case 616:
                            var v35 bool = true
                            _ = v35
                        case -747:
                            switch fi29 {
                            case 568:
                                v27 -= v27
                            }
                        }
                    }
                    v27 -= (v27 & 1)
                    _ = fi33
                }
                v36 := v2[v17]
                _ = len(v1)
                _ = v36
            }
        } else {
            switch (fi29 - fi29) {
            case -1:
                v37 := GenS3{F0: 212}
                _ = v37
            case -167:
                fmt.Println(fi29, v3, fi29)
            case 76:
                if false {
                    v38 := GenS2{F0: (v5 | v5), F1: 2147483648}
                    switch 0 {
                    case 32768:
                        var v39 string = ("hello" + (v17 + (("\n\t" + v17) + v17)))
                        _ = v39
                    }
                    _ = v38
                }
            }
        }
        fmt.Println(v3, v3, v17)
        _ = fi29
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v3 -= (v3 << ((4294967295) & 63))
        fmt.Println(v3)
        if true {
            v27 = (v27 >> ((2) & 63))
        }
        _ = fi40
    }
    var v41 uint = v27
    var v42 uint16 = 783
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
