package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
func genFn0(p0 uint, p1 int16, p2 uint64) int32 {
    if false {
        v5 := []int{}
        var v6 float32 = 1.5
        if (int8(p0) == -125) {
            v6 = v6
            if len(v5) > 0 {
                v5[0] = 1
            }
            v7 := GenS1{F0: 4294967295, F1: 434, F2: 2}
            _ = v7
        }
        _ = v5
        _ = v6
    } else {
        v8 := []int{2147483647, -1}
        p1 += (p1 >> ((p0) & 63))
        p0 = 26
        _ = v8
    }
    var v9 uint64 = 744
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            switch (fi10 * (fi10 - 1)) {
            case 4294967295:
                p2 = 2
            default:
                switch -9223372036854775808 {
                case 4294967295:
                    if (228 > 1) {
                        p1 = p1
                    }
                }
                fmt.Println(p2)
            }
            fmt.Println(p2, p1)
            var v12 string = "foo bar"
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -868
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -965
    v3 := make(map[uint32]string)
    v4 := []int{}
    var v13 uint = 4294967295
    switch "" {
    case "hello":
        for fi14 := 0; fi14 < 15; fi14++ {
            if false {
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 int16 = 32767
                    _ = fi15
                    _ = v16
                }
                v13 -= 4294967295
            }
            v4 = append(v4, (fi14 * (fi14 / ((fi14) | 1))))
            _ = fi14
        }
    case "\n\t":
        _ = len(v4)
    default:
        v2 += genFn0((v13 | (v13 % (((v13 | (v13 - (^(v13))))) | 1))), 456, 138)
    }
    switch -2147483648 {
    case 966:
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int16 = -1
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS1{F0: (v13 % ((42) | 1)), F1: (v2 & (v2 << ((255) & 63))), F2: 255}
    for fi20 := 0; fi20 < 2; fi20++ {
        if true {
            v4 = append(v4, 363)
            for fi21 := 0; fi21 < 7; fi21++ {
                _ = len(v4)
                v13 += v13
                if (!((false && false) || (!true))) {
                    switch "a" {
                    case "foo bar":
                        switch (fi21 & (fi20 << (((v13 ^ (v13 << (((v13 % ((v13) | 1))) & 63)))) & 63))) {
                        case 619:
                            delete(v1, 770)
                            switch "hello" {
                            case "a":
                                fmt.Println(fi21)
                            case "":
                                fmt.Println(fi20, fi20)
                            case "foo bar":
                                var v22 float64 = 3.14159
                                _ = v22
                            }
                            fmt.Println(v2, v2)
                        }
                    case "":
                        v13 += 42
                    default:
                        fmt.Println(fi21, fi20)
                    }
                    if ("" <= ("foo bar" + "")) {
                        v23 := GenS1{F0: 315, F1: genFn0(v13, 42, 665), F2: 32767}
                        fmt.Println(fi21)
                        v2 += v2
                        var v24 int16 = -1
                        var v25 uint64 = 42
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
                v26 := GenS2{F0: "hello", F1: v2, F2: genFn0(v13, 1, 4294967295)}
                delete(v1, 4294967295)
                _ = fi21
                _ = v26
            }
            switch ("hello" + "a") {
            case "hello":
                var v27 uint8 = 1
                _ = v27
            case "a":
                v2 += genFn0(v13, 485, 751)
            default:
                v2 += (v2 & v2)
            }
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 uint32 = 2147483648
            _ = fi28
            _ = v29
        }
        _ = fi20
    }
    _ = len(v4)
    delete(v1, 2)
    v2 -= v2
    var v30 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
