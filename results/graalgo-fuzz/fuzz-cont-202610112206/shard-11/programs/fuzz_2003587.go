package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 int16, p1 int) int8 {
    v7 := []int{}
    v7 = append(v7, (p1 >> ((255) & 63)))
    p1 += p1
    v7 = append(v7, (-(p1)))
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 119
    var v2 float64 = (((float64(1.5) + 0.1) * -2.25) - 0.1)
    var v3 int = -2147483648
    v4 := make(map[string]string)
    var v5 float64 = (1e10 - v2)
    var v6 int8 = 29
    if true {
        var v8 uint8 = (v1 >> ((255) & 63))
        if (true || false) {
            fmt.Println(v3, v2)
            var v9 uint32 = 2147483648
            switch (^(v3)) {
            case 970:
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch "" {
                    case "foo bar":
                        if true {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                v5 -= 0.0
                                _ = len(v4)
                                _ = fi11
                            }
                        } else {
                            v5 = v5
                        }
                    case "\n\t":
                        if (true || false) {
                            var v12 int64 = -1
                            v6 = v6
                            var v13 string = ("foo bar" + ("hello" + ""))
                            _ = v12
                            _ = v13
                        }
                    }
                    _ = fi10
                }
            case 225:
                fmt.Println(v5, v5)
            }
            _ = v9
        }
        v14 := GenS1{F0: v6}
        _ = v14
        _ = v8
    }
    switch ("" + ((("\n\t" + "\n\t") + "hello") + ("" + "hello"))) {
    case "hello":
        v15 := GenS2{F0: 213, F1: (v1 / (((v1 << ((753) & 63))) | 1)), F2: v3}
        _ = v15
    case "\n\t":
        v16 := v4["foo bar"]
        _ = v16
    case "a":
        var v17 int = -1
        _ = v17
    default:
        v3 = -2147483648
    }
    v18 := GenS3{F0: 2, F1: genFn0(622, v3), F2: 842}
    if ((((v5 / v5) - ((((v5 - ((((1.5 + v2) - v2) + (1.5 / ((v2 - v5) / v2))) / 3.14159)) + (v5 * v5)) + v2) + ((float64(float32(-2.25)) - (v5 + ((v2 * 0.1) / v5))) * (0.0 + (float64(0.0) / v5))))) > 3.14159) || false) {
        v1 = (v1 >> ((140) & 63))
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v5, v6, v2)
            v3 = (fi19 | (v3 * v3))
            fmt.Println(fi19, fi19, v5)
            _ = fi19
        }
        switch v3 {
        case -841:
            if (2 <= 2) {
                v20 := []int{-772, (v3 << ((4294967295) & 63)), v3, v3}
                _ = v20
            }
        default:
            if (!((false || true) || true)) {
                var v21 uint32 = 2147483648
                _ = v21
            } else {
                v22 := []int{(v3 - v3), v3}
                if (9223372036854775807 == 42) {
                    v23 := GenS2{F0: (v1 >> ((4294967295) & 63)), F1: (v1 - v1), F2: (v3 | 9223372036854775807)}
                    var v24 uint8 = (v1 + 0)
                    v1 = (v1 % (((v24 % (((v24 << ((253) & 63))) | 1))) | 1))
                    _ = v23
                    _ = v24
                } else {
                    v5 -= v2
                    switch (v3 | v3) {
                    case 32767:
                        for fi25 := 0; fi25 < 15; fi25++ {
                            for fi26 := 0; fi26 < 4; fi26++ {
                                var v27 float32 = -2.25
                                if false {
                                    fmt.Println(v2, fi26, v2)
                                    var v28 int16 = 0
                                    _ = v28
                                }
                                _ = fi26
                                _ = v27
                            }
                            _ = fi25
                        }
                    case 892:
                        for fi29 := 0; fi29 < 13; fi29++ {
                            v4["foo bar"] = "a"
                            for fi30 := 0; fi30 < 4; fi30++ {
                                v31 := GenS2{F0: (v1 % ((255) | 1)), F1: v1, F2: (v3 | v3)}
                                _ = fi30
                                _ = v31
                            }
                            var v32 uint64 = 1
                            _ = fi29
                            _ = v32
                        }
                    case 1:
                        v1 += (v1 - 216)
                    default:
                        for fi33 := 0; fi33 < 5; fi33++ {
                            var v34 uint32 = 42
                            var v35 uint64 = 131
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                }
                var v36 int16 = -1
                _ = v22
                _ = v36
            }
        }
    }
    fmt.Println(v6, v5)
    v1 = (v1 / (((v1 * (^(v1)))) | 1))
    var v37 uint8 = (v1 + v1)
    var v38 uint = 655
    fmt.Println(v3, v3)
    v5 += (float64(0.1) - (v2 / (((float64((-2.25 * 1.5)) + (0.1 / float64(3.14159))) - (v5 + v2)) / v5)))
    var v39 int8 = (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
