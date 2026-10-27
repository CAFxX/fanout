package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int8; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0() int8 {
    v4 := []int{127, 127}
    var v5 uint8 = 56
    v5 += (^(v5))
    if false {
        if (true || false) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 -= v5
        } else {
            var v6 uint64 = 18446744073709551615
            fmt.Println(v6)
            fmt.Println(v5, v6)
            _ = v6
        }
        v5 -= 2
        if (32767 != 32768) {
            v5 = (v5 % (((v5 * v5)) | 1))
            switch ("a" + "hello") {
            case "\n\t":
                switch "foo bar" {
                case "":
                    switch 128 {
                    case -2147483648:
                        fmt.Println(v5)
                    case 292:
                        var v7 uint = 575
                        _ = v7
                    case -9223372036854775808:
                        var v8 string = ((("foo bar" + "foo bar") + ("" + "hello")) + "foo bar")
                        _ = v8
                    }
                case "foo bar":
                    switch "a" {
                    case "hello":
                        v5 = v5
                    }
                case "a":
                    v9 := GenS1{F0: 2, F1: 42}
                    _ = v9
                }
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "a":
                v5 += 42
            }
            v10 := GenS2{F0: 127, F1: true, F2: 42}
            _ = v10
        } else {
            v5 = v5
        }
    } else {
        var v11 uint64 = 175
        _ = v11
    }
    _ = v4
    _ = v5
    return -20
}
func genFn1(p0 string) int16 {
    v12 := []int{9223372036854775807, -9223372036854775808}
    if true {
        var v13 bool = (true && false)
        fmt.Println(v13)
        var v14 int16 = 1000
        _ = v13
        _ = v14
    } else {
        if (!(631 == 32767)) {
            var v15 int16 = -1
            if len(v12) > 0 {
                v12[0] = -1
            }
            var v16 int64 = 1
            _ = v15
            _ = v16
        }
    }
    switch 569 {
    case 32767:
        var v17 uint8 = 0
        _ = v17
    case -525:
        p0 = "foo bar"
    }
    var v18 int8 = 55
    var v19 int64 = 1
    _ = p0
    _ = v12
    _ = v18
    _ = v19
    return 111
}

func main() {
    var v1 int = 32767
    var v2 string = "hello"
    var v3 uint32 = 2147483648
    v1 += (v1 >> ((561) & 63))
    v2 = (v2 + "\n\t")
    v20 := []int{v1, (v1 * v1), (v1 | v1), (v1 % (((v1 * (v1 * (-(v1))))) | 1))}
    var v21 uint8 = 128
    for fi22 := 0; fi22 < 19; fi22++ {
        v2 = (v2 + v2)
        if (442 > 2147483647) {
            if (0.0 >= 1.5) {
                if (!((v2 + ("\n\t" + v2)) <= "a")) {
                    var v23 uint16 = 255
                    _ = v23
                }
                v24 := GenS3{F0: 42, F1: -11}
                v3 += v3
                _ = v24
            } else {
                for fi25 := 0; fi25 < 9; fi25++ {
                    v3 -= v3
                    _ = fi25
                }
                v21 = 50
                if (!true) {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 uint16 = 0
                        v28 := GenS2{F0: genFn0(), F1: (genFn1(v2) < -32768), F2: 1}
                        var v29 bool = (v3 == (v3 / ((2) | 1)))
                        var v30 float64 = 3.14159
                        if true {
                            var v31 int16 = 1
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    var v32 uint16 = 255
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 uint8 = (v21 >> ((255) & 63))
                fmt.Println(v1, fi33)
                _ = fi33
                _ = v34
            }
        }
        _ = fi22
    }
    var v35 float64 = (1.5 / 1.0)
    var v36 float64 = v35
    v37 := GenS1{F0: 255, F1: 169}
    var v38 int = (-(v1))
    v39 := GenS3{F0: 719, F1: int8(v21)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
