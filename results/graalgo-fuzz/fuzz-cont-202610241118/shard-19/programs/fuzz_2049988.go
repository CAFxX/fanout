package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 string, p1 int64) int8 {
    p1 -= (p1 ^ 42)
    if (-1 != -9223372036854775808) {
        v4 := GenS1{F0: (p0 + p0)}
        _ = v4
    }
    p1 = 1
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]uint32)
    var v3 uint64 = 1
    var v5 int64 = -997
    if true {
        var v6 int8 = int8(v5)
        var v7 string = ""
        _ = v6
        _ = v7
    }
    var v8 uint8 = 165
    var v9 float32 = (float32(1e10) * 1.5)
    v5 = v5
    var v10 bool = (false && (((0.1 + ((float64(float32((1e10 - ((float64((1.5 + 3.14159)) + (0.1 + (0.0 / 1e10))) / -2.25)))) / 1.0) - 1.5)) / 3.14159) >= 0.1))
    for fi11 := 0; fi11 < 6; fi11++ {
        switch (fi11 << ((1) & 63)) {
        case 42:
            v12 := GenS1{F0: (("" + ("" + ("a" + ""))) + "a")}
            _ = v12
        case -285:
            v5 += v5
        case 32768:
            if (((v8 / ((v8) | 1)) > v8) && ((!v10) && true)) {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    switch "foo bar" {
                    case "hello":
                        v3 = (v3 * 0)
                    case "a":
                        if v10 {
                            if v10 {
                                v5 -= 42
                                fmt.Println(v5)
                                if (65535 != 0) {
                                    v10 = (4294967295 < 255)
                                    v13 := GenS1{F0: ""}
                                    _ = v13
                                }
                            } else {
                                var v14 int8 = v1
                                _ = v14
                            }
                            v15 := GenS2{F0: v1}
                            var v16 float32 = 1.5
                            _ = v15
                            _ = v16
                        } else {
                            var v17 float64 = ((float64(1e10) + 0.0) / (float64((1.5 * (3.14159 / (v9 / 1.0)))) * 0.0))
                            v1 = int8(v3)
                            _ = v17
                        }
                    default:
                        v2[fi11] = 42
                    }
                case "\n\t":
                    switch int(v1) {
                    case 4294967295:
                        v18 := []int{(fi11 * 2147483647), -1, 1}
                        _ = v18
                    default:
                        v19 := GenS3{F0: (v1 >> ((255) & 63))}
                        _ = v19
                    }
                case "a":
                    v5 = v5
                default:
                    v20 := []int{}
                    _ = v20
                }
                var v21 uint8 = (v8 << ((900) & 63))
                _ = v21
            }
        default:
            v1 -= -117
        }
        v8 -= (v8 / ((v8) | 1))
        for fi22 := 0; fi22 < 19; fi22++ {
            v3 += v3
            if v10 {
                var v23 uint = 591
                var v24 uint64 = v3
                fmt.Println(v1, v24)
                _ = v23
                _ = v24
            } else {
                var v25 bool = (((v9 / ((1.5 + v9) - (float32(3.14159) / (((1e10 - float32(((-2.25 * float64(((1e10 * v9) / v9))) * float64(float32((1.5 + 1.5)))))) / v9) - (1e10 + v9))))) + (((v9 / 1.5) - (1.5 / v9)) * ((float32(1e10) - v9) - 1.5))) >= (3.14159 - float32((1e10 + ((((3.14159 - 3.14159) / 1.0) - 1.5) / -2.25)))))
                v5 = (v5 >> ((65535) & 63))
                var v26 float64 = -2.25
                if (1 < 4294967295) {
                    v3 -= (v3 | v3)
                    v2[fi22] = 857
                }
                var v27 int = (fi22 - (fi11 >> ((289) & 63)))
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 float32 = (v9 / (v9 + v9))
            _ = fi22
            _ = v28
        }
        _ = fi11
    }
    switch "foo bar" {
    case "":
        v29 := GenS4{F0: -560}
        _ = v29
    case "hello":
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 float32 = v9
            _ = fi30
            _ = v31
        }
    case "foo bar":
        if v10 {
            fmt.Println(v9, v5, v5)
        } else {
            switch (("\n\t" + (("a" + "foo bar") + ("a" + ("" + ((((("a" + ("\n\t" + ("a" + "\n\t"))) + "hello") + "") + "\n\t") + "hello"))))) + "") {
            case "":
                var v32 uint16 = 42
                _ = v32
            default:
                v33 := GenS2{F0: (v1 - (v1 + v1))}
                _ = v33
            }
            v34 := []int{}
            if ("\n\t" <= "foo bar") {
                v10 = false
                fmt.Println(v8, v8)
                var v35 uint = 54
                _ = v35
            }
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
