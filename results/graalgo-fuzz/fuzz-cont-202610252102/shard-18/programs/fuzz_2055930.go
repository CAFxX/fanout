package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0(p0 int16) int16 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = (p0 >> ((65535) & 63))
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            p0 = p0
            if false {
                p0 -= (p0 >> ((292) & 63))
            }
            p0 = (p0 & -251)
            _ = fi9
        }
        p0 = -1
        _ = fi8
    }
    switch "foo bar" {
    case "\n\t":
        p0 = p0
    case "hello":
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 127:
                v11 := GenS1{F0: (!(false && true))}
                _ = v11
            case 42:
                fmt.Println(fi10)
            case -252:
                fmt.Println(p0, fi10)
            }
            if true {
                p0 -= (^(p0))
            }
            _ = fi10
        }
    case "foo bar":
        p0 += 1
    default:
        p0 = (p0 * p0)
    }
    switch (("" + ("\n\t" + "a")) + "a") {
    case "foo bar":
        p0 -= 1000
    case "\n\t":
        fmt.Println(p0)
    default:
        fmt.Println(p0)
    }
    _ = p0
    return (p0 ^ (p0 << ((18446744073709551615) & 63)))
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 uint8 = 128
    var v5 int64 = -732
    if (true && true) {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{-840, fi12}
            var v14 uint32 = 1
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = 0
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint8 = v4
            v18 := GenS1{F0: false}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    fmt.Println(v1, v1, v1)
    v5 = -9223372036854775808
    var v19 float64 = 1.5
    var v20 float64 = v19
    v21 := []int{4294967295, 598, 42, 42}
    var v22 uint = 255
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            v3[4294967295] = "hello"
            _ = fi23
        }
        if (934 < 732) {
            v5 += v5
            if ((v1 * v1) > 558) {
                fmt.Println(v4, v5)
                if true {
                    var v24 int64 = (v5 >> (((v22 + uint(v1))) & 63))
                    _ = v24
                } else {
                    v25 := v2[42]
                    fmt.Println(v5, v22)
                    _ = v25
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    v27 := GenS2{F0: v4, F1: v1, F2: -9223372036854775808}
                    _ = fi26
                    _ = v27
                }
            } else {
                if true {
                    v28 := GenS1{F0: true}
                    switch "hello" {
                    case "\n\t":
                        _ = len(v2)
                    }
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if ((1e10 <= (0.1 + (1.5 * 0.1))) || false) {
                            for fi30 := 0; fi30 < 7; fi30++ {
                                v31 := GenS1{F0: true}
                                v32 := GenS2{F0: v4, F1: v1, F2: (^(v5))}
                                v4 = v4
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            v20 -= -2.25
                        } else {
                            var v33 uint = v22
                            if true {
                                v19 -= v20
                            }
                            _ = v33
                        }
                        _ = fi29
                    }
                    var v34 uint64 = (^(v1))
                    _ = v34
                }
            }
        } else {
            if true {
                fmt.Println(v5)
                switch 4294967295 {
                case 860:
                    if true {
                        _ = len(v21)
                        v3[2147483648] = ("\n\t" + (("foo bar" + ("foo bar" + "")) + "\n\t"))
                        v19 = 3.14159
                    }
                case 0:
                    v35 := v2[-882]
                    _ = v35
                case -875:
                    v19 += 1e10
                }
            }
            v4 = v4
        }
    } else {
        v36 := GenS2{F0: (v4 / ((127) | 1)), F1: (v1 | (^(v1))), F2: (v5 / ((v5) | 1))}
        _ = v36
    }
    v22 -= 1
    var v37 uint = 42
    for fi38 := 0; fi38 < 9; fi38++ {
        if true {
            v22 = uint(v1)
            var v39 uint8 = v4
            _ = v39
        } else {
            v1 = (v1 + v1)
        }
        _ = fi38
    }
    v22 -= (v22 - 1)
    if true {
        var v40 bool = false
        _ = v40
    } else {
        fmt.Println(v4, v37)
        v4 = v4
    }
    switch (("a" + "a") + "\n\t") {
    case "\n\t":
        v19 = (v20 - float64((3.14159 / 0.1)))
    default:
        v41 := v2[-485]
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
