package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 string, p1 uint32, p2 uint16) int64 {
    if (-1 >= 1) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int16 = 32767
            _ = fi6
            _ = v7
        }
        p0 = p0
    } else {
        switch p0 {
        case "foo bar":
            if false {
                v8 := []int{400}
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(p1)
                    p1 += p1
                    _ = fi9
                }
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch fi10 {
                    case 2147483647:
                        var v11 uint = 422
                        _ = v11
                    case 4294967295:
                        if true {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                var v13 int8 = 100
                                _ = fi12
                                _ = v13
                            }
                            for fi14 := 0; fi14 < 14; fi14++ {
                                var v15 bool = ((127 > -114) || true)
                                _ = fi14
                                _ = v15
                            }
                            p2 += p2
                        }
                    }
                    _ = fi10
                }
                _ = v8
            }
        case "\n\t":
            p1 = 103
        }
        fmt.Println(p2, p2, p1)
        fmt.Println(p0, p1)
    }
    var v16 float32 = 0.0
    var v17 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 9223372036854775807
}
func genFn1(p0 int32, p1 int64, p2 int64) int32 {
    fmt.Println(p1)
    for fi18 := 0; fi18 < 6; fi18++ {
        p2 += (p1 << ((2) & 63))
        if false {
            p0 += 183
            if ((^(p0)) != (p0 + (p0 & (p0 >> ((78) & 63))))) {
                if false {
                    p1 = (p2 + p1)
                    if (100 <= 1) {
                        p1 -= -9223372036854775808
                        var v19 bool = true
                        p2 -= -1
                        _ = v19
                    } else {
                        var v20 uint32 = 42
                        for fi21 := 0; fi21 < 7; fi21++ {
                            var v22 uint = 1
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                    p1 += (p2 % (((p1 << ((42) & 63))) | 1))
                }
            }
        } else {
            if true {
                p0 = (p0 >> ((42) & 63))
            }
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 & (p0 | 1))
}

func main() {
    var v1 uint8 = 2
    var v2 float32 = (((0.1 - ((0.0 * (float32((0.0 / 1.0)) + 1.5)) - 1.5)) + float32((3.14159 / (1.5 * 3.14159)))) * 3.14159)
    var v3 int16 = 32767
    v4 := make(map[int]bool)
    var v5 string = ("\n\t" + ((("\n\t" + "hello") + ("foo bar" + "\n\t")) + (("\n\t" + "foo bar") + "a")))
    var v23 float32 = ((0.0 / ((0.0 / v2) + float32(0.1))) / v2)
    for fi24 := 0; fi24 < 19; fi24++ {
        if false {
            fmt.Println(v2, fi24)
        } else {
            switch fi24 {
            case 841:
                var v25 int = 9223372036854775807
                _ = v25
            }
            var v26 uint32 = 2147483648
            if (!true) {
                switch "a" {
                case "":
                    fmt.Println(v2)
                case "a":
                    v26 = v26
                case "foo bar":
                    var v27 int8 = -1
                    _ = v27
                }
                v28 := GenS1{F0: 2}
                for fi29 := 0; fi29 < 7; fi29++ {
                    fmt.Println(v3, v5, v2)
                    v30 := []int{fi29, 127}
                    var v31 int = 127
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v28
            } else {
                v3 -= v3
            }
            _ = v26
        }
        _ = fi24
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint32 = 0
            var v35 uint = 255
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    v2 += ((3.14159 / (v2 - (0.1 - 3.14159))) / (v2 + v23))
    var v36 string = ((v5 + "hello") + (v5 + ""))
    v3 = v3
    var v37 uint8 = 255
    v38 := []int{4294967295, -9223372036854775808, 1, 9223372036854775807}
    var v39 int16 = -32768
    v40 := GenS1{F0: 1}
    var v41 float64 = float64(v23)
    for fi42 := 0; fi42 < 3; fi42++ {
        if false {
            v36 = ((v5 + "") + "a")
            v5 = v36
            delete(v4, 2147483647)
        } else {
            if false {
                _ = len(v4)
            }
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
