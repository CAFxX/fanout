package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 int, p2 uint) uint8 {
    v4 := []int{476, p1, (p1 % ((p1) | 1)), (p1 / ((p1) | 1))}
    if true {
        if len(v4) > 0 {
            v4[0] = (p1 - (p1 + 1))
        }
    } else {
        if p0 {
            switch "foo bar" {
            case "a":
                for fi5 := 0; fi5 < 14; fi5++ {
                    p2 = 255
                    p0 = p0
                    _ = fi5
                }
            case "\n\t":
                for fi6 := 0; fi6 < 14; fi6++ {
                    v4 = append(v4, fi6)
                    v4 = append(v4, p1)
                    var v7 uint8 = 0
                    _ = fi6
                    _ = v7
                }
            case "foo bar":
                _ = len(v4)
            default:
                if len(v4) > 0 {
                    v4[0] = p1
                }
            }
            v4 = append(v4, p1)
            var v8 uint8 = 2
            _ = v8
        } else {
            var v9 int16 = 0
            fmt.Println(v9, p2)
            _ = v9
        }
        v4 = append(v4, (p1 & p1))
        if (51 == 1) {
            if p0 {
                switch ((((("\n\t" + ("a" + "foo bar")) + (("hello" + ("foo bar" + "a")) + ("foo bar" + ("foo bar" + "foo bar")))) + "") + "foo bar") + "a") {
                case "\n\t":
                    var v10 uint = (^(p2))
                    _ = v10
                case "":
                    for fi11 := 0; fi11 < 8; fi11++ {
                        for fi12 := 0; fi12 < 1; fi12++ {
                            _ = len(v4)
                            v13 := GenS1{F0: 42, F1: fi12}
                            _ = fi12
                            _ = v13
                        }
                        _ = len(v4)
                        _ = fi11
                    }
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint = p2
                    _ = fi14
                    _ = v15
                }
            } else {
                switch (-(p1)) {
                case -460:
                    var v16 int = (-(p1))
                    _ = v16
                }
                p0 = p0
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 7; fi18++ {
                    p0 = p0
                    fmt.Println(fi18, fi18, p0)
                    _ = fi18
                }
                p0 = p0
                _ = fi17
            }
            p1 += p1
        }
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float64 = 0.1
        v4 = append(v4, (p1 & 0))
        _ = fi19
        _ = v20
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v21 uint = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 42
    var v3 uint64 = 329
    v2 = (^(v2))
    var v22 uint64 = v3
    v22 += v3
    if (0 > 42) {
        v3 = v22
        if (3.14159 != float32(0.0)) {
            var v23 int64 = -1
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v2, v3, v22)
                for fi25 := 0; fi25 < 16; fi25++ {
                    v3 -= v22
                    _ = fi25
                }
                if true {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        switch ("a" + "a") {
                        case "":
                            _ = len(v1)
                        default:
                            v22 = v22
                        }
                        _ = fi26
                    }
                    var v27 uint64 = (v22 / (((v22 % ((v3) | 1))) | 1))
                    _ = v27
                }
                _ = fi24
            }
            _ = v23
        }
        v28 := []int{-9223372036854775808, -9223372036854775808, 32768}
        _ = v28
    } else {
        v29 := GenS2{F0: -14}
        v3 += (v3 * 42)
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := []int{}
                var v33 int = fi30
                var v34 int16 = 1000
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi30
        }
        _ = v29
    }
    var v35 int16 = -666
    var v36 uint32 = v2
    v36 += (v36 * 2)
    v2 = 2147483647
    fmt.Println(v2, v22)
    v1[32767] = true
    switch "hello" {
    case "a":
        for fi37 := 0; fi37 < 6; fi37++ {
            v1[127] = true
            fmt.Println(fi37, v2)
            var v38 int = fi37
            _ = fi37
            _ = v38
        }
    default:
        var v39 bool = (697 > 4294967295)
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
}
