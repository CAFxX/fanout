package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 533
    var v3 float64 = float64(0.1)
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 1, F1: 127}
        _ = fi5
        _ = v6
    }
    v7 := []int{32768}
    v4 = ((v4 + v4) + "\n\t")
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: (-48 == 1)}
        v10 := GenS3{F0: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = v2
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v3 = ((3.14159 + v3) + (v3 / (v3 / (v3 / (v3 - 1.5)))))
            fmt.Println(v3, v2)
            v13 := GenS4{F0: 2, F1: v4}
            _ = fi12
            _ = v13
        }
        if (false || (!(32767 > 0))) {
            v14 := GenS1{F0: 127, F1: 127}
            if (true || (42 > int16(fi11))) {
                v15 := GenS4{F0: 2, F1: ((v4 + ("\n\t" + "\n\t")) + ("foo bar" + "a"))}
                _ = v15
            } else {
                var v16 float64 = v3
                v16 -= (v3 * 3.14159)
                _ = v16
            }
            _ = v14
        }
        _ = fi11
    }
    v17 := v1[v4]
    switch ("foo bar" + ("foo bar" + (v4 + ("foo bar" + (((((v4 + (v4 + v4)) + (v4 + ("a" + "\n\t"))) + "a") + (((v4 + "foo bar") + "") + v4)) + "foo bar"))))) {
    case "a":
        if ((false || (357 >= 716)) && ((!(false && true)) || false)) {
            for fi18 := 0; fi18 < 17; fi18++ {
                if true {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        _ = len(v7)
                        _ = fi19
                    }
                }
                v7 = append(v7, -253)
                _ = fi18
            }
            v17 = 0
        }
    case "":
        if true {
            fmt.Println(v2)
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 uint8 = 2
                    _ = fi21
                    _ = v22
                }
                if true {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v17 += (fi23 - fi20)
                        fmt.Println(v3, v2)
                        var v24 uint64 = (v2 + v2)
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi20
            }
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if (true && (!(523 < 164))) {
            v26 := GenS5{F0: 255}
            v17 += v17
            _ = v26
        } else {
            var v27 int = -1
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 int16 = 1000
            var v30 uint16 = 1
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi25
    }
    switch v17 {
    case 32767:
        var v31 uint64 = (v2 | v2)
        _ = v31
    case 338:
        var v32 int8 = 42
        _ = v32
    case 546:
        var v33 string = (v4 + (v4 + v4))
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
